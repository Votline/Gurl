package main

import (
	"log"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/gl/v4.1-core/gl"

	"Glur/internal/ui"
	"Glur/internal/views"
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
	pg := render.Setup()

	hv := home.NewHomeView(pg)

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	for !win.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.UseProgram(pg)

		hv.Render()

		if err := gl.GetError(); err != gl.NO_ERROR {
			log.Println("OpenGL error. \nErr: ", err)
		}

		win.SwapBuffers()
		glfw.PollEvents()
	}
}
