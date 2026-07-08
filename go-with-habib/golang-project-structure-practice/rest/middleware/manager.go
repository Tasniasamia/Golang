package middleware

import "net/http"

type middleware func(http.Handler) http.Handler;

type Manager struct{
	GlobalMiddlewares [] middleware;
}


func NewManager()*Manager{
	return &Manager{}
}



func(m *Manager)Use(middlewares...middleware) {

m.GlobalMiddlewares = append(m.GlobalMiddlewares, middlewares...)

}





func(m *Manager)With(next http.Handler,middlewares...middleware) http.Handler{


for _,mid:=range m.GlobalMiddlewares{
	next=mid(next);
}

for _,mid:=range middlewares{
next=mid(next);
}

return next;

}
