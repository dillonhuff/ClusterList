package clusterlist

import v "vertex"

type ClusterList struct {
	length int
	front *vertexNode
	back *vertexNode
}

type vertexNode struct {
	next *vertexNode
	v *v.Vertex
}

func (c *ClusterList) Len() int {
	return c.length
}

func (c *ClusterList) Add(v *v.Vertex) {
	vNode := &vertexNode{c.front, v}
	c.front = vNode
	if (c.Len() == 0) {
		c.back = vNode
	}
	c.length = c.length + 1
}

func (c *ClusterList) Concat(other *ClusterList) {
	c.back.next = other.front
	c.back = other.back
	c.length = c.length + other.length
}

func NewClusterList() *ClusterList {
	return &ClusterList{0, nil, nil}
}