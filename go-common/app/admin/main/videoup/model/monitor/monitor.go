package monitor

import (
	"fmt"
	"go-common/app/admin/main/videoup/model/archive"
	"go-common/app/admin/main/videoup/model/manager"
)

const (
	// RedisPrefix 参数:business。
	BusPrefix = "monitor_stats_%d_"
	// SuffixVideo 视频停留统计。参数:state
	SuffixVideo = "%d"
	// SuffixArc 稿件停留统计。参数:round。参数:state。
	SuffixArc    = "%d_%d"
	RulesKey     = "monitor_rules"
	RuleIDIncKey = "monitor_rule_inc_id"

	BusVideo int8 = 1
	BusArc   int8 = 2

	RuleStateOK      = 1
	RuleStateDisable = 0

	NotifyTypeEmail = 1
	NotityTypeSms   = 2

	CompGT  = ">"
	CompLT  = "<"
	CompGET = ">="
	CompLET = "<="
	CompNE  = "!="
	CompE   = "="
)

var (
	// RedisKeyConf 监控业务的Redis Key配置
	RedisKeyConf = map[int8]*KeyConf{
		BusVideo: {
			KeyFormat: fmt.Sprintf(BusPrefix, BusVideo) + SuffixVideo,
			KFields: map[string][]int64{
				"state": {
					int64(archive.VideoStatusSubmit),
					int64(archive.VideoStatusWait),
				},
			},
		},
		BusArc: {
			KeyFormat: fmt.Sprintf(BusPrefix, BusArc) + SuffixArc,
			KFields: map[string][]int64{
				"round": {
					int64(archive.RoundAuditSecond),
					int64(archive.RoundAuditThird),
					int64(archive.RoundReviewFlow),
					int64(archive.RoundReviewFirst),
					int64(archive.RoundReviewFirstWaitTrigger),
					int64(archive.RoundReviewSecond),
					int64(archive.RoundTriggerClick),
				},
				"state": {
					int64(archive.StateForbidFixed),
					int64(archive.StateForbidWait),
					int64(archive.StateOpen),
					int64(archive.StateOrange),
				},
			},
		},
	}
)

type KeyConf struct {
	KeyFormat string
	KFields   map[string][]int64
}

type RuleResultParams struct {
	Type     int8 `form:"type" validate:"required"`
	Business int8 `form:"business" validate:"required"`
	//STime    utils.FormatTime `form:"stime"`
	//ETime    utils.FormatTime `form:"etime"`
}
type RuleResultData struct {
	Rule  *Rule         `json:"rule"`
	User  *manager.User `json:"user"`
	Stats *Stats        `json:"stats"`
}

// Rule 监控规则信息
type Rule struct {
	ID       int64     `json:"id"`
	Type     int8      `json:"type"`
	Business int8      `json:"business"`
	Name     string    `json:"name"`
	State    int8      `json:"state"`
	STime    string    `json:"s_time"`
	ETime    string    `json:"e_time"`
	RuleConf *RuleConf `json:"rule_conf"`
	UID      int64     `json:"uid"`
	CTime    string    `json:"ctime"`
}

// RuleConf 监控方案配置结构体
type RuleConf struct {
	Name    string              `json:"name"`
	MoniCdt map[string]struct { //监控方案的监控条件
		Comp  string `json:"comparison"`
		Value int64  `json:"value"`
	} `json:"moni_cdt"`
	NotifyCdt map[string]struct { //达到发送通知的条件
		Comp  string `json:"comparison"`
		Value int64  `json:"value"`
	} `json:"notify_cdt"`
	Notify struct { //通知类型配置
		Way    int8     `json:"way"`
		Member []string `json:"member"`
	} `json:"notify"`
}

type Stats struct {
	TotalCount int `json:"total_count"`
	MoniCount  int `json:"moni_count"`
	MaxTime    int `json:"max_time"`
}
