package storage

import (
	"context"
	"fmt"

	"github.com/oxelf/pixel-party/internal/canvas"
)

func (r *RedisConnection) SetCanvas(canvasId string, pixel []canvas.Pixel) error {
	ctx := context.Background()

	var cmds []string
	for i := 0; i < len(pixel); i++ {
		cmds = append(cmds, "set")
		cmds = append(cmds, "u16")
		cmds = append(cmds, fmt.Sprintf("#%d", i))
		cmds = append(cmds, fmt.Sprintf("%d", pixel[i].EncodePixel()))
	}

	return r.BitField(ctx, "canvas:"+canvasId, cmds).Err()
}

func (r *RedisConnection) SetPixel(canvasId string, pixel canvas.Pixel, offset int) error {
	ctx := context.Background()

	return r.BitField(ctx, "canvas:"+canvasId, "set", "u16", fmt.Sprintf("#%d", offset), fmt.Sprintf("%d", pixel.EncodePixel())).Err()
}

func (r *RedisConnection) GetCanvasAsPixel(canvasId string) ([]canvas.Pixel, error) {
	ctx := context.Background()
	val, err := r.Get(ctx, "canvas:"+canvasId).Bytes()
	if err != nil {
		return nil, err
	}

	pixel := []canvas.Pixel{}

	for i := 0; i < len(val); i += 2 {
		pixel = append(pixel, canvas.DecodeFromTwoBytePixels(val[i], val[i+1]))
	}

	return pixel, nil
}

func (r *RedisConnection) GetCanvas(canvasId string) ([]byte, error) {
	ctx := context.Background()

	val, err := r.Get(ctx, "canvas:"+canvasId).Bytes()
	if err != nil {
		return nil, err
	}

	return val, nil
}
