package game

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/KEINOS/go-noise"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type WorldMap struct {
	world       [][]bool
	Width       int
	Height      int
	GraphicSize int // actual pixel size of one block
}

func (w *WorldMap) Setup() {
	// Initialize world Map
	if w.Height == 0 || w.Width == 0 || w.GraphicSize == 0 {
		log.Fatal("must specify width and height and graphic size when initialize worldmap")
	}
	w.world = make([][]bool, w.Height)
	for i := range w.world {
		w.world[i] = make([]bool, w.Width)
	}

	w.GenerateMap()
}

func (w *WorldMap) GenerateMap() {
	seed := time.Now().Second()
	randGen := rand.New(rand.NewSource(int64(seed)))
	n, err := noise.New(noise.Perlin, randGen.Int63())

	if err != nil {
		log.Fatal("Generate map error")
	}
	width := len(w.world[0])
	height := len(w.world)
	topTerrains := make([]float32, width)
	for i := range topTerrains {
		topTerrains[i] = (n.Eval32(float32(i)*float32(w.GraphicSize)/100)*0.5+0.5)*float32(height/2) + float32(height)/2
	}

	for y := range w.world {
		for x := range w.world[y] {
			if y > int(topTerrains[x]) {
				w.world[y][x] = true
			}
		}
	}
}

func (w *WorldMap) Update(elapsed time.Duration) {
}

func (w *WorldMap) Render(screen *ebiten.Image) {
	for y := range w.world {
		for x := range w.world[y] {
			if w.world[y][x] {
				vector.DrawFilledRect(screen, float32(x*w.GraphicSize), float32(y*w.GraphicSize), float32(w.GraphicSize), float32(w.GraphicSize), color.RGBA{0x27, 0x37, 0x4d, 0xff}, false)
			}
		}
	}
}

func (w *WorldMap) UpdatePhysic() {

}
