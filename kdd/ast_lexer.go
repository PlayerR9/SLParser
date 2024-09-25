package kdd

import (
	"fmt"

	"github.com/PlayerR9/SlParser/ast"
	"github.com/PlayerR9/SlParser/grammar"
	internal "github.com/PlayerR9/SlParser/kdd/internal"
	gers "github.com/PlayerR9/go-errors"
)

//go:generate stringer -type=NodeType -linecomment

type NodeType int

const (
	/*InvalidNode represents an invalid node.
	Node[InvalidNode]
	*/
	InvalidNode NodeType = iota - 1 // Invalid

	/*RhsNode represents the terminal symbol.
	Node[RhsNode (<id>)]
	*/
	RhsNode // Rhs

	/*RuleNode represents a single rule.
	Node[RuleNode]
	 ├── RhsNode (<id>) // This is the LHS of the rule.
	 ├── RhsNode (<id>) // This is the RHS of the rule.
	 └── ...
	*/
	RuleNode // Rule

	/*SourceNode is the collection of all rules in the grammar.
	Node[SourceNode]
	 ├── Node[RuleNode]
	 └── ...
	*/
	SourceNode // Source
)

// rule : LOWERCASE_ID COLON rhs+ SEMICOLON ;

var (
	ast_maker *ast.AstMaker[*Node, internal.TokenType]
)

func init() {
	builder := ast.NewBuilder[*Node, internal.TokenType]()

	// TODO: Add here your own custom rules...

	// rhs : UPPERCASE_ID ;
	// rhs : LOWERCASE_ID ;
	builder.Register(internal.NtRhs, func(tk *grammar.ParseTree[internal.TokenType]) (*Node, error) {
		children := tk.GetChildren()

		if len(children) != 1 {
			return nil, fmt.Errorf("expected one child, got %d instead", len(children))
		}

		type_ := children[0].Type()
		gers.AssertNotNil(type_, "type_")

		var is_terminal bool

		switch type_ {
		case internal.TtUppercaseId:
			is_terminal = true
		case internal.TtLowercaseId:
			is_terminal = true
		default:
			return nil, fmt.Errorf("expected UPPERCASE_ID or LOWERCASE_ID, got %s instead", type_.String())
		}

		node := NewNode(tk.Pos(), RhsNode, children[0].Data())
		node.IsTerminal = is_terminal
		return node, nil
	})

	// rule1 : rhs ;
	// rule1 : rhs rule1 ;
	rule1 := func(children []*grammar.ParseTree[internal.TokenType]) (*Node, error) {
		if len(children) != 1 {
			return nil, fmt.Errorf("expected one child, got %d instead", len(children))
		}

		node, err := ast_maker.Convert(children[0])
		if err != nil {
			return nil, err
		} else if node.Type != RhsNode {
			return nil, fmt.Errorf("expected RhsNode, got %s instead", node.Type.String())
		}

		return node, nil
	}

	builder.Register(internal.NtRule, func(tk *grammar.ParseTree[internal.TokenType]) (*Node, error) {
		children := tk.GetChildren()

		// rule : LOWERCASE_ID COLON rule1 SEMICOLON ;
		ast.CheckType(children, 0, internal.TtLowercaseId)

		lhs := NewNode(children[0].Pos(), RhsNode, children[0].Data())
		lhs.IsTerminal = true

		node := NewNode(tk.Pos(), RuleNode, "")
		node.AddChild(lhs)

		ast.CheckType(children, 1, internal.TtColon)
		ast.CheckType(children, 3, internal.TtSemicolon)

		sub_children, err := ast.LhsToAst(2, children, internal.NtRule1, rule1)
		if err != nil {
			return nil, err
		}

		node.AddChildren(sub_children)

		return node, nil
	})

	source1 := func(children []*grammar.ParseTree[internal.TokenType]) (*Node, error) {
		var node *Node

		switch len(children) {
		case 1:
			// source1 : rule ;

			var err error

			node, err = ast_maker.Convert(children[0])
			if err != nil {
				return nil, err
			} else if node.Type != RuleNode {
				return nil, fmt.Errorf("expected RuleNode, got %s instead", node.Type.String())
			}
		case 2:
			// source1 : rule NEWLINE source1 ;

			ast.CheckType(children, 1, internal.TtNewline)

			var err error

			node, err = ast_maker.Convert(children[0])
			if err != nil {
				return nil, err
			} else if node.Type != RuleNode {
				return nil, fmt.Errorf("expected RuleNode, got %s instead", node.Type.String())
			}
		default:
			return nil, fmt.Errorf("expected one or two children, got %d instead", len(children))
		}

		return node, nil
	}

	builder.Register(internal.NtSource, func(tk *grammar.ParseTree[internal.TokenType]) (*Node, error) {
		children := tk.GetChildren()

		var node *Node

		switch len(children) {
		case 2:
			// source : rule EOF ;

			ast.CheckType(children, 1, internal.EtEOF)

			sub_node, err := ast_maker.Convert(children[0])
			if err != nil {
				return nil, err
			}

			node = NewNode(tk.Pos(), SourceNode, "")
			node.AddChild(sub_node)
		case 4:
			// source : rule NEWLINE source1 EOF ;

			ast.CheckType(children, 1, internal.TtNewline)
			ast.CheckType(children, 3, internal.EtEOF)

			sub_node, err := ast_maker.Convert(children[0])
			if err != nil {
				return nil, err
			} else if sub_node.Type != RuleNode {
				return nil, fmt.Errorf("expected RuleNode, got %s instead", sub_node.Type.String())
			}

			node = NewNode(tk.Pos(), SourceNode, "")
			node.AddChild(sub_node)

			tmp, err := ast.LhsToAst(2, children, internal.NtSource1, source1)
			if err != nil {
				return nil, err
			}

			node.AddChildren(tmp)
		default:
			return nil, fmt.Errorf("expected 2 or 4 children, got %d instead", len(children))
		}

		return node, nil
	})

	ast_maker = builder.Build()
}
