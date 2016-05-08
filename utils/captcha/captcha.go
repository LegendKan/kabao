package captcha

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"kabao/third/freetype"
	"math/rand"
	"time"
)

const (
	fontPath = "../../static/fonts/"
	fontFile = "Arcade Book.ttf"
	fontSize = 12
	spacing  = 1.5
)

//var font
var context *freetype.Context

func init() {
	fmt.Println("Init...")
	strFontFile := fontPath + fontFile
	fontBytes, err := ioutil.ReadFile(strFontFile)
	if err != nil {
		panic("Read Font file error: " + err.Error())
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil || font == nil {
		panic("ParseFont error: " + err.Error())
	}
	context = freetype.NewContext()
	context.SetDPI(120)
	context.SetFont(font)
	context.SetFontSize(fontSize)
}

func DrawToImg(strCode string) ([]byte, error) {
	//计算字符串的宽度，对于高度，还有此问题，懂的可以改改
	pt := freetype.Pt(10, 10+int(context.PointToFixed(fontSize)>>6))
	ptr, err := context.MeasureString(strCode, pt)
	if err != nil {
		return nil, err
	}
	//pt.Y += context.PointToFixed(fontSize * spacing)
	width := int(ptr.X >> 6)
	width += 10
	height := int(pt.Y+context.PointToFixed(fontSize*spacing))>>6 - int(context.PointToFixed(fontSize)>>6)
	fg, bg := image.Black, image.White
	rgba := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	//disturbBitmap(rgba)
	context.SetClip(rgba.Bounds())
	context.SetDst(rgba)
	context.SetSrc(fg)
	_, err = context.DrawString("1234", pt)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = png.Encode(buf, rgba)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//绘制干扰背景
func disturbBitmap(img *image.RGBA) {
	r := rand.New(rand.NewSource(int64(time.Now().Second())))
	for i := 0; i < img.Rect.Max.X; i++ {
		for j := 0; j < img.Rect.Max.Y; j++ {
			n := r.Intn(100)
			if n < 40 {
				c := color.NRGBA{uint8(r.Intn(150)), uint8(r.Intn(150)), uint8(r.Intn(150)), uint8(r.Intn(100))}
				img.Set(i, j, c)
			}
		}

	}

}
