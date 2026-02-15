package fetcher

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/joseph0x45/argus/internal/models"
)

func ExtractChannelInfo(url string) (*models.ChannelInfo, error) {
	cmd := exec.Command(
		"yt-dlp",
		"--skip-download",
		"--playlist-items",
		"1",
		"--print",
		"%(uploader)s|%(channel_id)s",
		url,
	)
	var out bytes.Buffer
	cmd.Stdout = &out
	if cmd.Run() != nil {
		return nil, fmt.Errorf("Error while extracting channel info.")
	}
	parts := strings.Split(strings.TrimSpace(out.String()), "|")
	if len(parts) != 2 {
		return nil, fmt.Errorf("Error while extracting channel info")
	}
	channelName, channelID := parts[0], parts[1]
	rssURL := fmt.Sprintf(
		"https://www.youtube.com/feeds/videos.xml?channel_id=%s",
		channelID,
	)
	return &models.ChannelInfo{
		ID:     channelID,
		Name:   channelName,
		RssURL: rssURL,
	}, nil
}
