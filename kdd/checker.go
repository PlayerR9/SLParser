package kdd

import (
	"fmt"

	ast "github.com/PlayerR9/SlParser/ast"
	"github.com/dustin/go-humanize"
)

var (
	// What Counts as a Valid Node?
	//
	// Overall Rules:
	//  1. All children must be valid.
	//  2. A node cannot be nil.
	//
	// SourceNode:
	//  1. Must have at least one children.
	//  2. All children must be of type RuleNode.
	//
	// RuleNode:
	//  1. Must have at least two children. (The first is the LHS while the rest are the RHSs).
	//  2. All children must be of type RhsNode.
	//
	// RhsNode:
	//  1. No children are expected.
	//  2. Data must not be empty.
	CheckAST ast.CheckASTWithLimit[*Node]
)

func init() {
	table := make(map[NodeType]ast.CheckNodeFn[*Node])

	table[SourceNode] = func(node *Node) error {
		// 1. All children must be rule nodes.
		// 2. At least one children is expected.

		if node.FirstChild == nil {
			return fmt.Errorf("at least one rule is expected")
		}

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == RuleNode {
				continue
			}

			return fmt.Errorf("expected %s child to be of type %q, got %q instead",
				humanize.Ordinal(1), RuleNode.String(), c.Type.String(),
			)
		}

		return nil
	}

	table[RuleNode] = func(node *Node) error {
		// 1. All children must be rhs nodes.
		// 2. At least two children are expected.

		if node.FirstChild == nil {
			return fmt.Errorf("missing LHS node")
		} else if node.FirstChild == node.LastChild {
			return fmt.Errorf("missing RHS nodes")
		}

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == RhsNode {
				continue
			}

			return fmt.Errorf("expected %s child to be of type %q, got %q instead",
				humanize.Ordinal(1), RhsNode.String(), c.Type.String(),
			)
		}

		return nil
	}

	table[RhsNode] = func(node *Node) error {
		// 1. No children are expected.
		// 2. Data must not be empty.

		if node.Data == "" {
			return fmt.Errorf("missing identifier")
		}

		if node.FirstChild != nil {
			return fmt.Errorf("expected no children, got %d instead", len(node.GetChildren()))
		}

		return nil
	}

	CheckAST = ast.MakeCheckFn(table)
}
