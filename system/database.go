package system

import (
	"time"
)

// User 用户信息结构体
type User struct {
	ID       int
	Username string
	Device   string
}

// Subscription 订阅信息结构体
type Subscription struct {
	UserID    int
	Channel   string
	MsgType   string
	CreatedAt time.Time
}

// PushHistory 推送历史记录结构体
type PushHistory struct {
	ID         int
	UserID     int
	Channel    string
	MsgContent string
	SentAt     time.Time
}

// GetUserSubscriptions 模拟从数据库中获取用户的订阅信息
func GetUserSubscriptions(userID int) []Subscription {
	// ...
}

// RecordPushHistory 模拟将消息发送历史记录到数据库
func RecordPushHistory(userID int, msgContent string) error {
	// ...
}
