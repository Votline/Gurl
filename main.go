package main

import (
	"log"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/gl/v4.1-core/gl"

	"Glur/internal/ui"
	"Glur/internal/font"
	"Glur/internal/shapes"
	"Glur/internal/render"
)

func Init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("GLFW init error. \nErr: ", err)
	}
	defer glfw.Terminate()
	if err := gl.Init(); err != nil {
		log.Fatalln("OpenGL init error. \nErr: ", err)
	}

	win := ui.PrimaryWindow()

	vao := render.Canvas(shapes.QuadUV())
	img := font.RenderTextImage("Hello boss")
	tex := render.TextureImage(img)

	pg := render.Setup()

	for !win.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.UseProgram(pg)

		gl.BindVertexArray(vao)

		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, tex)

		loc := gl.GetUniformLocation(pg, gl.Str("tex\x00"))
		gl.Uniform1i(loc, 0)

		gl.DrawArrays(gl.TRIANGLES, 0, 6)

		win.SwapBuffers()
		glfw.PollEvents()
	}
}
