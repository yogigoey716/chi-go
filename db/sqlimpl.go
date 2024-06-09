package db

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/yogigoey716/chi-go/model"
)

var (
	bl     model.Blogs
	record []model.Blogs
)

func createBlog(bl *model.Blogs) (model.BlogData, error) {
	db := GetDBConnection()
	if err := db.Table("chigo").Create(&bl).Error; err != nil {
		log.Info("failure", model.Blogs{}, err)
	}
	record := model.BlogData{
		Blog:    *bl,
		Message: "data saved",
	}
	return record, nil

}

func getAllBlog(id string) (model.Blogs, error) {
	db := GetDBConnection()
	if err := db.Table("chigo").Where("id=?", id).Find(&record).Error; err != nil {
		log.Info("failure", []model.Blogs{})
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Blogs{}, fmt.Errorf("blog with IDs %s not found", id)
		}
		return model.Blogs{}, fmt.Errorf("failed to get blog: %w", err)
	}
	if len(record) == 0 {
		return model.Blogs{}, gorm.ErrRecordNotFound
	}
	return record[0], nil
}

func getAvailableBlog() ([]model.Blogs, error) {
	db := GetDBConnection()

	if result := db.Table("chigo").Find(&record); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return []model.Blogs{}, fmt.Errorf("no blogs found")
		}
		return []model.Blogs{}, fmt.Errorf("failed to get blogs: %w", result.Error)
	}

	return record, nil
}

func updateBlog(id string, bl *model.Blogs) (model.BlogData, error) {

	db := GetDBConnection()
	if err := db.Table("chigo").Where("id=?", id).Updates(&bl).Error; err != nil {
		log.Info("failure", []model.Blogs{})
	}
	record := model.BlogData{
		Blog:    *bl,
		Message: "record updated successfully",
	}
	return record, nil
}

func deleteBlog(id string) (string, error) {
	db := GetDBConnection()
	if err := db.Table("chigo").Where("id=?", id).Delete(&bl).Error; err != nil {
		log.Info("failure", []model.Blogs{})
		return "not able to delete", err
	}

	return "deleted successfully", nil
}
