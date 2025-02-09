package daos

import (
	"ranking-service/internal/models"
)

type VideoDAOInterface interface {
	GetOrCreateVideo(videoID uint) (*models.Video, error)
	UpdateVideoStats(videoID uint, update models.ScoreUpdate) error
	GetTopVideos(limit int) ([]models.Video, error)
}
