# go-validators-request
Project Go yang diperuntukkan untuk membuat sistem validasi request oleh client dan mengembalikan error tersebut ke Client untuk diperbaiki


## Dalam project ini melibatkan 3 file yaitu:

```
go-validators-request/app/Http/Controller/penduduk/penduduk.go
go-validators-request/app/Http/Requests/penduduk/penduduk.go
go-validators-request/app/Model/penduduk.go
go-validators-request/main.go
```

Pada Project ini kita fokuskan masalah validasi request. validasi akan di import dari controller setiap data yang sebelum di **Created** dan **Update**

Pada Controller dalam method **Cretaed** akan dipanggil fungsi ini:
```
// Validasi data
if err := Helper.ValidateData(penduduk.Nik, penduduk.Nama, penduduk.Tgl_lahir, penduduk.Waktu_kunjungan); err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
    return
}
```

Untuk memastikan fungsi **Helper.ValidateData** berfungsi maka akan di import file validation requests terlebih dahulu:
```
Helper "main.core/go-validator-request/app/Http/Requests/penduduk"
```

Sedangkan Isi Validasi terdiri dari:
1. Memastikan data tidak kosong ketika dikirim
```
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
```
2. Format NIK mesti terdiri angka dan sebanyak 16 karakter
```
if match, _ := regexp.MatchString("^[0-9]{16}$", nik); !match {
    return errors.New("NIK tidak valid, haruns mengandung 16 karakter nomor")
}
```
3. Nama Wajib lebih dari 2 karakter
```
if len(nama) < 3 {
    return errors.New("nama terlalu pendek")
}
```
4. Format Tanggal Lahir Mesti YYYY-MM-DD
```
if match, _ := regexp.MatchString("^[0-9]{4}-[0-9]{2}-[0-9]{2}$", tgl_lahir); !match {
    return errors.New("format tanggal salah, harus dalam format 'YYYY-MM-DD'")
}
```
5. Format Waktu Kunjungan Mesti YYYY-MM-DD H:i:s
```
layout := "2006-01-02 15:04:05"
_, err := time.Parse(layout, waktu_kunjungan)
if err != nil {
    return errors.New("format waktu_kunjungan harus yyyy-mm-dd H:i:s")
}
```


## Cara Menjalankan Project:

1. Buka **CMD**
2. Masuk ke Folder project melalui **CMD**
3. ketik **go run main.go** 


## Collection Postman

Untuk mempermudah terting data saya telah menyiapkan di folder ini:
```
go-validators-request/postman-collection/Go.postman_collection.json
```
Silahkan import kedalam postman anda

Enjoyy..

Jika ada kendala jangan sungkan email ke andikaisra7@gmail.com


## Catatan Penggunaan Tiap Folder:

- app/Function -> Fungsinya sebagai penyimpanan query kompleks atau fungsi query yang memiliki relasi yang kompleks yang nanti sisa dipanggil oleh controller
- app/Helpers -> Fungsinya sebagai penyimpanan file helper seperti mengubah format tanggal, waktu, uang dan lain-lain
- app/Http/Controllers -> Fungsinya sebagai penyimpanan seluruh controller
- app/Http/Requests -> Fungsinya sebagai penyimpanan semua validator requests yang nanti akan dipanggil di controller untuk validasi input data seperti simpan dan update
- app/Model -> Fungsinya sebagai penyimpanan folder-folder Model, Connect Database serta migration

## Catatan Batasan Penggunaan Komponen:

- Model: Komponen yang berinteraksi dengan database dan pengubah format data seperti mengubah format tanggal yang dipanggil dari Helper
- Helper: Komponen fungsi yang dapat dipanggil dari mana saja
- Controller: Komponen yang mengelola pemanggilan data dari model, pemanggilan validasi requests, pemanggilan Functions Query Kompleks dan Pengelolaan Rest API
- Request: Komponen yang mengelola validasi input dari user termasuk keamanan SQL Injection dan XSS Script
- Main.go: Komponen yang mengelola Routing dan Pengelolaan hak akses serta authentikasi menggunakan JWT auth