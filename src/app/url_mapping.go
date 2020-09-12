package app

import (
	"net/http"

	"github.com/dzikrisyafi/RESTfulAPI-golang/src/controllers"
)

var (
	router = http.NewServeMux()
)

func mapUrls() {
	router.Handle("/api/mahasiswa/", controllers.NewMahasiswaHandler())
}
