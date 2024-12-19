package canvas

import (
	"github.com/onsi/gomega"
	"testing"
)

func TestEncodePixel(t *testing.T) {
	g := gomega.NewWithT(t)
	pixel := Pixel{12, 4, 8, 15}
	result := pixel.EncodePixel()
	g.Expect(result).To(gomega.Equal(uint16(0xC48F)))
}

func TestDecodePixel(t *testing.T) {
	g := gomega.NewWithT(t)

	pixel := DecodePixel(50319)
	g.Expect(pixel).To(gomega.Equal(Pixel{12, 4, 8, 15}))
}

func TestDecodeTwoBytePixel(t *testing.T) {
	g := gomega.NewWithT(t)

	pixel := DecodeFromTwoBytePixels(uint8(0xC4), uint8(0x8F))
	g.Expect(pixel).To(gomega.Equal(Pixel{12, 4, 8, 15}))
}

func TestMain(m *testing.M) {
	m.Run()
}
