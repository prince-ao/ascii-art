package main

import (
	"fmt"
	"github.com/alecthomas/kong"
	"image"
	// "image/color"
	_ "image/jpeg"
	"log"
	"math"
	"os"
	// "strconv"
	// "time"
)

var CLI struct {
	Info struct {
		Path string `arg:"" name:"path" help:"Path to image."`
	} `cmd:"" help:"Get info on file."`

	Convert struct {
		Width int `optional:"" help:"Desired width of image."`

		Path string `arg:"" name:"path" help:"Path to image."`
	} `cmd:"" help:"Convert image to ascii."`
}

func w(t float32) float64 {
	absT := math.Abs(float64(t))
	a := -0.5

	if absT <= 1 {
		return (a+2)*math.Pow(absT, 3) - (a+3)*math.Pow(absT, 2) + 1
	} else if absT < 2 {
		return a*math.Pow(absT, 3) - 5*a*math.Pow(absT, 2) + 8*a*absT - 4*a
	} else {
		return 0
	}
}

func handleInfo(path string) {
	reader, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	bounds := m.Bounds()

	fmt.Printf("Image Size: (%d x %d)\n", bounds.Max.X, bounds.Max.Y)
}

func handleConvert(path string, width int) {

}

func main() {
	ctx := kong.Parse(&CLI)

	switch ctx.Command() {
	case "info <path>":
		handleInfo(CLI.Info.Path)
	case "convert":
		// handleConvert(CLI.Convert.Path, CLI.Convert.Width)
	default:
		fmt.Printf("%v", ctx.Command())
	}
	/* args := os.Args

	if len(args[1:]) < 1 {
		log.Fatal("expected 2 command line arguments.")
	}

	file := args[1]

	W, err := strconv.Atoi(args[2])

	if err != nil {
		log.Fatal(err)
	}

	reader, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer reader.Close()

	m, _, err := image.Decode(reader)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully loaded image!")

	bounds := m.Bounds()

	fmt.Printf("Image size: %d x %d\n", bounds.Max.X, bounds.Max.Y)

	scale := float64(W) / float64(bounds.Max.X)

	if scale > 1 {
		log.Fatal("only down-sampling allowed")
	}

	H := float64(bounds.Max.Y) * scale

	newImage := image.NewRGBA(image.Rect(0, 0, W, int(H)))

	newBounds := newImage.Bounds()

	for y1 := newBounds.Min.Y; y1 < newBounds.Max.Y; y1++ {
		for x1 := newBounds.Min.X; x1 < newBounds.Max.X; x1++ {
			x0, y0 := (float32(x1)+0.5)/float32(scale)-0.5, (float32(y1)+0.5)/float32(scale)-0.5

			i, j := int(x0), int(y0)

			sigmaX := x0 - float32(i)
			sigmaY := y0 - float32(j)

			window := [4]int{-1, 0, 1, 2}

			Wx := [4]float64{w(-1 - sigmaX), w(0 - sigmaX), w(1 - sigmaX), w(2 - sigmaX)}
			Wy := [4]float64{w(-1 - sigmaY), w(0 - sigmaY), w(1 - sigmaY), w(2 - sigmaY)}

			W := make([][]float64, 4)

			for i := 0; i < 4; i++ {
				W[i] = make([]float64, 4)
			}

			for _, v1 := range window {
				for _, v2 := range window {
					W[v1+1][v2+1] = Wx[v1+1] * Wy[v2+1]
				}
			}

			var accR float64 = 0
			var accG float64 = 0
			var accB float64 = 0

			for _, v1 := range window {
				for _, v2 := range window {
					real_i := i + v1
					real_j := j + v2
					P := [3]uint32{}
					if real_i < bounds.Min.X || real_i >= bounds.Max.X || real_j < bounds.Min.Y || real_j >= bounds.Max.Y {
						P[0] = 0
						P[1] = 0
						P[2] = 0
					} else {
						r, g, b, _ := m.At(real_i, real_j).RGBA()

						P[0] = r >> 8
						P[1] = g >> 8
						P[2] = b >> 8
					}

					w := W[v1+1][v2+1]

					accR += float64(P[0]) * w
					accG += float64(P[1]) * w
					accB += float64(P[2]) * w
				}
			}

			outR := uint8(math.Min(math.Max(accR, 0), 255))
			outG := uint8(math.Min(math.Max(accG, 0), 255))
			outB := uint8(math.Min(math.Max(accB, 0), 255))

			newImage.Set(x1, y1, color.RGBA{outR, outG, outB, 255})
		}
	}

	time.Sleep(5 * 1e9)

	brightness_symbols := "`^\",:;Il!i~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"

	for y := newBounds.Min.Y; y < newBounds.Max.Y; y++ {
		for x := newBounds.Min.X; x < newBounds.Max.X; x++ {
			r16, g16, b16, _ := newImage.At(x, y).RGBA()

			r8, g8, b8 := uint(r16>>8), uint(g16>>8), uint(b16>>8)

			luminosity := uint8(0.21*float32(r8) + 0.72*float32(g8) + 0.07*float32(b8))

			symbol_index := int8(math.Round(float64(luminosity) * float64(len(brightness_symbols)-1) / 255))

			symbol := brightness_symbols[symbol_index]
			fmt.Printf("%c%c%c", symbol, symbol, symbol)
		}
		fmt.Println()
	} */

}
