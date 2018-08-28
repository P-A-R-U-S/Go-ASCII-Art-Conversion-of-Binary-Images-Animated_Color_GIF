package main

import (
	_ "image"
	_ "image/draw"
	_ "image/gif"
	_ "image/png"

	"bufio"
	"bytes"
	"fmt"
	"github.com/urfave/cli"
	"image"
	"image/draw"
	"image/gif"
	"image/png"
	"io"
	"log"
	"os"
)

var (
	// VERSION indicates which version of the binary is running.
	VERSION = "1.0.0.1"

	// GITCOMMIT indicates which git hash the binary was built off of
	GITCOMMIT string

	file   *os.File
	frames []bytes.Buffer
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

	a := cli.NewApp()
	a.Name = "GIF-Image to ANSI"
	a.Usage = "Converting animated GIF-Image to binary ANSI animation."
	a.Author = "Valentyn Ponomarenko"
	a.Copyright = "Valentyn Ponomarenko"
	a.Version = VERSION
	a.Email = "ValentynPonomarenko@gmail.com"
	a.Description = "Command line utility done as experiment and can be free use or modified under MIT License."

	a.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "imgPath",
			Usage: "path to animated gif-image, e.g. --imgPath ./Desktop/animation.gif",
		},
	}

	a.Action = func(c *cli.Context) error {
		var err error

		if len(c.Args()) == 0 {
			cli.ShowAppHelp(c)
		}

		if c.IsSet("imgPath") {
			file, err = os.Open(c.String("imgPath"))

			if err != nil {
				log.Fatalf("not able to open gil-file, %s", err)
			}

			frames, err = getFrames(file)

			if err != nil {
				log.Fatalf("not able to parse gil-file, %s", err)
			}
		}

		return nil
	}

	err := a.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

}
