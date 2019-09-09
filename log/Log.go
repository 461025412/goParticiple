package Log

import (
	"github.com/astaxie/beego/logs"
	"fmt"
	"encoding/json"
	"goParticiple/def"
)

/*
初始化日志
 */
func InitLog(configData def.Config) error {
	config := make(map[string]interface{})
	config["filename"] = configData.LogPath
	config["level"] = convertLogLevel(configData.LogLevel)

	configStr, err := json.Marshal(config)

	if err != nil {
		err = fmt.Errorf("initLogger failed, marshal err:" + err.Error())
		return err
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))
	logs.EnableFuncCallDepth(true)
	return nil
}
/**
 转换配置日志级别 为日志标准格式
 */
func convertLogLevel(level string) int {

	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}

	return logs.LevelDebug
}
