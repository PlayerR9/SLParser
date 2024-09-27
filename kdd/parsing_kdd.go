package kdd

import (
	"bytes"
	"fmt"
	"io"
	"iter"
	"log"
	"os"

	sl "github.com/PlayerR9/SlParser"
	"github.com/PlayerR9/SlParser/ast"
	"github.com/PlayerR9/SlParser/lexer"
	"github.com/PlayerR9/SlParser/parser"
	gers "github.com/PlayerR9/go-errors"
)

// DebugMode is the debug mode.
type DebugMode int

const (
	// ShowNone is the debug mode that does not show anything.
	ShowNone DebugMode = 0

	// ShowTokens is the debug mode that shows the list of tokens.
	ShowTokens DebugMode = 1

	// ShowForest is the debug mode that shows the forest.
	ShowForest DebugMode = 2

	// ShowAST is the debug mode that shows the AST.
	ShowAST DebugMode = 4

	// ShowAll is the debug mode that shows everything.
	ShowAll DebugMode = ShowTokens | ShowForest | ShowAST
)

// Parsing is the parser.
type Parsing struct {
	// debug_mode is the debug mode.
	debug_mode DebugMode

	// debugger is the debugger.
	debugger *log.Logger

	// lexer is the lexer.
	lexer *lexer.Lexer[TokenType]

	// parser is the parser.
	parser *parser.Parser[TokenType]

	// ast is the AST maker.
	ast ast.AstMaker[*Node, TokenType]
}

// NewParser creates a new parser.
//
// Returns:
//   - *Parsing: The new parser. Never returns nil.
func NewParser() *Parsing {
	return &Parsing{
		debug_mode: ShowNone,
		debugger:   log.New(os.Stdout, "[PARSER]: ", log.LstdFlags),
		lexer:      Lexer,
		parser:     Parser,
		ast:        ast_maker,
	}
}

// SetMode sets the debug mode.
//
// Parameters:
//   - mode: The debug mode.
//
// Does nothing if the receiver is nil.
func (p *Parsing) SetMode(mode DebugMode) {
	if p == nil {
		return
	}

	p.debug_mode = mode
}

// SetDebugger sets the debugger.
//
// Parameters:
//   - debugger: The debugger.
//
// Does nothing if the receiver is nil.
//
// Sets the debugger to 'log.New(io.Discard, "", 0)' if 'debugger' is nil.
func (p *Parsing) SetDebugger(debugger *log.Logger) {
	if p == nil {
		return
	}

	if debugger == nil {
		p.debugger = log.New(io.Discard, "", 0)
	} else {
		p.debugger = debugger
	}
}

// write is a helper function that writes a data to a writer.
//
// Does nothing if the data is empty.
//
// Parameters:
//   - w: the writer to write to.
//   - data: the data to write.
//
// Returns:
//   - error: the error that occurred.
//
// Errors:
//   - io.ErrShortWrite: if the writer is nil or the data could not be written fully.
//   - any other error returned by the writer.
func write(w io.Writer, data []byte) error {
	if len(data) == 0 {
		return nil
	} else if w == nil {
		return io.ErrShortWrite
	}

	n, err := w.Write(data)
	if err != nil {
		return err
	} else if n != len(data) {
		return io.ErrShortWrite
	}

	return nil
}

// LogPrint acts as DebugPrint but with a logger instead.
//
// Parameters:
//   - l: the logger to write to.
//   - title: the title of the debug message. If empty, no title is printed.
//   - lines: the lines of the debug message. If nil, no lines are printed.
//
// Returns:
//   - error: the error that occurred.
//
// Errors:
//   - io.ErrShortWrite: if the writer is nil or the lines could not be written fully.
//   - any other error returned by the logger.
func (p *Parsing) LogPrint(title string, lines iter.Seq[string]) error {
	var buff bytes.Buffer

	if title != "" {
		p.debugger.Println(title)
		_, _ = buff.WriteRune('\n')
	}

	if lines != nil {
		for line := range lines {
			_, _ = buff.WriteString(line)
			_, _ = buff.WriteRune('\n')
		}

		_, _ = buff.WriteRune('\n')
	}

	w := p.debugger.Writer()
	err := write(w, buff.Bytes())
	return err
}

// Full is the full parsing function.
//
// Parameters:
//   - data: The data to parse.
//
// Returns:
//   - *Node: The parsed node. Never returns nil.
//   - error: The error. Never returns nil.
func (p Parsing) Full(data []byte) (*Node, error) {
	defer p.lexer.Reset()
	tokens, err := sl.Lex(p.lexer, data)

	// DEBUG: Print the list of tokens.
	if p.debug_mode&ShowTokens != 0 {
		err := p.LogPrint("Here's the list of tokens:", func(yield func(string) bool) {
			for _, tk := range tokens {
				if !yield(tk.String()) {
					return
				}
			}
		})
		if err != nil {
			panic(err)
		}
	}

	if err != nil {
		exit_code, err := sl.DisplayErr(os.Stdout, data, err)
		gers.AssertErr(err, "DisplayErr(os.Stdout, data, err)")

		os.Exit(exit_code + 1)
	}

	defer p.parser.Reset()

	p.parser.SetTokens(tokens)

	var node *Node
	var last_error error

	for node == nil {
		forest, err := p.parser.Parse()
		if err != nil {
			if last_error == nil {
				last_error = err
			}

			break
		} else if len(forest) == 0 {
			break
		}

		// DEBUG: Print the forest.
		if p.debug_mode&ShowForest != 0 {
			err := p.LogPrint("Here's the forest:", func(yield func(string) bool) {
				for _, tree := range forest {
					if !yield(tree.String()) {
						return
					}
				}
			})
			if err != nil {
				panic(err)
			}
		}

		if len(forest) != 1 {
			last_error = fmt.Errorf("expected one forest, got %d instead", len(forest))

			continue
		}

		node, err = ast_maker.Convert(forest[0])
		if err != nil {
			last_error = err

			continue
		}
	}

	if node == nil {
		return nil, last_error
	}

	if p.debug_mode&ShowAST != 0 {
		err := p.LogPrint("Here's the AST:", func(yield func(string) bool) {
			_ = yield(PrintAst(node))
		})
		if err != nil {
			panic(err)
		}
	}

	return node, nil
}
