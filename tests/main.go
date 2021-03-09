package main

import (
	"bytes"
	"fmt"
	"github.com/rustzz/blocks"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	urls = [2]string{
		"https://2ch.hk/makaba/templates/img/anon.jpg",
		"https://2ch.hk/makaba/templates/img/anon.jpg",
	}
	texts = [3]string{"cum CUM", "cum CUM", "cum CUM"}
)

func GetImage(url string) (outImage image.Image, err error) {
	resp, err := http.Get(url)
	if err != nil { return }
	defer resp.Body.Close()

	imageBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil { return }

	imageBuffer := bytes.NewBuffer(imageBytes)
	outImage, _, err = image.Decode(imageBuffer)
	if err != nil { return }
	return
}

func FromConstructor() (imgBytes []byte, err error) {
	srcImages, err := func () (out [2]image.Image, err error) {
		for index, url := range urls {
			out[index], err = GetImage(url)
			if err != nil { return }
		}
		return
	}()
	if err != nil { return }
	dem := blocks.New(srcImages, texts)
	imgBytes, err = dem.Make()
	if err != nil { return }
	return
}

func main() {
	imgBytes, err := FromConstructor()
	if err != nil { log.Fatal(err) }

	homeDir, err := os.UserHomeDir()
	file, err := os.Create(fmt.Sprintf("%s/out.png", homeDir))
	if err != nil { log.Fatal(err) }
	defer file.Close()

	imgBuffer := bytes.NewBuffer(imgBytes)
	im, _, err := image.Decode(imgBuffer)
	if err != nil { log.Fatal(err) }
	if err = png.Encode(file, im); err != nil { log.Fatal(err) }
}
