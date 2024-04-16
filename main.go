package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type rect struct {
	x, y            int32
	speed           uint8
	sword           bool
	swordCountdown  uint8
	shield          bool
	shieldCountdown uint8
	life            uint8
	collided        bool
}

func main() {
	// Initialization of rects
	player := rect{x: 0, y: 0, speed: 10, sword: false, swordCountdown: 0, shield: false, shieldCountdown: 0, life: 10, collided: false}
	enemy := rect{x: 910, y: 0, speed: 3, sword: false, swordCountdown: 0, shield: false, shieldCountdown: 0, life: 10, collided: false}
	isGameOver, hasPlayerWon := false, false

	// Initialization of raylib and the game loop
	rl.InitWindow(960, 540, "Sword")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		// Initialize drawing
		rl.BeginDrawing()

		rl.ClearBackground(rl.LightGray)

		if !isGameOver {
			rl.DrawRectangle(enemy.x, enemy.y, 50, 50, rl.Red)
			rl.DrawRectangle(player.x, player.y, 50, 50, rl.Blue)

			rl.DrawText(fmt.Sprintf("Player Life: %v		Enemy Life: %v", player.life, enemy.life), 10, 10, 20, rl.Black)

			// Run
			if rl.IsKeyDown(rl.KeyLeftShift) {
				player.speed = 15
			} else {
				player.speed = 10
			}

			// Player movement
			switch {
			case rl.IsKeyDown(rl.KeyW):
				player.y -= int32(player.speed)
			case rl.IsKeyDown(rl.KeyS):
				player.y += int32(player.speed)
			case rl.IsKeyDown(rl.KeyA):
				player.x -= int32(player.speed)
			case rl.IsKeyDown(rl.KeyD):
				player.x += int32(player.speed)
			case rl.IsMouseButtonPressed(rl.MouseButtonLeft): // Sword Activation
				if !player.sword {
					player.sword = true
					player.swordCountdown = 60
					player.shield = false
				} else {
					player.sword = false
					player.swordCountdown = 0
				}
			case rl.IsMouseButtonPressed(rl.MouseButtonRight): // Shield Activation
				if !player.shield {
					player.shield = true
					player.shieldCountdown = 120
					player.sword = false
				} else {
					player.shield = false
					player.shieldCountdown = 0
				}
			}

			// Enemy movement
			if areNear(enemy, player, 100) {
				enemy.speed = 7
				enemy.sword = true
			} else {
				enemy.speed = 3
				enemy.sword = false
			}

			if enemy.x > player.x {
				enemy.x -= int32(enemy.speed)
			} else if enemy.x < player.x {
				enemy.x += int32(enemy.speed)
			}
			if enemy.y > player.y {
				enemy.y -= int32(enemy.speed)
			} else if enemy.y < player.y {
				enemy.y += int32(enemy.speed)
			}

			if player.sword {
				player.swordCountdown -= 3
			}
			if player.swordCountdown <= 0 {
				player.sword = false
			}

			if player.shield {
				player.shieldCountdown -= 2
			}
			if player.shieldCountdown <= 0 {
				player.shield = false
			}

			// Sword and shield drawing
			if player.sword {
				rl.DrawRectangle(player.x+10, player.y+12, 25, 25, rl.DarkPurple)
			}

			if player.shield {
				rl.DrawRectangle(player.x-5, player.y-5, 60, 60, rl.DarkPurple)
			}

			// Check collisions
			if !player.shield && checkCollision(player, enemy) {
				if !player.collided {
					player.life--
					player.collided = true
				}
			} else {
				player.collided = false
			}
			if player.sword && checkCollision(player, enemy) {
				if !enemy.collided {
					enemy.life--
					enemy.collided = true
				}
			} else {
				enemy.collided = false
			}

			if player.life < 1 || player.life > 10 {
				isGameOver = true
			}
			if enemy.life < 1 || enemy.life > 10 {
				isGameOver = true
				hasPlayerWon = true
			}
		} else if hasPlayerWon {
			rl.DrawText("You Won!", 480, 270, 20, rl.Black)
		} else if !hasPlayerWon {
			rl.DrawText("Game Over!", 480, 270, 20, rl.Black)
		}

		rl.EndDrawing()
	}
}

func checkCollision(rect1, rect2 rect) bool {
	if rect1.x+50 < rect2.x {
		return false
	}

	if rect1.x > rect2.x+50 {
		return false
	}

	if rect1.y+50 < rect2.y {
		return false
	}

	if rect1.y > rect2.y+50 {
		return false
	}

	return true
}

func areNear(rect1, rect2 rect, distanceThreshold float64) bool {
	center1X := float64(rect1.x) + 25
	center1Y := float64(rect1.y) + 25
	center2X := float64(rect2.x) + 25
	center2Y := float64(rect2.y) + 25

	distance := math.Sqrt(math.Pow(center2X-center1X, 2) + math.Pow(center2Y-center1Y, 2))

	return distance <= distanceThreshold
}
