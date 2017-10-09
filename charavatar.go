package charavatar

import ()

var (
	defaultResource *Resource
)

func Init(fcs, bcs []string, fontPath string, fontFiles []string) error {
	var err error
	defaultResource, err = NewResource(fcs, bcs, fontPath, fontFiles)
	if err != nil {
		return err
	}
	return nil
}

func BuildRandImage(w, h int, char rune) (*Image, error) {
	font := defaultResource.Fonts.GetByRand()
	frontColor := defaultResource.FrontColors.GetByRand()
	backgroundColor := defaultResource.BackgroundColors.GetByRand()

	img := NewAvatar(font, frontColor, backgroundColor, w, h, char).BuildImage()

	return img, nil
}

func BuildImage(f, fc, bc string, w, h int, char rune) (*Image, error) {
	font := defaultResource.Fonts.GetByKey(f)
	frontColor := defaultResource.FrontColors.GetByKey(fc)
	backgroundColor := defaultResource.BackgroundColors.GetByKey(bc)

	var img *Image
	var err error

	if font != nil && frontColor != nil && backgroundColor != nil {
		img = NewAvatar(font, frontColor, backgroundColor, w, h, char).BuildImage()
	} else {
		img, err = BuildRandImage(w, h, char)
	}

	return img, err
}
