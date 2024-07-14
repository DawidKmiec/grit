package game

import (
	global "grit/global"
)

func cycleGame(grid [global.GRID_WIDTH][global.GRID_HEIGHT]byte) [global.GRID_WIDTH][global.GRID_HEIGHT]byte {
	// OMG LETS GOOO BUG FIX!!!!
	// I was referencing a pointer dumbass
	new_grid := grid

	for x := 0; x < global.GRID_WIDTH; x++ {
		for y := 0; y < global.GRID_HEIGHT; y++ {

			switch pixel_type := grid[x][y]; pixel_type {
			case 0:
				continue
			case 's':
				sandLogic(&grid, &new_grid, gridXY{x, y})
			case 'b':
				// work on functionality later
			case 'w':
				// work on functionality later
			default:
				panic("invalid item found in grid")

			}

		}
	}

	return new_grid
}
