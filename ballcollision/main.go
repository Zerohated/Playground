package ballcollision

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	WIDTH  = 720
	HEIGHT = 720
)

// Ball describe the state of a ball object
type Ball struct {
	X       float64
	Y       float64
	Radius  float64
	X_Speed float64
	Y_Speed float64
}

func (ball *Ball) random() {
	seed := time.Now().Unix()
	rand.Seed(seed)
	ball.X, ball.Y = float64(rand.Intn(700)), float64(rand.Intn(700))
	ball.Radius = 10
	ball.X_Speed, ball.Y_Speed = float64(2*rand.Intn(10)), float64(2*rand.Intn(10))

}
func (ball Ball) render(renderer *sdl.Renderer) error {
	left := ball.X - ball.Radius
	right := ball.X + ball.Radius
	top := ball.Y - ball.Radius
	bottom := ball.Y + ball.Radius
	renderer.SetDrawColor(255, 255, 255, 255)
	err := renderer.Clear()
	if err != nil {
		return err
	}
	renderer.SetDrawColor(0, 0, 0, 255)
	for col := left; col <= right; col++ {
		for row := top; row <= bottom; row++ {
			if (col-ball.X)*(col-ball.X)+(row-ball.Y)*(row-ball.Y) <= ball.Radius*ball.Radius {
				renderer.DrawPoint(int32(col), int32(row))
			}
		}
	}
	renderer.Present()
	return nil
}
func (ball *Ball) nextFrame() error {
	nextX, nextY := ball.X+ball.X_Speed, ball.Y+ball.Y_Speed
	if nextX+ball.Radius >= WIDTH {
		nextX = WIDTH - (ball.Radius+nextX-WIDTH+ball.Radius)*0.9
		ball.X_Speed = -(ball.X_Speed * 0.9)
	}
	if nextX-ball.Radius <= 0 {
		nextX = 0 + (ball.Radius-nextX+ball.Radius)*0.9
		ball.X_Speed = -(ball.X_Speed * 0.9)
	}
	if nextY+ball.Radius >= HEIGHT {
		nextY = HEIGHT - (ball.Radius+nextY-HEIGHT+ball.Radius)*0.9
		ball.Y_Speed = -(ball.Y_Speed * 0.9)
	}
	if nextY-ball.Radius <= 0 {
		nextY = 0 + (ball.Radius-nextY)*0.9
		ball.Y_Speed = -(ball.Y_Speed * 0.9)
	}
	ball.Y_Speed++
	ball.X, ball.Y = nextX, nextY
	return nil
}

func Run() {
	runtime.LockOSThread()
	// Init sdl.EVERYTHING
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()
	// Init a window in WIDTH*HEIGHT
	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(WIDTH), int32(HEIGHT), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	// Create Renderer in this window
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
	}
	defer renderer.Destroy()

	thisBall := Ball{X: rand.Float64() * 720, Y: rand.Float64() * 720, Radius: 10, X_Speed: rand.Float64() * 2, Y_Speed: rand.Float64() * 2}
	// First render
	renderer.Clear()
	thisBall.render(renderer)
	renderer.Present()

	// go func() {
	// 	for {
	// 		// time.Sleep(time.Second * 1)
	// 	}
	// }()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			case *sdl.KeyboardEvent:
				// Press "F5" to restart the scence
				if event.(*sdl.KeyboardEvent).Keysym.Scancode == sdl.SCANCODE_F5 && event.(*sdl.KeyboardEvent).State == sdl.RELEASED {
					fmt.Println("Restart!")
					// Reset()
					thisBall.random()
				}
			}
		}
		// render each frame
		sdl.Delay(10)
		thisBall.nextFrame()
		thisBall.render(renderer)
	}
}
