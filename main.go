package main

import (
	"fmt"
	"github.com/urfave/cli"
	"image/color"
	"image/gif"
	"log"
	"math/rand"
	"os"
	"strconv"
)

const (
	PARAM_NAME_IMAGE_PATH  string = "path"
	PARAM_NAME_IMAGE_WIDTH string = "width"

	ANSI_BASIC_BASE    int     = 16
	ANSI_COLOR_SPACE   uint32  = 6
	ANSI_FOREGROUND    string  = "38"
	ANSI_RESET         string  = "\x1b[0m"
	DEFAULT_CHARACTERS string  = "01"
	DEFAULT_WIDTH      int     = 100
	DEFAULT_HEIGTH     int     = 100
	PROPORTION         float32 = 0.46
	RGBA_COLOR_SPACE   uint32  = 1 << 16

	// VERSION indicates which version of the binary is running.
	VERSION = "1.0.0.1"
)

var (
	fileName   string
	width      = DEFAULT_WIDTH
	heigth     = DEFAULT_HEIGTH
	characters = DEFAULT_CHARACTERS
)

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
			Name:  "path, p",
			Usage: "path to animated gif-image, e.g. --path ./Desktop/animation.gif or --p ./Desktop/animation.gif ",
		},
		//cli.IntFlag{
		//	Name:  "width, w",
		//	Value: DEFAULT_WIDTH,
		//	Usage: "image width, e.g. -- width 200 or --w 200",
		//},
		cli.StringFlag{
			Name:  "characters, c",
			Value: DEFAULT_CHARACTERS,
			Usage: "characters set, e.g. -- characters 200 or --c 200",
		},
	}

	a.Action = func(c *cli.Context) error {

		if len(c.Args()) == 0 {
			cli.ShowAppHelp(c)
		}

		if c.IsSet(PARAM_NAME_IMAGE_PATH) {
			fileName = c.String(PARAM_NAME_IMAGE_PATH)

			if _, err := os.Stat(fileName); err != nil {
				if os.IsNotExist(err) {
					log.Fatalf("not able to open file, %s", err)
				}
			}

		}

		//if c.IsSet(PARAM_NAME_IMAGE_WIDTH) {
		//	width = c.Int(PARAM_NAME_IMAGE_WIDTH)
		//}

		Process()

		return nil
	}

	err := a.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

}

func Process() error {

	// decoding image can panic pretty frequently if image are broken or contains missed pixels
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(fmt.Errorf("image decoding error: %s", r))
		}
	}()

	r, err := os.Open(fileName)
	if err != nil {
		return err
	}

	// get frames from existing gif-image
	gifImage, err := gif.DecodeAll(r)
	if err != nil {
		return err
	}

	// calculate image dimension
	//var lX, lY, hX, hY, iHeight, iWidth int
	//for _, ip := range gifImage.Image {
	//	if ip.Rect.Min.X < lX {
	//		lX = ip.Rect.Min.X
	//	}
	//	if ip.Rect.Min.Y < lY {
	//		lY = ip.Rect.Min.Y
	//	}
	//	if ip.Rect.Max.X > hX {
	//		hX = ip.Rect.Max.X
	//	}
	//	if ip.Rect.Max.Y > hY {
	//		hY = ip.Rect.Max.Y
	//	}
	//}
	//
	//iWidth = hX - lX
	//iHeight = hY - lY

	//opImage := image.NewRGBA(image.Rect(0, 0, iWidth, iHeight))
	//draw.Draw(opImage, opImage.Bounds(), gifImage.Image[0], image.ZP, draw.Src)

	charactersLength := len(characters)
	for _, img := range gifImage.Image {

		width, heigth = Size()

		if err != nil {
			fmt.Println(fmt.Errorf("image decoding error: %v", r))
		}
		m := Resize(uint(width), uint(float32(heigth)*PROPORTION), img, LANCZOS_3)
		var current, previous string
		bounds := m.Bounds()
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				current = AnsiToCode(m.At(x, y))
				if current != previous {
					fmt.Print(current)
				}
				if ANSI_RESET != current {
					char := string(characters[rand.Int()%charactersLength])
					fmt.Print(char)
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Print("\n")
		}
		fmt.Print(ANSI_RESET)
	}

	return err
}

func AnsiToCode(c color.Color) string {
	r, g, b, _ := c.RGBA()
	code := int(ANSI_BASIC_BASE + toAnsiSpace(r)*36 + toAnsiSpace(g)*6 + toAnsiSpace(b))
	if code == ANSI_BASIC_BASE {
		return ANSI_RESET
	}
	return "\033[" + ANSI_FOREGROUND + ";5;" + strconv.Itoa(code) + "m"
}

func toAnsiSpace(val uint32) int {
	return int(float32(ANSI_COLOR_SPACE) * (float32(val) / float32(RGBA_COLOR_SPACE)))
}
