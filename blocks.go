package blocks

import (
	"bytes"
	"image"
)

func New(srcImages [2]image.Image, texts [3]string) *Blocks {
	return &Blocks{
		TemplateConfig: &Template{
			TextConfig: &Text{
				Texts: texts,
				FontConfig: &Font{},
			},
			FrameConfig: &Frame{
				LineWidth: 2,
				TopHeight: 30,
				MiddleHeight: 30,
				ImageWidth: 250,
				ImageHeight: 250,
			},
		},
		SrcImagesConfig: &SrcImages{
			LeftImage: &SrcImage{
				Image: srcImages[0],
				Width: srcImages[0].Bounds().Size().X,
				Height: srcImages[0].Bounds().Size().Y,
			},
			RightImage: &SrcImage{
				Image: srcImages[1],
				Width: srcImages[1].Bounds().Size().X,
				Height: srcImages[1].Bounds().Size().Y,
			},
		},
	}
}

func (blocks *Blocks) GetImageBuffer() (imgBuffer *bytes.Buffer, err error) {
	imgBuffer = &bytes.Buffer{}
	if err = blocks.TemplateConfig.Image.EncodePNG(imgBuffer); err != nil { return }
	return
}

func (blocks *Blocks) Make() (imgBytes []byte, err error) {
	blocks.TemplateConfig.RenderTemplate()
	blocks.RenderSrcImage()
	if err = blocks.TemplateConfig.RenderTexts(); err != nil { return }

	imgBuffer, err := blocks.GetImageBuffer()
	if err != nil { return }
	imgBytes = imgBuffer.Bytes()
	return
}

