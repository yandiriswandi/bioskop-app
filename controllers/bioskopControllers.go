package controllers

import (
	"bioskop-app/config"
	"bioskop-app/models"
	"database/sql"
	"net/http"
	"strconv"

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
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Gagal menambahkan bioskop",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Data berhasil dibuat!",
		"data":    bioskop,
	})
}

func UpdateBioskop(c *gin.Context) {
	var bioskop models.Bioskop
	id := c.Param("id") // ambil id dari URL parameter

	// Validasi input
	if err := c.ShouldBindJSON(&bioskop); err != nil {
		validationError := err.Error()

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

		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": validationError,
		})
		return
	}

	query := `UPDATE bioskop SET nama = $1, lokasi = $2, rating = $3 WHERE id = $4`
	res, err := config.DB.Exec(query, bioskop.Nama, bioskop.Lokasi, bioskop.Rating, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Gagal memperbarui bioskop",
		})
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  404,
			"message": "Bioskop tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Data berhasil diperbarui!",
		"data":    bioskop,
	})
}

func DeleteBioskop(c *gin.Context) {
	id := c.Param("id") // ambil id dari URL parameter

	query := `DELETE FROM bioskop WHERE id = $1`
	res, err := config.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Gagal menghapus bioskop",
		})
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  404,
			"message": "Bioskop tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Bioskop berhasil dihapus",
	})
}

func GetBioskopList(c *gin.Context) {
	// Ambil query parameter page dan size, default kalau kosong
	pageStr := c.DefaultQuery("page", "1")
	sizeStr := c.DefaultQuery("size", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil || size < 1 {
		size = 10
	}

	offset := (page - 1) * size

	// Hitung total data
	var totalData int
	err = config.DB.QueryRow(`SELECT COUNT(*) FROM bioskop`).Scan(&totalData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Gagal mengambil total data",
		})
		return
	}

	// Ambil data bioskop
	rows, err := config.DB.Query(`SELECT id, nama, lokasi, rating FROM bioskop ORDER BY id LIMIT $1 OFFSET $2`, size, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Gagal mengambil data bioskop",
		})
		return
	}
	defer rows.Close()

	var bioskops []models.Bioskop
	for rows.Next() {
		var bioskop models.Bioskop
		if err := rows.Scan(&bioskop.ID, &bioskop.Nama, &bioskop.Lokasi, &bioskop.Rating); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  500,
				"message": "Gagal membaca data bioskop",
			})
			return
		}
		bioskops = append(bioskops, bioskop)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    200,
		"message":   "Berhasil mengambil list bioskop",
		"data":      bioskops,
		"page":      page,
		"size":      size,
		"totalData": totalData,
	})
}

func GetBioskopByID(c *gin.Context) {
	id := c.Param("id")

	var bioskop models.Bioskop
	query := `SELECT id, nama, lokasi, rating FROM bioskop WHERE id = $1`
	err := config.DB.QueryRow(query, id).Scan(&bioskop.ID, &bioskop.Nama, &bioskop.Lokasi, &bioskop.Rating)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  404,
				"message": "Bioskop tidak ditemukan",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  500,
				"message": "Gagal mengambil detail bioskop",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Berhasil mengambil detail bioskop",
		"data":    bioskop,
	})
}
