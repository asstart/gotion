package gotion_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/asstart/gotion"
	"github.com/asstart/gotion/utils"
)

type querydbTuple struct {
	Source   *gotion.QuertyDBRq
	Expected string
}

type createdbTuple struct {
	Source   *gotion.CreateDBRq
	Expected string
}

type updatedbTuple struct {
	Source   *gotion.UpdateDBRq
	Expected string
}

type retrievedbTuple struct {
	Source   string
	Expected gotion.Database
}

func TestMarshalEmptyQuery(t *testing.T) {
	source := gotion.QuertyDBRq{}

	expected := "{}"

	actual, err := json.Marshal(source)
	utils.AssertNil(t, err)
	utils.AssertEqualsString(t, expected, string(actual))
}

func TestMarshalSingleLevelFilterQuery(t *testing.T) {
	source := gotion.QuertyDBRq{
		Filter: &gotion.Filter{
			Property: "language",
			RichText: &gotion.TextCondition{
				Equals: "english",
			},
		},
	}

	expected := `
	{
		"filter":{
			"property":"language",
			"rich_text":{
				"equals":"english"
			}
		}
	}
	`

	actual, err := json.Marshal(source)
	utils.AssertNil(t, err)
	utils.AssertEqualsString(
		t,
		utils.Minimise(expected),
		utils.Minimise(string(actual)),
	)
}

func TestMarshal2NestedLevelFilter(t *testing.T) {
	source := gotion.QuertyDBRq{
		Filter: &gotion.Filter{
			And: []gotion.Filter{
				{
					Property: "level1_1",
					Number: &gotion.NumberCondition{
						Equals: 10,
					},
				},
				{
					Property: "level1_2",
					RichText: &gotion.TextCondition{
						Contains: "level1",
					},
				},
				{
					Or: []gotion.Filter{
						{
							Property: "level2_1",
							Number: &gotion.NumberCondition{
								GreaterThan: 15,
							},
						},
						{
							Property: "level2_2",
							Number: &gotion.NumberCondition{
								LessThan: 30,
							},
						},
					},
				},
			},
		},
	}

	expected := `
	{
		"filter":{
			"and":[
				{
					"property":"level1_1",
					"number":{
						"equals":10
					}
				},
				{
					"property":"level1_2",
					"rich_text":{
						"contains":"level1"
					}
				},
				{
					"or":[
						{
							"property":"level2_1",
							"number":{
								"greater_than":15
							}
						},
						{
							"property":"level2_2",
							"number":{
								"less_than":30
							}
						}
					]
				}
			]
		}
	}
	`

	actual, err := json.Marshal(source)
	utils.AssertNil(t, err)
	utils.AssertEqualsString(
		t,
		utils.Minimise(expected),
		utils.Minimise(string(actual)),
	)

}

func TestMarshalRichTextQuery(t *testing.T) {
	type tuple struct {
		Source   *gotion.QuertyDBRq
		Expected string
	}
	se := []tuple{
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "text", RichText: &gotion.TextCondition{Equals: "some_string"},
				}},
			`{"filter":{"property":"text","rich_text":{"equals":"some_string"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "text", RichText: &gotion.TextCondition{DoesntEqual: "some_string"},
				}},
			`{"filter":{"property":"text","rich_text":{"does_not_equal":"some_string"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "text", RichText: &gotion.TextCondition{Contains: "some_string"},
				}},
			`{"filter":{"property":"text","rich_text":{"contains":"some_string"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "text", RichText: &gotion.TextCondition{DoesntContain: "some_string"},
				}},
			`{"filter":{"property":"text","rich_text":{"does_not_contain":"some_string"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "text", RichText: &gotion.TextCondition{StartsWith: "some_string"},
				}},
			`{"filter":{"property":"text","rich_text":{"starts_with":"some_string"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "text", RichText: &gotion.TextCondition{EndsWith: "some_string"},
				}},
			`{"filter":{"property":"text","rich_text":{"ends_with":"some_string"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "text", RichText: &gotion.TextCondition{IsEmpty: true},
				}},
			`{"filter":{"property":"text","rich_text":{"is_empty": true}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "text", RichText: &gotion.TextCondition{IsNotEmpty: true},
				}},
			`{"filter":{"property":"text","rich_text":{"is_not_empty":true}}}`,
		},
	}

	for _, pair := range se {
		actual, err := json.Marshal(pair.Source)
		utils.AssertNil(t, err)
		utils.AssertEqualsString(
			t,
			utils.Minimise(pair.Expected),
			utils.Minimise(string(actual)),
		)
	}
}

func TestMarshalTitleQuery(t *testing.T) {
	se := []querydbTuple{
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "title", Title: &gotion.TextCondition{Equals: "some_string"},
				}},
			`{"filter":{"property":"title","title":{"equals":"some_string"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "title", Title: &gotion.TextCondition{DoesntEqual: "some_string"},
				}},
			`{"filter":{"property":"title","title":{"does_not_equal":"some_string"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "title", Title: &gotion.TextCondition{Contains: "some_string"},
				}},
			`{"filter":{"property":"title","title":{"contains":"some_string"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "title", Title: &gotion.TextCondition{DoesntContain: "some_string"},
				}},
			`{"filter":{"property":"title","title":{"does_not_contain":"some_string"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "title", Title: &gotion.TextCondition{StartsWith: "some_string"},
				}},
			`{"filter":{"property":"title","title":{"starts_with":"some_string"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "title", Title: &gotion.TextCondition{EndsWith: "some_string"},
				}},
			`{"filter":{"property":"title","title":{"ends_with":"some_string"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "title", Title: &gotion.TextCondition{IsEmpty: true},
				}},
			`{"filter":{"property":"title","title":{"is_empty": true}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "title", Title: &gotion.TextCondition{IsNotEmpty: true},
				}},
			`{"filter":{"property":"title","title":{"is_not_empty":true}}}`,
		},
	}

	for _, pair := range se {
		actual, err := json.Marshal(pair.Source)
		utils.AssertNil(t, err)
		utils.AssertEqualsString(
			t,
			utils.Minimise(pair.Expected),
			utils.Minimise(string(actual)),
		)
	}
}

func TestMarshalURLQuery(t *testing.T) {
	se := []querydbTuple{
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "url", URL: &gotion.TextCondition{Equals: "https://github.com/"},
				}},
			`{"filter":{"property":"url","url":{"equals":"https://github.com/"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "url", URL: &gotion.TextCondition{DoesntEqual: "https://github.com/"},
				}},
			`{"filter":{"property":"url","url":{"does_not_equal":"https://github.com/"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "url", URL: &gotion.TextCondition{Contains: "https://github.com/"},
				}},
			`{"filter":{"property":"url","url":{"contains":"https://github.com/"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "url", URL: &gotion.TextCondition{DoesntContain: "https://github.com/"},
				}},
			`{"filter":{"property":"url","url":{"does_not_contain":"https://github.com/"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "url", URL: &gotion.TextCondition{StartsWith: "https://github.com/"},
				}},
			`{"filter":{"property":"url","url":{"starts_with":"https://github.com/"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "url", URL: &gotion.TextCondition{EndsWith: "https://github.com/"},
				}},
			`{"filter":{"property":"url","url":{"ends_with":"https://github.com/"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "url", URL: &gotion.TextCondition{IsEmpty: true},
				}},
			`{"filter":{"property":"url","url":{"is_empty": true}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "url", URL: &gotion.TextCondition{IsNotEmpty: true},
				}},
			`{"filter":{"property":"url","url":{"is_not_empty":true}}}`,
		},
	}

	for _, pair := range se {
		actual, err := json.Marshal(pair.Source)
		utils.AssertNil(t, err)
		utils.AssertEqualsString(
			t,
			utils.Minimise(pair.Expected),
			utils.Minimise(string(actual)),
		)
	}
}

func TestMarshalEmailQuery(t *testing.T) {
	se := []querydbTuple{
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "email", Email: &gotion.TextCondition{Equals: "test@mail.ru"},
				}},
			`{"filter":{"property":"email","email":{"equals":"test@mail.ru"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "email", Email: &gotion.TextCondition{DoesntEqual: "test@mail.ru"},
				}},
			`{"filter":{"property":"email","email":{"does_not_equal":"test@mail.ru"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "email", Email: &gotion.TextCondition{Contains: "test@mail.ru"},
				}},
			`{"filter":{"property":"email","email":{"contains":"test@mail.ru"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "email", Email: &gotion.TextCondition{DoesntContain: "test@mail.ru"},
				}},
			`{"filter":{"property":"email","email":{"does_not_contain":"test@mail.ru"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "email", Email: &gotion.TextCondition{StartsWith: "test@mail.ru"},
				}},
			`{"filter":{"property":"email","email":{"starts_with":"test@mail.ru"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "email", Email: &gotion.TextCondition{EndsWith: "test@mail.ru"},
				}},
			`{"filter":{"property":"email","email":{"ends_with":"test@mail.ru"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "email", Email: &gotion.TextCondition{IsEmpty: true},
				}},
			`{"filter":{"property":"email","email":{"is_empty": true}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "email", Email: &gotion.TextCondition{IsNotEmpty: true},
				}},
			`{"filter":{"property":"email","email":{"is_not_empty":true}}}`,
		},
	}

	for _, pair := range se {
		actual, err := json.Marshal(pair.Source)
		utils.AssertNil(t, err)
		utils.AssertEqualsString(
			t,
			utils.Minimise(pair.Expected),
			utils.Minimise(string(actual)),
		)
	}
}

func TestMarshalPhoneNumberQuery(t *testing.T) {
	se := []querydbTuple{
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "phone_number", PhoneNumber: &gotion.TextCondition{Equals: "+19876543210"},
				}},
			`{"filter":{"property":"phone_number","phone_number":{"equals":"+19876543210"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "phone_number", PhoneNumber: &gotion.TextCondition{DoesntEqual: "+19876543210"},
				}},
			`{"filter":{"property":"phone_number","phone_number":{"does_not_equal":"+19876543210"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "phone_number", PhoneNumber: &gotion.TextCondition{Contains: "+19876543210"},
				}},
			`{"filter":{"property":"phone_number","phone_number":{"contains":"+19876543210"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "phone_number", PhoneNumber: &gotion.TextCondition{DoesntContain: "+19876543210"},
				}},
			`{"filter":{"property":"phone_number","phone_number":{"does_not_contain":"+19876543210"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "phone_number", PhoneNumber: &gotion.TextCondition{StartsWith: "+19876543210"},
				}},
			`{"filter":{"property":"phone_number","phone_number":{"starts_with":"+19876543210"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "phone_number", PhoneNumber: &gotion.TextCondition{EndsWith: "+19876543210"},
				}},
			`{"filter":{"property":"phone_number","phone_number":{"ends_with":"+19876543210"}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "phone_number", PhoneNumber: &gotion.TextCondition{IsEmpty: true},
				}},
			`{"filter":{"property":"phone_number","phone_number":{"is_empty": true}}}`,
		},
		{
			&gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "phone_number", PhoneNumber: &gotion.TextCondition{IsNotEmpty: true},
				}},
			`{"filter":{"property":"phone_number","phone_number":{"is_not_empty":true}}}`,
		},
	}

	for _, pair := range se {
		actual, err := json.Marshal(pair.Source)
		utils.AssertNil(t, err)
		utils.AssertEqualsString(
			t,
			utils.Minimise(pair.Expected),
			utils.Minimise(string(actual)),
		)
	}
}

func TestMarshalNumberQuery(t *testing.T) {
	se := []querydbTuple{
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "number",
				Number:   &gotion.NumberCondition{Equals: 10},
			}},
			Expected: `{"filter":{"property":"number","number":{"equals":10}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "number",
				Number:   &gotion.NumberCondition{DoesntEqual: 10},
			}},
			Expected: `{"filter":{"property":"number","number":{"does_not_equal":10}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "number",
				Number:   &gotion.NumberCondition{GreaterThan: 10},
			}},
			Expected: `{"filter":{"property":"number","number":{"greater_than":10}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "number",
				Number:   &gotion.NumberCondition{LessThan: 10},
			}},
			Expected: `{"filter":{"property":"number","number":{"less_than":10}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "number",
				Number:   &gotion.NumberCondition{GreaterThanOrEqualTo: 10},
			}},
			Expected: `{"filter":{"property":"number","number":{"greater_than_or_equal_to":10}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "number",
				Number:   &gotion.NumberCondition{LessThanOrEqualTo: 10},
			}},
			Expected: `{"filter":{"property":"number","number":{"less_than_or_equal_to":10}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "number",
				Number:   &gotion.NumberCondition{IsEmpty: true},
			}},
			Expected: `{"filter":{"property":"number","number":{"is_empty":true}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "number",
				Number:   &gotion.NumberCondition{IsNotEmpty: true},
			}},
			Expected: `{"filter":{"property":"number","number":{"is_not_empty":true}}}`,
		},
	}

	for _, pair := range se {
		actual, err := json.Marshal(pair.Source)
		utils.AssertNil(t, err)
		utils.AssertEqualsString(
			t,
			utils.Minimise(pair.Expected),
			utils.Minimise(string(actual)),
		)
	}
}

func TestMarshalCheckboxQuery(t *testing.T) {
	se := []querydbTuple{
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "checkbox",
				Checkbox: &gotion.CheckboxCondition{
					Equals: utils.BoolPtr(true),
				},
			},
			},
			Expected: `{"filter":{"property":"checkbox","checkbox":{"equals":true}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "checkbox",
				Checkbox: &gotion.CheckboxCondition{
					Equals: utils.BoolPtr(false),
				},
			},
			},
			Expected: `{"filter":{"property":"checkbox","checkbox":{"equals":false}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "checkbox",
				Checkbox: &gotion.CheckboxCondition{
					DoesntEqual: utils.BoolPtr(true),
				},
			},
			},
			Expected: `{"filter":{"property":"checkbox","checkbox":{"does_not_equal":true}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "checkbox",
				Checkbox: &gotion.CheckboxCondition{
					DoesntEqual: utils.BoolPtr(false),
				},
			},
			},
			Expected: `{"filter":{"property":"checkbox","checkbox":{"does_not_equal":false}}}`,
		},
	}

	for _, pair := range se {
		actual, err := json.Marshal(pair.Source)
		utils.AssertNil(t, err)
		utils.AssertEqualsString(
			t,
			utils.Minimise(pair.Expected),
			utils.Minimise(string(actual)),
		)
	}
}

func TestMarshalSelectQuery(t *testing.T) {
	se := []querydbTuple{
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "select",
				Select: &gotion.SelectCondition{
					Equals: "tag",
				},
			}},
			Expected: `{"filter":{"property":"select","select":{"equals":"tag"}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "select",
				Select: &gotion.SelectCondition{
					DoesntEqual: "tag",
				},
			}},
			Expected: `{"filter":{"property":"select","select":{"does_not_equal":"tag"}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "select",
				Select: &gotion.SelectCondition{
					IsEmpty: true,
				},
			}},
			Expected: `{"filter":{"property":"select","select":{"is_empty":true}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "select",
				Select: &gotion.SelectCondition{
					IsNotEmpty: true,
				},
			}},
			Expected: `{"filter":{"property":"select","select":{"is_not_empty":true}}}`,
		},
	}

	for _, pair := range se {
		actual, err := json.Marshal(pair.Source)
		utils.AssertNil(t, err)
		utils.AssertEqualsString(
			t,
			utils.Minimise(pair.Expected),
			utils.Minimise(string(actual)),
		)
	}
}

func TestMarshalMultiSelectQuery(t *testing.T) {
	se := []querydbTuple{
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "multi_select",
				MultiSelect: &gotion.MultiSelectCondition{
					Contains: "tag",
				},
			}},
			Expected: `{"filter":{"property":"multi_select","multi_select":{"contains":"tag"}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "multi_select",
				MultiSelect: &gotion.MultiSelectCondition{
					DoesntContain: "tag",
				},
			}},
			Expected: `{"filter":{"property":"multi_select","multi_select":{"does_not_contain":"tag"}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "multi_select",
				MultiSelect: &gotion.MultiSelectCondition{
					IsEmpty: true,
				},
			}},
			Expected: `{"filter":{"property":"multi_select","multi_select":{"is_empty":true}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "multi_select",
				MultiSelect: &gotion.MultiSelectCondition{
					IsNotEmpty: true,
				},
			}},
			Expected: `{"filter":{"property":"multi_select","multi_select":{"is_not_empty":true}}}`,
		},
	}

	for _, pair := range se {
		actual, err := json.Marshal(pair.Source)
		utils.AssertNil(t, err)
		utils.AssertEqualsString(
			t,
			utils.Minimise(pair.Expected),
			utils.Minimise(string(actual)),
		)
	}
}

func TestMarshalFileQuery(t *testing.T) {
	se := []querydbTuple{
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "file",
				File: &gotion.FileCondition{
					IsEmpty: true,
				},
			}},
			Expected: `{"filter":{"property":"file","files":{"is_empty":true}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "file",
				File: &gotion.FileCondition{
					IsNotEmpty: true,
				},
			}},
			Expected: `{"filter":{"property":"file","files":{"is_not_empty":true}}}`,
		},
	}

	for _, pair := range se {
		actual, err := json.Marshal(pair.Source)
		utils.AssertNil(t, err)
		utils.AssertEqualsString(
			t,
			utils.Minimise(pair.Expected),
			utils.Minimise(string(actual)),
		)
	}
}

func TestMarshalRelationQuery(t *testing.T) {
	se := []querydbTuple{
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "relation",
				Relation: &gotion.RelationCondition{
					Contains: "869ec06096fd48a7a7492b15c9502cc9",
				},
			}},
			Expected: `{"filter":{"property":"relation","relation":{"contains":"869ec06096fd48a7a7492b15c9502cc9"}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "relation",
				Relation: &gotion.RelationCondition{
					DoesntContains: "869ec06096fd48a7a7492b15c9502cc9",
				},
			}},
			Expected: `{"filter":{"property":"relation","relation":{"does_not_contain":"869ec06096fd48a7a7492b15c9502cc9"}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "relation",
				Relation: &gotion.RelationCondition{
					IsEmpty: true,
				},
			}},
			Expected: `{"filter":{"property":"relation","relation":{"is_empty":true}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "relation",
				Relation: &gotion.RelationCondition{
					IsNotEmpty: true,
				},
			}},
			Expected: `{"filter":{"property":"relation","relation":{"is_not_empty":true}}}`,
		},
	}

	for _, pair := range se {
		actual, err := json.Marshal(pair.Source)
		utils.AssertNil(t, err)
		utils.AssertEqualsString(
			t,
			utils.Minimise(pair.Expected),
			utils.Minimise(string(actual)),
		)
	}
}

func TestMarshalPeopleQuery(t *testing.T) {
	se := []querydbTuple{
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "people",
				People: &gotion.PeopleCondition{
					Contains: "id_1",
				},
			}},
			Expected: `{"filter":{"property":"people","people":{"contains":"id_1"}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "people",
				People: &gotion.PeopleCondition{
					DoesntContains: "id_2",
				},
			}},
			Expected: `{"filter":{"property":"people","people":{"does_not_contain":"id_2"}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "people",
				People: &gotion.PeopleCondition{
					IsEmpty: true,
				},
			}},
			Expected: `{"filter":{"property":"people","people":{"is_empty":true}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "people",
				People: &gotion.PeopleCondition{
					IsNotEmpty: true,
				},
			}},
			Expected: `{"filter":{"property":"people","people":{"is_not_empty":true}}}`,
		},
	}

	for _, pair := range se {
		actual, err := json.Marshal(pair.Source)
		utils.AssertNil(t, err)
		utils.AssertEqualsString(
			t,
			utils.Minimise(pair.Expected),
			utils.Minimise(string(actual)),
		)
	}
}

func TestMarshalCreatedByQuery(t *testing.T) {
	se := []querydbTuple{
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "created_by",
				CreatedBy: &gotion.PeopleCondition{
					Contains: "id_1",
				},
			}},
			Expected: `{"filter":{"property":"created_by","created_by":{"contains":"id_1"}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "created_by",
				CreatedBy: &gotion.PeopleCondition{
					DoesntContains: "id_2",
				},
			}},
			Expected: `{"filter":{"property":"created_by","created_by":{"does_not_contain":"id_2"}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "created_by",
				CreatedBy: &gotion.PeopleCondition{
					IsEmpty: true,
				},
			}},
			Expected: `{"filter":{"property":"created_by","created_by":{"is_empty":true}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "created_by",
				CreatedBy: &gotion.PeopleCondition{
					IsNotEmpty: true,
				},
			}},
			Expected: `{"filter":{"property":"created_by","created_by":{"is_not_empty":true}}}`,
		},
	}

	for _, pair := range se {
		actual, err := json.Marshal(pair.Source)
		utils.AssertNil(t, err)
		utils.AssertEqualsString(
			t,
			utils.Minimise(pair.Expected),
			utils.Minimise(string(actual)),
		)
	}
}

func TestMarshalLastEditedByQuery(t *testing.T) {
	se := []querydbTuple{
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "last_edited_by",
				LastEditedBy: &gotion.PeopleCondition{
					Contains: "id_1",
				},
			}},
			Expected: `{"filter":{"property":"last_edited_by","last_edited_by":{"contains":"id_1"}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "last_edited_by",
				LastEditedBy: &gotion.PeopleCondition{
					DoesntContains: "id_2",
				},
			}},
			Expected: `{"filter":{"property":"last_edited_by","last_edited_by":{"does_not_contain":"id_2"}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "last_edited_by",
				LastEditedBy: &gotion.PeopleCondition{
					IsEmpty: true,
				},
			}},
			Expected: `{"filter":{"property":"last_edited_by","last_edited_by":{"is_empty":true}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "last_edited_by",
				LastEditedBy: &gotion.PeopleCondition{
					IsNotEmpty: true,
				},
			}},
			Expected: `{"filter":{"property":"last_edited_by","last_edited_by":{"is_not_empty":true}}}`,
		},
	}

	for _, pair := range se {
		actual, err := json.Marshal(pair.Source)
		utils.AssertNil(t, err)
		utils.AssertEqualsString(
			t,
			utils.Minimise(pair.Expected),
			utils.Minimise(string(actual)),
		)
	}
}

func TestMarshalFormulaQuery(t *testing.T) {
	dt, _ := time.Parse(time.RFC3339, "2033-05-18T03:33:00.100+00:00")
	se := []querydbTuple{
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "formula",
				Formula: &gotion.FormulaCondition{
					String: &gotion.TextCondition{
						Equals: "10",
					},
				},
			}},
			Expected: `{"filter":{"property":"formula","formula":{"string":{"equals":"10"}}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "formula",
				Formula: &gotion.FormulaCondition{
					Checkbox: &gotion.CheckboxCondition{
						Equals: utils.BoolPtr(true),
					},
				},
			}},
			Expected: `{"filter":{"property":"formula","formula":{"checkbox":{"equals":true}}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "formula",
				Formula: &gotion.FormulaCondition{
					Number: &gotion.NumberCondition{
						Equals: 10,
					},
				},
			}},
			Expected: `{"filter":{"property":"formula","formula":{"number":{"equals":10}}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "formula",
				Formula: &gotion.FormulaCondition{
					Date: &gotion.DateCondition{
						Equals: &gotion.DateTimeWrap{
							Datetime: dt,
						},
					},
				},
			}},
			Expected: `{"filter":{"property":"formula","formula":{"date":{"equals":"2033-05-18T03:33:00.1Z"}}}}`,
		},
	}

	for _, pair := range se {
		actual, err := json.Marshal(pair.Source)
		utils.AssertNil(t, err)
		utils.AssertEqualsString(
			t,
			utils.Minimise(pair.Expected),
			utils.Minimise(string(actual)),
		)
	}
}

func TestMarshalTRollupQuery(t *testing.T) {
	se := []querydbTuple{
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "rollup",
				Rollup: &gotion.RollupCondition{
					Any: &gotion.RollupProperty{
						RichText: &gotion.TextCondition{
							Equals: "some_string",
						},
					},
				},
			}},
			Expected: `{"filter":{"property":"rollup","rollup":{"any":{"rich_text":{"equals":"some_string"}}}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "rollup",
				Rollup: &gotion.RollupCondition{
					Every: &gotion.RollupProperty{
						RichText: &gotion.TextCondition{
							Equals: "some_string",
						},
					},
				},
			}},
			Expected: `{"filter":{"property":"rollup","rollup":{"every":{"rich_text":{"equals":"some_string"}}}}}`,
		},
		{
			Source: &gotion.QuertyDBRq{Filter: &gotion.Filter{
				Property: "rollup",
				Rollup: &gotion.RollupCondition{
					None: &gotion.RollupProperty{
						RichText: &gotion.TextCondition{
							Equals: "some_string",
						},
					},
				},
			}},
			Expected: `{"filter":{"property":"rollup","rollup":{"none":{"rich_text":{"equals":"some_string"}}}}}`,
		},
	}

	for _, pair := range se {
		actual, err := json.Marshal(pair.Source)
		utils.AssertNil(t, err)
		utils.AssertEqualsString(
			t,
			utils.Minimise(pair.Expected),
			utils.Minimise(string(actual)),
		)
	}
}

func TestMarshallDateCondition(t *testing.T) {
	dt, _ := time.Parse(time.RFC3339, "2033-05-18T03:33:00.100+00:00")
	se := []querydbTuple{
		querydbTuple{
			Source: &gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "dateprop",
					Date: &gotion.DateCondition{
						Equals: &gotion.DateTimeWrap{
							Datetime: dt,
						},
					},
				},
			},
			Expected: `{"filter":{"property":"dateprop", "date":{"equals":"2033-05-18T03:33:00.1Z"}}}`,
		},
		querydbTuple{
			Source: &gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "dateprop",
					Date: &gotion.DateCondition{
						Before: &gotion.DateTimeWrap{
							Datetime: dt,
						},
					},
				},
			},
			Expected: `{"filter":{"property":"dateprop", "date":{"before":"2033-05-18T03:33:00.1Z"}}}`,
		},
		querydbTuple{
			Source: &gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "dateprop",
					Date: &gotion.DateCondition{
						OnOrBefore: &gotion.DateTimeWrap{
							Datetime: dt,
						},
					},
				},
			},
			Expected: `{"filter":{"property":"dateprop", "date":{"on_or_before":"2033-05-18T03:33:00.1Z"}}}`,
		},
		querydbTuple{
			Source: &gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "dateprop",
					Date: &gotion.DateCondition{
						After: &gotion.DateTimeWrap{
							Datetime: dt,
						},
					},
				},
			},
			Expected: `{"filter":{"property":"dateprop", "date":{"after":"2033-05-18T03:33:00.1Z"}}}`,
		},
		querydbTuple{
			Source: &gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "dateprop",
					Date: &gotion.DateCondition{
						OnOrAfter: &gotion.DateTimeWrap{
							Datetime: dt,
						},
					},
				},
			},
			Expected: `{"filter":{"property":"dateprop", "date":{"on_or_after":"2033-05-18T03:33:00.1Z"}}}`,
		},
		querydbTuple{
			Source: &gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "dateprop",
					Date: &gotion.DateCondition{
						PastWeek: &gotion.DateTimeEmptyWrap{},
					},
				},
			},
			Expected: `{"filter":{"property":"dateprop", "date":{"past_week":{}}}}`,
		},
		querydbTuple{
			Source: &gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "dateprop",
					Date: &gotion.DateCondition{
						PastMonth: &gotion.DateTimeEmptyWrap{},
					},
				},
			},
			Expected: `{"filter":{"property":"dateprop", "date":{"past_month":{}}}}`,
		},
		querydbTuple{
			Source: &gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "dateprop",
					Date: &gotion.DateCondition{
						PastYear: &gotion.DateTimeEmptyWrap{},
					},
				},
			},
			Expected: `{"filter":{"property":"dateprop", "date":{"past_year":{}}}}`,
		},
		querydbTuple{
			Source: &gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "dateprop",
					Date: &gotion.DateCondition{
						NextWeek: &gotion.DateTimeEmptyWrap{},
					},
				},
			},
			Expected: `{"filter":{"property":"dateprop", "date":{"next_week":{}}}}`,
		},
		querydbTuple{
			Source: &gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "dateprop",
					Date: &gotion.DateCondition{
						NextMonth: &gotion.DateTimeEmptyWrap{},
					},
				},
			},
			Expected: `{"filter":{"property":"dateprop", "date":{"next_month":{}}}}`,
		},
		querydbTuple{
			Source: &gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "dateprop",
					Date: &gotion.DateCondition{
						NextYear: &gotion.DateTimeEmptyWrap{},
					},
				},
			},
			Expected: `{"filter":{"property":"dateprop", "date":{"next_year":{}}}}`,
		},
		querydbTuple{
			Source: &gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "dateprop",
					Date: &gotion.DateCondition{
						IsEmpty: true,
					},
				},
			},
			Expected: `{"filter":{"property":"dateprop", "date":{"is_empty":true}}}`,
		},
		querydbTuple{
			Source: &gotion.QuertyDBRq{
				Filter: &gotion.Filter{
					Property: "dateprop",
					Date: &gotion.DateCondition{
						IsNotEmpty: true,
					},
				},
			},
			Expected: `{"filter":{"property":"dateprop", "date":{"is_not_empty":true}}}`,
		},
	}
	for _, pair := range se {
		actual, err := json.Marshal(pair.Source)
		utils.AssertNil(t, err)
		utils.AssertEqualsString(
			t,
			utils.Minimise(pair.Expected),
			utils.Minimise(string(actual)),
		)
	}

}

func TestQueryWithSort(t *testing.T) {
	se := []querydbTuple{
		querydbTuple{
			Source: &gotion.QuertyDBRq{
				Sorts: []gotion.Sort{
					gotion.Sort{
						Property:  "title",
						Direction: gotion.Descending,
					},
				},
			},
			Expected: `{"sorts":[{"property":"title","direction":"descending"}]}`,
		},
		querydbTuple{
			Source: &gotion.QuertyDBRq{
				Sorts: []gotion.Sort{
					gotion.Sort{
						Property: "title",
					},
				},
			},
			Expected: `{"sorts":[{"property":"title","direction":"ascending"}]}`,
		},
		querydbTuple{
			Source: &gotion.QuertyDBRq{
				Sorts: []gotion.Sort{
					gotion.Sort{
						Property:  "title",
						Direction: gotion.Ascending,
					},
				},
			},
			Expected: `{"sorts":[{"property":"title","direction":"ascending"}]}`,
		},
		querydbTuple{
			Source: &gotion.QuertyDBRq{
				Sorts: []gotion.Sort{
					gotion.Sort{
						Timestamp: gotion.CreatedTime,
					},
				},
			},
			Expected: `{"sorts":[{"timestamp":"created_time","direction":"ascending"}]}`,
		},
		querydbTuple{
			Source: &gotion.QuertyDBRq{
				Sorts: []gotion.Sort{
					gotion.Sort{
						Timestamp: gotion.CreatedTime,
						Direction: gotion.Ascending,
					},
				},
			},
			Expected: `{"sorts":[{"timestamp":"created_time","direction":"ascending"}]}`,
		},
		querydbTuple{
			Source: &gotion.QuertyDBRq{
				Sorts: []gotion.Sort{
					gotion.Sort{
						Timestamp: gotion.CreatedTime,
						Direction: gotion.Descending,
					},
				},
			},
			Expected: `{"sorts":[{"timestamp":"created_time","direction":"descending"}]}`,
		},
		querydbTuple{
			Source: &gotion.QuertyDBRq{
				Sorts: []gotion.Sort{
					gotion.Sort{
						Timestamp: gotion.CreatedTime,
						Direction: gotion.Descending,
					},
					gotion.Sort{
						Property:  "title",
						Direction: gotion.Ascending,
					},
				},
			},
			Expected: `{"sorts":[{"timestamp":"created_time","direction":"descending"},{"property":"title","direction":"ascending"}]}`,
		},
	}
	for _, pair := range se {
		actual, err := json.Marshal(pair.Source)
		utils.AssertNil(t, err)
		utils.AssertEqualsString(
			t,
			utils.Minimise(pair.Expected),
			utils.Minimise(string(actual)),
		)
	}
}

func TestQueryFillFilled(t *testing.T) {
	query := gotion.QuertyDBRq{
		Filter: &gotion.Filter{
			Property: "title",
			RichText: &gotion.TextCondition{

				DoesntEqual: "Row1",
			},
		},
		Sorts: []gotion.Sort{
			gotion.Sort{
				Timestamp: gotion.CreatedTime,
				Direction: gotion.Descending,
			},
		},
		StartCursor: "3-295-0235",
		PageSize:    50,
	}

	marshaled, _ := json.Marshal(query)

	expected := `
	{
		"filter":{
			"property": "title",
			"rich_text": {
				"does_not_equal": "Row1"
			}
		},
		"sorts":[
			{
				"timestamp": "created_time",
				"direction": "descending"
			}
		],
		"start_cursor":"3-295-0235",
		"page_size":50
	}
	`

	utils.AssertEqualsString(t, utils.Minimise(expected), utils.Minimise(string(marshaled)))
}

func TestMarshallCreateDB(t *testing.T) {
	se := []createdbTuple{
		{
			Source: &gotion.CreateDBRq{
				PageId: "b0b48eac42514c2da3ab126a9986cf72",
				Title: []gotion.RichText{
					gotion.RichText{
						Text: &gotion.Text{
							Content: "Hello world",
						},
					},
				},
				Properties: gotion.DBProperties{
					"Name": gotion.DBProperty{
						Title: &gotion.DBDefaultProperty{},
					},
				},
			},
			Expected: `{"parent":{"type":"page_id","page_id":"b0b48eac42514c2da3ab126a9986cf72"},"title":[{"text":{"content":"Hello world"}}],"properties":{"Name":{"title":{}}}}`,
		},
		{
			Source: &gotion.CreateDBRq{
				PageId: "b0b48eac42514c2da3ab126a9986cf72",
				Title: []gotion.RichText{
					gotion.RichText{
						Text: &gotion.Text{
							Content: "Hello",
						},
						Annotations: &gotion.Annotations{
							Italic:        true,
							Bold:          true,
							Underline:     true,
							Strikethrough: true,
							Code:          true,
							Color:         "red",
						},
					},
					gotion.RichText{
						Mention: &gotion.Mention{
							User: &gotion.User{
								ID: "c7f2ae70-6b98-438f-8564-c59a71d7b3a4",
							},
						},
					},
					gotion.RichText{
						Equation: &gotion.Equation{
							Expression: "w + orld",
						},
					},
				},
				Properties: gotion.DBProperties{
					"Name": gotion.DBProperty{
						Title: &gotion.DBDefaultProperty{},
					},
					"Rich text": gotion.DBProperty{
						RichText: &gotion.DBDefaultProperty{},
					},
					"Bare number": gotion.DBProperty{
						Number: &gotion.DBNumberProperty{
							Format: gotion.Number,
						},
					},
					"Formatted number": gotion.DBProperty{
						Number: &gotion.DBNumberProperty{
							Format: gotion.Number,
						},
					},
					"Select": gotion.DBProperty{
						Select: &gotion.DBSelectProperties{
							Options: []gotion.DBSelectProperty{

								gotion.DBSelectProperty{
									Name:  "tag1",
									Color: gotion.Blue,
								},
								gotion.DBSelectProperty{
									Name:  "tag2",
									Color: gotion.Orange,
								},
							},
						},
					},
					"Multiselect": gotion.DBProperty{
						MultiSelect: &gotion.DBSelectProperties{
							Options: []gotion.DBSelectProperty{
								gotion.DBSelectProperty{
									Name:  "tag1_multi",
									Color: gotion.Blue,
								},
								gotion.DBSelectProperty{
									Name:  "tag2_multi",
									Color: gotion.Orange,
								},
							},
						},
					},
					"Date": gotion.DBProperty{
						Date: &gotion.DBDefaultProperty{},
					},
					"People": gotion.DBProperty{
						Date: &gotion.DBDefaultProperty{},
					},
					"Files": gotion.DBProperty{
						Files: &gotion.DBDefaultProperty{},
					},
					"Checkbox": gotion.DBProperty{
						Checkbox: &gotion.DBDefaultProperty{},
					},
					"URL": gotion.DBProperty{
						URL: &gotion.DBDefaultProperty{},
					},
					"Email": gotion.DBProperty{
						Email: &gotion.DBDefaultProperty{},
					},
					"Phone": gotion.DBProperty{
						PhoneNumber: &gotion.DBDefaultProperty{},
					},
					"Formula": gotion.DBProperty{
						Formula: &gotion.DBFormulatProperty{
							Expression: "1",
						},
					},
					"Created at": gotion.DBProperty{
						CreatedTime: &gotion.DBDefaultProperty{},
					},
					"Created by": gotion.DBProperty{
						CreatedBy: &gotion.DBDefaultProperty{},
					},
					"Updated at": gotion.DBProperty{
						LastEditedTime: &gotion.DBDefaultProperty{},
					},
					"Updated by": gotion.DBProperty{
						LastEditedBy: &gotion.DBDefaultProperty{},
					},
				},
			},
			Expected: `{"parent":{"type":"page_id","page_id":"b0b48eac42514c2da3ab126a9986cf72"},"title":[{"annotations":{"bold":true,"italic":true,"strikethrough":true,"underline":true,"code":true,"color":"red"},"text":{"content":"Hello"}},{"mention":{"user":{"id":"c7f2ae70-6b98-438f-8564-c59a71d7b3a4"}}},{"equation":{"expression":"w+orld"}}],"properties":{"Barenumber":{"number":{"format":"number"}},"Checkbox":{"checkbox":{}},"Createdat":{"created_time":{}},"Createdby":{"created_by":{}},"Date":{"date":{}},"Email":{"email":{}},"Files":{"files":{}},"Formattednumber":{"number":{"format":"number"}},"Formula":{"formula":{"expression":"1"}},"Multiselect":{"multi_select":{"options":[{"name":"tag1_multi","color":"blue"},{"name":"tag2_multi","color":"orange"}]}},"Name":{"title":{}},"People":{"date":{}},"Phone":{"phone_number":{}},"Richtext":{"rich_text":{}},"Select":{"select":{"options":[{"name":"tag1","color":"blue"},{"name":"tag2","color":"orange"}]}},"URL":{"url":{}},"Updatedat":{"last_edited_time":{}},"Updatedby":{"last_edited_by":{}}}}`,
		},
	}

	for _, pair := range se {
		actual, err := json.Marshal(pair.Source)
		utils.AssertNil(t, err)
		utils.AssertEqualsString(
			t,
			utils.Minimise(pair.Expected),
			utils.Minimise(string(actual)),
		)
	}
}

func TestMarshallUpdateDb(t *testing.T) {
	se := []updatedbTuple{
		{
			Source:   &gotion.UpdateDBRq{},
			Expected: "{}",
		},
		{
			Source: &gotion.UpdateDBRq{
				Icon: &gotion.IconDescriptor{
					Type:  gotion.EmojiIconType,
					Emoji: "🎎",
				},
			},
			Expected: `{"icon":{"type":"emoji","emoji":"🎎"}}`,
		},
		{
			Source: &gotion.UpdateDBRq{
				Cover: &gotion.FileDescriptor{
					Type: gotion.ExternalFileDescriptorType,
					ExternalFile: &gotion.ExternalFile{
						URL: "url",
					},
				},
			},
			Expected: `{"cover":{"type":"external","external":{"url":"url"}}}`,
		},
		{
			Source: &gotion.UpdateDBRq{
				Title: []gotion.RichText{
					{
						Type: gotion.TextRichTextType,
						Text: &gotion.Text{
							Content: "hello",
						},
					},
				},
			},
			Expected: `{"title":[{"type":"text","text":{"content":"hello"}}]}`,
		},
		{
			Source: &gotion.UpdateDBRq{
				Properties: gotion.DBProperties{
					"Name": gotion.DBProperty{
						Title: &gotion.DBDefaultProperty{},
					},
					"Rich text": gotion.DBProperty{
						RichText: &gotion.DBDefaultProperty{},
					},
					"Bare number": gotion.DBProperty{
						Number: &gotion.DBNumberProperty{
							Format: gotion.Number,
						},
					},
					"Formatted number": gotion.DBProperty{
						Number: &gotion.DBNumberProperty{
							Format: gotion.Number,
						},
					},
					"Select": gotion.DBProperty{
						Select: &gotion.DBSelectProperties{
							Options: []gotion.DBSelectProperty{

								gotion.DBSelectProperty{
									Name:  "tag1",
									Color: gotion.Blue,
								},
								gotion.DBSelectProperty{
									Name:  "tag2",
									Color: gotion.Orange,
								},
							},
						},
					},
					"Multiselect": gotion.DBProperty{
						MultiSelect: &gotion.DBSelectProperties{
							Options: []gotion.DBSelectProperty{
								gotion.DBSelectProperty{
									Name:  "tag1_multi",
									Color: gotion.Blue,
								},
								gotion.DBSelectProperty{
									Name:  "tag2_multi",
									Color: gotion.Orange,
								},
							},
						},
					},
					"Date": gotion.DBProperty{
						Date: &gotion.DBDefaultProperty{},
					},
					"People": gotion.DBProperty{
						Date: &gotion.DBDefaultProperty{},
					},
					"Files": gotion.DBProperty{
						Files: &gotion.DBDefaultProperty{},
					},
					"Checkbox": gotion.DBProperty{
						Checkbox: &gotion.DBDefaultProperty{},
					},
					"URL": gotion.DBProperty{
						URL: &gotion.DBDefaultProperty{},
					},
					"Email": gotion.DBProperty{
						Email: &gotion.DBDefaultProperty{},
					},
					"Phone": gotion.DBProperty{
						PhoneNumber: &gotion.DBDefaultProperty{},
					},
					"Formula": gotion.DBProperty{
						Formula: &gotion.DBFormulatProperty{
							Expression: "1",
						},
					},
					"Created at": gotion.DBProperty{
						CreatedTime: &gotion.DBDefaultProperty{},
					},
					"Created by": gotion.DBProperty{
						CreatedBy: &gotion.DBDefaultProperty{},
					},
					"Updated at": gotion.DBProperty{
						LastEditedTime: &gotion.DBDefaultProperty{},
					},
					"Updated by": gotion.DBProperty{
						LastEditedBy: &gotion.DBDefaultProperty{},
					},
				},
			},
			Expected: `{"properties":{"Barenumber":{"number":{"format":"number"}},"Checkbox":{"checkbox":{}},"Createdat":{"created_time":{}},"Createdby":{"created_by":{}},"Date":{"date":{}},"Email":{"email":{}},"Files":{"files":{}},"Formattednumber":{"number":{"format":"number"}},"Formula":{"formula":{"expression":"1"}},"Multiselect":{"multi_select":{"options":[{"name":"tag1_multi","color":"blue"},{"name":"tag2_multi","color":"orange"}]}},"Name":{"title":{}},"People":{"date":{}},"Phone":{"phone_number":{}},"Richtext":{"rich_text":{}},"Select":{"select":{"options":[{"name":"tag1","color":"blue"},{"name":"tag2","color":"orange"}]}},"URL":{"url":{}},"Updatedat":{"last_edited_time":{}},"Updatedby":{"last_edited_by":{}}}}`,
		},
	}

	for _, pair := range se {
		actual, err := json.Marshal(pair.Source)
		utils.AssertNil(t, err)
		utils.AssertEqualsString(
			t,
			utils.Minimise(pair.Expected),
			utils.Minimise(string(actual)),
		)
	}
}

func TestUnmarshalDatabaseMeta(t *testing.T) {

	se := []retrievedbTuple{
		retrievedbTuple{
			Source: `
					{
						"object": "database",
						"id": "d8959563-9ad5-4ed3-9176-08bb23e9f348",
						"cover": null,
						"icon": null,
						"created_time": "2022-05-24T22:40:00.000Z",
						"created_by": {
							"object": "user",
							"id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
						},
						"last_edited_by": {
							"object": "user",
							"id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
						},
						"last_edited_time": "2022-05-26T21:53:00.000Z",
						"title": [],
						"properties": {
							"Name": {
								"id": "title",
								"name": "Name",
								"type": "title",
								"title": {}
								}
						},
						"parent": {
							"type": "page_id",
							"page_id": "b0b48eac-4251-4c2d-a3ab-126a9986cf72"
						},
						"url": "https://www.notion.so/d89595639ad54ed3917608bb23e9f348",
						"archived": false
					}
		  `,
			Expected: gotion.Database{
				Object: "database",
				ID:     "d8959563-9ad5-4ed3-9176-08bb23e9f348",
				CreatedTime: gotion.DateTimeWrap{
					Datetime: time.Date(2022, 5, 24, 22, 40, 0, 0, time.UTC),
				},
				CreatedBy: gotion.User{
					Object: "user",
					ID:     "c7f2ae70-6b98-438f-8564-c59a71d7b3a4",
				},
				LastEditedBy: gotion.User{
					Object: "user",
					ID:     "c7f2ae70-6b98-438f-8564-c59a71d7b3a4",
				},
				LastEditedTime: gotion.DateTimeWrap{
					Datetime: time.Date(2022, 5, 26, 21, 53, 0, 0, time.UTC),
				},
				Title: []gotion.RichText{},
				Properties: gotion.DBProperties{
					"Name": gotion.DBProperty{
						ID:    "title",
						Name:  "Name",
						Type:  gotion.PropTypeTitle,
						Title: &gotion.DBDefaultProperty{},
					},
				},
				Parent: gotion.DBParent{
					Type:   gotion.PageDBParentType,
					PageID: "b0b48eac-4251-4c2d-a3ab-126a9986cf72",
				},
				URL:      "https://www.notion.so/d89595639ad54ed3917608bb23e9f348",
				Archived: false,
			},
		},
	}

	for _, pair := range se {
		res := gotion.Database{}
		json.Unmarshal([]byte(pair.Source), &res)
		utils.AssertEqualsStruct(t, pair.Expected, res)
	}
}
