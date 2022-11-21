package pgrdf

// Bookshelf is a Project Gutenberg bookshelf.
type Bookshelf struct {
	Name     string `json:"subject"` // The bookshelf name.
	Resource string `json:"name"`    // Name of bookshelf at gutenberg.org. Usually "2009/pgterms/Bookshelf".
}
