package generator

import (
    "fmt"
    "github.com/jdxyw/generativeart"
    "github.com/jdxyw/generativeart/arts"
    "github.com/jdxyw/generativeart/common"
    "image/color"
    "math/rand"
    "time"
)

var DRAWINGS = map[string]generativeart.Engine{
    "maze":      arts.NewMaze(10),
    "julia":     arts.NewJulia(func(z complex128) complex128 { return z*z + complex(-0.1, 0.651) }, 40, 1.5, 1.5),
    "randcicle": arts.NewRandCicle(30, 80, 0.2, 2, 10, 30, true),
    "blackhole": arts.NewBlackHole(200, 400, 0.01),
    "janus":     arts.NewJanus(5, 10),
    "random":    arts.NewRandomShape(150),
    "silksky":   arts.NewSilkSky(15, 5),
    "circles":   arts.NewColorCircle2(30),
    "contour":   arts.NewContourLine(5000),
    "noise":     arts.NewNoiseLine(1000),
    "circle":    arts.NewCircleLoop(100),
    "dot":       arts.NewDotLine(100, 20, 50, false),
    "colorcanva":arts.NewColorCanve(10),
}

func DrawOne(art string) string {
    rand.New(rand.NewSource(time.Now().Unix()))
    c := generativeart.NewCanva(1920, 1080)
    c.SetColorSchema([]color.RGBA{
        {0x58, 0x18, 0x45, 0xFF},
        {0x90, 0x0C, 0x3F, 0xFF},
        {0xC7, 0x00, 0x39, 0xFF},
        {0xFF, 0x57, 0x33, 0xFF},
        {0xFF, 0xC3, 0x0F, 0xFF},
    })

    c.SetBackground(color.RGBA{R: 30, G: 30, B: 30, A: 225})
    c.SetIterations(250)
    c.FillBackground()
    c.SetLineWidth(1.0)
    c.SetLineColor(common.Tomato)
    c.Draw(DRAWINGS[art])

    fileName := fmt.Sprintf("/tmp/%s_%f.png", art, rand.Float64())
    if err := c.ToPNG(fileName); err != nil {
        fmt.Printf("Error saving file %s: %v\n", fileName, err)
    }
    return fileName
}

