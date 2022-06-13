package gotion

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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
		return fmt.Errorf("%v isn't enum value", res)
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
	Bold          bool   `json:"bold"`
	Italic        bool   `json:"italic"`
	Strikethrough bool   `json:"strikethrough"`
	Underline     bool   `json:"underline"`
	Code          bool   `json:"code"`
	Color         string `json:"color"`
}

type Text struct {
	Content string `json:"content"`
	Link    *Link  `json:"link,omitempty"`
}

type Mention struct {
	MentionType string `json:"type,omitempty"`

	User     *User         `json:"user,omitempty"`
	Page     string        `json:"page,omitempty"`
	Database string        `json:"database,omitempty"`
	Date     *DateProperty `json:"date,omitempty"`
	Template *Template     `json:"template_mention,omitempty"`
}

type Equation struct {
	Expression string `json:"expression"`
}

type FileDescriptor struct {
	Type         FileDescriptorType `json:"type"`
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
		return fmt.Errorf("%v isn't enum value", res)
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
	URL        string `json:"url"`
	ExpireTime string `json:"expire_time"`
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
	var t = time.Time{}
	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}
	*dt = DateTimeWrap{Datetime: t}
	return nil
}
