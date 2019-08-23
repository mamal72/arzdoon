package main

import (
	"github.com/mamal72/arzdoon/pkg/adapters/bonbast"
	"github.com/mamal72/arzdoon/pkg/utils"
)

func main() {
	provider, _ := bonbast.New()
	prices, _ := provider.GetPriceList()
	utils.PrintPriceTable(provider.GetAdapterName(), prices)
}
