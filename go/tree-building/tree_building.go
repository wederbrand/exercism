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
	if len(records) == 0 {
		return nil, nil
	}

	// Sort it to make sure we get lower IDs first
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	allNodes := make(map[int]*Node)
	rootNode := &Node{}
	
	for _, record := range records {
		_, found := allNodes[record.ID]
		if found {
			return nil, errors.New("duplicate node")
		}

		thisNode := &Node{ID: record.ID}
		allNodes[record.ID] = thisNode

		if record.ID == 0 {
			if record.Parent != 0 {
				return nil, errors.New("root node can't have a parent")
			}
			rootNode = thisNode
			continue
		}

		_, found = allNodes[record.ID-1]
		if !found {
			return nil, errors.New("non-continuous")
		}

		if record.Parent >= record.ID {
			return nil, errors.New("parent can't be equal or bigger than ID")
		}

		parent, found := allNodes[record.Parent]
		if !found {
			return nil, errors.New("parent not found")
		}

		parent.Children = append(parent.Children, thisNode)
	}

	return rootNode, nil
}
