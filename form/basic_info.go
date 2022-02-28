package form

type BasicInfo struct {
	Key                string `json:"key" form:"key"`
	ChineseText        string `json:"chinese_text" form:"chinese_text"`
	EnglishText        string `json:"english_text" form:"english_text"`
	Remark             string `json:"remark" form:"remark"`
	IsMachineTranslate bool   `json:"is_machine_translate" form:"is_machine_translate"`
}

type BatchBasicInfo struct {
	BasicInfo []BasicInfo `json:"basic_info" form:"basic_info"`
}

type GetBasicInfoListByCondition struct {
	PageSize    int
	PageNum     int
	Key         string
	ChineseText string
	EnglishText string
}

type GetBasicInfoListBySearchKey struct {
	PageSize  int
	PageNum   int
	SearchKey string
}

type NewBasicInfo struct {
	ID                 int    `json:"id"`
	Key                string `json:"key"`
	ChineseText        string `json:"chinese_text"`
	EnglishText        string `json:"english_text"`
	Remark             string `json:"remark"`
	IsMachineTranslate bool   `json:"is_machine_translate"`
}

type KeyLang struct {
	Key  []string `json:"key"`
	Lang string   `json:"lang"`
}
