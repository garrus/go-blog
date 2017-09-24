package core

import (
	"net/http"
	"fmt"
)

type HttpServer struct {
	Router Router
}

func (hs *HttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	context := NewContext(r, w)

	defer func(context *Context){
		if err := recover(); err != nil {
			fmt.Errorf("[%s] %s %s %s 500\n",
				context.startTime.Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, r.RequestURI)
			context.SetHeader("Content-Type", "text/plain")
			context.Send("500 Internal Service Error", 500)
		}
		context.End()
	}(context)

	if handler, ok := hs.Router.Route(r.Method, r.RequestURI); ok {
		handler(context)
		fmt.Printf("[%s] %s %s %s %d %f %d\n",
			context.startTime.Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, r.RequestURI, context.StatusCode(), context.ElapsedTime().Seconds(), context.BodyLen())
	} else {
		fmt.Printf("[%s] %s %s %s 404\n",
			context.startTime.Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, r.RequestURI)
		context.SetHeader("Content-Type", "text/plain")
		context.Send("Invalid Request URI", 404)
	}
}