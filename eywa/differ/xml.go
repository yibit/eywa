package differ

import (
	"fmt"

	"github.com/ajankovic/xdiff"
	"github.com/ajankovic/xdiff/parser"
)

func XmlDiffer(aString, bString []byte) (string, error) {
	xParser := parser.New()
	left, err := xParser.ParseBytes(aString)
	if err != nil {
		return "", err
	}

	right, err := xParser.ParseBytes(bString)
	if err != nil {
		return "", err
	}

	deltas, err := xdiff.Compare(left, right)
	if err != nil {
		return "", err
	}

	if len(deltas) == 0 {
		return "", nil
	}

	text := ""
	for _, d := range deltas {
		if d.Operation == xdiff.Update {
			text += fmt.Sprintf("%s('%s'->'%s')\n", d.Operation, d.Subject, d.Object)
			continue
		}
		text += fmt.Sprintf(text, "%s('%s')\n", d.Operation, d.Subject)
	}

	return text, nil
}
