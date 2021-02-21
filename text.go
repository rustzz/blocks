package blocks

import (
	"fmt"
	"github.com/fogleman/gg"
	"path/filepath"
	"runtime"
)

// get root path of project
var (
	_, b, _, _	= runtime.Caller(0)
	basePath	= filepath.Dir(b)
	fontName	= "times.ttf"
)
// ========================

func (tbd *TwoBlocksDown) settingTopFont(outImage *gg.Context, text string, heightTextFrame int) (fontSize int, err error) {
	fontSize = 10
	for ;; {
		if err = outImage.LoadFontFace(fmt.Sprintf("%s/fonts/%s", basePath, fontName), float64(fontSize));
			err != nil {
			return
		}
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
				if err = outImage.LoadFontFace(fmt.Sprintf("%s/fonts/%s", basePath, fontName), float64(fontSize));
					err != nil {
					return
				}
				widthText, heightText = outImage.MeasureString(text)
			}
			return
		}
	}
}

func (tbd *TwoBlocksDown) settingMiddleFont(outImage *gg.Context, text string, heightTextFrame int) (fontSize int, err error) {
	fontSize = 10
	for ;; {
		if err = outImage.LoadFontFace(fmt.Sprintf("%s/fonts/%s", basePath, fontName), float64(fontSize));
			err != nil {
			return
		}
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
				if err = outImage.LoadFontFace(fmt.Sprintf("%s/fonts/%s", basePath, fontName), float64(fontSize));
					err != nil {
					return
				}
				widthText, heightText = outImage.MeasureString(text)
			}
			return
		}
	}
}

func (tbd *TwoBlocksDown) setTexts(outImage *gg.Context, texts []string) (*gg.Context, error) {
	outImage.SetHexColor("#000000")
	fontSize, err := tbd.settingTopFont(outImage, texts[0], tbd.TemplateConfig.TopLineHeight)
	if err != nil {
		return outImage, err
	}
	if fontSize < 10 {
		fontSize = 0
	}

	if err = outImage.LoadFontFace(fmt.Sprintf("%s/fonts/%s", basePath, fontName), float64(fontSize));
		err != nil {
		return outImage, err
	}

	widthTopText, heightTopText := outImage.MeasureString(texts[0])
	outImage.DrawString(
		texts[0],
		float64((outImage.Width() / 2)-int(widthTopText / 2)),
		float64(((tbd.TemplateConfig.TopLineHeight + tbd.TextConfig.MarginTop + tbd.TextConfig.MarginBottom) / 2) +
			int(heightTopText / 2)),
	)

	fontSize, err = tbd.settingMiddleFont(outImage, texts[1], tbd.TemplateConfig.MiddleLinesHeight)
	if err != nil {
		return outImage, err
	}
	fontSize -= 15
	if fontSize < 10 {
		fontSize = 0
	}

	if err = outImage.LoadFontFace(fmt.Sprintf("%s/fonts/%s", basePath, fontName), float64(fontSize)); err != nil {
		return outImage, err
	}

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

	if err = outImage.LoadFontFace(fmt.Sprintf("%s/fonts/%s", basePath, fontName), float64(fontSize)); err != nil {
		return outImage, err
	}

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
