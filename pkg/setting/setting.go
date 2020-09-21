package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	LogSavePath string
	LogSaveName string
	LogFileExt  string
}

var AppSetting = &App{}

type Judger struct {
	TestChamberBaseDir string
}

var JudgerSetting = &Judger{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type         string
	User         string
	Password     string
	Host         string
	Name         string
	TablePrefix  string
	MaxOpenConns int
	MaxIdleConns int
}

var DatabaseSetting = &Database{}

type Session struct {
	Token       string
	SessionUser string
}

var SessionSetting = &Session{}

type Tool struct {
	PicBaseDir string
}

var ToolSetting = &Tool{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("judger", JudgerSetting)
	mapTo("session", SessionSetting)
	mapTo("tool", ToolSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

func mapTo(s string, v interface{}) {
	err := cfg.Section(s).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", s, err)
	}
}
