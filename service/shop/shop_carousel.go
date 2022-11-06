package shop

import "go-admin/model/shop"

type CarouselService struct {
}

func (c *CarouselService) GetCarouselList() (carousel []*shop.ShopCarousel, err error) {
	return shop.GetCarousel()
}
