package twod

import (
	"math"
	"math/cmplx"

	"github.com/kindermoumoute/adventofcode/pkg"
)

type Vector complex128

func NewVector(x, y int) Vector {
	return Vector(complex(float64(x), float64(y)))
}

func (v Vector) C() complex128 {
	return complex128(v)
}

func (v Vector) X() int {
	return toInt(real(complex128(v)))
}

func (v Vector) Y() int {
	return toInt(imag(complex128(v)))
}

func (v Vector) ScalarMult(i int) Vector {
	return NewVector(v.X()*i, v.Y()*i)
}

func (v Vector) DistFromOrigin() int {
	return v.DistFrom(0)
}

func (v Vector) DistFrom(from Vector) int {
	dist := (v - from)
	return pkg.Abs(dist.X()) + pkg.Abs(dist.Y())
}

func (v Vector) AngleFrom(from Vector) float64 {
	return cmplx.Phase((v - from).C())
}

func (v Vector) EuclideanDistFrom(from Vector) float64 {
	return cmplx.Abs((v - from).C())
}

func (v Vector) EuclideanDistFromOrigin() float64 {
	return v.EuclideanDistFrom(0)
}

func (v Vector) Rotate(angle float64) Vector {
	radius, currentAngle := cmplx.Polar(v.C())
	return Vector(complex128(cmplx.Rect(radius, currentAngle+angle)))
}

func toInt(f float64) int {
	return int(math.Round(f))
}