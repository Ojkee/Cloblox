package functools_test

import (
	"testing"

	"Cloblox/functools"
)

func TestMinMax_1(t *testing.T) {
	data := []float64{2, 3, 0, -2, 55, 0.3, -4, 2}
	mmin, mmax := functools.GetMinMax(&data)
	if mmin != -4.0 || mmax != 55.0 {
		t.Errorf("mmin = %f\nmmax = %f", mmin, mmax)
	}
}

func TestMinMax_2(t *testing.T) {
	data := []float64{2, 3, 0, -2, 55, 0.3, -4, 2, -44}
	mmin, mmax := functools.GetMinMax(&data)
	if mmin != -44.0 || mmax != 55.0 {
		t.Errorf("mmin = %f\nmmax = %f", mmin, mmax)
	}
}

func TestMinMax_3(t *testing.T) {
	data := []float64{2, 3, 0, -2, 55, 0.3, -4, 2, 66}
	mmin, mmax := functools.GetMinMax(&data)
	if mmin != -4.0 || mmax != 66.0 {
		t.Errorf("mmin = %f\nmmax = %f", mmin, mmax)
	}
}

func TestMinMax_4(t *testing.T) {
	data := []float64{}
	mmin, mmax := functools.GetMinMax(&data)
	if mmin != 0.0 || mmax != 0.0 {
		t.Errorf("mmin = %f\nmmax = %f", mmin, mmax)
	}
}

func TestMinMax_5(t *testing.T) {
	data := []float64{-1, 2, -2, 4, -3, 6, 0, -1, 3, -2}
	mmin, mmax := functools.GetMinMax(&data)
	if mmin != -3.0 || mmax != 6.0 {
		t.Errorf("mmin = %f\nmmax = %f", mmin, mmax)
	}
}

func TestMinMax_6(t *testing.T) {
	data := []float64{-1, 2, -2, 4, -3, 6, 0, -1, 3, 88}
	mmin, mmax := functools.GetMinMax(&data)
	if mmin != -3.0 || mmax != 88.0 {
		t.Errorf("mmin = %f\nmmax = %f", mmin, mmax)
	}
}

func TestMinMax_7(t *testing.T) {
	data := []float64{-1, 2, -2, 4, -3, 6, 0, -1, 3, -88}
	mmin, mmax := functools.GetMinMax(&data)
	if mmin != -88.0 || mmax != 6.0 {
		t.Errorf("mmin = %f\nmmax = %f", mmin, mmax)
	}
}

func TestMinMax_8(t *testing.T) {
	data := []float64{1}
	mmin, mmax := functools.GetMinMax(&data)
	if mmin != 1.0 || mmax != 1.0 {
		t.Errorf("mmin = %f\nmmax = %f", mmin, mmax)
	}
}

func TestMinMax_9(t *testing.T) {
	data := []float64{1, 2}
	mmin, mmax := functools.GetMinMax(&data)
	if mmin != 1.0 || mmax != 2.0 {
		t.Errorf("mmin = %f\nmmax = %f", mmin, mmax)
	}
}
