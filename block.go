package gotion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Block struct {
	Object         string       `json:"object,omitempty"`
	ID             string       `json:"id"`
	Parent         BlockParent  `json:"parent"`
	Type           BlockType    `json:"type"`
	CreatedTime    DateTimeWrap `json:"created_time"`
	CreatedBy      User         `json:"created_by"`
	LastEditedTime DateTimeWrap `json:"last_edited_time"`
	LastEditedBy   User         `json:"last_edited_by"`
	Archived       *bool        `json:"archived,omitempty"`
	HasChildred    *bool        `json:"has_children,omitempty"`

	Paragraph        interface{} `json:"paragraph,omitempty"`
	Heading_1        interface{} `json:"heading_1,omitempty"`
	Heading_2        interface{} `json:"heading_2,omitempty"`
	Heading_3        interface{} `json:"heading_3,omitempty"`
	Callout          interface{} `json:"callout,omitempty"`
	Quote            interface{} `json:"quote,omitempty"`
	BulletedListItem interface{} `json:"bulleted_list_item,omitempty"`
	NumberedListItem interface{} `json:"numbered_list_item,omitempty"`
	ToDo             interface{} `json:"to_do,omitempty"`
	Toggle           interface{} `json:"toggle,omitempty"`
	Code             interface{} `json:"code,omitempty"` //WTF
	ChildPage        interface{} `json:"child_page,omitempty"`
	ChildDatabase    interface{} `json:"child_database,omitempty"`
	Embed            interface{} `json:"embed,omitempty"`
	Image            interface{} `json:"image,omitempty"`
	Video            interface{} `json:"video,omitempty"`
	File             interface{} `json:"file,omitempty"`
	Pdf              interface{} `json:"pdf,omitempty"`
	Bookmark         interface{} `json:"bookmark,omitempty"`
	Equation         interface{} `json:"equation,omitempty"`
	Divider          interface{} `json:"divider,omitempty"`
	TableOfContents  interface{} `json:"table_of_contents,omitempty"` // Table of contents block?
	BreadCrumbBlocks interface{} `json:"breadcrumb_blocks,omitempty"` //WTF
	ColumnList       interface{} `json:"column_list,omitempty"`       //Column Block?
	LinkPreview      interface{} `json:"link_preview,omitempty"`      //Link Privew blocks?
	Template         interface{} `json:"template,omitempty"`          //Tempalte Blocks?
	LinkToPage       interface{} `json:"link_to_page,omitempty"`      //Link To page BlocK?
	Table_row        interface{} `json:"table_row,omitempty"`
	SyncedBlock      interface{} `json:"synced_block,omitempty"`
	Column           interface{} `json:"column,omitempty"` //row?
	Table            interface{} `json:"table,omitempty"`
}

type BlockParent struct {
	Type       string `json:"type"`
	BlockId    string `json:"block_id,omitempty"`
	PageId     string `json:"page_id,omitempty"`
	DatabaseId string `json:"database_id,omitempty"`
}

type BlockParentType int

const (
	NoBlockParentType = iota
	PageBlockParentType
	DatabaseBlockParentType
	BlockBlockParentType
)

var BlockParentTypeToString = map[BlockParentType]string{
	PageBlockParentType:     "page_id",
	DatabaseBlockParentType: "database_id",
	BlockBlockParentType:    "block_id",
}

var StringToBlockParentType = map[string]BlockParentType{
	"page_id":     PageBlockParentType,
	"database_id": DatabaseBlockParentType,
	"block_id":    BlockBlockParentType,
}

func (p *BlockParentType) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	res, ok := StringToBlockParentType[v]
	if !ok {
		return fmt.Errorf("%v isn't enum value in %T", res, p)
	}
	*p = res
	return nil
}

func (p BlockParentType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(BlockParentTypeToString[p])
	b.WriteString(`"`)
	return b.Bytes(), nil
}

type BlockType int

const (
	NoBlockType = iota
	ParagraphBlockType
	H1BlockType
	H2BlockType
	H3BlockType
	BulletedListBlockType
	NumberedListBlockType
	ToDoBlockType
	ToggleBlockType
	ChildPageBlockType
	ChildDatabaseBlockType
	EmbedBlockType
	ImageBlockType
	VideoBlockType
	FileBlockType
	PDFBlockType
	BookmarkBlockType
	CalloutBlockType
	QuoteBlockType
	EquationBlockType
	DividerBlockType
	TableOfContentsBlockType
	ColumnBlockType
	ColumnListBlockType
	LinkPreviewBlockType
	SuncedBlockBlockType
	TemplateBlockType
	LinkToPageBlockType
	TableBlockType
	TableRowBlockType
	UnsupportedBlockType
)

var BlockTypeToString = map[BlockType]string{
	ParagraphBlockType:       "paragraph",
	H1BlockType:              "heading_1",
	H2BlockType:              "heading_2",
	H3BlockType:              "heading_3",
	BulletedListBlockType:    "bulleted_list_item",
	NumberedListBlockType:    "numbered_list_item",
	ToDoBlockType:            "to_do",
	ToggleBlockType:          "toggle",
	ChildPageBlockType:       "child_page",
	ChildDatabaseBlockType:   "child_database",
	EmbedBlockType:           "embed",
	ImageBlockType:           "image",
	VideoBlockType:           "video",
	FileBlockType:            "file",
	PDFBlockType:             "pdf",
	BookmarkBlockType:        "bookmark",
	CalloutBlockType:         "callout",
	QuoteBlockType:           "quote",
	EquationBlockType:        "equation",
	DividerBlockType:         "divider",
	TableOfContentsBlockType: "table_of_contents",
	ColumnBlockType:          "column",
	ColumnListBlockType:      "column_list",
	LinkPreviewBlockType:     "link_preview",
	SuncedBlockBlockType:     "synced_block",
	TemplateBlockType:        "template",
	LinkToPageBlockType:      "link_to_page",
	TableBlockType:           "table",
	TableRowBlockType:        "table_row",
	UnsupportedBlockType:     "unsupported",
}

var StringToBlockType = map[string]BlockType{
	"paragraph":          ParagraphBlockType,
	"heading_1":          H1BlockType,
	"heading_2":          H2BlockType,
	"heading_3":          H3BlockType,
	"bulleted_list_item": BulletedListBlockType,
	"numbered_list_item": NumberedListBlockType,
	"to_do":              ToDoBlockType,
	"toggle":             ToggleBlockType,
	"child_page":         ChildPageBlockType,
	"child_database":     ChildDatabaseBlockType,
	"embed":              EmbedBlockType,
	"image":              ImageBlockType,
	"video":              VideoBlockType,
	"file":               FileBlockType,
	"pdf":                PDFBlockType,
	"bookmark":           BookmarkBlockType,
	"callout":            CalloutBlockType,
	"quote":              QuoteBlockType,
	"equation":           EquationBlockType,
	"divider":            DividerBlockType,
	"table_of_contents":  TableOfContentsBlockType,
	"column":             ColumnBlockType,
	"column_list":        ColumnListBlockType,
	"link_preview":       LinkPreviewBlockType,
	"synced_block":       SuncedBlockBlockType,
	"template":           TemplateBlockType,
	"link_to_page":       LinkToPageBlockType,
	"table":              TableBlockType,
	"table_row":          TableRowBlockType,
	"unsupported":        UnsupportedBlockType,
}

func (p *BlockType) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	res, ok := StringToBlockType[v]
	if !ok {
		return fmt.Errorf("%v isn't enum value in %T", res, p)
	}
	*p = res
	return nil
}

func (p BlockType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(BlockTypeToString[p])
	b.WriteString(`"`)
	return b.Bytes(), nil
}
