package socket5

import "time"

// Monitor 监控器接口
type Monitor interface {
	// AddTraffic 添加流量统计
	AddTraffic(traffic int64)

	// StartTime 启动时间
	StartTime() time.Time

	// TotalTraffic 总流量
	TotalTraffic() int64

	// RunningTime 运行时间
	RunningTime() time.Duration
}

// DefaultMonitor 默认监控器，不进行任何统计
type DefaultMonitor struct{}

func (m *DefaultMonitor) AddTraffic(traffic int64) {}

func (m *DefaultMonitor) StartTime() time.Time {
	return time.Time{}
}

func (m *DefaultMonitor) TotalTraffic() int64 {
	return 0
}

func (m *DefaultMonitor) RunningTime() time.Duration {
	return 0
}
