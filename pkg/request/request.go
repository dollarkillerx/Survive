package request

type HeartbeatRequest struct {
	Time           int64  `json:"time"`             // 发送请求时间 毫秒
	CPU            string `json:"cpu"`              // cpu类型
	CPUUsedPercent string `json:"cpu_used_percent"` // cpu使用率
	MemUsedPercent string `json:"mem_used_percent"` // Mem使用率
}
