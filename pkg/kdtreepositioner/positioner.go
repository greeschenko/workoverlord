package kdtreepositioner

import (
	"math"
	"sort"
)

// SpatialObject represents an object in 2D space with an ID
// and coordinates.
type SpatialObject interface {
	ID() string
	Coordinates() [2]int
}

// Node represents a k-d tree node
type Node struct {
	Object SpatialObject
	Left   *Node
	Right  *Node
	Axis   int
}

// NewKDTree constructs a k-d tree from a list of spatial objects
func NewKDTree(objects []SpatialObject, depth int) *Node {
	if len(objects) == 0 {
		return nil
	}

	axis := depth % 2
	sort.Slice(objects, func(i, j int) bool {
		coord1 := objects[i].Coordinates()
		coord2 := objects[j].Coordinates()
		if axis == 0 {
			return coord1[0] < coord2[0]
		}
		return coord1[1] < coord2[1]
	})

	mid := len(objects) / 2
	return &Node{
		Object: objects[mid],
		Left:   NewKDTree(objects[:mid], depth+1),
		Right:  NewKDTree(objects[mid+1:], depth+1),
		Axis:   axis,
	}
}

// Finds the closest object to a given target
func (n *Node) closestObject(target SpatialObject, best *Node, bestDist float64) *Node {
	if n == nil {
		return best
	}
	dist := distance(n.Object, target)
	if dist < bestDist {
		best = n
		bestDist = dist
	}

	axis := n.Axis
	var next, other *Node
	targetCoord := target.Coordinates()
	nodeCoord := n.Object.Coordinates()

	if (axis == 0 && targetCoord[0] < nodeCoord[0]) || (axis == 1 && targetCoord[1] < nodeCoord[1]) {
		next, other = n.Left, n.Right
	} else {
		next, other = n.Right, n.Left
	}

	best = next.closestObject(target, best, bestDist)

	if math.Abs(float64(targetCoord[axis]-nodeCoord[axis])) < bestDist {
		best = other.closestObject(target, best, bestDist)
	}

	return best
}

// NearestNeighbor returns the ID of the closest object to the given target
func (n *Node) NearestNeighbor(target SpatialObject) string {
	node := n.closestObject(target, n, distance(n.Object, target))
	return node.Object.ID()
}

// Find the nearest object in a specified direction
func (n *Node) FindNearestInDirection(target SpatialObject, direction string) string {
	var best SpatialObject
	var bestDist = math.Inf(1)
	bestID := ""

	var search func(*Node)
	search = func(node *Node) {
		if node == nil {
			return
		}

		nodeCoord := node.Object.Coordinates()
		targetCoord := target.Coordinates()

		valid := false
		switch direction {
		case "up":
			valid = nodeCoord[1] > targetCoord[1]
		case "down":
			valid = nodeCoord[1] < targetCoord[1]
		case "left":
			valid = nodeCoord[0] < targetCoord[0]
		case "right":
			valid = nodeCoord[0] > targetCoord[0]
		}

		if valid {
			dist := distance(node.Object, target)
			if dist < bestDist {
				best = node.Object
				bestDist = dist
				bestID = node.Object.ID()
			}
		}

		search(node.Left)
		search(node.Right)
	}

	search(n)
	return bestID
}

// Computes the Euclidean distance between two spatial objects
func distance(a, b SpatialObject) float64 {
	coordA := a.Coordinates()
	coordB := b.Coordinates()
	dx, dy := float64(coordA[0]-coordB[0]), float64(coordA[1]-coordB[1])
	return math.Sqrt(dx*dx + dy*dy)
}
