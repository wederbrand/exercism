// Package tree can build nice views for web presentations.
package tree

import (
	"errors"
	"sort"
)

// Record keeps data on a record, it's own id and it's parent.
type Record struct {
	ID     int
	Parent int
}

// Node keeps data on a record ID and it's children.
type Node struct {
	ID       int
	Children []*Node
}

// Build builds the tree.
func Build(records []Record) (*Node, error) {
	// Sort it to make sure we get lower IDs first
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	tree := make(map[int]*Node)

	for i, r := range records {
		if r.ID != i || (r.ID == 0 && r.Parent != 0) || (r.ID != 0 && r.Parent >= r.ID) {
			return nil, errors.New("not consecutive or messed up parent")
		}

		node := &Node{ID: r.ID}
		tree[r.ID] = node

		if r.ID != 0 {
			parent := tree[r.Parent]
			parent.Children = append(parent.Children, node)
		}
	}

	return tree[0], nil
}
