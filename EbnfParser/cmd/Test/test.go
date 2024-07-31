// Code generated by EbnfParser. DO NOT EDIT.
package Test

// TokenType is the type of a token.
type TokenType int

const (
	EtkEOF TokenType = iota 
	
	TtkClParen
	TtkDot
	TtkEqual
	TtkLowercaseId
	TtkNewline
	TtkOpParen
	TtkPipe
	TtkUppercaseId
	
	NtkIdentifier
	NtkOrExpr
	NtkRhs
	NtkRhsCls
	NtkRule
	NtkRuleLine
	NtkSource
	NtkSource1
)

func (t TokenType) String() string {
	return [...]string{
		"End of File",
		// Add here your custom token names.
	}[t]
}

func (t TokenType) GoString() string {
	return [...]string{
		"EtkEOF",
		
		"TtkClParen",
		"TtkDot",
		"TtkEqual",
		"TtkLowercaseId",
		"TtkNewline",
		"TtkOpParen",
		"TtkPipe",
		"TtkUppercaseId",
		
		"NtkIdentifier",
		"NtkOrExpr",
		"NtkRhs",
		"NtkRhsCls",
		"NtkRule",
		"NtkRuleLine",
		"NtkSource",
		"NtkSource1",
	}[t]
}
	
const Grammar string = `[ EtkEOF ] NtkSource1 -> NtkSource : ACCEPT .

NtkOrExpr TtkPipe [ NtkIdentifier ] -> NtkOrExpr : SHIFT .
[ NtkIdentifier ] -> NtkRhs : REDUCE .
[ NtkIdentifier ] TtkPipe NtkIdentifier -> NtkOrExpr : REDUCE .
NtkIdentifier TtkPipe [ NtkIdentifier ] -> NtkOrExpr : SHIFT .

TtkClParen [ NtkOrExpr ] TtkOpParen -> NtkRhs : SHIFT .
[ NtkOrExpr ] TtkPipe NtkIdentifier -> NtkOrExpr : REDUCE .

[ NtkRhs ] -> NtkRhsCls : REDUCE .
NtkRhsCls [ NtkRhs ] -> NtkRhsCls : SHIFT .

TtkDot [ NtkRhsCls ] TtkEqual TtkUppercaseId -> NtkRule : SHIFT .
NtkRuleLine [ NtkRhsCls ] TtkEqual TtkNewline TtkUppercaseId -> NtkRule : SHIFT .
NtkRuleLine [ NtkRhsCls ] TtkPipe TtkNewline -> NtkRuleLine : SHIFT .
[ NtkRhsCls ] NtkRhs -> NtkRhsCls : REDUCE .

[ NtkRule ] -> NtkSource1 : REDUCE .
NtkSource1 TtkNewline [ NtkRule ] -> NtkSource1 : SHIFT .

[ NtkRuleLine ] NtkRhsCls TtkEqual TtkNewline TtkUppercaseId -> NtkRule : REDUCE .
[ NtkRuleLine ] NtkRhsCls TtkPipe TtkNewline -> NtkRuleLine : REDUCE .



EtkEOF [ NtkSource1 ] -> NtkSource : SHIFT .
[ NtkSource1 ] TtkNewline NtkRule -> NtkSource1 : REDUCE .

[ TtkClParen ] NtkOrExpr TtkOpParen -> NtkRhs : REDUCE .

[ TtkDot ] NtkRhsCls TtkEqual TtkUppercaseId -> NtkRule : REDUCE .
[ TtkDot ] TtkNewline -> NtkRuleLine : REDUCE .

TtkDot NtkRhsCls [ TtkEqual ] TtkUppercaseId -> NtkRule : SHIFT .
NtkRuleLine NtkRhsCls [ TtkEqual ] TtkNewline TtkUppercaseId -> NtkRule : SHIFT .

[ TtkLowercaseId ] -> NtkIdentifier : REDUCE .

NtkSource1 [ TtkNewline ] NtkRule -> NtkSource1 : SHIFT .
NtkRuleLine NtkRhsCls TtkEqual [ TtkNewline ] TtkUppercaseId -> NtkRule : SHIFT .
NtkRuleLine NtkRhsCls TtkPipe [ TtkNewline ] -> NtkRuleLine : SHIFT .
TtkDot [ TtkNewline ] -> NtkRuleLine : SHIFT .

TtkClParen NtkOrExpr [ TtkOpParen ] -> NtkRhs : SHIFT .

NtkRuleLine NtkRhsCls [ TtkPipe ] TtkNewline -> NtkRuleLine : SHIFT .
NtkIdentifier [ TtkPipe ] NtkIdentifier -> NtkOrExpr : SHIFT .
NtkOrExpr [ TtkPipe ] NtkIdentifier -> NtkOrExpr : SHIFT .

TtkDot NtkRhsCls TtkEqual [ TtkUppercaseId ] -> NtkRule : SHIFT .
NtkRuleLine NtkRhsCls TtkEqual TtkNewline [ TtkUppercaseId ] -> NtkRule : SHIFT .
[ TtkUppercaseId ] -> NtkIdentifier : REDUCE .`
