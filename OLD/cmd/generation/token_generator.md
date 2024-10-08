package generation

import (
	"github.com/PlayerR9/SlParser/OLD/cmd/pkg"
	uslc "github.com/PlayerR9/go-commons/cmp"
	ggen "github.com/PlayerR9/go-generator/generator"
)

type TokenGen struct {
	PackageName  string
	Data         *pkg.EnumExtractor
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

	tmp.AddDoFunc(func(a *TokenGen) error {
		a.SpecialEnums = a.Data.GetSpecialEnums()

		a.SpecialEnums = uslc.DeleteElem(a.SpecialEnums, "etk_EOF")

		return nil
	})

	tmp.AddDoFunc(func(a *TokenGen) error {
		a.LexerEnums = a.Data.GetLexerEnums()

		return nil
	})

	tmp.AddDoFunc(func(a *TokenGen) error {
		a.ParserEnums = a.Data.GetParserEnums()

		return nil
	})

	TokenGenerator = tmp
}

const token_templ string = `// Code generated by SlParser.
package {{ .PackageName }}

// token_type is the type of a token.
type token_type int

const (
	etk_EOF token_type = iota
   {{- range $index, $element := .SpecialEnums }}
	{{ $element }}
	{{- end }}
	{{ range $index, $element := .LexerEnums }}
	{{ $element }}
	{{- end }}
	{{ range $index, $element := .ParserEnums }}
	{{ $element }}
	{{- end }}
)

// String implements the Grammar.TokenTyper interface.
func (t token_type) String() string {
	return [...]string{
		"End of File",
		// Add here your custom token names.
	}[t]
}

// GoString implements the Grammar.TokenTyper interface.
func (t token_type) GoString() string {
	return [...]string{
		"etk_EOF",
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
}`
