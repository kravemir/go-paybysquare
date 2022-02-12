package bsqr

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func EncodeSequence(plainSequence string) (string, error) {
	checksumAndSequenceData := prependBSQRChecksum(plainSequence)

	compressedData, err := compressBSQRLZMA(checksumAndSequenceData)
	if err != nil {
		return "", fmt.Errorf("LZMA compression failed: %w", err)
	}

	var completeOutput bytes.Buffer
	completeOutput.WriteByte(0)
	completeOutput.WriteByte(0)
	err = binary.Write(&completeOutput, binary.LittleEndian, uint16(len(checksumAndSequenceData.Bytes())))
	if err != nil {
		panic(err)
	}
	completeOutput.Write(compressedData)

	resultToQRCode := base32BSQREncoding.EncodeToString(completeOutput.Bytes())
	return resultToQRCode, nil
}
