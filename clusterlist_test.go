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
	cl2.Add(&v.Vertex{1.342, 908.1})
	cl1.Concat(cl2)
	if cList.Len() != 3 {
		t.Errorf("Cluster list has length", cList.Len())
	}
}