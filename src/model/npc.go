package model

type NPC struct {
	Model

	Name        string `json:"name"`
	LikeCount   int    `json:"like_count"`
	FollowCount int    `json:"follow_count"`
	HotCount    int    `json:"hot_count"`
	Introduce   string `json:"introduce"`
	CreatedBy   int    `json:"created_by"`
	IsDel       int    `json:"is_del"`
}

func GetNpcList(currentPage int, pageSize int, params interface{}) (npcList []NPC) {
	db.Where(params).Offset(currentPage).Limit(pageSize).Find(&npcList)
	return
}

func GetNpcTotal(params interface{}) (count int64) {
	db.Model(&NPC{}).Where(params).Count(&count)
	return
}

func CreateNpc(params interface{}) bool {
	db.Model(&NPC{}).Create(params)
	return true
}
