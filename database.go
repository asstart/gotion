package gotion

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type DB struct {
	Object         string       `json:"object"`
	ID             string       `json:"id"`
	CreatedTime    DateTimeWrap `json:"created_time"`
	CreatedBy      User         `json:"created_by"`
	LastEditedTime DateTimeWrap `json:"last_edited_time"`
	LastEditedBy   User         `json:"last_edited_by"`
	Title          []RichText   `json:"title"`
	//Doesn't exist in doc, but exists in response
	Description []RichText      `json:"description"`
	Icon        *IconDescriptor `json:"icon,omitempty"`
	Cover       *FileDescriptor `json:"cover,omitempty"`
	Properties  DBProperties    `json:"properties"`
	Parent      DBParent        `json:"parent"`
	URL         string          `json:"url"`
	Archived    bool            `json:"archived"`
}

type DBParent struct {
	Type      DBParentType `json:"type"`
	PageID    string       `json:"page_id,omitempty"`
	Workspace bool         `json:"workspace,omitempty"`
}

type DBParentType int

const (
	NoDBParentType DBParentType = iota
	PageDBParentType
	WorksapceDBParentType
)

var DBParentTypeToString = map[DBParentType]string{
	PageDBParentType:      "page_id",
	WorksapceDBParentType: "workspace",
}

var StringToDBParentType = map[string]DBParentType{
	"page_id":   PageDBParentType,
	"workspace": WorksapceDBParentType,
}

func (dbp DBParent) ValidateRequest() error {
	var buff = strings.Builder{}
	if dbp.Type == NoDBParentType {
		buff.WriteString("DBParent.Type is empty\n")
	}
	if dbp.PageID == "" {
		buff.WriteString("DBParent.PageID is empty\n")
	}
	var final = buff.String()
	if final == "" {
		return nil
	}
	return errors.New(buff.String())
}

func (p *DBParentType) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	res, ok := StringToDBParentType[v]
	if !ok {
		return fmt.Errorf("%v isn't enum value in %T", res, p)
	}
	*p = res
	return nil
}

func (p DBParentType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(DBParentTypeToString[p])
	b.WriteString(`"`)
	return b.Bytes(), nil
}

type DBProperties map[string]DBProperty

type DBProperty struct {
	ID   string         `json:"id,omitempty"`
	Type DBPropertyType `json:"type,omitempty"`
	Name string         `json:"name,omitempty"`

	Title          *DBDefaultProperty  `json:"title,omitempty"`
	RichText       *DBDefaultProperty  `json:"rich_text,omitempty"`
	Number         *DBNumberProperty   `json:"number,omitempty"`
	Select         *DBSelectProperties `json:"select,omitempty"`
	MultiSelect    *DBSelectProperties `json:"multi_select,omitempty"`
	Date           *DBDefaultProperty  `json:"date,omitempty"`
	People         *DBDefaultProperty  `json:"people,omitempty"`
	Files          *DBDefaultProperty  `json:"files,omitempty"`
	Checkbox       *DBDefaultProperty  `json:"checkbox,omitempty"`
	URL            *DBDefaultProperty  `json:"url,omitempty"`
	Email          *DBDefaultProperty  `json:"email,omitempty"`
	PhoneNumber    *DBDefaultProperty  `json:"phone_number,omitempty"`
	Formula        *DBFormulatProperty `json:"formula,omitempty"`
	Relation       *DBRelationProperty `json:"relation,omitempty"`
	Rollup         *RollupProperty     `json:"rollup,omitempty"` // TODO need fix this
	CreatedTime    *DBDefaultProperty  `json:"created_time,omitempty"`
	CreatedBy      *DBDefaultProperty  `json:"created_by,omitempty"`
	LastEditedTime *DBDefaultProperty  `json:"last_edited_time,omitempty"`
	LastEditedBy   *DBDefaultProperty  `json:"last_edited_by,omitempty"`
}

type DBDefaultProperty struct{}

type DBNumberProperty struct {
	Format NumberPropertyFormat `json:"format"`
}

type DBSelectProperties struct {
	Options []DBSelectProperty `json:"options"`
}

type DBSelectProperty struct {
	Name  string        `json:"name"`
	ID    string        `json:"id,omitempty"`
	Color PropertyColor `json:"color"`
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

type QueryDBRq struct {
	Filter      *Filter `json:"filter,omitempty"`
	Sorts       []Sort  `json:"sorts,omitempty"`
	StartCursor string  `json:"start_cursor,omitempty"`
	PageSize    int     `json:"page_size,omitempty"`
}

type QueryDBRs struct {
	Object     string  `json:"object"`
	Results    []Page  `json:"results"`
	NextCursor *string `json:"next_cursor"`
	HasMore    bool    `json:"has_more"`
}

type Sort struct {
	Property string `json:"property,omitempty"`
	/*
		If you will pass both timestamp and property in one sort object
		like this:

		{
			"sorts": [
				{
					"timestamp": "created_time",
					"property": "title",
					"direction": "ascending"
				}
			]
		}

		Notiion will accept it, but behaviour unpredictable(need to test)
		TODO Think about validtion on exactly one of this
	*/
	Timestamp TimestampConditionType `json:"timestamp,omitempty"`
	Direction SortDirection          `json:"direction"`
}

type SortDirection int

const (
	Ascending SortDirection = iota
	Descending
)

var SortDirectionToString = map[SortDirection]string{
	Ascending:  "ascending",
	Descending: "descending",
}

var StringToSortDirection = map[string]SortDirection{
	"ascending":  Ascending,
	"descending": Descending,
}

func (p *SortDirection) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	res, ok := StringToSortDirection[v]
	if !ok {
		return fmt.Errorf("%v isn't enum value in %T", res, p)
	}
	*p = res
	return nil
}

func (p SortDirection) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(SortDirectionToString[p])
	b.WriteString(`"`)
	return b.Bytes(), nil
}

//Add filter terst with timestamp
type Filter struct {
	Property  string                 `json:"property,omitempty"`
	Timestamp TimestampConditionType `json:"timestamp,omitempty"`

	/*
		If you will pass several conditions for one propery,
		For example:
		{
			"filter": {
				"property": "dateprop",
				"date": {
					"before": "2033-05-18T03:33:00.1Z",
					"is_not_empty": false
				}
			}
		}

		Notion will handle this as valid request, but only one condition will be applied
	*/

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
//that's why I use *bool hre, instead bool.
//https://github.com/golang/go/issues/13284
type CheckboxCondition struct {
	Equals      *bool `json:"equals,omitempty"`
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
	Equals     *DateTimeWrap      `json:"equals,omitempty"`
	Before     *DateTimeWrap      `json:"before,omitempty"`
	After      *DateTimeWrap      `json:"after,omitempty"`
	OnOrBefore *DateTimeWrap      `json:"on_or_before,omitempty"`
	OnOrAfter  *DateTimeWrap      `json:"on_or_after,omitempty"`
	PastWeek   *DateTimeEmptyWrap `json:"past_week,omitempty"`
	PastMonth  *DateTimeEmptyWrap `json:"past_month,omitempty"`
	PastYear   *DateTimeEmptyWrap `json:"past_year,omitempty"`
	NextWeek   *DateTimeEmptyWrap `json:"next_week,omitempty"`
	NextMonth  *DateTimeEmptyWrap `json:"next_month,omitempty"`
	NextYear   *DateTimeEmptyWrap `json:"next_year,omitempty"`
	IsEmpty    bool               `json:"is_empty,omitempty"`
	IsNotEmpty bool               `json:"is_not_empty,omitempty"`
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

type CreateDBRq struct {
	PageId     string
	Title      []RichText
	Properties DBProperties
	Icon       *IconDescriptor
	Cover      *FileDescriptor
}

func (cdb CreateDBRq) ValidateRequest() error {
	var buff = strings.Builder{}

	if cdb.PageId == "" {
		buff.WriteString("CreateDB.PageId couldn't be nil")
	}

	if cdb.Properties == nil {
		buff.WriteString("CreateDB.Properties couldn't be nil")
	}

	if cdb.Icon != nil {
		err := cdb.Icon.ValidateRequest()
		if err != nil {
			buff.WriteString(err.Error())
		}
	}

	if cdb.Cover != nil {
		err := cdb.Cover.ValidateRequest()
		if err != nil {
			buff.WriteString(err.Error())
		}
	}

	var final = buff.String()
	if final == "" {
		return nil
	}
	return errors.New(buff.String())
}

func (cdb CreateDBRq) MarshalJSON() ([]byte, error) {
	err := cdb.ValidateRequest()

	if err != nil {
		return nil, err
	}

	type DTO struct {
		Parent     *DBParent       `json:"parent"`
		Title      []RichText      `json:"title,omitempty"`
		Properties DBProperties    `json:"properties"`
		Icon       *IconDescriptor `json:"icon,omitempty"`
		Cover      *FileDescriptor `json:"cover,omitempty"`
	}
	dto := DTO{
		Parent: &DBParent{
			Type:   PageDBParentType,
			PageID: cdb.PageId,
		},
		Title:      cdb.Title,
		Properties: cdb.Properties,
		Icon:       cdb.Icon,
		Cover:      cdb.Cover,
	}
	return json.Marshal(dto)
}

type UpdateDBRq struct {
	Title      []RichText      `json:"title,omitempty"`
	Properties DBProperties    `json:"properties,omitempty"`
	Icon       *IconDescriptor `json:"icon,omitempty"`
	Cover      *FileDescriptor `json:"cover,omitempty"`
}

type NumberPropertyFormat int

const (
	NoNumberFormat NumberPropertyFormat = iota
	Number
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

var NumberFormatToString = map[NumberPropertyFormat]string{
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

var StringToNumberFormat = map[string]NumberPropertyFormat{
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

func (p *NumberPropertyFormat) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	res, ok := StringToNumberFormat[v]
	if !ok {
		return fmt.Errorf("%v isn't enum value in %T", res, p)
	}
	*p = res
	return nil
}

func (p NumberPropertyFormat) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(NumberFormatToString[p])
	b.WriteString(`"`)
	return b.Bytes(), nil
}

type DBPropertyType int

const (
	NoDBPropType DBPropertyType = iota
	DBPropTypeTitle
	DBPropTypeRichText
	DBPropTypeNumber
	DBPropTypeSelect
	DBPropTypeMultiSelect
	DBPropTypeDate
	DBPropTypePeople
	DBPropTypeFiles
	DBPropTypeCheckbox
	DBPropTypeURL
	DBPropTypeEmail
	DBPropTypePhoneNumber
	DBPropTypeFormula
	DBPropTypeRelation
	DBPropTypeRollup
	DBPropTypeCreatedTime
	DBPropTypeCreatedBy
	DBPropTypeLastEditedTime
	DBPropTypeLastEditedBy
)

var PropTypeToString = map[DBPropertyType]string{
	DBPropTypeTitle:          "title",
	DBPropTypeRichText:       "rich_text",
	DBPropTypeNumber:         "number",
	DBPropTypeSelect:         "select",
	DBPropTypeMultiSelect:    "multi_select",
	DBPropTypeDate:           "date",
	DBPropTypePeople:         "people",
	DBPropTypeFiles:          "files",
	DBPropTypeCheckbox:       "checkbox",
	DBPropTypeURL:            "url",
	DBPropTypeEmail:          "email",
	DBPropTypePhoneNumber:    "phone_number",
	DBPropTypeFormula:        "formula",
	DBPropTypeRelation:       "relation",
	DBPropTypeRollup:         "rollup",
	DBPropTypeCreatedTime:    "created_time",
	DBPropTypeCreatedBy:      "created_by",
	DBPropTypeLastEditedTime: "last_edited_time",
	DBPropTypeLastEditedBy:   "last_edited_by",
}

var StringToPropType = map[string]DBPropertyType{
	"title":            DBPropTypeTitle,
	"rich_text":        DBPropTypeRichText,
	"number":           DBPropTypeNumber,
	"select":           DBPropTypeSelect,
	"multi_select":     DBPropTypeMultiSelect,
	"date":             DBPropTypeDate,
	"people":           DBPropTypePeople,
	"files":            DBPropTypeFiles,
	"checkbox":         DBPropTypeCheckbox,
	"url":              DBPropTypeURL,
	"email":            DBPropTypeEmail,
	"phone_number":     DBPropTypePhoneNumber,
	"formula":          DBPropTypeFormula,
	"relation":         DBPropTypeRelation,
	"rollup":           DBPropTypeRollup,
	"created_time":     DBPropTypeCreatedTime,
	"created_by":       DBPropTypeCreatedBy,
	"last_edited_time": DBPropTypeLastEditedTime,
	"last_edited_by":   DBPropTypeLastEditedBy,
}

func (p *DBPropertyType) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	res, ok := StringToPropType[v]
	if !ok {
		return fmt.Errorf("%v isn't enum value in %T", res, p)
	}
	*p = res
	return nil
}

func (p DBPropertyType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(PropTypeToString[p])
	b.WriteString(`"`)
	return b.Bytes(), nil
}

type RollupFunction int

const (
	NoRollupFunc RollupFunction = iota
	Count
	CountAll
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
	EarliestDate
	LatestDate
	DateRange
)

var RollupFunctionToString = map[RollupFunction]string{
	Count:             "count",
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
	ShowOriginal:      "show_original",
	EarliestDate:      "earliest_date",
	LatestDate:        "latest_date",
	DateRange:         "date_range",
}

var StringToRollupFunction = map[string]RollupFunction{
	"count":               Count,
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
	"show_original":        ShowOriginal,
	"earliest_date":       EarliestDate,
	"latest_date":         LatestDate,
	"date_range":          DateRange,
}

func (p *RollupFunction) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	res, ok := StringToRollupFunction[v]
	if !ok {
		return fmt.Errorf("%v isn't enum value in %T", res, p)
	}
	*p = res
	return nil
}

func (p RollupFunction) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(RollupFunctionToString[p])
	b.WriteString(`"`)
	return b.Bytes(), nil
}

type TimestampConditionType int

const (
	NoTimestampFilterType TimestampConditionType = iota
	CreatedTime
	LastEditedTime
)

var TimestampFilterTypeToString = map[TimestampConditionType]string{
	CreatedTime:    "created_time",
	LastEditedTime: "last_edited_time",
}

var StringToTimestampFilterType = map[string]TimestampConditionType{
	"created_time":     CreatedTime,
	"last_edited_time": LastEditedTime,
}

func (p *TimestampConditionType) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	res, ok := StringToTimestampFilterType[v]
	if !ok {
		return fmt.Errorf("%v isn't enum value in %T", res, p)
	}
	*p = res
	return nil
}

func (p TimestampConditionType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(TimestampFilterTypeToString[p])
	b.WriteString(`"`)
	return b.Bytes(), nil
}
