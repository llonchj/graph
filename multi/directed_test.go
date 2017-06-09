// Copyright ©2014 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package multi

import (
	"math"
	"testing"
)

func TestMultigraph(t *testing.T) {
	nodes := []struct {
		srcID, targetID int
		weight          float64
	}{
		{0, 1, .1},
		{0, 1, .3},
		{1, 2, .2},
		{2, 0, .1},
		{0, 2, .3},
		{2, 3, .1},
	}

	g := NewDirectedGraph(0, math.Inf(1))
	for _, n := range nodes {
		f := Node(n.srcID)
		t := Node(n.targetID)
		m := Edge{F: f, T: t, W: n.weight}
		g.SetEdge(m)
	}

	if len(g.Edges()) != len(nodes) {
		t.Errorf("Edges mismatch. Got %d, expected %d\n", len(g.Edges()), len(nodes))
	}

	e := g.Edge(Node(0), Node(1)).(MultiEdge)
	if len(e.Edges()) != 2 {
		t.Errorf("MultiEdge mismatch. Got %d, expected %d\n", len(e.Edges()), 2)
	}

	r := Edge{F: 2, T: 3, W: 0.1}
	if !g.HasEdgeBetween(r.From(), r.To()) {
		t.Errorf("Graph should have edge %v\n", r)
	}
	g.RemoveEdge(r)
	if g.HasEdgeBetween(r.From(), r.To()) {
		t.Errorf("Edge %v was not removed\n", r)
	}
}
