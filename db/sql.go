package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
	"github.com/yogigoey716/chi-go/config"
)
/*
	Pada file ini dilakukan penghubungan koneksi dengan database. 
	Terdapat banyak library yang bisa digunakan untuk koneksi ini, pada kali ini kita menggunakan GORM untuk postgresSQL.
	GORM memberikan method yang mudah untuk melakukan operasi database kali ini.

	Pada kode ini terdapat banyak pointer(tanda *) yang digunakan. Hal ini bertujuan untuk mengakses memory untuk kode koneksi database.
	Hal ini bertujuan untuk tidak melakukan pembuatan variable baru (jadi 2) apabila melakukan inisialisasi New().
	Yang bertujuan untuk tidak melakukan penambahan memory (effisiensi pembuatan variable cost memory (RAM)).
	Apabila dibandingkan dengan bahasa pemrograman lain (Java, Kotlin, Python, C++ (OOP)) hal yang dilakukan ini adalah Singleton.
	
	Apabila ingin menggantikan database ke database yang lainnya kalian dapat mengubah file ini saja.
	 
*/

type sqlConn struct {
	DbPool *gorm.DB
}

var connector *sqlConn

func InitMysql() *sqlConn {
	if connector != nil {
		log.Info("DataBase is initialized")
		return connector
	}
	log.Info("DataBase was not initialized ..initializing again")
	var err error
	connector, err = initDB()
	if err != nil {
		panic(err)
	}
	return connector
}

func initDB() (*sqlConn, error) {
	log.Info(config.GetYamlValues().DBConfig, config.GetYamlValues().DBConfig.Port)
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		config.GetYamlValues().DBConfig.Server, config.GetYamlValues().DBConfig.Username, config.GetYamlValues().DBConfig.Schema, config.GetYamlValues().DBConfig.Password) //Build connection string

	db, err := gorm.Open("postgres", dbUri)
	if err != nil {
		panic(err)
	}
	if maxCons := config.GetYamlValues().DBConfig.MaxConnection; maxCons > 0 {
		db.DB().SetMaxOpenConns(maxCons)
		db.DB().SetMaxIdleConns(maxCons / 3)
	}
	return &sqlConn{db}, nil
}

func GetDBConnection() *gorm.DB {
	return connector.DbPool
}
