package main

import (
	"fmt"
	"os"
	"time"

	bsqr2 "github.com/kravemir/go-paybysquare/bsqr"
	"github.com/kravemir/go-paybysquare/bsqrbuilder"
	"github.com/kravemir/go-paybysquare/image/electronic"
)

const OUTPUTS_DIR = "output/examples"

func main() {
	ensureOutputsDirectory()

	oneTimePaymentExample()
}

func oneTimePaymentExample() {
	NixOSDonation := bsqrbuilder.OneTimePayment{
		IBAN: "NL13 BUNQ 2061 2530 32",
		BIC:  "BUNQNL2AXXX",

		Amount:       "5.00",
		CurrencyCode: "EUR",
		DueDate:      time.Now(),

		VariableSymbol: "",
		ConstantSymbol: "",
		SpecificSymbol: "",

		Note: "NixOS: humble donation",

		BeneficiaryName:         "Stichting NixOS Foundation",
		BeneficiaryAddressLine1: "",
		BeneficiaryAddressLine2: "Frederiksoord",
	}

	sequenceData := NixOSDonation.BuildSequenceDataModelText()

	bsqrTextContent, err := bsqr2.EncodeSequence(sequenceData)
	assertNilError(err)

	bsqr, err := electronic.NewBSQRImage(bsqrTextContent, electronic.ImageOptions{
		PointSizePX: 3,
	})
	assertNilError(err)

	err = bsqr.WritePNGFile(exampleOutputFilePath("one-time-payment.png"))
	assertNilError(err)
}

func exampleOutputFilePath(filename string) string {
	return fmt.Sprintf("%s/%s.png", OUTPUTS_DIR, filename)
}

func ensureOutputsDirectory() {
	os.RemoveAll(OUTPUTS_DIR)
	if _, err := os.Stat(OUTPUTS_DIR); os.IsNotExist(err) {
		err := os.MkdirAll(OUTPUTS_DIR, 0775)
		if err != nil {
			panic(err)
		}
	}
}

func assertNilError(err error) {
	if err != nil {
		panic(err)
	}
}
