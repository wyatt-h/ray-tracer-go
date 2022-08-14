package main

import (
	"fmt"
	"os"
)

func main() {
	imageWidth := 256
	imageHeight := 256

	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	for j := imageHeight - 1; j >= 0; j-- {
		fmt.Fprintln(os.Stderr, "\rScanlines remaining: ", j)
		for i := 0; i < imageWidth; i++ {
			r := float64(i) / float64(imageWidth-1)
			g := float64(j) / float64(imageHeight-1)
			b := 0.25

			ir := int64(255.999 * r)
			ig := int64(255.999 * g)
			ib := int64(255.999 * b)

			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
	fmt.Fprintln(os.Stderr, "\nDone.")
}
