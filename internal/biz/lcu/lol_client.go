package lcu

import (
	"errors"
	"regexp"
	"strconv"

	"Soraka/pkg/windows/process"
)

const (
	lolUxProcessName = "LeagueClientUx.exe"
)

var (
	lolCommandlineReg     = regexp.MustCompile(`--remoting-auth-token=(.+?)" ".*?--app-port=(\d+)"`)
	ErrLolProcessNotFound = errors.New("没有找到lol客户端进程")
)

type ClientInfoUseCase struct{}

func NewClientInfoUseCase() *ClientInfoUseCase {
	return &ClientInfoUseCase{}
}

// 获取 LCU API 信息
func (c *ClientInfoUseCase) GetLolClientApiInfo() (port int, token string, err error) {
	cmdline, err := process.GetProcessCommand(lolUxProcessName)
	if err != nil {
		return 0, "", ErrLolProcessNotFound
	}

	btsChunk := lolCommandlineReg.FindSubmatch([]byte(cmdline))
	if len(btsChunk) < 3 {
		return 0, "", ErrLolProcessNotFound
	}

	token = string(btsChunk[1])
	port, err = strconv.Atoi(string(btsChunk[2]))
	return
}
