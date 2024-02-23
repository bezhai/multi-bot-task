package pixiv_model

import (
	"time"
)

type MultiTag struct {
	Name        string `bson:"name" json:"name"`
	Translation string `bson:"translation" json:"translation"`
	Visible     bool   `bson:"visible" json:"visible"`
}

type PixivImageInfo struct {
	ImageKey     string      `json:"image_key" bson:"image_key"`
	PixivAddr    string      `json:"pixiv_addr" bson:"pixiv_addr"`
	Visible      bool        `json:"visible" bson:"visible"`
	Author       string      `json:"author" bson:"author"`
	MultiTags    []*MultiTag `json:"multi_tags" bson:"multi_tags"`
	CreateTime   time.Time   `json:"create_time" bson:"create_time"`
	UpdateTime   time.Time   `json:"update_time" bson:"update_time"`
	Height       int         `json:"height" bson:"height"`
	Width        int         `json:"width" bson:"width"`
	Size         int64       `json:"size" bson:"size"`
	NeedDownload bool        `json:"need_download" bson:"need_download"`
	TosFileName  string      `json:"tos_file_name" bson:"tos_file_name"`
	AuthorId     *string     `json:"author_id" bson:"author_id"`
	DelFlag      bool        `json:"del_flag" bson:"del_flag"`
	IllustId     int         `json:"illust_id" bson:"illust_id"`
	Title        string      `json:"title" bson:"title"`
}

func (i *PixivImageInfo) GetVisibleTags() (res []string) {
	for _, tag := range i.MultiTags {
		if tag.Visible && tag.Translation != "" {
			res = append(res, tag.Translation)
		}
	}

	return
}

type StatusMode int

const (
	StatusNotDelete StatusMode = 0
	StatusVisible   StatusMode = 1
	StatusDelete    StatusMode = 2
	StatusAll       StatusMode = 3
	StatusNoVisible StatusMode = 4
)

type ListImageReq struct {
	Author   string     `json:"author"`    // 模糊查询
	AuthorId string     `json:"author_id"` // 精确查询
	Tags     []string   `json:"tags"`      // 按每个tag模糊查询的交集
	Status   StatusMode `json:"status"`
	Page     int        `json:"page"`
	PageSize int        `json:"page_size"`
	IllustId int        `json:"illust_id"`
}

type PixivImageInfoWithUrl struct {
	*PixivImageInfo
	Url         string `json:"url"`
	DownloadUrl string `json:"download_url"`
}

type UpdateImageStatusReq struct {
	PixivAddrList []string   `json:"pixiv_addr_list"`
	Status        StatusMode `json:"status"`
}
