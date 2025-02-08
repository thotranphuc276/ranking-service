package services

import (
	"real-time-ranking/internal/models"
	"testing"
	"time"
)

type mockVideoDAO struct {
	videos map[uint]models.Video
}

func newMockVideoDAO() *mockVideoDAO {
	return &mockVideoDAO{
		videos: make(map[uint]models.Video),
	}
}

func (m *mockVideoDAO) GetOrCreateVideo(videoID uint) (*models.Video, error) {
	video, exists := m.videos[videoID]
	if !exists {
		video = models.Video{
			ID:        videoID,
			Score:     0,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}
		m.videos[videoID] = video
	}
	return &video, nil
}

func (m *mockVideoDAO) UpdateScore(videoID uint, score float64) error {
	video, _ := m.GetOrCreateVideo(videoID)
	video.Score = score
	m.videos[videoID] = *video
	return nil
}

func (m *mockVideoDAO) UpdateVideoStats(videoID uint, update models.ScoreUpdate) error {
	video, _ := m.GetOrCreateVideo(videoID)
	if update.Views != nil {
		video.Views = *update.Views
	}
	if update.Likes != nil {
		video.Likes = *update.Likes
	}
	if update.Comments != nil {
		video.Comments = *update.Comments
	}
	if update.Shares != nil {
		video.Shares = *update.Shares
	}
	if update.WatchTime != nil {
		video.WatchTime = *update.WatchTime
	}
	video.UpdatedAt = time.Now().Unix()
	m.videos[videoID] = *video
	return nil
}

func (m *mockVideoDAO) GetTopVideos(limit int) ([]models.Video, error) {
	videos := make([]models.Video, 0, len(m.videos))
	for _, v := range m.videos {
		videos = append(videos, v)
	}
	return videos, nil
}

func (m *mockVideoDAO) GetUserTopVideos(userID uint, limit int) ([]models.Video, error) {
	videos := make([]models.Video, 0)
	for _, v := range m.videos {
		if v.UserID == userID {
			videos = append(videos, v)
		}
	}
	return videos, nil
}

func TestCalculateScore(t *testing.T) {
	tests := []struct {
		name     string
		update   models.ScoreUpdate
		expected float64
	}{
		{
			name: "all metrics",
			update: models.ScoreUpdate{
				Views:     ptr(int64(100)),
				Likes:     ptr(int64(50)),
				Comments:  ptr(int64(20)),
				Shares:    ptr(int64(10)),
				WatchTime: ptr(int64(1000)),
			},
			expected: 100 + 100 + 60 + 40 + 100,
		},
		{
			name: "partial metrics",
			update: models.ScoreUpdate{
				Views: ptr(int64(100)),
				Likes: ptr(int64(50)),
			},
			expected: 100 + 100,
		},
	}

	service := NewRankingService(newMockVideoDAO())

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := service.CalculateScore(tt.update)
			if score != tt.expected {
				t.Errorf("Expected score %f, got %f", tt.expected, score)
			}
		})
	}
}

func TestUpdateVideoScore(t *testing.T) {
	dao := newMockVideoDAO()
	service := NewRankingService(dao)

	videoID := uint(1)
	update := models.ScoreUpdate{
		Views:     ptr(int64(100)),
		Likes:     ptr(int64(50)),
		Comments:  ptr(int64(20)),
		Shares:    ptr(int64(10)),
		WatchTime: ptr(int64(1000)),
	}

	if err := service.UpdateVideoScore(videoID, update); err != nil {
		t.Errorf("Failed to update video score: %v", err)
	}

	videos, err := dao.GetTopVideos(1)
	if err != nil {
		t.Errorf("Failed to get videos: %v", err)
	}

	if len(videos) != 1 {
		t.Errorf("Expected 1 video, got %d", len(videos))
	}

	if videos[0].Views != *update.Views {
		t.Errorf("Expected views %d, got %d", *update.Views, videos[0].Views)
	}
}

func ptr[T any](v T) *T {
	return &v
}
