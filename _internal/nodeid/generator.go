package nodeid

import "math/rand"

// Saved NodeIDs to prevent duplicates.
var generatedNodeIDs []string

// Generate a unique RDF `nodeID` value.
func Generate() string {
	const letters = "abcdef0123456789"

	var id string
	for {
		b := make([]byte, 32)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		id = "N" + string(b)

		// check for duplicate ID
		uniq := true
		for _, i := range generatedNodeIDs {
			if id == i {
				uniq = false
				break
			}
		}
		if uniq {
			break
		}
	}
	generatedNodeIDs = append(generatedNodeIDs, id)

	return id
}
