package blocks

import (
	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
)

func (template *Template) makeTemplateImage() {
	templateWidth := (template.FrameConfig.ImageWidth * 2) +
		template.FrameConfig.LineWidth
	templateHeight := (template.FrameConfig.LineWidth * 2) +
		(template.TextConfig.MarginTop + template.TextConfig.MarginBottom) +
		template.FrameConfig.MiddleHeight + template.FrameConfig.TopHeight +
		template.FrameConfig.ImageHeight
	template.Image = gg.NewContext(templateWidth, templateHeight)
}

func (template *Template) fillBackground() {
	template.Image.SetHexColor("#ffffff")
	template.Image.DrawRectangle(
		0, 0,
		float64(template.Image.Width()),
		float64(template.Image.Height()),
	)
	template.Image.Fill()
}

func (template *Template) drawFrame() {
	template.Image.SetHexColor("#000000")
	template.Image.DrawRectangle(
		float64((template.Image.Width() / 2) -
			(template.FrameConfig.LineWidth / 2)),
		float64(template.FrameConfig.TopHeight +
			template.TextConfig.MarginTop +
			template.TextConfig.MarginBottom),
		float64(template.FrameConfig.LineWidth),
		float64(template.Image.Height()),
	)
	template.Image.Fill()
	template.Image.DrawRectangle(
		0, float64(template.FrameConfig.TopHeight +
			template.TextConfig.MarginTop +
			template.TextConfig.MarginBottom),
		float64(template.Image.Width()),
		float64(template.FrameConfig.LineWidth),
	)
	template.Image.Fill()
	template.Image.DrawRectangle(
		0, float64(template.FrameConfig.TopHeight +
			template.TextConfig.MarginTop +
			template.TextConfig.MarginBottom +
			template.FrameConfig.LineWidth +
			template.FrameConfig.MiddleHeight),
		float64(template.Image.Width()),
		float64(template.FrameConfig.LineWidth),
	)
	template.Image.Fill()
}

func (blocks *Blocks) resizeSrcImages() {
	blocks.SrcImagesConfig.LeftImage.Image = resize.Resize(
		uint(blocks.TemplateConfig.FrameConfig.ImageWidth),
		uint(blocks.TemplateConfig.FrameConfig.ImageHeight),
		blocks.SrcImagesConfig.LeftImage.Image, resize.Lanczos3,
	)
	blocks.SrcImagesConfig.RightImage.Image = resize.Resize(
		uint(blocks.TemplateConfig.FrameConfig.ImageWidth),
		uint(blocks.TemplateConfig.FrameConfig.ImageHeight),
		blocks.SrcImagesConfig.RightImage.Image, resize.Lanczos3,
	)
}

func (template *Template) RenderTemplate() {
	template.makeTemplateImage()
	template.fillBackground()
	template.drawFrame()
}

func (blocks *Blocks) RenderSrcImage() {
	blocks.resizeSrcImages()
	blocks.TemplateConfig.Image.DrawImage(
		blocks.SrcImagesConfig.LeftImage.Image, 0,
		blocks.TemplateConfig.FrameConfig.TopHeight+
			blocks.TemplateConfig.TextConfig.MarginTop+
			blocks.TemplateConfig.TextConfig.MarginBottom+
			blocks.TemplateConfig.FrameConfig.MiddleHeight+
			(blocks.TemplateConfig.FrameConfig.LineWidth*2),
	)
	blocks.TemplateConfig.Image.DrawImage(
		blocks.SrcImagesConfig.RightImage.Image,
		blocks.TemplateConfig.FrameConfig.ImageWidth +
			blocks.TemplateConfig.FrameConfig.LineWidth,
		blocks.TemplateConfig.FrameConfig.TopHeight+
			blocks.TemplateConfig.TextConfig.MarginTop+
			blocks.TemplateConfig.TextConfig.MarginBottom+
			blocks.TemplateConfig.FrameConfig.MiddleHeight+
			(blocks.TemplateConfig.FrameConfig.LineWidth*2),
	)
}
