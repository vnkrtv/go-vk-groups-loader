package vkapi

import (
	"fmt"
	"strconv"

	"github.com/go-vk-api/vk"
)

type VKLoader interface {
	GetGroups(screenNames []string) ([]VKGroup, error)
	GetGroupsPosts(domains []string, postsCount int) (map[string]VKWall, error)
}

type VKAPi struct {
	api *vk.Client
}

func NewVKApi(token string) (*VKAPi, error) {
	api, err := vk.NewClientWithOptions(
		vk.WithToken(token),
	)
	return &VKAPi{
		api: api,
	}, err
}

func (a *VKAPi) GetGroups(groupsScreenNames []string) ([]VKGroup, error)  {
	var groupIDs string
	for i, group := range groupsScreenNames {
		groupIDs += group
		if i != len(groupsScreenNames) - 1 {
			groupIDs += ","
		}
	}
	var groups []VKGroup
	err := a.api.CallMethod("groups.getById", vk.RequestParams{
		"group_ids": groupIDs,
		"fields": "description,activity,members_count",
	}, &groups)
	return groups, err
}

func (a *VKAPi) GetGroupsPosts(groupsScreenNames []string, postsCount int) (map[string]VKWall, error) {
	groups := make([]string, len(groupsScreenNames))
	for i, str := range groupsScreenNames {
		groups[i] = fmt.Sprintf("%s", strconv.Quote(str))
		if i != len(groups) - 1{
			groups[i] += ","
		}
	}
	var response []VKWall
	code := `
        var domains = %s;
		var res = [];
		var i = 0;
		while (i < domains.length) {
			var posts = API.wall.get({
				domain: domains[i], 
				count: %d,
				offset: 1
			});
			res.push(posts);
			i = i + 1; 
		}
		return res;`
	err := a.api.CallMethod("execute", vk.RequestParams{
		"code": fmt.Sprintf(code, groups, postsCount),
	}, &response)
	wallMap := make(map[string]VKWall, len(groups))
	for i, wall := range response {
		wallMap[groupsScreenNames[i]] = wall
	}
	return wallMap, err
}
