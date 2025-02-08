package services

import (
	"real-time-ranking/internal/daos"
	"real-time-ranking/internal/models"
)

type RankingService struct {
	videoDAO daos.VideoDAOInterface
}

func NewRankingService(videoDAO daos.VideoDAOInterface) *RankingService {
	return &RankingService{videoDAO: videoDAO}
}

func (s *RankingService) CalculateScore(update models.ScoreUpdate) float64 {
	viewWeight := 1.0
	likeWeight := 2.0
	commentWeight := 3.0
	shareWeight := 4.0
	watchTimeWeight := 0.1

	score := 0.0
	if update.Views != nil {
		score += float64(*update.Views) * viewWeight
	}
	if update.Likes != nil {
		score += float64(*update.Likes) * likeWeight
	}
	if update.Comments != nil {
		score += float64(*update.Comments) * commentWeight
	}
	if update.Shares != nil {
		score += float64(*update.Shares) * shareWeight
	}
	if update.WatchTime != nil {
		score += float64(*update.WatchTime) * watchTimeWeight
	}
	return score
}

func (s *RankingService) UpdateVideoScore(videoID uint, update models.ScoreUpdate) error {
	if err := s.videoDAO.UpdateVideoStats(videoID, update); err != nil {
		return err
	}

	score := s.CalculateScore(update)
	return s.videoDAO.UpdateScore(videoID, score)
}

func (s *RankingService) GetTopVideos(limit int) ([]models.Video, error) {
	return s.videoDAO.GetTopVideos(limit)
}

func (s *RankingService) GetUserTopVideos(userID uint, limit int) ([]models.Video, error) {
	return s.videoDAO.GetUserTopVideos(userID, limit)
}
