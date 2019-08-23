package pkg

import "time"

// PriceItem TBA
type PriceItem struct {
	Code      string
	Title     string
	SellPrice uint64
	BuyPrice  uint64
	UpdatedAt time.Time
}

// PriceList TBA
type PriceList []PriceItem

// PriceProviderConfig TBA
type PriceProviderConfig map[string]string

// PriceProvider TBA
type PriceProvider interface {
	GetAdapterName() string
	HealthCheck() (bool, error)

	GetPriceList() (*PriceList, error)
}
