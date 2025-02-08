package daos

import (
	"ranking-service/internal/models"
)

type VideoDAOInterface interface {
	GetOrCreateVideo(videoID uint) (*models.Video, error)
	UpdateScore(videoID uint, score float64) error
	UpdateVideoStats(videoID uint, update models.ScoreUpdate) error
	GetTopVideos(limit int) ([]models.Video, error)
	GetUserTopVideos(userID uint, limit int) ([]models.Video, error)
}
