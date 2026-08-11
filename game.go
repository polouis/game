package main

import (
	"fmt"
	"os"

	"github.com/polouis/engine"
	"github.com/polouis/engine/types"
)

type HealthComponent struct {
	Current int
	Max     int
}

type GameWorld struct {
	HealthStore      *engine.ComponentArray[HealthComponent]
	InputIntentStore *engine.ComponentArray[InputIntentComponent]
}

func NewGameWorld() *GameWorld {
	return &GameWorld{
		HealthStore:      engine.NewComponentArray[HealthComponent](),
		InputIntentStore: engine.NewComponentArray[InputIntentComponent](),
	}
}

var playerEnt engine.EntityID
var squareEnt engine.EntityID

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func loadMeshFromFile(ctx *engine.Context, filePath string) {
	dat, err := os.ReadFile(filePath)
	checkErr(err)
	err = ctx.RM.LoadMesh(dat)
	checkErr(err)
}

func loadMeshes(ctx *engine.Context) {
	loadMeshFromFile(ctx, "./asset/mesh/1.json")
	loadMeshFromFile(ctx, "./asset/mesh/square.json")
}

func initialize(ctx *engine.Context, gw *GameWorld) {

	loadMeshes(ctx)

	atlas := engine.NewAtlas()
	err := atlas.Load(engine.AtlasLoadCrunch, engine.ImageLoadAseprite)
	checkErr(err)

	dat, err := os.ReadFile("./asset/entity/player.json")
	checkErr(err)
	err = ctx.RM.LoadEntity("player", dat)
	checkErr(err)
	playerEnt, err = ctx.RM.Spawn(ctx, "player")
	checkErr(err)
	gw.InputIntentStore.Upsert(playerEnt, InputIntentComponent{})

	dat, err = os.ReadFile("./asset/entity/square.json")
	checkErr(err)
	err = ctx.RM.LoadEntity("square", dat)
	checkErr(err)
	squareEnt, err = ctx.RM.Spawn(ctx, "square")
	checkErr(err)
}

func update(ctx *engine.Context, gw *GameWorld, deltaTime uint64) {
	UpdateInputSystem(ctx, gw, deltaTime)
	UpdateMovementSystem(ctx, gw, deltaTime)

	engine.UpdatePhysicsSystem(ctx, deltaTime)
	engine.UpdateRenderSystem(ctx, deltaTime)
}

func release(ctx *engine.Context) {
	engine.ReleaseRenderSystem(ctx)
}

func main() {

	fmt.Println("Game starting !")
	ctx := engine.New(types.SDL)
	gameWorld := NewGameWorld()

	err := engine.Run(
		ctx,
		func(ctx *engine.Context) { initialize(ctx, gameWorld) },
		func(ctx *engine.Context, deltaTime uint64) { update(ctx, gameWorld, deltaTime) },
		release)

	if err != nil {
		fmt.Printf("Got error during run : %v", err)
	}
}
