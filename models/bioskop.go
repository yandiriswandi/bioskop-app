package models

type Bioskop struct {
	ID     int     `db:"id" json:"id"`
	Nama   string  `db:"nama" json:"nama" binding:"required"`
	Lokasi string  `db:"lokasi" json:"lokasi" binding:"required"`
	Rating float64 `db:"rating" json:"rating" binding:"lte=100"`
}
