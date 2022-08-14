package utilities

import "math"

type Vec3 struct {
	X, Y, Z float64
}

type Point3 Vec3
type Color Vec3

func (u *Vec3) Add(v Vec3) {
	u.X += v.X
	u.Y += v.Y
	u.Z += v.Z
}

func (u *Vec3) Sub(v Vec3) {
	u.X -= v.X
	u.Y -= v.Y
	u.Z -= v.Z
}

func (u *Vec3) Negate() {
	u.X *= -1
	u.Y *= -1
	u.Z *= -1
}

func (u *Vec3) Multiply(t float64) {
	u.X *= t
	u.Y *= t
	u.Z *= t
}

func (u *Vec3) Divide(c float64) {
	u.Multiply(1 / c)
}

func (u *Vec3) Length() float64 {
	return math.Sqrt(u.LengthSqured())
}

func (u *Vec3) LengthSqured() float64 {
	return u.X*u.X + u.Y*u.Y + u.Z*u.Z
}

// Vec3 Utility Function
func Add(u *Vec3, v *Vec3) Vec3 {
	return Vec3{u.X + v.X, u.Y + v.Y, u.Z + v.Z}
}

func Sub(u *Vec3, v *Vec3) Vec3 {
	return Vec3{u.X + v.X, u.Y + v.Y, u.Z + v.Z}
}

func Multiply(u *Vec3, t float64) Vec3 {
	return Vec3{t * u.X, t * u.Y, t * u.Z}
}

func Divide(u *Vec3, t float64) Vec3 {
	return Multiply(u, 1/t)
}

func Dot(u *Vec3, v *Vec3) float64 {
	return u.X*v.X + u.Y*v.Y + u.Z*v.Z
}

func Cross(u *Vec3, v *Vec3) Vec3 {
	return Vec3{u.Y*v.Z - u.Z*v.Y, u.Z*v.X - u.X*v.Z, u.X*v.Y - u.Y*v.X}
}

func Unit(u *Vec3) Vec3 {
	return Divide(u, u.Length())
}
