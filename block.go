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

	Paragraph        *TextBlock            `json:"paragraph,omitempty"`
	Heading_1        *HeadingBlock         `json:"heading_1,omitempty"`
	Heading_2        *HeadingBlock         `json:"heading_2,omitempty"`
	Heading_3        *HeadingBlock         `json:"heading_3,omitempty"`
	Callout          *CalloutBlock         `json:"callout,omitempty"`
	Quote            *TextBlock            `json:"quote,omitempty"`
	BulletedListItem *TextBlock            `json:"bulleted_list_item,omitempty"`
	NumberedListItem *TextBlock            `json:"numbered_list_item,omitempty"`
	ToDo             *TodoBlock            `json:"to_do,omitempty"`
	Toggle           *TextBlock            `json:"toggle,omitempty"`
	Code             *CodeBlock            `json:"code,omitempty"` //WTF
	ChildPage        *ChildBlock           `json:"child_page,omitempty"`
	ChildDatabase    *ChildBlock           `json:"child_database,omitempty"`
	Embed            *EmbedBlock           `json:"embed,omitempty"`
	Image            *ImageBlock           `json:"image,omitempty"`
	Video            *VideoBlock           `json:"video,omitempty"`
	File             *FileBlock            `json:"file,omitempty"`
	Pdf              *PDFBlock             `json:"pdf,omitempty"`
	Bookmark         *BookmarkBlock        `json:"bookmark,omitempty"`
	Equation         *EquationBlock        `json:"equation,omitempty"`
	Divider          *DividerBlock         `json:"divider,omitempty"`
	TableOfContents  *TableOfContentsBlock `json:"table_of_contents,omitempty"` // Table of contents block?
	BreadCrumbBlocks *BreadCrumbBlock      `json:"breadcrumb_blocks,omitempty"` //WTF
	ColumnList       *ColumnListBlock      `json:"column_list,omitempty"`       //Column Block?
	Column           *ColumnBlock          `json:"column,omitempty"`            //row?
	LinkPreview      *LinkPreviewBlock     `json:"link_preview,omitempty"`      //Link Privew blocks?
	Template         *TemplateBlock        `json:"template,omitempty"`          //Tempalte Blocks?
	LinkToPage       *LinkToPageBlock      `json:"link_to_page,omitempty"`      //Link To page BlocK?
	SyncedBlock      *SyncedBlock          `json:"synced_block,omitempty"`
	Table            *TableBlock           `json:"table,omitempty"`
	TableRow         *TableRowBlock        `json:"table_row,omitempty"`
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

type TextBlock struct {
	RichText []RichText `json:"rich_text"`
	Color    string     `json:"color,omitempty"` //TODO colors in tags has lower count of available values or not? Check and make single color enum
	Children []Block    `json:"children,omitempty"`
}

type HeadingBlock struct {
	RichText []RichText `json:"rich_text"`
	Color    string     `json:"color,omitempty"` //TODO colors in tags has lower count of available values or not? Check and make single color enum
}

type CalloutBlock struct {
	RichText []RichText      `json:"rich_text"`
	Icon     *IconDescriptor `json:"icon,omitempty"`
	Color    string          `json:"color,omitempty"` //TODO colors in tags has lower count of available values or not? Check and make single color enum
	Children []Block         `json:"children,omitempty"`
}

type TodoBlock struct {
	RichText []RichText `json:"rich_text"`
	Checked  bool       `json:"checked"`
	Color    string     `json:"color,omitempty"` //TODO colors in tags has lower count of available values or not? Check and make single color enum
	Children []Block    `json:"children,omitempty"`
}

type CodeBlock struct {
	RichText []RichText `json:"rich_text"`
	Caption  []RichText `json:"caption"`
	Language string     `json:"color,omitempty"` //TODO add language enum
}

type ChildBlock struct {
	Title string `json:"title"`
}

/*
Embed blocks created via Notion API may differ of embed blocks created via UI
Details: https://developers.notion.com/reference/block#embed-blocks
*/
type EmbedBlock struct {
	Url string `json:"url"`
}

type ImageBlock struct {
	Image FileDescriptor `json:"image"`
}

type VideoBlock struct {
	Video FileDescriptor `json:"image"`
}

type FileBlock struct {
	File FileDescriptor `json:"file"`
}

type PDFBlock struct {
	PDF FileDescriptor `json:"pdf"`
}

type BookmarkBlock struct {
	URL     string `json:"url"`
	Caption string `json:"caption"`
}

type EquationBlock struct {
	Expression string `json:"expression"`
}

type DividerBlock struct{}

type TableOfContentsBlock struct {
	Color string `json:"color,omitempty"` //TODO colors in tags has lower count of available values or not? Check and make single color enum
}

type BreadCrumbBlock struct{}

/*
TODO need to test ColumnListBlock and ColumnBlock carefuly https://developers.notion.com/reference/block#column-list-and-column-blocks
*/
type ColumnListBlock struct {
	Children []ColumnBlock `json:"children"`
}

/*
Can contains any block type itself "column", need validation
*/
type ColumnBlock struct {
	Children []Block `json:"children"`
}

/*
Link Preview block objects return the originally pasted url.
NOTE: The link_preview block will only be returned as part of a response.
It cannot be created via the API.
https://developers.notion.com/reference/block#link-preview-blocks
*/
type LinkPreviewBlock struct {
	URL string `json:"link_preview"`
}

type TemplateBlock struct {
	RichText []RichText `json:"rich_text"`
	Children []Block    `json:"children"`
}

type LinkToPageBlock struct {
	Type       string `json:"type"` //TODO need to create enum
	PageID     string `json:"page_id,omitempty"`
	DatabaseID string `json:"database_id,omitempty"`
}

/*
Need validations to prevent creating original synced block with filled SyncedFrom and visa verse
Details: https://developers.notion.com/reference/block#synced-block-blocks
*/
type SyncedBlock struct {
	SyncedFrom *SyncedFrom `json:"synced_from,omitempty"`
	Children   []Block     `json:"children"`
}

type SyncedFrom struct {
	Type    string `json:"type"` //TODO the only possible value now is block, but might need to create enum
	BlockID string `json:"block_id"`
}

type TableBlock struct {
	//Note that this cannot be changed via the public API once a table is created.
	TableWidth      int             `json:"table_width"`
	HasColumnHeader bool            `json:"has_column_header"`
	HasRowHeader    bool            `json:"has_row_header"`
	Children        []TableRowBlock `json:"children"`
}

type TableRowBlock struct {
	Cells []RichText `json:"cells"`
}
