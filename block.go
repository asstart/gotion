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
	Color    *Color     `json:"color,omitempty"`
	Children []Block    `json:"children,omitempty"`
}

type HeadingBlock struct {
	RichText []RichText `json:"rich_text"`
	Color    *Color     `json:"color,omitempty"`
}

type CalloutBlock struct {
	RichText []RichText      `json:"rich_text"`
	Icon     *IconDescriptor `json:"icon,omitempty"`
	Color    *Color          `json:"color,omitempty"`
	Children []Block         `json:"children,omitempty"`
}

type TodoBlock struct {
	RichText []RichText `json:"rich_text"`
	Checked  bool       `json:"checked"`
	Color    *Color     `json:"color,omitempty"`
	Children []Block    `json:"children,omitempty"`
}

type CodeBlock struct {
	RichText []RichText `json:"rich_text"`
	Caption  []RichText `json:"caption"`
	Language *Language  `json:"color,omitempty"`
}

type Language int

const (
	NoLanguage Language = iota
	ABAP
	Arduino
	Bash
	Basic
	C
	Clojure
	CoffeeScript
	CPlusPlus
	CSharp
	CSS
	Dart
	Diff
	Docker
	Elixir
	Elm
	Erlang
	Flow
	Fortran
	FSharp
	Gherkin
	Glsl
	Go
	GraphQL
	Groovy
	Haskell
	HTML
	Java
	JavaScript
	JSON
	Julia
	Kotlin
	LaTeX
	LESS
	Lisp
	LiveScript
	Lua
	Makefile
	Markdown
	Markup
	Matlab
	Mermaid
	nix
	ObjectiveC
	OCaml
	Pascal
	Perl
	PHP
	PlainText
	PowerShell
	Prolog
	Protobuf
	Python
	R
	Reason
	Ruby
	Rust
	SASS
	Scala
	Scheme
	SCSS
	Shell
	SQL
	Swift
	TypeScript
	VBNet
	Verilog
	VHDL
	VisualBasic
	Webassembly
	XML
	YAML
)

var LanguageToString = map[Language]string{
	ABAP:         "abap",
	Arduino:      "arduino",
	Bash:         "bash",
	Basic:        "basic",
	C:            "c",
	Clojure:      "clojure",
	CoffeeScript: "coffeescript",
	CPlusPlus:    "c++",
	CSharp:       "c#",
	CSS:          "css",
	Dart:         "dart",
	Diff:         "diff",
	Docker:       "docker",
	Elixir:       "elixir",
	Elm:          "elm",
	Erlang:       "erlang",
	Flow:         "flow",
	Fortran:      "fortran",
	FSharp:       "f#",
	Gherkin:      "gherkin",
	Glsl:         "glsl",
	Go:           "go",
	GraphQL:      "graphql",
	Groovy:       "groovy",
	Haskell:      "haskell",
	HTML:         "html",
	Java:         "java",
	JavaScript:   "javascript",
	JSON:         "json",
	Julia:        "julia",
	Kotlin:       "kotlin",
	LaTeX:        "latex",
	LESS:         "less",
	Lisp:         "lisp",
	LiveScript:   "livescript",
	Lua:          "lua",
	Makefile:     "makefile",
	Markdown:     "markdown",
	Markup:       "markup",
	Matlab:       "matlab",
	Mermaid:      "mermaid",
	nix:          "nix",
	ObjectiveC:   "objective-c",
	OCaml:        "ocaml",
	Pascal:       "pascal",
	Perl:         "perl",
	PHP:          "php",
	PlainText:    "plain text",
	PowerShell:   "powershell",
	Prolog:       "prolog",
	Protobuf:     "protobuf",
	Python:       "python",
	R:            "r",
	Reason:       "reason",
	Ruby:         "ruby",
	Rust:         "rust",
	SASS:         "sass",
	Scala:        "scala",
	Scheme:       "scheme",
	SCSS:         "scss",
	Shell:        "shell",
	SQL:          "sql",
	Swift:        "swift",
	TypeScript:   "typescript",
	VBNet:        "vb.net",
	Verilog:      "verilog",
	VHDL:         "vhdl",
	VisualBasic:  "visual basic",
	Webassembly:  "webassembly",
	XML:          "xml",
	YAML:         "yaml",
}

var StringToLanguage = map[string]Language{
	"abap":         ABAP,
	"arduino":      Arduino,
	"bash":         Bash,
	"basic":        Basic,
	"c":            C,
	"clojure":      Clojure,
	"coffeescript": CoffeeScript,
	"c++":          CPlusPlus,
	"c#":           CSharp,
	"css":          CSS,
	"dart":         Dart,
	"diff":         Diff,
	"docker":       Docker,
	"elixir":       Elixir,
	"elm":          Elm,
	"erlang":       Erlang,
	"flow":         Flow,
	"fortran":      Fortran,
	"f#":           FSharp,
	"gherkin":      Gherkin,
	"glsl":         Glsl,
	"go":           Go,
	"graphql":      GraphQL,
	"groovy":       Groovy,
	"haskell":      Haskell,
	"html":         HTML,
	"java":         Java,
	"javascript":   JavaScript,
	"json":         JSON,
	"julia":        Julia,
	"kotlin":       Kotlin,
	"latex":        LaTeX,
	"less":         LESS,
	"lisp":         Lisp,
	"livescript":   LiveScript,
	"lua":          Lua,
	"makefile":     Makefile,
	"markdown":     Markdown,
	"markup":       Markup,
	"matlab":       Matlab,
	"mermaid":      Mermaid,
	"nix":          nix,
	"objective-c":  ObjectiveC,
	"ocaml":        OCaml,
	"pascal":       Pascal,
	"perl":         Perl,
	"php":          PHP,
	"plain text":   PlainText,
	"powershell":   PowerShell,
	"prolog":       Prolog,
	"protobuf":     Protobuf,
	"python":       Python,
	"r":            R,
	"reason":       Reason,
	"ruby":         Ruby,
	"rust":         Rust,
	"sass":         SASS,
	"scala":        Scala,
	"scheme":       Scheme,
	"scss":         SCSS,
	"shell":        Shell,
	"sql":          SQL,
	"swift":        Swift,
	"typescript":   TypeScript,
	"vb.net":       VBNet,
	"verilog":      Verilog,
	"vhdl":         VHDL,
	"visual basic": VisualBasic,
	"webassembly":  Webassembly,
	"xml":          XML,
	"yaml":         YAML,
}

func (l Language) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(LanguageToString[l])
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (l *Language) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	res, ok := StringToLanguage[s]
	if !ok {
		return fmt.Errorf("%v isn't enum value in %T", res, l)
	}
	*l = res
	return nil
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
	Color *Color `json:"color,omitempty"`
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
