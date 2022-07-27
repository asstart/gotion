package gotion

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type RichText struct {
	Type        RichTextType `json:"type,omitempty"`
	PlainText   string       `json:"plain_text,omitempty"`
	Href        string       `json:"href,omitempty"`
	Annotations *Annotations `json:"annotations,omitempty"`
	//One of followings should be define
	Text     *Text     `json:"text,omitempty"`
	Mention  *Mention  `json:"mention,omitempty"`
	Equation *Equation `json:"equation,omitempty"`
}

type RichTextType int

const (
	NoRichTextType RichTextType = iota
	TextRichTextType
	MentionRichTextType
	EquationRichTextType
)

var StringToRichTextType = map[string]RichTextType{
	"text":     TextRichTextType,
	"mention":  MentionRichTextType,
	"equation": EquationRichTextType,
}

var RichTextTypeToString = map[RichTextType]string{
	TextRichTextType:     "text",
	MentionRichTextType:  "mention",
	EquationRichTextType: "equation",
}

func (p *RichTextType) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	res, ok := StringToRichTextType[v]
	if !ok {
		return fmt.Errorf("%v isn't enum value in %T", res, p)
	}
	*p = res
	return nil
}

func (p RichTextType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(RichTextTypeToString[p])
	b.WriteString(`"`)
	return b.Bytes(), nil
}

type Annotations struct {
	Bold          bool          `json:"bold"`
	Italic        bool          `json:"italic"`
	Strikethrough bool          `json:"strikethrough"`
	Underline     bool          `json:"underline"`
	Code          bool          `json:"code"`
	Color         PropertyColor `json:"color"`
}

type Text struct {
	Content string `json:"content"`
	Link    *Link  `json:"link,omitempty"`
}

type Link struct {
	URL string `json:"url"`
}

type Mention struct {
	Type MentionType `json:"type,omitempty"`

	User     *User             `json:"user,omitempty"`
	Page     *IDWrap           `json:"page,omitempty"`
	Database *IDWrap           `json:"database,omitempty"`
	Date     *DatePageProperty `json:"date,omitempty"`
	Template *Template         `json:"template_mention,omitempty"`
}

type IDWrap struct {
	ID string `json:"id"`
}

type MentionType int

const (
	NoMentionType MentionType = iota
	UserMentionType
	PageMentionType
	DatabaseMentionType
	DateMentionType
	LinkPriviewMentionType
)

var MentionTypeToString = map[MentionType]string{
	UserMentionType:        "user",
	PageMentionType:        "page",
	DatabaseMentionType:    "database",
	DateMentionType:        "date",
	LinkPriviewMentionType: "link_preview",
}

var StringToMentionType = map[string]MentionType{
	"user":        UserMentionType,
	"page":        PageMentionType,
	"database":    DatabaseMentionType,
	"date":        DateMentionType,
	"link_privew": LinkPriviewMentionType,
}

func (p *MentionType) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	res, ok := StringToMentionType[v]
	if !ok {
		return fmt.Errorf("%v isn't enum value in %T", res, p)
	}
	*p = res
	return nil
}

func (p MentionType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(MentionTypeToString[p])
	b.WriteString(`"`)
	return b.Bytes(), nil
}

type Template struct {
	TemplateMentionDate *string `json:"template_mention_date"`
	TemplateMentionUser *string `json:"template_mention_user"`
}

type Equation struct {
	Expression string `json:"expression"`
}

type FileDescriptor struct {
	Type         FileDescriptorType `json:"type"`
	Name         string             `json:"name,omitempty"`
	ExternalFile *ExternalFile      `json:"external,omitempty"`
	NotionFile   *NotionFile        `json:"file,omitempty"`
}

type FileDescriptorType int

const (
	NoFileDescriptorType FileDescriptorType = iota
	NotionFileDescriptorType
	ExternalFileDescriptorType
)

var FileDescriptorTypeToString = map[FileDescriptorType]string{
	NotionFileDescriptorType:   "file",
	ExternalFileDescriptorType: "external",
}

var StringToFileDescriptorType = map[string]FileDescriptorType{
	"file":     NotionFileDescriptorType,
	"external": ExternalFileDescriptorType,
}

func (fdt FileDescriptor) ValidateRequest() error {
	if fdt.Type == NoFileDescriptorType {
		return errors.New("FileDescriptor.Type couldn't be empty")
	}
	return nil
}

func (p *FileDescriptorType) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	res, ok := StringToFileDescriptorType[v]
	if !ok {
		return fmt.Errorf("%v isn't enum value in %T", res, p)
	}
	*p = res
	return nil
}

func (p FileDescriptorType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(FileDescriptorTypeToString[p])
	b.WriteString(`"`)
	return b.Bytes(), nil
}

type ExternalFile struct {
	URL string `json:"url"`
}

type NotionFile struct {
	URL        string        `json:"url"`
	ExpiryTime *DateTimeWrap `json:"expiry_time,omitempty"`
}

type Emoji struct {
	Type  string `json:"type"`
	Emoji string `json:"emoji"`
}

type DateTimeWrap struct {
	Datetime time.Time
}

func (dt DateTimeWrap) MarshalJSON() ([]byte, error) {
	return dt.Datetime.MarshalJSON()
}

type DateTimeEmptyWrap struct{}

func (dt *DateTimeWrap) UnmarshalJSON(b []byte) error {
	input := string(b)
	//Of course this pattern isn't enough to get full validation of input date, but
	//I use it here only to distinguish between date and datetime input
	m, err := regexp.Match(`^"\d{4}-\d{2}-\d{2}"$`, b)
	if err != nil {
		return err
	}
	if m {
		t, err := time.Parse(`"2006-01-02"`, input)
		if err != nil {
			return err
		}
		*dt = DateTimeWrap{Datetime: t}
		return nil
	}
	var t = time.Time{}
	err = json.Unmarshal(b, &t)
	if err != nil {
		return err
	}
	*dt = DateTimeWrap{Datetime: t}
	return nil
}

type PropertyColor int

const (
	DefaultColor PropertyColor = iota
	Gray
	Brown
	Orange
	Yellow
	Green
	Blue
	Purple
	Pink
	Red
)

var ColorToString = map[PropertyColor]string{
	DefaultColor: "default",
	Gray:         "gray",
	Brown:        "brown",
	Orange:       "orange",
	Yellow:       "yellow",
	Green:        "green",
	Blue:         "blue",
	Purple:       "purple",
	Pink:         "pink",
	Red:          "red",
}

var StringToColor = map[string]PropertyColor{
	"default": DefaultColor,
	"gray":    Gray,
	"brown":   Brown,
	"orange":  Orange,
	"yellow":  Yellow,
	"green":   Green,
	"blue":    Blue,
	"purple":  Purple,
	"pink":    Pink,
	"red":     Red,
}

func (p *PropertyColor) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	res, ok := StringToColor[v]
	if !ok {
		return fmt.Errorf("%v isn't enum value in %T", res, p)
	}
	*p = res
	return nil
}

func (p PropertyColor) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(ColorToString[p])
	b.WriteString(`"`)
	return b.Bytes(), nil
}

type IconDescriptor struct {
	Type     IconType      `json:"type"`
	Emoji    string        `json:"emoji,omitempty"`
	External *ExternalFile `json:"external,omitempty"`
	File     *NotionFile   `json:"file,omitempty"`
}

//Notion supports icon as a notion file and this can be returned with response.
//But in request supports only emoji and external types
func (id IconDescriptor) ValidateRequest() error {
	var buff = strings.Builder{}
	if id.Type == NoIconType {
		buff.WriteString("IconDescriptor.Type is empty")
	}
	if id.Type == FileIconType {
		buff.WriteString(fmt.Sprintf("IconDescriptor.Type in request supports only %v and %v, but not %v\n", IconTypeToString[EmojiIconType], IconTypeToString[ExternalIconType], IconTypeToString[FileIconType]))
	}
	if id.Type == EmojiIconType && id.Emoji == "" {
		buff.WriteString(fmt.Sprintf("IconDescriptor.Emoji shouldn't be empty if IconDescriptor.Type=%v\n", IconTypeToString[EmojiIconType]))
	}
	if id.Type == ExternalIconType && id.External == nil {
		buff.WriteString(fmt.Sprintf("IconDescriptor.External shouldn't be empty if IconDescriptor.Type=%v\n", IconTypeToString[ExternalIconType]))
	}
	var final = buff.String()
	if final == "" {
		return nil
	}
	return errors.New(buff.String())
}

type IconType int

const (
	NoIconType IconType = iota
	EmojiIconType
	FileIconType
	ExternalIconType
)

var IconTypeToString = map[IconType]string{
	EmojiIconType:    "emoji",
	FileIconType:     "file",
	ExternalIconType: "external",
}

var StringToIconType = map[string]IconType{
	"emoji":    EmojiIconType,
	"file":     FileIconType,
	"external": ExternalIconType,
}

func (p *IconType) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	res, ok := StringToIconType[v]
	if !ok {
		return fmt.Errorf("%v isn't enum value in %T", res, p)
	}
	*p = res
	return nil
}

func (p IconType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(IconTypeToString[p])
	b.WriteString(`"`)
	return b.Bytes(), nil
}
