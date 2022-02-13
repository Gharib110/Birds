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

func (b *Bird) calcAcceleration() *Vector2D {
	accelerator := &Vector2D{0, 0}

	return accelerator
}

//moveOne for moving a *Bird to the next position
func (b *Bird) moveOne() {
	b.Velocity = *b.Velocity.AddV(b.calcAcceleration()).limit(-1, 1)

	birdsMap[int(b.Position.X)][int(b.Position.Y)] = b.id
	b.Position = *b.Position.AddV(&b.Velocity)

	birdsMap[int(b.Position.X)][int(b.Position.Y)] = b.id
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

	birdsMap[int(bird.Position.X)][int(bird.Position.Y)] = bird.id

	birds[bid] = bird
	go bird.start()
}
