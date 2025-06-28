package home

import (
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"

	"Glur/internal/font"
	"Glur/internal/shapes"
	"Glur/internal/render"
	"Glur/internal/components"
)

type HomeView struct {
	vao uint32
	tex uint32
	pg  uint32
	inputs []input.Input
}
func NewHomeView(pg uint32) *HomeView {
	vao := render.Canvas(shapes.QuadUV())
	img := font.RenderTextImage("Glur")
	tex := render.TextureImage(img)
	inputs := make([]input.Input, 0)

	hv := HomeView{
		vao: vao,
		tex: tex,
		pg: pg,
		inputs: inputs,
	}
	hv.AddInput("URL here", -0.5, 0.0, 0.5, 0.3)
	hv.AddInput("Nothing here", 0.5, 0.0, 0.52, 0.3)
	return &hv
}

func (hv *HomeView) AddInput(text string, x, y, scaleX, scaleY float32) {
	inputField := input.NewInputField(
		hv.pg, text, x, y, scaleX, scaleY)
	hv.inputs = append(hv.inputs, inputField)
	log.Printf("\n%v\n%v/%v\n%v/%v", text, x, y, scaleX, scaleY)
}

func (hv *HomeView) Render() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(hv.pg)
	
	for _, inp := range hv.inputs {
		inp.Render()
	}

	gl.BindVertexArray(hv.vao)
	gl.ActiveTexture(gl.TEXTURE0)
	
	gl.BindTexture(gl.TEXTURE_2D, hv.tex)
	
	loc := gl.GetUniformLocation(hv.pg, gl.Str("tex\x00"))
	gl.Uniform1i(loc, 0)
	
	locCol := gl.GetUniformLocation(hv.pg, gl.Str("color\x00"))
	gl.Uniform4f(locCol, 1.0, 1.0, 1.0, 1.0)

	locOff := gl.GetUniformLocation(hv.pg, gl.Str("offset\x00"))
	gl.Uniform2f(locOff, 0.0, 0.6)

	locScal := gl.GetUniformLocation(hv.pg, gl.Str("scale\x00"))
	gl.Uniform2f(locScal, 0.3, 0.3)

	gl.DrawArrays(gl.TRIANGLES, 0, 6)



}
