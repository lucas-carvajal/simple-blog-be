package api

import (
	"encoding/json"
	"fmt"
	"simple-blog-be/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

// ToEntity converts ArticleDto to repository.Article
func (dto *ArticleDto) ToEntity() (*repository.ArticleEntity, error) {
	var id primitive.ObjectID
	var err error

	if dto.ID != "" {
		id, err = primitive.ObjectIDFromHex(dto.ID)
		if err != nil {
			return nil, err
		}
	}

	// Convert content DTOs to entities
	content := make([]repository.ContentEntity, len(dto.Content))
	for i, c := range dto.Content {
		if paragraphDto, ok := c.(ParagraphDto); ok {
			content[i] = repository.ParagraphEntity{
				Metadata: repository.BaseContentEntity{
					Order: paragraphDto.Metadata.Order,
				},
				Text: paragraphDto.Text,
			}
		}
		// Add more content type conversions here as needed
	}

	// Parse timestamps if they exist
	var createdAt, updatedAt time.Time
	if dto.CreatedAt != "" {
		createdAt, err = time.Parse(time.RFC3339, dto.CreatedAt)
		if err != nil {
			return nil, err
		}
	}
	if dto.UpdatedAt != "" {
		updatedAt, err = time.Parse(time.RFC3339, dto.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &repository.ArticleEntity{
		ID:        id,
		Title:     dto.Title,
		Subheader: dto.Subheader,
		Content:   content,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

// Custom unmarshal method for ArticleDto
func (a *ArticleDto) UnmarshalJSON(data []byte) error {
	// Create an auxiliary struct with same fields but Content as []map[string]interface{}
	type Aux struct {
		ID        string                   `json:"id"`
		Title     string                   `json:"title"`
		Subheader string                   `json:"subheader"`
		Content   []map[string]interface{} `json:"content"`
		CreatedAt string                   `json:"createdAt"`
		UpdatedAt string                   `json:"updatedAt"`
	}

	var aux Aux
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Copy the simple fields
	a.ID = aux.ID
	a.Title = aux.Title
	a.Subheader = aux.Subheader
	a.CreatedAt = aux.CreatedAt
	a.UpdatedAt = aux.UpdatedAt

	// Convert the content items
	a.Content = make([]ContentDto, len(aux.Content))
	for i, c := range aux.Content {
		// Extract metadata
		metadataMap, ok := c["metadata"].(map[string]interface{})
		if !ok {
			return fmt.Errorf("invalid metadata format at content index %d", i)
		}

		// Convert order to int
		order, ok := metadataMap["order"].(float64)
		if !ok {
			return fmt.Errorf("invalid order format at content index %d", i)
		}

		// Extract text
		text, ok := c["text"].(string)
		if !ok {
			return fmt.Errorf("invalid text format at content index %d", i)
		}

		// Create ParagraphDto
		a.Content[i] = ParagraphDto{
			Metadata: BaseContentDto{
				Order: int(order),
			},
			Text: text,
		}
	}

	return nil
}

// FromEntity converts repository.Article to ArticleDto
func FromEntity(entity *repository.ArticleEntity) *ArticleDto {
	// Convert content entities to DTOs
	content := make([]ContentDto, len(entity.Content))
	for i, c := range entity.Content {
		if paragraph, ok := c.(repository.ParagraphEntity); ok {
			content[i] = ParagraphDto{
				Metadata: BaseContentDto{
					Order: paragraph.Metadata.Order,
				},
				Text: paragraph.Text,
			}
		}
		// Add more content type conversions here as needed
	}

	return &ArticleDto{
		ID:        entity.ID.Hex(),
		Title:     entity.Title,
		Subheader: entity.Subheader,
		Content:   content,
		CreatedAt: entity.CreatedAt.Format(time.RFC3339),
		UpdatedAt: entity.UpdatedAt.Format(time.RFC3339),
	}
}

// FromEntities converts a slice of repository.Article to a slice of ArticleDto
func FromEntities(entities []repository.ArticleEntity) []ArticleDto {
	dtos := make([]ArticleDto, len(entities))
	for i, entity := range entities {
		dtos[i] = *FromEntity(&entity)
	}
	return dtos
}

// FromEntityWithFirstContent converts repository.ArticleEntity to ArticleDto with only the first content
func FromEntityWithFirstContent(entity *repository.ArticleEntity) *ArticleDto {
	// Find content with order = 1
	var firstContent ContentDto
	for _, c := range entity.Content {
		if paragraph, ok := c.(repository.ParagraphEntity); ok {
			if paragraph.Metadata.Order == 1 {
				firstContent = ParagraphDto{
					Metadata: BaseContentDto{
						Order: paragraph.Metadata.Order,
					},
					Text: paragraph.Text,
				}
				break
			}
		}
	}

	// Create content slice with only the first content if found
	var content []ContentDto
	if firstContent != nil {
		content = []ContentDto{firstContent}
	} else {
		content = []ContentDto{}
	}

	return &ArticleDto{
		ID:        entity.ID.Hex(),
		Title:     entity.Title,
		Subheader: entity.Subheader,
		Content:   content,
		CreatedAt: entity.CreatedAt.Format(time.RFC3339),
		UpdatedAt: entity.UpdatedAt.Format(time.RFC3339),
	}
}

// FromEntitiesWithFirstContent converts a slice of repository.ArticleEntity to a slice of ArticleDto with only first content
func FromEntitiesWithFirstContent(entities []repository.ArticleEntity) []ArticleDto {
	dtos := make([]ArticleDto, len(entities))
	for i, entity := range entities {
		dtos[i] = *FromEntityWithFirstContent(&entity)
	}
	return dtos
}
