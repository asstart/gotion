package gotion

/**
https://developers.notion.com/reference/user#where-user-objects-appear-in-the-api

User objects will always contain object and id keys, as described below.
The remaining properties may appear if the user is being rendered in a rich text or page property context,
and the bot has the correct capabilities to access those properties.
**/
type User struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	Type      string `json:"type,omitempty"`
	Name      string `json:"name,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	Person    Person `json:"person,omitempty"`
	Bot       *Bot   `json:"bot,omitempty"`
}

type Person struct {
	Email string `json:"email"`
}

type Bot struct {
	Owner BotOwner `json:"owner"`
}

type BotOwner struct {
	Type      string `json:"type"`
	Workspace bool   `json:"workspace"`
	Owner     User   `json:"user"`
}
