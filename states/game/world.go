package game

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/KEINOS/go-noise"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/tonible14012002/go_game/engine/common"
)

type WorldMap struct {
	world        [][]bool
	width        int
	height       int
	graphicSize  int // actual pixel size of one block
	gravityAccel common.Vectorf
}

func (w *WorldMap) Setup(width, height, graphicSize int, gravityAccel common.Vectorf) *WorldMap {
	// Initialize world Map
	if width == 0 || height == 0 || graphicSize == 0 {
		log.Fatal("width, height, graphic size must larger than zero")
	}
	w.gravityAccel = gravityAccel
	w.width = width
	w.height = height
	w.graphicSize = graphicSize
	w.world = make([][]bool, w.height)
	for i := range w.world {
		w.world[i] = make([]bool, w.width)
	}

	w.GenerateMap()
	return w
}

func (w *WorldMap) SetGravityAccel(gravityAccel common.Vectorf) {
	w.gravityAccel = gravityAccel
}

func (w WorldMap) GetGravityAccel() common.Vectorf {
	return w.gravityAccel
}

func (w *WorldMap) ResetWorld() {
	// TODO: Reset world map properties
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
		topTerrains[i] = (n.Eval32(float32(i)*float32(w.graphicSize)/100)*0.5+0.5)*float32(height/2) + float32(height)/2
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
				vector.DrawFilledRect(
					screen,
					float32(x*w.graphicSize),
					float32(y*w.graphicSize),
					float32(w.graphicSize),
					float32(w.graphicSize),
					color.RGBA{0x27, 0x37, 0x4d, 0xff},
					false,
				)
			}
		}
	}
}

func (w *WorldMap) UpdatePhysic(elapsed time.Duration, entities Entities) {
	t := elapsed.Seconds()
	for _, entity := range entities {
		pos := entity.GetPosition()
		velo := entity.GetVelo()
		accel := entity.GetAccel()
		// NOTE: Current Entity dont have accel
		potentialAccel := common.Vectorf{
			X: accel.X + w.gravityAccel.X,
			Y: accel.Y + w.gravityAccel.Y,
		}
		potentialVelo := common.Vectorf{
			X: velo.X + potentialAccel.X*t,
			Y: velo.Y + potentialAccel.Y*t,
		}
		potentialPos := common.Vectorf{
			X: pos.X + potentialVelo.X*t,
			Y: pos.Y + potentialVelo.Y*t,
		}
		entity.SetVelo(potentialVelo)
		entity.SetPosition(potentialPos)
		fmt.Println(entity.GetVelo())
	}
}
