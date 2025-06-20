package service

import "Soraka/biz/client"

// ClientService exposes League of Legends client related functions to the frontend.
type ClientService struct{}

func (ClientService) GetClientPath() string {
	return client.GetClientPath()
}

func (ClientService) StartClient() string {
	if err := client.Start(); err != nil {
		return err.Error()
	}
	return ""
}

func (ClientService) GetConcurrency() int {
	return client.GetConcurrency()
}

func (ClientService) SaveConcurrency(c int) {
	client.SaveConcurrency(c)
}
