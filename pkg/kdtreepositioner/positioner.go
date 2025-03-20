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

// SpatialObject represents an object with a unique ID, coordinates, and size.
type SpatialObject interface {
	ID() string
	Coordinates() [2]int
	WHSize() [2]int
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
	nodeCenter := getObjectCenter(n.Object)
	dist := pointDistance(nodeCenter, target)
	if dist < bestDist {
		best = n
		bestDist = dist
	}

	axis := n.Axis
	nodeCoord := nodeCenter

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
	node := n.closestObject(target, n, pointDistance(getObjectCenter(n.Object), target))
	if node != nil {
		return node.Object
	}
	return nil
}

// FindNearestInDirection finds the nearest object in a specified direction.
func (n *Node) FindNearestInDirection(target SpatialObject, direction string) SpatialObject {
	var bestDist = math.Inf(1)
	var bestObj SpatialObject

	targetCenter := getObjectCenter(target)

	var search func(node *Node)
	search = func(node *Node) {
		if node == nil {
			return
		}

		nodeCenter := getObjectCenter(node.Object)
		valid := false

		switch direction {
		case "down":
			valid = nodeCenter[1] > targetCenter[1]
		case "up":
			valid = nodeCenter[1] < targetCenter[1]
		case "left":
			valid = nodeCenter[0] < targetCenter[0]
		case "right":
			valid = nodeCenter[0] > targetCenter[0]
		}

		if valid {
			dist := pointDistance(nodeCenter, targetCenter)
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

// getObjectCenter calculates the center of a SpatialObject
func getObjectCenter(obj SpatialObject) [2]int {
	coord := obj.Coordinates()
	size := obj.WHSize()
	return [2]int{coord[0] + size[0]/2, coord[1] + size[1]/2}
}

// Computes the Euclidean distance between two points.
func pointDistance(a, b [2]int) float64 {
	dx := float64(a[0] - b[0])
	dy := float64(a[1] - b[1])
	return math.Sqrt(dx*dx + dy*dy)
}
