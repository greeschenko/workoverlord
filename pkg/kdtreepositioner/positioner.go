package kdtreepositioner

import (
	"math"
	"sort"
	"fyne.io/fyne/v2"
)

// KDTree defines the interface for a k-d tree.
type KDTree interface {
	NearestNeighbor(target fyne.Position) fyne.CanvasObject
	FindNearestInDirection(target fyne.CanvasObject, direction string) fyne.CanvasObject
	Rebuild(objects []fyne.CanvasObject)
}

// Node represents a k-d tree node.
type Node struct {
	Object fyne.CanvasObject
	Left   *Node
	Right  *Node
	Axis   int
}

// NewKDTree constructs a k-d tree from a list of canvas objects.
func NewKDTree(objects []fyne.CanvasObject, depth int) *Node {
	if len(objects) == 0 {
		return nil
	}

	axis := depth % 2
	sort.Slice(objects, func(i, j int) bool {
		posI := objects[i].Position()
		posJ := objects[j].Position()
		if axis == 0 {
			return posI.X < posJ.X
		}
		return posI.Y < posJ.Y
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
func (n *Node) Rebuild(objects []fyne.CanvasObject) {
	newTree := NewKDTree(objects, 0)
	*n = *newTree
}

// closestObject finds the nearest object to the given target.
func (n *Node) closestObject(target fyne.Position, best *Node, bestDist float64) *Node {
	if n == nil {
		return best
	}
	dist := pointDistance(n.Object.Position(), target)
	if dist < bestDist {
		best = n
		bestDist = dist
	}

	axis := n.Axis
	nodeCoord := n.Object.Position()

	var next, other *Node
	if (axis == 0 && target.X < nodeCoord.X) || (axis == 1 && target.Y < nodeCoord.Y) {
		next, other = n.Left, n.Right
	} else {
		next, other = n.Right, n.Left
	}

	best = next.closestObject(target, best, bestDist)

	if math.Abs(float64(target.X)-float64(nodeCoord.X)) < bestDist || math.Abs(float64(target.Y)-float64(nodeCoord.Y)) < bestDist {
		best = other.closestObject(target, best, bestDist)
	}

	return best
}

// NearestNeighbor finds the closest canvas object to the given target coordinates.
func (n *Node) NearestNeighbor(target fyne.Position) fyne.CanvasObject {
	node := n.closestObject(target, n, pointDistance(n.Object.Position(), target))
	if node != nil {
		return node.Object
	}
	return nil
}

// FindNearestInDirection finds the nearest object in a specified direction.
func (n *Node) FindNearestInDirection(target fyne.CanvasObject, direction string) fyne.CanvasObject {
	var bestDist = math.Inf(1)
	var bestObj fyne.CanvasObject

	var search func(node *Node)
	search = func(node *Node) {
		if node == nil {
			return
		}

		nodeCoord := node.Object.Position()
		targetCoord := target.Position()
		valid := false

		switch direction {
		case "up":
			valid = nodeCoord.Y > targetCoord.Y
		case "down":
			valid = nodeCoord.Y < targetCoord.Y
		case "left":
			valid = nodeCoord.X < targetCoord.X
		case "right":
			valid = nodeCoord.X > targetCoord.X
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
func pointDistance(a, b fyne.Position) float64 {
	dx := float64(a.X) - float64(b.X)
	dy := float64(a.Y) - float64(b.Y)
	return math.Sqrt(dx*dx + dy*dy)
}
