package input

import (
	"github.com/go-gl/gl/v4.1-core/gl"

	"Glur/internal/font"
	"Glur/internal/shapes"
	"Glur/internal/render"
)

type Input interface {
	Render()
}

type InputField struct {
	vao uint32
	pg  uint32
	tex uint32
	posX, posY     float32
	scaleX, scaleY float32
}

func NewInputField(sharedPG uint32, text string, x, y, scaleX, scaleY float32) *InputField {
	img := font.RenderTextImage(text)
	tex := render.TextureImage(img)
	vao := render.Canvas(shapes.QuadUV())

	return &InputField {
		vao: vao,
		pg: sharedPG,
		tex: tex,
		posX: x,
		posY: y,
		scaleX: scaleX,
		scaleY: scaleY,
	}

}

func (f *InputField) Render() {
	gl.BindVertexArray(f.vao)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, f.tex)

	loc := gl.GetUniformLocation(f.pg, gl.Str("tex\x00"))
	gl.Uniform1i(loc, 0)

	locCol := gl.GetUniformLocation(f.pg, gl.Str("color\x00"))
	gl.Uniform4f(locCol, 1.0, 0.0, 0.0, 1.0)

	locOff := gl.GetUniformLocation(f.pg, gl.Str("offset\x00"))
	gl.Uniform2f(locOff, f.posX, f.posY)
	
	locScal := gl.GetUniformLocation(f.pg, gl.Str("scale\x00"))
	gl.Uniform2f(locScal, f.scaleX, f.scaleY)

	gl.DrawArrays(gl.TRIANGLES, 0, 6)
}
