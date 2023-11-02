package pgrdf

import (
	"io"
	"strings"

	"github.com/mrcook/pgrdf/internal/unmarshaler"
)

// rdfUnmarshal will deserialise an RDF object to an Ebook object.
func rdfUnmarshal(r io.Reader) (*Ebook, error) {
	rdf, err := unmarshaler.New(r)
	if err != nil {
		return nil, err
	}

	ebook := &Ebook{
		ID:                      rdf.Ebook.Id(),
		Titles:                  splitTitles(rdf.Ebook.Titles),
		AlternateTitles:         splitTitles(rdf.Ebook.Alternatives),
		TableOfContents:         rdf.Ebook.TableOfContents,
		Publisher:               rdf.Ebook.Publisher,
		PublishedYear:           rdf.Ebook.PublishedYear,
		ReleaseDate:             rdf.Ebook.Issued.Value,
		Summary:                 rdf.Ebook.Summary,
		Series:                  rdf.Ebook.Series,
		Languages:               nil,
		LanguageDialect:         rdf.Ebook.LanguageDialect,
		LanguageNotes:           rdf.Ebook.LanguageNotes,
		PublicationNote:         rdf.Ebook.PublicationNote,
		EditionNote:             rdf.Ebook.EditionNote,
		ProductionNotes:         rdf.Ebook.ProductionNotes,
		Copyright:               rdf.Ebook.Rights,
		CopyrightClearanceCode:  rdf.Ebook.DpClearanceCode,
		BookType:                "",
		Notes:                   rdf.Ebook.Descriptions,
		PhysicalDescriptionNote: rdf.Ebook.PhysicalDescriptionNote,
		SourceLinks:             rdf.Ebook.SourceLinks,
		LCCN:                    rdf.Ebook.LCCN,
		ISBN:                    rdf.Ebook.ISBN,
		BookCovers:              nil,
		TitlePageImage:          rdf.Ebook.TitlePageImage,
		BackCover:               rdf.Ebook.BackCoverImage,
		Creators:                nil,
		Subjects:                nil,
		Files:                   nil,
		Bookshelves:             nil,
		Downloads:               rdf.Ebook.Downloads.Value,
		AuthorLinks:             nil,
		CCComment:               rdf.Work.Comment,
		CCLicense:               rdf.Work.License.Resource,
	}
	if rdf.Ebook.Type.Description.Value != nil {
		ebook.SetBookType(rdf.Ebook.Type.Description.Value.Data)
	}

	for _, cover := range rdf.Ebook.BookCoverImages {
		ebook.BookCovers = append(ebook.BookCovers, bookCoverFilename(cover))
	}

	for _, lang := range rdf.Ebook.Languages {
		// NOTE: this should never happen, but let's check for nil anyway
		if lang.Description.Value != nil {
			ebook.Languages = append(ebook.Languages, lang.Description.Value.Data)
		}
	}

	for _, l := range rdf.Descriptions {
		ebook.AddAuthorLink(l.Description, l.About)
	}

	for _, c := range rdf.Ebook.Creators {
		creator := Creator{
			ID:      c.Agent.Id(),
			Name:    c.Agent.Name,
			Aliases: c.Agent.Aliases,
			Role:    RoleAut,
		}
		if c.Agent.BirthYear != nil {
			creator.Born = c.Agent.BirthYear.Value
		}
		if c.Agent.DeathYear != nil {
			creator.Died = c.Agent.DeathYear.Value
		}
		for _, webpage := range c.Agent.Webpages {
			creator.WebPages = append(creator.WebPages, webpage.Resource)
		}
		ebook.AddCreator(creator)
	}

	for _, rel := range rdf.Ebook.RelAdapters {
		addRelatorToCreators(ebook, rel, RoleAdp)
	}
	for _, rel := range rdf.Ebook.RelAfterwords {
		addRelatorToCreators(ebook, rel, RoleAft)
	}
	for _, rel := range rdf.Ebook.RelAnnotators {
		addRelatorToCreators(ebook, rel, RoleAnn)
	}
	for _, rel := range rdf.Ebook.RelArrangers {
		addRelatorToCreators(ebook, rel, RoleArr)
	}
	for _, rel := range rdf.Ebook.RelArtists {
		addRelatorToCreators(ebook, rel, RoleArt)
	}
	for _, rel := range rdf.Ebook.RelIntroductions {
		addRelatorToCreators(ebook, rel, RoleAui)
	}
	for _, rel := range rdf.Ebook.RelCommentators {
		addRelatorToCreators(ebook, rel, RoleCmm)
	}
	for _, rel := range rdf.Ebook.RelComposers {
		addRelatorToCreators(ebook, rel, RoleCmp)
	}
	for _, rel := range rdf.Ebook.RelConductors {
		addRelatorToCreators(ebook, rel, RoleCnd)
	}
	for _, rel := range rdf.Ebook.RelCompilers {
		addRelatorToCreators(ebook, rel, RoleCom)
	}
	for _, rel := range rdf.Ebook.RelContributors {
		addRelatorToCreators(ebook, rel, RoleCtb)
	}
	for _, rel := range rdf.Ebook.RelDubious {
		addRelatorToCreators(ebook, rel, RoleDub)
	}
	for _, rel := range rdf.Ebook.RelEditors {
		addRelatorToCreators(ebook, rel, RoleEdt)
	}
	for _, rel := range rdf.Ebook.RelEngravers {
		addRelatorToCreators(ebook, rel, RoleEgr)
	}
	for _, rel := range rdf.Ebook.RelIllustrators {
		addRelatorToCreators(ebook, rel, RoleIll)
	}
	for _, rel := range rdf.Ebook.RelLibrettists {
		addRelatorToCreators(ebook, rel, RoleLbt)
	}
	for _, rel := range rdf.Ebook.RelOther {
		addRelatorToCreators(ebook, rel, RoleOth)
	}
	for _, rel := range rdf.Ebook.RelPublishers {
		addRelatorToCreators(ebook, rel, RolePbl)
	}
	for _, rel := range rdf.Ebook.RelPhotographers {
		addRelatorToCreators(ebook, rel, RolePht)
	}
	for _, rel := range rdf.Ebook.RelPerformers {
		addRelatorToCreators(ebook, rel, RolePrf)
	}
	for _, rel := range rdf.Ebook.RelPrinters {
		addRelatorToCreators(ebook, rel, RolePrt)
	}
	for _, rel := range rdf.Ebook.RelResearchers {
		addRelatorToCreators(ebook, rel, RoleRes)
	}
	for _, rel := range rdf.Ebook.RelTranscribers {
		addRelatorToCreators(ebook, rel, RoleTrc)
	}
	for _, rel := range rdf.Ebook.RelTranslators {
		addRelatorToCreators(ebook, rel, RoleTrl)
	}

	// NOTE: this should be deprecated
	for _, rel := range rdf.Ebook.RelCollaborators {
		addRelatorToCreators(ebook, rel, RoleClb)
	}

	for _, rel := range rdf.Ebook.RelUnknown {
		addRelatorToCreators(ebook, rel, RoleUnk)
	}

	for _, s := range rdf.Ebook.Subjects {
		if s.Description.Value == nil || s.Description.MemberOf == nil {
			continue // NOTE: this should never happen, but let's check for nil anyway
		}
		ebook.AddSubject(s.Description.Value.Data, s.Description.MemberOf.Resource)
	}
	for _, f := range rdf.Ebook.HasFormats {
		file := File{
			URL:      f.File.About,
			Extent:   f.File.Extent.Value,
			Modified: f.File.Modified.Value,
		}
		for _, f := range f.File.Formats {
			if f.Description.Value == nil {
				continue // NOTE: this should never happen, but let's check for nil anyway
			}
			file.AddEncoding(f.Description.Value.Data)
		}
		ebook.AddBookFile(file)
	}
	for _, s := range rdf.Ebook.Bookshelves {
		if s.Description.Value == nil || s.Description.MemberOf == nil {
			continue // NOTE: this should never happen, but let's check for nil anyway
		}
		ebook.AddBookshelf(s.Description.Value.Data, s.Description.MemberOf.Resource)
	}

	return ebook, nil
}

func splitTitles(titles []string) []string {
	var newTitles []string
	for _, title := range titles {
		title = strings.ReplaceAll(title, "\r", "\n")
		for _, t := range strings.Split(title, "\n") {
			t = strings.TrimSpace(t)
			if len(t) > 0 {
				newTitles = append(newTitles, t)
			}
		}
	}
	return newTitles
}

// Extract the book cover filename from the file path.
// marc901 tags contain a book cover filename from the HTML version of the ebook.
func bookCoverFilename(cover string) string {
	parts := strings.Split(cover, "-h")
	cover = parts[len(parts)-1]
	cover = strings.TrimPrefix(cover, "/")
	return cover
}

// addRelatorToCreators appends a MARC relator to the creators list with the given role and agent (if present).
func addRelatorToCreators(e *Ebook, relator unmarshaler.MarcRelator, role MarcRelator) {
	creator := Creator{
		ID:   relator.AgentId(),
		Role: role,
	}

	if relator.Agent != nil {
		creator.Name = relator.Agent.Name
		creator.Aliases = relator.Agent.Aliases

		if relator.Agent.BirthYear != nil {
			creator.Born = relator.Agent.BirthYear.Value
		}
		if relator.Agent.DeathYear != nil {
			creator.Died = relator.Agent.DeathYear.Value
		}
		for _, webpage := range relator.Agent.Webpages {
			creator.WebPages = append(creator.WebPages, webpage.Resource)
		}
	}

	e.AddCreator(creator)
}
