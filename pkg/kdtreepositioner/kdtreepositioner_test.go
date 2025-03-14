package kdtreepositioner_test

import (
	"github.com/stretchr/testify/assert"
	"greeschenko/workoverlord2/pkg/kdtreepositioner"
	"testing"
)

type testObject struct {
	id    string
	coord [2]int
}

func (t testObject) ID() string {
	return t.id
}

func (t testObject) Coordinates() [2]int {
	return t.coord
}

func TestNearestNeighbor(t *testing.T) {
	objects := []kdtreepositioner.SpatialObject{
		testObject{"A", [2]int{1, 2}},
		testObject{"B", [2]int{3, 4}},
		testObject{"C", [2]int{5, 6}},
	}

	tree := kdtreepositioner.NewKDTree(objects, 0)

	nearest := tree.NearestNeighbor([2]int{4, 5})
	assert.Equal(t, "B", nearest)
}

func TestFindNearestInDirection(t *testing.T) {
	objects := []kdtreepositioner.SpatialObject{
		testObject{"A", [2]int{1, 1}},
		testObject{"B", [2]int{3, 3}},
		testObject{"C", [2]int{5, 5}},
	}

	tree := kdtreepositioner.NewKDTree(objects, 0)

	nearestRight := tree.FindNearestInDirection(testObject{"X", [2]int{2, 2}}, "right")
	assert.Equal(t, "B", nearestRight)
}

func TestRebuild(t *testing.T) {
	objects := []kdtreepositioner.SpatialObject{
		testObject{"A", [2]int{1, 1}},
	}

	tree := kdtreepositioner.NewKDTree(objects, 0)
	assert.Equal(t, "A", tree.NearestNeighbor([2]int{1, 1}))

	newObjects := []kdtreepositioner.SpatialObject{
		testObject{"B", [2]int{3, 3}},
	}
	tree.Rebuild(newObjects)
	assert.Equal(t, "B", tree.NearestNeighbor([2]int{3, 3}))
}
