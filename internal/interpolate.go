package interpolate

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"os"
	"text/template"
)

func Execute(input map[string]string, tmplString string) (err error) {
	templateString, err := template.New("stdin").Funcs(sprig.TxtFuncMap()).Parse(tmplString)
	if err != nil {
		return err
	}

	buffer := new(bytes.Buffer)

	err = templateString.Execute(buffer, input)
	if err != nil {
		return err
	}

	os.Stdout.Write(buffer.Bytes())

	return
}
