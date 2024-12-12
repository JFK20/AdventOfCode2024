package vec

type Vector2D struct {
	X int
	Y int
}

func AddVector2D(left Vector2D, right Vector2D) Vector2D {
	return Vector2D{left.X + right.X, left.Y + right.Y}
}

func SubVector2D(left Vector2D, right Vector2D) Vector2D {
	return Vector2D{left.X - right.X, left.Y - right.Y}
}

func (vec *Vector2D) EqualsVector2D(left Vector2D) bool {
	if vec.X != left.X || vec.Y != left.Y {
		return false
	}
	return true
}

func (vec *Vector2D) RotateVector2D() {
	tmp := vec.X
	vec.X = -vec.Y
	vec.Y = tmp
}

func (vec *Vector2D) GetAllNeighbours() []Vector2D {
	right := Vector2D{vec.X + 1, vec.Y}
	left := Vector2D{vec.X - 1, vec.Y}
	top := Vector2D{vec.X, vec.Y + 1}
	bottom := Vector2D{vec.X, vec.Y - 1}
	return []Vector2D{right, left, top, bottom}
}

func (vec *Vector2D) IsInBounds(bounds Vector2D) bool {
	if vec.X >= bounds.X || vec.Y >= bounds.Y {
		return false
	}
	if vec.X < 0 || vec.Y < 0 {
		return false
	}
	return true
}
