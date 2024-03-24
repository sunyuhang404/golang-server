package util

import (
	"fmt"
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

func Setup() {
	var err error

	Cfg, err := ini.Load("src/config/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'config/app.ini': %v", err)
	}

	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

	println(Cfg)

	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	Port = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

	sec2, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
	JwtSecret = sec2.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec2.Key("PAGE_SIZE").MustInt(10)

	sec3, err := Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'database': %v", err)
	}

	DBType = sec3.Key("TYPE").MustString("mysql")
	DBName = sec3.Key("NAME").String()
	DBUser = sec3.Key("USER").String()
	DBPassword = sec3.Key("PASSWORD").String()
	DBHost = sec3.Key("HOST").String()

	//LoadServer()
	//LoadApp()
	//LoadDatabase()
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	fmt.Println(sec)
	//Port = sec.Key("HTTP_PORT").MustInt(8000)
	//ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	//WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
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
