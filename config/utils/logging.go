package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	// 第２引数でファイルの作成と読み書き、追記を設定
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	// ログの書き込み先をlogfileへos.Stdoutで標準出力する
	// ログのフォーマットを指定
	// ログの出力先
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
