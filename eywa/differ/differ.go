package differ

import (
	"github.com/charmbracelet/log"

	diffPatch "github.com/sergi/go-diff/diffmatchpatch"
)

func Diff(dataType string, aData, bData []byte) bool {
	switch dataType {
	case "json":
		x, err := JsonDiffer(aData, bData)
		if err != nil {
			log.Errorf("%s", err.Error())
			return false
		}
		if x != "" {
			log.Errorf("\n%s", x)
			return false
		}
	case "xml":
		x, err := XmlDiffer(aData, bData)
		if err != nil {
			log.Errorf("%s", err.Error())
			return false
		}
		if x != "" {
			log.Errorf("\n%s", x)
			return false
		}
	default:
		if string(aData) == string(bData) {
			return true
		} else {
			dmp := diffPatch.New()
			diffs := dmp.DiffMain(string(aData), string(bData), true)
			if len(diffs) > 1 {
				log.Errorf("\n%s", dmp.DiffPrettyText(diffs))
				return false
			}
		}
	}

	return true
}
