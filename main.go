package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type rect struct {
	x, y   int32
	sword  bool
	shield bool
	life   uint8
}

func main() {
	player := rect{x: 0, y: 0, sword: false, shield: false, life: 10}
	enemy := rect{x: 0, y: 0, sword: false, shield: false, life: 10} // Adjusted enemy's starting position

	rl.InitWindow(960, 540, "Sword")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.LightGray)

		rl.DrawRectangle(enemy.x, enemy.y, 50, 50, rl.Red)
		rl.DrawRectangle(player.x, player.y, 50, 50, rl.Blue)

		rl.DrawText(fmt.Sprintf("Player Life: %v		Enemy Life: %v\n", player.life, enemy.life), 10, 10, 20, rl.Black)

		// Player movement
		switch {
		case rl.IsKeyDown(rl.KeyW):
			player.y -= 10
		case rl.IsKeyDown(rl.KeyS):
			player.y += 10
		case rl.IsKeyDown(rl.KeyA):
			player.x -= 10
		case rl.IsKeyDown(rl.KeyD):
			player.x += 10
		case rl.IsMouseButtonPressed(rl.MouseButtonLeft):
			if !player.sword {
				player.sword = true
				player.shield = false
			} else {
				player.sword = false
			}
		case rl.IsMouseButtonPressed(rl.MouseButtonRight):
			if !player.shield {
				player.shield = true
				player.sword = false
			} else {
				player.shield = false
			}
		}

		// Enemy movement
		if enemy.x > player.x {
			enemy.x -= 5
		} else if enemy.x < player.x {
			enemy.x += 5
		}
		if enemy.y > player.y {
			enemy.y -= 5
		} else if enemy.y < player.y {
			enemy.y += 5
		}

		if player.sword && rl.GetMouseX() > player.x-100 && rl.GetMouseX() < player.x+120 && rl.GetMouseY() > player.y-100 && rl.GetMouseY() < player.y+120 {
			rl.DrawRectangle(rl.GetMouseX(), rl.GetMouseY(), 25, 25, rl.DarkPurple)
		}

		if player.shield {
			rl.DrawRectangle(player.x-5, player.y-5, 60, 60, rl.DarkPurple)
		}

		if !player.shield && checkCollision(player, enemy) {
			player.life--
		} else if player.sword && checkCollision(player, enemy) {
			enemy.life--
		}

		rl.EndDrawing()
	}
}

func checkCollision(rect1, rect2 rect) bool {
	// Check if rect1 is to the left of rect2
	if rect1.x+50 < rect2.x {
		return false
	}

	// Check if rect1 is to the right of rect2
	if rect1.x > rect2.x+50 {
		return false
	}

	// Check if rect1 is above rect2
	if rect1.y+50 < rect2.y {
		return false
	}

	// Check if rect1 is below rect2
	if rect1.y > rect2.y+50 {
		return false
	}

	// If none of the above conditions are true, rectangles overlap and there is a collision
	return true
}
