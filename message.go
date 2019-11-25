package samhook

// Warning is a predefined color for a warning (yellow)
const Warning string = "#FFBB00"

// Danger is a predefined color for a dangerous condition (red)
const Danger string = "#FF0000"

// Good is a predefined color for a normal information (green)
const Good string = "#00FF00"

// Message message主體
type Message struct {
	Parse       string       `json:"parse,omitempty"`
	Username    string       `json:"username,omitempty"`
	IconURL     string       `json:"icon_url,omitempty"`
	IconEmoji   string       `json:"icon_emoji,omitempty"`
	Channel     string       `json:"channel,omitempty"`
	Text        string       `json:"text,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

// Attachment attachment主體
type Attachment struct {
	Fallback   string  `json:"fallback,omitempty"`
	Color      string  `json:"color,omitempty"`
	Pretext    string  `json:"pretext,omitempty"`
	AuthorName string  `json:"author_name,omitempty"`
	AuthorLink string  `json:"author_link,omitempty"`
	AuthorIcon string  `json:"author_icon,omitempty"`
	Title      string  `json:"title,omitempty"`
	TitleLink  string  `json:"title_link,omitempty"`
	Text       string  `json:"text,omitempty"`
	ImageURL   string  `json:"image_url,omitempty"`
	Fields     []Field `json:"fields,omitempty"`
	Footer     string  `json:"footer,omitempty"`
	FooterIcon string  `json:"footer_icon,omitempty"`
	ThumbURL   string  `json:"thumb_url,omitempty"`
}

// Field field主體
type Field struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
	Short bool   `json:"short,omitempty"`
}
