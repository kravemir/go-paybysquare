package bsqr

import (
	"bytes"

	"github.com/ulikunitz/xz/lzma"
)

func compressBSQRLZMA(checksumAndData bytes.Buffer) ([]byte, error) {
	lzmaConfig := lzma.WriterConfig{
		Properties: &lzma.Properties{
			LC: 3,
			LP: 0,
			PB: 2,
		},
		DictCap: 128 * 1024,
	}

	var buf bytes.Buffer
	lzmaWriter, err := lzmaConfig.NewWriter(&buf)
	if err != nil {
		panic(err)
	}
	_, err = lzmaWriter.Write(checksumAndData.Bytes())
	if err != nil {
		panic(err)
	}
	err = lzmaWriter.Close()
	if err != nil {
		panic(err)
	}
	return buf.Bytes()[13:], nil
}
