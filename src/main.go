package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func saveImage(data []byte, filename string) error {
	err := os.MkdirAll("./img", os.ModePerm)

	if err != nil {
		return err
	}

	return ioutil.WriteFile("./img/"+filename, []byte(data), os.ModePerm)
}

func main() {
	nx := 200
	ny := 100
	data := fmt.Sprint("P3\n", nx, " ", ny, "\n255\n")

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			r := float64(i) / float64(nx)
			g := float64(j) / float64(ny)
			b := 0.2
			ir := int64(255.99 * r)
			ig := int64(255.99 * g)
			ib := int64(255.99 * b)
			data = fmt.Sprint(data, ir, " ", ig, " ", ib, "\n")
		}
	}

	err := saveImage([]byte(data), "test.ppm")

	if err != nil {
		panic(err)
	}
}
