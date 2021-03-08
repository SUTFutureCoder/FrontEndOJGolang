package models

// Contest 比赛表
type Contest struct {
	Model
	// ContestName 比赛名称
	ContestName string `json:"contest_name"`
	// ContestDesc 比赛描述
	ContestDesc string `json:"contest_desc"`
	// ContestStartTime 比赛开始时间
	ContestStartTime uint64 `json:"contest_start_time"`
	// ContestEndTime 比赛结束时间
	ContestEndTime uint64 `json:"contest_end_time"`
	// SignupStartTime
	SignupStartTime uint64 `json:"signup_start_time"`
	// SignupEndTime
	SignupEndTime uint64 `json:"signup_end_time"`
}

