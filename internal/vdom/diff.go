// Copyright 2026 Daniel\n// Licensed under the GNU Affero General Public License v3.0.\n// Copying or distributing this file requires compliance with AGPLv3.\n
package vdom

type PatchOp string

const (
	OpReplaceNode     PatchOp = "replace"
	OpInsertNode      PatchOp = "insert"
	OpRemoveNode      PatchOp = "remove"
	OpSetAttribute    PatchOp = "setAttr"
	OpRemoveAttribute PatchOp = "removeAttr"
	OpSetText         PatchOp = "setText"
)

type Patch struct {
	Op    PatchOp     `json:"op"`
	Path  []int       `json:"path"`
	Name  string      `json:"name,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Node  *Node       `json:"node,omitempty"`
}

func Diff(old, new Node) []Patch {
	return diffNode(old, new, []int{})
}

func diffNode(old, new Node, path []int) []Patch {
	var patches []Patch

	if old.Type != new.Type || old.Tag != new.Tag {
		patches = append(patches, Patch{
			Op:   OpReplaceNode,
			Path: append([]int{}, path...), // garantiert nicht nil
			Node: &new,
		})
		return patches
	}

	if new.Type == TextNode && old.Text != new.Text {
		patches = append(patches, Patch{
			Op:    OpSetText,
			Path:  path,
			Value: new.Text,
		})
	}

	patches = append(patches, diffProps(old, new, path)...)
	patches = append(patches, diffChildren(old, new, path)...)

	return patches
}

func diffProps(old, new Node, path []int) []Patch {
	var patches []Patch

	if old.Props == nil {
		old.Props = map[string]string{}
	}
	if new.Props == nil {
		new.Props = map[string]string{}
	}

	for k, v := range new.Props {
		if old.Props[k] != v {
			patches = append(patches, Patch{
				Op:    OpSetAttribute,
				Path:  path,
				Name:  k,
				Value: v,
			})
		}
	}

	for k := range old.Props {
		if _, ok := new.Props[k]; !ok {
			patches = append(patches, Patch{
				Op:   OpRemoveAttribute,
				Path: path,
				Name: k,
			})
		}
	}

	return patches
}

func diffChildren(old, new Node, path []int) []Patch {
	var patches []Patch

	max := len(old.Children)
	if len(new.Children) > max {
		max = len(new.Children)
	}

	for i := 0; i < max; i++ {
		childPath := append(append([]int(nil), path...), i)

		switch {
		case i >= len(old.Children):
			patches = append(patches, Patch{
				Op:   OpInsertNode,
				Path: childPath,
				Node: &new.Children[i],
			})
		case i >= len(new.Children):
			patches = append(patches, Patch{
				Op:   OpRemoveNode,
				Path: childPath,
			})
		default:
			patches = append(patches, diffNode(old.Children[i], new.Children[i], childPath)...)
		}
	}

	return patches
}
