package cli

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/joseph0x45/argus/internal/db"
	"github.com/joseph0x45/argus/internal/fetcher"
	"github.com/joseph0x45/argus/internal/models"
	"github.com/teris-io/shortid"
)

func addChannel(args []string) int {
	conn := db.GetConn()
	defer conn.Close()
	flagSet := flag.NewFlagSet("add-channel", flag.ContinueOnError)
	channelURL := flagSet.String("url", "", "URL of the channel to add")
	flagSet.Parse(args)
	if *channelURL == "" {
		fmt.Fprintln(os.Stderr, "'url' is required")
		printUsage()
		return 1
	}
	channelInfo, err := fetcher.ExtractChannelInfo(*channelURL)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	channel := &models.Channel{
		ID:          shortid.MustGenerate(),
		ChannelID:   channelInfo.ID,
		ChannelName: channelInfo.Name,
		RssURL:      channelInfo.RssURL,
		AddedAt:     time.Now().Format("Jan 02 2006"),
	}
	if err := conn.InsertChannel(channel); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return 0
}
