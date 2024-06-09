package server

import (
	"fmt"

	"github.com/yogigoey716/chi-go/config"
	"github.com/yogigoey716/chi-go/db"
	"github.com/yogigoey716/chi-go/handler"

	"net"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

/*
	File ini bertujuan hanya untuk menjalankan (listen and server) pada function main().
	Kode ini berisikan Inisialisasi dari beberapa kode sebelumnya.
*/

type Server struct {
	httpServer *http.Server
	router     *chi.Mux
}

/*
Kode ini yang akan digunakan pada main.

Bila dilihat terdapat operasi dari sql.go yaitu inimysql()
dari client.go yaitu NewClient() yang berisikan method dari interface SqlClient
dan dimasukan ke parameter NewService dari service.go
*/
func New() *Server {
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))
	db.InitMysql()

	mysqlFGClient := db.NewClient(
		&db.Config{
			DBConnection: "",
		})

	run := handler.NewService(mysqlFGClient)
	setupRoutesForUpdate(run, r)

	server := newServer(r)

	return server
}

/*
Router untuk endpoint (/api) dengan service yang kita buat pada /v1/...
Sehingga endpoint akan berbentuk /api/v1/{METHOD CRUD PADA SERVICE}

Kode ini dapat membantu apabila terdapat service baru
*/
func setupRoutesForUpdate(service handler.Service, r *chi.Mux) {

	r.Route("/api", func(r chi.Router) {
		r.Mount("/v1", handler.Handler(service))
	})
}

func (s *Server) ListenAndServe() error {
	l, err := net.Listen("tcp", ":"+s.httpServer.Addr)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	return s.httpServer.Serve(l)
}

func newServer(r *chi.Mux) *Server {
	fmt.Println("****Server Started on", config.GetYamlValues().ServerConfig.Port, "****")
	return &Server{
		httpServer: &http.Server{Addr: config.GetYamlValues().ServerConfig.Port, Handler: r},
		router:     r,
	}
}
