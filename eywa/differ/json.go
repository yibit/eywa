package differ

import (
	"encoding/json"
	"fmt"

	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
)

func JsonDiffer(aString, bString []byte) (string, error) {
	differ := diff.New()
	result, err := differ.Compare(aString, bString)
	if err != nil {
		fmt.Printf("failed to unmarshal file: %s\n", err.Error())
		return "", err
	}

	if result.Modified() {
		var aJson map[string]interface{}
		json.Unmarshal(aString, &aJson)

		config := formatter.AsciiFormatterConfig{
			ShowArrayIndex: true,
			Coloring:       true,
		}

		formatter := formatter.NewAsciiFormatter(aJson, config)
		text, _ := formatter.Format(result)

		return text, nil
	}

	return "", nil
}
