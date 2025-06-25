package shaders

import (
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
)

const vertexShaderSource = `
#version 410 core
layout (location = 0) in vec3 aPos;
layout (location = 1) in vec2 aTex;

out vec2 TexCoord;

void main() {
	gl_Position = vec4(aPos, 1.0);
	TexCoord = aTex;
}` + "\x00"

const fragmentShaderSource = `
#version 410 core
in vec2 TexCoord;
out vec4 FragColor;

uniform sampler2D tex;

void main() {
	FragColor = texture(tex, TexCoord);
}` + "\x00"


func AddShaders(pg uint32) {
	vertexShader := compileShader(
		vertexShaderSource, gl.VERTEX_SHADER)
	fragmentShader := compileShader(
		fragmentShaderSource, gl.FRAGMENT_SHADER)

	gl.AttachShader(pg, vertexShader)
	gl.AttachShader(pg, fragmentShader)
}

func compileShader(source string, shaderType uint32) uint32 {
	shader := gl.CreateShader(shaderType)
	cSource, freeMemory := gl.Strs(source)
	gl.ShaderSource(shader, 1, cSource, nil)
	
	freeMemory()
	gl.CompileShader(shader)

	checkShaderStatus(shader)

	return shader
}

func checkShaderStatus(shader uint32) {
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)
		
		logMsg := make([]byte, logLength)
		gl.GetShaderInfoLog(shader, logLength, nil, &logMsg[0])

		log.Fatalf("\nShader compile error: \nType: %d\n%s\n",
			logLength, logMsg)
	}
}
