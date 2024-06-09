package handler

import (
	log "github.com/sirupsen/logrus"
	"github.com/yogigoey716/chi-go/db"
	"github.com/yogigoey716/chi-go/model"
)

/*
Pada bagian ini karena service yang disediakan hanya 1 yaitu maka kita dapat membuat interface yang sama
sepert pada client.go.
Pada file ini berisi implementasi dari method yang disediakan oleh intercface SqlClient dari client.go.

Apabila terdapat pada service lain maka best practicenya file service dipisah untuk kemudahan development dan bug fix (SoC).
File Sevice ini akan digunakan pada file handler.go.
*/

type Service interface {
	CreateRecordCoreTeam(bl *model.Blogs) (model.BlogData, error)
	GetRecordSetPost(id string) (model.Blogs, error)
	UpdateBlog(id string, bl *model.Blogs) (model.BlogData, error)
	DeleteBlogs(id string) (string, error)
	GetAllBlogs() ([]model.Blogs, error)
}

func NewService(sqlDB db.SqlClient) Service {
	return &service{
		sqlDB: sqlDB,
	}
}

type service struct {
	sqlDB db.SqlClient
}

func (s *service) CreateRecordCoreTeam(bl *model.Blogs) (model.BlogData, error) {
	record, err := s.sqlDB.CreateData(bl)
	if err != nil {
		log.Info("Failure: mot getting data from Table", err)
		return model.BlogData{}, err
	}
	return record, nil
}

func (s *service) GetRecordSetPost(id string) (model.Blogs, error) {
	record, err := s.sqlDB.GetData(id)
	if err != nil {
		log.Info("Failure: not getting data from Table", err)
		return model.Blogs{}, err
	}
	return record, nil
}

func (s *service) UpdateBlog(id string, bl *model.Blogs) (model.BlogData, error) {
	record, err := s.sqlDB.UpdateData(id, bl)
	if err != nil {
		log.Info("Failure: mot getting data from Table", err)
		return model.BlogData{}, err
	}
	return record, nil
}

func (s *service) DeleteBlogs(id string) (string, error) {
	record, err := s.sqlDB.DeleteData(id)
	if err != nil {
		log.Info("Failure: not getting data from Table", err)
		return "record is not available in the system", err
	}
	return record, nil

}

func (s *service) GetAllBlogs() ([]model.Blogs, error) {
	var record []model.Blogs
	record, err := s.sqlDB.GetAllBlog()
	if err != nil {
		log.Info("Failure: not getting data from Table", err)
		return nil, err
	}
	return record, nil
}
