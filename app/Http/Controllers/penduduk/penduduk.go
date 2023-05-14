package penduduk

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"

	Helper "main.core/go-validator-request/app/Http/Requests/penduduk"
	Model "main.core/go-validator-request/app/Model/penduduk"
)

func Create(c *gin.Context) {
	var penduduk Model.Penduduk

	if err := c.ShouldBindJSON(&penduduk); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Validasi data
	if err := Helper.ValidateData(penduduk.Nik, penduduk.Nama, penduduk.Tgl_lahir, penduduk.Waktu_kunjungan); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := Model.DB.Table("penduduk").Create(&penduduk).Error; err != nil {
		if IsDuplicateError(err) {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "Data dengan NIK yang sama sudah ada"})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Berhasil Menyimpan Data"})
}

func IsDuplicateError(err error) bool {
	var mysqlError *mysql.MySQLError
	if errors.As(err, &mysqlError) && mysqlError.Number == 1062 {
		return true
	}

	return false
}

func Update(c *gin.Context) {

	var penduduk Model.Penduduk
	nik := c.Param("nik")

	if err := c.ShouldBindJSON(&penduduk); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Validasi data
	if err := Helper.ValidateData(penduduk.Nik, penduduk.Nama, penduduk.Tgl_lahir, penduduk.Waktu_kunjungan); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if Model.DB.Table("penduduk").Model(&penduduk).Where("nik = ?", nik).Updates(&penduduk).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Tidak dapat mengubah data penduduk"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Data Berhasil Diperbaharui"})

}
