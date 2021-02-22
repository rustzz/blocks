package main

import (
	"bytes"
	"fmt"
	"github.com/rustzz/blocks"
	"image"
	"log"
	"os"
)


func main() {
	urls := []string{
		"https://upload.wikimedia.org/wikipedia/commons/thumb/a/ad/Placeholder_no_text.svg/768px-Placeholder_no_text.svg.png",
		"https://upload.wikimedia.org/wikipedia/commons/thumb/a/ad/Placeholder_no_text.svg/768px-Placeholder_no_text.svg.png",
	}
	imageReaders := func () (out []*bytes.Reader) {
		for _, url := range urls {
			imageReader, err := blocks.LoadSrcImageFromURL(url)
			if err != nil {
				log.Fatal(err)
				return
			}
			out = append(out, imageReader)
		}
		return
	}()
	images := func () (out []image.Image) {
		for _, imageReader := range imageReaders {
			im, _, err := image.Decode(imageReader)
			if err != nil {
				log.Fatal(err)
				return
			}
			out = append(out, im)
		}
		return
	}()

	homeDir, _ := os.UserHomeDir()
	tbd := blocks.New()
	if _, err := tbd.Make(images, []string{
		"PLACE HOLDER | place holder", "PLACE HOLDER | place holder", "PLACE HOLDER | place holder",
	}, fmt.Sprintf("%s/out.png", homeDir)); err != nil {
		log.Fatal(err)
	}
}
