// Code generated by SlParser.
package test

import (
	"fmt"

	"github.com/PlayerR9/grammar/grammar"
	"github.com/PlayerR9/grammar/parser"
)

var (
	// internal_parser is the parser of the grammar.
	internal_parser *parser.Parser[token_type]
)

func init() {
	decision_func := func(p *parser.Parser[token_type], lookahead *grammar.Token[token_type]) (parser.Actioner, error) {
		top1, ok := p.Pop()
		if !ok {
			return nil, fmt.Errorf("p.stack is empty")
		}

		var act parser.Actioner

		switch top1.Type {
		case etk_EOF:
			// [ etk_EOF ] ttk_TypeId ttk_Newline ttk_ClSquare ttk_ClParen ttk_Number ttk_OpParen ttk_Range ttk_Space ttk_In ttk_Space ttk_VariableId ttk_Space ttk_For ttk_Space ttk_VariableId ttk_Space ttk_AsteriskOp ttk_Space ttk_VariableId ttk_OpSquare ttk_Space ttk_AssignmentOp ttk_Space ttk_TypeId -> ntk_Source : REDUCE .
		case ntk_Source:
		case ttk_AssignmentOp:
			// etk_EOF ttk_TypeId ttk_Newline ttk_ClSquare ttk_ClParen ttk_Number ttk_OpParen ttk_Range ttk_Space ttk_In ttk_Space ttk_VariableId ttk_Space ttk_For ttk_Space ttk_VariableId ttk_Space ttk_AsteriskOp ttk_Space ttk_VariableId ttk_OpSquare ttk_Space [ ttk_AssignmentOp ] ttk_Space ttk_TypeId -> ntk_Source : SHIFT .
		case ttk_AsteriskOp:
			// etk_EOF ttk_TypeId ttk_Newline ttk_ClSquare ttk_ClParen ttk_Number ttk_OpParen ttk_Range ttk_Space ttk_In ttk_Space ttk_VariableId ttk_Space ttk_For ttk_Space ttk_VariableId ttk_Space [ ttk_AsteriskOp ] ttk_Space ttk_VariableId ttk_OpSquare ttk_Space ttk_AssignmentOp ttk_Space ttk_TypeId -> ntk_Source : SHIFT .
		case ttk_ClParen:
			// etk_EOF ttk_TypeId ttk_Newline ttk_ClSquare [ ttk_ClParen ] ttk_Number ttk_OpParen ttk_Range ttk_Space ttk_In ttk_Space ttk_VariableId ttk_Space ttk_For ttk_Space ttk_VariableId ttk_Space ttk_AsteriskOp ttk_Space ttk_VariableId ttk_OpSquare ttk_Space ttk_AssignmentOp ttk_Space ttk_TypeId -> ntk_Source : SHIFT .
		case ttk_ClSquare:
			// etk_EOF ttk_TypeId ttk_Newline [ ttk_ClSquare ] ttk_ClParen ttk_Number ttk_OpParen ttk_Range ttk_Space ttk_In ttk_Space ttk_VariableId ttk_Space ttk_For ttk_Space ttk_VariableId ttk_Space ttk_AsteriskOp ttk_Space ttk_VariableId ttk_OpSquare ttk_Space ttk_AssignmentOp ttk_Space ttk_TypeId -> ntk_Source : SHIFT .
		case ttk_For:
			// etk_EOF ttk_TypeId ttk_Newline ttk_ClSquare ttk_ClParen ttk_Number ttk_OpParen ttk_Range ttk_Space ttk_In ttk_Space ttk_VariableId ttk_Space [ ttk_For ] ttk_Space ttk_VariableId ttk_Space ttk_AsteriskOp ttk_Space ttk_VariableId ttk_OpSquare ttk_Space ttk_AssignmentOp ttk_Space ttk_TypeId -> ntk_Source : SHIFT .
		case ttk_In:
			// etk_EOF ttk_TypeId ttk_Newline ttk_ClSquare ttk_ClParen ttk_Number ttk_OpParen ttk_Range ttk_Space [ ttk_In ] ttk_Space ttk_VariableId ttk_Space ttk_For ttk_Space ttk_VariableId ttk_Space ttk_AsteriskOp ttk_Space ttk_VariableId ttk_OpSquare ttk_Space ttk_AssignmentOp ttk_Space ttk_TypeId -> ntk_Source : SHIFT .
		case ttk_Newline:
			// etk_EOF ttk_TypeId [ ttk_Newline ] ttk_ClSquare ttk_ClParen ttk_Number ttk_OpParen ttk_Range ttk_Space ttk_In ttk_Space ttk_VariableId ttk_Space ttk_For ttk_Space ttk_VariableId ttk_Space ttk_AsteriskOp ttk_Space ttk_VariableId ttk_OpSquare ttk_Space ttk_AssignmentOp ttk_Space ttk_TypeId -> ntk_Source : SHIFT .
		case ttk_Number:
			// etk_EOF ttk_TypeId ttk_Newline ttk_ClSquare ttk_ClParen [ ttk_Number ] ttk_OpParen ttk_Range ttk_Space ttk_In ttk_Space ttk_VariableId ttk_Space ttk_For ttk_Space ttk_VariableId ttk_Space ttk_AsteriskOp ttk_Space ttk_VariableId ttk_OpSquare ttk_Space ttk_AssignmentOp ttk_Space ttk_TypeId -> ntk_Source : SHIFT .
		case ttk_OpParen:
			// etk_EOF ttk_TypeId ttk_Newline ttk_ClSquare ttk_ClParen ttk_Number [ ttk_OpParen ] ttk_Range ttk_Space ttk_In ttk_Space ttk_VariableId ttk_Space ttk_For ttk_Space ttk_VariableId ttk_Space ttk_AsteriskOp ttk_Space ttk_VariableId ttk_OpSquare ttk_Space ttk_AssignmentOp ttk_Space ttk_TypeId -> ntk_Source : SHIFT .
		case ttk_OpSquare:
			// etk_EOF ttk_TypeId ttk_Newline ttk_ClSquare ttk_ClParen ttk_Number ttk_OpParen ttk_Range ttk_Space ttk_In ttk_Space ttk_VariableId ttk_Space ttk_For ttk_Space ttk_VariableId ttk_Space ttk_AsteriskOp ttk_Space ttk_VariableId [ ttk_OpSquare ] ttk_Space ttk_AssignmentOp ttk_Space ttk_TypeId -> ntk_Source : SHIFT .
		case ttk_Range:
			// etk_EOF ttk_TypeId ttk_Newline ttk_ClSquare ttk_ClParen ttk_Number ttk_OpParen [ ttk_Range ] ttk_Space ttk_In ttk_Space ttk_VariableId ttk_Space ttk_For ttk_Space ttk_VariableId ttk_Space ttk_AsteriskOp ttk_Space ttk_VariableId ttk_OpSquare ttk_Space ttk_AssignmentOp ttk_Space ttk_TypeId -> ntk_Source : SHIFT .
		case ttk_Space:
			// etk_EOF ttk_TypeId ttk_Newline ttk_ClSquare ttk_ClParen ttk_Number ttk_OpParen ttk_Range [ ttk_Space ] ttk_In ttk_Space ttk_VariableId ttk_Space ttk_For ttk_Space ttk_VariableId ttk_Space ttk_AsteriskOp ttk_Space ttk_VariableId ttk_OpSquare ttk_Space ttk_AssignmentOp ttk_Space ttk_TypeId -> ntk_Source : SHIFT .
			// etk_EOF ttk_TypeId ttk_Newline ttk_ClSquare ttk_ClParen ttk_Number ttk_OpParen ttk_Range ttk_Space ttk_In [ ttk_Space ] ttk_VariableId ttk_Space ttk_For ttk_Space ttk_VariableId ttk_Space ttk_AsteriskOp ttk_Space ttk_VariableId ttk_OpSquare ttk_Space ttk_AssignmentOp ttk_Space ttk_TypeId -> ntk_Source : SHIFT .
			// etk_EOF ttk_TypeId ttk_Newline ttk_ClSquare ttk_ClParen ttk_Number ttk_OpParen ttk_Range ttk_Space ttk_In ttk_Space ttk_VariableId [ ttk_Space ] ttk_For ttk_Space ttk_VariableId ttk_Space ttk_AsteriskOp ttk_Space ttk_VariableId ttk_OpSquare ttk_Space ttk_AssignmentOp ttk_Space ttk_TypeId -> ntk_Source : SHIFT .
			// etk_EOF ttk_TypeId ttk_Newline ttk_ClSquare ttk_ClParen ttk_Number ttk_OpParen ttk_Range ttk_Space ttk_In ttk_Space ttk_VariableId ttk_Space ttk_For [ ttk_Space ] ttk_VariableId ttk_Space ttk_AsteriskOp ttk_Space ttk_VariableId ttk_OpSquare ttk_Space ttk_AssignmentOp ttk_Space ttk_TypeId -> ntk_Source : SHIFT .
			// etk_EOF ttk_TypeId ttk_Newline ttk_ClSquare ttk_ClParen ttk_Number ttk_OpParen ttk_Range ttk_Space ttk_In ttk_Space ttk_VariableId ttk_Space ttk_For ttk_Space ttk_VariableId [ ttk_Space ] ttk_AsteriskOp ttk_Space ttk_VariableId ttk_OpSquare ttk_Space ttk_AssignmentOp ttk_Space ttk_TypeId -> ntk_Source : SHIFT .
			// etk_EOF ttk_TypeId ttk_Newline ttk_ClSquare ttk_ClParen ttk_Number ttk_OpParen ttk_Range ttk_Space ttk_In ttk_Space ttk_VariableId ttk_Space ttk_For ttk_Space ttk_VariableId ttk_Space ttk_AsteriskOp [ ttk_Space ] ttk_VariableId ttk_OpSquare ttk_Space ttk_AssignmentOp ttk_Space ttk_TypeId -> ntk_Source : SHIFT .
			// etk_EOF ttk_TypeId ttk_Newline ttk_ClSquare ttk_ClParen ttk_Number ttk_OpParen ttk_Range ttk_Space ttk_In ttk_Space ttk_VariableId ttk_Space ttk_For ttk_Space ttk_VariableId ttk_Space ttk_AsteriskOp ttk_Space ttk_VariableId ttk_OpSquare [ ttk_Space ] ttk_AssignmentOp ttk_Space ttk_TypeId -> ntk_Source : SHIFT .
			// etk_EOF ttk_TypeId ttk_Newline ttk_ClSquare ttk_ClParen ttk_Number ttk_OpParen ttk_Range ttk_Space ttk_In ttk_Space ttk_VariableId ttk_Space ttk_For ttk_Space ttk_VariableId ttk_Space ttk_AsteriskOp ttk_Space ttk_VariableId ttk_OpSquare ttk_Space ttk_AssignmentOp [ ttk_Space ] ttk_TypeId -> ntk_Source : SHIFT .
		case ttk_TypeId:
			// etk_EOF [ ttk_TypeId ] ttk_Newline ttk_ClSquare ttk_ClParen ttk_Number ttk_OpParen ttk_Range ttk_Space ttk_In ttk_Space ttk_VariableId ttk_Space ttk_For ttk_Space ttk_VariableId ttk_Space ttk_AsteriskOp ttk_Space ttk_VariableId ttk_OpSquare ttk_Space ttk_AssignmentOp ttk_Space ttk_TypeId -> ntk_Source : SHIFT .
			// etk_EOF ttk_TypeId ttk_Newline ttk_ClSquare ttk_ClParen ttk_Number ttk_OpParen ttk_Range ttk_Space ttk_In ttk_Space ttk_VariableId ttk_Space ttk_For ttk_Space ttk_VariableId ttk_Space ttk_AsteriskOp ttk_Space ttk_VariableId ttk_OpSquare ttk_Space ttk_AssignmentOp ttk_Space [ ttk_TypeId ] -> ntk_Source : SHIFT .
		case ttk_VariableId:
			// etk_EOF ttk_TypeId ttk_Newline ttk_ClSquare ttk_ClParen ttk_Number ttk_OpParen ttk_Range ttk_Space ttk_In ttk_Space [ ttk_VariableId ] ttk_Space ttk_For ttk_Space ttk_VariableId ttk_Space ttk_AsteriskOp ttk_Space ttk_VariableId ttk_OpSquare ttk_Space ttk_AssignmentOp ttk_Space ttk_TypeId -> ntk_Source : SHIFT .
			// etk_EOF ttk_TypeId ttk_Newline ttk_ClSquare ttk_ClParen ttk_Number ttk_OpParen ttk_Range ttk_Space ttk_In ttk_Space ttk_VariableId ttk_Space ttk_For ttk_Space [ ttk_VariableId ] ttk_Space ttk_AsteriskOp ttk_Space ttk_VariableId ttk_OpSquare ttk_Space ttk_AssignmentOp ttk_Space ttk_TypeId -> ntk_Source : SHIFT .
			// etk_EOF ttk_TypeId ttk_Newline ttk_ClSquare ttk_ClParen ttk_Number ttk_OpParen ttk_Range ttk_Space ttk_In ttk_Space ttk_VariableId ttk_Space ttk_For ttk_Space ttk_VariableId ttk_Space ttk_AsteriskOp ttk_Space [ ttk_VariableId ] ttk_OpSquare ttk_Space ttk_AssignmentOp ttk_Space ttk_TypeId -> ntk_Source : SHIFT .
		default:
			return nil, fmt.Errorf("unexpected token: %s", top1.String())
		}

		return act, nil
	}

	internal_parser = parser.NewParser(decision_func)
}