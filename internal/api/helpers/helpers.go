package helpers

import "github.com/leekchan/accounting"

var (
	lc = accounting.LocaleInfo["IDR"]
	rp = accounting.Accounting{Symbol: "Rp. ",
		Precision: 0,
		Thousand:  lc.ThouSep,
		Decimal:   lc.DecSep,
	}
)

func FormatMoney(price int) string {
	return rp.FormatMoney(price)
}
