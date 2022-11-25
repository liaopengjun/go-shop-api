package request

type AddUserAddressParam struct {
	UserName      string `json:"userName"`
	UserPhone     string `json:"userPhone"`
	ProvinceName  string `json:"provinceName"`
	CityName      string `json:"cityName"`
	RegionName    string `json:"regionName"`
	DetailAddress string `json:"detailAddress"`
	DefaultFlag   int    `json:"defaultFlag"`
}

type EditUserAddressParam struct {
	AddressId     int    `json:"addressId"`
	UserName      string `json:"userName"`
	UserPhone     string `json:"userPhone"`
	ProvinceName  string `json:"provinceName"`
	CityName      string `json:"cityName"`
	RegionName    string `json:"regionName"`
	DetailAddress string `json:"detailAddress"`
	DefaultFlag   int    `json:"defaultFlag"`
}
