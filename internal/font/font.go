package font

import (
	"log"
	"image"
	"image/color"
	"io/ioutil"

	"golang.org/x/image/font"
	"github.com/golang/freetype"
)

const fontSize = 40
const fontPath = "assets/NotoSans-Regular.ttf"

func RenderTextImage(txt string) *image.RGBA {
	fontBytes, err := ioutil.ReadFile(fontPath)
	if err != nil {
		log.Fatalln("Read ttf file error. \nErr: ", err)
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Fatalln("Parse font error. \nErr: ", err)
	}

	img := image.NewRGBA(image.Rect(0, 0, 200, 90))
	drawer := freetype.NewContext()
	drawer.SetClip(img.Bounds())
	
	drawer.SetDPI(72)
	drawer.SetFont(f)
	drawer.SetFontSize(fontSize)
	drawer.SetDst(img)
	drawer.SetSrc(image.NewUniform(color.White))
	drawer.SetHinting(font.HintingFull)

	pt := freetype.Pt(10, 45)
	_, err = drawer.DrawString(txt, pt)
	if err != nil {
		log.Fatalln("Failed to draw string. \nErr: ", err)
	}

	return img
}
