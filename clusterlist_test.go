package clusterlist

import v "vertex"
import "testing"

func TestEmptyListSize(t *testing.T) {
	cList := NewClusterList()
	if cList.Len() != 0 {
		t.Errorf("Empty ClusterList has nonzero length")
	}
}

func TestAddIncrementsSize(t *testing.T) {
	cList := NewClusterList()
	cList.Add(&v.Vertex{12.2, 13.4})
	if cList.Len() != 1 {
		t.Errorf("Cluster length is not 1 it is = ")
	}
}

func TestAppend(t *testing.T) {
	cl1 := NewClusterList()
	cl1.Add(&v.Vertex{1.2, 1.3})
	cl2 := NewClusterList()
	cl2.Add(&v.Vertex{-12.3, 23.3})
	cl2.Add(&v.Vertex{1.0, 1.0})
	cl1.Concat(cl2)
	if cl1.Len() != 3 {
		t.Errorf("Cluster list has length", cl1.Len())
	}
	correctBack := &v.Vertex{-12.3, 23.3}
	if !correctBack.Eq(cl1.Back()) {
		println(cl1.Back().Show())
		t.Errorf("Concatenated lists back not same")
	}
}

func TestFron(t *testing.T) {
	l := NewClusterList()
	l.Add(&v.Vertex{.5, -1.5})
	correct := &v.Vertex{.5, -1.5}
	if !correct.Eq(l.Front()) {
		t.Errorf("Cluster list front is wrong")
	}
}