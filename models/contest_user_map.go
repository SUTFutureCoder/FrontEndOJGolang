package models

// ContestUserMap 比赛用户关联表
type ContestUserMap struct {
	Model
	// ContestId 比赛Id
	ContestId uint64 `json:"contest_id"`
}


