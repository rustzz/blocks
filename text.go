package blocks

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

var (
	font, _ = truetype.Parse(goregular.TTF)
)

func (template *Template) configureTopFont(text string, heightTextFrame int) (err error) {
	template.TextConfig.FontConfig.FontSize = 10
	for ;; {
		faceFace := truetype.NewFace(font, &truetype.Options{
			Size: float64(template.TextConfig.FontConfig.FontSize),
		})
		template.Image.SetFontFace(faceFace)

		widthText, heightText := template.Image.MeasureString(text)
		if int(heightText) > heightTextFrame { template.TextConfig.FontConfig.FontSize -= 1 }
		if int(heightText) < heightTextFrame { template.TextConfig.FontConfig.FontSize += 1 }
		if (heightTextFrame-int(heightText)) < 5 &&
			(heightTextFrame-int(heightText)) > -5 {
			for ; template.Image.Width() < int(widthText); {
				template.TextConfig.FontConfig.FontSize -= 1
				faceFace = truetype.NewFace(font, &truetype.Options{
					Size: float64(template.TextConfig.FontConfig.FontSize),
				})
				template.Image.SetFontFace(faceFace)
				widthText, heightText = template.Image.MeasureString(text)
			}
			return
		}
	}
}

func (template *Template) configureMiddleFont(text string, heightTextFrame int) (err error) {
	template.TextConfig.FontConfig.FontSize = 10
	for ;; {
		fontFace := truetype.NewFace(font, &truetype.Options{
			Size: float64(template.TextConfig.FontConfig.FontSize),
		})
		template.Image.SetFontFace(fontFace)

		widthText, heightText := template.Image.MeasureString(text)
		if int(heightText) > heightTextFrame { template.TextConfig.FontConfig.FontSize -= 1 }
		if int(heightText) < heightTextFrame { template.TextConfig.FontConfig.FontSize += 1 }
		if (heightTextFrame - int(heightText)) < 5 &&
			(heightTextFrame - int(heightText)) > -5 {
			for ; template.Image.Width() / 2 < int(widthText); {
				template.TextConfig.FontConfig.FontSize -= 1
				fontFace = truetype.NewFace(font, &truetype.Options{
					Size: float64(template.TextConfig.FontConfig.FontSize),
				})
				template.Image.SetFontFace(fontFace)
				widthText, heightText = template.Image.MeasureString(text)
			}
			return
		}
	}
}

func (template *Template) RenderTexts() (err error) {
	template.Image.SetHexColor("#000000")

	if err = template.configureTopFont(
		template.TextConfig.Texts[0],
		template.FrameConfig.TopHeight,
	); err != nil { return }
	fontSizeTop := template.TextConfig.FontConfig.FontSize
	if fontSizeTop < 10 { fontSizeTop = 0 }

	if err = template.configureMiddleFont(
		template.TextConfig.Texts[1],
		template.FrameConfig.MiddleHeight,
	); err != nil { return }
	fontSizeLeft := template.TextConfig.FontConfig.FontSize
	if fontSizeLeft < 10 { fontSizeLeft = 0 }

	if err = template.configureMiddleFont(
		template.TextConfig.Texts[2],
		template.FrameConfig.MiddleHeight,
	); err != nil { return }
	fontSizeRight := template.TextConfig.FontConfig.FontSize
	if fontSizeRight < 10 { fontSizeRight = 0 }

	fontFaceTop := truetype.NewFace(font, &truetype.Options{Size: float64(fontSizeTop)})
	fontFaceLeft := truetype.NewFace(font, &truetype.Options{Size: float64(fontSizeLeft)})
	fontFaceRight := truetype.NewFace(font, &truetype.Options{Size: float64(fontSizeRight)})

	template.Image.SetFontFace(fontFaceTop)
	widthTopText, heightTopText := template.Image.MeasureString(template.TextConfig.Texts[0])
	template.Image.DrawString(
		template.TextConfig.Texts[0],
		float64((template.Image.Width() / 2)-int(widthTopText / 2)),
		float64(int(heightTopText)),
	)

	template.Image.SetFontFace(fontFaceLeft)
	widthLeftText, heightLeftText := template.Image.MeasureString(template.TextConfig.Texts[1])
	template.Image.DrawString(
		template.TextConfig.Texts[1],
		float64((template.Image.Width() / 4)-int(widthLeftText / 2)),
		float64(template.FrameConfig.TopHeight + template.FrameConfig.LineWidth + int(heightLeftText)),
	)

	template.Image.SetFontFace(fontFaceRight)
	widthRightText, heightRightText := template.Image.MeasureString(template.TextConfig.Texts[2])
	template.Image.DrawString(
		template.TextConfig.Texts[2],
		float64((template.Image.Width() / 2) + (template.Image.Width() / 4) - int(widthRightText / 2)),
		float64(template.FrameConfig.TopHeight + template.FrameConfig.LineWidth + int(heightRightText)),
	)
	return
}
