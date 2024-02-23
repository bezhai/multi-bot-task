package pixiv_model

type TranslateWord struct {
	Origin       string  `bson:"origin" json:"origin"`               // 原字段
	Translation  *string `bson:"translation" json:"translation"`     // 翻译
	HasTranslate bool    `bson:"has_translate" json:"has_translate"` // 是否已翻译
	ExtraInfo    *Extra  `bson:"extra_info" json:"extra_info"`       // 额外信息
}

type Extra struct {
	Zh string `json:"zh" bson:"zh"` // 中文翻译提示
	En string `json:"en" bson:"en"` // 英文翻译提示
}

type TranslateWithNum struct {
	Word *TranslateWord `bson:"word" json:"word"`
	Num  int64          `bson:"num" json:"num"` // 翻译关联数量
}
