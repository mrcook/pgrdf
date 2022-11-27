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
		Titles:            titles(rdf.Ebook.Title),
		OtherTitles:       rdf.Ebook.Alternative,
		Publisher:         rdf.Ebook.Publisher,
		ReleaseDate:       rdf.Ebook.Issued.Value,
		Series:            rdf.Ebook.Series,
		PublishedYear:     rdf.Ebook.PublishedYear,
		Copyright:         rdf.Ebook.Rights,
		BookType:          rdf.Ebook.Type.Description.Value.Data,
		Note:              rdf.Ebook.Description,
		BookCoverFilename: bookCoverFilename(rdf.Ebook.BookCover),
		Downloads:         rdf.Ebook.Downloads.Value,
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
		ebook.AddCreator(*createCreator(&c.Agent, RoleAut, c.Agent.Id()))
	}
	for _, c := range rdf.Ebook.RelAdapters {
		ebook.AddCreator(*createCreator(&c.Agent, RoleAdp, c.AgentId()))
	}
	for _, c := range rdf.Ebook.RelAfterwords {
		ebook.AddCreator(*createCreator(&c.Agent, RoleAft, c.AgentId()))
	}
	for _, c := range rdf.Ebook.RelAnnotators {
		ebook.AddCreator(*createCreator(&c.Agent, RoleAnn, c.AgentId()))
	}
	for _, c := range rdf.Ebook.RelArrangers {
		ebook.AddCreator(*createCreator(&c.Agent, RoleArr, c.AgentId()))
	}
	for _, c := range rdf.Ebook.RelArtists {
		ebook.AddCreator(*createCreator(&c.Agent, RoleArt, c.AgentId()))
	}
	for _, c := range rdf.Ebook.RelIntroductions {
		ebook.AddCreator(*createCreator(&c.Agent, RoleAui, c.AgentId()))
	}
	for _, c := range rdf.Ebook.RelCommentators {
		ebook.AddCreator(*createCreator(&c.Agent, RoleCmm, c.AgentId()))
	}
	for _, c := range rdf.Ebook.RelComposers {
		ebook.AddCreator(*createCreator(&c.Agent, RoleCmp, c.AgentId()))
	}
	for _, c := range rdf.Ebook.RelConductors {
		ebook.AddCreator(*createCreator(&c.Agent, RoleCnd, c.AgentId()))
	}
	for _, c := range rdf.Ebook.RelCompilers {
		ebook.AddCreator(*createCreator(&c.Agent, RoleCom, c.AgentId()))
	}
	for _, c := range rdf.Ebook.RelContributors {
		ebook.AddCreator(*createCreator(&c.Agent, RoleCtb, c.AgentId()))
	}
	for _, c := range rdf.Ebook.RelDubious {
		ebook.AddCreator(*createCreator(&c.Agent, RoleDub, c.AgentId()))
	}
	for _, c := range rdf.Ebook.RelEditors {
		ebook.AddCreator(*createCreator(&c.Agent, RoleEdt, c.AgentId()))
	}
	for _, c := range rdf.Ebook.RelEngravers {
		ebook.AddCreator(*createCreator(&c.Agent, RoleEgr, c.AgentId()))
	}
	for _, c := range rdf.Ebook.RelIllustrators {
		ebook.AddCreator(*createCreator(&c.Agent, RoleIll, c.AgentId()))
	}
	for _, c := range rdf.Ebook.RelLibrettists {
		ebook.AddCreator(*createCreator(&c.Agent, RoleLbt, c.AgentId()))
	}
	for _, c := range rdf.Ebook.RelOther {
		ebook.AddCreator(*createCreator(&c.Agent, RoleOth, c.AgentId()))
	}
	for _, c := range rdf.Ebook.RelPublishers {
		ebook.AddCreator(*createCreator(&c.Agent, RolePbl, c.AgentId()))
	}
	for _, c := range rdf.Ebook.RelPhotographers {
		ebook.AddCreator(*createCreator(&c.Agent, RolePht, c.AgentId()))
	}
	for _, c := range rdf.Ebook.RelPerformers {
		ebook.AddCreator(*createCreator(&c.Agent, RolePrf, c.AgentId()))
	}
	for _, c := range rdf.Ebook.RelPrinters {
		ebook.AddCreator(*createCreator(&c.Agent, RolePrt, c.AgentId()))
	}
	for _, c := range rdf.Ebook.RelResearchers {
		ebook.AddCreator(*createCreator(&c.Agent, RoleRes, c.AgentId()))
	}
	for _, c := range rdf.Ebook.RelTranscribers {
		ebook.AddCreator(*createCreator(&c.Agent, RoleTrc, c.AgentId()))
	}
	for _, c := range rdf.Ebook.RelTranslators {
		ebook.AddCreator(*createCreator(&c.Agent, RoleTrl, c.AgentId()))
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
func addCreator(e *Ebook, agent *unmarshaller.Agent, role MarcRelator) {
	creator := Creator{
		ID:      agent.Id(),
		Name:    agent.Name,
		Aliases: agent.Aliases,
		Born:    agent.BirthYear.Value,
		Died:    agent.DeathYear.Value,
		Role:    role,
		WebPage: agent.Webpage.Resource,
	}
	e.AddCreator(creator)
}

func createCreator(agent *unmarshaller.Agent, role MarcRelator, agentId int) *Creator {
	return &Creator{
		ID:      agentId,
		Name:    agent.Name,
		Aliases: agent.Aliases,
		Born:    agent.BirthYear.Value,
		Died:    agent.DeathYear.Value,
		Role:    role,
		WebPage: agent.Webpage.Resource,
	}
}
