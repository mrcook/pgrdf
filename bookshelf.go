package pgrdf

// Bookshelf is a Project Gutenberg bookshelf name.
// <pgterms:bookshelf>
type Bookshelf struct {
	// The bookshelf name.
	// <rdf:Description><rdf:value>
	Name string `json:"subject"`

	// Name of bookshelf at gutenberg.org.
	// <rdf:Description><dcam:memberOf rdf:resource="2009/pgterms/Bookshelf"/>
	Resource string `json:"resource"`
}
