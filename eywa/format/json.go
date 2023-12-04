package format

import (
	"encoding/json"

	"github.com/TylerBrock/colorjson"
	"github.com/fatih/color"
)

func Json(data string) string {
	var obj map[string]interface{}

	json.Unmarshal([]byte(data), &obj)

	// Make a custom formatter with indent set
	fmt := colorjson.NewFormatter()
	fmt.Indent = 2
	fmt.KeyColor = color.New(color.FgBlue)

	// Marshall the Colorized JSON
	s, err := fmt.Marshal(obj)
	if err != nil {
		return err.Error()
	}

	return string(s)
}
