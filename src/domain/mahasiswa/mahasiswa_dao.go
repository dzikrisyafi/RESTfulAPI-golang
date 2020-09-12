package mahasiswa

import (
	"github.com/dzikrisyafi/RESTfulAPI-golang/src/datasources/mysql/db"
)

const (
	queryGetAllMahasiswa = "SELECT * FROM mahasiswa;"
	queryGetMahasiswa    = "SELECT nama, jurusan FROM mahasiswa WHERE npm=?;"
	queryInsertMahasiswa = "INSERT INTO mahasiswa(npm, nama, jurusan) VALUES(?, ?, ?);"
	queryUpdateMahasiswa = "UPDATE mahasiswa SET nama=?, jurusan=? WHERE npm=?;"
	queryDeleteMahasiswa = "DELETE FROM mahasiswa WHERE npm=?;"
)

func (m *Mahasiswa) GetAllMahasiswa() (AllMahasiswa, error) {
	stmt, err := db.DbConn().Prepare(queryGetAllMahasiswa)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]Mahasiswa, 0)
	for rows.Next() {
		if err := rows.Scan(&m.NPM, &m.Nama, &m.Jurusan); err != nil {
			return nil, err
		}

		result = append(result, *m)
	}

	return result, nil
}

func (m *Mahasiswa) GetMahasiswa() error {
	stmt, err := db.DbConn().Prepare(queryGetMahasiswa)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result := stmt.QueryRow(m.NPM)
	if err = result.Scan(&m.Nama, &m.Jurusan); err != nil {
		return err
	}

	return nil
}

func (m *Mahasiswa) CreateMahasiswa() error {
	stmt, err := db.DbConn().Prepare(queryInsertMahasiswa)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(m.NPM, m.Nama, m.Jurusan); err != nil {
		return err
	}

	return nil
}

func (m *Mahasiswa) UpdateMahasiswa() error {
	stmt, err := db.DbConn().Prepare(queryUpdateMahasiswa)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(m.Nama, m.Jurusan, m.NPM); err != nil {
		return err
	}

	return nil
}

func (m *Mahasiswa) DeleteMahasiswa() error {
	stmt, err := db.DbConn().Prepare(queryDeleteMahasiswa)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(m.NPM); err != nil {
		return err
	}

	return nil
}
