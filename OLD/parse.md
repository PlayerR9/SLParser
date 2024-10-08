package SlParser

import (
	ebnf "github.com/PlayerR9/SlParser/ebnf"
	gr "github.com/PlayerR9/grammar"
)

// ParseEbnf parses an EBNF file.
//
// Parameters:
//   - data: The data to parse.
//
// Returns:
//   - *ast.Node[prx.NodeType]: The root node of the AST tree.
func ParseEbnf(data []byte) (*ebnf.Node, error) {
	ebnf.Parser.SetDebug(gr.ShowNone)

	root, err := ebnf.Parser.Parse(data)
	if err != nil {
		return root, err
	}

	return root, nil
}
