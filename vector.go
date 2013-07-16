package vector

import (
	"math"
)

type Vec2d [2]float64
type Vec3d [3]float64

//Sum returns the elementwise sum of a and b.
func (a Vec3d) Sum(b Vec2d) Vec2d {return Vec2d{a[0] + b[0], a[1] + b[1]}}
func (a Vec3d) Sum(b Vec3d) Vec3d {return Vec3d{a[0] + b[0], a[1] + b[1], a[2] + b[2]}}

//Dif returns the difference of a and b (b - a).
func (a Vec2d) Dif(b Vec2d) Vec2d {return Vec2d{b[0] - a[0], b[1] - a[1]}}
func (a Vec3d) Dif(b Vec3d) Vec3d {return Vec3d{b[0] - a[0], b[1] - a[1], b[2] - a[2]}}

//Cross returns the cross product (vector product) of a and b.
func (a Vec3d) Cross(b Vec3d) Vec3d {return Vec3d{a[1]*b[2] - a[2]*b[1], a[2]*b[0] - a[0]*b[2], a[0]*b[1] - a[1]*b[0]}}

//Dot returns the dot product (scalar product) of a and b.
func (a Vec2d) Dot(b Vec2d) float64 {return a[0]*b[0] + a[1]*b[1]}
func (a Vec3d) Dot(b Vec3d) float64 {return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]}

//Scale scales (multiplies) a by k.
func (a Vec2d) Scale(k float64) Vec2d {return Vec2d{k*a[0], k*a[1]}}
func (a Vec3d) Scale(k float64) Vec3d {return Vec3d{k*a[0], k*a[1], k*a[2]}}

//Inv returns the multiplicative inverse (reciprocal) of a.
func (a Vec3d) Inv() Vec3d {return Vec3d{1/a[0], 1/a[1], 1/a[2]}}

//AbsSquared returns the square of the absolute value of a vector.
func (a Vec3d) AbsSquared() float64 {return a[0]*a[0] + a[1]*a[1] + a[2]*a[2]}

//Abs() returns the absolute value (length or modulus) of a.
func (a Vec3d) Abs() float64 {return math.Sqrt(a.AbsSquared())}

//Dir() returns the direction vector (unit vector) of a.
func (a Vec3d) Dir() Vec3d {
	invAbs := 1/a.Abs()
	return Vec3d{a[0] * invAbs, a[1] * invAbs, a[2] * invAbs}
}

/*Rotate() returns a rotated angle radians about unit vector axis using
Rodrigues' rotation formula*/
func (a Vec3d) Rotate(axis Vec3d, angle float64) Vec3d {
	return a.Scale(math.Cos(angle)).Sum(axis.Cross(a).Scale(math.Sin(angle))).Sum(axis.Scale(axis.Dot(a)*(1-math.Cos(angle))))
}
