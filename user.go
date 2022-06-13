package gotion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

/**
https://developers.notion.com/reference/user#where-user-objects-appear-in-the-api

User objects will always contain object and id keys, as described below.
The remaining properties may appear if the user is being rendered in a rich text or page property context,
and the bot has the correct capabilities to access those properties.
**/
type User struct {
	ID        string   `json:"id"`
	Object    string   `json:"object,omitempty"`
	Type      UserType `json:"type,omitempty"`
	Name      string   `json:"name,omitempty"`
	AvatarUrl string   `json:"avatar_url,omitempty"`
	Person    *Person  `json:"person,omitempty"`
	Bot       *Bot     `json:"bot,omitempty"`
}

type Person struct {
	Person string `json:"person,omitempty"`
	Email  string `json:"email"`
}

type Bot struct {
	Bot   string   `json:"bot,omitempty"`
	Owner BotOwner `json:"owner"`
}

type BotOwner struct {
	Owner     string `json:"owner,omitempty"`
	Type      string `json:"type"`
	Workspace bool   `json:"workspace"`
	OwnerUser User   `json:"user"`
}

type UserType int

const (
	PersonType UserType = iota
	BotType
)

var UserTypeToString = map[UserType]string{
	PersonType: "person",
	BotType:    "bot",
}

var StringToUserType = map[string]UserType{
	"person": PersonType,
	"bot":    BotType,
}

func (p *UserType) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	res, ok := StringToUserType[v]
	if !ok {
		return fmt.Errorf("%v isn't enum value", res)
	}
	p = &res
	return nil
}

func (p UserType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(UserTypeToString[p])
	b.WriteString(`"`)
	return b.Bytes(), nil
}
