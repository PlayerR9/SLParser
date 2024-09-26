package internal

import (
	kdd "github.com/PlayerR9/SlParser/kdd"
	gers "github.com/PlayerR9/go-errors"
	"github.com/PlayerR9/go-generator"
)

// ASTGen is a generator for the AST.
type ASTGen struct {
	// PackageName is the name of the package.
	PackageName string

	// Ast is the list of candidates for the node types in the AST.
	Ast []string
}

// SetPackageName implements the generator.PackageNameSetter interface.
func (g *ASTGen) SetPackageName(pkg_name string) {
	if g == nil {
		return
	}

	g.PackageName = pkg_name
}

// NewASTGen creates a new ASTGen with the given tokens.
//
// Parameters:
//   - nodes: The list of tokens.
//
// Returns:
//   - *ASTGen: the ASTGen. Never returns nil.
func NewASTGen(table map[*kdd.Node]*Info) *ASTGen {
	candidates := CandidatesForAst(table)

	gen := &ASTGen{
		Ast: candidates,
	}

	return gen
}

var (
	// ASTGenerator is the generator for the AST.
	ASTGenerator *generator.CodeGenerator[*ASTGen]
)

func init() {
	var err error

	ASTGenerator, err = generator.NewCodeGeneratorFromTemplate[*ASTGen]("ast", ast_templ)
	gers.AssertErr(err, "generator.NewCodeGeneratorFromTemplate[*ASTGen](%q, ast_templ)", "ast")
}

// ast_templ is the template for the AST.
const ast_templ = `// Code generated by SlParser. Do not edit.
package {{ .PackageName }}

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"iter"
	"log"
	"os"

	sl "github.com/PlayerR9/SlParser"
	"github.com/PlayerR9/SlParser/ast"
	"github.com/PlayerR9/SlParser/grammar"
	"github.com/PlayerR9/SlParser/lexer"
	"github.com/PlayerR9/SlParser/parser"
)

type NodeType int

const (
	/*InvalidNode represents an invalid node.
	Node[InvalidNode]
	*/
	InvalidNode NodeType = iota - 1 // Invalid {{ range $index, $value := .Ast }}

	/*{{ $value }}Node is [...].
	Node[{{ $value }}Node]
	*/
	{{ $value }}Node // {{ $value }}
	{{- end }}
)

var (
	ast_maker *ast.AstMaker[*Node, internal.TokenType]
)
	
func init() {
	builder := ast.NewBuilder[*Node, internal.TokenType]()

	// TODO: Add here your own custom rules...
	{{ range $index, $value := .Ast }}
	builder.Register(internal.Nt{{ $value }}, func(tk *grammar.ParseTree[internal.TokenType]) (*Node, error) {
		children := tk.GetChildren()
		if len(children) == 0 {
			return nil, errors.New("expected at least one child")
		}

		// TODO: Complete this function...

		node := NewNode(tk.Pos(), {{ $value }}Node, "")
		return node, nil
	})
	{{- end }}

	ast_maker = builder.Build()
}`
