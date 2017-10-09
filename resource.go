package charavatar

import ()

type Resource struct {
	FrontColors      *Colors
	BackgroundColors *Colors

	Fonts *Fonts
}

func NewResource(fcs, bcs []string, fontPath string, fontFiles []string) (*Resource, error) {
	frontColors, err := NewColors(fcs...)
	if err != nil {
		return nil, err
	}

	backgroundColors, err := NewColors(bcs...)
	if err != nil {
		return nil, err
	}

	fonts, err := NewFonts(fontPath, fontFiles...)
	if err != nil {
		return nil, err
	}

	return &Resource{frontColors, backgroundColors, fonts}, nil
}
