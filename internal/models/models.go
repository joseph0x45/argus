package models

type ChannelInfo struct {
	ID     string
	Name   string
	RssURL string
}

type Channel struct {
	ID          string `db:"id"`
	ChannelID   string `db:"channel_id"`
	ChannelName string `db:"channel_name"`
	RssURL      string `db:"rss_url"`
	AddedAt     string `db:"added_at"`
}
