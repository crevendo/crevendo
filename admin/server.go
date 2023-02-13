package admin

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

type Response = any
type TemplateController = func() (Response, error)

type Server struct {
	echoServer *echo.Echo
}

func NewServer(options ...func(*Server)) *Server {
	e := echo.New()
	srv := &Server{echoServer: e}
	srv.setEchoOptions()
	for _, o := range options {
		o(srv)
	}
	return srv
}

func (s Server) setEchoOptions() {
	wd, _ := os.Getwd()
	s.echoServer.Debug = true
	s.echoServer.Static("/assets", wd+"/admin/dist/assets")
	s.setIndex(wd + "/admin/dist/index.html")
}

func WithRoute(route string, controller ...TemplateController) func(*Server) {
	return func(s *Server) {
		s.echoServer.GET(route, func(c echo.Context) error {
			status := http.StatusOK
			data, err := controller[0]()
			if err != nil {
				s.echoServer.Logger.Error(err)
				return err
			}

			jsonString, err := json.Marshal(&data)
			if err != nil {
				return err
			}
			return c.Render(status, "index.html", string(jsonString))
		})
	}
}

func (s Server) setIndex(index string) {
	t := &Template{
		templates: template.Must(template.ParseGlob(index)),
	}
	s.echoServer.Renderer = t
}

func (s Server) Listen(port int64) {
	go func() {
		s.echoServer.Logger.Fatal(s.echoServer.Start(":" + strconv.FormatInt(port, 10)))
	}()
}
