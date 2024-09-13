package models

type DownloadYoutubeVideoInput struct {
	YoutubeURL string `json:"youtube_url"`
}

type VideoResponse struct {
	URL string `json:"url"`
}
