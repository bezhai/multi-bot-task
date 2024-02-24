package temp

import (
	"time"

	"github.com/bezhai/multi-bot-task/biz/model/image_store"
)

type PixivImageMetaInfo struct {
	PixivAddr   string                  `thrift:"pixiv_addr,1,required" bson:"pixiv_addr" form:"pixiv_addr,required" json:"pixiv_addr,required" query:"pixiv_addr,required"`
	Visible     bool                    `thrift:"visible,2,required" bson:"visible" form:"visible,required" json:"visible,required" query:"visible,required"`
	Author      string                  `thrift:"author,3,required" bson:"author" form:"author,required" json:"author,required" query:"author,required"`
	MultiTags   []*image_store.MultiTag `thrift:"multi_tags,4,optional" bson:"multi_tags" form:"multi_tags" json:"multi_tags,omitempty" query:"multi_tags"`
	CreateTime  time.Time               `thrift:"create_time,5,required" bson:"create_time" form:"create_time,required" json:"create_time,required" query:"create_time,required"`
	UpdateTime  time.Time               `thrift:"update_time,6,required" bson:"update_time" form:"update_time,required" json:"update_time,required" query:"update_time,required"`
	TosFileName string                  `thrift:"tos_file_name,7,required" bson:"tos_file_name" form:"tos_file_name,required" json:"tos_file_name,required" query:"tos_file_name,required"`
	AuthorID    *string                 `thrift:"author_id,8,optional" bson:"author_id" form:"author_id" json:"author_id,omitempty" query:"author_id"`
	DelFlag     bool                    `thrift:"del_flag,9,required" bson:"del_flag" form:"del_flag,required" json:"del_flag,required" query:"del_flag,required"`
	IllustID    int32                   `thrift:"illust_id,10,required" bson:"illust_id" form:"illust_id,required" json:"illust_id,required" query:"illust_id,required"`
	Title       string                  `thrift:"title,11,required" bson:"title" form:"title,required" json:"title,required" query:"title,required"`
}

type PixivImageMetaInfoWithUrl struct {
	PixivImageMetaInfo *PixivImageMetaInfo `thrift:"pixiv_image_meta_info,1,required" form:"pixiv_image_meta_info,required" json:"pixiv_image_meta_info,required" query:"pixiv_image_meta_info,required"`
	ShowURL            *string             `thrift:"show_url,2,optional" form:"show_url" json:"show_url,omitempty" query:"show_url"`
	DownloadURL        *string             `thrift:"download_url,3,optional" form:"download_url" json:"download_url,omitempty" query:"download_url"`
}

type ListPixivImageMetaInfoResponseData struct {
	PixivImageMetaInfos []*PixivImageMetaInfoWithUrl `thrift:"pixiv_image_meta_infos,1,required" form:"pixiv_image_meta_infos,required" json:"pixiv_image_meta_infos,required" query:"pixiv_image_meta_infos,required"`
	Total               int32                        `thrift:"total,2,required" form:"total,required" json:"total,required" query:"total,required"`
}
