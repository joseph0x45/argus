package db

import "github.com/joseph0x45/sad"

var migrations = []sad.Migration{
	{
		Version: 1,
		Name:    "create_channels_table",
		SQL: `
      create table channels (
        id text primary key,
        channel_id text not null unique,
        channel_name text not null,
        rss_url text not null,
        added_at text not null
      );
    `,
	},
}
