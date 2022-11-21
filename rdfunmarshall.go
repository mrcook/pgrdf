package pgrdf

import (
	"io"
	"strings"

	"github.com/mrcook/pgrdf/_internal/unmarshaller"
)

// rdfUnmarshall unmarshalls an RDF document to an Ebook.
func rdfUnmarshall(r io.Reader) (*Ebook, error) {
	rdf, err := unmarshaller.New(r)
	if err != nil {
		return nil, err
	}

	ebook := &Ebook{
		ID:                rdf.Ebook.Id(),
		BookType:          rdf.Ebook.Type.Description.Value.Data,
		ReleaseDate:       rdf.Ebook.Issued.Value,
		Publisher:         rdf.Ebook.Publisher,
		PublishedYear:     rdf.Ebook.PublishedYear,
		Copyright:         rdf.Ebook.Rights,
		Titles:            titles(rdf.Ebook.Title),
		OtherTitles:       rdf.Ebook.Alternative,
		BookCoverFilename: bookCoverFilename(rdf.Ebook.BookCover),
		Downloads:         rdf.Ebook.Downloads.Value,
		Series:            rdf.Ebook.Series,
		Note:              rdf.Ebook.Description,
		Comment:           rdf.Work.Comment,
		CCLicense:         rdf.Work.License.Resource,
	}

	ebook.Language = Language{
		Code:    rdf.Ebook.Language.Description.Value.Data,
		Dialect: rdf.Ebook.LanguageDialect,
		Notes:   rdf.Ebook.LanguageNotes,
	}

	for _, l := range rdf.Descriptions {
		ebook.AddAuthorLink(l.Description, l.About)
	}
	for _, c := range rdf.Ebook.Creators {
		ebook.AddCreator(*createCreator(&c.Agent, RoleAut))
	}
	for _, t := range rdf.Ebook.Compilers {
		ebook.AddCreator(*createCreator(&t.Agent, RoleCom))
	}
	for _, t := range rdf.Ebook.Contributors {
		ebook.AddCreator(*createCreator(&t.Agent, RoleCtb))
	}
	for _, e := range rdf.Ebook.Editors {
		ebook.AddCreator(*createCreator(&e.Agent, RoleEdt))
	}
	for _, i := range rdf.Ebook.Illustrators {
		ebook.AddCreator(*createCreator(&i.Agent, RoleIll))
	}
	for _, t := range rdf.Ebook.Translators {
		ebook.AddCreator(*createCreator(&t.Agent, RoleTrl))
	}
	for _, s := range rdf.Ebook.Subjects {
		ebook.AddSubject(s.Description.Value.Data, s.Description.MemberOf.Resource)
	}
	for _, f := range rdf.Ebook.HasFormats {
		file := File{
			URL:      f.File.About,
			Extent:   f.File.Extent.Value,
			Modified: f.File.Modified.Value,
		}
		for _, f := range f.File.Formats {
			file.AddEncoding(f.Description.Value.Data)
		}
		ebook.AddBookFile(file)
	}
	for _, s := range rdf.Ebook.Bookshelves {
		ebook.AddBookshelf(s.Description.Value.Data, s.Description.MemberOf.Resource)
	}

	return ebook, nil
}

func titles(title string) []string {
	return strings.Split(title, "\n")
}

// Extract the book cover filename from the file path.
// marc901 tags contain a book cover filename from the HTML version of the ebook.
func bookCoverFilename(cover string) string {
	parts := strings.Split(cover, "-h")
	cover = parts[len(parts)-1]
	cover = strings.TrimPrefix(cover, "/")
	return cover
}

// addCreator appends an Agent to the creators list with the given role.
func addCreator(e *Ebook, agent *unmarshaller.Agent, role MarcRelatorCode) {
	creator := Creator{
		ID:      agent.Id(),
		Name:    agent.Name,
		Aliases: agent.Aliases,
		Born:    agent.Birthdate.Value,
		Died:    agent.Deathdate.Value,
		Role:    role,
		WebPage: agent.Webpage.Resource,
	}
	e.AddCreator(creator)
}

func createCreator(agent *unmarshaller.Agent, role MarcRelatorCode) *Creator {
	return &Creator{
		ID:      agent.Id(),
		Name:    agent.Name,
		Aliases: agent.Aliases,
		Born:    agent.Birthdate.Value,
		Died:    agent.Deathdate.Value,
		Role:    role,
		WebPage: agent.Webpage.Resource,
	}
}
