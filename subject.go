package pgrdf

// Subject is a Dublin Core Vocabulary Encoding Scheme such as LCSH and LCC.
type Subject struct {
	Heading string `json:"heading"` // Heading or other label
	Schema  string `json:"schema"`  // Vocabulary Encoding Scheme (full URL).
}
