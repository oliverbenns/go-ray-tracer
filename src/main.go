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

func calculateColor(r *Ray) Color {
	unitDirection := r.direction.UnitVec()
	t := 0.5 * (unitDirection.y + 1.0)
	v := Vec3{1.0, 1.0, 1.0}.MultF(1.0 - t).Add(Vec3{0.5, 0.7, 1.0}.MultF(t))

	return Color{v.x, v.y, v.z}
}

func main() {
	nx := 200
	ny := 100
	img := image.NewRGBA(image.Rect(0, 0, nx, ny))

	lowerLeftCorner := Vec3{-2.0, -1.0, -1.0}
	horizontal := Vec3{4.0, 0.0, 0.0}
	vertical := Vec3{0.0, 2.0, 0.0}
	origin := Vec3{0.0, 0.0, 0.0}

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u := float64(i) / float64(nx)
			v := float64(ny-j-1) / float64(ny)
			r := Ray{origin, lowerLeftCorner.Add(horizontal.MultF(u)).Add(vertical.MultF(v))}
			col := calculateColor(&r).MultF(0xffff)

			imageColor := color.RGBA64{
				uint16(col.r),
				uint16(col.g),
				uint16(col.b),
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
