package gotion

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Page struct {
	Object         string
	ID             string          `json:"id"`
	CreatedTime    DateTimeWrap    `json:"created_time"`
	CreatedBy      User            `json:"created_by"`
	LastEditedTime DateTimeWrap    `json:"last_edited_time"`
	LastEditedBy   User            `json:"last_edited_by"`
	Parent         PageParent      `json:"parent"`
	Icon           *IconDescriptor `json:"icon,omitempty"`
	Cover          *FileDescriptor `json:"cover,omitempty"`
	Archived       bool            `json:"archived"`
	Properties     PageProperties  `json:"properties"`
	URL            string          `json:"url"`
}

type PageParent struct {
	Type       PageParentType `json:"type"`
	DatabaseID string         `json:"database_id,omitempty"`
	PageID     string         `json:"page_id,omitempty"`
	Workspace  bool           `json:"workspace,omitempty"`
}

type PageParentType int

const (
	NoPageParentType PageParentType = iota
	PageParentDB
	PageParentWorkspace
	PageParentPage
)

var PageParentToString = map[PageParentType]string{
	PageParentDB:        "database_id",
	PageParentWorkspace: "workspace",
	PageParentPage:      "page_id",
}

var StringToPageParent = map[string]PageParentType{
	"database_id": PageParentDB,
	"workspace":   PageParentWorkspace,
	"page_id":     PageParentPage,
}

func (pt PageParentType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(PageParentToString[pt])
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (pt *PageParentType) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	res, ok := StringToPageParent[v]
	if !ok {
		return fmt.Errorf("%v isn't enum value in %T", res, pt)
	}
	*pt = res
	return nil
}

type PageProperties map[string]PageProperty

type PageProperty struct {
	ID   string         `json:"id"`
	Type DBPropertyType `json:"type"`

	Title          []RichText           `json:"title,omitempty"`
	RichText       []RichText           `json:"rich_text,omitempty"`
	Number         *float64             `json:"number,omitempty"`
	Select         *SelectPageProperty  `json:"select,omitempty"`
	MultiSelect    []SelectPageProperty `json:"multi_select,omitempty"`
	Date           *DatePageProperty    `json:"date,omitempty"`
	People         []User               `json:"people,omitempty"`
	Files          []FileDescriptor     `json:"files,omitempty"`
	Checkbox       *bool                `json:"checkbox,omitempty"`
	URL            *string              `json:"url,omitempty"`
	Email          *string              `json:"email,omitempty"`
	PhoneNumber    *string              `json:"phone_number,omitempty"`
	Formula        *FormulaPageProperty `json:"formula,omitempty"`
	Relation       []IDWrap             `json:"relation,omitempty"`
	Rollup         *RollupPageProperty  `json:"rollup,omitempty"`
	CreatedTime    *DateTimeWrap        `json:"created_time,omitempty"`
	CreatedBy      *User                `json:"created_by,omitempty"`
	LastEditedTime *DateTimeWrap        `json:"last_edited_time,omitempty"`
	LastEditedBy   *User                `json:"last_edited_by,omitempty"`
}

type SelectPageProperty struct {
	Name  string        `json:"name"`
	ID    string        `json:"id"`
	Color PropertyColor `json:"color"`
}

type DatePageProperty struct {
	Start    DateTimeWrap  `json:"start"`
	End      *DateTimeWrap `json:"end,omitempty"`
	TimeZone *string       `json:"time_zone,omitempty"`
}

// type FormulaPageProperty struct {
// 	ID   string         `json:"id"`
// 	Type DBPropertyType `json:"type"`
// 	Formula FormulaPropertyDescription `json:"formula"`
// }

type FormulaPageProperty struct {
	Type           PagePropertyFormulaType `json:"type"`
	StringFormula  *string                 `json:"string,omitempty"`
	BooleanFormula *bool                   `json:"boolean,omitempty"`
	NumberFormula  *float64                `json:"number,omitempty"`
	DateFormula    *DatePageProperty       `json:"date,omitempty"`
}

type PagePropertyFormulaType int

const (
	NoFormulaType PagePropertyFormulaType = iota
	StringFormulaType
	NumberFormulaType
	BoolFormulaType
	DateFormuleType
)

var PagePropertyFormulaTypeToString = map[PagePropertyFormulaType]string{
	StringFormulaType: "string",
	NumberFormulaType: "number",
	BoolFormulaType:   "boolean",
	DateFormuleType:   "date",
}

var StringToPagePropertyFormulaType = map[string]PagePropertyFormulaType{
	"string":  StringFormulaType,
	"number":  NumberFormulaType,
	"boolean": BoolFormulaType,
	"date":    DateFormuleType,
}

func (p PagePropertyFormulaType) MarshalJSON() ([]byte, error) {
	var b = bytes.NewBufferString(`"`)
	b.WriteString(PagePropertyFormulaTypeToString[p])
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (p *PagePropertyFormulaType) UnmarshalJSON(b []byte) error {
	var s = ""
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	res, ok := StringToPagePropertyFormulaType[s]
	if !ok {
		return fmt.Errorf("%v isn't enum value in %T", res, p)
	}
	*p = res
	return nil
}

type RollupPageProperty struct {
	Type     RollupPropertyType `json:"type"`
	Number   *float64           `json:"number,omitempty"`
	Date     *DatePageProperty  `json:"date,omitempty"`
	String   *string            `json:"string,omitempty"`
	Array    []RollupArrayValue `json:"array,omitempty"`
	Function RollupFunction     `json:"function"`
}

type RollupArrayValue struct {
	Type DBPropertyType `json:"type"`

	Title          []RichText           `json:"title,omitempty"`
	RichText       []RichText           `json:"rich_text,omitempty"`
	Number         *float64             `json:"number,omitempty"`
	Select         *SelectPageProperty  `json:"select,omitempty"`
	MultiSelect    []SelectPageProperty `json:"multi_select,omitempty"`
	Date           *DatePageProperty    `json:"date,omitempty"`
	People         []User               `json:"people,omitempty"`
	Files          []FileDescriptor     `json:"files,omitempty"`
	Checkbox       *bool                `json:"checkbox,omitempty"`
	URL            *string              `json:"url,omitempty"`
	Email          *string              `json:"email,omitempty"`
	PhoneNumber    *string              `json:"phone_number,omitempty"`
	Formula        *FormulaPageProperty `json:"formula,omitempty"`
	Relation       []IDWrap             `json:"relation,omitempty"`
	CreatedTime    *DateTimeWrap        `json:"created_time,omitempty"`
	CreatedBy      *User                `json:"created_by,omitempty"`
	LastEditedTime *DateTimeWrap        `json:"last_edited_time,omitempty"`
	LastEditedBy   *User                `json:"last_edited_by,omitempty"`
}

type RollupPropertyType int

const (
	NoRollupPropertyType RollupPropertyType = iota
	NumberRollupPropertyType
	DateRollupPropertyType
	ArrayRollupPropertyType
	UnsupportedRollupPropertyType
	IncompleteRollupPropertyType
)

var RollupPropertyTypeToString = map[RollupPropertyType]string{
	NumberRollupPropertyType:      "number",
	DateRollupPropertyType:        "date",
	ArrayRollupPropertyType:       "array",
	UnsupportedRollupPropertyType: "unsupported",
	IncompleteRollupPropertyType:  "incomplete",
}

var StringToRollupPropertyType = map[string]RollupPropertyType{
	"number":      NumberRollupPropertyType,
	"date":        DateRollupPropertyType,
	"array":       ArrayRollupPropertyType,
	"unsupported": UnsupportedRollupPropertyType,
	"incomplete":  IncompleteRollupPropertyType,
}

func (p RollupPropertyType) MarshalJSON() ([]byte, error) {
	var b = bytes.NewBufferString(`"`)
	b.WriteString(RollupPropertyTypeToString[p])
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (p *RollupPropertyType) UnmarshalJSON(b []byte) error {
	var s = ""
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	res, ok := StringToRollupPropertyType[s]
	if !ok {
		return fmt.Errorf("%v isn't enum value in %T", res, p)
	}
	*p = res
	return nil
}

type CreatePageRq struct {
	ID string
	Properties PageProperties
	// TODO children
	Icon *IconDescriptor
	Cover *FileDescriptor
}

func (rq CreatePageRq)ValidateRequest() error {

	buff := strings.Builder{}

	if rq.ID == "" {
		buff.WriteString(fmt.Sprintf("CreatePageRq.ID couldn't be empty\n"))
	}

	if rq.Properties == nil {
		buff.WriteString("CreatePageRq.Properties couldn't be nil\n")
	}

	if rq.Icon != nil {
		err := rq.Icon.ValidateRequest()
		if err != nil {
			buff.WriteString(err.Error())
		}
	}

	if rq.Cover != nil {
		err := rq.Cover.ValidateRequest()
		if err != nil {
			buff.WriteString(err.Error())
		}
	}

	return errors.New(buff.String())
}
