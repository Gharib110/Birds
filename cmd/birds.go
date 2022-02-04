package main

import (
	"math/rand"
	"time"
)

type Bird struct {
	Position Vector2D
	Velocity Vector2D
	id       int
}

//moveOne for moving a *Bird to the next position
func (b *Bird) moveOne() {
	b.Position = *b.Position.AddV(&b.Velocity)
	next := b.Position.AddV(&b.Velocity)

	if next.X >= screenWidth || next.X < 0 {
		b.Velocity = Vector2D{
			X: -b.Velocity.X,
			Y: b.Velocity.Y,
		}
	}

	if next.Y >= screenHeight || next.Y < 0 {
		b.Velocity = Vector2D{
			X: b.Velocity.X,
			Y: -b.Velocity.Y,
		}
	}
}

//start using for starting the moving
func (b *Bird) start() {
	for {
		b.moveOne()
		time.Sleep(time.Millisecond * 5)
	}
}

//createBirds using for creating birds and push them into the queue
func createBirds(bid int) {
	bird := &Bird{
		Position: Vector2D{rand.Float64() * screenWidth, rand.Float64() * screenHeight},
		Velocity: Vector2D{(rand.Float64() * 2) - 1.0, (rand.Float64() * 2) - 1.0},
		id:       bid,
	}

	birds[bid] = bird
	go bird.start()
}
