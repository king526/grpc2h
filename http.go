package grpc2h

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"google.golang.org/grpc"
)

var (
	typeContext      = reflect.TypeOf((*context.Context)(nil)).Elem()
	typeServerStream = reflect.TypeOf((*grpc.ServerStream)(nil)).Elem()
)

type Handler struct {
	streamImpl []reflect.Type
	rt         reflect.Type
	rv         reflect.Value
}

type HandlerBuilder struct {
	h Handler
}

func (b *HandlerBuilder) SetGRPCServer(s interface{}) {
	b.h.rt = reflect.TypeOf(s)
	b.h.rv = reflect.ValueOf(s)
}

func (b *HandlerBuilder) RegisterSeverStream(v grpc.ServerStream) {
	rt := reflect.TypeOf(v)
	for rt.Kind() != reflect.Struct {
		rt = rt.Elem()
	}
	b.h.streamImpl = append(b.h.streamImpl, rt)
}

func (b *HandlerBuilder) Build() (*Handler, error) {
	for i := 0; i < b.h.rt.NumMethod(); i++ {
		m := b.h.rv.Method(i)
		// fmt.Println(m.Type())
		for j := 0; j < m.Type().NumIn(); j++ {
			arg := m.Type().In(j)
			if !arg.Implements(typeServerStream) {
				continue
			}
			if b.h.getServerStramArgType(arg) == nil {
				return nil, fmt.Errorf("unregistered server stream:%v", arg)
			}
		}
	}
	return &b.h, nil
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, ok := h.rt.MethodByName(r.URL.Path[1:]); !ok {
		return
	}
	ss := newServerStream(r, w)
	f := h.rv.MethodByName(r.URL.Path[1:])
	ft := f.Type()
	var params []reflect.Value
	for i := 0; i < ft.NumIn(); i++ {
		arg := ft.In(i)
		if arg == typeContext {
			params = append(params, reflect.ValueOf(r.Context()))
			continue
		}
		if arg.Implements(typeServerStream) {
			typ := h.getServerStramArgType(arg)
			pa := reflect.New(typ)
			setFieldValueByType(pa, typeServerStream, ss)
			params = append(params, pa)
			continue
		}
		p := reflect.New(arg.Elem())
		if err := json.NewDecoder(r.Body).Decode(p.Interface()); err != nil {
			h.writeResp(ss, []reflect.Value{reflect.ValueOf(err)})
			return
		}
		params = append(params, p)
	}
	resp := f.Call(params)
	h.writeResp(ss, resp)
}

func (h *Handler) getServerStramArgType(arg reflect.Type) reflect.Type {
	for _, t := range h.streamImpl {
		if reflect.New(t).Type().Implements(arg) {
			return t
		}
	}
	return nil
}

func (h *Handler) writeResp(ss *serverStream, resp []reflect.Value) {
	var hasBody bool
	err := resp[0]
	if len(resp) == 2 {
		hasBody = true
		err = resp[1]
	}
	if !err.IsNil() {
		ss.w.Write([]byte(fmt.Sprint(err)))
		return
	}
	if hasBody {
		ss.w.Write([]byte(fmt.Sprint(resp[0])))
		return
	}
	bs, _ := json.Marshal(ss.items)
	ss.w.Write(bs)
}
