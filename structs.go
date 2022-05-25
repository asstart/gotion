package gotion

import (
	"encoding/json"
	"time"
)

type DatabasePages struct {
	Results    []Page  `json:"results"`
	NextCursor *string `json:"next_cursor"`
	HasMore    *bool   `json:"has_more"`
}

type UpdatePage struct {
	Properties PageProperties `json:"properties"`
	Archived   *bool          `json:"archive,omitempty"`
	Icon       *string        `json:"icon,omitempty"`
	Cover      *string        `json:"cover,omitempty"`
}

//TODO object property
type Page struct {
	ID             *string     `json:"id"`
	CreatedTime    *time.Time  `json:"created_time"`
	CreatedBy      *User       `json:"created_by"`
	LastEditedTime *time.Time  `json:"last_edited_time"`
	LastEditedBy   *User       `json:"last_edited_by"`
	Parent         *Parent     `json:"parent"`
	Icon           *Emoji      `json:"icon,omitempty"`
	Cover          *File       `json:"cover,omitempty"`
	Archived       *bool       `json:"archived"`
	Properties     interface{} `json:"properties"`
	URL            *string     `json:"url"`
}

type Parent struct {
	Type       *string `json:"type"`
	DatabaseID *string `json:"database_id,omitempty"`
	PageID     *string `json:"page_id,omitempty"`
	Workspace  *bool   `json:"workspace,omitempty"`
}

type PageProperties map[string]PageProperty

type PageProperty struct {
	ID   *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
	Name *string `json:"name,omitempty"`

	Title          []RichText     `json:"title,omitempty"`
	RichText       []RichText     `json:"rich_text,omitempty"`
	Number         *float64       `json:"number,omitempty"`
	Select         *SelectOption  `json:"select,omitempty"`
	MultiSelect    []SelectOption `json:"multi_select,omitempty"`
	Date           *DateProperty  `json:"date,omitempty"`
	People         []User         `json:"people,omitempty"`
	Files          []File         `json:"files,omitempty"`
	Checkbox       *bool          `json:"checkbox,omitempty"`
	URL            *string        `json:"url,omitempty"`
	Email          *string        `json:"email,omitempty"`
	PhoneNumber    *string        `json:"phone_number,omitempty"`
	Formula        *Formula       `json:"formula,omitempty"`
	Relation       *Relation      `json:"relation,omitempty"`
	Rollup         *Rollup        `json:"rollup,omitempty"`
	CreatedTime    *time.Time     `json:"created_time,omitempty"`
	CreatedBy      *User          `json:"created_by,omitempty"`
	LastEditedTime *time.Time     `json:"last_edited_time,omitempty"`
	LastEditedBy   *User          `json:"last_edited_by,omitempty"`
}

//TODO Check Start,End correctness
type DateProperty struct {
	Start    *string `json:"start"`
	End      *string `json:"end"`
	TimeZone *string `json:"time_zone"`
}

type SelectOption struct {
	Name  *string `json:"name,omitempty"`
	ID    *string `json:"id,omitempty"`
	Color *string `json:"color,omitempty"`
}

type Formula struct {
	Type           *string    `json:"type"`
	StringFormula  *string    `json:"string,omitempty"`
	BooleanFormula *bool      `json:"boolean,omitempty"`
	NumberFormula  *float64   `json:"number,omitempty"`
	DateFormula    *time.Time `json:"date,omitempty"`
	// Expression *string `json:"expression"`
}

type Relation struct {
	DatabaseID         *string `json:"database_id"`
	SyncedPropertyName *string `json:"synced_property_name"`
	SyncedPropertyID   *string `json:"synced_property_id"`
}

type Rollup struct {
	RelationPropertyName string `json:"relation_property_name"`
	RelationPropertyID   string `json:"relation_property_id"`
	RollupPropertyName   string `json:"rollup_property_name"`
	RollupPropertyID     string `json:"rollup_property_id"`
	Function             string `json:"function"`
}

type Template struct {
	TemplateMentionDate *string `json:"template_mention_date"`
	TemplateMentionUser *string `json:"template_mention_user"`
}

type Link struct {
	URL *string `json:"url"`
}

type File struct {
	Type         *string       `json:"type,omitempty"`
	Name         *string       `json:"name,omitempty"`
	ExternalFile *ExternalFile `json:"external,omitempty"`
	NotionFile   *NotionFile   `json:"file,omitempty"`
}

func (p *Page) UnmarshalJSON(b []byte) error {
	type (
		P   Page
		DTO struct {
			P
			Properties json.RawMessage `json:"properties"`
		}
	)

	var dto DTO

	if err := json.Unmarshal(b, &dto); err != nil {
		return err
	}

	page := dto.P

	var props PageProperties

	if err := json.Unmarshal(dto.Properties, &props); err != nil {
		return err
	}

	page.Properties = props

	*p = Page(page)
	return nil
}
