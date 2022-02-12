package bsqrbuilder

import (
	"regexp"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type OneTimePayment struct {
	IBAN string
	BIC  string

	Amount       string
	CurrencyCode string
	DueDate      time.Time

	VariableSymbol string
	ConstantSymbol string
	SpecificSymbol string

	Note string

	BeneficiaryName         string
	BeneficiaryAddressLine1 string
	BeneficiaryAddressLine2 string
}

func (p OneTimePayment) BuildSequenceDataModelText() string {
	return strings.Join(
		[]string{
			"",
			"1",
			"1",
			p.Amount,
			p.CurrencyCode,
			p.DueDate.Format("20060102"),
			p.VariableSymbol,
			p.ConstantSymbol,
			p.SpecificSymbol,
			"",
			safeString(p.Note),
			"1", // one BankAccounts
			p.IBAN,
			p.BIC,
			"0", // not StandingOrderExt
			"0", // not DirectDebitExt
			safeString(p.BeneficiaryName),
			safeString(p.BeneficiaryAddressLine1),
			safeString(p.BeneficiaryAddressLine2),
		},
		"\t",
	)
}

func safeString(text string) string {
	text = removeDiacritics(text)
	text = removeUnsafeCharacters(text)
	return text
}

func removeUnsafeCharacters(text string) string {
	return regexp.MustCompile("[^-a-zA-Z0-9 .:]").ReplaceAllString(text, "")
}

func removeDiacritics(text string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	text, _, _ = transform.String(t, text)
	return text
}
