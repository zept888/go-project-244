package diff

type NodeType string

const (
	Added     NodeType = "added"
	Removed   NodeType = "removed"
	Unchanged NodeType = "unchanged"
	Updated   NodeType = "updated"
	Nested    NodeType = "nested"
)

type Node struct {
	Key      string
	Type     NodeType
	OldValue any
	NewValue any
	Children []Node
}
