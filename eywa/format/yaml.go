package format

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func Yaml(file string, indent int) string {
	r, err := os.Open(file)
	if err != nil {
		return err.Error()
	}

	defer r.Close()

	var out bytes.Buffer
	if err := FormatStream(r, &out, indent); err != nil {
		return err.Error()
	}

	return out.String()
}

func FormatStream(r io.Reader, out io.Writer, indent int) error {
	d := yaml.NewDecoder(r)

	in := yaml.Node{
		Kind:        0,
		Style:       0,
		Tag:         "",
		Value:       "",
		Anchor:      "",
		Alias:       &yaml.Node{},
		Content:     []*yaml.Node{},
		HeadComment: "",
		LineComment: "",
		FootComment: "",
		Line:        0,
		Column:      0,
	}

	err := d.Decode(&in)
	for err == nil {
		e := yaml.NewEncoder(out)
		e.SetIndent(indent)
		if err := e.Encode(&in); err != nil {
			log.Fatal(err)
		}
		e.Close()

		if err = d.Decode(&in); err == nil {
			fmt.Fprintln(out, "---")
		}
	}

	if err == io.EOF {
		return nil
	}

	return err
}
