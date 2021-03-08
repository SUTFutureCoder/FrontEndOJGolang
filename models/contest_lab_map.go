package models

// ContestLabMap 比赛实验室关联表
type ContestLabMap struct {
	Model
	// ContestId 比赛Id
	ContestId uint64 `json:"contest_id"`
	// LabId 实验室Id
	LabId uint64 `json:"lab_id"`
}


