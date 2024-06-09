package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/yogigoey716/chi-go/config"
	"github.com/yogigoey716/chi-go/server"
)

/*
Function main pada golang hanya memiliki 1 tugas yaitu melakukan operasi secara terus menerus tanpa gangguan (listening and serving).

Sehingga pada function main ini dia hanya melakukan inisialisasi server untuk service endpoint dan koneksi ke database.

Semua operasi database dilakukan pada file tertentu.

Apabila dipetakan sistem code ini maka.

# Dalam pembacaan kode perhatikan parameter dan return type

config -> model -> db (sql -> sqlimpl -> client) -> handler (service -> handler)  -> server (includes handler, db, config) -> main
*/
func main() {
	s := server.New()
	log.Info("Listening on port:", config.GetYamlValues().ServerConfig.Port)
	err := s.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatalf("Listen: %s\n", err)
	}
	log.Info("service stopped")

}
