package service

import "Soraka/biz/lcu"

// LcuService exposes LCU related helpers.
type LcuService struct{}

// CheckLogin returns the login status along with LCU credentials.
func (LcuService) CheckLogin() (bool, string, string) {
	return lcu.CheckLogin()
}

func (LcuService) GetCredentials() (string, string, error) {
	return lcu.GetCredentials()
}
