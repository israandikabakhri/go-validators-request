// helpers/validator.go

package penduduk

import (
	"errors"
	"regexp"
	"time"
)

// Validasi NIK, nama, dan alamat
func ValidateData(nik, nama, tgl_lahir, waktu_kunjungan string) error {

	if nik == "" {
		return errors.New("NIK tidak boleh kosong")
	}

	if nama == "" {
		return errors.New("nama tidak boleh kosong")
	}

	if tgl_lahir == "" {
		return errors.New("tanggal Lahir tidak boleh kosong")
	}

	if waktu_kunjungan == "" {
		return errors.New("waktu kunjungan tidak boleh kosong")
	}

	// Validasi NIK
	if match, _ := regexp.MatchString("^[0-9]{16}$", nik); !match {
		return errors.New("NIK tidak valid, haruns mengandung 16 karakter nomor")
	}

	// Validasi nama
	if len(nama) < 3 {
		return errors.New("nama terlalu pendek")
	}

	// Validasi tgl_lahir
	if match, _ := regexp.MatchString("^[0-9]{4}-[0-9]{2}-[0-9]{2}$", tgl_lahir); !match {
		return errors.New("format tanggal salah, harus dalam format 'YYYY-MM-DD'")
	}

	// validasi waktu_kunjungan
	layout := "2006-01-02 15:04:05"
	_, err := time.Parse(layout, waktu_kunjungan)
	if err != nil {
		return errors.New("format waktu_kunjungan harus yyyy-mm-dd H:i:s")
	}

	return nil
}
