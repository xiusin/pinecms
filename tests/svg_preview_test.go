package tests

import (
	"fmt"
	"testing"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/rasterizer"
)

var fontLatin *canvas.FontFamily
var fontArabic *canvas.FontFamily
var fontDevanagari *canvas.FontFamily

const mmPrePx = 25.4 / 96.0 // 毫米每像素

func Test_full(t *testing.T) {
	fontLatin = canvas.NewFontFamily("站酷高端黑")
	if err := fontLatin.LoadLocalFont("站酷高端黑", canvas.FontRegular); err != nil {
		panic(err)
	}
	c := canvas.NewFromSize(canvas.Size{H: 1080 * mmPrePx, W: 1920 * mmPrePx}) // 接收一个宽高比参数
	ctx := canvas.NewContext(c)
	ctx.SetFillColor(canvas.Transparent)
	draw(ctx)

	c.WriteFile("preview.png", rasterizer.PNGWriter(canvas.DefaultResolution))
}

func drawText(c *canvas.Context, x, y float64, face *canvas.FontFace, rich *canvas.RichText) {
	metrics := face.Metrics()
	//width, height := 1280.0, 400.0
	fmt.Println(x, y, metrics.String())
	text := rich.ToText(200 * mmPrePx, 100*mmPrePx, canvas.Justify, canvas.Center, 0.0, 0.0)
	fmt.Println("text size", text.Bounds())
	c.DrawText(x, y, text)
}

func draw(c *canvas.Context) {
	// Draw a comprehensive text box   27个字  每个字体200 27 * 200 =
	pt := 200.0 * mmPrePx

	// 字体描边
	c.SetStrokeWidth(0.75)
	face := fontLatin.Face(pt, canvas.Black, canvas.FontRegular, canvas.FontNormal, canvas.FontSineUnderline, canvas.FontOverline, canvas.FontSawtoothUnderline)
	rt := canvas.NewRichText(face)
	rt.Add(face, "你好世界你好世界你好世界你好世界你好你")
	rt.Add(fontLatin.Face(pt, canvas.Red, canvas.FontLight, canvas.FontNormal), "欢迎来到德莱联盟")

	drawText(c, 1, 100*mmPrePx, face, rt)

}
