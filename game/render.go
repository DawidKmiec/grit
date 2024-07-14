package game

import (
	global "grit/global"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func render(grid *[global.GRID_WIDTH][global.GRID_HEIGHT]byte) {
	rl.BeginDrawing()
	defer rl.EndDrawing()
	rl.ClearBackground(rl.Black)

	renderGrid(grid)
}

// 0 0 is top right
func renderGrid(grid *[global.GRID_WIDTH][global.GRID_HEIGHT]byte) {
	for x := 0; x < global.GRID_WIDTH; x++ {
		for y := 0; y < global.GRID_HEIGHT; y++ {

			switch pixel_type := grid[x][y]; pixel_type {
			case 0:
				continue
			case 's':
				drawPixel(gridXY{x, y}, rl.Yellow)
			case 'b':
				drawPixel(gridXY{x, y}, rl.Gray)
			case 'w':
				drawPixel(gridXY{x, y}, rl.Blue)
			default:
				panic("item in grid is invalid")
			}
		}
	}
}

func drawPixel(grid_location gridXY, colour rl.Color) {
	rl.DrawRectangle(int32(grid_location.x*global.PIXEL_SIZE), int32(grid_location.y*global.PIXEL_SIZE), global.PIXEL_SIZE, global.PIXEL_SIZE, colour)
}
