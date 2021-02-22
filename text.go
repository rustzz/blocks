package blocks

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

func (tbd *TwoBlocksDown) settingTopFont(outImage *gg.Context, text string, heightTextFrame int) (fontSize int, err error) {
	fontSize = 10
	for ;; {
		font, _ := truetype.Parse(goregular.TTF)
		face := truetype.NewFace(font, &truetype.Options{Size: float64(fontSize)})
		outImage.SetFontFace(face)

		widthText, heightText := outImage.MeasureString(text)
		if int(heightText) > heightTextFrame {
			fontSize -= 1
		}
		if int(heightText) < heightTextFrame {
			fontSize += 1
		}
		if (heightTextFrame-int(heightText)) < 5 &&
			(heightTextFrame-int(heightText)) > -5 {
			for ; outImage.Width() < int(widthText); {
				fontSize -= 1
				face = truetype.NewFace(font, &truetype.Options{Size: float64(fontSize)})
				outImage.SetFontFace(face)
				widthText, heightText = outImage.MeasureString(text)
			}
			return
		}
	}
}

func (tbd *TwoBlocksDown) settingMiddleFont(outImage *gg.Context, text string, heightTextFrame int) (fontSize int, err error) {
	fontSize = 10
	for ;; {
		font, _ := truetype.Parse(goregular.TTF)
		face := truetype.NewFace(font, &truetype.Options{Size: float64(fontSize)})
		outImage.SetFontFace(face)

		widthText, heightText := outImage.MeasureString(text)
		if int(heightText) > heightTextFrame {
			fontSize -= 1
		}
		if int(heightText) < heightTextFrame {
			fontSize += 1
		}
		if (heightTextFrame - int(heightText)) < 5 &&
			(heightTextFrame - int(heightText)) > -5 {
			for ; outImage.Width() / 2 < int(widthText); {
				fontSize -= 1
				face = truetype.NewFace(font, &truetype.Options{Size: float64(fontSize)})
				outImage.SetFontFace(face)
				widthText, heightText = outImage.MeasureString(text)
			}
			return
		}
	}
}

func (tbd *TwoBlocksDown) setTexts(outImage *gg.Context, texts []string) (*gg.Context, error) {
	font, _ := truetype.Parse(goregular.TTF)
	outImage.SetHexColor("#000000")

	fontSize, err := tbd.settingTopFont(outImage, texts[0], tbd.TemplateConfig.TopLineHeight)
	if err != nil {
		return outImage, err
	}
	if fontSize < 10 {
		fontSize = 0
	}

	face := truetype.NewFace(font, &truetype.Options{Size: float64(fontSize)})
	outImage.SetFontFace(face)

	widthTopText, heightTopText := outImage.MeasureString(texts[0])
	outImage.DrawString(
		texts[0],
		float64((outImage.Width() / 2)-int(widthTopText / 2)),
		float64(((tbd.TemplateConfig.TopLineHeight + tbd.TextConfig.MarginTop + tbd.TextConfig.MarginBottom) / 2) +
			int(heightTopText / 3)),
	)

	fontSize, err = tbd.settingMiddleFont(outImage, texts[1], tbd.TemplateConfig.MiddleLinesHeight)
	if err != nil {
		return outImage, err
	}
	fontSize -= 15
	if fontSize < 10 {
		fontSize = 0
	}

	face = truetype.NewFace(font, &truetype.Options{Size: float64(fontSize)})
	outImage.SetFontFace(face)

	widthLeftText, _ := outImage.MeasureString(texts[1])
	outImage.DrawString(
		texts[1],
		float64((outImage.Width() / 4)-int(widthLeftText / 2)),
		float64((tbd.TemplateConfig.TopLineHeight + tbd.TextConfig.MarginTop + tbd.TextConfig.MarginBottom) +
			tbd.TemplateConfig.LineWidth +
			((tbd.TemplateConfig.MiddleLinesHeight + tbd.TextConfig.MarginTop) / 2)),
	)

	fontSize, err = tbd.settingMiddleFont(outImage, texts[2], tbd.TemplateConfig.MiddleLinesHeight)
	if err != nil {
		return outImage, err
	}
	fontSize -= 15
	if fontSize < 10 {
		fontSize = 0
	}

	face = truetype.NewFace(font, &truetype.Options{Size: float64(fontSize)})
	outImage.SetFontFace(face)

	widthRightText, _ := outImage.MeasureString(texts[2])
	outImage.DrawString(
		texts[2],
		float64((outImage.Width() / 2) + (outImage.Width() / 4) - int(widthRightText / 2)),
		float64((tbd.TemplateConfig.TopLineHeight + tbd.TextConfig.MarginTop + tbd.TextConfig.MarginBottom) +
			tbd.TemplateConfig.LineWidth +
			((tbd.TemplateConfig.MiddleLinesHeight + tbd.TextConfig.MarginTop) / 2)),
	)
	return outImage, nil
}
