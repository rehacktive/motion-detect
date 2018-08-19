package motion

import (
	"errors"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

const (
	thresold    = 10
	sensitivity = 50
	minArea     = 100000
)

func init() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

// DetectMotion comparing the two images, with a thresold
func DetectMotion(image1 string, image2 string, output string) (bool, error) {
	img1, err := getImage(image1)
	if err != nil {
		return false, err
	}
	img2, err := getImage(image2)
	if err != nil {
		return false, err
	}

	b1 := img1.Bounds()
	b2 := img2.Bounds()

	if b1.Dx() != b2.Dx() || b1.Dy() != b2.Dy() {
		return false, errors.New("different images sizes")
	}

	out := image.NewRGBA(image.Rect(0, 0, b1.Dx(), b1.Dy()))
	diff := 0
	for i := 0; i < b1.Dx(); i++ {
		for j := 0; j < b1.Dy(); j++ {
			c, d := colorDiff(img1.At(i, j), img2.At(i, j))
			if d > sensitivity {
				diff++
			}
			out.Set(i, j, c)
		}
	}

	f, _ := os.OpenFile(output, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	jpeg.Encode(f, out, nil)
	return diff > minArea, err
}

func colorDiff(color1 color.Color, color2 color.Color) (color.Color, int) {
	gray1, _, _, _ := color.GrayModel.Convert(color1).RGBA()
	gray2, _, _, _ := color.GrayModel.Convert(color2).RGBA()
	diff := abs(int(gray1>>8) - int(gray2>>8))
	if diff > thresold {
		return color.Gray{uint8(diff)}, diff
	}
	return color.Black, 0
}

func getImage(imagePath string) (image.Image, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
