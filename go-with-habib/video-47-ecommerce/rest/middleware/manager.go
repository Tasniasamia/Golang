package middleware

import (
	"fmt"
	"net/http"
);

type middleware func(http.Handler) http.Handler;

type Manager struct {
	globalMiddleware []middleware;
}

func NewManager() *Manager {
	return &Manager{
		globalMiddleware: []middleware{},
	}
}

func (m *Manager) Use(md ...middleware) {
	m.globalMiddleware = append(m.globalMiddleware, md...)
}

func (m *Manager) With(h http.Handler, md ...middleware) http.Handler {
    next:=h;
      for _,md:=range md{
	    fmt.Println(md)
	    middleware:= md;
		next=middleware(next);

	}
	 



  return next;
	


}

func (m *Manager) WrappedMux(h http.Handler) http.Handler {
    next:=h;
   for _,md:=range m.globalMiddleware {
		next = md(next)
	}
   return next;
}
	
   
