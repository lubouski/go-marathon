package vlc

import (
	"strings"
)

type DecodingTree struct {
	Value string
	Zero, One *DecodingTree
}

func (et encodingTable) DecodingTree() DecodingTree {
	res := DecodingTree{}

	for ch, code := range et {
		// dt.Add(code, ch)
		res.Add(code, ch)
	}
	return res
}

func (dt *DecodingTree) Decode(str string) string {
	var buf strings.Builder

	currNode := dt

	for _, ch := range str {
		if currNode.Value != "" {
			buf.WriteString(currNode.Value)
			currNode = dt
		}

		switch ch {
		case '0':
			currNode = currNode.Zero
		case '1':
			currNode = currNode.One
		}
	}

        if currNode.Value != "" {
        	buf.WriteString(currNode.Value)
        	currNode = dt
        }

	return buf.String()
}

func (dt *DecodingTree) Add(code string, value rune) {
	currNode := dt

	for _, ch := range code {
		switch ch {
		case '0':
			if currNode.Zero == nil {
				currNode.Zero = &DecodingTree{}
			}
			currNode = currNode.Zero
		case '1':
			if currNode.One == nil {
				currNode.One = &DecodingTree{}
			}
			currNode = currNode.One
		}
	}
	currNode.Value = string(value)
}
