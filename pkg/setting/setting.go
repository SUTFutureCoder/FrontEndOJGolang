package setting

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-ini/ini"
)

const FRONTENDOJ = "FrontEndOJGolang"

type App struct {
	LogSavePath string
	LogSaveName string
	LogFileExt  string
}

var AppSetting = &App{}

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
	FileBaseDir string
}

var ToolSetting = &Tool{}

type Judger struct {
	JudgerAddr string
	HttpPort string
}

var JudgerSetting = &Judger{}

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
	mapTo("session", SessionSetting)
	mapTo("tool", ToolSetting)
	mapTo("judger", JudgerSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

func mapTo(s string, v interface{}) {
	err := cfg.Section(s).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", s, err)
	}
}

func Check() {
	checkStaticFileDir()
}

func checkStaticFileDir() {
	ToolSetting.FileBaseDir = checkAndFixDirExists(ToolSetting.FileBaseDir, "static/file")
}

func checkAndFixDirExists(targetDir string, suffix string) string {
	_, err := os.Stat(targetDir)
	if err == nil || os.IsExist(err) {
		return targetDir
	}
	// try create dir if not exist
	err = os.MkdirAll(targetDir, 0777)
	if err == nil {
		return targetDir
	}

	// failed to create dir but try user home dir
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Can not create target dir and your home dir, please check your config in conf/*.ini was set correctly.")
	}
	newDir := fmt.Sprintf("%s/%s", homeDir, suffix)
	// try create dir if not exist
	err = os.MkdirAll(newDir, 0777)
	if err == nil {
		return newDir
	}
	log.Fatalf("Can not create home dir, please check your config in conf/*.ini was set correctly.")
	return ""
}