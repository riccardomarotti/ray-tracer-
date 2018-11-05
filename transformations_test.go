package main

import (
	"math"
	"testing"
)

func TestTranslation(t *testing.T) {
	translation := Translation(5, -3, 2)
	p := Point(-3, 4, 5)

	expected := Point(2, 1, 7)
	actual := translation.MultiplyByTuple(p)

	AssertTupleEqual(expected, actual, t)

	invTranslation := translation.Inverse()
	expectedInverseTrnslatedPoint := Point(-8, 7, 3)

	AssertTupleEqual(expectedInverseTrnslatedPoint, invTranslation.MultiplyByTuple(p), t)
}

func TestTranslationDoesNotAffectVectors(t *testing.T) {
	translation := Translation(5, -3, 2)
	v := Vector(-3, 4, 5)

	AssertTupleEqual(v, translation.MultiplyByTuple(v), t)
}

func TestScalingPoint(t *testing.T) {
	S := Scaling(2, 3, 4)
	p := Point(-4, 6, 8)

	expectedScaledPoint := Point(-8, 18, 32)
	actualScaledPoint := S.MultiplyByTuple(p)

	AssertTupleEqual(expectedScaledPoint, actualScaledPoint, t)
}

func TestScalingVector(t *testing.T) {
	S := Scaling(2, 3, 4)
	v := Vector(-4, 6, 8)

	expectedScaledVector := Vector(-8, 18, 32)
	actualScaledVector := S.MultiplyByTuple(v)

	AssertTupleEqual(expectedScaledVector, actualScaledVector, t)
}

func TestScalingInverse(t *testing.T) {
	S := Scaling(2, 3, 4)
	Sinv := S.Inverse()
	p := Point(-4, 6, 8)

	expectedScaledPoint := Point(-2, 2, 2)
	actualScaledPoint := Sinv.MultiplyByTuple(p)

	AssertTupleEqual(expectedScaledPoint, actualScaledPoint, t)
}

func TestReflectionIsScalingBtNegativeNumber(t *testing.T) {
	R := Scaling(-1, 1, 1)
	p := Point(2, 3, 4)

	AssertTupleEqual(Point(-2, 3, 4), R.MultiplyByTuple(p), t)
}

func TestRotatePointAroundXAxis(t *testing.T) {
	p := Point(0, 1, 0)

	halfQuarterRotation := RotationX(math.Pi / 4)
	quarterRotation := RotationX(math.Pi / 2)

	AssertTupleEqual(Point(0, math.Sqrt(2)/2, math.Sqrt(2)/2), halfQuarterRotation.MultiplyByTuple(p), t)
	AssertTupleEqual(Point(0, 0, 1), quarterRotation.MultiplyByTuple(p), t)
}

func TestInverseOfRotationRatatesInTheOppositeDirection(t *testing.T) {
	p := Point(0, 1, 0)

	halfQuarterRotation := RotationX(math.Pi / 4)
	inverseHalfQuarterRotation := halfQuarterRotation.Inverse()

	AssertTupleEqual(Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2), inverseHalfQuarterRotation.MultiplyByTuple(p), t)
}

func TestRotatePointAroundYAxis(t *testing.T) {
	p := Point(0, 0, 1)

	halfQuarterRotation := RotationY(math.Pi / 4)
	quarterRotation := RotationY(math.Pi / 2)

	AssertTupleEqual(Point(math.Sqrt(2)/2, 0, math.Sqrt(2)/2), halfQuarterRotation.MultiplyByTuple(p), t)
	AssertTupleEqual(Point(1, 0, 0), quarterRotation.MultiplyByTuple(p), t)
}

func TestRotatePointAroundZAxis(t *testing.T) {
	p := Point(0, 1, 0)

	halfQuarterRotation := RotationZ(math.Pi / 4)
	quarterRotation := RotationZ(math.Pi / 2)

	AssertTupleEqual(Point(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0), halfQuarterRotation.MultiplyByTuple(p), t)
	AssertTupleEqual(Point(-1, 0, 0), quarterRotation.MultiplyByTuple(p), t)
}

func TestShearingXtoY(t *testing.T) {
	shearing := Shearing(1, 0, 0, 0, 0, 0)
	p := Point(2, 3, 4)

	AssertTupleEqual(Point(5, 3, 4), shearing.MultiplyByTuple(p), t)
}

func TestShearingXtoZ(t *testing.T) {
	shearing := Shearing(0, 1, 0, 0, 0, 0)
	p := Point(2, 3, 4)

	AssertTupleEqual(Point(6, 3, 4), shearing.MultiplyByTuple(p), t)
}

func TestShearingYtoX(t *testing.T) {
	shearing := Shearing(0, 0, 1, 0, 0, 0)
	p := Point(2, 3, 4)

	AssertTupleEqual(Point(2, 5, 4), shearing.MultiplyByTuple(p), t)
}

func TestShearingYtoZ(t *testing.T) {
	shearing := Shearing(0, 0, 0, 1, 0, 0)
	p := Point(2, 3, 4)

	AssertTupleEqual(Point(2, 7, 4), shearing.MultiplyByTuple(p), t)
}

func TestShearingZtoX(t *testing.T) {
	shearing := Shearing(0, 0, 0, 0, 1, 0)
	p := Point(2, 3, 4)

	AssertTupleEqual(Point(2, 3, 6), shearing.MultiplyByTuple(p), t)
}

func TestShearingZtoY(t *testing.T) {
	shearing := Shearing(0, 0, 0, 0, 0, 1)
	p := Point(2, 3, 4)

	AssertTupleEqual(Point(2, 3, 7), shearing.MultiplyByTuple(p), t)
}
