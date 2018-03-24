package goalgorithms

import (
	"container/heap"
	"fmt"
	"image"
	"log"
	"math"
)

// RVertex repsesent a vertex on the resulting voronoi diagram.
// Will be replaced with a double-connected edge list structure in a later
// version of the library.
type RVertex struct {
	X int
	Y int
}

type Voronoi struct {
	Bounds       image.Rectangle
	Sites        SiteSlice
	EventQueue   EventQueue
	ParabolaTree *VNode
	SweepLine    int // tracks the current position of the sweep line; updated when a new site is added.
	Result       []RVertex
}

func NewVoronoi(sites SiteSlice, bounds image.Rectangle) *Voronoi {
	voronoi := &Voronoi{Bounds: bounds}
	voronoi.Sites = make(SiteSlice, len(sites), len(sites))
	copy(voronoi.Sites, sites)
	voronoi.init()
	return voronoi
}

func NewFromPoints(points []image.Point, bounds image.Rectangle) *Voronoi {
	var sites SiteSlice
	for _, point := range points {
		sites = append(sites, Site{point.X, point.Y})
	}
	return NewVoronoi(sites, bounds)
}

func (v *Voronoi) init() {
	// 1. Push sites to a priority queue, sorted by by Y
	// 2. Create empty binary tree for parabola arcs
	// 3. Create empty doubly-connected edge list (DCEL) for the voronoi diagram

	// 1. Push sites to a priority queue, sorted by by Y
	v.EventQueue = NewEventQueue(v.Sites)

	// 2. Create empty binary tree for parabola arcs
	v.ParabolaTree = nil

	// 3. Create empty doubly-connected edge list (DCEL) for the voronoi diagram
	// TODO: Create DCEL list
}

func (v *Voronoi) Reset() {
	v.EventQueue = NewEventQueue(v.Sites)
	v.ParabolaTree = nil
	v.Result = make([]RVertex, 0)
	v.SweepLine = 0
}

func (v *Voronoi) HandleNextEvent() {
	if v.EventQueue.Len() > 0 {
		// Process events by Y (priority)
		event := heap.Pop(&v.EventQueue).(*Event)

		// Event with Y above the sweep line should be ignored.
		if event.Y < v.SweepLine {
			log.Printf("Ignoring event as it's above the sweep line (%d)\r\n", v.SweepLine)
			return
		}

		v.SweepLine = event.Y
		if event.EventType == EventSite {
			v.handleSiteEvent(event)
		} else {
			v.handleCircleEvent(event)
		}
	}
}

func (v *Voronoi) Generate() {
	v.Reset()

	// While queue is not empty
	for v.EventQueue.Len() > 0 {
		v.HandleNextEvent()
	}
}

// findNodeAbove finds the node for the parabola that is vertically above the specified site.
func (v *Voronoi) findNodeAbove(site Site) *VNode {
	node := v.ParabolaTree

	for !node.IsLeaf() {
		if node.IsLeaf() {
			log.Printf("At leaf %d,%d\r\n", node.Site.X, node.Site.Y)
		} else {
			log.Printf(
				"At internal node %d,%d <-> %d,%d\r\n",
				node.PrevChildArc().Site.X, node.PrevChildArc().Site.Y,
				node.NextChildArc().Site.X, node.NextChildArc().Site.Y,
			)
		}

		x := GetXOfIntersection(node, v.SweepLine)
		if site.X < x {
			log.Printf("site.X (%d) < x (%d), going left\r\n", site.X, x)
			node = node.Left
		} else {
			log.Printf("site.X (%d) >= x (%d), going right\r\n", site.X, x)
			node = node.Right
		}
		if node.IsLeaf() {
			log.Printf("X of intersection: %d\r\n", x)
		}
	}

	return node
}

func (v *Voronoi) handleSiteEvent(event *Event) {
	log.Println()
	log.Printf("Handling site event %d:%d\r\n", event.X, event.Y)
	log.Printf("Sweep line: %d", v.SweepLine)
	log.Printf("Tree: %v", v.ParabolaTree)

	eventSite := Site{event.X, event.Y}

	// If the binary tree is empty, just add an arc for this site as the only leaf in the tree
	if v.ParabolaTree == nil {
		log.Print("Adding event as root\r\n")
		v.ParabolaTree = &VNode{Site: eventSite}
		return
	}

	// If the tree is not empty, find the arc vertically above the new site
	arcAbove := v.findNodeAbove(eventSite)
	if arcAbove == nil {
		log.Print("Could not find arc above event site!\r\n")
		// Do something
		return
	}
	log.Printf("Arc above: %d:%d\r\n", arcAbove.Site.X, arcAbove.Site.Y)

	if len(arcAbove.Events) > 0 {
		log.Printf("Removing %d events from queue.\r\n", len(arcAbove.Events))

		// Remove false circle events from queue
		for _, e := range arcAbove.Events {
			if e.index > -1 {
				v.EventQueue.Remove(e)
			}
		}
		arcAbove.Events = nil
	}

	y := GetYByX(arcAbove.Site, eventSite.X, v.SweepLine)
	point := RVertex{eventSite.X, y}
	log.Printf("Y of intersection = %d:%d\r\n", point.X, point.Y)
	v.Result = append(v.Result, point)

	// The node above (NA) is replaced wit ha branch with one internal node and three leafs.
	// The middle leaf stores the new parabola and the other two store the one being split.
	//    (NA)
	//   /   \
	//  (  )  [old]
	// /    \
	//[old]  [new]

	// Copy of the old arc
	arcAbove.Right = &VNode{
		Site:   arcAbove.Site,
		Events: arcAbove.Events,
		Parent: arcAbove,
	}

	// Internal node
	arcAbove.Left = &VNode{Parent: arcAbove}

	// The new arc
	arcAbove.Left.Right = &VNode{
		Site:   eventSite,
		Parent: arcAbove.Left,
	}
	newArc := arcAbove.Left.Right

	// Copy of the old arc
	arcAbove.Left.Left = &VNode{
		Site:   arcAbove.Site,
		Events: arcAbove.Events,
		Parent: arcAbove.Left,
	}

	arcAbove.Site.X = 0
	arcAbove.Site.Y = 0
	arcAbove.Events = nil

	// Check for circle events where the new arc is the right most arc
	prevArc := newArc.PrevArc()
	log.Printf("Prev arc for %v is %v\r\n", newArc, prevArc)
	if prevArc != nil {
		prevPrevArc := prevArc.PrevArc()
		log.Printf("Prev->prev arc for %v is %v\r\n", newArc, prevPrevArc)
		v.addCircleEvent(prevPrevArc, prevArc, newArc)
	}

	// Check for circle events where the new arc is the left most arc
	nextArc := newArc.NextArc()
	if nextArc != nil {
		nextNextArc := nextArc.NextArc()
		v.addCircleEvent(newArc, nextArc, nextNextArc)
	}
}

func (v *Voronoi) calcCircle(site1, site2, site3 Site) (x int, y int, r int, err error) {
	// Solution by https://math.stackexchange.com/a/1268279/543428
	// Explanation at http://mathforum.org/library/drmath/view/55002.html
	x = 0
	y = 0
	r = 0
	err = nil

	x1 := float64(site1.X)
	y1 := float64(site1.Y)

	x2 := float64(site2.X)
	y2 := float64(site2.Y)

	x3 := float64(site3.X)
	y3 := float64(site3.Y)

	if x2-x1 == 0 || x3-x2 == 0 {
		err = fmt.Errorf("no circle found connecting points %f,%f %f,%f and %f,%f", x1, y1, x2, y2, x3, y3)
		return
	}

	mr := (y2 - y1) / (x2 - x1)
	mt := (y3 - y2) / (x3 - x2)

	if mr == mt || mr-mt == 0 || mr == 0 {
		err = fmt.Errorf("no circle found connecting points %f,%f %f,%f and %f,%f", x1, y1, x2, y2, x3, y3)
		return
	}

	cx := (mr*mt*(y3-y1) + mr*(x2+x3) - mt*(x1+x2)) / (2 * (mr - mt))
	cy := (y1+y2)/2 - (cx-(x1+x2)/2)/mr
	cr := math.Pow((math.Pow((x2-cx), 2) + math.Pow((y2-cy), 2)), 0.5)

	x = int(cx + 0.5)
	y = int(cy + 0.5)
	r = int(cr + 0.5)

	return
}

func (v *Voronoi) addCircleEvent(arc1, arc2, arc3 *VNode) {
	if arc1 == nil || arc2 == nil || arc3 == nil {
		return
	}

	log.Printf("Checking for circle at <%v> <%v> <%v>\r\n", arc1.Site, arc2.Site, arc3.Site)
	x, y, r, err := v.calcCircle(arc1.Site, arc2.Site, arc3.Site)
	if err != nil {
		return
	}

	// Only add events with bottom point below the sweep line
	bottomY := y + r
	if bottomY <= v.SweepLine {
		return
	}

	event := &Event{
		EventType: EventCircle,
		X:         x,
		Y:         bottomY,
		Radius:    r,
	}
	v.EventQueue.Push(event)

	arc2.AddEvent(event)
	event.Node = arc2

	log.Printf("Added circle with center %d,%d an r=%d\r\n", x, y, r)
}

func (v *Voronoi) handleCircleEvent(event *Event) {
	log.Println()
	log.Printf("Handling circle event %d:%d with radius %d\r\n", event.X, event.Y, event.Radius)
	log.Printf("Sweep line: %d", v.SweepLine)
	log.Printf("Tree: %v", v.ParabolaTree)

	log.Printf("Node to be removed: %v", event.Node)

	// Remove other circle events, in which this node participates
	if len(event.Node.Events) > 0 {
		log.Printf("Removing %d events from queue.\r\n", len(event.Node.Events))

		// Remove false circle events from queue
		for _, e := range event.Node.Events {
			if e.index > -1 {
				v.EventQueue.Remove(e)
			}
		}
		event.Node.Events = nil
	}

	// Add center of circle as vertex
	point := RVertex{event.X, event.Y - event.Radius}
	log.Printf("Vertex at %d:%d (center of circle)\r\n", point.X, point.Y)
	v.Result = append(v.Result, point)

	return
}
