package core

import (
	"net/http"
	"time"
	"bytes"
	"io"
)

type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	startTime      time.Time
	bodyBuffer	   bytes.Buffer
	statusCode int
	bodyLen int
	isSent bool
	sentTime time.Time
}


func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	c := &Context{}
	c.request = r
	c.responseWriter = w
	c.startTime = time.Now()
	c.statusCode = http.StatusOK
	return c
}

func (c Context) StartTime() time.Time{
	return c.startTime
}

func (c Context) ElapsedTime() time.Duration {
	return time.Now().Sub(c.startTime)
}

func (c Context) Request() *http.Request {
	return c.request
}

func (c *Context) SetStatusCode(code int) {
	c.statusCode = code
}

func (c Context) StatusCode() int {
	return c.statusCode
}

func (c *Context) SetHeader(name string, val string) {
	c.responseWriter.Header().Add(name, val)
}

func (c *Context) AppendBody(body string){
	c.bodyBuffer.WriteString(body)
	c.bodyLen = c.bodyBuffer.Len()
}

func (c Context) BodyLen() int{
	return c.bodyLen
}

func (c *Context) End(){
	if !c.isSent {
		c.doSent()
	}
}

func (c *Context) Send(body string, statusCode int){

	if len(body) > 0 {
		c.AppendBody(body)
	}

	if statusCode > 0 {
		c.SetStatusCode(statusCode)
	}

	c.doSent()
}

func (c *Context) doSent(){

	if c.statusCode != http.StatusOK {
		c.responseWriter.WriteHeader(c.statusCode)
	}
	if c.responseWriter.Header().Get("Content-Type") == "" {
		c.responseWriter.Header().Add("Content-Type", "text/html")
	}

	io.WriteString(c.responseWriter, c.bodyBuffer.String())
	c.bodyBuffer.Truncate(0)

	c.isSent = true
	c.sentTime = time.Now()
}

