package model

import (
	"database/sql"
	"fmt"
	"strings"
)

var TabelMatakuliah string = `
	CREATE TABLE matakuliah(
		id VARCHAR(10) PRIMARY KEY,
		nama_matkul VARCHAR(30),
		jurusan VARCHAR(10)
	);
`

type Matakuliah struct {
	Id          string `json:"Id"`
	Nama_matkul string `json:"Nama_matkul"`
	Jurusan     string `json:"Jurusan"`
}

func (m *Matakuliah) Fields() ([]string, []interface{}) {
	fields := []string{"id", "nama_matkul", "jurusan"}
	temp := []interface{}{&m.Id, &m.Nama_matkul, &m.Jurusan}
	return fields, temp
}

func (m *Matakuliah) Structur() *Matakuliah {
	return &Matakuliah{}
}

func (m *Matakuliah) Insert(db *sql.DB) error {
	query := fmt.Sprintf("INSERT INTO %v values(?,?,?)", "matakuliah")
	_, err := db.Exec(query, &m.Id, &m.Nama_matkul, &m.Jurusan)
	return err
}

func (m *Matakuliah) Update(db *sql.DB, data map[string]interface{}) error {
	var kolom = []string{}
	var args []interface{}
	i := 1
	// Ini loop data untuk dimasukan kedalam set
	for key, value := range data {
		updateData := fmt.Sprintf("%v = ?", strings.ToLower(key))
		kolom = append(kolom, updateData)
		args = append(args, value)
		i++
	}
	// Ubah array menjadi string dengan pemisah koma
	dataUpdate := strings.Join(kolom, ",")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s = ?", "matakuliah", dataUpdate, "id")
	args = append(args, m.Id)
	// Exec dengan query yang ada
	_, err := db.Exec(query, args...)
	return err
}

func (m *Matakuliah) Delete(db *sql.DB) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", "matakuliah", "id")
	// Exec dengan query yang ada
	_, err := db.Exec(query, m.Id)
	return err
}

func GetMatakuliah(db *sql.DB, id string) (*Matakuliah, error) {
	m := &Matakuliah{}
	each := m.Structur()
	_, dst := each.Fields()
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ?", "matakuliah", "id")
	// isinya akan dimasukan kedalam var dst yang dideklarasikan diatas
	err := db.QueryRow(query, id).Scan(dst...)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return each, nil
}

func GetAllMatakuliah(db *sql.DB) ([]*Matakuliah, error) {
	m := &Matakuliah{}
	query := fmt.Sprintf("SELECT * FROM %s", "matakuliah")
	data, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer data.Close()
	var result []*Matakuliah
	for data.Next() {
		each := m.Structur()
		_, dst := each.Fields()
		err := data.Scan(dst...)
		if err != nil {
			return nil, err
		}
		fmt.Println(each)
		result = append(result, each)
	}
	return result, nil
}
