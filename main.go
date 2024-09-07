package main

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(200, 100, "rimer")
	defer rl.CloseWindow()
	rl.SetExitKey(rl.KeyNull)
	rl.SetTargetFPS(30)

	var startTime time.Time
	var isRunning bool = false
	var minutes int = 1
	var seconds int

	for !rl.WindowShouldClose() {

		if !isRunning {
			if (rl.IsKeyPressed(rl.KeyJ) || rl.IsKeyPressed(rl.KeyDown)) && minutes > 0 {
				minutes--
			}
			if rl.IsKeyPressed(rl.KeyK) || rl.IsKeyPressed(rl.KeyUp) {
				minutes++
			}
			if rl.IsKeyPressed(rl.KeyH) || rl.IsKeyDown(rl.KeyLeft) {
				if seconds > 0 {
					seconds--
				} else {
					seconds = 59
					if minutes > 0 {
						minutes--
					}
				}
			}
			if rl.IsKeyPressed(rl.KeyL) || rl.IsKeyDown(rl.KeyRight) {
				if seconds < 59 {
					seconds++
				} else {
					seconds = 0
					minutes++
				}
			}

			if rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeyEnter) {
				startTime = time.Now()
				isRunning = true
			}
		}

		if isRunning && rl.IsKeyPressed(rl.KeyR) {
			rl.ClearBackground(rl.Red)
			isRunning = false
		}

		duration := time.Duration(minutes)*time.Minute + time.Duration(seconds)*time.Second
		// timer := time.NewTimer(duration)
		remaining := duration - time.Since(startTime)
		if remaining <= 0 && isRunning {
			isRunning = false
			// timer.Stop()
		}

		screenWidth := int32(rl.GetScreenWidth())
		screenHeight := int32(rl.GetScreenHeight())
		fontSize := (screenWidth / 5)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		if isRunning {

			remainingSeconds := int(remaining.Seconds())
			displayMinutes := remainingSeconds / 60
			displaySeconds := remainingSeconds % 60

			var timeColor rl.Color
			if displayMinutes < 1 && displaySeconds < 15 {
				timeColor = rl.Red
			} else {
				timeColor = rl.Black
			}

			timerText := fmt.Sprintf("%02d:%02d", displayMinutes, displaySeconds)
			rl.DrawText(timerText, screenWidth/2-rl.MeasureText(timerText, fontSize)/2, screenHeight/2-rl.MeasureText(timerText, fontSize)/4, fontSize, timeColor)

		} else {
			timerText := fmt.Sprintf("Set Timer: %02d:%02d", minutes, seconds)
			helpText := "Use arrow keys to adjust"

			rl.DrawText(timerText, screenWidth/2-rl.MeasureText(timerText, 20)/2, screenHeight/2, 20, rl.Black)
			rl.DrawText(helpText, screenWidth/2-rl.MeasureText(helpText, 10)/2, screenHeight/4, 10, rl.Black)
		}
		rl.EndDrawing()
	}
}
