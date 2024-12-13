package mathUtil

type Matrix[T Number] struct {
	Prefactor float64
	A         T
	B         T
	C         T
	D         T
}

func (m *Matrix[T]) Invert() {
	var under = float64(m.A*m.D - m.B*m.C)
	m.Prefactor = 1.0 / under
	tmp := m.D
	m.D = m.A
	m.A = tmp
	m.B = -m.B
	m.C = -m.C
}

func (m *Matrix[T]) Multiply(v Vector2D[T]) Vector2D[T] {
	ret := Vector2D[T]{}
	a := m.Prefactor * float64(m.A) * float64(v.X)
	b := m.Prefactor * float64(m.B) * float64(v.Y)
	c := m.Prefactor * float64(m.C) * float64(v.X)
	d := m.Prefactor * float64(m.D) * float64(v.Y)
	ret.X = T(a + b)
	ret.Y = T(c + d)
	return ret
}
