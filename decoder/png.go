package png

import (
	"errors"
	"fmt"
)

type Png struct {
	pngBytes  [8]byte // PNG file signature
	length    uint32  // Length of the data field
	chunkType [4]byte // The type of the chunk
	chunkData []byte  // The data bytes are preprocessed to have no zero bytes at the end
	crc       [4]byte // Cyclic redundancy check
}

func New() *Png {
	return &Png{}
}

func (png *Png) Decode(content []byte) (string, error) {
	// read first 8 byte
	eightByte := content[:8]
	err := png.isPngFile(eightByte)
	if err != nil {
		return "", err
	}
	// remove first 8 byte
	content = content[8:]
	png.readLength(content[:4])
	content = content[4:]

	// read chunk chunkType
	png.chunkType = [4]byte{content[0], content[1], content[2], content[3]}
	content = content[4:]

	// read chunk chunkData
	png.chunkData = content[:int(png.length)]
	content = content[int(png.length):]

	// read chunk crc
	png.crc = [4]byte{content[0], content[1], content[2], content[3]}
	content = content[4:]

	fmt.Println("PNG Signature: ", png.pngBytes)
	fmt.Println("Length: ", png.length)
	fmt.Println("ChunkType: ", png.chunkType)
	fmt.Println("ChunkData: ", png.chunkData)
	fmt.Println("CRC: ", png.crc)
	png.decodeData()

	return "PNG file decoded", nil
}

func (png *Png) readLength(content []byte) {
	for i := 0; i < 4; i++ {
		png.length += uint32(content[i]) << (24 - 8*i)
	}
}

func (png *Png) isPngFile(content []byte) error {
	pngBytes := []byte{137, 80, 78, 71, 13, 10, 26, 10}
	for i := 0; i < len(pngBytes); i++ {
		if content[i] != pngBytes[i] {
			return errors.New("File is not a PNG")
		}
		png.pngBytes[i] = content[i]
	}

	return nil
}

func (png *Png) decodeData() {
	switch string(png.chunkType[:]) {
	case "IHDR":
		fmt.Println("IHDR chunk")
	case "PLTE":
		fmt.Println("PLTE chunk")
	case "IDAT":
		fmt.Println("IDAT chunk")
	case "IEND":
		fmt.Println("IEND chunk")
	default:
		fmt.Println("Unknown chunk")
	}
}
