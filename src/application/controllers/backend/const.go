package backend

type MenuV2 struct {
	Path            string   `json:"path"`
	Label           string   `json:"label"`
	NodePath        string   `json:"nodePath"`
	Exact           bool     `json:"exact"`
	Icon            string   `json:"icon"`
	PathToComponent string   `json:"pathToComponent"`
	SideVisible     bool     `json:"sideVisible"`
	Children        []MenuV2 `json:"children"`
}

type ThemeConfig struct {
	Name        string                 `json:"name"`
	Author      string                 `json:"author"`
	Description string                 `json:"description"`
	Extra       map[string]interface{} `json:"extra"`
	IsDefault   bool                   `json:"is_default"`
	Dir         string                 `json:"dir"`
}

type FieldShowInPageList struct {
	Forms []KV `json:"forms"`
	List  []KV `json:"list"`
}

type KV struct {
	Label   string      `json:"label"`
	Value   interface{} `json:"value"`
	Name    string      `json:"-"`
	Checked bool        `json:"-"`
}

type TabsSchema struct {
	Title string         `json:"title"`
	Hash  string         `json:"hash"`
	Body  FormController `json:"body"`
}

type FormController struct {
	Title    string        `json:"title"`
	Api      string        `json:"api"`
	Type     string        `json:"type"`
	Mode     string        `json:"mode"`
	Controls []FormControl `json:"controls"`
}

type FormControl struct {
	Type             string        `json:"type"`
	Name             string        `json:"name"`
	Label            string        `json:"label"`
	Value            interface{}   `json:"value"`
	Options          []KV          `json:"options"`
	Validations      string        `json:"validations"`
	Required         bool          `json:"required"`
	Description      string        `json:"description"`
	Placeholder      string        `json:"placeholder"`
	ValidationErrors string        `json:"validationErrors"`
	Multiple         bool          `json:"multiple"`
	Precision        int           `json:"precision"`
	Inline           bool          `json:"inline"`
	Buttons          []interface{} `json:"buttons"`
	Limits           []string      `json:"limits"`
	LimitsLogic      string        `json:"limitsLogic"`
}
