package Soraka

import "Soraka/global"

var (
	APPVersion = "0.1.0"
	Commit     = "dev"
	BuildTime  = ""
	BuildUser  = ""
)

func init() {
	global.SetAppInfo(global.AppInfo{
		Version:   APPVersion,
		Commit:    Commit,
		BuildUser: BuildUser,
		BuildTime: BuildTime,
	})
}
