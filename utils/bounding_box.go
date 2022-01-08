package utils

import (
	"github.com/fogleman/gg"
)

type Box struct {
	// clarifai api bounding box coordinates
	TopRow    float64
	RightCol  float64
	BottomRow float64
	LeftCol   float64
	LineWidth float64
	LineColor string
}

func DrawBox(box Box) {
	// set Box default values
	if box.LineColor == "" {
		box.LineColor = "#0488D0" // shade of blue
	}
	if box.LineWidth == 0 {
		box.LineWidth = 7.5
	}
	// load image from path
	im, err := gg.LoadPNG("face.png")
	if err != nil {
		panic(err)
	}
	// get image width
	w := im.Bounds().Size().X
	// get image heigth
	h := im.Bounds().Size().Y
	// draw the canvas
	dc := gg.NewContextForImage(im)
	dc.DrawImage(im, w, h)
	dc.SetHexColor(box.LineColor)

	// draw left column -- start
	x1 := float64(box.LeftCol) * float64(w)
	y1 := float64(box.TopRow) * float64(h)
	x2 := x1
	y2 := float64(box.BottomRow) * float64(h)

	dc.SetLineWidth(box.LineWidth)
	dc.DrawLine(x1, y1, x2, y2)
	dc.Stroke()
	// draw left column -- end

	// draw top column -- start
	x1 = float64(box.LeftCol) * float64(w)
	y1 = float64(box.TopRow) * float64(h)
	x2 = float64(box.RightCol) * float64(w)
	y2 = y1

	dc.SetLineWidth(float64(box.LineWidth))
	dc.DrawLine(x1, y1, x2, y2)
	dc.Stroke()
	// draw top column -- end

	// draw right column -- start
	x1 = float64(box.RightCol) * float64(w)
	y1 = float64(box.TopRow) * float64(h)
	x2 = x1
	y2 = float64(box.BottomRow) * float64(h)

	dc.SetLineWidth(box.LineWidth)
	dc.DrawLine(x1, y1, x2, y2)
	dc.Stroke()
	// draw right column -- end

	// draw bottom column -- start
	x1 = float64(box.LeftCol) * float64(w)
	y1 = float64(box.BottomRow) * float64(h)
	x2 = float64(box.RightCol) * float64(w)
	y2 = y1

	dc.SetLineWidth(box.LineWidth)
	dc.DrawLine(x1, y1, x2, y2)
	dc.Stroke()
	// draw bottom column -- end

	dc.SavePNG("out.png")
}

func CropBox(box Box) {

}
