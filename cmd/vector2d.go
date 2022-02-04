package main

import "math"

type Vector2D struct {
	X float64
	Y float64
}

func (v *Vector2D) AddV(other *Vector2D) *Vector2D {
	return &Vector2D{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v *Vector2D) SubtractV(other *Vector2D) *Vector2D {
	return &Vector2D{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (v *Vector2D) MultiplyV(other *Vector2D) *Vector2D {
	return &Vector2D{
		X: v.X * other.X,
		Y: v.Y * other.Y,
	}
}

func (v *Vector2D) Add(num float64) *Vector2D {
	return &Vector2D{
		X: v.X + num,
		Y: v.Y + num,
	}
}

func (v *Vector2D) Subtract(num float64) *Vector2D {
	return &Vector2D{
		X: v.X + num,
		Y: v.Y + num,
	}
}

func (v *Vector2D) Multiply(num float64) *Vector2D {
	return &Vector2D{
		X: v.X + num,
		Y: v.Y + num,
	}
}

func (v *Vector2D) Division(num float64) *Vector2D {
	return &Vector2D{
		X: v.X / num,
		Y: v.Y / num,
	}
}

func (v *Vector2D) limit(lower, upper float64) *Vector2D {
	return &Vector2D{
		X: math.Min(math.Max(v.X, lower), upper),
		Y: math.Min(math.Max(v.Y, lower), upper),
	}
}

func (v *Vector2D) DistanceV(other *Vector2D) float64 {
	return math.Sqrt(math.Pow(v.X-other.X, 2) + math.Pow(v.Y-other.Y, 2))
}
