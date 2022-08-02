package gotion_test

import (
	"encoding/json"
	"github.com/asstart/gotion"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnmarshalRetrieveParagraphBlock(t *testing.T) {
	source := `
	{
		"object": "block",
		"id": "38457669-f3db-486d-879d-3ab81aab0868",
		"parent": {
		  "type": "page_id",
		  "page_id": "6d8c8ae0-b4b9-410e-af67-f85cb151398f"
		},
		"created_time": "2022-08-02T22:21:00.000Z",
		"last_edited_time": "2022-08-02T22:21:00.000Z",
		"created_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"last_edited_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"has_children": false,
		"archived": false,
		"type": "paragraph",
		"paragraph": {
		  "rich_text": [
			{
			  "type": "text",
			  "text": {
				"content": "Paragraph Test",
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
			  "plain_text": "Paragraph Test",
			  "href": null
			}
		  ],
		  "color": "default"
		}
	  }
	  `
	var expected gotion.Block
	var res gotion.Block
	err := json.Unmarshal([]byte(source), &res)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}

func TestUnmarshalH1Block(t *testing.T) {
	source := `
	{
		"object": "block",
		"id": "2a664611-b28d-496d-9d59-9a34ce958eb4",
		"parent": {
		  "type": "page_id",
		  "page_id": "6d8c8ae0-b4b9-410e-af67-f85cb151398f"
		},
		"created_time": "2022-08-02T22:22:00.000Z",
		"last_edited_time": "2022-08-02T22:22:00.000Z",
		"created_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"last_edited_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"has_children": false,
		"archived": false,
		"type": "heading_1",
		"heading_1": {
		  "rich_text": [
			{
			  "type": "text",
			  "text": {
				"content": "Heading 1 test",
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
			  "plain_text": "Heading 1 test",
			  "href": null
			}
		  ],
		  "color": "default"
		}
	  }
	  `
	var expected gotion.Block
	var res gotion.Block
	err := json.Unmarshal([]byte(source), &res)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}

func TestUnmarshalH2Block(t *testing.T) {
	source := `
	{
		"object": "block",
		"id": "0a031d54-c627-4e67-9f93-48a7b747fdea",
		"parent": {
		  "type": "page_id",
		  "page_id": "6d8c8ae0-b4b9-410e-af67-f85cb151398f"
		},
		"created_time": "2022-08-02T22:22:00.000Z",
		"last_edited_time": "2022-08-02T22:22:00.000Z",
		"created_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"last_edited_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"has_children": false,
		"archived": false,
		"type": "heading_2",
		"heading_2": {
		  "rich_text": [
			{
			  "type": "text",
			  "text": {
				"content": "Heading 2 test",
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
			  "plain_text": "Heading 2 test",
			  "href": null
			}
		  ],
		  "color": "default"
		}
	  }
	  
	  `
	var expected gotion.Block
	var res gotion.Block
	err := json.Unmarshal([]byte(source), &res)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}
func TestUnmarshalH3Block(t *testing.T) {
	source := `
	{
		"object": "block",
		"id": "572e40f5-1695-4b86-a22c-f9de4ff71324",
		"parent": {
		  "type": "page_id",
		  "page_id": "6d8c8ae0-b4b9-410e-af67-f85cb151398f"
		},
		"created_time": "2022-08-02T22:22:00.000Z",
		"last_edited_time": "2022-08-02T22:22:00.000Z",
		"created_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"last_edited_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"has_children": false,
		"archived": false,
		"type": "heading_3",
		"heading_3": {
		  "rich_text": [
			{
			  "type": "text",
			  "text": {
				"content": "Heading 3 test",
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
			  "plain_text": "Heading 3 test",
			  "href": null
			}
		  ],
		  "color": "default"
		}
	  }
	  `
	var expected gotion.Block
	var res gotion.Block
	err := json.Unmarshal([]byte(source), &res)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}
func TestUnmarshalCalloutBlock(t *testing.T) {
	source := `
	{
		"object": "block",
		"id": "e004f495-8876-45a6-b5ee-94ae7b8ce6ff",
		"parent": {
		  "type": "page_id",
		  "page_id": "6d8c8ae0-b4b9-410e-af67-f85cb151398f"
		},
		"created_time": "2022-08-02T22:22:00.000Z",
		"last_edited_time": "2022-08-02T22:22:00.000Z",
		"created_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"last_edited_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"has_children": false,
		"archived": false,
		"type": "callout",
		"callout": {
		  "rich_text": [
			{
			  "type": "text",
			  "text": {
				"content": "Callout test",
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
			  "plain_text": "Callout test",
			  "href": null
			}
		  ],
		  "icon": {
			"type": "emoji",
			"emoji": "💡"
		  },
		  "color": "gray_background"
		}
	  }
	  `
	var expected gotion.Block
	var res gotion.Block
	err := json.Unmarshal([]byte(source), &res)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}
func TestUnmarshalBulletedList(t *testing.T) {
	source := `

	"object": "block",
	"id": "a362db34-4253-4fb5-ab0e-9c86ccfd3233",
	"parent": {
	  "type": "page_id",
	  "page_id": "6d8c8ae0-b4b9-410e-af67-f85cb151398f"
	},
	"created_time": "2022-08-02T22:22:00.000Z",
	"last_edited_time": "2022-08-02T22:22:00.000Z",
	"created_by": {
	  "object": "user",
	  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
	},
	"last_edited_by": {
	  "object": "user",
	  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
	},
	"has_children": false,
	"archived": false,
	"type": "bulleted_list_item",
	"bulleted_list_item": {
	  "rich_text": [
		{
		  "type": "text",
		  "text": {
			"content": "Bullet list test item 1",
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
		  "plain_text": "Bullet list test item 1",
		  "href": null
		}
	  ],
	  "color": "default"
	}
  }
	  `
	var expected gotion.Block
	var res gotion.Block
	err := json.Unmarshal([]byte(source), &res)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}
func TestUnmarshalNumListBlock(t *testing.T) {
	source := `
	{
		"object": "block",
		"id": "34fa8f22-b6c6-46b3-999d-8885524acfee",
		"parent": {
		  "type": "page_id",
		  "page_id": "6d8c8ae0-b4b9-410e-af67-f85cb151398f"
		},
		"created_time": "2022-08-02T22:23:00.000Z",
		"last_edited_time": "2022-08-02T22:23:00.000Z",
		"created_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"last_edited_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"has_children": false,
		"archived": false,
		"type": "numbered_list_item",
		"numbered_list_item": {
		  "rich_text": [
			{
			  "type": "text",
			  "text": {
				"content": "Numbered list test item 1",
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
			  "plain_text": "Numbered list test item 1",
			  "href": null
			}
		  ],
		  "color": "default"
		}
	  }
	  `
	var expected gotion.Block
	var res gotion.Block
	err := json.Unmarshal([]byte(source), &res)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}
func TestUnmarshalUnchectedTodoBlock(t *testing.T) {
	source := `
	{
		"object": "block",
		"id": "3cb64c3d-1f94-4e0f-8886-b4a44e2b594c",
		"parent": {
		  "type": "page_id",
		  "page_id": "6d8c8ae0-b4b9-410e-af67-f85cb151398f"
		},
		"created_time": "2022-08-02T22:23:00.000Z",
		"last_edited_time": "2022-08-02T22:23:00.000Z",
		"created_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"last_edited_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"has_children": false,
		"archived": false,
		"type": "to_do",
		"to_do": {
		  "rich_text": [
			{
			  "type": "text",
			  "text": {
				"content": "Todo test false checkbox",
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
			  "plain_text": "Todo test false checkbox",
			  "href": null
			}
		  ],
		  "checked": false,
		  "color": "default"
		}
	  }
	  `
	var expected gotion.Block
	var res gotion.Block
	err := json.Unmarshal([]byte(source), &res)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}
func TestUnmarshalCheckedTodoBlock(t *testing.T) {
	source := `
	{
		"object": "block",
		"id": "1728cbe7-e607-477b-b135-5b57e04e0eda",
		"parent": {
		  "type": "page_id",
		  "page_id": "6d8c8ae0-b4b9-410e-af67-f85cb151398f"
		},
		"created_time": "2022-08-02T22:23:00.000Z",
		"last_edited_time": "2022-08-02T22:23:00.000Z",
		"created_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"last_edited_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"has_children": false,
		"archived": false,
		"type": "to_do",
		"to_do": {
		  "rich_text": [
			{
			  "type": "text",
			  "text": {
				"content": "Todo test true checkox",
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
			  "plain_text": "Todo test true checkox",
			  "href": null
			}
		  ],
		  "checked": true,
		  "color": "default"
		}
	  }
	  `
	var expected gotion.Block
	var res gotion.Block
	err := json.Unmarshal([]byte(source), &res)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}
func TestUnmarshalToggleBlock(t *testing.T) {
	source := `
	{
		"object": "block",
		"id": "36f309f3-6f0c-44de-acf7-65abd707505b",
		"parent": {
		  "type": "page_id",
		  "page_id": "6d8c8ae0-b4b9-410e-af67-f85cb151398f"
		},
		"created_time": "2022-08-02T22:23:00.000Z",
		"last_edited_time": "2022-08-02T22:23:00.000Z",
		"created_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"last_edited_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"has_children": false,
		"archived": false,
		"type": "toggle",
		"toggle": {
		  "rich_text": [
			{
			  "type": "text",
			  "text": {
				"content": "Toggle test item 1",
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
			  "plain_text": "Toggle test item 1",
			  "href": null
			}
		  ],
		  "color": "default"
		}
	  }
	  `
	var expected gotion.Block
	var res gotion.Block
	err := json.Unmarshal([]byte(source), &res)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}
func TestUnmarshalCodeBlock(t *testing.T) {
	source := `
	{
		"object": "block",
		"id": "777a5bd1-04e3-4762-9c22-5a0f0a9f76c6",
		"parent": {
		  "type": "page_id",
		  "page_id": "6d8c8ae0-b4b9-410e-af67-f85cb151398f"
		},
		"created_time": "2022-08-02T22:24:00.000Z",
		"last_edited_time": "2022-08-02T22:24:00.000Z",
		"created_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"last_edited_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"has_children": false,
		"archived": false,
		"type": "code",
		"code": {
		  "caption": [],
		  "rich_text": [
			{
			  "type": "text",
			  "text": {
				"content": "Code test",
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
			  "plain_text": "Code test",
			  "href": null
			}
		  ],
		  "language": "shell"
		}
	  }
	  `
	var expected gotion.Block
	var res gotion.Block
	err := json.Unmarshal([]byte(source), &res)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}
func TestUnmarshalChildPageBlock(t *testing.T) {
	source := `
	{
		"object": "block",
		"id": "612f67f9-f046-41e4-810d-1a49de093c53",
		"parent": {
		  "type": "page_id",
		  "page_id": "6d8c8ae0-b4b9-410e-af67-f85cb151398f"
		},
		"created_time": "2022-08-02T22:25:00.000Z",
		"last_edited_time": "2022-08-02T22:25:00.000Z",
		"created_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"last_edited_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"has_children": false,
		"archived": false,
		"type": "child_page",
		"child_page": {
		  "title": "Single child page for test"
		}
	  }
	  `
	var expected gotion.Block
	var res gotion.Block
	err := json.Unmarshal([]byte(source), &res)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}
func TestUnmarshalChildDatabaseBlock(t *testing.T) {
	source := `

	"object": "block",
	"id": "b7052dc3-1b53-4840-afcb-039d893dbe2a",
	"parent": {
	  "type": "page_id",
	  "page_id": "6d8c8ae0-b4b9-410e-af67-f85cb151398f"
	},
	"created_time": "2022-08-02T22:26:00.000Z",
	"last_edited_time": "2022-08-02T22:26:00.000Z",
	"created_by": {
	  "object": "user",
	  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
	},
	"last_edited_by": {
	  "object": "user",
	  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
	},
	"has_children": false,
	"archived": false,
	"type": "child_database",
	"child_database": {
	  "title": "Single child database for test"
	}
  }
	  `
	var expected gotion.Block
	var res gotion.Block
	err := json.Unmarshal([]byte(source), &res)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}

//Looks wrong. TODO check
func TestUnmarshalEmbedBlock(t *testing.T) {
	source := `
	{
		"object": "block",
		"id": "4165c38a-c6b3-447f-a4cf-2de17d986589",
		"parent": {
		  "type": "page_id",
		  "page_id": "6d8c8ae0-b4b9-410e-af67-f85cb151398f"
		},
		"created_time": "2022-08-02T22:26:00.000Z",
		"last_edited_time": "2022-08-02T22:26:00.000Z",
		"created_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"last_edited_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"has_children": false,
		"archived": false,
		"type": "link_preview",
		"link_preview": {
		  "url": "https://github.com/asstart"
		}
	  }
	  `
	var expected gotion.Block
	var res gotion.Block
	err := json.Unmarshal([]byte(source), &res)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}
func TestUnmarshalExternalImageBlock(t *testing.T) {
	source := `
	{
		"object": "block",
		"id": "9ccc2240-2cf2-4c62-9ada-eba7fc7204a8",
		"parent": {
		  "type": "page_id",
		  "page_id": "6d8c8ae0-b4b9-410e-af67-f85cb151398f"
		},
		"created_time": "2022-08-02T22:29:00.000Z",
		"last_edited_time": "2022-08-02T22:30:00.000Z",
		"created_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"last_edited_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"has_children": false,
		"archived": false,
		"type": "image",
		"image": {
		  "caption": [],
		  "type": "external",
		  "external": {
			"url": "https://images.unsplash.com/photo-1616530940355-351fabd9524b?ixlib=rb-1.2.1&q=80&cs=tinysrgb&fm=jpg&crop=entropy"
		  }
		}
	  }
	  `
	var expected gotion.Block
	var res gotion.Block
	err := json.Unmarshal([]byte(source), &res)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}
func TestUnmarshalInternalImageBlock(t *testing.T) {
	source := `
	{
		"object": "block",
		"id": "4f4c9bf8-9c2f-4a39-982f-0639fcf68c76",
		"parent": {
		  "type": "page_id",
		  "page_id": "6d8c8ae0-b4b9-410e-af67-f85cb151398f"
		},
		"created_time": "2022-08-02T22:30:00.000Z",
		"last_edited_time": "2022-08-02T22:31:00.000Z",
		"created_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"last_edited_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"has_children": false,
		"archived": false,
		"type": "image",
		"image": {
		  "caption": [],
		  "type": "file",
		  "file": {
			"url": "https://s3.us-west-2.amazonaws.com/secure.notion-static.com/5df70500-ed2e-4ae9-ac1a-a5f29c073e99/1.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Content-Sha256=UNSIGNED-PAYLOAD&X-Amz-Credential=AKIAT73L2G45EIPT3X45%2F20220802%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20220802T231436Z&X-Amz-Expires=3600&X-Amz-Signature=8097e63d5e7a3b39422f12b4b04c9bdf9a2784de116bd29c9946fa20ac1287b2&X-Amz-SignedHeaders=host&x-id=GetObject",
			"expiry_time": "2022-08-03T00:14:36.458Z"
		  }
		}
	  }
	  `
	var expected gotion.Block
	var res gotion.Block
	err := json.Unmarshal([]byte(source), &res)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}
func TestUnmarshalExternalVideoBlock(t *testing.T) {
	source := `

	"object": "block",
	"id": "3990ee53-7c81-44a7-a179-33fa4603f6d2",
	"parent": {
	  "type": "page_id",
	  "page_id": "6d8c8ae0-b4b9-410e-af67-f85cb151398f"
	},
	"created_time": "2022-08-02T22:31:00.000Z",
	"last_edited_time": "2022-08-02T22:31:00.000Z",
	"created_by": {
	  "object": "user",
	  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
	},
	"last_edited_by": {
	  "object": "user",
	  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
	},
	"has_children": false,
	"archived": false,
	"type": "video",
	"video": {
	  "caption": [],
	  "type": "external",
	  "external": {
		"url": "https://www.youtube.com/watch?v=jfKfPfyJRdk"
	  }
	}
  }
	  `
	var expected gotion.Block
	var res gotion.Block
	err := json.Unmarshal([]byte(source), &res)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}
func TestUnmarshalInternalFileBlock(t *testing.T) {
	source := `
	{
		"object": "block",
		"id": "330005dc-b2c7-40ef-bbb3-982a5bee10c9",
		"parent": {
		  "type": "page_id",
		  "page_id": "6d8c8ae0-b4b9-410e-af67-f85cb151398f"
		},
		"created_time": "2022-08-02T22:32:00.000Z",
		"last_edited_time": "2022-08-02T22:32:00.000Z",
		"created_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"last_edited_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"has_children": false,
		"archived": false,
		"type": "file",
		"file": {
		  "caption": [],
		  "type": "file",
		  "file": {
			"url": "https://s3.us-west-2.amazonaws.com/secure.notion-static.com/13bc3e79-6a69-4798-93bd-60d1f84e02f5/IMG20220316205019.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Content-Sha256=UNSIGNED-PAYLOAD&X-Amz-Credential=AKIAT73L2G45EIPT3X45%2F20220802%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20220802T231516Z&X-Amz-Expires=3600&X-Amz-Signature=8eeb8081d0283958849acad0a1768e84a289b0208a84deeb561e10739e33b24f&X-Amz-SignedHeaders=host&x-id=GetObject",
			"expiry_time": "2022-08-03T00:15:16.614Z"
		  }
		}
	  }
	  `
	var expected gotion.Block
	var res gotion.Block
	err := json.Unmarshal([]byte(source), &res)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}
func TestUnmarshalInternalPDFBlock(t *testing.T) {
	source := `
	{
		"object": "block",
		"id": "da8f8108-c1f0-4317-942b-afe65a5c63eb",
		"parent": {
		  "type": "page_id",
		  "page_id": "6d8c8ae0-b4b9-410e-af67-f85cb151398f"
		},
		"created_time": "2022-08-02T22:34:00.000Z",
		"last_edited_time": "2022-08-02T22:36:00.000Z",
		"created_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"last_edited_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"has_children": false,
		"archived": false,
		"type": "pdf",
		"pdf": {
		  "caption": [],
		  "type": "file",
		  "file": {
			"url": "https://s3.us-west-2.amazonaws.com/secure.notion-static.com/456c2366-e977-464d-9522-baa22593ba74/cidr2021_paper17.pdf?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Content-Sha256=UNSIGNED-PAYLOAD&X-Amz-Credential=AKIAT73L2G45EIPT3X45%2F20220802%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20220802T231547Z&X-Amz-Expires=3600&X-Amz-Signature=6de2dcb6464d0d60343dba8359d427c79946a19ad27669304939398cd3b547f2&X-Amz-SignedHeaders=host&x-id=GetObject",
			"expiry_time": "2022-08-03T00:15:47.389Z"
		  }
		}
	  }
	  `
	var expected gotion.Block
	var res gotion.Block
	err := json.Unmarshal([]byte(source), &res)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}
func TestUnmarshalBookmarkBlock(t *testing.T) {
	source := `
	{
		"object": "block",
		"id": "fcd32ff3-9e7d-4164-918f-6b002f820457",
		"parent": {
		  "type": "page_id",
		  "page_id": "6d8c8ae0-b4b9-410e-af67-f85cb151398f"
		},
		"created_time": "2022-08-02T22:36:00.000Z",
		"last_edited_time": "2022-08-02T22:36:00.000Z",
		"created_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"last_edited_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"has_children": false,
		"archived": false,
		"type": "bookmark",
		"bookmark": {
		  "caption": [],
		  "url": "https://developers.notion.com/reference/property-value-object#multi-select-property-values"
		}
	  }
	  `
	var expected gotion.Block
	var res gotion.Block
	err := json.Unmarshal([]byte(source), &res)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}

func TestUnmarshalParagraphEquation(t *testing.T) {
	source := `
	{
		"object": "block",
		"id": "abe05812-769e-414c-a3a1-c5315c402176",
		"parent": {
		  "type": "page_id",
		  "page_id": "6d8c8ae0-b4b9-410e-af67-f85cb151398f"
		},
		"created_time": "2022-08-02T22:36:00.000Z",
		"last_edited_time": "2022-08-02T22:41:00.000Z",
		"created_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"last_edited_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"has_children": false,
		"archived": false,
		"type": "paragraph",
		"paragraph": {
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
			}
		  ],
		  "color": "default"
		}
	  }
	  `
	var expected gotion.Block
	var res gotion.Block
	err := json.Unmarshal([]byte(source), &res)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}
func TestUnmarshalFairEquation(t *testing.T) {
	source := `
	{
		"object": "block",
		"id": "ea1d226b-4d40-466f-a026-420fbe2cd603",
		"parent": {
		  "type": "page_id",
		  "page_id": "6d8c8ae0-b4b9-410e-af67-f85cb151398f"
		},
		"created_time": "2022-08-02T22:37:00.000Z",
		"last_edited_time": "2022-08-02T22:39:00.000Z",
		"created_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"last_edited_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"has_children": false,
		"archived": false,
		"type": "equation",
		"equation": {
		  "expression": "|x|=\\begin{cases}\n x, x > 0 \\\\\n-x, x < 0 \n\\end{cases}"
		}
	  }
	  `
	var expected gotion.Block
	var res gotion.Block
	err := json.Unmarshal([]byte(source), &res)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}

func TestUnmarshalTableOfContentsBlock(t *testing.T) {
	source := `
	{
		"object": "block",
		"id": "afbf5f4e-737f-44a6-a7a8-f8ecf2433940",
		"parent": {
		  "type": "page_id",
		  "page_id": "6d8c8ae0-b4b9-410e-af67-f85cb151398f"
		},
		"created_time": "2022-08-02T22:39:00.000Z",
		"last_edited_time": "2022-08-02T22:39:00.000Z",
		"created_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"last_edited_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"has_children": false,
		"archived": false,
		"type": "table_of_contents",
		"table_of_contents": {
		  "color": "gray"
		}
	  }
	  `
	var expected gotion.Block
	var res gotion.Block
	err := json.Unmarshal([]byte(source), &res)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}
func TestUnmarshalBreadcrumbBlock(t *testing.T) {
	source := `
	{
		"object": "block",
		"id": "763f1c8e-1063-43d9-beba-f88f52794d28",
		"parent": {
		  "type": "page_id",
		  "page_id": "6d8c8ae0-b4b9-410e-af67-f85cb151398f"
		},
		"created_time": "2022-08-02T22:39:00.000Z",
		"last_edited_time": "2022-08-02T22:39:00.000Z",
		"created_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"last_edited_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"has_children": false,
		"archived": false,
		"type": "breadcrumb",
		"breadcrumb": {}
	  }
	  `
	var expected gotion.Block
	var res gotion.Block
	err := json.Unmarshal([]byte(source), &res)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}

func TestUnmarshalParagraphLink(t *testing.T) {
	source := `
	{
		"object": "block",
		"id": "2f8f9ea3-f0ec-4bd7-a875-330f3729a747",
		"parent": {
		  "type": "page_id",
		  "page_id": "6d8c8ae0-b4b9-410e-af67-f85cb151398f"
		},
		"created_time": "2022-08-02T22:43:00.000Z",
		"last_edited_time": "2022-08-02T22:43:00.000Z",
		"created_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"last_edited_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"has_children": false,
		"archived": false,
		"type": "paragraph",
		"paragraph": {
		  "rich_text": [
			{
			  "type": "text",
			  "text": {
				"content": "https://developers.notion.com/reference/property-value-object#multi-select-property-values",
				"link": {
				  "url": "https://developers.notion.com/reference/property-value-object#multi-select-property-values"
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
			  "plain_text": "https://developers.notion.com/reference/property-value-object#multi-select-property-values",
			  "href": "https://developers.notion.com/reference/property-value-object#multi-select-property-values"
			}
		  ],
		  "color": "default"
		}
	  }
	  `
	var expected gotion.Block
	var res gotion.Block
	err := json.Unmarshal([]byte(source), &res)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}

//TODO do different types of template exists? Or just to-do? Need test
func TestUnmarshalTemplateBlock(t *testing.T) {
	source := `
	{
		"object": "block",
		"id": "89420287-76ec-4b0b-882f-9d38992b67cd",
		"parent": {
		  "type": "page_id",
		  "page_id": "6d8c8ae0-b4b9-410e-af67-f85cb151398f"
		},
		"created_time": "2022-08-02T22:43:00.000Z",
		"last_edited_time": "2022-08-02T22:44:00.000Z",
		"created_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"last_edited_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"has_children": true,
		"archived": false,
		"type": "template",
		"template": {
		  "rich_text": [
			{
			  "type": "text",
			  "text": {
				"content": "Add a new to-do - test template",
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
			  "plain_text": "Add a new to-do - test template",
			  "href": null
			}
		  ]
		}
	  }
	  `
	var expected gotion.Block
	var res gotion.Block
	err := json.Unmarshal([]byte(source), &res)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}

func TestUnmarshalTableBlock(t *testing.T) {
	source := `
	{
		"object": "block",
		"id": "e1f1e00a-f567-451a-a3bb-e02346185352",
		"parent": {
		  "type": "page_id",
		  "page_id": "6d8c8ae0-b4b9-410e-af67-f85cb151398f"
		},
		"created_time": "2022-08-02T22:45:00.000Z",
		"last_edited_time": "2022-08-02T23:07:00.000Z",
		"created_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"last_edited_by": {
		  "object": "user",
		  "id": "c7f2ae70-6b98-438f-8564-c59a71d7b3a4"
		},
		"has_children": true,
		"archived": false,
		"type": "table",
		"table": {
		  "table_width": 2,
		  "has_column_header": false,
		  "has_row_header": false
		}
	  }
	  `
	var expected gotion.Block
	var res gotion.Block
	err := json.Unmarshal([]byte(source), &res)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}
