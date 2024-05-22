package main

import (
	"image/color"
	"image/jpeg"
	"log"
	"os"

	gim "github.com/766b/go-image-merge"
	"github.com/766b/go-image-merge/imagefilter"
)

func main() {
	grids := []*gim.Grid{
		{
			ImageFilePath:   "./cmd/gim/input/ginger.png",
			BackgroundColor: color.White,
			Filters: []imagefilter.FilterType{
				imagefilter.Vignette,
			},
		},
		{
			ImageFilePath:   "./cmd/gim/input/ginger.png",
			BackgroundColor: color.RGBA{R: 0x8b, G: 0xd0, B: 0xc6},
		},
	}
	rgba, err := gim.New(grids, 2, 1).Merge()
	if err != nil {
		log.Panicf(err.Error())
	}

	file, err := os.Create("cmd/gim/output/merged.jpg")
	if err != nil {
		log.Panicf(err.Error())
	}

	err = jpeg.Encode(file, rgba, &jpeg.Options{Quality: 80})
	if err != nil {
		log.Panicf(err.Error())
	}
}
