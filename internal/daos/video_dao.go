package daos

import (
	"context"
	"fmt"
	"ranking-service/internal/models"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type VideoDAO struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewVideoDAO(db *gorm.DB, redis *redis.Client) *VideoDAO {
	return &VideoDAO{db: db, redis: redis}
}

func (d *VideoDAO) GetOrCreateVideo(videoID uint) (*models.Video, error) {
	var video models.Video
	result := d.db.First(&video, videoID)

	if result.Error == gorm.ErrRecordNotFound {
		video = models.Video{
			ID:        videoID,
			Score:     0,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}
		if err := d.db.Create(&video).Error; err != nil {
			return nil, fmt.Errorf("failed to create video: %w", err)
		}
	} else if result.Error != nil {
		return nil, fmt.Errorf("failed to query video: %w", result.Error)
	}

	return &video, nil
}

func (d *VideoDAO) UpdateVideoStats(videoID uint, update models.ScoreUpdate) error {
	video, err := d.GetOrCreateVideo(videoID)
	if err != nil {
		return err
	}

	updates := make(map[string]interface{})
	updates["updated_at"] = time.Now().Unix()

	if update.Views != nil {
		updates["views"] = *update.Views + video.Views
	}
	if update.Likes != nil {
		updates["likes"] = *update.Likes + video.Likes
	}
	if update.Comments != nil {
		updates["comments"] = *update.Comments + video.Comments
	}
	if update.Shares != nil {
		updates["shares"] = *update.Shares + video.Shares
	}
	if update.WatchTime != nil {
		updates["watch_time"] = *update.WatchTime + video.WatchTime
	}

	if update.Score != nil {
		updates["score"] = *update.Score + video.Score
	}

	return d.db.Model(video).Updates(updates).Error
}

func (d *VideoDAO) GetTopVideos(limit int) ([]models.Video, error) {
	ctx := context.Background()
	result, err := d.redis.ZRevRangeWithScores(ctx, "video_rankings", 0, int64(limit-1)).Result()
	if err != nil {
		return nil, err
	}

	var videos []models.Video
	for _, z := range result {
		videoID, _ := strconv.ParseUint(fmt.Sprint(z.Member), 10, 64)
		var video models.Video
		if err := d.db.First(&video, uint(videoID)).Error; err != nil {
			continue
		}
		videos = append(videos, video)
	}
	return videos, nil
}
