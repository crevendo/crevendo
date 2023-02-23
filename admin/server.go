package admin

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"os"
	"strconv"
)

type Response = any
type TemplateController = func() (Response, error)

type Server struct {
	echoServer *echo.Echo
}

func NewServer(options ...func(*Server)) *Server {
	fmt.Println("Serving admin...")
	e := echo.New()
	srv := &Server{echoServer: e}
	InitDB()
	srv.setEchoOptions()
	for _, o := range options {
		o(srv)
	}
	return srv
}

func WithEndpoint(path string, fn func(ctx Context) error) func(*Server) {
	return func(s *Server) {
		s.echoServer.GET(path, adaptHandler(fn))
	}
}

func adaptHandler(fn func(ctx Context) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		return fn(Context{c})
	}
}

func (s Server) setEchoOptions() {
	wd, _ := os.Getwd()
	s.echoServer.Debug = true
	s.echoServer.Static("/assets", wd+"/admin/dist/assets")
	s.echoServer.File("/", wd+"/admin/dist/index.html")
	s.echoServer.POST("/api/options", saveOptionsEndpoint)
	s.echoServer.GET("/api/options", getOptionsEndpoint)
}

func (s Server) Listen(port int64) {
	go func() {
		s.echoServer.Logger.Fatal(s.echoServer.Start(":" + strconv.FormatInt(port, 10)))
	}()
}
