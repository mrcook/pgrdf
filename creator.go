package pgrdf

// Creator is a person involved in the creation of the work, such as the author,
// illustrator, etc. When a creator has an `aut` role, the RDF tag used is
// `<dcterms:creator>`, all other roles use `<marcrel:*>` tags.
// `aut` roles MUST include an agent ID and Name but for other roles the
// agent information is optional.
type Creator struct {
	// Unique Project Gutenberg ID.
	// `<pgterms:agent rdf:about="..">`
	ID int `json:"id"`

	// Name of the creator.
	// `<pgterms:name>`
	Name string `json:"name"`

	// Any aliases for the creator.
	// `<pgterms:alias>`
	Aliases []string `json:"aliases,omitempty"`

	// Date of Birth.
	// `<pgterms:birthdate>`
	Born int `json:"born_year,omitempty"`

	// Date of Death.
	// `<pgterms:deathdate>`
	Died int `json:"died_year,omitempty"`

	// Code indicating the role. e.g. `aut`, `edt`, `ill`, etc.
	// `aut` roles are added to the RDF as a `<dcterms:creator>` tag,
	// all other roles use a `<marcrel:*>` tag, e.g. `<marcrel:edt>`.
	Role MarcRelator `json:"role,omitempty"`

	// URLs for this creator (e.g. Wikipedia).
	// `<pgterms:webpage>`
	WebPages []string `json:"webpages,omitempty"`
}
