package templating

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"unicode/utf8"

	"github.com/spf13/pflag"
)

// TemplateFlags stores command line formatting flags and provides for their
// registration and handling.
type TemplateFlags struct {
	// template stores the value of the --template flag.
	template string
	// templateFile stores the value of the --template-file flag.
	templateFile string
}

// Register registers the flags into the specified flag set.
func (f *TemplateFlags) Register(flags *pflag.FlagSet) {
	flags.StringVar(&f.template, "template", "", "Specify an output template")
	flags.StringVar(&f.templateFile, "template-file", "", "Specify a file containing an output template")

	// If the executable is built as balena-go, display prettier output, otherwise show the output as JSON
	if filepath.Base(os.Args[0]) != "balena-go" {
		f.template = "{{ json . }}"
	}
}

// LoadTemplate loads the template specified by the flags. If no template has
// been specified, then it returns nil with no error. Template literals
// specified via the command line will have a trailing newline added.
func (f *TemplateFlags) LoadTemplate() (*template.Template, error) {
	// Figure out if there's a template to be processed. If not, then no valid
	// template has been specified and we can just return. If a template literal
	// was provided directly on the command line, then add a trailing newline to
	// make typical command line usage more friendly.
	var literal string
	if f.template != "" {
		literal = f.template + "\n"
	} else if f.templateFile != "" {
		if l, err := os.ReadFile(f.templateFile); err != nil {
			return nil, fmt.Errorf("unable to load template: %w", err)
		} else if !utf8.Valid(l) {
			return nil, errors.New("template file is not UTF-8 encoded")
		} else {
			literal = string(l)
		}
	} else {
		return nil, nil
	}

	// Create the template and register built-in functions.
	result := template.New("")
	result.Funcs(builtins)

	// Parse the template literal.
	return result.Parse(literal)
}
