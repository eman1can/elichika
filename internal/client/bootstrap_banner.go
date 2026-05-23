package client

import "elichika/internal/generic"

type BootstrapBanner struct {
	Banners generic.Array[Banner1] `json:"banners"`
}
