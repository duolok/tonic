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
	"maze":       arts.NewMaze(10),
	"blackhole":  arts.NewBlackHole(200, 400, 0.01),
	"random":     arts.NewRandomShape(150),
	"silksky":    arts.NewSilkSky(15, 5),
	"circles":    arts.NewColorCircle2(30),
	"contour":    arts.NewContourLine(5000),
    "circlenoise":arts.NewCircleNoise(2000, 60, 80),
	"noise":      arts.NewNoiseLine(1000),
	"circle":     arts.NewCircleLoop(100),
	"dot":        arts.NewDotLine(100, 20, 50, false),
	"perlin":     arts.NewPerlinPerls(10, 200, 40, 80),
	"dotswave":   arts.NewDotsWave(300),
	"colorcanva": arts.NewColorCanve(10),
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

	case "pixelhole":
		colors := []color.RGBA{
			{0xF9, 0xC8, 0x0E, 0xFF},
			{0xF8, 0x66, 0x24, 0xFF},
			{0xEA, 0x35, 0x46, 0xFF},
			{0x66, 0x2E, 0x9B, 0xFF},
			{0x43, 0xBC, 0xCD, 0xFF},
		}
		c := generativeart.NewCanva(800, 800)
		c.SetBackground(common.Black)
		c.FillBackground()
		c.SetColorSchema(colors)
		c.SetIterations(1200)

	case "blackhole":
		c.SetLineWidth(1.2)
        c.SetBackground(common.Black)
		c.SetColorSchema([]color.RGBA{
			{0x00, 0x00, 0x00, 0xFF},
			{0xFF, 0xFF, 0xFF, 0xFF},
		})

    case "circlenoise":
        c.SetBackground(common.White)
        c.SetAlpha(80)
        c.SetLineWidth(0.3)
        c.FillBackground()
        c.SetIterations(400)


	case "random":
		c.SetColorSchema([]color.RGBA{
			{0xCF, 0x2B, 0x34, 0xFF},
			{0xF0, 0x8F, 0x46, 0xFF},
			{0xF0, 0xC1, 0x29, 0xFF},
			{0x19, 0x6E, 0x94, 0xFF},
			{0x35, 0x3A, 0x57, 0xFF},
		})

	case "perlin":
		c.SetAlpha(120)
		c.SetLineWidth(0.3)
		c.FillBackground()
		c.SetIterations(200)

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
		colors := []color.RGBA{
			{0x58, 0x18, 0x45, 0xFF},
			{0x90, 0x0C, 0x3F, 0xFF},
			{0xC7, 0x00, 0x39, 0xFF},
			{0xFF, 0x57, 0x33, 0xFF},
			{0xFF, 0xC3, 0x0F, 0xFF},
		}
		c := generativeart.NewCanva(1600, 1600)
		c.SetBackground(color.RGBA{0x1a, 0x06, 0x33, 0xFF})
		c.FillBackground()
		c.SetColorSchema(colors)

	case "noise":
		c.SetLineWidth(0.3)
		c.SetColorSchema([]color.RGBA{
			{0x80, 0x80, 0x80, 0xFF},
			{0xFF, 0x00, 0x00, 0xFF},
		})
	case "circle":
		c.SetLineWidth(1)
		c.SetLineColor(common.Lavender)
		c.SetAlpha(300)
		c.SetIterations(1000)

	case "dot":
		c.SetBackground(color.RGBA{230, 230, 230, 255})
		c.SetLineWidth(10)
		c.SetIterations(15000)
		c.FillBackground()

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
	c.FillBackground()
	c.Draw(DRAWINGS[art])

	fileName := fmt.Sprintf("/tmp/%s_%f.png", art, rand.Float64())
	if err := c.ToPNG(fileName); err != nil {
		fmt.Printf("Error saving file %s: %v\n", fileName, err)
	}
	return fileName
}
