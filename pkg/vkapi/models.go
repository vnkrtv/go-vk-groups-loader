package vkapi

type VKGroup struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	ScreenName   string `json:"screen_name"`
	Type         string `json:"type"`
	Description  string `json:"description"`
	Activity     string `json:"activity"`
	MembersCount int    `json:"members_count"`
}

type VKWall struct {
	Count int      `json:"count"`
	Items []VKPost `json:"items"`
}

type VKPost struct {
	ID          int             `json:"id"`
	Date        int             `json:"date"`
	PostType    string          `json:"post_type"`
	Text        string          `json:"text"`
	IsPinned    int8            `json:"is_pinned"`
	Comments    struct {
		Count		int	     		 `json:"count"`
	}                           `json:"comments"`
	Likes       struct {
		Count		int  			 `json:"count"`
	}                           `json:"likes"`
	Reposts     struct {
		Count		int  			 `json:"count"`
	}                           `json:"reposts"`
	Views       struct {
		Count		int 			 `json:"count"`
	}                           `json:"views"`
	Attachments []VKAttachments `json:"attachments"`
}

type VKAttachments struct {
	Type  string   `json:"type"`
	Link  struct {
		Url          string   `json:"url"`
		Title        string   `json:"title"`
		Description  string   `json:"description"`
	}              `json:"link"`
}
