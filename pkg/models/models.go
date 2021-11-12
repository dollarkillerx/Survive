package models

import "time"

type Node struct {
	Name           string     `json:"name"`
	Token          string     `json:"token"`
	Timeout        *time.Time `json:"timeout"`
	TimeoutBl      bool       `json:"timeout_bl"`       // 是否超时
	TimeoutSec     int64      `json:"timeout_sec"`      // 超时sec数
	LastPingTime   string     `json:"last_ping_time"`   // 最后请求时间
	Delay          int64      `json:"delay"`            // 延迟
	TimeoutCount   int64      `json:"timeout_count"`    // 超时次数
	CPU            string     `json:"cpu"`              // cpu类型
	CPUUsedPercent string     `json:"cpu_used_percent"` // cpu使用率
	MemUsedPercent string     `json:"mem_used_percent"` // Mem使用率
	SendEmail      bool       `json:"send_email"`       // 是否发送了邮件
	IP             string     `json:"ip"`               // ip
}
