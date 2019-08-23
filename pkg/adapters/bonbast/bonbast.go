package bonbast

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gocolly/colly"
	"github.com/hashicorp/go-multierror"
	"github.com/mamal72/stringish"

	"github.com/mamal72/arzdoon/pkg"
)

const address = "https://www.bonbast.com/"
const timeLayout = "Fri, 23 Aug 2019 00:01:01"

// Adapter TBA
type Adapter struct {
	collector *colly.Collector
}

// GetAdapterName TBA
func (a *Adapter) GetAdapterName() string {
	return "Bonbast (https://www.bonbast.com/)"
}

// New TBA
func New() (pkg.PriceProvider, error) {
	adapter := &Adapter{}

	collector := colly.NewCollector()

	adapter.collector = collector
	return adapter, nil
}

// GetPriceList TBA
func (a *Adapter) GetPriceList() (*pkg.PriceList, error) {
	priceList := pkg.PriceList{}
	var errors error

	a.collector.OnHTML("body", func(body *colly.HTMLElement) {
		updatedAt, err := time.Parse(time.RFC1123, fmt.Sprintf("%s UTC", body.ChildText(".miladi.utc")))
		if err != nil {
			errors = multierror.Append(errors, err)
		}

		body.ForEach(".col-xs-12 table.table-condensed", func(_ int, table *colly.HTMLElement) {
			table.ForEach("tr:not(:first-of-type)", func(_ int, priceRow *colly.HTMLElement) {
				sellPrice, err := strconv.ParseUint(priceRow.DOM.Children().Eq(2).Text(), 10, 64)
				if err != nil {
					errors = multierror.Append(errors, err)
				}

				buyPrice, err := strconv.ParseUint(priceRow.DOM.Children().Eq(3).Text(), 10, 64)
				if err != nil {
					errors = multierror.Append(errors, err)
				}

				title := stringish.New(priceRow.DOM.Children().Eq(1).Text()).Filter(func(char string) bool {
					if char == " " {
						return true
					}

					if char >= "A" && char <= "z" {
						return true
					}

					return false
				}).TrimSpaces().GetString()

				priceItem := pkg.PriceItem{
					Code:      priceRow.DOM.Children().Eq(0).Text(),
					Title:     title,
					SellPrice: sellPrice,
					BuyPrice:  buyPrice,
					UpdatedAt: updatedAt,
				}
				priceList = append(priceList, priceItem)
			})
		})
	})

	err := a.collector.Visit(address)
	if err != nil {
		errors = multierror.Append(errors, err)
	}

	return &priceList, errors
}

// HealthCheck TBA
func (a *Adapter) HealthCheck() (bool, error) {
	_, err := http.Get(address)
	if err != nil {
		return false, err
	}

	return true, nil
}
