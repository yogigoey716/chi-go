package db

import (
	"github.com/yogigoey716/chi-go/model"
)

/*
Pada file ini berisi perilaku atau behaviour (interface) dari implementasi operasi database yang dilaukan di file sqlimpl.go

Function yang harus diperlihat dengan jelas yaitu pada function NewClient().
Function tersebut akan digunakan untuk inisialisasi pada file server.go.
Function tersebut mengembalikan struct sqlClient yang mempunyai interface SqlClient (perhatikan kapital) yang berisi method untuk operasi database.
*/
type SqlClient interface {
	GetAllBlog() ([]model.Blogs, error)
	CreateData(bl *model.Blogs) (model.BlogData, error)
	GetData(string) (model.Blogs, error)
	UpdateData(id string, bl *model.Blogs) (model.BlogData, error)
	DeleteData(id string) (string, error)
}

func NewClient(config *Config) SqlClient {
	return &sqlClient{
		config: config,
	}
}

type Config struct {
	DBConnection string
}

type sqlClient struct {
	config *Config
}

/*
Method menggunakan pointer untuk melakukan referensi ke struct sqlClient yang mempunyai interface,
dan akan digunakan pada NewClient().

return createBlog getAllBlog updateBlog tersebut merupakan function dari sqlimpl.go yang berisi operasi postgresSQL

method disini akan memliki return type yang sama seperti function dari sqlimpl.go
*/

func (c *sqlClient) CreateData(bl *model.Blogs) (model.BlogData, error) {
	return createBlog(bl)
}

func (c *sqlClient) GetData(id string) (model.Blogs, error) {
	return getAllBlog(id)
}

func (c *sqlClient) UpdateData(id string, bl *model.Blogs) (model.BlogData, error) {
	return updateBlog(id, bl)
}

func (c *sqlClient) DeleteData(id string) (string, error) {
	return deleteBlog(id)
}

func (c *sqlClient) GetAllBlog() ([]model.Blogs, error) {
	return getAvailableBlog()
}
