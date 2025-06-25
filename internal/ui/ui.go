package ui

import (
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"
)

const windowWidth = 200
const windowHeight = 90

func PrimaryWindow() *glfw.Window {
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.Decorated, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.TransparentFramebuffer, glfw.True)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCompatProfile)
	win, err := glfw.CreateWindow(windowWidth, windowHeight, "Glur", nil, nil)
	if err != nil {
		log.Println("Create window error. \nErr: ", err)
		win = nil
	}
	win.MakeContextCurrent()
	win.SetAttrib(glfw.Floating, 1)
	vidMode := glfw.GetPrimaryMonitor().GetVideoMode()
	win.SetPos(vidMode.Width-220, vidMode.Height-1075)

	glfw.SwapInterval(1)
	return win
}
