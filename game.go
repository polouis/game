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
	HealthStore *engine.ComponentArray[HealthComponent]
}

func NewGameWorld() *GameWorld {
	return &GameWorld{
		HealthStore: engine.NewComponentArray[HealthComponent](),
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

func initialize(ctx *engine.Context) {

	loadMeshes(ctx)

	atlas := engine.NewAtlas()
	err := atlas.Load(engine.AtlasLoadCrunch, engine.ImageLoadAseprite)
	checkErr(err)

	dat, err := os.ReadFile("./asset/entity/player.json")
	checkErr(err)
	err = ctx.RM.LoadEntity("player", dat)
	checkErr(err)
	playerEnt, err = ctx.RM.Spawn(ctx, "player")

	dat, err = os.ReadFile("./asset/entity/square.json")
	checkErr(err)
	err = ctx.RM.LoadEntity("square", dat)
	checkErr(err)
	squareEnt, err = ctx.RM.Spawn(ctx, "square")
	checkErr(err)
}

func update(ctx *engine.Context, deltaTime uint64) {
	cmd := engine.HandleInput(ctx, deltaTime)
	if cmd != nil {
		cmd.Execute(ctx, playerEnt)
	}

	engine.UpdatePhysicsSystem(ctx, deltaTime)
	engine.UpdateRenderSystem(ctx, deltaTime)
}

func release(ctx *engine.Context) {
	engine.ReleaseRenderSystem(ctx)
}

func main() {

	fmt.Println("Game starting !")
	ctx := engine.New(types.SDL)

	err := engine.Run(ctx, initialize, update, release)

	if err != nil {
		fmt.Printf("Got error during run : %v", err)
	}
}
