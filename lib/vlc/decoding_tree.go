package vlc

type DecodingTree struct {
	Value string
	Zero  *DecodingTree
	One   *DecodingTree
}

func (et encodingTable) DecodingTree() DecodingTree {
	res := DecodingTree{}

	for ch, code := range et {
		res.Add(code, ch)
	}

	return res
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
