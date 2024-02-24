package imagex

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"math"

	"github.com/nfnt/resize"
)

func ResizeImage(file io.Reader, originSize int) (outFile io.Reader, imgWidth int, imgHeight int, err error) {
	img, form, err := image.Decode(file)
	if err != nil {
		return
	}

	var newImg image.Image

	if originSize >= 9*1024*1024 {
		width := img.Bounds().Size().X

		targetSize := 4 * 1024 * 1024
		rate := math.Sqrt(float64(targetSize) / float64(originSize))

		newWidth := uint(float64(width) * rate)

		newImg = resize.Resize(newWidth, 0, img, resize.Lanczos3)
	} else {
		newImg = img
	}

	imgWidth = newImg.Bounds().Size().X
	imgHeight = newImg.Bounds().Size().Y

	pipeReader, pipeWriter := io.Pipe()

	go func() {
		defer pipeWriter.Close()

		switch form {
		case "jpeg":
			jpeg.Encode(pipeWriter, newImg, nil)
		case "png":
			png.Encode(pipeWriter, newImg)
		case "gif":
			gif.Encode(pipeWriter, newImg, nil)
		}
	}()

	return pipeReader, imgWidth, imgHeight, nil

}
