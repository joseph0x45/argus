package db

import (
	"fmt"

	"github.com/joseph0x45/argus/internal/models"
)

func (c *Conn) InsertChannel(channel *models.Channel) error {
	const query = `
    insert into channels (
      id, channel_id, channel_name,
      rss_url, added_at
    )
    values (
      :id, :channel_id, :channel_name,
      :rss_url, :added_at
    );
  `
	if _, err := c.db.NamedExec(
		query, channel,
	); err != nil {
		return fmt.Errorf("Error while inserting channel: %w", err)
	}
	return nil
}
