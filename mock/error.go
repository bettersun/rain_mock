package mock

// LogWarn
//  错误非nil时输出日志
func LogWarn(err error) {
	if err != nil {
		logger.Warn(err)
	}
}
