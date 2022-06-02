package gotion

type RichText struct {
	Type        string       `json:"type,omitempty"`
	PlainText   string       `json:"plain_text,omitempty"`
	Href        string       `json:"href,omitempty"`
	Annotations *Annotations `json:"annotations,omitempty"`
	//One of followings should be define
	Text     *Text     `json:"text,omitempty"`
	Mention  *Mention  `json:"mention,omitempty"`
	Equation *Equation `json:"equation,omitempty"`
}

type Annotations struct {
	Bold          bool   `json:"bold"`
	Italic        bool   `json:"italic"`
	Strikethrough bool   `json:"strikethrough"`
	Underline     bool   `json:"underline"`
	Code          bool   `json:"code"`
	Color         string `json:"color"`
}

type Text struct {
	Content string `json:"content"`
	Link    *Link  `json:"link,omitempty"`
}

type Mention struct {
	MentionType string `json:"type,omitempty"`

	User     *User         `json:"user,omitempty"`
	Page     string        `json:"page,omitempty"`
	Database string        `json:"database,omitempty"`
	Date     *DateProperty `json:"date,omitempty"`
	Template *Template     `json:"template_mention,omitempty"`
}

type Equation struct {
	Expression string `json:"expression"`
}

type FileDescriptor struct {
	Type         string       `json:"type"`
	ExternalFile ExternalFile `json:"external,omitempty"`
	NotionFile   NotionFile   `json:"file,omitempty"`
}

type ExternalFile struct {
	URL string `json:"url"`
}

type NotionFile struct {
	URL        string `json:"url"`
	ExpireTime string `json:"expire_time"`
}

type Emoji struct {
	Type  string `json:"type"`
	Emoji string `json:"emoji"`
}
