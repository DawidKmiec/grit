package game

import (
	global "grit/global"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func GameLoop() {
	defer rl.CloseWindow()
	initWindow()
	var grid [global.GRID_WIDTH][global.GRID_HEIGHT]byte

	for !rl.WindowShouldClose() {
		//update
		mousePlace(&grid)
		grid = cycleGame(grid)

		// render
		render(&grid)
	}
}

type gridXY struct {
	x int
	y int
}

func initWindow() {
	rl.InitWindow(global.WINDOW_WIDTH, global.WINDOW_HEIGHT, global.WINDOW_NAME)
	rl.SetTargetFPS(global.TARGET_RENDER_FPS)
}
