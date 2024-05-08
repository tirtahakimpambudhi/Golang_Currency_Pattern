package channel

import (
	"go_pipeline/pkg"
	"image"
	"strings"
)

type InputPath string
type Job struct {
	Path       InputPath
	Image      image.Image
	OutputPath string
	Err        error
}

func LoadImage(paths []InputPath) <-chan Job {
	out := make(chan Job)
	go func() {
		defer close(out)
		for _, path := range paths {
			job := Job{
				Path:       InputPath(path),
				OutputPath: strings.Replace(string(path), "img/", "img/output/", 1),
			}

			job.Image, job.Err = pkg.ReadImage(string(path))
			out <- job
		}
		//if chan not close the statment waiting channel because not close
		// error deadlock

	}()
	return out
}

func Resize(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		defer close(out)
		for i := range input {

			i.Image = pkg.ResizeImage(500, 500, i.Image)

			out <- i
		}

	}()
	return out
}

func ConvertGrayScale(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		defer close(out)
		for i := range input {

			i.Image = pkg.GrayImage(i.Image)

			out <- i
		}

	}()
	return out
}

func SaveImage(input <-chan Job) <-chan bool {
	out := make(chan bool)
	go func() {
		defer close(out)
		for i := range input {
			err := pkg.WriteImage(i.OutputPath, i.Image)
			out <- err == nil
		}

	}()
	return out
}
