package render

import (
	"github.com/go-gl/gl/v4.1-core/gl"

	"Glur/internal/shaders"
)

func Setup() uint32 {
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	pg := gl.CreateProgram()
	shaders.AddShaders(pg)
	gl.LinkProgram(pg)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	gl.EnableVertexAttribArray(0)

	gl.LineWidth(3.0)
	gl.UseProgram(pg)

	gl.ClearColor(0.0, 0.0, 0.0, 0.7)

	return pg
}
