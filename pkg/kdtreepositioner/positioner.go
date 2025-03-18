package kdtreepositioner

import (
	"math"
	"sort"
)

// KDTree defines the interface for a k-d tree.
type KDTree interface {
	NearestNeighbor(target [2]int) SpatialObject
	FindNearestInDirection(target SpatialObject, direction string) SpatialObject
	Rebuild(objects []SpatialObject)
}

// SpatialObject represents an object with a unique ID and coordinates.
type SpatialObject interface {
	ID() string
	Coordinates() [2]int
    SetSelected(bool)
}

// Node represents a k-d tree node.
type Node struct {
	Object SpatialObject
	Left   *Node
	Right  *Node
	Axis   int
}

// NewKDTree constructs a k-d tree from a list of spatial objects.
func NewKDTree(objects []SpatialObject, depth int) *Node {
	if len(objects) == 0 {
		return nil
	}

	axis := depth % 2
	sort.Slice(objects, func(i, j int) bool {
		return objects[i].Coordinates()[axis] < objects[j].Coordinates()[axis]
	})

	mid := len(objects) / 2
	n := &Node{
		Object: objects[mid],
		Axis:   axis,
	}

	n.Left = NewKDTree(objects[:mid], depth+1)
	n.Right = NewKDTree(objects[mid+1:], depth+1)

	return n
}

// Rebuild reconstructs the KDTree from the given objects.
func (n *Node) Rebuild(objects []SpatialObject) {
	newTree := NewKDTree(objects, 0)
	*n = *newTree
}

// closestObject finds the nearest object to the given target.
func (n *Node) closestObject(target [2]int, best *Node, bestDist float64) *Node {
	if n == nil {
		return best
	}
	dist := pointDistance(n.Object.Coordinates(), target)
	if dist < bestDist {
		best = n
		bestDist = dist
	}

	axis := n.Axis
	nodeCoord := n.Object.Coordinates()

	var next, other *Node
	if target[axis] < nodeCoord[axis] {
		next, other = n.Left, n.Right
	} else {
		next, other = n.Right, n.Left
	}

	best = next.closestObject(target, best, bestDist)

	if math.Abs(float64(target[axis]-nodeCoord[axis])) < bestDist {
		best = other.closestObject(target, best, bestDist)
	}

	return best
}

// NearestNeighbor finds the closest spatial object to the given target coordinates.
func (n *Node) NearestNeighbor(target [2]int) SpatialObject {
	node := n.closestObject(target, n, pointDistance(n.Object.Coordinates(), target))
	if node != nil {
		return node.Object
	}
	return nil
}

// FindNearestInDirection finds the nearest object in a specified direction.
func (n *Node) FindNearestInDirection(target SpatialObject, direction string) SpatialObject {
	var bestDist = math.Inf(1)
	var bestObj SpatialObject

	var search func(node *Node)
	search = func(node *Node) {
		if node == nil {
			return
		}

		nodeCoord := node.Object.Coordinates()
		targetCoord := target.Coordinates()
		valid := false

		switch direction {
		case "down":
			valid = nodeCoord[1] > targetCoord[1]
		case "up":
			valid = nodeCoord[1] < targetCoord[1]
		case "left":
			valid = nodeCoord[0] < targetCoord[0]
		case "right":
			valid = nodeCoord[0] > targetCoord[0]
		}

		if valid {
			dist := pointDistance(nodeCoord, targetCoord)
			if dist < bestDist {
				bestDist = dist
				bestObj = node.Object
			}
		}

		search(node.Left)
		search(node.Right)
	}

	search(n)
	return bestObj
}

// Computes the Euclidean distance between two points.
func pointDistance(a, b [2]int) float64 {
	dx := float64(a[0] - b[0])
	dy := float64(a[1] - b[1])
	return math.Sqrt(dx*dx + dy*dy)
}
