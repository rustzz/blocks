package blocks

import (
	"bytes"
	"github.com/fogleman/gg"
	"image"
)

type TextConfig struct {
	MarginTop			int
	MarginBottom		int
}

type TemplateConfig struct {
	LineWidth			int
	TopLineHeight		int
	MiddleLinesHeight	int
	ImageHeight			int
	ImageWidth			int
}

type TwoBlocksDown struct {
	OutImage			*gg.Context
	TemplateConfig		TemplateConfig
	TextConfig			TextConfig
}

func New() *TwoBlocksDown {
	MarginTop := 10
	MarginBottom := 10
	return &TwoBlocksDown{
		TemplateConfig: TemplateConfig{
			LineWidth:         2,
			TopLineHeight:     30,
			MiddleLinesHeight: 30,
			ImageHeight:       250,
			ImageWidth:        250,
		},
		TextConfig: TextConfig{
			MarginTop: MarginTop,
			MarginBottom: MarginBottom,
		},
	}
}

func saveImage(outImage *gg.Context, path string) (imageReader *bytes.Reader, err error) {
	if len(path) != 0 {
		err = outImage.SavePNG(path)
		if err != nil {
			return
		}
	} else {
		imageBuffer := new(bytes.Buffer)
		err = outImage.EncodePNG(imageBuffer)
		if err != nil {
			return
		}
		imageReader = bytes.NewReader(imageBuffer.Bytes())
		return
	}
	return
}

func (tbd *TwoBlocksDown) Make(srcImages []image.Image, texts []string, outPath string) (imageReader *bytes.Reader, err error) {
	outImage := tbd.createTemplate()
	outImage = tbd.placeSrcImages(outImage, srcImages)
	outImage, err = tbd.setTexts(outImage, texts)
	if err != nil {
		return
	}

	tbd.OutImage = outImage
	imageReader, err = saveImage(outImage, outPath)
	if err != nil {
		return
	}
	return
}
