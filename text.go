package blocks

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

var (
	font, _ = truetype.Parse(goregular.TTF)
)

func (tbd *TwoBlocksDown) settingTopFont(outImage *gg.Context, text string, heightTextFrame int) (fontSize int, err error) {
	fontSize = 10
	for ;; {
		faceFace := truetype.NewFace(font, &truetype.Options{Size: float64(fontSize)})
		outImage.SetFontFace(faceFace)

		widthText, heightText := outImage.MeasureString(text)
		if int(heightText) > heightTextFrame { fontSize -= 1 }
		if int(heightText) < heightTextFrame { fontSize += 1 }
		if (heightTextFrame-int(heightText)) < 5 &&
			(heightTextFrame-int(heightText)) > -5 {
			for ; outImage.Width() < int(widthText); {
				fontSize -= 1
				faceFace = truetype.NewFace(font, &truetype.Options{Size: float64(fontSize)})
				outImage.SetFontFace(faceFace)
				widthText, heightText = outImage.MeasureString(text)
			}
			return
		}
	}
}

func (tbd *TwoBlocksDown) settingMiddleFont(outImage *gg.Context, text string, heightTextFrame int) (fontSize int, err error) {
	fontSize = 10
	for ;; {
		fontFace := truetype.NewFace(font, &truetype.Options{Size: float64(fontSize)})
		outImage.SetFontFace(fontFace)

		widthText, heightText := outImage.MeasureString(text)
		if int(heightText) > heightTextFrame { fontSize -= 1 }
		if int(heightText) < heightTextFrame { fontSize += 1 }
		if (heightTextFrame - int(heightText)) < 5 &&
			(heightTextFrame - int(heightText)) > -5 {
			for ; outImage.Width() / 2 < int(widthText); {
				fontSize -= 1
				fontFace = truetype.NewFace(font, &truetype.Options{Size: float64(fontSize)})
				outImage.SetFontFace(fontFace)
				widthText, heightText = outImage.MeasureString(text)
			}
			return
		}
	}
}

func (tbd *TwoBlocksDown) setTexts(outImage *gg.Context, texts []string) (*gg.Context, error) {
	outImage.SetHexColor("#000000")

	fontSizeTop, err := tbd.settingTopFont(outImage, texts[0], tbd.TemplateConfig.TopLineHeight)
	if err != nil {
		return outImage, err
	}
	if fontSizeTop < 10 { fontSizeTop = 0 }


	fontSizeLeft, err := tbd.settingMiddleFont(outImage, texts[1], tbd.TemplateConfig.MiddleLinesHeight)
	if err != nil { return outImage, err }
	fontSizeLeft -= 15
	if fontSizeLeft < 10 { fontSizeLeft = 0 }


	fontSizeRight, err := tbd.settingMiddleFont(outImage, texts[2], tbd.TemplateConfig.MiddleLinesHeight)
	if err != nil { return outImage, err }
	fontSizeRight -= 15
	if fontSizeRight < 10 { fontSizeRight = 0 }

	fontFaceTop := truetype.NewFace(font, &truetype.Options{Size: float64(fontSizeTop)})
	fontFaceLeft := truetype.NewFace(font, &truetype.Options{Size: float64(fontSizeLeft)})
	fontFaceRight := truetype.NewFace(font, &truetype.Options{Size: float64(fontSizeRight)})

	outImage.SetFontFace(fontFaceTop)
	widthTopText, heightTopText := outImage.MeasureString(texts[0])
	outImage.DrawString(
		texts[0],
		float64((outImage.Width() / 2)-int(widthTopText / 2)),
		float64(((tbd.TemplateConfig.TopLineHeight + tbd.TextConfig.MarginTop + tbd.TextConfig.MarginBottom) / 2) +
			int(heightTopText / 3)),
	)

	outImage.SetFontFace(fontFaceLeft)
	widthLeftText, _ := outImage.MeasureString(texts[1])
	outImage.DrawString(
		texts[1],
		float64((outImage.Width() / 4)-int(widthLeftText / 2)),
		float64((tbd.TemplateConfig.TopLineHeight + tbd.TextConfig.MarginTop + tbd.TextConfig.MarginBottom) +
			tbd.TemplateConfig.LineWidth +
			((tbd.TemplateConfig.MiddleLinesHeight + tbd.TextConfig.MarginTop) / 2)),
	)

	outImage.SetFontFace(fontFaceRight)
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
