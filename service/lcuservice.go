package service

import "Soraka/biz/lcu"

// LcuService exposes LCU related helpers.
type LcuService struct{}

func (LcuService) CheckLogin() bool {
	return lcu.CheckLogin()
}

func (LcuService) GetCredentials() (string, string, error) {
	return lcu.GetCredentials()
}
