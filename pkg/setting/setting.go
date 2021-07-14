package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg 			*ini.File
	RunMode 		string
	HTTPort			int
	JWTSecret		string
	PageSize		int
	ReadTimeout 	time.Duration
	WriteTimeout	time.Duration
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	LoadBase()
	LoadServer()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
	PageSize = Cfg.Section("").Key("PAGE_SIZE").MustInt(10)
	JWTSecret = Cfg.Section("").Key("JWT_SECRET").MustString("sldkjfl")
}

func LoadServer() {
	server, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get server: %v", err)
	}

	HTTPort = server.Key("HTTP_PORT").MustInt(8080)
	ReadTimeout = time.Duration(server.Key("ReadTimeout").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(server.Key("WriteTimeout").MustInt(60)) * time.Second
}