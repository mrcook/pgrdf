package pgrdf

// Subject is a Dublin Core Vocabulary Encoding Scheme such as LCSH and LCC.
// <dcterms:subject>
type Subject struct {
	// Heading or other label
	// <rdf:Description><rdf:value>
	Heading string `json:"heading"`

	// Vocabulary Encoding Scheme.
	// Usually http://purl.org/dc/terms/LCSH or http://purl.org/dc/terms/LCC
	// <rdf:Description><dcam:memberOf rdf:resource="..."/>
	Schema string `json:"schema"`
}
