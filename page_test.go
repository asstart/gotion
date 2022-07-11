package gotion_test

import (
	"encoding/json"
	// "fmt"
	"math"
	"testing"
	"time"

	"github.com/asstart/gotion"
	"github.com/asstart/gotion/utils"
	"github.com/stretchr/testify/assert"
)

type PageTestTuple struct {
	Source   string
	Expected gotion.Page
}

func TestRetrievePage(t *testing.T) {
	se := []PageTestTuple{
		{
			Source: `
			{
				"object": "page",
				"id": "1bcea86a-2f57-4171-a06b-c6052bea314b",
				"created_time": "2022-06-19T12:23:00.000Z",
				"last_edited_time": "2022-06-19T13:10:00.000Z",
				"created_by": {
				  "object": "user",
				  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
				},
				"last_edited_by": {
				  "object": "user",
				  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
				},
				"cover": null,
				"icon": null,
				"parent": {
				  "type": "page_id",
				  "page_id": "b0b48eac-4251-4c2d-a3ab-126a9986cf72"
				},
				"archived": false,
				"properties": {
				  "title": {
					"id": "title",
					"type": "title",
					"title": [
					  {
						"type": "text",
						"text": {
						  "content": "Empty page ",
						  "link": null
						},
						"annotations": {
						  "bold": false,
						  "italic": false,
						  "strikethrough": false,
						  "underline": false,
						  "code": false,
						  "color": "default"
						},
						"plain_text": "Empty page ",
						"href": null
					  }
					]
				  }
				},
				"url": "https://www.notion.so/Empty-page-1bcea86a2f574171a06bc6052bea314b"
			  }
			  `,
			Expected: gotion.Page{
				Object: "page",
				ID:     "1bcea86a-2f57-4171-a06b-c6052bea314b",
				CreatedTime: gotion.DateTimeWrap{
					Datetime: time.Date(2022, 6, 19, 12, 23, 0, 0, time.UTC),
				},
				LastEditedTime: gotion.DateTimeWrap{
					Datetime: time.Date(2022, 6, 19, 13, 10, 0, 0, time.UTC),
				},
				CreatedBy: gotion.User{
					Object: "user",
					ID:     "c7f2ae70-6b98-438f-8564-c59a71d7b3a4",
				},
				LastEditedBy: gotion.User{
					Object: "user",
					ID:     "c7f2ae70-6b98-438f-8564-c59a71d7b3a4",
				},
				Parent: gotion.PageParent{
					Type:   gotion.PageParentPage,
					PageID: "b0b48eac-4251-4c2d-a3ab-126a9986cf72",
				},
				Properties: gotion.PageProperties{
					"title": gotion.PageProperty{
						ID:   "title",
						Type: gotion.DBPropTypeTitle,
						Title: []gotion.RichText{
							gotion.RichText{
								Type: gotion.TextRichTextType,
								Text: &gotion.Text{
									Content: "Empty page ",
								},
								Annotations: &gotion.Annotations{},
								PlainText:   "Empty page ",
							},
						},
					},
				},
				URL: "https://www.notion.so/Empty-page-1bcea86a2f574171a06bc6052bea314b",
			},
		},
		{
			Source: `
			{
				"object": "page",
				"id": "c1d21d5d-fd19-486a-9888-556f193c1b4d",
				"created_time": "2022-06-19T15:15:00.000Z",
				"last_edited_time": "2022-06-19T15:15:00.000Z",
				"created_by": {
				  "object": "user",
				  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
				},
				"last_edited_by": {
				  "object": "user",
				  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
				},
				"cover": {
				  "type": "external",
				  "external": {
					"url": "https://www.notion.so/images/page-cover/met_william_morris_1877_willow.jpg"
				  }
				},
				"icon": {
				  "type": "emoji",
				  "emoji": "🚉"
				},
				"parent": {
				  "type": "page_id",
				  "page_id": "b0b48eac-4251-4c2d-a3ab-126a9986cf72"
				},
				"archived": false,
				"properties": {
				  "title": {
					"id": "title",
					"type": "title",
					"title": [
					  {
						"type": "text",
						"text": {
						  "content": "Filled page",
						  "link": null
						},
						"annotations": {
						  "bold": false,
						  "italic": false,
						  "strikethrough": false,
						  "underline": false,
						  "code": false,
						  "color": "default"
						},
						"plain_text": "Filled page",
						"href": null
					  }
					]
				  }
				},
				"url": "https://www.notion.so/Filled-page-c1d21d5dfd19486a9888556f193c1b4d"
			  }
			`,
			Expected: gotion.Page{
				Object: "page",
				ID:     "c1d21d5d-fd19-486a-9888-556f193c1b4d",
				CreatedTime: gotion.DateTimeWrap{
					Datetime: time.Date(2022, 6, 19, 15, 15, 0, 0, time.UTC),
				},
				LastEditedTime: gotion.DateTimeWrap{
					Datetime: time.Date(2022, 6, 19, 15, 15, 0, 0, time.UTC),
				},
				CreatedBy: gotion.User{
					Object: "user",
					ID:     "c7f2ae70-6b98-438f-8564-c59a71d7b3a4",
				},
				LastEditedBy: gotion.User{
					Object: "user",
					ID:     "c7f2ae70-6b98-438f-8564-c59a71d7b3a4",
				},
				Cover: &gotion.FileDescriptor{
					Type: gotion.ExternalFileDescriptorType,
					ExternalFile: &gotion.ExternalFile{
						URL: "https://www.notion.so/images/page-cover/met_william_morris_1877_willow.jpg",
					},
				},
				Icon: &gotion.IconDescriptor{
					Type:  gotion.EmojiIconType,
					Emoji: "🚉",
				},
				Parent: gotion.PageParent{
					Type:   gotion.PageParentPage,
					PageID: "b0b48eac-4251-4c2d-a3ab-126a9986cf72",
				},
				Properties: gotion.PageProperties{
					"title": gotion.PageProperty{
						ID:   "title",
						Type: gotion.DBPropTypeTitle,
						Title: []gotion.RichText{
							gotion.RichText{
								Type: gotion.TextRichTextType,
								Text: &gotion.Text{
									Content: "Filled page",
								},
								Annotations: &gotion.Annotations{},
								PlainText:   "Filled page",
							},
						},
					},
				},
				URL: "https://www.notion.so/Filled-page-c1d21d5dfd19486a9888556f193c1b4d",
			},
		},
		{
			Source: `
			{
				"object": "page",
				"id": "e0f73c2d-7a2e-4fc2-9f28-73380ed76512",
				"created_time": "2022-06-19T13:07:00.000Z",
				"last_edited_time": "2022-06-19T18:58:00.000Z",
				"created_by": {
				  "object": "user",
				  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
				},
				"last_edited_by": {
				  "object": "user",
				  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
				},
				"cover": null,
				"icon": null,
				"parent": {
				  "type": "database_id",
				  "database_id": "ba96e9ad-4600-47a7-a868-5f3328618910"
				},
				"archived": false,
				"properties": {
				  "Email Prop": {
					"id": "%3AbTo",
					"type": "email",
					"email": "test@mail.dev"
				  },
				  "User Prop": {
					"id": "%3BC%3Dh",
					"type": "people",
					"people": [
					  {
						"object": "user",
						"id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4",
						"name": "Имя",
						"avatar_url": null,
						"type": "person",
						"person": {
						  "email": "test@mail.dev"
						}
					  }
					]
				  },
				  "Phone Prop": {
					"id": "%3F%3EgN",
					"type": "phone_number",
					"phone_number": "+79998887766"
				  },
				  "Number Prop": {
					"id": "%40E%7Cn",
					"type": "number",
					"number": 30
				  },
				  "Date formula": {
					"id": "%40VDr",
					"type": "formula",
					"formula": {
					  "type": "date",
					  "date": {
						"start": "2022-06-19T19:00:00.000Z",
						"end": null,
						"time_zone": null
					  }
					}
				  },
				  "Created By Prop": {
					"id": "%40ZPE",
					"type": "created_by",
					"created_by": {
					  "object": "user",
					  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4",
					  "name": "Имя",
					  "avatar_url": null,
					  "type": "person",
					  "person": {
						"email": "test@mail.dev"
					  }
					}
				  },
				  "Date Range Prop": {
					"id": "Cl%5E%5C",
					"type": "date",
					"date": {
					  "start": "2022-06-15",
					  "end": "2022-06-15",
					  "time_zone": null
					}
				  },
				  "Number formula": {
					"id": "KRxq",
					"type": "formula",
					"formula": {
					  "type": "number",
					  "number": 100
					}
				  },
				  "Marked Checkbox": {
					"id": "KZ%3Fr",
					"type": "checkbox",
					"checkbox": true
				  },
				  "Last Edited Time Prop": {
					"id": "O%3DYK",
					"type": "last_edited_time",
					"last_edited_time": "2022-06-19T18:58:00.000Z"
				  },
				  "Single Datetime Prop": {
					"id": "Pb%3E%3A",
					"type": "date",
					"date": {
					  "start": "2022-06-19T00:00:00.000-11:00",
					  "end": null,
					  "time_zone": null
					}
				  },
				  "External file": {
					"id": "QK%3Ak",
					"type": "files",
					"files": [
					  {
						"name": "https://github.com/asstart/gotion/blob/main/LICENSE",
						"type": "external",
						"external": {
						  "url": "https://github.com/asstart/gotion/blob/main/LICENSE"
						}
					  }
					]
				  },
				  "Unmarked checkbox": {
					"id": "So%40y",
					"type": "checkbox",
					"checkbox": false
				  },
				  "Last Edited By Prop": {
					"id": "UjPK",
					"type": "last_edited_by",
					"last_edited_by": {
					  "object": "user",
					  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4",
					  "name": "Имя",
					  "avatar_url": null,
					  "type": "person",
					  "person": {
						"email": "test@mail.dev"
					  }
					}
				  },
				  "Text formula": {
					"id": "VDLR",
					"type": "formula",
					"formula": {
					  "type": "string",
					  "string": "Title with many diff styles"
					}
				  },
				  "Multiselect Prop": {
					"id": "Y%3B%5Du",
					"type": "multi_select",
					"multi_select": [
					  {
						"id": "6c0151da-c522-48e3-812b-428044bf9338",
						"name": "t11",
						"color": "red"
					  },
					  {
						"id": "01ba9b69-9953-4d2d-9cb7-8a32ef179388",
						"name": "t12",
						"color": "pink"
					  }
					]
				  },
				  "Bool formula": {
					"id": "c%3AI%5E",
					"type": "formula",
					"formula": {
					  "type": "boolean",
					  "boolean": true
					}
				  },
				  "Single Date Prop": {
					"id": "cKQx",
					"type": "date",
					"date": {
					  "start": "2022-06-16",
					  "end": null,
					  "time_zone": null
					}
				  },
				  "Link Prop": {
					"id": "fFC%5B",
					"type": "url",
					"url": "https://github.com/asstart/gotion"
				  },
				  "Text Prop": {
					"id": "jWtn",
					"type": "rich_text",
					"rich_text": [
					  {
						"type": "text",
						"text": {
						  "content": "Prop ",
						  "link": null
						},
						"annotations": {
						  "bold": true,
						  "italic": false,
						  "strikethrough": false,
						  "underline": false,
						  "code": false,
						  "color": "default"
						},
						"plain_text": "Prop ",
						"href": null
					  },
					  {
						"type": "text",
						"text": {
						  "content": "with ",
						  "link": null
						},
						"annotations": {
						  "bold": false,
						  "italic": true,
						  "strikethrough": false,
						  "underline": false,
						  "code": false,
						  "color": "default"
						},
						"plain_text": "with ",
						"href": null
					  },
					  {
						"type": "text",
						"text": {
						  "content": "many ",
						  "link": null
						},
						"annotations": {
						  "bold": false,
						  "italic": true,
						  "strikethrough": true,
						  "underline": false,
						  "code": false,
						  "color": "default"
						},
						"plain_text": "many ",
						"href": null
					  },
					  {
						"type": "text",
						"text": {
						  "content": "styles",
						  "link": null
						},
						"annotations": {
						  "bold": false,
						  "italic": true,
						  "strikethrough": true,
						  "underline": false,
						  "code": true,
						  "color": "default"
						},
						"plain_text": "styles",
						"href": null
					  }
					]
				  },
				  "Created Time Prop": {
					"id": "w%3Fp%7C",
					"type": "created_time",
					"created_time": "2022-06-19T13:07:00.000Z"
				  },
				  "Notion File Prop": {
					"id": "wXSX",
					"type": "files",
					"files": [
					  {
						"name": "exmpl.txt",
						"type": "file",
						"file": {
						  "url": "https://s3.us-west-2.amazonaws.com/secure.notion-static.com/84650ce1-7307-48a5-9139-01f503db7019/exmpl.txt?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Content-Sha256=UNSIGNED-PAYLOAD&X-Amz-Credential=AKIAT73L2G45EIPT3X45%2F20220619%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20220619T190041Z&X-Amz-Expires=3600&X-Amz-Signature=f101ca2bc0c7037c62b75e203fd5b27595d4aff000b586b1d8e1071649bfc69b&X-Amz-SignedHeaders=host&x-id=GetObject",
						  "expiry_time": "2022-06-19T20:00:41.932Z"
						}
					  }
					]
				  },
				  "Select Prop": {
					"id": "%7CRrh",
					"type": "select",
					"select": {
					  "id": "aeb7a7de-519f-4458-a77a-5058af1894e2",
					  "name": "t1",
					  "color": "blue"
					}
				  },
				  "Prop Text With Link": {
					"id": "Ew%3BV",
					"type": "rich_text",
					"rich_text": [
					  {
						"type": "text",
						"text": {
						  "content": "RFC ",
						  "link": null
						},
						"annotations": {
						  "bold": false,
						  "italic": false,
						  "strikethrough": false,
						  "underline": false,
						  "code": false,
						  "color": "default"
						},
						"plain_text": "RFC ",
						"href": null
					  },
					  {
						"type": "text",
						"text": {
						  "content": "https://www.rfc-editor.org/rfc/rfc3339.html",
						  "link": {
							"url": "https://www.rfc-editor.org/rfc/rfc3339.html"
						  }
						},
						"annotations": {
						  "bold": false,
						  "italic": false,
						  "strikethrough": false,
						  "underline": false,
						  "code": false,
						  "color": "default"
						},
						"plain_text": "https://www.rfc-editor.org/rfc/rfc3339.html",
						"href": "https://www.rfc-editor.org/rfc/rfc3339.html"
					  }
					]
				  },
				  "Prop Text With Date Mention": {
					"id": "Sv_%3C",
					"type": "rich_text",
					"rich_text": [
					  {
						"type": "text",
						"text": {
						  "content": "Mention ",
						  "link": null
						},
						"annotations": {
						  "bold": false,
						  "italic": false,
						  "strikethrough": false,
						  "underline": false,
						  "code": false,
						  "color": "default"
						},
						"plain_text": "Mention ",
						"href": null
					  },
					  {
						"type": "mention",
						"mention": {
						  "type": "date",
						  "date": {
							"start": "2022-06-20",
							"end": null,
							"time_zone": null
						  }
						},
						"annotations": {
						  "bold": false,
						  "italic": false,
						  "strikethrough": false,
						  "underline": false,
						  "code": false,
						  "color": "default"
						},
						"plain_text": "2022-06-20 → ",
						"href": null
					  },
					  {
						"type": "text",
						"text": {
						  "content": " ",
						  "link": null
						},
						"annotations": {
						  "bold": false,
						  "italic": false,
						  "strikethrough": false,
						  "underline": false,
						  "code": false,
						  "color": "default"
						},
						"plain_text": " ",
						"href": null
					  }
					]
				  },
				  "Prop Text With Mention": {
					"id": "TpMl",
					"type": "rich_text",
					"rich_text": [
					  {
						"type": "text",
						"text": {
						  "content": "Mention ",
						  "link": null
						},
						"annotations": {
						  "bold": false,
						  "italic": false,
						  "strikethrough": false,
						  "underline": false,
						  "code": false,
						  "color": "default"
						},
						"plain_text": "Mention ",
						"href": null
					  },
					  {
						"type": "mention",
						"mention": {
						  "type": "user",
						  "user": {
							"object": "user",
							"id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4",
							"name": "Имя",
							"avatar_url": null,
							"type": "person",
							"person": {
							  "email": "test@mail.dev"
							}
						  }
						},
						"annotations": {
						  "bold": false,
						  "italic": false,
						  "strikethrough": false,
						  "underline": false,
						  "code": false,
						  "color": "default"
						},
						"plain_text": "@Имя",
						"href": null
					  },
					  {
						"type": "text",
						"text": {
						  "content": " ",
						  "link": null
						},
						"annotations": {
						  "bold": false,
						  "italic": false,
						  "strikethrough": false,
						  "underline": false,
						  "code": false,
						  "color": "default"
						},
						"plain_text": " ",
						"href": null
					  }
					]
				  },
				  "Prop Text With Mention Database": {
					"id": "r%5CU%3D",
					"type": "rich_text",
					"rich_text": [
					  {
						"type": "text",
						"text": {
						  "content": "Mention ",
						  "link": null
						},
						"annotations": {
						  "bold": false,
						  "italic": false,
						  "strikethrough": false,
						  "underline": false,
						  "code": false,
						  "color": "default"
						},
						"plain_text": "Mention ",
						"href": null
					  },
					  {
						"type": "mention",
						"mention": {
						  "type": "database",
						  "database": {
							"id": "ba96e9ad-4600-47a7-a868-5f3328618910"
						  }
						},
						"annotations": {
						  "bold": false,
						  "italic": false,
						  "strikethrough": false,
						  "underline": false,
						  "code": false,
						  "color": "default"
						},
						"plain_text": "Rich Text prop db",
						"href": "https://www.notion.so/ba96e9ad460047a7a8685f3328618910"
					  },
					  {
						"type": "text",
						"text": {
						  "content": " ",
						  "link": null
						},
						"annotations": {
						  "bold": false,
						  "italic": false,
						  "strikethrough": false,
						  "underline": false,
						  "code": false,
						  "color": "default"
						},
						"plain_text": " ",
						"href": null
					  }
					]
				  },
				  "Prop Text With Mention Page": {
					"id": "yb%3Ew",
					"type": "rich_text",
					"rich_text": [
					  {
						"type": "text",
						"text": {
						  "content": "Mention ",
						  "link": null
						},
						"annotations": {
						  "bold": false,
						  "italic": false,
						  "strikethrough": false,
						  "underline": false,
						  "code": false,
						  "color": "default"
						},
						"plain_text": "Mention ",
						"href": null
					  },
					  {
						"type": "mention",
						"mention": {
						  "type": "page",
						  "page": {
							"id": "548b9b8e-89a3-4e01-8dc4-35f76b1cb653"
						  }
						},
						"annotations": {
						  "bold": false,
						  "italic": false,
						  "strikethrough": false,
						  "underline": false,
						  "code": false,
						  "color": "default"
						},
						"plain_text": "Row1",
						"href": "https://www.notion.so/548b9b8e89a34e018dc435f76b1cb653"
					  },
					  {
						"type": "text",
						"text": {
						  "content": " ",
						  "link": null
						},
						"annotations": {
						  "bold": false,
						  "italic": false,
						  "strikethrough": false,
						  "underline": false,
						  "code": false,
						  "color": "default"
						},
						"plain_text": " ",
						"href": null
					  }
					]
				  },
				  "Property Equation": {
					"id": "zgh%7C",
					"type": "rich_text",
					"rich_text": [
					  {
						"type": "equation",
						"equation": {
						  "expression": "E = mc^2"
						},
						"annotations": {
						  "bold": false,
						  "italic": false,
						  "strikethrough": false,
						  "underline": false,
						  "code": false,
						  "color": "default"
						},
						"plain_text": "E = mc^2",
						"href": null
					  },
					  {
						"type": "text",
						"text": {
						  "content": " ",
						  "link": null
						},
						"annotations": {
						  "bold": false,
						  "italic": false,
						  "strikethrough": false,
						  "underline": false,
						  "code": false,
						  "color": "default"
						},
						"plain_text": " ",
						"href": null
					  }
					]
				  },
				  "Related Rich Text prop db": {
					"id": "%5B%3Fnd",
					"type": "relation",
					"relation": [
					  {
						"id": "e0f73c2d-7a2e-4fc2-9f28-73380ed76512"
					  }
					]
				  },
				  "Rollup Property": {
					"id": "KOL%3F",
					"type": "rollup",
					"rollup": {
					  "type": "number",
					  "number": 30,
					  "function": "max"
					}
				  },
				  "Name": {
					"id": "title",
					"type": "title",
					"title": [
					  {
						"type": "text",
						"text": {
						  "content": "Title with many diff styles",
						  "link": null
						},
						"annotations": {
						  "bold": false,
						  "italic": false,
						  "strikethrough": false,
						  "underline": false,
						  "code": false,
						  "color": "default"
						},
						"plain_text": "Title with many diff styles",
						"href": null
					  }
					]
				  }
				},
				"url": "https://www.notion.so/Title-with-many-diff-styles-e0f73c2d7a2e4fc29f2873380ed76512"
			  }
			`,
			Expected: gotion.Page{
				Object: "page",
				ID:     "e0f73c2d-7a2e-4fc2-9f28-73380ed76512",
				CreatedTime: gotion.DateTimeWrap{
					Datetime: time.Date(2022, 6, 19, 13, 7, 0, 0, time.UTC),
				},
				LastEditedTime: gotion.DateTimeWrap{
					Datetime: time.Date(2022, 6, 19, 18, 58, 0, 0, time.UTC),
				},
				CreatedBy: gotion.User{
					Object: "user",
					ID:     "c7f2ae70-6b98-438f-8564-c59a71d7b3a4",
				},
				LastEditedBy: gotion.User{
					Object: "user",
					ID:     "c7f2ae70-6b98-438f-8564-c59a71d7b3a4",
				},
				Parent: gotion.PageParent{
					Type:       gotion.PageParentDB,
					DatabaseID: "ba96e9ad-4600-47a7-a868-5f3328618910",
				},
				Properties: gotion.PageProperties{
					"Email Prop": gotion.PageProperty{
						ID:    "%3AbTo",
						Type:  gotion.DBPropTypeEmail,
						Email: utils.StrPtr("test@mail.dev"),
					},
					"User Prop": gotion.PageProperty{
						ID:   "%3BC%3Dh",
						Type: gotion.DBPropTypePeople,
						People: []gotion.User{
							gotion.User{
								Object: "user",
								ID:     "c7f2ae70-6b98-438f-8564-c59a71d7b3a4",
								Name:   "Имя",
								Type:   gotion.PersonType,
								Person: &gotion.Person{
									Email: "test@mail.dev",
								},
							},
						},
					},
					"Phone Prop": gotion.PageProperty{
						ID:          "%3F%3EgN",
						Type:        gotion.DBPropTypePhoneNumber,
						PhoneNumber: utils.StrPtr("+79998887766"),
					},
					"Number Prop": gotion.PageProperty{
						ID:     "%40E%7Cn",
						Type:   gotion.DBPropTypeNumber,
						Number: utils.FloatPtr(30),
					},
					"Date formula": gotion.PageProperty{
						ID:   "%40VDr",
						Type: gotion.DBPropTypeFormula,
						Formula: &gotion.FormulaPageProperty{
							Type: gotion.DateFormuleType,
							DateFormula: &gotion.DatePageProperty{
								Start: gotion.DateTimeWrap{
									Datetime: time.Date(2022, 6, 19, 19, 0, 0, 0, time.UTC),
								},
							},
						},
					},
					"Created By Prop": gotion.PageProperty{
						ID:   "%40ZPE",
						Type: gotion.DBPropTypeCreatedBy,
						CreatedBy: &gotion.User{
							Object: "user",
							ID:     "c7f2ae70-6b98-438f-8564-c59a71d7b3a4",
							Name:   "Имя",
							Type:   gotion.PersonType,
							Person: &gotion.Person{
								Email: "test@mail.dev",
							},
						},
					},
					"Date Range Prop": {
						ID:   "Cl%5E%5C",
						Type: gotion.DBPropTypeDate,
						Date: &gotion.DatePageProperty{
							Start: gotion.DateTimeWrap{
								Datetime: time.Date(2022, 6, 15, 0, 0, 0, 0, time.UTC),
							},
							End: &gotion.DateTimeWrap{
								Datetime: time.Date(2022, 6, 15, 0, 0, 0, 0, time.UTC),
							},
						},
					},
					"Number formula": gotion.PageProperty{
						ID:   "KRxq",
						Type: gotion.DBPropTypeFormula,
						Formula: &gotion.FormulaPageProperty{
							Type:          gotion.NumberFormulaType,
							NumberFormula: utils.FloatPtr(100),
						},
					},
					"Marked Checkbox": gotion.PageProperty{
						ID:       "KZ%3Fr",
						Type:     gotion.DBPropTypeCheckbox,
						Checkbox: utils.BoolPtr(true),
					},
					"Last Edited Time Prop": gotion.PageProperty{
						ID:   "O%3DYK",
						Type: gotion.DBPropTypeLastEditedTime,
						LastEditedTime: &gotion.DateTimeWrap{
							Datetime: time.Date(2022, 6, 19, 18, 58, 0, 0, time.UTC),
						},
					},
					"Single Datetime Prop": gotion.PageProperty{
						ID:   "Pb%3E%3A",
						Type: gotion.DBPropTypeDate,
						Date: &gotion.DatePageProperty{
							Start: gotion.DateTimeWrap{
								Datetime: time.Date(2022, 6, 19, 0, 0, 0, 0, time.FixedZone("", -11*60*60)),
							},
						},
					},
					"External file": gotion.PageProperty{
						ID:   "QK%3Ak",
						Type: gotion.DBPropTypeFiles,
						Files: []gotion.FileDescriptor{
							gotion.FileDescriptor{
								Name: "https://github.com/asstart/gotion/blob/main/LICENSE",
								Type: gotion.ExternalFileDescriptorType,
								ExternalFile: &gotion.ExternalFile{
									URL: "https://github.com/asstart/gotion/blob/main/LICENSE",
								},
							},
						},
					},
					"Unmarked checkbox": gotion.PageProperty{
						ID:       "So%40y",
						Type:     gotion.DBPropTypeCheckbox,
						Checkbox: utils.BoolPtr(false),
					},
					"Last Edited By Prop": gotion.PageProperty{
						ID:   "UjPK",
						Type: gotion.DBPropTypeLastEditedBy,
						LastEditedBy: &gotion.User{
							Object: "user",
							ID:     "c7f2ae70-6b98-438f-8564-c59a71d7b3a4",
							Name:   "Имя",
							Type:   gotion.PersonType,
							Person: &gotion.Person{
								Email: "test@mail.dev",
							},
						},
					},
					"Text formula": gotion.PageProperty{
						ID:   "VDLR",
						Type: gotion.DBPropTypeFormula,
						Formula: &gotion.FormulaPageProperty{
							Type:          gotion.StringFormulaType,
							StringFormula: utils.StrPtr("Title with many diff styles"),
						},
					},
					"Multiselect Prop": {
						ID:   "Y%3B%5Du",
						Type: gotion.DBPropTypeMultiSelect,
						MultiSelect: []gotion.SelectPageProperty{
							gotion.SelectPageProperty{
								ID:    "6c0151da-c522-48e3-812b-428044bf9338",
								Name:  "t11",
								Color: gotion.Red,
							},
							gotion.SelectPageProperty{
								ID:    "01ba9b69-9953-4d2d-9cb7-8a32ef179388",
								Name:  "t12",
								Color: gotion.Pink,
							},
						},
					},
					"Bool formula": gotion.PageProperty{
						ID:   "c%3AI%5E",
						Type: gotion.DBPropTypeFormula,
						Formula: &gotion.FormulaPageProperty{
							Type:           gotion.BoolFormulaType,
							BooleanFormula: utils.BoolPtr(true),
						},
					},
					"Single Date Prop": gotion.PageProperty{
						ID:   "cKQx",
						Type: gotion.DBPropTypeDate,
						Date: &gotion.DatePageProperty{
							Start: gotion.DateTimeWrap{
								Datetime: time.Date(2022, 6, 16, 0, 0, 0, 0, time.UTC),
							},
						},
					},
					"Link Prop": gotion.PageProperty{
						ID:   "fFC%5B",
						Type: gotion.DBPropTypeURL,
						URL:  utils.StrPtr("https://github.com/asstart/gotion"),
					},
					"Text Prop": gotion.PageProperty{
						ID:   "jWtn",
						Type: gotion.DBPropTypeRichText,
						RichText: []gotion.RichText{
							gotion.RichText{
								Type: gotion.TextRichTextType,
								Text: &gotion.Text{
									Content: "Prop ",
								},
								Annotations: &gotion.Annotations{
									Bold: true,
								},
								PlainText: "Prop ",
							},
							gotion.RichText{
								Type: gotion.TextRichTextType,
								Text: &gotion.Text{
									Content: "with ",
								},
								Annotations: &gotion.Annotations{
									Italic: true,
								},
								PlainText: "with ",
							},
							gotion.RichText{
								Type: gotion.TextRichTextType,
								Text: &gotion.Text{
									Content: "many ",
								},
								Annotations: &gotion.Annotations{
									Italic:        true,
									Strikethrough: true,
								},
								PlainText: "many ",
							},
							gotion.RichText{
								Type: gotion.TextRichTextType,
								Text: &gotion.Text{
									Content: "styles",
								},
								Annotations: &gotion.Annotations{
									Italic:        true,
									Strikethrough: true,
									Code:          true,
								},
								PlainText: "styles",
							},
						},
					},
					"Created Time Prop": gotion.PageProperty{
						ID:   "w%3Fp%7C",
						Type: gotion.DBPropTypeCreatedTime,
						CreatedTime: &gotion.DateTimeWrap{
							Datetime: time.Date(2022, 6, 19, 13, 7, 0, 0, time.UTC),
						},
					},
					"Notion File Prop": {
						ID:   "wXSX",
						Type: gotion.DBPropTypeFiles,
						Files: []gotion.FileDescriptor{
							gotion.FileDescriptor{
								Name: "exmpl.txt",
								Type: gotion.NotionFileDescriptorType,
								NotionFile: &gotion.NotionFile{
									URL: "https://s3.us-west-2.amazonaws.com/secure.notion-static.com/84650ce1-7307-48a5-9139-01f503db7019/exmpl.txt?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Content-Sha256=UNSIGNED-PAYLOAD&X-Amz-Credential=AKIAT73L2G45EIPT3X45%2F20220619%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20220619T190041Z&X-Amz-Expires=3600&X-Amz-Signature=f101ca2bc0c7037c62b75e203fd5b27595d4aff000b586b1d8e1071649bfc69b&X-Amz-SignedHeaders=host&x-id=GetObject",
									ExpiryTime: gotion.DateTimeWrap{
										Datetime: time.Date(2022, 6, 19, 20, 0, 41, int(932*math.Pow(10, 6)), time.UTC),
									},
								},
							},
						},
					},
					"Select Prop": gotion.PageProperty{
						ID:   "%7CRrh",
						Type: gotion.DBPropTypeSelect,
						Select: &gotion.SelectPageProperty{
							ID:    "aeb7a7de-519f-4458-a77a-5058af1894e2",
							Name:  "t1",
							Color: gotion.Blue,
						},
					},
					"Prop Text With Link": {
						ID:   "Ew%3BV",
						Type: gotion.DBPropTypeRichText,
						RichText: []gotion.RichText{
							gotion.RichText{
								Type: gotion.TextRichTextType,
								Text: &gotion.Text{
									Content: "RFC ",
								},
								Annotations: &gotion.Annotations{},
								PlainText:   "RFC ",
							},
							gotion.RichText{
								Type: gotion.TextRichTextType,
								Text: &gotion.Text{
									Content: "https://www.rfc-editor.org/rfc/rfc3339.html",
									Link: &gotion.Link{
										URL: "https://www.rfc-editor.org/rfc/rfc3339.html",
									},
								},
								Annotations: &gotion.Annotations{},
								PlainText:   "https://www.rfc-editor.org/rfc/rfc3339.html",
								Href:        "https://www.rfc-editor.org/rfc/rfc3339.html",
							},
						},
					},
					"Prop Text With Date Mention": gotion.PageProperty{
						ID:   "Sv_%3C",
						Type: gotion.DBPropTypeRichText,
						RichText: []gotion.RichText{
							gotion.RichText{
								Type: gotion.TextRichTextType,
								Text: &gotion.Text{
									Content: "Mention ",
								},
								Annotations: &gotion.Annotations{},
								PlainText:   "Mention ",
							},
							gotion.RichText{
								Type: gotion.MentionRichTextType,
								Mention: &gotion.Mention{
									Type: gotion.DateMentionType,
									Date: &gotion.DatePageProperty{
										Start: gotion.DateTimeWrap{
											time.Date(2022, 6, 20, 0, 0, 0, 0, time.UTC),
										},
									},
								},
								Annotations: &gotion.Annotations{},
								PlainText:   "2022-06-20 → ",
							},
							gotion.RichText{
								Type: gotion.TextRichTextType,
								Text: &gotion.Text{
									Content: " ",
								},
								Annotations: &gotion.Annotations{},
								PlainText:   " ",
							},
						},
					},
					"Prop Text With Mention": gotion.PageProperty{
						ID:   "TpMl",
						Type: gotion.DBPropTypeRichText,
						RichText: []gotion.RichText{
							gotion.RichText{
								Type: gotion.TextRichTextType,
								Text: &gotion.Text{
									Content: "Mention ",
								},
								Annotations: &gotion.Annotations{},
								PlainText:   "Mention ",
							},
							gotion.RichText{
								Type: gotion.MentionRichTextType,
								Mention: &gotion.Mention{
									Type: gotion.UserMentionType,
									User: &gotion.User{
										Object: "user",
										ID:     "c7f2ae70-6b98-438f-8564-c59a71d7b3a4",
										Name:   "Имя",
										Type:   gotion.PersonType,
										Person: &gotion.Person{
											Email: "test@mail.dev",
										},
									},
								},
								Annotations: &gotion.Annotations{},
								PlainText:   "@Имя",
							},
							gotion.RichText{
								Type: gotion.TextRichTextType,
								Text: &gotion.Text{
									Content: " ",
								},
								Annotations: &gotion.Annotations{},
								PlainText:   " ",
							},
						},
					},
					"Prop Text With Mention Database": gotion.PageProperty{
						ID:   "r%5CU%3D",
						Type: gotion.DBPropTypeRichText,
						RichText: []gotion.RichText{
							gotion.RichText{
								Type: gotion.TextRichTextType,
								Text: &gotion.Text{
									Content: "Mention ",
								},
								Annotations: &gotion.Annotations{},
								PlainText:   "Mention ",
							},
							gotion.RichText{
								Type: gotion.MentionRichTextType,
								Mention: &gotion.Mention{
									Type: gotion.DatabaseMentionType,
									Database: &gotion.IDWrap{
										ID: "ba96e9ad-4600-47a7-a868-5f3328618910",
									},
								},
								Annotations: &gotion.Annotations{},
								PlainText:   "Rich Text prop db",
								Href:        "https://www.notion.so/ba96e9ad460047a7a8685f3328618910",
							},
							gotion.RichText{
								Type: gotion.TextRichTextType,
								Text: &gotion.Text{
									Content: " ",
								},
								Annotations: &gotion.Annotations{},
								PlainText:   " ",
							},
						},
					},
					"Prop Text With Mention Page": gotion.PageProperty{
						ID:   "yb%3Ew",
						Type: gotion.DBPropTypeRichText,
						RichText: []gotion.RichText{
							gotion.RichText{
								Type: gotion.TextRichTextType,
								Text: &gotion.Text{
									Content: "Mention ",
								},
								Annotations: &gotion.Annotations{},
								PlainText:   "Mention ",
							},
							gotion.RichText{
								Type: gotion.MentionRichTextType,
								Mention: &gotion.Mention{
									Type: gotion.PageMentionType,
									Page: &gotion.IDWrap{
										"548b9b8e-89a3-4e01-8dc4-35f76b1cb653",
									},
								},
								Annotations: &gotion.Annotations{},
								PlainText:   "Row1",
								Href:        "https://www.notion.so/548b9b8e89a34e018dc435f76b1cb653",
							},
							gotion.RichText{
								Type: gotion.TextRichTextType,
								Text: &gotion.Text{
									Content: " ",
								},
								Annotations: &gotion.Annotations{},
								PlainText:   " ",
							},
						},
					},
					"Property Equation": gotion.PageProperty{
						ID:   "zgh%7C",
						Type: gotion.DBPropTypeRichText,
						RichText: []gotion.RichText{
							gotion.RichText{
								Type: gotion.EquationRichTextType,
								Equation: &gotion.Equation{
									Expression: "E = mc^2",
								},
								Annotations: &gotion.Annotations{},
								PlainText:   "E = mc^2",
							},
							gotion.RichText{
								Type: gotion.TextRichTextType,
								Text: &gotion.Text{
									Content: " ",
								},
								Annotations: &gotion.Annotations{},
								PlainText:   " ",
							},
						},
					},
					"Related Rich Text prop db": gotion.PageProperty{
						ID:   "%5B%3Fnd",
						Type: gotion.DBPropTypeRelation,
						Relation: []gotion.IDWrap{
							gotion.IDWrap{
								ID: "e0f73c2d-7a2e-4fc2-9f28-73380ed76512",
							},
						},
					},
					"Rollup Property": gotion.PageProperty{
						ID:   "KOL%3F",
						Type: gotion.DBPropTypeRollup,
						Rollup: &gotion.RollupPageProperty{
							Type:     gotion.NumberRollupPropertyType,
							Number:   utils.FloatPtr(30),
							Function: gotion.Max,
						},
					},
					"Name": gotion.PageProperty{
						ID:   "title",
						Type: gotion.DBPropTypeTitle,
						Title: []gotion.RichText{
							gotion.RichText{
								Type: gotion.TextRichTextType,
								Text: &gotion.Text{
									Content: "Title with many diff styles",
								},
								Annotations: &gotion.Annotations{},
								PlainText:   "Title with many diff styles",
							},
						},
					},
				},
				URL: "https://www.notion.so/Title-with-many-diff-styles-e0f73c2d7a2e4fc29f2873380ed76512",
			},
		},
		{
			Source: `
			{
				"object": "page",
				"id": "96c917eb-d00b-4c7f-9e3d-6cdce17308fb",
				"created_time": "2022-06-20T01:11:00.000Z",
				"last_edited_time": "2022-06-20T01:11:00.000Z",
				"created_by": {
				  "object": "user",
				  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
				},
				"last_edited_by": {
				  "object": "user",
				  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
				},
				"cover": null,
				"icon": null,
				"parent": {
				  "type": "database_id",
				  "database_id": "eb65248f-c5f5-488b-a6e0-c30d1c9faaa2"
				},
				"archived": false,
				"properties": {
				  "Rollup Relation": {
					"id": "%3AKWH",
					"type": "rollup",
					"rollup": {
					  "type": "array",
					  "array": [
						{
						  "type": "relation",
						  "relation": [
							{
							  "id": "e0f73c2d-7a2e-4fc2-9f28-73380ed76512"
							}
						  ]
						}
					  ],
					  "function": "show_original"
					}
				  },
				  "Rollup Phone ": {
					"id": "%3CVAv",
					"type": "rollup",
					"rollup": {
					  "type": "array",
					  "array": [
						{
						  "type": "phone_number",
						  "phone_number": "+79998887766"
						}
					  ],
					  "function": "show_original"
					}
				  },
				  "Rollup Earliest Date": {
					"id": "%3DBkP",
					"type": "rollup",
					"rollup": {
					  "type": "date",
					  "date": {
						"start": "2022-06-20T01:30:00.000+00:00",
						"end": null,
						"time_zone": null
					  },
					  "function": "earliest_date"
					}
				  },
				  "Rollup Title": {
					"id": "%3F_%5EM",
					"type": "rollup",
					"rollup": {
					  "type": "array",
					  "array": [
						{
						  "type": "title",
						  "title": [
							{
							  "type": "text",
							  "text": {
								"content": "Title with many diff styles",
								"link": null
							  },
							  "annotations": {
								"bold": false,
								"italic": false,
								"strikethrough": false,
								"underline": false,
								"code": false,
								"color": "default"
							  },
							  "plain_text": "Title with many diff styles",
							  "href": null
							}
						  ]
						}
					  ],
					  "function": "show_original"
					}
				  },
				  "Rollup Date": {
					"id": "Dv_u",
					"type": "rollup",
					"rollup": {
					  "type": "array",
					  "array": [
						{
						  "type": "date",
						  "date": {
							"start": "2022-06-16",
							"end": null,
							"time_zone": null
						  }
						}
					  ],
					  "function": "show_original"
					}
				  },
				  "Rollup Email": {
					"id": "EAe%5B",
					"type": "rollup",
					"rollup": {
					  "type": "array",
					  "array": [
						{
						  "type": "email",
						  "email": "test@mail.dev"
						}
					  ],
					  "function": "show_original"
					}
				  },
				  "Rollup Created By": {
					"id": "JJew",
					"type": "rollup",
					"rollup": {
					  "type": "array",
					  "array": [
						{
						  "type": "created_by",
						  "created_by": {
							"object": "user",
							"id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
						  }
						}
					  ],
					  "function": "show_original"
					}
				  },
				  "Rollup Link": {
					"id": "LfXP",
					"type": "rollup",
					"rollup": {
					  "type": "array",
					  "array": [
						{
						  "type": "url",
						  "url": "https://github.com/asstart/gotion"
						}
					  ],
					  "function": "show_original"
					}
				  },
				  "Rollup Percent": {
					"id": "N%60%3A%3B",
					"type": "rollup",
					"rollup": {
					  "type": "number",
					  "number": 0,
					  "function": "percent_empty"
					}
				  },
				  "Rollup User": {
					"id": "OIYq",
					"type": "rollup",
					"rollup": {
					  "type": "array",
					  "array": [
						{
						  "type": "people",
						  "people": [
							{
							  "object": "user",
							  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
							}
						  ]
						}
					  ],
					  "function": "show_original"
					}
				  },
				  "Rollup Created Date": {
					"id": "PhsD",
					"type": "rollup",
					"rollup": {
					  "type": "array",
					  "array": [
						{
						  "type": "created_time",
						  "created_time": "2022-06-19T13:07:00.000Z"
						}
					  ],
					  "function": "show_original"
					}
				  },
				  "Rollup Count": {
					"id": "QmRZ",
					"type": "rollup",
					"rollup": {
					  "type": "number",
					  "number": 1,
					  "function": "count"
					}
				  },
				  "Relation": {
					"id": "SE%5DE",
					"type": "relation",
					"relation": [
					  {
						"id": "e0f73c2d-7a2e-4fc2-9f28-73380ed76512"
					  }
					]
				  },
				  "Rollup File": {
					"id": "TO%5Eq",
					"type": "rollup",
					"rollup": {
					  "type": "array",
					  "array": [
						{
						  "type": "files",
						  "files": []
						}
					  ],
					  "function": "show_original"
					}
				  },
				  "Rollup Checkbox": {
					"id": "Xt%5Cm",
					"type": "rollup",
					"rollup": {
					  "type": "array",
					  "array": [
						{
						  "type": "checkbox",
						  "checkbox": true
						}
					  ],
					  "function": "show_original"
					}
				  },
				  "Rollup Single Select": {
					"id": "YAdu",
					"type": "rollup",
					"rollup": {
					  "type": "array",
					  "array": [
						{
						  "type": "select",
						  "select": {
							"id": "aeb7a7de-519f-4458-a77a-5058af1894e2",
							"name": "t1",
							"color": "blue"
						  }
						}
					  ],
					  "function": "show_original"
					}
				  },
				  "Rollup Number": {
					"id": "dMVf",
					"type": "rollup",
					"rollup": {
					  "type": "array",
					  "array": [
						{
						  "type": "number",
						  "number": 30
						}
					  ],
					  "function": "show_original"
					}
				  },
				  "Rollup Multi Select": {
					"id": "htzg",
					"type": "rollup",
					"rollup": {
					  "type": "array",
					  "array": [
						{
						  "type": "multi_select",
						  "multi_select": [
							{
							  "id": "6c0151da-c522-48e3-812b-428044bf9338",
							  "name": "t11",
							  "color": "red"
							},
							{
							  "id": "01ba9b69-9953-4d2d-9cb7-8a32ef179388",
							  "name": "t12",
							  "color": "pink"
							}
						  ]
						}
					  ],
					  "function": "show_original"
					}
				  },
				  "Rollup Median": {
					"id": "l%5Cmm",
					"type": "rollup",
					"rollup": {
					  "type": "number",
					  "number": 30,
					  "function": "median"
					}
				  },
				  "Rollup Formula": {
					"id": "qH%60G",
					"type": "rollup",
					"rollup": {
					  "type": "array",
					  "array": [
						{
						  "type": "formula",
						  "formula": {
							"type": "boolean",
							"boolean": true
						  }
						}
					  ],
					  "function": "show_original"
					}
				  },
				  "Rollup Rich Text": {
					"id": "w%7BWX",
					"type": "rollup",
					"rollup": {
					  "type": "array",
					  "array": [
						{
						  "type": "rich_text",
						  "rich_text": [
							{
							  "type": "text",
							  "text": {
								"content": "Prop ",
								"link": null
							  },
							  "annotations": {
								"bold": true,
								"italic": false,
								"strikethrough": false,
								"underline": false,
								"code": false,
								"color": "default"
							  },
							  "plain_text": "Prop ",
							  "href": null
							},
							{
							  "type": "text",
							  "text": {
								"content": "with ",
								"link": null
							  },
							  "annotations": {
								"bold": false,
								"italic": true,
								"strikethrough": false,
								"underline": false,
								"code": false,
								"color": "default"
							  },
							  "plain_text": "with ",
							  "href": null
							},
							{
							  "type": "text",
							  "text": {
								"content": "many ",
								"link": null
							  },
							  "annotations": {
								"bold": false,
								"italic": true,
								"strikethrough": true,
								"underline": false,
								"code": false,
								"color": "default"
							  },
							  "plain_text": "many ",
							  "href": null
							},
							{
							  "type": "text",
							  "text": {
								"content": "styles",
								"link": null
							  },
							  "annotations": {
								"bold": false,
								"italic": true,
								"strikethrough": true,
								"underline": false,
								"code": true,
								"color": "default"
							  },
							  "plain_text": "styles",
							  "href": null
							}
						  ]
						}
					  ],
					  "function": "show_original"
					}
				  },
				  "Name": {
					"id": "title",
					"type": "title",
					"title": [
					  {
						"type": "text",
						"text": {
						  "content": "Rollup Test",
						  "link": null
						},
						"annotations": {
						  "bold": false,
						  "italic": false,
						  "strikethrough": false,
						  "underline": false,
						  "code": false,
						  "color": "default"
						},
						"plain_text": "Rollup Test",
						"href": null
					  }
					]
				  }
				},
				"url": "https://www.notion.so/Rollup-Test-96c917ebd00b4c7f9e3d6cdce17308fb"
			  }
			`,
			Expected: gotion.Page{
				Object: "page",
				ID:     "96c917eb-d00b-4c7f-9e3d-6cdce17308fb",
				CreatedTime: gotion.DateTimeWrap{
					Datetime: time.Date(2022, 6, 20, 1, 11, 0, 0, time.UTC),
				},
				LastEditedTime: gotion.DateTimeWrap{
					Datetime: time.Date(2022, 6, 20, 1, 11, 0, 0, time.UTC),
				},
				CreatedBy: gotion.User{
					Object: "user",
					ID:     "c7f2ae70-6b98-438f-8564-c59a71d7b3a4",
				},
				LastEditedBy: gotion.User{
					Object: "user",
					ID:     "c7f2ae70-6b98-438f-8564-c59a71d7b3a4",
				},
				Parent: gotion.PageParent{
					Type:       gotion.PageParentDB,
					DatabaseID: "eb65248f-c5f5-488b-a6e0-c30d1c9faaa2",
				},
				Properties: gotion.PageProperties{
					"Rollup Relation": {
						ID:   "%3AKWH",
						Type: gotion.DBPropTypeRollup,
						Rollup: &gotion.RollupPageProperty{
							Type: gotion.ArrayRollupPropertyType,
							Array: []gotion.RollupArrayValue{
								gotion.RollupArrayValue{
									Type: gotion.DBPropTypeRelation,
									Relation: []gotion.IDWrap{
										gotion.IDWrap{
											ID: "e0f73c2d-7a2e-4fc2-9f28-73380ed76512",
										},
									},
								},
							},
							Function: gotion.ShowOriginal,
						},
					},
					"Rollup Phone ": {
						ID:   "%3CVAv",
						Type: gotion.DBPropTypeRollup,
						Rollup: &gotion.RollupPageProperty{
							Type: gotion.ArrayRollupPropertyType,
							Array: []gotion.RollupArrayValue{
								gotion.RollupArrayValue{
									Type:        gotion.DBPropTypePhoneNumber,
									PhoneNumber: utils.StrPtr("+79998887766"),
								},
							},
							Function: gotion.ShowOriginal,
						},
					},
					"Rollup Earliest Date": {
						ID:   "%3DBkP",
						Type: gotion.DBPropTypeRollup,
						Rollup: &gotion.RollupPageProperty{
							Type: gotion.DateRollupPropertyType,
							Date: &gotion.DatePageProperty{
								Start: gotion.DateTimeWrap{
									Datetime: time.Date(2022, 6, 20, 1, 30, 0, 0, time.FixedZone("", 0)),
								},
							},
							Function: gotion.EarliestDate,
						},
					},
					"Rollup Title": {
						ID:   "%3F_%5EM",
						Type: gotion.DBPropTypeRollup,
						Rollup: &gotion.RollupPageProperty{
							Type: gotion.ArrayRollupPropertyType,
							Array: []gotion.RollupArrayValue{
								gotion.RollupArrayValue{
									Type: gotion.DBPropTypeTitle,
									Title: []gotion.RichText{
										gotion.RichText{
											Type: gotion.TextRichTextType,
											Text: &gotion.Text{
												Content: "Title with many diff styles",
											},
											Annotations: &gotion.Annotations{},
											PlainText:   "Title with many diff styles",
										},
									},
								},
							},
							Function: gotion.ShowOriginal,
						},
					},
					"Rollup Date": gotion.PageProperty{
						ID:   "Dv_u",
						Type: gotion.DBPropTypeRollup,
						Rollup: &gotion.RollupPageProperty{
							Type: gotion.ArrayRollupPropertyType,
							Array: []gotion.RollupArrayValue{
								gotion.RollupArrayValue{
									Type: gotion.DBPropTypeDate,
									Date: &gotion.DatePageProperty{
										Start: gotion.DateTimeWrap{
											Datetime: time.Date(2022, 6, 16, 0, 0, 0, 0, time.UTC),
										},
									},
								},
							},
							Function: gotion.ShowOriginal,
						},
					},
					"Rollup Email": gotion.PageProperty{
						ID:   "EAe%5B",
						Type: gotion.DBPropTypeRollup,
						Rollup: &gotion.RollupPageProperty{
							Type: gotion.ArrayRollupPropertyType,
							Array: []gotion.RollupArrayValue{
								gotion.RollupArrayValue{
									Type:  gotion.DBPropTypeEmail,
									Email: utils.StrPtr("test@mail.dev"),
								},
							},
							Function: gotion.ShowOriginal,
						},
					},
					"Rollup Created By": {
						ID:   "JJew",
						Type: gotion.DBPropTypeRollup,
						Rollup: &gotion.RollupPageProperty{
							Type: gotion.ArrayRollupPropertyType,
							Array: []gotion.RollupArrayValue{
								gotion.RollupArrayValue{
									Type: gotion.DBPropTypeCreatedBy,
									CreatedBy: &gotion.User{
										Object: "user",
										ID:     "c7f2ae70-6b98-438f-8564-c59a71d7b3a4",
									},
								},
							},
							Function: gotion.ShowOriginal,
						},
					},
					"Rollup Link": {
						ID:   "LfXP",
						Type: gotion.DBPropTypeRollup,
						Rollup: &gotion.RollupPageProperty{
							Type: gotion.ArrayRollupPropertyType,
							Array: []gotion.RollupArrayValue{
								gotion.RollupArrayValue{
									Type: gotion.DBPropTypeURL,
									URL:  utils.StrPtr("https://github.com/asstart/gotion"),
								},
							},
							Function: gotion.ShowOriginal,
						},
					},
					"Rollup Percent": {
						ID:   "N%60%3A%3B",
						Type: gotion.DBPropTypeRollup,
						Rollup: &gotion.RollupPageProperty{
							Type:     gotion.NumberRollupPropertyType,
							Number:   utils.FloatPtr(0),
							Function: gotion.PercentEmpty,
						},
					},
					"Rollup User": {
						ID:   "OIYq",
						Type: gotion.DBPropTypeRollup,
						Rollup: &gotion.RollupPageProperty{
							Type: gotion.ArrayRollupPropertyType,
							Array: []gotion.RollupArrayValue{
								gotion.RollupArrayValue{
									Type: gotion.DBPropTypePeople,
									People: []gotion.User{
										gotion.User{
											Object: "user",
											ID:     "c7f2ae70-6b98-438f-8564-c59a71d7b3a4",
										},
									},
								},
							},
							Function: gotion.ShowOriginal,
						},
					},
					"Rollup Created Date": gotion.PageProperty{
						ID:   "PhsD",
						Type: gotion.DBPropTypeRollup,
						Rollup: &gotion.RollupPageProperty{
							Type: gotion.ArrayRollupPropertyType,
							Array: []gotion.RollupArrayValue{
								gotion.RollupArrayValue{
									Type: gotion.DBPropTypeCreatedTime,
									CreatedTime: &gotion.DateTimeWrap{
										Datetime: time.Date(2022, 6, 19, 13, 7, 0, 0, time.UTC),
									},
								},
							},
							Function: gotion.ShowOriginal,
						},
					},
					"Rollup Count": gotion.PageProperty{
						ID:   "QmRZ",
						Type: gotion.DBPropTypeRollup,
						Rollup: &gotion.RollupPageProperty{
							Type:     gotion.NumberRollupPropertyType,
							Number:   utils.FloatPtr(1),
							Function: gotion.Count,
						},
					},
					"Relation": gotion.PageProperty{
						ID:   "SE%5DE",
						Type: gotion.DBPropTypeRelation,
						Relation: []gotion.IDWrap{
							gotion.IDWrap{
								ID: "e0f73c2d-7a2e-4fc2-9f28-73380ed76512",
							},
						},
					},
					"Rollup File": gotion.PageProperty{
						ID:   "TO%5Eq",
						Type: gotion.DBPropTypeRollup,
						Rollup: &gotion.RollupPageProperty{
							Type: gotion.ArrayRollupPropertyType,
							Array: []gotion.RollupArrayValue{
								gotion.RollupArrayValue{
									Type:  gotion.DBPropTypeFiles,
									Files: []gotion.FileDescriptor{},
								},
							},
							Function: gotion.ShowOriginal,
						},
					},
					"Rollup Checkbox": gotion.PageProperty{
						ID:   "Xt%5Cm",
						Type: gotion.DBPropTypeRollup,
						Rollup: &gotion.RollupPageProperty{
							Type: gotion.ArrayRollupPropertyType,
							Array: []gotion.RollupArrayValue{
								gotion.RollupArrayValue{
									Type:     gotion.DBPropTypeCheckbox,
									Checkbox: utils.BoolPtr(true),
								},
							},
							Function: gotion.ShowOriginal,
						},
					},
					"Rollup Single Select": gotion.PageProperty{
						ID:   "YAdu",
						Type: gotion.DBPropTypeRollup,
						Rollup: &gotion.RollupPageProperty{
							Type: gotion.ArrayRollupPropertyType,
							Array: []gotion.RollupArrayValue{
								gotion.RollupArrayValue{
									Type: gotion.DBPropTypeSelect,
									Select: &gotion.SelectPageProperty{
										ID:    "aeb7a7de-519f-4458-a77a-5058af1894e2",
										Name:  "t1",
										Color: gotion.Blue,
									},
								},
							},
							Function: gotion.ShowOriginal,
						},
					},
					"Rollup Number": gotion.PageProperty{
						ID:   "dMVf",
						Type: gotion.DBPropTypeRollup,
						Rollup: &gotion.RollupPageProperty{
							Type: gotion.ArrayRollupPropertyType,
							Array: []gotion.RollupArrayValue{
								gotion.RollupArrayValue{
									Type:   gotion.DBPropTypeNumber,
									Number: utils.FloatPtr(30),
								},
							},
							Function: gotion.ShowOriginal,
						},
					},
					"Rollup Multi Select": gotion.PageProperty{
						ID:   "htzg",
						Type: gotion.DBPropTypeRollup,
						Rollup: &gotion.RollupPageProperty{
							Type: gotion.ArrayRollupPropertyType,
							Array: []gotion.RollupArrayValue{
								gotion.RollupArrayValue{
									Type: gotion.DBPropTypeMultiSelect,
									MultiSelect: []gotion.SelectPageProperty{
										gotion.SelectPageProperty{
											ID:    "6c0151da-c522-48e3-812b-428044bf9338",
											Name:  "t11",
											Color: gotion.Red,
										},
										gotion.SelectPageProperty{
											ID:    "01ba9b69-9953-4d2d-9cb7-8a32ef179388",
											Name:  "t12",
											Color: gotion.Pink,
										},
									},
								},
							},
							Function: gotion.ShowOriginal,
						},
					},
					"Rollup Median": gotion.PageProperty{
						ID: "l%5Cmm",
						Type: gotion.DBPropTypeRollup,
						Rollup: &gotion.RollupPageProperty{
							Type: gotion.NumberRollupPropertyType,
							Number: utils.FloatPtr(30),
							Function: gotion.Median,
						},
					},
					"Rollup Formula": gotion.PageProperty{
						ID:   "qH%60G",
						Type: gotion.DBPropTypeRollup,
						Rollup: &gotion.RollupPageProperty{
							Type: gotion.ArrayRollupPropertyType,
							Array: []gotion.RollupArrayValue{
								gotion.RollupArrayValue{
									Type: gotion.DBPropTypeFormula,
									Formula: &gotion.FormulaPageProperty{
										Type:           gotion.BoolFormulaType,
										BooleanFormula: utils.BoolPtr(true),
									},
								},
							},
							Function: gotion.ShowOriginal,
						},
					},
					"Rollup Rich Text": gotion.PageProperty{
						ID:   "w%7BWX",
						Type: gotion.DBPropTypeRollup,
						Rollup: &gotion.RollupPageProperty{
							Type: gotion.ArrayRollupPropertyType,
							Array: []gotion.RollupArrayValue{
								gotion.RollupArrayValue{
									Type: gotion.DBPropTypeRichText,
									RichText: []gotion.RichText{
										gotion.RichText{
											Type: gotion.TextRichTextType,
											Text: &gotion.Text{
												Content: "Prop ",
											},
											Annotations: &gotion.Annotations{
												Bold: true,
											},
											PlainText: "Prop ",
										},
										gotion.RichText{
											Type: gotion.TextRichTextType,
											Text: &gotion.Text{
												Content: "with ",
											},
											Annotations: &gotion.Annotations{
												Italic: true,
											},
											PlainText: "with ",
										},
										gotion.RichText{
											Type: gotion.TextRichTextType,
											Text: &gotion.Text{
												Content: "many ",
											},
											Annotations: &gotion.Annotations{
												Italic:        true,
												Strikethrough: true,
											},
											PlainText: "many ",
										},
										gotion.RichText{
											Type: gotion.TextRichTextType,
											Text: &gotion.Text{
												Content: "styles",
											},
											Annotations: &gotion.Annotations{
												Italic:        true,
												Strikethrough: true,
												Code:          true,
											},
											PlainText: "styles",
										},
									},
								},
							},
							Function: gotion.ShowOriginal,
						},
					},
					"Name": gotion.PageProperty{
						ID:   "title",
						Type: gotion.DBPropTypeTitle,
						Title: []gotion.RichText{
							gotion.RichText{
								Type: gotion.TextRichTextType,
								Text: &gotion.Text{
									Content: "Rollup Test",
								},
								Annotations: &gotion.Annotations{},
								PlainText:   "Rollup Test",
							},
						},
					},
				},
				URL: "https://www.notion.so/Rollup-Test-96c917ebd00b4c7f9e3d6cdce17308fb",
			},
		},
	}
	for _, tuple := range se {
		var res = gotion.Page{}
		err := json.Unmarshal([]byte(tuple.Source), &res)
		utils.AssertNil(t, err)
		assert.Equal(t, tuple.Expected, res)
	}
}
