package electronic

import (
	"bytes"
	"image"
	"image/png"

	_ "embed"
)

//go:embed logo_electronic_960dpi.png
var LogoImage960DPIBytes []byte

var LogoImage960DPI image.Image

func init() {
	img, err := png.Decode(bytes.NewReader(LogoImage960DPIBytes))
	if err != nil {
		panic(err)
	}

	LogoImage960DPI = img
}
