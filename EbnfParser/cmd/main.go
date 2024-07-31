package main

import (
	"flag"
	"log"
	"os"
	"text/template"

	prx "github.com/PlayerR9/EbnfParser/cmd/Parser"
	pkg "github.com/PlayerR9/EbnfParser/cmd/pkg"
	ggen "github.com/PlayerR9/lib_units/generator"
)

var (
	// t is the template used to generate the output.
	t *template.Template

	// logger is the logger used to log messages.
	logger *log.Logger
)

func init() {
	t = template.Must(template.New("").Parse(templ))

	logger = ggen.InitLogger("ebnf parser")
}

var (
	// InputFileFlag is the flag used to specify the input file.
	InputFileFlag *string
)

func init() {
	ggen.SetOutputFlag("<dir>.go", true)

	InputFileFlag = flag.String("i", "", "The input file to parse. This flag is required.")
}

type Gen struct {
	PackageName  string
	SpecialEnums []string
	LexerEnums   []string
	ParserEnums  []string
	Rules        string
}

func (g Gen) SetPackageName(pkg_name string) ggen.Generater {
	g.PackageName = pkg_name

	return g
}

func main() {
	err := ggen.ParseFlags()
	if err != nil {
		logger.Fatalf("Error parsing flags: %s", err.Error())
	}

	output_loc, err := ggen.FixOutputLoc("test.go", "")
	if err != nil {
		logger.Fatalf("Error fixing output location: %s", err.Error())
	}

	if *InputFileFlag == "" {
		logger.Fatalf("Input file is required")
	}

	data, err := os.ReadFile(*InputFileFlag)
	if err != nil {
		logger.Fatalf("Error reading file: %s", err.Error())
	}

	root, err := prx.Parse(data)
	if err != nil {
		logger.Fatalf("Error parsing file: %s", err.Error())
	}

	ee_data, err := prx.ExtractEnums.Apply(root)
	if err != nil {
		logger.Fatalf("Error extracting enums: %s", err.Error())
	}

	g := Gen{
		SpecialEnums: ee_data.GetSpecialEnums(),
		LexerEnums:   ee_data.GetLexerEnums(),
		ParserEnums:  ee_data.GetParserEnums(),
	}

	_, err = prx.RenameNodes.Apply(root)
	if err != nil {
		logger.Fatalf("Error renaming nodes: %s", err.Error())
	}

	err = ggen.Generate(output_loc, g, t,
		func(g *Gen) error {
			rules, err := pkg.ExtractRules(root)
			if err != nil {
				return err
			}

			g.Rules = pkg.StringifyRules(rules)

			return nil
		},
	)
	if err != nil {
		logger.Fatalf("Error generating file: %s", err.Error())
	}

	logger.Printf("Successfully generated file: %q", output_loc)
}

const templ string = `// Code generated by EbnfParser. DO NOT EDIT.
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

func (t TokenType) String() string {
	return [...]string{
		"End of File",
		// Add here your custom token names.
	}[t]
}

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
	
const Grammar string = ` + "`" + `{{ .Rules }}` + "`" + `
`
