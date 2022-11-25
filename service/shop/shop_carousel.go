package shop

import "go-shop-api/model/shop"

type ShopCarouselService struct {
}

func (c *ShopCarouselService) GetCarouselList() (carousel []*shop.ShopCarousel, err error) {
	return shop.GetCarousel()
}
