package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	// ログファイルの読み込み、作成など
	logfile,err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	// ログの書き込み先
	multiLogFile := io.MultiWriter(os.Stdout,logfile)
	// ログフォーマット
	log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)
	log.SetOutput(multiLogFile)
}