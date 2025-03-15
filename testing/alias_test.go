package testing

import (
	"fmt"
	"github.com/dprotaso/go-yit"
	"log"
	"testing"

	"gopkg.in/yaml.v3"
)

var data = `
first: second
dummy: &var world
hello: *var
`

func TestAlias(t *testing.T) {
	var Doc yaml.Node

	err := yaml.Unmarshal([]byte(data), &Doc)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// FromNode -> builds an iterator for all nodes at that level (just the DocumentNode in this case)
	// RecurseNodes() iterates on that DocumentNode, recursively building an iterator on all children
	it := yit.FromNodes(Doc.Content[0].Content...) //.RecurseNodes()

	fmt.Println("--")
	for node, ok := it(); ok; node, ok = it() {
		fmt.Println(node.Value)
	}
	fmt.Println("--")
}
