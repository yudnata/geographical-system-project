package points

import "time"

type ObjectType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type GeoPoint struct {
	ID                int       `json:"id"`
	TypeID            *int      `json:"type_id"`
	Name              string    `json:"name"`
	Latitude          float64   `json:"latitude"`
	Longitude         float64   `json:"longitude"`
	Address           string    `json:"address"`
	OwnerID           string    `json:"owner_id"`
	TahunBerdiri      *int      `json:"tahun_berdiri"`
	StatusKepemilikan string    `json:"status_kepemilikan"`
	Description       string    `json:"description"`
	IsActive          bool      `json:"is_active"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type CreatePointReq struct {
	TypeID            int     `json:"type_id"`
	Name              string  `json:"name"`
	Latitude          float64 `json:"latitude"`
	Longitude         float64 `json:"longitude"`
	Address           string  `json:"address"`
	TahunBerdiri      int     `json:"tahun_berdiri"`
	StatusKepemilikan string  `json:"status_kepemilikan"`
	Description       string  `json:"description"`
}

type UpdatePointReq struct {
	TypeID            int     `json:"type_id"`
	Name              string  `json:"name"`
	Latitude          float64 `json:"latitude"`
	Longitude         float64 `json:"longitude"`
	Address           string  `json:"address"`
	TahunBerdiri      int     `json:"tahun_berdiri"`
	StatusKepemilikan string  `json:"status_kepemilikan"`
	Description       string  `json:"description"`
	IsActive          bool    `json:"is_active"`
}
