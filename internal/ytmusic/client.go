package ytmusic

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type Song struct {
	Title    string
	URL      string
	Duration string
	Artist   string
	Album    string
}

func FetchSong(query string) (*Song, error) {
	cmd := exec.Command(
		"yt-dlp",
		"ytsearch1:"+query,
		"--skip-download",
		"--print-json",
	)

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var result struct {
		Title      string `json:"title"`
		WebpageURL string `json:"webpage_url"`
		Duration   int    `json:"duration"`
	}

	if err := json.Unmarshal(out, &result); err != nil {
		return nil, err

	}
	song := &Song{
		Title:    result.Title,
		URL:      result.WebpageURL,
		Duration: fmt.Sprintf("%d:%02d", result.Duration/60, result.Duration%60),
	}
	return song, nil
}
