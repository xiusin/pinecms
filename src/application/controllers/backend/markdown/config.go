package markdown

type Config struct {
	Authors     Authors    `json:"authors"`
	Homepage    string     `json:"homepage"`
	Description string     `json:"description"`
	Website     Website    `json:"website"`
	Doc         Doc        `json:"doc"`
	Comment     Comment    `json:"comment"`
	EditConfig  EditConfig `json:"edit_config"`
}
type Authors struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
type Website struct {
	Title       string `json:"title"`
	Keywords    string `json:"keywords"`
	Description string `json:"description"`
}
type Doc struct {
	Default string        `json:"Default"`
	Index   string        `json:"index"`
	Version string        `json:"version"`
	Ignore  []interface{} `json:"ignore"`
}
type Comment struct {
	Enable         bool     `json:"enable"`
	ClientID       string   `json:"client_id"`
	ClientSecret   string   `json:"client_secret"`
	Repo           string   `json:"repo"`
	Owner          string   `json:"owner"`
	Admin          []string `json:"admin"`
	Language       string   `json:"language"`
	PerPage        int      `json:"perPage"`
	PagerDirection string   `json:"pagerDirection"`
}
type Values struct {
	Installation string `json:"installation"`
}
type EditConfig struct {
	Label  string `json:"label"`
	Values Values `json:"values"`
}
