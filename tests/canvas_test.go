package tests

import (
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/softwarebackend"
	_ "image/jpeg"
	"image/png"
	"os"
	"testing"
)

func Test_Canvas(t *testing.T) {
	backend := softwarebackend.New(1920, 1080)
	cv := canvas.New(backend)

	//w, h := float64(cv.Width()), float64(cv.Height())
	//cv.SetFillStyle("#000")
	//cv.FillRect(0, 0, w, h)
	cv.DrawImage("cat.jpeg", 0, 0, 1920, 1080)

	cv.SetFont("/Users/xiusin/projects/src/github.com/xiusin/pinecms/tests/Righteous-Regular.ttf", 48)

	cv.SetFillStyle("#F00")

	cv.SetTextAlign(canvas.Center)

	cv.FillText("hello world", 50, 50)

	//w, h := float64(cv.Width()), float64(cv.Height())
	//cv.SetFillStyle("#000")
	//cv.FillRect(0, 0, w, h)
	//
	//for r := 0.0; r < math.Pi*2; r += math.Pi * 0.1 {
	//	cv.SetFillStyle(int(r*10), int(r*20), int(r*40))
	//	cv.BeginPath()
	//	cv.MoveTo(w*0.5, h*0.5)
	//	cv.Arc(w*0.5, h*0.5, math.Min(w, h)*0.4, r, r+0.1*math.Pi, false)
	//	cv.ClosePath()
	//	cv.Fill()
	//}
	//
	//cv.SetStrokeStyle("#FFF")
	//cv.SetLineWidth(10)
	//cv.BeginPath()
	//cv.Arc(w*0.5, h*0.5, math.Min(w, h)*0.4, 0, math.Pi*2, false)
	cv.Stroke()

	f, err := os.OpenFile("result.png", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	err = png.Encode(f, backend.Image)
	if err != nil {
		panic(err)
	}
}
