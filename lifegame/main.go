package lifegame

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	// WIDTH  = 1280 / SIZE
	SIZE   = 5
	WIDTH  = 720 / SIZE
	HEIGHT = 720 / SIZE
)

var (
	MATRIX = [][]int{}
	BUFFER = [][]int{}
)

func Reset() {
	// Init Cells
	MATRIX = [][]int{}
	BUFFER = [][]int{}
	line := []int{}
	bufferLine := []int{}
	seed := time.Now().Unix()
	rand.Seed(seed)
	for row := 0; row < HEIGHT; row++ {
		for col := 0; col < WIDTH; col++ {
			newCell := rand.Intn(7) % 2
			bufferCell := 0
			line = append(line, newCell)
			bufferLine = append(bufferLine, bufferCell)
		}
		MATRIX = append(MATRIX, line)
		BUFFER = append(BUFFER, bufferLine)
		line = []int{}
		bufferLine = []int{}
	}
}
func init() {
	Reset()
}

func Run() {
	runtime.LockOSThread()
	// Init sdl.EVERYTHING
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()
	// Init a window in 1280*720
	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(WIDTH*SIZE), int32(HEIGHT*5), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
	}
	defer renderer.Destroy()

	// First render
	renderer.Clear()
	for row := 0; row < HEIGHT; row++ {
		for col := 0; col < WIDTH; col++ {
			if MATRIX[row][col] == 1 {
				renderer.SetDrawColor(255, 255, 255, 255)
			} else {
				renderer.SetDrawColor(0, 0, 0, 255)
			}
			rect := sdl.Rect{X: int32(col * SIZE), Y: int32(row * SIZE), W: 5, H: 5}
			renderer.FillRect(&rect)
		}
	}
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
				if event.(*sdl.KeyboardEvent).Keysym.Scancode == sdl.SCANCODE_F5 && event.(*sdl.KeyboardEvent).State == sdl.RELEASED {
					fmt.Println("Restart!")
					Reset()
					renderer.Clear()
					for row := 0; row < HEIGHT; row++ {
						for col := 0; col < WIDTH; col++ {
							if MATRIX[row][col] == 1 {
								renderer.SetDrawColor(255, 255, 255, 255)
							} else {
								renderer.SetDrawColor(0, 0, 0, 255)
							}
							rect := sdl.Rect{X: int32(col * SIZE), Y: int32(row * SIZE), W: 5, H: 5}
							renderer.FillRect(&rect)
						}
					}
					renderer.Present()
				}
			}
		}
		renewWindow(renderer)
	}
}
func renewWindow(renderer *sdl.Renderer) {
	renderer.Clear()
	for row := 0; row < HEIGHT; row++ {
		for col := 0; col < WIDTH; col++ {
			nearby := [][]int{
				[]int{row - 1, col - 1},
				[]int{row - 1, col},
				[]int{row - 1, col + 1},
				[]int{row, col - 1},
				[]int{row, col + 1},
				[]int{row + 1, col - 1},
				[]int{row + 1, col},
				[]int{row + 1, col + 1},
			}
			aliveCount := 0
			for index, coordinate := range nearby {
				if coordinate[0] == -1 {
					nearby[index][0] += HEIGHT
				}
				if coordinate[1] == -1 {
					nearby[index][1] += WIDTH
				}
				if coordinate[0] == HEIGHT {
					nearby[index][0] -= HEIGHT
				}
				if coordinate[1] == WIDTH {
					nearby[index][1] -= WIDTH
				}
				if MATRIX[coordinate[0]][coordinate[1]] == 1 {
					aliveCount++
				}
			}
			if MATRIX[row][col] == 1 {
				renderer.SetDrawColor(255, 255, 255, 255)
				if aliveCount < 2 || aliveCount > 3 {
					// DEAD
					BUFFER[row][col] = 0
				}
			} else {
				renderer.SetDrawColor(0, 0, 0, 255)
				if aliveCount > 3 {
					// ALIVE
					BUFFER[row][col] = 1
				}
			}
			rect := sdl.Rect{X: int32(col * SIZE), Y: int32(row * SIZE), W: 5, H: 5}
			renderer.FillRect(&rect)
		}
	}
	copy(MATRIX, BUFFER)
	renderer.Present()
}
