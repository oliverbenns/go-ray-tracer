package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"time"
)

func saveImage(img *image.RGBA, filename string) error {
	err := os.MkdirAll("./img", os.ModePerm)

	if err != nil {
		return err
	}

	file, err := os.Create("./img/" + filename)

	if err != nil {
		return err
	}

	defer file.Close()

	return png.Encode(file, img)
}

func main() {
	nx := 200
	ny := 100
	img := image.NewRGBA(image.Rect(0, 0, nx, ny))

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			col := Vec3{
				float64(i) / float64(nx),
				float64(ny-j-1) / float64(ny),
				0.2,
			}

			imageColor := color.RGBA64{
				uint16(0xffff * col.R()),
				uint16(0xffff * col.G()),
				uint16(0xffff * col.B()),
				0xffff,
			}
			img.Set(i, j, imageColor)
		}
	}

	filename := time.Now().Format(time.RFC3339) + ".png"
	err := saveImage(img, filename)

	if err != nil {
		panic(err)
	}
}
