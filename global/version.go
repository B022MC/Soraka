package global

var (
	APPVersion = "0.1.0"
	Commit     = "dev"
	BuildTime  = ""
	BuildUser  = ""
)

func init() {
	SetAppInfo(AppInfo{
		Version:   APPVersion,
		Commit:    Commit,
		BuildUser: BuildUser,
		BuildTime: BuildTime,
	})
}
