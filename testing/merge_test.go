package testing

import (
	"fmt"
	"github.com/dprotaso/go-yit"
	"log"
	"testing"

	"gopkg.in/yaml.v3"
)

var dataMerge = `
first: second
dummy: &var
  a: b
  c: d
hello: 
  <<: *var
  e: f
  g: h
`

func TestMerge(t *testing.T) {
	var Doc yaml.Node

	err := yaml.Unmarshal([]byte(dataMerge), &Doc)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// FromNode -> builds an iterator for all nodes at that level (just the DocumentNode in this case)
	// RecurseNodes() iterates on that DocumentNode, recursively building an iterator on all children
	//it := yit.FromNode(&Doc).RecurseNodes()
	it := yit.FromNodes(Doc.Content[0].Content[5].Content...)

	for node, ok := it(); ok; node, ok = it() {
		fmt.Println(node.Value)
	}
}
