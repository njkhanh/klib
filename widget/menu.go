package widget

type Menu struct {
	Text       string
	Icon       string
	Url        string
	Route      string
	Active     string
	Activated  bool
	IsHidden   bool
	Permission string
	Target     string
	Params     []string
	Children   []Menu
}
