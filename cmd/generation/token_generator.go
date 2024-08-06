package generation

import (
	ggen "github.com/PlayerR9/lib_units/generator"
)

type TokenGen struct {
	PackageName  string
	SpecialEnums []string
	LexerEnums   []string
	ParserEnums  []string
}

func (g *TokenGen) SetPackageName(pkg_name string) {
	g.PackageName = pkg_name
}

var (
	TokenGenerator *ggen.CodeGenerator[*TokenGen]
)

func init() {
	tmp, err := ggen.NewCodeGeneratorFromTemplate[*TokenGen]("", token_templ)
	if err != nil {
		Logger.Fatalf("Error creating code generator: %s", err.Error())
	}

	TokenGenerator = tmp
}

const token_templ string = `// Code generated by SlParser.
package {{ .PackageName }}

// TokenType is the type of a token.
type TokenType int

const (
   {{- range $index, $element := .SpecialEnums }}
	{{ if eq $index 0 }} {{- $element }} TokenType = iota {{ else }} {{- $element }} {{ end }}
	{{- end }}
	{{ range $index, $element := .LexerEnums }}
	{{ $element }}
	{{- end }}
	{{ range $index, $element := .ParserEnums }}
	{{ $element }}
	{{- end }}
)

// String implements the Grammar.TokenTyper interface.
func (t TokenType) String() string {
	return [...]string{
		"End of File",
		// Add here your custom token names.
	}[t]
}

// GoString implements the Grammar.TokenTyper interface.
func (t TokenType) GoString() string {
	return [...]string{
		{{- range $index, $element := .SpecialEnums }}
		"{{ $element }}",
		{{- end }}
		{{ range $index, $element := .LexerEnums }}
		"{{ $element }}",
		{{- end }}
		{{ range $index, $element := .ParserEnums }}
		"{{ $element }}",
		{{- end }}
	}[t]
}
`
