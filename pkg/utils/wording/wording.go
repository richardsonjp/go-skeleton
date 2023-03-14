package wording

import (
	"fmt"
	"go-skeleton/pkg/utils/array"
	"math"
	"math/big"
	"regexp"
	"strconv"
	"strings"

	"github.com/leekchan/accounting"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

func formatBasicWords(amount float64) string {
	switch amount {
	case 11:
		return "sebelas"
	case 10:
		return "sepuluh"
	case 9:
		return "sembilan"
	case 8:
		return "delapan"
	case 7:
		return "tujuh"
	case 6:
		return "enam"
	case 5:
		return "lima"
	case 4:
		return "empat"
	case 3:
		return "tiga"
	case 2:
		return "dua"
	case 1:
		return "satu"
	case 0:
		return "nol"
	}

	return ""
}

func formatDozenWords(amount float64) string {
	if amount >= 20 {
		division := math.Floor(amount / 10)
		remainder := float64(int(amount) % 10)
		remainderStr := ""
		if remainder > 0 {
			remainderStr = formatBasicWords(remainder)
		}

		return strings.Join(array.DeleteEmpty([]string{formatBasicWords(division), "puluh", remainderStr}), " ")
	}

	if amount >= 12 {
		return fmt.Sprintf("%s belas", formatBasicWords(float64(int(amount)%10)))
	}

	return formatBasicWords(amount)
}

func formatHundredWords(amount float64) string {
	if amount >= 100 {
		division := math.Floor(amount / 100)
		remainder := float64(int(amount) % 100)
		prefixStr := ""
		if division == 1 {
			prefixStr = "seratus"
		} else {
			prefixStr = fmt.Sprintf("%s ratus", formatBasicWords(division))
		}

		remainderStr := ""
		if remainder > 0 {
			remainderStr = formatDozenWords(remainder)
		}

		return strings.Join(array.DeleteEmpty([]string{prefixStr, remainderStr}), " ")
	}

	return formatDozenWords(amount)
}

func formatThousandWords(amount float64) string {
	if amount >= 1000 {
		division := math.Floor(amount / 1000)
		remainder := float64(int(amount) % 1000)
		prefixStr := ""
		if division == 1 {
			prefixStr = "seribu"
		} else {
			prefixStr = fmt.Sprintf("%s ribu", formatHundredWords(division))
		}

		remainderStr := ""
		if remainder > 0 {
			remainderStr = formatHundredWords(remainder)
		}

		return strings.Join(array.DeleteEmpty([]string{prefixStr, remainderStr}), " ")
	}

	return formatHundredWords(amount)
}

func formatMillionWords(amount float64) string {
	if amount >= 1000000 {
		division := math.Floor(amount / 1000000)
		remainder := float64(int(amount) % 1000000)

		prefixStr := fmt.Sprintf("%s juta", formatHundredWords(division))
		remainderStr := ""
		if remainder > 0 {
			remainderStr = formatThousandWords(remainder)
		}

		return strings.Join(array.DeleteEmpty([]string{prefixStr, remainderStr}), " ")
	}

	return formatThousandWords(amount)
}

func formatBillionWords(amount float64) string {
	if amount >= 1000000000 {
		division := math.Floor(amount / 1000000000)
		remainder := float64(int(amount) % 1000000000)

		prefixStr := fmt.Sprintf("%s miliar", formatHundredWords(division))
		remainderStr := ""
		if remainder > 0 {
			remainderStr = formatMillionWords(remainder)
		}

		return strings.Join(array.DeleteEmpty([]string{prefixStr, remainderStr}), " ")
	}

	return formatMillionWords(amount)
}

func formatTrillionWords(amount float64) string {
	if amount >= 1000000000000 {
		division := math.Floor(amount / 1000000000000)
		remainder := float64(int(amount) % 1000000000000)

		prefixStr := fmt.Sprintf("%s triliun", formatHundredWords(division))
		remainderStr := ""
		if remainder > 0 {
			remainderStr = formatBillionWords(remainder)
		}

		return strings.Join(array.DeleteEmpty([]string{prefixStr, remainderStr}), " ")
	}

	return formatBillionWords(amount)
}

func formatQuadTrillionWords(amount float64) string {
	if amount >= 1000000000000000 {
		division := math.Floor(amount / 1000000000000000)
		remainder := float64(int(amount) % 1000000000000000)

		prefixStr := fmt.Sprintf("%s kuadtriliun", formatHundredWords(division))
		remainderStr := ""
		if remainder > 0 {
			remainderStr = formatTrillionWords(remainder)
		}

		return strings.Join(array.DeleteEmpty([]string{prefixStr, remainderStr}), " ")
	}

	return formatTrillionWords(amount)
}

func formatLimit(amount float64) string {
	if amount >= 1000000000000000000 {
		return fmt.Sprint(amount)
	}

	return formatQuadTrillionWords(amount)
}

func FormatWords(amount float64) string {
	part1 := ""
	if amount < 0 {
		part1 = "negatif"
	}

	absolute := math.Floor(math.Abs(amount))
	part2 := formatLimit(absolute)

	part3 := ""
	decimals := strings.Split(fmt.Sprintf("%.2f", amount), ".")
	if len(decimals) > 1 {
		var decimalWordings []string
		decimalStr := decimals[1]
		precision, err := strconv.ParseFloat(decimalStr, 64)
		if err == nil && precision > 0 {
			for _, c := range decimalStr {
				vStr := string(c)
				vValue, err := strconv.ParseFloat(vStr, 64)
				if err == nil {
					decimalWordings = append(decimalWordings, formatBasicWords(vValue))
				}
			}
		}

		if len(decimalWordings) > 0 {
			part3 = fmt.Sprintf("koma %s", strings.Join(decimalWordings, " "))
		}
	}

	return strings.Join(array.DeleteEmpty([]string{part1, part2, part3}), " ")
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

//08 -> 62
func NormalizePhoneNumber(phone string) string {
	phone = strings.TrimSpace(phone)

	if len(phone) > 2 && phone[0:1] == "0" {
		phone = strings.Replace(phone, "0", "62", 1)
	}
	return phone
}

//62 -> 08
func DenormalizePhoneNumber(phone string) string {
	phone = strings.TrimSpace(phone)
	phone = strings.Replace(phone, "+62", "0", 1)
	phone = strings.Replace(phone, "+", "", 1)

	if len(phone) > 2 && phone[0:2] == "62" {
		phone = strings.Replace(phone, "62", "0", 1)
	}
	return phone
}

func FormatIDRCurrency(value *float64) string {
	currencyCode := "IDR"
	f := big.NewFloat(*value)
	cur, err := currency.ParseISO(currencyCode)
	if err != nil {
		return fmt.Sprintf("%s%s", currencyCode, accounting.FormatNumberBigFloat(f, 2, ",", "."))
	}

	scale, _ := currency.Cash.Rounding(cur) // fractional digits
	unit, _ := f.Float64()
	dec := number.Decimal(unit, number.Scale(scale))

	p := message.NewPrinter(language.English)
	if currencyCode == "IDR" {
		p = message.NewPrinter(language.Indonesian)
	}

	return p.Sprintf("%v%v", currency.Symbol(cur), dec)
}
