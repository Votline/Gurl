package render

import (
	"image"

	"github.com/go-gl/gl/v4.1-core/gl"

	"Glur/internal/shaders"
)

func Setup() uint32 {
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	pg := gl.CreateProgram()
	shaders.AddShaders(pg)
	gl.LinkProgram(pg)

	gl.LineWidth(3.0)
	gl.UseProgram(pg)

	gl.ClearColor(0.0, 0.0, 0.0, 0.7)

	return pg
}

func Canvas(quad []float32) uint32 {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(
		gl.ARRAY_BUFFER, len(quad)*4,
		gl.Ptr(quad), gl.STATIC_DRAW)

	gl.VertexAttribPointer(
		0, 3, gl.FLOAT,
		false, 5*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	gl.VertexAttribPointer(
		1, 2, gl.FLOAT,
		false, 5*4, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(1)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	return vao
}

func TextureImage(img *image.RGBA) uint32 {
	var tex uint32
	gl.GenTextures(1, &tex)
	gl.BindTexture(gl.TEXTURE_2D, tex)

	width := int32(img.Rect.Size().X)
	height := int32(img.Rect.Size().Y)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA,
		width, height, 0,
		gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(img.Pix))
	
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S,
		gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T,
		gl.CLAMP_TO_EDGE)
	
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER,
		gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER,
		gl.LINEAR)

	gl.BindTexture(gl.TEXTURE_2D, 0)

	return tex

}
