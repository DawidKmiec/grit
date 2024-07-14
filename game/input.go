package game

import (
	global "grit/global"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func placePixel(grid *[global.GRID_WIDTH][global.GRID_HEIGHT]byte, item_location gridXY, selectedBlock byte) {
	if !checkValidity(item_location) {
		return
	}
	grid[item_location.x][item_location.y] = selectedBlock
}

func findGridSpace(location rl.Vector2) gridXY {
	grid_x := int(location.X) / global.PIXEL_SIZE
	grid_y := int(location.Y) / global.PIXEL_SIZE

	return gridXY{grid_x, grid_y}

}

func mousePlace(grid *[global.GRID_WIDTH][global.GRID_HEIGHT]byte) {
	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		mousePosition := rl.GetMousePosition()
		grid_location := findGridSpace(mousePosition)
		placePixel(grid, grid_location, 's')
	}
}
