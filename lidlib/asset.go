package lidlib

import (
	"image/jpeg"
	"io"
	"log"
	"os"
	"path"

	"github.com/nfnt/resize"
)

type Asset struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func NewAsset(path, filename string, file io.Reader) (*Asset, error) {
	return nil, nil
}

func (a Asset) Thumbnail() {
	fp := path.Join(a.Path, a.Name)

	file, err := os.Open(fp)
	if err != nil {
		log.Println(err)
		return
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Println(err)
		return
	}

	out := resize.Resize(1000, 0, img, resize.Bilinear)
	defer file.Close()

	jpeg.Encode(file, out, nil)
}

func (a Asset) Resample() {
	fp := path.Join(a.Path, a.Name)

	file, err := os.Open(fp)
	if err != nil {
		log.Println(err)
		return
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Println(err)
		return
	}

	out := resize.Resize(300, 0, img, resize.Bilinear)

	resampledDir := path.Join(a.Path, "_vorschaubilder")
	os.MkdirAll(resampledDir, 0755)

	resampledFp := path.Join(resampledDir, a.Name)
	file, err = os.Create(resampledFp)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	jpeg.Encode(file, out, nil)
}
