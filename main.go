package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main.core/go-validator-request/app/Http/Controllers/penduduk"
	"main.core/go-validator-request/config/database"
	"main.core/go-validator-request/migration"
)

func init() {
	database.ConnectDatabase()
	migration.Migration()
}

func main() {

	r := gin.Default()

	// Menambahkan middleware CORS
	r.Use(cors.Default())

	r.POST("/api/penduduk", penduduk.Create)
	r.PUT("/api/penduduk/:nik", penduduk.Update)

	r.Run()

}
