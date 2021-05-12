package config

import (
	"log"

	"example.com/config/utils"
	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
}

// ConfigListを外部から呼び出せるようグローバルで宣言する
var Config ConfigList

// main()関数より先に実行されるinit()関数を定義して ConfigList を読み込む
func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	// config.iniに設定した値を読み込む
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("3000"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
}
