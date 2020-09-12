package mahasiswa

import (
	"encoding/json"
	"io"
)

type Mahasiswa struct {
	NPM     string `json:"npm"`
	Nama    string `json:"nama"`
	Jurusan string `json:"jurusan"`
}

type AllMahasiswa []Mahasiswa

func (m *Mahasiswa) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(m)
}

func (m *Mahasiswa) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

func (m *AllMahasiswa) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}
