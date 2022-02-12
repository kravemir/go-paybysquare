package electronic

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/skip2/go-qrcode"
)

type BSQRImage struct {
	bsqr         *qrcode.QRCode
	img          image.Image
	decoratedImg image.Image
}

type ImageOptions struct {
	PointSizePX int
}

func NewBSQRImage(bsqrTextContent string, options ImageOptions) (*BSQRImage, error) {
	bsqr, err := qrcode.New(bsqrTextContent, qrcode.Low)
	if err != nil {
		return nil, fmt.Errorf("generate QR code: %w", err)
	}

	img := bsqr.Image(-options.PointSizePX)
	decoratedImg := decorateImage(img)

	return &BSQRImage{
		bsqr:         bsqr,
		img:          img,
		decoratedImg: decoratedImg,
	}, nil
}

func (i *BSQRImage) Image() image.Image {
	return i.img
}

func (i *BSQRImage) PNG() []byte {
	encoder := png.Encoder{CompressionLevel: png.BestCompression}

	var b bytes.Buffer
	err := encoder.Encode(&b, i.decoratedImg)
	if err != nil {
		panic(err)
	}

	return b.Bytes()
}

func (i *BSQRImage) WritePNGFile(filename string) error {
	pngBytes := i.PNG()

	return ioutil.WriteFile(filename, pngBytes, os.FileMode(0644))
}
