package tree

var (
	// Vector ..
	Vector = [30][1000]int{}
	// ID ..
	ID = 0
	// Flag ..
	Flag = [1000]bool{}
)

// Insert ..
func Insert(name string) {
	layer := 0
	for _, v := range name {
		index := int(v - 'a')
		if Vector[index][layer] == 0 {
			ID = ID + 1
			Vector[index][layer] = ID
		}

		layer = Vector[index][layer]
	}

	Flag[layer] = true
}

// Find ..
func Find(name string) bool {
	layer := 0
	for _, v := range name {
		index := int(v - 'a')
		if Vector[index][layer] == 0 {
			return false
		}

		layer = Vector[index][layer]
	}

	return Flag[layer]
}
