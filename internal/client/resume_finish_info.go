package client

import "elichika/internal/generic"

type ResumeFinishInfo struct {
	CachedJudgeResult generic.Dictionary[int32, int32] `json:"cached_judge_result" enum:"JudgeType"`
}
