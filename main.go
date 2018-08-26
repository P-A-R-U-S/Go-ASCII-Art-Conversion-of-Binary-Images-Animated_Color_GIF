package main

import (
	"fmt"
	_ "image"
	_ "image/draw"
	_ "image/gif"
	_ "image/png"
	"io"

	"bufio"
	"bytes"
	"image"
	"image/draw"
	"image/gif"
	"image/png"
	"log"
)

func getFrames(reader io.Reader) (pngFrames []bytes.Buffer, err error) {

	// decoding image can panic pretty frequently if image are broken or contains missed pixels
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("image decoding error: %s", r)
		}
	}()

	// get frames from existing gif-image
	gi, err := gif.DecodeAll(reader)

	if err != nil {
		return nil, err
	}

	// calculate image dimension
	var lX, lY, hX, hY, iHeight, iWidth int
	for _, ip := range gi.Image {
		if ip.Rect.Min.X < lX {
			lX = ip.Rect.Min.X
		}
		if ip.Rect.Min.Y < lY {
			lY = ip.Rect.Min.Y
		}
		if ip.Rect.Max.X > hX {
			hX = ip.Rect.Max.X
		}
		if ip.Rect.Max.Y > hY {
			hY = ip.Rect.Max.Y
		}
	}

	iWidth = hX - lX
	iHeight = hY - lY

	overpaintImage := image.NewRGBA(image.Rect(0, 0, iWidth, iHeight))
	draw.Draw(overpaintImage, overpaintImage.Bounds(), gi.Image[0], image.ZP, draw.Src)

	result := make([]bytes.Buffer, len(gi.Image))

	for i, srcImg := range gi.Image {
		draw.Draw(overpaintImage, overpaintImage.Bounds(), srcImg, image.ZP, draw.Over)

		w := bufio.NewWriter(&result[i])

		if err := png.Encode(w, overpaintImage); err != nil {
			log.Fatal(err)
		}

	}

	return result[:], nil

}

func main() {

}
