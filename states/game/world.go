package game

import (
	"image/color"
	"log"
	"math"
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
		entity.SetStable(false)
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

		veloAngle := potentialVelo.ToArcTan2f() // Current velo angle
		radius := entity.GetRadius()
		responseVelo := common.Vectorf{}
		isCollision := false

		for iteAngle := veloAngle - math.Pi/2; iteAngle <= veloAngle+math.Pi/2; iteAngle += math.Pi / 8 {
			checkPos := potentialPos.Add(
				common.Vectorf{
					X: float64(radius) * math.Cos(iteAngle),
					Y: float64(radius) * math.Sin(iteAngle),
				},
			)
			posMapX, posMapY := int(checkPos.X)/w.graphicSize, int(checkPos.Y)/w.graphicSize
			if posMapX >= w.width {
				posMapX = w.width - 1
			}
			if posMapX < 0 {
				posMapX = 0
			}
			if posMapY >= w.height {
				posMapY = w.height - 1
			}
			if posMapY < 0 {
				posMapY = 0
			}
			if w.world[posMapY][posMapX] {
				responseVelo = responseVelo.Add(potentialPos.Minus(checkPos))
				isCollision = true
			}
		}

		if isCollision {
			entity.SetStable(false)
			respMag := math.Sqrt(math.Pow(responseVelo.X, 2) + math.Pow(responseVelo.Y, 2))
			normalizeResp := common.Vectorf{
				X: responseVelo.X / respMag,
				Y: responseVelo.Y / respMag,
			}
			dot := potentialVelo.Dot(normalizeResp) // n * d
			entity.SetVelo(
				(potentialVelo.Minus(normalizeResp.Multi(2 * dot))).Multi(entity.GetFriction()),
			)
		} else {
			entity.SetVelo(potentialVelo)
			entity.SetPosition(potentialPos)
		}
		finalVelo := entity.GetVelo()
		veloMag := math.Sqrt(math.Pow(finalVelo.X, 2) + math.Pow(finalVelo.Y, 2))

		if veloMag < 0.1 {
			entity.SetStable(true)
		}
	}
}
