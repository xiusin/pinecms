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

type KV struct {
	Label string      `json:"label"`
	Value interface{} `json:"value"`
}
