package format

import (
	"encoding/json"

	"github.com/TylerBrock/colorjson"
	"github.com/fatih/color"
)

func Json(data string) string {
	var obj map[string]interface{}

	if err := json.Unmarshal([]byte(data), &obj); err != nil {
		return err.Error()
	}

	formatter := colorjson.NewFormatter()
	formatter.Indent = 2
	formatter.KeyColor = color.New(color.FgBlue)
	formatter.NullColor = color.New(color.FgRed)

	out, err := formatter.Marshal(obj)
	if err != nil {
		return err.Error()
	}

	return string(out)
}
