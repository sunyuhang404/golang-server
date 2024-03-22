package util

import (
	"gopkg.in/ini.v1"
	"log"
	"time"
)

var (
	Cfg *ini.File

	RunMode      string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string

	// database
	DBType     string
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
)

func Init() {
	var err error

	Cfg, err := ini.Load("src/config/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'config/app.ini': %v", err)
	}

	println(Cfg)

	LoadBase()
	LoadServer()
	LoadApp()
	LoadDatabase()
}

func LoadBase() {
	sec, err := Cfg.GetSection("mode")

	if err != nil {
		log.Fatalf("Fail to get section 'mode': %v", err)
	}
	RunMode = sec.Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	Port = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

func LoadDatabase() {
	sec, err := Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'database': %v", err)
	}

	DBType = sec.Key("TYPE").MustString("mysql")
	DBName = sec.Key("NAME").String()
	DBUser = sec.Key("USER").String()
	DBPassword = sec.Key("PASSWORD").String()
	DBHost = sec.Key("HOST").String()
}
