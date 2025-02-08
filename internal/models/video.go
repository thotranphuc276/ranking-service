package models

type Video struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	UserID    uint    `json:"user_id"`
	Title     string  `json:"title"`
	Score     float64 `json:"score"`
	Views     int64   `json:"views"`
	Likes     int64   `json:"likes"`
	Comments  int64   `json:"comments"`
	Shares    int64   `json:"shares"`
	WatchTime int64   `json:"watch_time"`
	CreatedAt int64   `json:"created_at"`
	UpdatedAt int64   `json:"updated_at"`
}

type ScoreUpdate struct {
	Views     *int64 `json:"views"`
	Likes     *int64 `json:"likes"`
	Comments  *int64 `json:"comments"`
	Shares    *int64 `json:"shares"`
	WatchTime *int64 `json:"watch_time"`
}
