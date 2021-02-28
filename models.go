package blocks

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"image"
)

type Font struct {
	Font				*truetype.Font
	FontSize			int
}

type Text struct {
	FontConfig		*Font

	Texts			[3]string
	MarginTop		int
	MarginBottom	int
}

type Frame struct {
	LineWidth		int
	TopHeight		int
	MiddleHeight	int
	ImageHeight		int
	ImageWidth		int
}

type SrcImage struct {
	Image	image.Image

	Width	int
	Height	int
}

type SrcImages struct {
	LeftImage	*SrcImage
	RightImage	*SrcImage
}

type Template struct {
	TextConfig		*Text
	FrameConfig		*Frame
	Image			*gg.Context

	Width			int
	Height			int
}

type Blocks struct {
	TemplateConfig		*Template
	SrcImagesConfig		*SrcImages

	configsConfigured	bool
}
