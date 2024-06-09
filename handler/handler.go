package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"github.com/yogigoey716/chi-go/httphandler"
	"github.com/yogigoey716/chi-go/model"
)

/*

	Pada file ini berisi routing dan method yang akan disajikan kepada pengguna api.

	Routing pada file ini mneggunakan library chi untuk membuat routernya.

	Pada file ini juga dimana untuk menentukan method CRUD REST-API.
	Sehingga beberapa function disini menggunakan net/http untuk menentukan body dari endpoint (untuk beberapa operasi seperti POST DELETE)


	File ini akan digunakan pada server.go

*/

/*

	Helper struct untuk menentukan service yang digunakan dan function dari function yang tersedia (hanya menarohkan nama function)

	Kode dibawah tidak termasuk kode yang terbaik karena setiap function tidak mengembalikan return error.
	Mengembalikan return error merupakan best practice di golang


	type ctx struct {
	store Service
	h     func(Service, http.ResponseWriter, *http.Request) error
}
*/

type ctx struct {
	store Service
	h     func(Service, http.ResponseWriter, *http.Request)
}

/*
Function biasa untuk mengembalikan HandlerFunc yang akan digunakan pada chi (helper)
*/
func (g *ctx) handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		g.h(g.store, w, r)
	}
}

/*
Disini kita menggunakan Chi untuk membuat Routing.

Disini juga kita dapat menambahkan middleware yang disediakan oleh chi. (Logger, dll)
*/
func Handler(store Service) http.Handler {
	r := chi.NewRouter()
	getRecordSetPost := ctx{store: store, h: getRecordSetPost}
	createBlogs := ctx{store: store, h: createBlogs}
	updateBlogs := ctx{store: store, h: updateBlogs}
	deleteBlogs := ctx{store: store, h: deleteBlogs}
	getAll := ctx{store: store, h: getAllBlogs}

	r.Get(httphandler.WrapHandlerFunc("/blog/{data}", "get blog", getRecordSetPost.handle()))
	r.Get(httphandler.WrapHandlerFunc("/blog", "get all blog", getAll.handle()))
	r.Post(httphandler.WrapHandlerFunc("/blog", "create blog", createBlogs.handle()))
	r.Put(httphandler.WrapHandlerFunc("/blog/{data}", "update blog", updateBlogs.handle()))
	r.Delete(httphandler.WrapHandlerFunc("/blog/remove/{data}", "delete blog", deleteBlogs.handle()))

	return r
}

/*
Pada semua function dibawah ini mengimplementasikan method dari file Service.go untuk melakukan operasi database dari client.go

Kali ini kita menggunakan library net/http dan chi untuk menulis response dan menerima request (payload) user (istilah ini bisa kita sebut Writing API).
parameter w dan r merupakan hal yang wajib dalam menggunakan library net/http (bawaan)

Terdapat banyak library yang dapat digunakan salah satunya Gin yang dapat membantu dalam proses Writing dan Routing.

Pada Kode ini tersedia struct Requset yang akan berisikan data.

Data tersebut akan dimunculkan dengan render.Render dari Chi yang berisikan handler untuk response sukses (NewSuccessResponse dari httphandler folder)
diikutsertakan dengan render.Status.
*/
func createBlogs(store Service, w http.ResponseWriter, r *http.Request) {

	data := &Request{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, httphandler.ErrInvalidRequest(err, "Invalid Request"))
		return
	}
	recordSchema := data.Blogs
	services, err := store.CreateRecordCoreTeam(recordSchema)
	if err != nil {
		log.Errorf("Unable To Fetch stats ", httphandler.Error(err).Code, services, err)
		httphandler.ErrInvalidRequest(err, "Unable To Fetch Services ")
		return
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, httphandler.NewSuccessResponse(http.StatusOK, services))
}

func getRecordSetPost(store Service, w http.ResponseWriter, r *http.Request) {
	data := chi.URLParam(r, "data")
	services, err := store.GetRecordSetPost(data)
	if err != nil {
		log.Errorf("Unable To Fetch stats ", httphandler.Error(err).Code, services, err)
		httphandler.ErrInvalidRequest(err, "Unable To Fetch Services ")
		return
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, httphandler.NewSuccessResponse(http.StatusOK, services))
}

func updateBlogs(store Service, w http.ResponseWriter, r *http.Request) {
	data := chi.URLParam(r, "data")
	payload := &Request{}
	if err := render.Bind(r, payload); err != nil {
		render.Render(w, r, httphandler.ErrInvalidRequest(err, "Invalid Request"))
		return
	}
	recordSchema := payload.Blogs
	services, err := store.UpdateBlog(data, recordSchema)
	if err != nil {
		log.Errorf("Unable To Fetch stats ", httphandler.Error(err).Code, services, err)
		httphandler.ErrInvalidRequest(err, "Unable To Fetch Services ")
		return
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, httphandler.NewSuccessResponse(http.StatusOK, services))
}

func deleteBlogs(store Service, w http.ResponseWriter, r *http.Request) {
	data := chi.URLParam(r, "data")
	services, err := store.DeleteBlogs(data)
	if err != nil {
		log.Errorf("Unable To Fetch stats ", httphandler.Error(err).Code, services, err)
		httphandler.ErrInvalidRequest(err, "Unable To Fetch Services ")
		return
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, httphandler.NewSuccessResponse(http.StatusOK, services))
}

func getAllBlogs(store Service, w http.ResponseWriter, r *http.Request) {
	services, err := store.GetAllBlogs()
	if err != nil {
		log.Errorf("Unable To Fetch stats ", httphandler.Error(err).Code, services, err)
		httphandler.ErrInvalidRequest(err, "Unable To Fetch Services ")
		return
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, httphandler.NewSuccessResponse(http.StatusOK, services))
}

type Request struct {
	*model.Blogs
}

func (a *Request) Bind(r *http.Request) error {
	//TODO: to be expanded
	return nil
}

type Response struct {
	Meta interface{}
	Data interface{}
}
