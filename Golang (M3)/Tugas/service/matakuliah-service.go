package service

import (
	"app/model"
	"app/repository"
	"errors"
)

func CariMataKuliah(kode string) (model.MataKuliah, error) {
	mataKuliahs := repository.GetMataKuliahs()

	for _, mk := range mataKuliahs {
		if mk.Kode == kode {
			return mk, nil
		}
	}

	return model.MataKuliah{}, errors.New("Data mata kuliah tidak ditemukan.")
}
