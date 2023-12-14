package nginx

// Items is an interface that represents either a Block or a Directive.
type Items interface {
	isItem() bool
	Type() string
}

// Block represents a block in the Nginx configuration.
type Block struct {
	Name  string
	Items []Items
}

func (b Block) isItem() bool {
	return true
}

func (b Block) Type() string {
	return "block"
}

func (b Block) GetServers() (block Block) {
	for _, item := range b.Items {
		block := item.(Block)
		if block.Name == "server" {
			return block
		}
	}
	return
}

func (b Block) GetServerName() string {
	for _, item := range b.Items {
		directive := item.(*Directive)
		if directive.Name == "server_name" {
			return directive.Value
		}
	}
	return ""
}

// Directive represents a directive in the Nginx configuration.
type Directive struct {
	Name  string
	Value string
}

func (d Directive) isItem() bool {
	return true
}

func (d Directive) Type() string {
	return "directive"
}
