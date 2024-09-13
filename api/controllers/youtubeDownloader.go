package controllers

import (
	"fmt"
	"net/http"
	"os/exec"

	"GOLANG/api/models"

	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func downloadVideo(videoURL, outputPath string) (string, error) {
	uuidV4 := uuid.New()
	outputFile := filepath.Join(outputPath, uuidV4.String()+".mp4")

	// yt-dlp komutunu çağır
	cmd := exec.Command("yt-dlp", "-f", "best", "-o", outputFile, "-u", "yasincakir.py@gmail.com", "-p", "lolobogolu123", videoURL)

	// Komutu çalıştır
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("video indirilemedi: %w", err)
	}

	return outputFile, nil
}

// @Tags youtube-downloader
// @Summary Youtube Download Video
// @Schemes
// @Accept json
// @Produce json
// @Param download body models.DownloadYoutubeVideoInput true "Download Youtube Video"
// @Success 200 {object} models.VideoResponse
// @Router /youtube/download [post]
func DownloadYoutubeVideo(c *gin.Context) {

	var DownloadYoutubeVideoInput models.DownloadYoutubeVideoInput

	if err := c.ShouldBindJSON(&DownloadYoutubeVideoInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	videoURL := DownloadYoutubeVideoInput.YoutubeURL

	outputPath := "./static/downloaded_videos"
	videoFilePath, err := downloadVideo(videoURL, outputPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	videoFileURL := fmt.Sprintf("/static/downloaded_videos/%s", filepath.Base(videoFilePath))

	response := models.VideoResponse{
		URL: videoFileURL,
	}

	c.JSON(http.StatusOK, response)
}
