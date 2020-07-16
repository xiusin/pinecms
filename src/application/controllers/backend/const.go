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
