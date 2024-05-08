package pkg

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

func ReadImage(path string) (image.Image,error) {
	file , err := os.Open(path)
	if err != nil {
		return nil, err
	}
	img,_,err  := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img,nil
}

func WriteImage(path string,img image.Image) error {
	file,err  := os.Create(path)
	if err != nil {
		return err
	}
	return jpeg.Encode(file,img,&jpeg.Options{Quality: jpeg.DefaultQuality})
}

func GrayImage(img image.Image) image.Image {
	bounds := img.Bounds()
	//set size rectangle img
	grayImg := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			//to original image pixel at x,y
			originalImg := img.At(x,y)
			// original image pixel to gray pixel
			grayPixel := color.GrayModel.Convert(originalImg)
			// fill gray img at x,y use gray pixel
			grayImg.Set(x,y,grayPixel)
		}
	}
	return grayImg
}

func ResizeImage(width, height int,img image.Image) image.Image {
	return resize.Resize(uint(width),uint(height),img,resize.Lanczos3)
}