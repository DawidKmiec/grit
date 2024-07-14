package game

import global "grit/global"

func checkValidity(item_location gridXY) bool {
	if item_location.x > global.GRID_WIDTH-1 || item_location.x < 0 {
		return false
	}
	if item_location.y > global.GRID_HEIGHT-1 || item_location.y < 0 {
		return false
	}
	return true
}

func checkItemAtDisplacement(grid *[global.GRID_WIDTH][global.GRID_HEIGHT]byte, item_location gridXY, displacement gridXY) byte {
	if !checkValidity(gridXY{item_location.x + displacement.x, item_location.y + displacement.y}) {
		return 'e'
	}
	return grid[item_location.x+displacement.x][item_location.y+displacement.y]
}

func collisionLogic(new_grid *[global.GRID_WIDTH][global.GRID_HEIGHT]byte, original_item_location gridXY, target_item_location gridXY, item_type byte, item_priority int) {
	// higher priority means it takes more precedent
	// for example sand = 2, water = 3
	// sand will take precedent
	// if it is in target area, it will not move
	// if it is the moving item, it will swap

	var target_item_priority int
	switch new_grid[target_item_location.x][target_item_location.y] {
	case 0:
		target_item_priority = global.EMPTY_PRIORITY
	case 's':
		target_item_priority = global.SAND_PRIORITY
	case 'b':
		target_item_priority = global.BRICK_PRIORITY
	}

	//  target_item_priority
	if item_priority < target_item_priority {
		// swap
		new_grid[original_item_location.x][original_item_location.y] = new_grid[target_item_location.x][target_item_location.y]
		new_grid[target_item_location.x][target_item_location.y] = item_type
	} else {
		// stay the same
		new_grid[original_item_location.x][original_item_location.y] = item_type
	}

}

func sandLogic(grid *[global.GRID_WIDTH][global.GRID_HEIGHT]byte, new_grid *[global.GRID_WIDTH][global.GRID_HEIGHT]byte, item_location gridXY) {
	// grid is static, new_grid will be changed
	// we perform the check for places the sand can move using grid, and set the target position to target_item_location
	// if it is an object we can sink through it will still be a valid target item location
	// we check with newgrid if this space is available
	// if it is, we swap the two values the space is available with
	target_item_location := item_location
	var value byte

	if value = checkItemAtDisplacement(grid, item_location, gridXY{0, 1}); value == 0 || value == 'w' {
		target_item_location = gridXY{item_location.x, item_location.y + 1}
	} else if value = checkItemAtDisplacement(grid, item_location, gridXY{-1, 1}); value == 0 || value == 'w' {
		target_item_location = gridXY{item_location.x - 1, item_location.y + 1}
	} else if value = checkItemAtDisplacement(grid, item_location, gridXY{1, 1}); value == 0 || value == 'w' {
		target_item_location = gridXY{item_location.x + 1, item_location.y + 1}
	}

	collisionLogic(new_grid, item_location, target_item_location, 's', global.SAND_PRIORITY)
}
