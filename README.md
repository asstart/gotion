# gotion

## Not implemented at all

- https://developers.notion.com/reference/retrieve-a-page-property


## Strange cases

- Can't create page under another page with parent block like this:
```
"parent": {
        "type": "page_id",
        "page_id": "02e0fefa-38e6-47fb-99b1-ac066ac1cd86"
    }
```

## Need to test

- Create page with properties: rollup, relation

## Need to add

- Page can be settled under parent with type=block. Now only database, page and workspace possible
- Block types which support children are "paragraph", "bulleted_list_item", "numbered_list_item", "toggle", "to_do", "quote", "callout", "synced_block", "template", "column", "child_page", "child_database", and "table". - Add validation

## Think about

- There'r 3 different parents for: database, page, block. Think about to combine them into a single one. https://developers.notion.com/reference/parent-object