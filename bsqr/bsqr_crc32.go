package bsqr

import (
	"bytes"
	"encoding/base32"
	"encoding/binary"
	"hash/crc32"
)

var base32BSQREncoding = base32.NewEncoding("0123456789ABCDEFGHIJKLMNOPQRSTUV")

func prependBSQRChecksum(data string) bytes.Buffer {
	dataBytes := []byte(data)

	dataChecksum := crc32.ChecksumIEEE(dataBytes)

	var checksumAndData bytes.Buffer
	err := binary.Write(&checksumAndData, binary.LittleEndian, dataChecksum)
	if err != nil {
		panic(err)
	}
	checksumAndData.Write(dataBytes)

	return checksumAndData
}
