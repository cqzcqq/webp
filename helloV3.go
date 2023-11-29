//go:build ignore
// +build ignore

package main

import (
	"image/jpeg"
	"log"
	"os"

	"github.com/cqzcqq/webp"
)

func main() {
	file, err := os.Open("testdata/m4_q75.webp")
	if err != nil {
		log.Fatalln(err)
	}

	output, err := os.Create("testdata/m4_q75.jpg")
	if err != nil {
		log.Fatal(err)
	}
	//noinspection GoUnhandledErrorResult
	defer output.Close()

	img, err := webp.Decode(file)
	if err != nil {
		log.Println(err)
	}

	if err = jpeg.Encode(output, img, &jpeg.Options{Quality: 75}); err != nil {
		log.Fatalln(err)
	}
}
