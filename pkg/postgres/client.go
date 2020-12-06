package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type GroupsStorage interface {
	InsertGroup(group Group) error
	InsertGroups(groups []Group) error
	GetGroupsScreenNames () ([]string, error)
}

type PostsStorage interface {
	InsertPost(post Post) error
	InsertPosts(post []Post) error
	UpdatePost(post Post) error
	UpdatePosts(post []Post) error
}

type NewsStorage interface {
	GroupsStorage
	PostsStorage
	CreateSchema() error
}

type Storage struct {
	 db *sqlx.DB
}

func OpenConnection(user, password, host, port, dbName string) (*Storage, error) {
	conStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbName)
	db, err := sqlx.Open("postgres", conStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	return &Storage{db: db}, err
}

func (s *Storage) CreateSchema() error {
	_, err := s.db.Exec(dbSchema)
	return err
}

func (s *Storage) InsertGroup(group Group) error {
	sql := `
		INSERT INTO 
			groups (group_id, screen_name, name, members_count) 
		VALUES 
			(:group_id, :screen_name, :name, :members_count)
		ON CONFLICT (group_id)
    		DO UPDATE SET
    			name = :name,
    			members_count = :members_count`
	_, err := s.db.NamedExec(sql, &group)
	return err
}

func (s *Storage) InsertGroups(groups []Group) error {
	for _, group := range groups {
		if err := s.InsertGroup(group); err != nil {
			return err
		}
	}
	return nil
}

func (s *Storage) GetGroupsScreenNames() ([]string, error) {
	var groups []string
	err := s.db.Select(&groups, "SELECT screen_name FROM groups")
	return groups, err
}

func (s *Storage) InsertPost(post Post) error {
	sql := `
		INSERT INTO 
			posts (post_id, group_screen_name, date, title, text, likes_count, views_count, comments_count, reposts_count) 
		VALUES 
			(:post_id, :group_screen_name, :date, :title, :text, :likes_count, :views_count, :comments_count, :reposts_count)
		ON CONFLICT (post_id, date)
		DO UPDATE SET
			title = :title,
			text = :text,
			likes_count = :likes_count,
			views_count = :views_count,
			comments_count = :comments_count,
			reposts_count = :reposts_count`
	_, err := s.db.NamedExec(sql, &post)
	return err
}

func (s *Storage) InsertPosts(posts []Post) error {
	for _, post := range posts {
		if err := s.InsertPost(post); err != nil {
			fmt.Printf("error: %#v", post)
			return err
		}
	}
	return nil
}

func (s *Storage) UpdatePost(post Post) error {
	sql := `
		UPDATE posts SET 
			group_screen_name = :group_screen_name, title = :title, 
			text = :text, likes_count = :likes_count, views_count = :views_count, 
			comments_count = :comments_count, reposts_count = :reposts_count 
		WHERE 
			post_id = :post_id AND date = :date`
	_, err := s.db.NamedExec(sql, &post)
	return err
}

func (s *Storage) UpdatePosts(posts []Post) error {
	for _, post := range posts {
		if err := s.UpdatePost(post); err != nil {
			return err
		}
	}
	return nil
}
