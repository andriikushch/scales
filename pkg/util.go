package scales

import (
	"strconv"
	"strings"
)

func initShapes(shapes []ChordShape) map[string][]ChordShape {
	res := make(map[string][]ChordShape)

	for _, v := range shapes {
		chordIDAsStrings := []string{}
		for _, k := range v.ChordID {
			chordIDAsStrings = append(chordIDAsStrings, strconv.Itoa(k))
		}

		key := strings.Join(chordIDAsStrings, "-")
		res[key] = append(res[key], v)
	}

	return res
}
