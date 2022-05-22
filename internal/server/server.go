package server

import (
	"net"

	"github.com/valyala/fasthttp"
)

type (
	Server struct {
		Server *fasthttp.Server
	}
)

func NewServer(handler fasthttp.RequestHandler, name string) (s *Server) {
	s = &Server{
		Server: &fasthttp.Server{
			Name:            name,
			Handler:         handler,
			CloseOnShutdown: true,
		},
	}

	return
}

func (s *Server) Serve(l net.Listener) (err error) {
	err = s.Server.Serve(l)
	return
}

func (s *Server) Close() (err error) {
	err = s.Server.Shutdown()
	return
}
