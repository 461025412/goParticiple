package bootstrap

import (
	"fmt"
	"goParticiple/def"
	"goParticiple/log"
	configRead "github.com/astaxie/beego/config"
)

// 全局配置信息变量
var ConfigData = def.Config{}

func InitService(configPath string) error {
	//加载配置文件信息
	err := InitConfig("ini", configPath)
	if err != nil {
		return err
	}
	//初始化日志模块
	err = Log.InitLog(ConfigData)
	if err != nil {
		return err
	}

	return nil
}

/**
加载配置信息
*/
func InitConfig(confType string, filename string) error {
	conf, err := configRead.NewConfig(confType, filename)

	if err != nil {
		err = fmt.Errorf("initLogger failed, marshal err:" + err.Error())
		return err
	}

	ConfigData.ConfigPath = filename
	//	加载日志路径
	log_path := conf.String("Log::log_path")
	if len(log_path) == 0 {
		log_path = "./IpParser.log"
	}
	ConfigData.LogPath = log_path

	// 加载日志级别
	log_level := conf.String("Log::log_level")
	if len(log_level) == 0 {
		log_level = "debug"
	}
	ConfigData.LogLevel = log_level

	// 分词word
	wordFilePath := conf.String("WordFilePath::word_file_path")
	if len(wordFilePath) == 0 {
		wordFilePath = "../data/word.dic"
	}
	ConfigData.WordPath = wordFilePath

	// term depth
	termDepth, err := conf.Int("TermDepth::term_depth")
	if err != nil {
		termDepth = 2
	}

	ConfigData.TermDepth = termDepth

	// update interval
	updateInterval, err := conf.Int64("UpdateInterval::update_interval")
	if err != nil {
		updateInterval = 60
	}

	ConfigData.UpdateInterval = updateInterval

	// http port 监听
	httpPort, err := conf.Int("HttpPort::http_port")
	if err != nil {
		httpPort = 8080
	}

	ConfigData.HttpPort = httpPort

	return nil
}
