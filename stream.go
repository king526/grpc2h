package grpc2h

import (
	"context"
	"google.golang.org/grpc/metadata"
	"net/http"
)

type serverStream struct {
	r     *http.Request
	w     http.ResponseWriter
	items []interface{}
}

func newServerStream(r *http.Request, w http.ResponseWriter) *serverStream {
	return &serverStream{
		r:     r,
		w:     w,
		items: []interface{}{},
	}
}

func (s *serverStream) SetHeader(m metadata.MD) error {
	header := s.w.Header()
	for k, v := range m {
		header[k] = v
	}
	return nil
}

func (s *serverStream) SendHeader(m metadata.MD) error {
	s.SetHeader(m)
	return nil
}

func (s *serverStream) SetTrailer(metadata.MD) {
	panic("implement me")
}

func (s *serverStream) Context() context.Context {
	return s.r.Context()
}

func (s *serverStream) SendMsg(m interface{}) error {
	s.items = append(s.items, m)
	return nil
}

func (s *serverStream) RecvMsg(m interface{}) error {
	panic("implement me")
}
