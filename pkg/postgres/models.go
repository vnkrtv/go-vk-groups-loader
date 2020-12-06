package postgres

import (
	"time"

	"database/sql"
)

const dbSchema = `
CREATE TABLE IF NOT EXISTS groups (
	group_id INTEGER 
             PRIMARY KEY,

	screen_name TEXT 
             UNIQUE 
             NOT NULL,

    name     TEXT
             UNIQUE
             NOT NULL,

	members_count INTEGER
             NOT NULL
             CHECK (members_count >= 0)
);

CREATE TABLE IF NOT EXISTS posts (
    post_id  INTEGER
             NOT NULL,

	group_screen_name TEXT,

    date     TIMESTAMP
             NOT NULL,

    title    TEXT
             NOT NULL,

    text     TEXT
			 NOT NULL,

    likes_count INTEGER
             NOT NULL
             CHECK (likes_count >= 0),

    views_count INTEGER
             NOT NULL
             CHECK (views_count >= 0),

    comments_count INTEGER
             NOT NULL
             CHECK (comments_count >= 0),

    reposts_count INTEGER
             NOT NULL
             CHECK (reposts_count >= 0),

	PRIMARY KEY (post_id, date),

	CONSTRAINT fk_group FOREIGN KEY (group_screen_name) 
		REFERENCES groups (screen_name)
			ON DELETE SET NULL
			ON UPDATE CASCADE
);`

type Group struct {
	ID           int    `db:"group_id"`
	ScreenName   string `db:"screen_name"`
	Name         string `db:"name"`
	MembersCount int    `db:"members_count"`
}

type Post struct {
	ID              int            `db:"post_id"`
	GroupScreenName sql.NullString `db:"group_screen_name"`
	Date            time.Time      `db:"date"`
	Title           string         `db:"title"`
	Text            string         `db:"text"`
	LikesCount      int            `db:"likes_count"`
	ViewsCount      int            `db:"views_count"`
	CommentsCount   int            `db:"comments_count"`
	RepostsCount    int            `db:"reposts_count"`
}
