package points

import "time"

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type MapPoint struct {
	ID                int       `json:"id"`
	CategoryID        *int      `json:"category_id"`
	Name              string    `json:"name"`
	Latitude          float64   `json:"latitude"`
	Longitude         float64   `json:"longitude"`
	Address           string    `json:"address"`
	OwnerID           string    `json:"owner_id"`
	OwnerName         string    `json:"owner_name"`
	TahunBerdiri      *int      `json:"tahun_berdiri"`
	StatusKepemilikan string    `json:"status_kepemilikan"`
	Description       string    `json:"description"`
	IsActive          bool      `json:"is_active"`
	Status            string    `json:"status"`
	RejectionNote     *string   `json:"rejection_note"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type Blog struct {
	ID         int       `json:"id"`
	MapPointID int       `json:"map_point_id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CoverPhoto *string   `json:"cover_photo"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreatePointReq struct {
	CategoryID        int     `json:"category_id"`
	Name              string  `json:"name"`
	Latitude          float64 `json:"latitude"`
	Longitude         float64 `json:"longitude"`
	Address           string  `json:"address"`
	TahunBerdiri      int     `json:"tahun_berdiri"`
	StatusKepemilikan string  `json:"status_kepemilikan"`
	Description       string  `json:"description"`
	Status            string  `json:"status"` // draft or pending
}

type UpdatePointReq struct {
	CategoryID        int     `json:"category_id"`
	Name              string  `json:"name"`
	Latitude          float64 `json:"latitude"`
	Longitude         float64 `json:"longitude"`
	Address           string  `json:"address"`
	TahunBerdiri      int     `json:"tahun_berdiri"`
	StatusKepemilikan string  `json:"status_kepemilikan"`
	Description       string  `json:"description"`
	IsActive          bool    `json:"is_active"`
	Status            string  `json:"status"` // draft or pending
}
