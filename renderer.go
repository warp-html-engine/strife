package strife

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type Style uint

const (
	Line Style = iota
	Fill
)

type RenderConfig struct {
	Alias        bool
	Accelerated  bool
	VerticalSync bool
}

func DefaultConfig() *RenderConfig {
	return &RenderConfig{
		Alias:        true,
		Accelerated:  true,
		VerticalSync: true,
	}
}

type Renderer struct {
	RenderConfig
	*sdl.Renderer

	color *Color
	font  *Font
}

func (r *Renderer) Clear() {
	r.SetColor(Black)
	w, h, err := r.Renderer.GetRendererOutputSize()
	if err != nil {
		panic(err)
	}
	r.Rect(0, 0, w, h, Fill)

	r.SetColor(White)
}

func (r *Renderer) Display() {
	r.Renderer.Present()
}

func (r *Renderer) SetColor(col *Color) {
	r.color = col
}

func (r *Renderer) Rect(x, y, w, h int, mode Style) {
	color := r.color
	r.SetDrawColor(color.R, color.G, color.B, color.A)

	if mode == Line {
		r.DrawRect(&sdl.Rect{int32(x), int32(y), int32(w), int32(h)})
	} else {
		r.FillRect(&sdl.Rect{int32(x), int32(y), int32(w), int32(h)})
	}
}

func (r *Renderer) SetFont(font *Font) {
	r.font = font
}

func (r *Renderer) String(message string, x, y int) {
	if r.font == nil {
		panic("Attempted to render '" + message + "' but no font is set!")
	}

	var surface *sdl.Surface
	var err error
	if r.Alias {
		surface, err = r.font.RenderUTF8_Blended(message, r.color.ToSDLColor())
	} else {
		surface, err = r.font.RenderUTF8_Solid(message, r.color.ToSDLColor())
	}
	defer surface.Free()
	if err != nil {
		panic(err)
	}

	texture, err := r.Renderer.CreateTextureFromSurface(surface)
	defer texture.Destroy()
	if err != nil {
		panic(err)
	}
	r.Renderer.Copy(texture, nil, &sdl.Rect{int32(x), int32(y), surface.W, surface.H})
}

func (r *Renderer) Image(image *Image, x, y int) {
	_, _, w, h, err := image.Texture.Query()
	if err != nil {
		panic(err)
	}
	r.Copy(image.Texture, nil, &sdl.Rect{int32(x), int32(y), int32(w), int32(h)})
}

func CreateRenderer(parent *RenderWindow, config *RenderConfig) (*Renderer, error) {
	var mode uint32
	if config.Accelerated {
		mode |= sdl.RENDERER_ACCELERATED
	} else {
		mode |= sdl.RENDERER_SOFTWARE
	}
	if config.VerticalSync {
		mode |= sdl.RENDERER_PRESENTVSYNC
	}

	renderInst, err := sdl.CreateRenderer(parent.Window, -1, mode)
	if err != nil {
		return nil, fmt.Errorf("Failed to create render context")
	}

	renderer := &Renderer{
		RenderConfig: *config,
		Renderer:     renderInst,
		color:        RGB(255, 255, 255),
	}
	return renderer, nil
}