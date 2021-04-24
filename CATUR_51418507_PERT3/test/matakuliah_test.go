package test

import (
	"catur_51418507_pert3/model" //sesuaikan dengan nama folder (case sensitive)
	"testing"
)

func TestMatakuliah(t *testing.T) {
	var dataInsertMtk = []model.Matakuliah{
		model.Matakuliah{
			Id:          "1",
			Nama_matkul: "SBD",
			Jurusan:     "FTI",
		},
		model.Matakuliah{
			Id:          "2",
			Nama_matkul: "TBO",
			Jurusan:     "FTI",
		},
		model.Matakuliah{
			Id:          "3",
			Nama_matkul: "Grafkom2",
			Jurusan:     "FTI",
		},
	}
	db, err := initDatabase()
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Testing insert matakuliah", func(t *testing.T) {
		for _, dataInsert := range dataInsertMtk {
			err := dataInsert.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
		}
	})
	t.Run("Testing update matakuliah", func(t *testing.T) {
		var updateData = map[string]interface{}{
			"nama_matkul": "SBD"}
		data := dataInsertMtk[0]
		if err := data.Update(db, updateData); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Testing Get matakuliah", func(t *testing.T) {
		_, err := model.GetMatakuliah(db, "1")
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Testing Get matakuliah", func(t *testing.T) {
		_, err := model.GetAllMatakuliah(db)
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Testing delete matakuliah", func(t *testing.T) {
		data := dataInsertMtk[0]
		if err := data.Delete(db); err != nil {
			t.Fatal(err)
		}
	})
	defer db.Close()
}
