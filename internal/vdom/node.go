// Copyright 2026 Daniel\n// Licensed under the GNU Affero General Public License v3.0.\n// Copying or distributing this file requires compliance with AGPLv3.\n
package vdom

type NodeType int

const (
    ElementNode NodeType = iota
    TextNode
)

type Node struct {
    Type     NodeType        `json:"type"`
    Tag      string          `json:"tag,omitempty"`
    Text     string          `json:"text,omitempty"`
    Props    map[string]string `json:"props,omitempty"`
    Children []Node          `json:"children,omitempty"`
}

func Element(tag string, props map[string]string, children ...Node) Node {
    return Node{
        Type:     ElementNode,
        Tag:      tag,
        Props:    props,
        Children: children,
    }
}

func Text(text string) Node {
    return Node{
        Type: TextNode,
        Text: text,
    }
}
