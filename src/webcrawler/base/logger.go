package base

import "github.com/Jsoneft/goc2p/src/logging"

// 创建日志记录器。
func NewLogger() logging.Logger {
	return logging.NewSimpleLogger()
}
