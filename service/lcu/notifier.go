package lcu

import (
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
)

// StartNotifier emits LCU status, credentials and summoner info events.
func StartNotifier(app *application.App) {
	go func() {
		var lastStatus bool
		var lastPort int
		var lastToken string
		for {
			port, token, err := GetLolClientApiInfo()
			status := err == nil
			if status != lastStatus {
				app.EmitEvent("lcuStatus", status)
				lastStatus = status
			}
			if status {
				if port != lastPort || token != lastToken {
					app.EmitEvent("lcuCreds", AuthInfo{Port: port, Token: token})
					InitCli(port, token)
					if info, err := (WailsAPI{}).GetCurrentSummoner(); err == nil {
						app.EmitEvent("summonerInfo", info)
					}
					lastPort = port
					lastToken = token
				}
			}
			time.Sleep(time.Second)
		}
	}()
}
