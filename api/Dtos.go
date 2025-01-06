package api

type ArticleDto struct {
	ID        string       `json:"id"`
	Title     string       `json:"title"`
	Subheader string       `json:"subheader"`
	Content   []ContentDto `json:"content"`
	CreatedAt string       `json:"createdAt"`
	UpdatedAt string       `json:"updatedAt"`
}

type BaseContentDto struct {
	Order int `json:"order"`
}
type ContentDto interface{}

type ParagraphDto struct {
	Metadata BaseContentDto `json:"metadata"`
	Text     string         `json:"text"`
}
