package gotion

type RichText struct {
	Type        string      `json:"type"`
	PlainText   string      `json:"plain_text"`
	Href        string      `json:"href"`
	Annotations Annotations `json:"annotations"`
	Text        Text        `json:"text,omitempty"`
	Mention     Mention     `json:"mention,omitempty"`
	Equation    Equation    `json:"equation,omitempty"`
}

type Annotations struct {
	Bold          bool   `json:"bold,omitempty"`
	Italic        bool   `json:"italic,omitempty"`
	Strikethrough bool   `json:"strikethrough,omitempty"`
	Underline     bool   `json:"underline,omitempty"`
	Code          bool   `json:"code,omitempty"`
	Color         string `json:"color,omitempty"`
}

type Text struct {
	Content string `json:"content"`
	Link    Link   `json:"link,omitempty"`
}

type Mention struct {
	MentionType string `json:"type"`

	User     string   `json:"user,omitempty"`
	Page     string   `json:"page,omitempty"`
	Database string   `json:"database,omitempty"`
	Date     string   `json:"date,omitempty"`
	Template Template `json:"template_mention,omitempty"`
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
