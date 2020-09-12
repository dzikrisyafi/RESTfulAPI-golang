package controllers

import (
	"net/http"
	"regexp"

	"github.com/dzikrisyafi/RESTfulAPI-golang/src/domain/mahasiswa"
)

type mahasiswaHandler struct {
}

type MahasiswaHandler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
	Get(string, http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Update(string, http.ResponseWriter, *http.Request)
	Delete(string, http.ResponseWriter, *http.Request)
}

func NewMahasiswaHandler() MahasiswaHandler {
	return &mahasiswaHandler{}
}

func GetNPM(r *http.Request) string {
	reg := regexp.MustCompile(`([a-z0-9]+)`)
	g := reg.FindAllString(r.URL.Path, -1)
	return g[len(g)-1]
}

func (h *mahasiswaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		npm := GetNPM(r)
		h.Get(npm, w, r)
		return
	case http.MethodPost:
		h.Create(w, r)
		return
	case http.MethodPut:
		npm := GetNPM(r)
		h.Update(npm, w, r)
		return
	case http.MethodDelete:
		npm := GetNPM(r)
		h.Delete(npm, w, r)
		return
	default:
		http.Error(w, "Method tidak tersedia", http.StatusBadRequest)
	}
}

func (h *mahasiswaHandler) Get(npm string, w http.ResponseWriter, r *http.Request) {
	mhs := &mahasiswa.Mahasiswa{NPM: npm}
	if len(npm) == 8 {
		if err := mhs.GetMahasiswa(); err != nil {
			http.Error(w, "error when trying to get mahasiswa", http.StatusInternalServerError)
			return
		}

		err := mhs.ToJSON(w)
		if err != nil {
			http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
			return
		}
	} else {
		res, err := mhs.GetAllMahasiswa()
		if err != nil {
			http.Error(w, "error when trying to get mahasiswa", http.StatusInternalServerError)
			return
		}

		err = res.ToJSON(w)
		if err != nil {
			http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
			return
		}
	}
}

func (h *mahasiswaHandler) Create(w http.ResponseWriter, r *http.Request) {
	mhs := &mahasiswa.Mahasiswa{}

	if err := mhs.FromJSON(r.Body); err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusInternalServerError)
		return
	}

	if err := mhs.CreateMahasiswa(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`{"message": "Success to created mahasiswa"}`))
}

func (h *mahasiswaHandler) Update(npm string, w http.ResponseWriter, r *http.Request) {
	mhs := &mahasiswa.Mahasiswa{NPM: npm}
	if err := mhs.FromJSON(r.Body); err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusInternalServerError)
		return
	}

	if err := mhs.UpdateMahasiswa(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`{"message": "Success to updated mahasiswa"}`))
}

func (h *mahasiswaHandler) Delete(npm string, w http.ResponseWriter, r *http.Request) {
	mhs := &mahasiswa.Mahasiswa{NPM: npm}
	if err := mhs.DeleteMahasiswa(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`{"message": "Success to deleted mahasiswa"}`))
}
