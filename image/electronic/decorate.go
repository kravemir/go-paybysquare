package electronic

import (
	"image"
	"image/color"

	"golang.org/x/image/draw"
)

func decorateImage(bsqrImage image.Image) image.Image {
	bsqrImageWidth := bsqrImage.Bounds().Dx()
	bsqrImageHeight := bsqrImage.Bounds().Dy()

	logoWidth := (bsqrImageWidth * 3) / 4
	logoHeight := int((int64(LogoImage960DPI.Bounds().Dy()) * int64(logoWidth)) / int64(LogoImage960DPI.Bounds().Dx()))

	resultImage := image.NewRGBA(image.Rect(
		0,
		0,
		bsqrImageWidth,
		bsqrImageHeight+logoHeight,
	))

	draw.Draw(
		resultImage,
		resultImage.Bounds(),
		&image.Uniform{color.RGBA{255, 255, 255, 255}},
		image.Point{},
		draw.Src,
	)

	draw.Copy(
		resultImage,
		image.Pt(0, 0),
		bsqrImage,
		bsqrImage.Bounds(),
		draw.Src,
		nil,
	)

	draw.BiLinear.Scale(
		resultImage,
		image.Rect(
			bsqrImageWidth-logoWidth,
			bsqrImageHeight,
			bsqrImageWidth,
			bsqrImageHeight+logoHeight,
		),
		LogoImage960DPI,
		LogoImage960DPI.Bounds(),
		draw.Over,
		nil,
	)

	return resultImage
}
