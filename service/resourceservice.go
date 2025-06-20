package service

import (
	"Soraka/service/utils"
)

// ResourceService handles retrieving LoL static resources.
type ResourceService struct{}

// FetchHeroList downloads the hero list js from Tencent's CDN.
func (ResourceService) FetchHeroList() (string, error) {
	const url = "https://game.gtimg.cn/images/lol/act/img/js/heroList/hero_list.js"
	data := utils.Get(url)
	return data, nil
}
