package mock

import "net/http"

// Config 配置
type Config struct {
	Port     string `yaml:"port"`     // 工具运行端口
	LogFile  string `yaml:"logFile"`  // 日志文件
	LogLevel int    `yaml:"logLevel"` // 日志等级

	Mock ProxyConfig `yaml:"mock"` // Mock 配置
	Path PathConfig  `yaml:"path"` // 路径配置
}

// ProxyConfig 代理配置
type ProxyConfig struct {
	DestHost     string `yaml:"destHost"`     // 目标主机(包含 协议://IP:端口 )
	MockItemFile string `yaml:"mockItemFile"` // Mock项目文件

	// 使用模拟服务通用响应头 仅当URL对应的响应头不存在时使用
	UseCommonHeader bool `yaml:"useCommonHeader"`
}

// PathConfig 路径配置
type PathConfig struct {
	Request  string `yaml:"request"`  // 保存请求信息的相对位置（相对于工具所在目录）
	Response string `yaml:"response"` // 保存响应信息的相对位置（相对于工具所在目录）
	Backup   string `yaml:"backup"`   // 备份 Mock 项目的目录（相对于工具所在目录）

	CommonHeaderFile   string `yaml:"commonHeaderFile"`   // 通用响应头部文件
	ResponseHeaderFile string `yaml:"responseHeaderFile"` // 响应头部信息文件
}

// MockItem Mock 项目
type MockItem struct {
	Path         string      `yaml:"path" json:"path"`                 // 请求API的路径
	Method       string      `yaml:"method" json:"method"`             // 请求方法
	UseMock      bool        `yaml:"useMock" json:"useMock"`           // 使用 Mock
	DestHost     string      `yaml:"destHost" json:"destHost"`         // 目标主机
	Duration     int         `yaml:"duration" json:"duration"`         // 响应等待时间（毫秒）
	StatusCode   int         `yaml:"statusCode" json:"statusCode"`     // 响应状态码
	Header       http.Header `yaml:"-" json:"-"`                       // 响应头部
	ResponseFile string      `yaml:"responseFile" json:"responseFile"` // 响应文件
	Description  string      `yaml:"description" json:"description"`   // 说明
}
