package handlers

import (
	"net/http"
	"ranking-service/internal/models"
	"ranking-service/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VideoHandler struct {
	rankingService *services.RankingService
}

func NewVideoHandler(rankingService *services.RankingService) *VideoHandler {
	return &VideoHandler{rankingService: rankingService}
}

// @Summary Update video score
// @Description Update video score based on new interactions
// @Tags videos
// @Accept json
// @Produce json
// @Param id path int true "Video ID"
// @Param update body models.ScoreUpdate true "Score update data"
// @Success 200 {string} string "Score updated successfully"
// @Router /api/v1/videos/{id}/score [post]
func (h *VideoHandler) UpdateScore(c *gin.Context) {
	videoID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid video ID"})
		return
	}

	var update models.ScoreUpdate
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.rankingService.UpdateVideoScore(uint(videoID), update); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Score updated successfully"})
}

// @Summary Get top videos
// @Description Get top ranked videos globally
// @Tags videos
// @Produce json
// @Param limit query int false "Limit number of results" default(10)
// @Success 200 {array} models.Video
// @Router /api/v1/videos/top [get]
func (h *VideoHandler) GetTopVideos(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	videos, err := h.rankingService.GetTopVideos(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, videos)
}
