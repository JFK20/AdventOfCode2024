package mathUtil

import "slices"

type Number interface {
	int | int32 | int64 | float32 | float64
}

type Vector2D[T Number] struct {
	X T
	Y T
}

func AddVector2D[T Number](left, right Vector2D[T]) Vector2D[T] {
	return Vector2D[T]{X: left.X + right.X, Y: left.Y + right.Y}
}

func SubVector2D[T Number](left, right Vector2D[T]) Vector2D[T] {
	return Vector2D[T]{X: left.X - right.X, Y: left.Y - right.Y}
}

func (vec *Vector2D[T]) EqualsVector2D(other Vector2D[T]) bool {
	return vec.X == other.X && vec.Y == other.Y
}

func (vec *Vector2D[T]) RotateVector2D() {
	tmp := vec.X
	vec.X = -vec.Y
	vec.Y = tmp
}

func (vec *Vector2D[T]) GetAllNeighbours() []Vector2D[T] {
	right := Vector2D[T]{X: vec.X + 1, Y: vec.Y}
	left := Vector2D[T]{X: vec.X - 1, Y: vec.Y}
	top := Vector2D[T]{X: vec.X, Y: vec.Y + 1}
	bottom := Vector2D[T]{X: vec.X, Y: vec.Y - 1}
	return []Vector2D[T]{right, left, top, bottom}
}

func (vec *Vector2D[T]) IsInBounds(bounds Vector2D[T]) bool {
	return vec.X >= 0 && vec.Y >= 0 && vec.X < bounds.X && vec.Y < bounds.Y
}

func (vec *Vector2D[T]) ConvertToInt() (Vector2D[int], error) {
	numX, errX := GetNearestInt(float64(vec.X))
	if errX != nil {
		return Vector2D[int]{}, errX
	}
	numY, errY := GetNearestInt(float64(vec.Y))
	if errY != nil {
		return Vector2D[int]{}, errY
	}
	return Vector2D[int]{X: numX, Y: numY}, nil
}

func Distinct(arr []Vector2D[int]) []Vector2D[int] {
	newArr := make([]Vector2D[int], 0)
	for _, val := range arr {
		if !slices.Contains(newArr, val) {
			newArr = append(newArr, val)
		}
	}
	return newArr
}
