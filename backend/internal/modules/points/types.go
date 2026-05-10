package points

import "time"

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type MapPoint struct {
	ID            int       `json:"id"`
	CategoryID    *int      `json:"category_id"`
	Name          string    `json:"name"`
	Latitude      float64   `json:"latitude"`
	Longitude     float64   `json:"longitude"`
	Address       string    `json:"address"`
	OwnerID       string    `json:"owner_id"`
	OwnerName     string    `json:"owner_name"`
	OwnerAvatar   string    `json:"owner_avatar"`
	TahunBerdiri  *string   `json:"tahun_berdiri"`
	Description   string    `json:"description"`
	CoverImage    *string   `json:"cover_image"`
	Status        string    `json:"status"`
	RejectionNote *string   `json:"rejection_note"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Blog struct {
	ID         int       `json:"id"`
	MapPointID int       `json:"map_point_id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreatePointReq struct {
	CategoryID   int     `json:"category_id"`
	Name         string  `json:"name"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Address      string  `json:"address"`
	TahunBerdiri string  `json:"tahun_berdiri"`
	Description  string  `json:"description"`
	CoverImage   string  `json:"cover_image"`
	Status       string  `json:"status"`
}

type UpdatePointReq struct {
	CategoryID   int     `json:"category_id"`
	Name         string  `json:"name"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Address      string  `json:"address"`
	TahunBerdiri string  `json:"tahun_berdiri"`
	Description  string  `json:"description"`
	CoverImage   string  `json:"cover_image"`
	Status       string  `json:"status"`
}

type UpsertBlogReq struct {
	Content string `json:"content"`
}

type VerifyPointReq struct {
	Status        string `json:"status"`
	RejectionNote string `json:"rejection_note"`
}

type BlogDetailResp struct {
	Blog  *Blog     `json:"blog"`
	Point *MapPoint `json:"point"`
}
