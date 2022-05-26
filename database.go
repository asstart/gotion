package gotion

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
	Type      *string `json:"type"`
	PageID    *string `json:"page_id,omitempty"`
	Workspace *bool   `json:"workspace,omitempty"`
}

type DBProperties map[string]DBProperty

type DBProperty struct {
	Id   string
	Type string
	Name string

	Title          struct{}           `json:"title,omitempty"`
	RichText       struct{}           `json:"rich_text,omitempty"`
	Number         DBNumberProperty   `json:"number,omitempty"`
	Select         []DBSelectProperty `json:"select,omitempty"`
	MultiSelect    []DBSelectProperty `json:"multi_select,omitempty"`
	Date           struct{}           `json:"date,omitempty"`
	People         struct{}           `json:"people,omitempty"`
	Files          struct{}           `json:"files,omitempty"`
	Checkbox       struct{}           `json:"checkbox,omitempty"`
	URL            struct{}           `json:"url,omitempty"`
	Email          struct{}           `json:"email,omitempty"`
	PhoneNumber    struct{}           `json:"phone_number,omitempty"`
	Formula        DBFormulatProperty `json:"formula,omitempty"`
	Relation       DBRelationProperty `json:"relation,omitempty"`
	Rollup         Rollup             `json:"rollup,omitempty"`
	CreatedTime    struct{}           `json:"created_time,omitempty"`
	CreatedBy      struct{}           `json:"created_by,omitempty"`
	LastEditedTime struct{}           `json:"last_edited_time,omitempty"`
	LastEditedBy   struct{}           `json:"last_edited_by,omitempty"`
}

type DBNumberProperty struct {
	Format string `json:"format"`
}

type DBSelectProperty struct {
	Name  string `json:"name"`
	Id    string `json:"id"`
	Color string `json:"color"`
}

type DBFormulatProperty struct {
	Expression string `json:"expression"`
}

type DBRelationProperty struct {
	DBId               string `json:"database_id"`
	SyncedPropertyName string `json:"synced_property_name"`
	SyncedPropertyId   string `json:"synced_property_id"`
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

//Add rollup, formula, people
type Filter struct {
	Property  string `json:"property,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`

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
	People       *PeopleCondition       `json:"people,omitempty"`
	CreatedBy    *PeopleCondition       `json:"created_by,omitempty"`
	LastEditedBy *PeopleCondition       `json:"last_edited_by,omitempty"`

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
	Contains     string `json:"contains,omitempty"`
	DoesntContain string `json:"does_not_contain,omitempty"`
	IsEmpty      bool   `json:"is_empty,omitempty"`
	IsNotEmpty   bool   `json:"is_not_empty,omitempty"`
}

type DateCondition struct {
	Equals     string   `json:"equals,omitempty"`
	Before     string   `json:"before,omitempty"`
	After      string   `json:"after,omitempty"`
	OnOrBefore string   `json:"on_or_before,omitempty"`
	OnOrAfter  string   `json:"on_or_after,omitempty"`
	PastWeek   struct{} `json:"past_week,omitempty"` //This values should be part of json despite it empty. Test it
	PastMonth  struct{} `json:"past_month,omitempty"`
	PastYear   struct{} `json:"past_year,omitempty"`
	NextWeek   struct{} `json:"next_week,omitempty"`
	NextMonth  struct{} `json:"next_month,omitempty"`
	NextYear   struct{} `json:"next_year,omitempty"`
	IsEmpty    bool     `json:"is_empty,omitempty"`
	IsNotEmpty bool     `json:"is_not_empty,omitempty"`
}

type PeopleCondition struct {
	IsEmpty        bool   `json:"is_empty,omitempty"`
	IsNotEmpty     bool   `json:"is_not_empty,omitempty"`
	Contains       string `json:"contains,omitempty"`
	DoesntContains string `json:"does_not_contain,omitempty"`
}

type FileCondition struct {
	IsEmpty    bool `json:"is_empty,omitempty"`
	IsNotEmpty bool `json:"is_not_empty,omitempty"`
}

type RelationCondition struct {
	IsEmpty        bool   `json:"is_empty,omitempty"`
	IsNotEmpty     bool   `json:"is_not_empty,omitempty"`
	Contains       string `json:"contains,omitempty"`
	DoesntContains string `json:"does_not_contain,omitempty"`
}

type RollupCondition struct {
	Any   RollupProperty `json:"any,omitempty"`
	Every RollupProperty `json:"every,omitempty"`
	None  RollupProperty `json:"none,omitempty"`

	Number NumberCondition `json:"number,omitempty"`
	Date   DateCondition   `json:"date,omitempty"`
}

type RollupProperty struct {
	RichText    TextCondition        `json:"rich_text,omitempty"`
	Title       TextCondition        `json:"title,omitempty"`
	URL         TextCondition        `json:"url,omitempty"`
	Email       TextCondition        `json:"email,omitempty"`
	PhoneNumber TextCondition        `json:"phone_number,omitempty"`
	Number      NumberCondition      `json:"number,omitempty"`
	CheckBox    CheckboxCondition    `json:"checkbox,omitempty"`
	Select      SelectCondition      `json:"select,omitempty"`
	MultiSelect MultiSelectCondition `json:"multi_select,omitempty"`
	File        FileCondition        `json:"files,omitempty"`
	Relation    RelationCondition    `json:"relation,omitempty"`
	Date        DateCondition        `json:"date,omitempty"`
}

type FormulaCondition struct {
	String   TextCondition     `json:"string,omitempty"`
	Checkbox CheckboxCondition `json:"checkbox,omitempty"`
	Number   NumberCondition   `json:"number,omitempty"`
	Date     DateCondition     `json:"date,omitempty"`
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
