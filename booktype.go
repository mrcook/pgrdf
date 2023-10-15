package pgrdf

type BookType string

const (
	BookTypeUnknown     BookType = ""
	BookTypeCollection  BookType = "Collection"
	BookTypeDataset     BookType = "Dataset"
	BookTypeImage       BookType = "Image"
	BookTypeMovingImage BookType = "MovingImage"
	BookTypeSound       BookType = "Sound"
	BookTypeStillImage  BookType = "StillImage"
	BookTypeText        BookType = "Text"
)
