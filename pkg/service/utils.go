package service

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"database/sql"

	pg "github.com/vnkrtv/go-vk-news-loader/pkg/postgres"
	vk "github.com/vnkrtv/go-vk-news-loader/pkg/vkapi"
)

type Config struct {
	PGUser   string `json:"pguser"`
	PGPass   string `json:"pgpass"`
	PGName   string `json:"pgname"`
	PGHost   string `json:"pghost"`
	PGPort   string `json:"pgport"`
	VKToken  string `json:"vktoken"`
	Interval int    `json:"interval"`
}

func ParseVKGroup(vkGroup vk.VKGroup) pg.Group {
	return pg.Group{
		ID:           vkGroup.ID,
		ScreenName:   vkGroup.ScreenName,
		Name:         vkGroup.Name,
		MembersCount: vkGroup.MembersCount,
	}
}

func ParseVKWall(vkWall vk.VKWall, groupScreenName string) []pg.Post {
	var posts []pg.Post
	for _, post := range vkWall.Items {
		if len(post.Attachments) != 0 &&
			post.Attachments[0].Link.Title != "" &&
			post.Attachments[0].Link.Description != "" {
			group := sql.NullString{
				String: groupScreenName,
				Valid:  true,
			}
			post := pg.Post{
				ID:              post.ID,
				GroupScreenName: group,
				Date:            time.Unix(int64(post.Date), 0),
				Title:           post.Attachments[0].Link.Title,
				Text:            post.Attachments[0].Link.Description,
				LikesCount:      post.Likes.Count,
				ViewsCount:      post.Views.Count,
				CommentsCount:   post.Comments.Count,
				RepostsCount:    post.Reposts.Count,

			}
			posts = append(posts, post)
		}
	}
	return posts
}

func GetGroupsScreenNames(groupsPath string) ([]string, error) {
	bytes, err := ioutil.ReadFile(groupsPath)
	if err != nil {
		return nil, err
	}
	var groupsScreenNames []string
	err = json.Unmarshal(bytes, &groupsScreenNames)
	return groupsScreenNames, err
}

func GetConfig(configPath string) (Config, error) {
	var config Config
	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(bytes, &config)
	return config, err
}
