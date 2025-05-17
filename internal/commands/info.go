package commands

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"os"
)

func HandleInfo(path string) {
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
