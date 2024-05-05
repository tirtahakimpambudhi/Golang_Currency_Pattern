package test

import (
	"go_pipeline/channel"
	"testing"
)

func TestPipeline(t *testing.T) {
	path := []channel.InputPath{
		channel.InputPath("./img/giga_gophers.jpeg"),
		channel.InputPath("./img/super_gophers.jpeg"),
	}
	loadImage := channel.LoadImage(path)
	resizeImage := channel.Resize(loadImage)
	grayImage := channel.ConvertGrayScale(resizeImage)
	result := channel.SaveImage(grayImage)
	for r := range result {
		if r {
			t.Log(r)
		} else {
			t.Fail()
		}
	}
}
