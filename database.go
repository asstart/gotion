package gotion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Database struct {
	Object         string         `json:"object"`
	Id             string         `json:"id"`
	CreatedTime    string         `json:"created_time"`
	CreatedBy      User           `json:"created_by"`
	LastEditedTime string         `json:"last_edited_time"`
	LastEditedBy   User           `json:"last_edited_by"`
	Title          []RichText     `json:"title"`
	Icon           FileDescriptor `json:"icon"`
	Cover          FileDescriptor `json:"cover"`
	Properties     DBProperties   `json:"properties"`
	ParentType     DBParent       `json:"parent"`
	URL            string         `json:"url"`
	Archived       bool           `json:"archived"`
}

type DBParent struct {
	Type      string `json:"type"`
	PageID    string `json:"page_id,omitempty"`
	Workspace *bool  `json:"workspace,omitempty"`
}

type DBProperties map[string]DBProperty

type DBProperty struct {
	Id   string     `json:"id,omitempty"`
	Type DBPropType `json:"type,omitempty"`
	Name string     `json:"name,omitempty"`

	Title          interface{}         `json:"title,omitempty"`
	RichText       interface{}         `json:"rich_text,omitempty"`
	Number         *DBNumberProperty   `json:"number,omitempty"`
	Select         *DBSelectProperties `json:"select,omitempty"`
	MultiSelect    *DBSelectProperties `json:"multi_select,omitempty"`
	Date           interface{}         `json:"date,omitempty"`
	People         interface{}         `json:"people,omitempty"`
	Files          interface{}         `json:"files,omitempty"`
	Checkbox       interface{}         `json:"checkbox,omitempty"`
	URL            interface{}         `json:"url,omitempty"`
	Email          interface{}         `json:"email,omitempty"`
	PhoneNumber    interface{}         `json:"phone_number,omitempty"`
	Formula        *DBFormulatProperty `json:"formula,omitempty"`
	Relation       *DBRelationProperty `json:"relation,omitempty"`
	Rollup         *Rollup             `json:"rollup,omitempty"`
	CreatedTime    interface{}         `json:"created_time,omitempty"`
	CreatedBy      interface{}         `json:"created_by,omitempty"`
	LastEditedTime interface{}         `json:"last_edited_time,omitempty"`
	LastEditedBy   interface{}         `json:"last_edited_by,omitempty"`
}

type DBNumberProperty struct {
	Format NumberConfigFormat `json:"format"`
}

type DBSelectProperties struct {
	Options []DBSelectProperty `json:"options"`
}

type DBSelectProperty struct {
	Name  string `json:"name"`
	Id    string `json:"id,omitempty"`
	Color Color  `json:"color"`
}

type DBFormulatProperty struct {
	Expression string `json:"expression"`
}

type DBRelationProperty struct {
	DBId               string `json:"database_id"`
	SyncedPropertyName string `json:"synced_property_name,omitempty"`
	SyncedPropertyId   string `json:"synced_property_id,omitempty"`
}

type DBRollupProperty struct {
	RelationPropertyName string `json:"relation_property_name"`
	RelationPropertyId   string `json:"relation_property_id"`
	RollupPropertyName   string `json:"rollup_property_name"`
	RollupPropertyId     string `json:"rollup_property_id"`
	Function             string `json:"function"`
}

type DatabaseQuery struct {
	Filter      *Filter `json:"filter,omitempty"`
	Sorts       []Sort  `json:"sorts,omitempty"`
	StartCursor string  `json:"start_cursor,omitempty"`
	PageSize    int     `json:"page_size,omitempty"`
}

type Sort struct {
	Property  string `json:"property,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Direction string `json:"direction"`
}

//Add filter terst with timestamp
type Filter struct {
	Property  string              `json:"property,omitempty"`
	Timestamp TimestampFilterType `json:"timestamp,omitempty"`

	RichText     *TextCondition        `json:"rich_text,omitempty"`
	Title        *TextCondition        `json:"title,omitempty"`
	URL          *TextCondition        `json:"url,omitempty"`
	Email        *TextCondition        `json:"email,omitempty"`
	PhoneNumber  *TextCondition        `json:"phone_number,omitempty"`
	Number       *NumberCondition      `json:"number,omitempty"`
	Checkbox     *CheckboxCondition    `json:"checkbox,omitempty"`
	Select       *SelectCondition      `json:"select,omitempty"`
	MultiSelect  *MultiSelectCondition `json:"multi_select,omitempty"`
	File         *FileCondition        `json:"files,omitempty"`
	Relation     *RelationCondition    `json:"relation,omitempty"`
	Date         *DateCondition        `json:"date,omitempty"`
	People       *PeopleCondition      `json:"people,omitempty"`
	CreatedBy    *PeopleCondition      `json:"created_by,omitempty"`
	LastEditedBy *PeopleCondition      `json:"last_edited_by,omitempty"`
	Formula      *FormulaCondition     `json:"formula,omitempty"`
	Rollup       *RollupCondition      `json:"rollup,omitempty"`

	And []Filter `json:"and,omitempty"`
	Or  []Filter `json:"or,omitempty"`
}

type TextCondition struct {
	Equals        string `json:"equals,omitempty"`
	DoesntEqual   string `json:"does_not_equal,omitempty"`
	Contains      string `json:"contains,omitempty"`
	DoesntContain string `json:"does_not_contain,omitempty"`
	StartsWith    string `json:"starts_with,omitempty"`
	EndsWith      string `json:"ends_with,omitempty"`
	IsEmpty       bool   `json:"is_empty,omitempty"`     //In doc written "only true", test what happened if false will be passed here
	IsNotEmpty    bool   `json:"is_not_empty,omitempty"` //In doc written "only true", test what happened if false will be passed here
}

//If we pass 0 as value, fields will be erased completely from result json.
type NumberCondition struct {
	Equals               float64 `json:"equals,omitempty"`
	DoesntEqual          float64 `json:"does_not_equal,omitempty"`
	GreaterThan          float64 `json:"greater_than,omitempty"`
	LessThan             float64 `json:"less_than,omitempty"`
	GreaterThanOrEqualTo float64 `json:"greater_than_or_equal_to,omitempty"`
	LessThanOrEqualTo    float64 `json:"less_than_or_equal_to,omitempty"`
	IsEmpty              bool    `json:"is_empty,omitempty"`
	IsNotEmpty           bool    `json:"is_not_empty,omitempty"`
}

//by default go treats bool_field=false as empty
//and if tag omitempty is present, bool_field=false will be erased from result json
//that's why I use *bool here, instead bool.
//https://github.com/golang/go/issues/13284
type CheckboxCondition struct {
	Equals      *bool `json:"equals,omitempty"` //Doesnt work with bool type, need check
	DoesntEqual *bool `json:"does_not_equal,omitempty"`
}

type SelectCondition struct {
	Equals      string `json:"equals,omitempty"`
	DoesntEqual string `json:"does_not_equal,omitempty"`
	IsEmpty     bool   `json:"is_empty,omitempty"`
	IsNotEmpty  bool   `json:"is_not_empty,omitempty"`
}

type MultiSelectCondition struct {
	Contains      string `json:"contains,omitempty"`
	DoesntContain string `json:"does_not_contain,omitempty"`
	IsEmpty       bool   `json:"is_empty,omitempty"`
	IsNotEmpty    bool   `json:"is_not_empty,omitempty"`
}

type DateCondition struct {
	//should be a valid ISO 8601 date string
	Equals string `json:"equals,omitempty"`
	//should be a valid ISO 8601 date string
	Before string `json:"before,omitempty"`
	//should be a valid ISO 8601 date string
	After string `json:"after,omitempty"`
	//should be a valid ISO 8601 date string
	OnOrBefore string `json:"on_or_before,omitempty"`
	//should be a valid ISO 8601 date string
	OnOrAfter  string      `json:"on_or_after,omitempty"`
	PastWeek   interface{} `json:"past_week,omitempty"` //This values should be part of json despite it empty. Test it
	PastMonth  interface{} `json:"past_month,omitempty"`
	PastYear   interface{} `json:"past_year,omitempty"`
	NextWeek   interface{} `json:"next_week,omitempty"`
	NextMonth  interface{} `json:"next_month,omitempty"`
	NextYear   interface{} `json:"next_year,omitempty"`
	IsEmpty    bool        `json:"is_empty,omitempty"`
	IsNotEmpty bool        `json:"is_not_empty,omitempty"`
}

type PeopleCondition struct {
	IsEmpty    bool `json:"is_empty,omitempty"`
	IsNotEmpty bool `json:"is_not_empty,omitempty"`
	//Should be valid uuid
	Contains string `json:"contains,omitempty"`
	//Should be valid uuid
	DoesntContains string `json:"does_not_contain,omitempty"`
}

type FileCondition struct {
	IsEmpty    bool `json:"is_empty,omitempty"`
	IsNotEmpty bool `json:"is_not_empty,omitempty"`
}

type RelationCondition struct {
	IsEmpty    bool `json:"is_empty,omitempty"`
	IsNotEmpty bool `json:"is_not_empty,omitempty"`
	//Should be valid uuid
	Contains string `json:"contains,omitempty"`
	//Should be valid uuid
	DoesntContains string `json:"does_not_contain,omitempty"`
}

type RollupCondition struct {
	Any   *RollupProperty `json:"any,omitempty"`
	Every *RollupProperty `json:"every,omitempty"`
	None  *RollupProperty `json:"none,omitempty"`

	Number *NumberCondition `json:"number,omitempty"`
	Date   *DateCondition   `json:"date,omitempty"`
}

type RollupProperty struct {
	RichText    *TextCondition        `json:"rich_text,omitempty"`
	Title       *TextCondition        `json:"title,omitempty"`
	URL         *TextCondition        `json:"url,omitempty"`
	Email       *TextCondition        `json:"email,omitempty"`
	PhoneNumber *TextCondition        `json:"phone_number,omitempty"`
	Number      *NumberCondition      `json:"number,omitempty"`
	CheckBox    *CheckboxCondition    `json:"checkbox,omitempty"`
	Select      *SelectCondition      `json:"select,omitempty"`
	MultiSelect *MultiSelectCondition `json:"multi_select,omitempty"`
	File        *FileCondition        `json:"files,omitempty"`
	Relation    *RelationCondition    `json:"relation,omitempty"`
	Date        *DateCondition        `json:"date,omitempty"`
}

type FormulaCondition struct {
	String   *TextCondition     `json:"string,omitempty"`
	Checkbox *CheckboxCondition `json:"checkbox,omitempty"`
	Number   *NumberCondition   `json:"number,omitempty"`
	Date     *DateCondition     `json:"date,omitempty"`
}

type CreateDB struct {
	Parent     DBParent     `json:"parent"`
	Title      []RichText   `json:"title,omitempty"`
	Properties DBProperties `json:"properties"`
}

type UpdateDB struct {
	Title      []RichText   `json:"title,omitempty"`
	Properties DBProperties `json:"properties,omitempty"`
}

type NumberConfigFormat int

const (
	Number NumberConfigFormat = iota
	NumberWithCommas
	Percent
	Dollar
	CanadianDollar
	Euro
	Pound
	Yen
	Ruble
	Rupee
	Won
	Yuan
	Real
	Lira
	Rupiah
	Franc
	HongKondDollar
	NewZellanDollar
	Krona
	NorwegianKrone
	MexicanPeso
	Rand
	NewTaiwanDollar
	DanishKrone
	Zloty
	Baht
	Forint
	Koruna
	Shekel
	ChileanPeso
	PhilippinePeso
	Dirham
	ColumbianPeso
	Riyal
	Ringgit
	Leu
	ArgentinePeso
	UruguayanPeso
)

var NumberFormatToString = map[NumberConfigFormat]string{
	Number:           "number",
	NumberWithCommas: "number_with_commas",
	Percent:          "percent",
	Dollar:           "dollar",
	CanadianDollar:   "canadian_dollar",
	Euro:             "euro",
	Pound:            "pound",
	Yen:              "yen",
	Ruble:            "ruble",
	Rupee:            "rupee",
	Won:              "won",
	Yuan:             "yuan",
	Real:             "real",
	Lira:             "lira",
	Rupiah:           "rupiah",
	Franc:            "franc",
	HongKondDollar:   "hong_kong_dollar",
	NewZellanDollar:  "new_zealand_dollar",
	Krona:            "krona",
	NorwegianKrone:   "norwegian_krone",
	MexicanPeso:      "mexican_peso",
	Rand:             "rand",
	NewTaiwanDollar:  "new_taiwan_dollar",
	DanishKrone:      "danish_krone",
	Zloty:            "zloty",
	Baht:             "baht",
	Forint:           "forint",
	Koruna:           "koruna",
	Shekel:           "shekel",
	ChileanPeso:      "chilean_peso",
	PhilippinePeso:   "philippine_peso",
	Dirham:           "dirham",
	ColumbianPeso:    "colombian_peso",
	Riyal:            "riyal",
	Ringgit:          "ringgit",
	Leu:              "leu",
	ArgentinePeso:    "argentine_peso",
	UruguayanPeso:    "uruguayan_peso",
}

var StringToNumberFormat = map[string]NumberConfigFormat{
	"number":             Number,
	"number_with_commas": NumberWithCommas,
	"percent":            Percent,
	"dollar":             Dollar,
	"canadian_dollar":    CanadianDollar,
	"euro":               Euro,
	"pound":              Pound,
	"yen":                Yen,
	"ruble":              Ruble,
	"rupee":              Rupee,
	"won":                Won,
	"yuan":               Yuan,
	"real":               Real,
	"lira":               Lira,
	"rupiah":             Rupiah,
	"franc":              Franc,
	"hong_kong_dollar":   HongKondDollar,
	"new_zealand_dollar": NewZellanDollar,
	"krona":              Krona,
	"norwegian_krone":    NorwegianKrone,
	"mexican_peso":       MexicanPeso,
	"rand":               Rand,
	"new_taiwan_dollar":  NewTaiwanDollar,
	"danish_krone":       DanishKrone,
	"zloty":              Zloty,
	"baht":               Baht,
	"forint":             Forint,
	"koruna":             Koruna,
	"shekel":             Shekel,
	"chilean_peso":       ChileanPeso,
	"philippine_peso":    PhilippinePeso,
	"dirham":             Dirham,
	"colombian_peso":     ColumbianPeso,
	"riyal":              Riyal,
	"ringgit":            Ringgit,
	"leu":                Leu,
	"argentine_peso":     ArgentinePeso,
	"uruguayan_peso":     UruguayanPeso,
}

func (p *NumberConfigFormat) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, v)
	if err != nil {
		return err
	}
	res, ok := StringToNumberFormat[v]
	if !ok {
		return fmt.Errorf("%v isn't enum value", res)
	}
	p = &res
	return nil
}

func (p *NumberConfigFormat) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(NumberFormatToString[*p])
	b.WriteString(`"`)
	return b.Bytes(), nil
}

type DBPropType int

const (
	PropTypeTitle DBPropType = iota
	PropTypeRichText
	PropTypeNumber
	PropTypeSelect
	PropTypeMultiSelect
	PropTypeDate
	PropTypePeople
	PropTypeFiles
	PropTypeCheckbox
	PropTypeURL
	PropTypeEmail
	PropTypePhoneNumber
	PropTypeFormula
	PropTypeRelation
	PropTypeRollup
	PropTypeCreatedTime
	PropTypeCreatedBy
	PropTypeLastEditedTime
	PropTypeLastEditedBy
)

var PropTypeToString = map[DBPropType]string{
	PropTypeTitle:          "title",
	PropTypeRichText:       "rich_text",
	PropTypeNumber:         "number",
	PropTypeSelect:         "select",
	PropTypeMultiSelect:    "multi_select",
	PropTypeDate:           "date",
	PropTypePeople:         "people",
	PropTypeFiles:          "files",
	PropTypeCheckbox:       "checkbox",
	PropTypeURL:            "url",
	PropTypeEmail:          "email",
	PropTypePhoneNumber:    "phone_number",
	PropTypeFormula:        "formula",
	PropTypeRelation:       "relation",
	PropTypeRollup:         "rollup",
	PropTypeCreatedTime:    "created_time",
	PropTypeCreatedBy:      "created_by",
	PropTypeLastEditedTime: "last_edited_time",
	PropTypeLastEditedBy:   "last_edited_by",
}

var StringToPropType = map[string]DBPropType{
	"title":            PropTypeTitle,
	"rich_text":        PropTypeRichText,
	"number":           PropTypeNumber,
	"select":           PropTypeSelect,
	"multi_select":     PropTypeMultiSelect,
	"date":             PropTypeDate,
	"people":           PropTypePeople,
	"files":            PropTypeFiles,
	"checkbox":         PropTypeCheckbox,
	"url":              PropTypeURL,
	"email":            PropTypeEmail,
	"phone_number":     PropTypePhoneNumber,
	"formula":          PropTypeFormula,
	"relation":         PropTypeRelation,
	"rollup":           PropTypeRollup,
	"created_time":     PropTypeCreatedTime,
	"created_by":       PropTypeCreatedBy,
	"last_edited_time": PropTypeLastEditedTime,
	"last_edited_by":   PropTypeLastEditedBy,
}

func (p *DBPropType) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, v)
	if err != nil {
		return err
	}
	res, ok := StringToPropType[v]
	if !ok {
		return fmt.Errorf("%v isn't enum value", res)
	}
	p = &res
	return nil
}

func (p *DBPropType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(PropTypeToString[*p])
	b.WriteString(`"`)
	return b.Bytes(), nil
}

type Color int

const (
	Default Color = iota
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

var ColorToString = map[Color]string{
	Default: "default",
	Gray:    "gray",
	Brown:   "brown",
	Orange:  "orange",
	Yellow:  "yellow",
	Green:   "green",
	Blue:    "blue",
	Purple:  "purple",
	Pink:    "pink",
	Red:     "red",
}

var StringToColor = map[string]Color{
	"default": Default,
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

func (p *Color) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, v)
	if err != nil {
		return err
	}
	res, ok := StringToColor[v]
	if !ok {
		return fmt.Errorf("%v isn't enum value", res)
	}
	p = &res
	return nil
}

func (p *Color) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(ColorToString[*p])
	b.WriteString(`"`)
	return b.Bytes(), nil
}

type RollupFunction int

const (
	CountAll RollupFunction = iota
	CountValues
	CountUniqueValues
	CountEmpty
	CountNotEmpty
	PercentEmpty
	PercentNotEmpty
	Sum
	Average
	Median
	Min
	Max
	Range
	ShowOriginal
)

var RollupFunctionToString = map[RollupFunction]string{
	CountAll:          "count_all",
	CountValues:       "count_values",
	CountUniqueValues: "count_unique_values",
	CountEmpty:        "count_empty",
	CountNotEmpty:     "count_not_empty",
	PercentEmpty:      "percent_empty",
	PercentNotEmpty:   "percent_not_empty",
	Sum:               "sum",
	Average:           "average",
	Median:            "median",
	Min:               "min",
	Max:               "max",
	Range:             "range",
	ShowOriginal:      "show_origina",
}

var StringToRollupFunction = map[string]RollupFunction{
	"count_all":           CountAll,
	"count_values":        CountValues,
	"count_unique_values": CountUniqueValues,
	"count_empty":         CountEmpty,
	"count_not_empty":     CountNotEmpty,
	"percent_empty":       PercentEmpty,
	"percent_not_empty":   PercentNotEmpty,
	"sum":                 Sum,
	"average":             Average,
	"median":              Median,
	"min":                 Min,
	"max":                 Max,
	"range":               Range,
	"show_origina":        ShowOriginal,
}

func (p *RollupFunction) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, v)
	if err != nil {
		return err
	}
	res, ok := StringToRollupFunction[v]
	if !ok {
		return fmt.Errorf("%v isn't enum value", res)
	}
	p = &res
	return nil
}

func (p *RollupFunction) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(RollupFunctionToString[*p])
	b.WriteString(`"`)
	return b.Bytes(), nil
}

type TimestampFilterType int

const (
	CreatedTime TimestampFilterType = iota
	LastEditedTime
)

var TimestampFilterTypeToString = map[TimestampFilterType]string{
	CreatedTime:    "created_time",
	LastEditedTime: "last_edited_time",
}

var StringToTimestampFilterType = map[string]TimestampFilterType{
	"created_time":     CreatedTime,
	"last_edited_time": LastEditedTime,
}

func (p *TimestampFilterType) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, v)
	if err != nil {
		return err
	}
	res, ok := StringToTimestampFilterType[v]
	if !ok {
		return fmt.Errorf("%v isn't enum value", res)
	}
	p = &res
	return nil
}

func (p *TimestampFilterType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(TimestampFilterTypeToString[*p])
	b.WriteString(`"`)
	return b.Bytes(), nil
}
