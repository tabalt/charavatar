package charavatar

import (
	"errors"
	"image/color"
	"strconv"
	"strings"
)

type Colors map[string]color.Color

func NewColors(codes ...string) (*Colors, error) {
	c := &Colors{}
	for _, v := range codes {
		if err := c.Add(v); err != nil {
			return nil, err
		}
	}
	return c, nil
}

// add color by code
func (c *Colors) Add(code string) error {
	code = strings.ToUpper(code)
	rgba, err := c.CodeToRGBA(code)
	if err != nil {
		return err
	}
	colors := *c
	colors[code] = rgba
	return nil
}

// transform color code to rgba object
func (c *Colors) CodeToRGBA(code string) (*color.RGBA, error) {
	const codePrefix = "#"
	const codeLength = 7

	if !strings.HasPrefix(code, codePrefix) {
		return nil, errors.New("color code must has prefix '" + codePrefix + "'")
	}

	if len(code) != codeLength {
		return nil, errors.New("color code length must be '" + strconv.Itoa(codeLength) + "'")
	}

	code = strings.TrimPrefix(code, codePrefix)
	code64, err := strconv.ParseInt(code, 16, 32)
	if err != nil {
		return nil, err
	}
	code32 := int(code64)

	red := uint8(code32 >> 16)
	green := uint8((code32 & 0x00FF00) >> 8)
	blue := uint8(code32 & 0x0000FF)

	return &color.RGBA{red, green, blue, 255}, nil
}

// get color by rand
func (c *Colors) GetByRand() color.Color {
	colors := *c
	for _, v := range colors {
		return v
	}
	return nil
}

// get color by key
func (c *Colors) GetByKey(key string) color.Color {
	key = strings.ToUpper(key)
	colors := *c
	return colors[key]
}
