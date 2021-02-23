package blocks

import (
	"github.com/fogleman/gg"
	"image"
	"github.com/nfnt/resize"
)

func (tbd *TwoBlocksDown) makeTemplateImage() (templateImage *gg.Context) {
	templateWidth := (tbd.TemplateConfig.ImageWidth * 2) + tbd.TemplateConfig.LineWidth
	templateHeight := (tbd.TemplateConfig.LineWidth * 2) + (tbd.TextConfig.MarginTop + tbd.TextConfig.MarginBottom) +
		tbd.TemplateConfig.MiddleLinesHeight + tbd.TemplateConfig.TopLineHeight + tbd.TemplateConfig.ImageHeight
	templateImage = gg.NewContext(templateWidth, templateHeight)
	return
}

func (tbd *TwoBlocksDown) fillBackground(template *gg.Context) *gg.Context {
	template.SetHexColor("#ffffff")
	template.DrawRectangle(
		0, 0,
		float64(template.Width()), float64(template.Height()),
	)
	template.Fill()
	return template
}

func (tbd *TwoBlocksDown) drawFrame(template *gg.Context) *gg.Context {
	template.SetHexColor("#000000")
	template.DrawRectangle(
		float64((template.Width() / 2) - (tbd.TemplateConfig.LineWidth / 2)),
		float64(tbd.TemplateConfig.TopLineHeight + tbd.TextConfig.MarginTop + tbd.TextConfig.MarginBottom),
		float64(tbd.TemplateConfig.LineWidth), float64(template.Height()),
	)
	template.Fill()
	template.DrawRectangle(
		0, float64(tbd.TemplateConfig.TopLineHeight + tbd.TextConfig.MarginTop + tbd.TextConfig.MarginBottom),
		float64(template.Width()), float64(tbd.TemplateConfig.LineWidth),
	)
	template.Fill()
	template.DrawRectangle(
		0, float64(tbd.TemplateConfig.TopLineHeight + tbd.TextConfig.MarginTop + tbd.TextConfig.MarginBottom +
			tbd.TemplateConfig.LineWidth + tbd.TemplateConfig.MiddleLinesHeight),
		float64(template.Width()), float64(tbd.TemplateConfig.LineWidth),
	)
	template.Fill()
	return template
}

func (tbd *TwoBlocksDown) resizeSrcImage(srcImage image.Image) image.Image {
	return resize.Resize(
		uint(tbd.TemplateConfig.ImageWidth), uint(tbd.TemplateConfig.ImageHeight),
		srcImage, resize.Lanczos3,
	)
}

func (tbd *TwoBlocksDown) placeSrcImages(outImage *gg.Context, srcImages []*image.Image) *gg.Context {
	outImage.DrawImage(
		tbd.resizeSrcImage(*srcImages[0]), 0,
		tbd.TemplateConfig.TopLineHeight + tbd.TextConfig.MarginTop + tbd.TextConfig.MarginBottom +
			tbd.TemplateConfig.MiddleLinesHeight + (tbd.TemplateConfig.LineWidth * 2),
	)
	outImage.DrawImage(
		tbd.resizeSrcImage(*srcImages[1]),
		tbd.TemplateConfig.ImageWidth + tbd.TemplateConfig.LineWidth,
		tbd.TemplateConfig.TopLineHeight+tbd.TextConfig.MarginTop + tbd.TextConfig.MarginBottom +
			tbd.TemplateConfig.MiddleLinesHeight + (tbd.TemplateConfig.LineWidth * 2),
	)
	return outImage
}

func (tbd *TwoBlocksDown) createTemplate() (template *gg.Context) {
	template = tbd.makeTemplateImage()
	template = tbd.fillBackground(template)
	template = tbd.drawFrame(template)
	return
}