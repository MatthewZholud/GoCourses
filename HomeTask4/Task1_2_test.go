package main

import (
	"testing"
)

func TestSquare(t *testing.T) {
	//Testing Perimeter
	testCasesPerimeter := []struct {
		in  int64
		out int64
		err error
	}{
		{-1, 0, ErrInputSmall},
		{101, 0, ErrInputBig},
		{0, 0, nil},
		{7, 28, nil},
	}

	for _, c := range testCasesPerimeter {
		res, err := Perimeter(c.in)
		if res != c.out || err != c.err {
			t.Errorf("Perimeter(%d) returns %v, %v; expected %v, %v", c.in, res, err, c.out, c.err)
		}
	}

	//Testing Area
	testCasesArea := []struct {
		in  int64
		out int64
		err error
	}{
		{-1, 0, ErrInputSmall},
		{101, 0, ErrInputBig},
		{0, 0, nil},
		{7, 49, nil},
	}
	for _, c := range testCasesArea {
		res, err := Area(c.in)
		if res != c.out || err != c.err {
			t.Errorf("Area(%d) returns %v, %v; expected %v, %v", c.in, res, err, c.out, c.err)
		}
	}

	//Testing End point
	testCasesEnd := []struct {
		inA  int64
		inX  int64
		inY  int64
		outX int64
		outY int64
		err  error
	}{
		{-1, -1, -1, 0, 0, ErrInputSmall},
		{101, 101, 101, 0, 0, ErrInputBig},
		{0, 7, 7, 7, 7, nil},
		{7, 0, 0, 7, -7, nil},
	}
	for _, c := range testCasesEnd {
		res1, res2, err := End(c.inX, c.inY, c.inA)
		if res1 != c.outX || err != c.err || res2 != c.outY {
			t.Errorf("End(%v, %v, %v) returns %v, %v, %v; expected %v, %v, %v", c.inX, c.inY, c.inA, res1, res2, c.err, c.outX, c.outY, c.err)
		}
	}
}
