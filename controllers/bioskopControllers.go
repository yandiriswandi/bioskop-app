package controllers

import (
	"bioskop-app/config"
	"bioskop-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBioskop(c *gin.Context) {
	var bioskop models.Bioskop
	// Validasi input
	if err := c.ShouldBindJSON(&bioskop); err != nil {
		// Tapi jangan langsung return, simpan errornya dulu
		validationError := err.Error()

		// Cek manual field kosong
		if bioskop.Nama == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  400,
				"message": "Nama tidak boleh kosong",
			})
			return
		}
		if bioskop.Lokasi == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  400,
				"message": "Lokasi tidak boleh kosong",
			})
			return
		}

		if bioskop.Rating > 100 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  400,
				"message": "Rating tidak boleh lebih dari 100",
			})
			return
		}

		// Kalau field udah aman, baru kalau masih ada error di ShouldBindJSON, kirim errornya
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": validationError,
		})
		return
	}

	query := `INSERT INTO bioskop (nama, lokasi, rating) VALUES ($1, $2, $3) RETURNING id`
	err := config.DB.QueryRow(query, bioskop.Nama, bioskop.Lokasi, bioskop.Rating).Scan(&bioskop.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan bioskop"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Data berhasil dibuat!",
		"data":    bioskop,
	})
}
