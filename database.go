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
