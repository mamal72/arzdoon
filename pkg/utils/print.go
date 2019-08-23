package utils

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"

	"github.com/mamal72/arzdoon/pkg"
)

// PrintPriceTable TBA
func PrintPriceTable(title string, list *pkg.PriceList) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Code", "Title", "Sell", "Buy", "Updated At"})

	for index, priceItem := range *list {
		table.Append([]string{
			fmt.Sprintf("%d", index),
			priceItem.Code,
			priceItem.Title,
			fmt.Sprintf("%d", priceItem.SellPrice),
			fmt.Sprintf("%d", priceItem.BuyPrice),
			priceItem.UpdatedAt.Format("Mon Jan 2 15:04:05"),
		})
	}

	fmt.Printf("Provider: %s\n", title)
	table.Render()
}
