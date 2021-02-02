package vendors

import (
	"image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"
)

// ResizeImage ..
func ResizeImage(imageName string, outputName string, path string) {
	file, err := os.Open(path + imageName)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(200, 0, img, resize.Lanczos3)

	out, err := os.Create(path + outputName)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)

	// Delete Old Image
	e := os.Remove(path + imageName)
	if e != nil {
		log.Fatal(e)
	}
}
