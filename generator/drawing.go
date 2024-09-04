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
    "janus":     arts.NewJanus(10, 0.2),
    "random":    arts.NewRandomShape(150),
    "silksky":   arts.NewSilkSky(15, 5),
    "circles":   arts.NewColorCircle2(30),
    "contour":   arts.NewContourLine(5000),
    "noise":     arts.NewNoiseLine(1000),
    "circle":    arts.NewCircleLoop(100),
    "dot":       arts.NewDotLine(100, 20, 50, false),
    "dotswave":  arts.NewDotsWave(300),
    "colorcanva":arts.NewColorCanve(10),
}

func optimizeDrawing(art string, c *generativeart.Canva) {
    switch art {
    case "maze":
        c.SetLineWidth(2.0)
        c.SetColorSchema([]color.RGBA{
            {0x00, 0x72, 0xFF, 0xFF}, 
            {0x00, 0xD8, 0xA0, 0xFF}, 
        })

    case "dotswave":
        colors := []color.RGBA{
            {0xFF, 0xBE, 0x0B, 0xFF},
            {0xFB, 0x56, 0x07, 0xFF},
            {0xFF, 0x00, 0x6E, 0xFF},
            {0x83, 0x38, 0xEC, 0xFF},
            {0x3A, 0x86, 0xFF, 0xFF},
        }
        c := generativeart.NewCanva(500, 500)

        c.SetColorSchema(colors)

    case "julia":
        c.SetLineWidth(1.5)
        c.SetColorSchema([]color.RGBA{
            {0xD4, 0xAF, 0x37, 0xFF}, 
            {0x80, 0x00, 0x80, 0xFF}, 
        })
    case "randcicle":
        c.SetLineWidth(1.0)
        c.SetColorSchema([]color.RGBA{
            {0xFF, 0xA5, 0x00, 0xFF}, 
            {0xFF, 0x45, 0x00, 0xFF}, 
        })
    case "blackhole":
        c.SetLineWidth(1.2)
        c.SetColorSchema([]color.RGBA{
            {0x00, 0x00, 0x00, 0xFF}, 
            {0xFF, 0xFF, 0xFF, 0xFF}, 
        })
    case "janus":
        c.SetLineWidth(2.0)
        c.SetColorSchema([]color.RGBA{
            {0xFF, 0xC0, 0xCB, 0xFF}, 
            {0xFF, 0x69, 0xB4, 0xFF}, 
        })
    case "random":
             c.SetColorSchema([]color.RGBA{
              {0xCF, 0x2B, 0x34, 0xFF},
              {0xF0, 0x8F, 0x46, 0xFF},
              {0xF0, 0xC1, 0x29, 0xFF},
              {0x19, 0x6E, 0x94, 0xFF},
              {0x35, 0x3A, 0x57, 0xFF},
             })




    case "silksky":
        c.SetLineWidth(0.5)
        c.SetColorSchema([]color.RGBA{
            {0x87, 0xCE, 0xFA, 0xFF}, 
            {0x1E, 0x90, 0xFF, 0xFF}, 
        })
    case "circles":
        colors := []color.RGBA{
          {0x11, 0x60, 0xC6, 0xFF},
          {0xFD, 0xD9, 0x00, 0xFF},
          {0xF5, 0xB4, 0xF8, 0xFF},
          {0xEF, 0x13, 0x55, 0xFF},
          {0xF4, 0x9F, 0x0A, 0xFF},
         }
         c := generativeart.NewCanva(800, 800)
         c.SetColorSchema(colors)
    case "contour":
        c.SetLineWidth(0.4)
        c.SetColorSchema([]color.RGBA{
            {0x00, 0x80, 0x80, 0xFF}, 
            {0x00, 0xFF, 0xFF, 0xFF}, 
        })
    case "noise":
        c.SetLineWidth(0.3)
        c.SetColorSchema([]color.RGBA{
            {0x80, 0x80, 0x80, 0xFF}, 
            {0xFF, 0x00, 0x00, 0xFF}, 
        })
    case "circle":
        c.SetLineWidth(1.0)
        c.SetColorSchema([]color.RGBA{
            {0x00, 0x80, 0x00, 0xFF}, 
            {0xFF, 0xFF, 0x00, 0xFF}, 
        })
    case "dot":
        c.SetLineWidth(1.5)
        c.SetColorSchema([]color.RGBA{
            {0xFF, 0xA5, 0x00, 0xFF}, 
            {0xFF, 0x00, 0xFF, 0xFF}, 
        })
    case "colorcanva":
        c.SetLineWidth(2.0)
        c.SetColorSchema([]color.RGBA{
            {0x6A, 0x5A, 0xCD, 0xFF}, 
            {0x8B, 0x00, 0x8B, 0xFF}, 
        })
    default:
        c.SetLineWidth(1.0)
        c.SetColorSchema([]color.RGBA{
            {0xFF, 0xFF, 0xFF, 0xFF}, 
        })
    }
}

func DrawOne(art string) string {
    rand.New(rand.NewSource(time.Now().Unix()))
    c := generativeart.NewCanva(1920, 1080)
    optimizeDrawing(art, c) 
    c.SetBackground(common.Black)
    c.FillBackground()
    c.Draw(DRAWINGS[art])

    fileName := fmt.Sprintf("/tmp/%s_%f.png", art, rand.Float64())
    if err := c.ToPNG(fileName); err != nil {
        fmt.Printf("Error saving file %s: %v\n", fileName, err)
    }
    return fileName
}
