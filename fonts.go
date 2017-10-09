package charavatar

import (
	"io/ioutil"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

type Fonts map[string]*truetype.Font

func NewFonts(path string, files ...string) (*Fonts, error) {
	f := &Fonts{}
	for _, v := range files {
		if err := f.Add(path, v); err != nil {
			return nil, err
		}
	}
	return f, nil
}

// add font by path and file
func (f *Fonts) Add(path, file string) error {

	font, err := f.FileToFont(path + file)
	if err != nil {
		return err
	}

	fonts := *f
	fonts[file] = font
	return nil
}

// transform font file data to font object
func (f *Fonts) FileToFont(file string) (*truetype.Font, error) {
	fontdata, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	font, err := freetype.ParseFont(fontdata)
	if err != nil {
		return nil, err
	}

	return font, nil
}

// get font by rand
func (f *Fonts) GetByRand() *truetype.Font {
	fonts := *f
	for _, v := range fonts {
		return v
	}
	return nil
}

// get font by key
func (f *Fonts) GetByKey(key string) *truetype.Font {
	fonts := *f
	return fonts[key]
}
