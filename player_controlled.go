package main

import (
	"github.com/polouis/engine"
	"github.com/polouis/engine/types"
)

type InputIntentComponent struct {
	Move engine.Vector3
}

func UpdateInputSystem(ctx *engine.Context, gw *GameWorld, deltaTime uint64) {
	var x, y float32
	if engine.GetKeyState(ctx, types.Up) || engine.GetButtonState(ctx, types.ButtonUp) {
		y += 1
	}
	if engine.GetKeyState(ctx, types.Down) || engine.GetButtonState(ctx, types.ButtonDown) {
		y -= 1
	}
	if engine.GetKeyState(ctx, types.Left) || engine.GetButtonState(ctx, types.ButtonLeft) {
		x -= 1
	}
	if engine.GetKeyState(ctx, types.Right) || engine.GetButtonState(ctx, types.ButtonRight) {
		x += 1
	}

	if x == 0 && y == 0 {
		return
	}

	for _, intent := range gw.InputIntentStore.All() {
		intent.Move = engine.Vector3{X: x, Y: y}
	}
}

func UpdateMovementSystem(ctx *engine.Context, gw *GameWorld, dt uint64) {
	for e, intent := range gw.InputIntentStore.All() {
		t, _ := ctx.W.TransformStore.Get(e)
		t.Position.X += intent.Move.X * 0.5 / 1e9 * float32(dt)
		t.Position.Y += intent.Move.Y * 0.5 / 1e9 * float32(dt)
	}
}
