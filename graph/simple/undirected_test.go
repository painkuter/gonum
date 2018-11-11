// Copyright ©2014 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package simple_test

import (
	"math"
	"testing"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/internal/set"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/testgraph"
)

func undirectedBuilder(nodes []graph.Node, edges []graph.WeightedEdge, _, _ float64) (g graph.Graph, n []graph.Node, e []graph.Edge, s, a float64, ok bool) {
	seen := make(set.Nodes)
	ug := simple.NewUndirectedGraph()
	for _, n := range nodes {
		seen.Add(n)
		ug.AddNode(n)
	}
	for _, edge := range edges {
		f := ug.Node(edge.From().ID())
		if f == nil {
			f = edge.From()
		}
		t := ug.Node(edge.To().ID())
		if t == nil {
			t = edge.To()
		}
		ce := simple.Edge{F: f, T: t}
		seen.Add(ce.F)
		seen.Add(ce.T)
		e = append(e, ce)
		ug.SetEdge(ce)
	}
	if len(seen) != 0 {
		n = make([]graph.Node, 0, len(seen))
	}
	for _, sn := range seen {
		n = append(n, sn)
	}
	return ug, n, e, math.NaN(), math.NaN(), true
}

func TestUndirected(t *testing.T) {
	t.Run("EdgeExistence", func(t *testing.T) {
		testgraph.EdgeExistence(t, undirectedBuilder)
	})
	t.Run("NodeExistence", func(t *testing.T) {
		testgraph.NodeExistence(t, undirectedBuilder)
	})
	t.Run("ReturnAdjacentNodes", func(t *testing.T) {
		testgraph.ReturnAdjacentNodes(t, undirectedBuilder)
	})
	t.Run("ReturnAllEdges", func(t *testing.T) {
		testgraph.ReturnAllEdges(t, undirectedBuilder)
	})
	t.Run("ReturnAllNodes", func(t *testing.T) {
		testgraph.ReturnAllNodes(t, undirectedBuilder)
	})
	t.Run("ReturnEdgeSlice", func(t *testing.T) {
		testgraph.ReturnEdgeSlice(t, undirectedBuilder)
	})
	t.Run("ReturnNodeSlice", func(t *testing.T) {
		testgraph.ReturnNodeSlice(t, undirectedBuilder)
	})
}

func TestAssertMutableNotDirected(t *testing.T) {
	var g graph.UndirectedBuilder = simple.NewUndirectedGraph()
	if _, ok := g.(graph.Directed); ok {
		t.Fatal("Graph is directed, but a MutableGraph cannot safely be directed!")
	}
}

func TestMaxID(t *testing.T) {
	g := simple.NewUndirectedGraph()
	nodes := make(map[graph.Node]struct{})
	for i := simple.Node(0); i < 3; i++ {
		g.AddNode(i)
		nodes[i] = struct{}{}
	}
	g.RemoveNode(int64(0))
	delete(nodes, simple.Node(0))
	g.RemoveNode(int64(2))
	delete(nodes, simple.Node(2))
	n := g.NewNode()
	g.AddNode(n)
	if g.Node(n.ID()) == nil {
		t.Error("added node does not exist in graph")
	}
	if _, exists := nodes[n]; exists {
		t.Errorf("Created already existing node id: %v", n.ID())
	}
}

// Test for issue #123 https://github.com/gonum/graph/issues/123
func TestIssue123UndirectedGraph(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("unexpected panic: %v", r)
		}
	}()
	g := simple.NewUndirectedGraph()

	n0 := g.NewNode()
	g.AddNode(n0)

	n1 := g.NewNode()
	g.AddNode(n1)

	g.RemoveNode(n0.ID())

	n2 := g.NewNode()
	g.AddNode(n2)
}
