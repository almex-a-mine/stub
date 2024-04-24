package infrastructure

import (
	"stub/config"
	"stub/domain"
	"stub/domain/handler"
	"sync"
)

var wgrps sync.WaitGroup

func Router() {
	wgrps = sync.WaitGroup{}
	wgrps.Add(1)
	//設定取得
	config := config.Initialize(domain.SrvName)
	//ログハンドラ作成
	logger := NewLogger(config.SystemConf.MaxLength,
		config.SystemConf.MaxRotation,
		config.SystemConf.LogStopInfo,
		config.SystemConf.LogStopTrace,
		config.SystemConf.LogStopMqtt,
		config.SystemConf.LogStopDebug,
		config.SystemConf.LogStopMutex,
		config.SystemConf.LogStopWarn,
		config.SystemConf.LogStopError,
		config.SystemConf.LogStopFatal,
		config.SystemConf.LogStopSequence)
	logger.Info("Program Start")
	logger.Info("config %+v", config)
	//MQTTハンドラ作成
	mqtt := NewMQTTHandler(logger, config.MqttConf.TCP, config.MqttConf.Port, domain.SrvName)
	sysMqtt := NewMQTTHandler(logger, config.MqttConf.TCP, config.MqttConf.Port, domain.SrvName+"_sys")

	// Start Contoller
	Controller(mqtt, sysMqtt, logger, config)
}

// 停止処理
func RouterStop() {
	wgrps.Done()
}

func Controller(mqtt handler.MqttRepository, sysMqtt handler.MqttRepository, logger handler.LoggerRepository, config config.Configuration) {

}
