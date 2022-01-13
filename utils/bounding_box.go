package utils

import (
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
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

func DrawBox(imgUrl string, box []Box) {
	// load image from path
	im, err := gg.LoadImage(imgUrl)
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

	for _, b := range box {
		// set Box default values
		if b.LineColor == "" {
			b.LineColor = "#0488D0" // shade of blue
		}
		if b.LineWidth == 0 {
			b.LineWidth = 7.5
		}
		dc.SetHexColor(b.LineColor)

		// draw left column -- start
		x1 := float64(b.LeftCol) * float64(w)
		y1 := float64(b.TopRow) * float64(h)
		x2 := x1
		y2 := float64(b.BottomRow) * float64(h)

		dc.SetLineWidth(b.LineWidth)
		dc.DrawLine(x1, y1, x2, y2)
		dc.Stroke()
		// draw left column -- end

		// draw top column -- start
		x1 = float64(b.LeftCol) * float64(w)
		y1 = float64(b.TopRow) * float64(h)
		x2 = float64(b.RightCol) * float64(w)
		y2 = y1

		dc.SetLineWidth(float64(b.LineWidth))
		dc.DrawLine(x1, y1, x2, y2)
		dc.Stroke()
		// draw top column -- end

		// draw right column -- start
		x1 = float64(b.RightCol) * float64(w)
		y1 = float64(b.TopRow) * float64(h)
		x2 = x1
		y2 = float64(b.BottomRow) * float64(h)

		dc.SetLineWidth(b.LineWidth)
		dc.DrawLine(x1, y1, x2, y2)
		dc.Stroke()
		// draw right column -- end

		// draw bottom column -- start
		x1 = float64(b.LeftCol) * float64(w)
		y1 = float64(b.BottomRow) * float64(h)
		x2 = float64(b.RightCol) * float64(w)
		y2 = y1

		dc.SetLineWidth(b.LineWidth)
		dc.DrawLine(x1, y1, x2, y2)
		dc.Stroke()
		// draw bottom column -- end
	}

	dc.SavePNG("out.png")
}

// func CropBox(imgUrl string, box []Box) {
// 	// load image from path
// 	im, err := gg.LoadImage(imgUrl)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// get image width
// 	w := im.Bounds().Size().X
// 	// get image heigth
// 	h := im.Bounds().Size().Y
// 	// draw the canvas
// 	dc := gg.NewContextForImage(im)
// 	dc.DrawImage(im, w, h)

// 	for _, b := range box {
// 		// set Box default values
// 		if b.LineColor == "" {
// 			b.LineColor = "#0488D0" // shade of blue
// 		}
// 		if b.LineWidth == 0 {
// 			b.LineWidth = 7.5
// 		}
// 		dc.SetHexColor(b.LineColor)

// 		// draw left column -- start
// 		x1 := float64(b.LeftCol) * float64(w)
// 		y1 := float64(b.TopRow) * float64(h)
// 		x2 := x1
// 		y2 := float64(b.BottomRow) * float64(h)

// 		dc.SetLineWidth(b.LineWidth)
// 		dc.DrawLine(x1, y1, x2, y2)
// 		dc.Stroke()
// 		// draw left column -- end

// 		// draw top column -- start
// 		x1 = float64(b.LeftCol) * float64(w)
// 		y1 = float64(b.TopRow) * float64(h)
// 		x2 = float64(b.RightCol) * float64(w)
// 		y2 = y1

// 		dc.SetLineWidth(float64(b.LineWidth))
// 		dc.DrawLine(x1, y1, x2, y2)
// 		dc.Stroke()
// 		// draw top column -- end

// 		// draw right column -- start
// 		x1 = float64(b.RightCol) * float64(w)
// 		y1 = float64(b.TopRow) * float64(h)
// 		x2 = x1
// 		y2 = float64(b.BottomRow) * float64(h)

// 		dc.SetLineWidth(b.LineWidth)
// 		dc.DrawLine(x1, y1, x2, y2)
// 		dc.Stroke()
// 		// draw right column -- end

// 		// draw bottom column -- start
// 		x1 = float64(b.LeftCol) * float64(w)
// 		y1 = float64(b.BottomRow) * float64(h)
// 		x2 = float64(b.RightCol) * float64(w)
// 		y2 = y1

// 		dc.SetLineWidth(b.LineWidth)
// 		dc.DrawLine(x1, y1, x2, y2)
// 		dc.Stroke()
// 		// draw bottom column -- end
// 	}

// 	dc.SavePNG("out.png")

// }

func readImage(name string) (image.Image, error) {
	fd, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	// image.Decode requires that you import the right image package. We've
	// imported "image/png", so Decode will work for png files. If we needed to
	// decode jpeg files then we would need to import "image/jpeg".
	//
	// Ignored return value is image format name.
	img, _, err := image.Decode(fd)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// writeImage writes an Image back to the disk.
func writeImage(img image.Image, name string) error {
	fd, err := os.Create(name)
	if err != nil {
		return err
	}
	defer fd.Close()

	return png.Encode(fd, img)
}

func CropImage(name string, box Box, outputName string) {
	im, err := readImage(name)
	if err != nil {
		panic(err)
	}

	// get image width
	w := im.Bounds().Size().X
	// get image heigth
	h := im.Bounds().Size().Y
	// cropping coordinates
	crop := image.Rect(int(box.LeftCol*float64(w)), int(box.TopRow*float64(h)), int(box.RightCol*float64(w)), int(box.BottomRow*float64(h)))

	type subImager interface {
		SubImage(r image.Rectangle) image.Image
	}

	// img is an Image interface. This checks if the underlying value has a
	// method called SubImage. If it does, then we can use SubImage to crop the
	// image.
	cim, ok := im.(subImager)
	if !ok {
		panic("not supported")
	}

	writeImage(cim.SubImage(crop), outputName)
}

func DrawText(path string) {
	TTF, err := ioutil.ReadFile(path)
	font, err := truetype.Parse(TTF)
	if err != nil {
		log.Fatal(err)
	}

	face := truetype.NewFace(font, &truetype.Options{Size: 48})

	h := 250
	w := 500
	dc := gg.NewContext(w, h)
	dc.DrawRectangle(0, 0, float64(w), float64(h))
	dc.SetHexColor("#ffffff")
	dc.Fill()
	dc.Clear()
	dc.SetHexColor("#00")
	dc.SetFontFace(face)
	dc.DrawStringAnchored("Hello, world!", 50, 125, 0, 0)
	dc.SavePNG("text.png")
}
