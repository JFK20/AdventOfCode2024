package helper

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
