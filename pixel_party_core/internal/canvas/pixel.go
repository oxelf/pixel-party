package canvas

import "log"

type Pixel struct {
	R, G, B, A uint8
}

func DecodePixel(pixel uint16) Pixel {
	var rgba [4]int

	rgba[0] = int((pixel >> 12) & 0xF) // Red (bits 12-15)
	rgba[1] = int((pixel >> 8) & 0xF)  // Green (bits 8-11)
	rgba[2] = int((pixel >> 4) & 0xF)  // Blue (bits 4-7)
	rgba[3] = int(pixel & 0xF)         // Alpha (bits 0-3)

	return Pixel{
		R: uint8(rgba[0]),
		G: uint8(rgba[1]),
		B: uint8(rgba[2]),
		A: uint8(rgba[3]),
	}
}

func DecodeFromTwoBytePixels(byte1 uint8, byte2 uint8) Pixel {
	return Pixel{
		R: uint8((byte1 >> 4) & 0xF),
		G: uint8(byte1 & 0xF),
		B: uint8((byte2 >> 4) & 0xF),
		A: uint8(byte2 & 0xF),
	}
}

func (p *Pixel) EncodePixel() uint16 {
	if p.R < 0 || p.R > 15 || p.G < 0 || p.G > 15 || p.B < 0 || p.B > 15 || p.A < 0 || p.A > 15 {
		log.Fatalf("Invalid RGBA values: R=%d, G=%d, B=%d, A=%d", p.R, p.G, p.B, p.A)
	}
	var result uint16
	result = result | uint16(p.R)<<12
	result = result | uint16(p.G)<<8
	result = result | uint16(p.B)<<4
	result = result | uint16(p.A)
	return result
}
