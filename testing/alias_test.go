package testing

import (
	"fmt"
	"github.com/dprotaso/go-yit"
	"gopkg.in/yaml.v3"
	"os"
	"testing"
)

func TestAlias(t *testing.T) {
	y, err := os.ReadFile("./input.yaml")
	if err != nil {
		t.Error(err)
	}

	var doc yaml.Node

	if err = yaml.Unmarshal(y, &doc); err != nil {
		t.Fatal(err)
	}

	it := yit.FromNode(&doc).Values().Values().Values() //.RecurseNodes().Values() //.Filter(yit.WithKind(yaml.AliasNode)).MapKeys()

	for node, ok := it(); ok; node, ok = it() {
		print("----------")
		fmt.Println(node.Value)
	}
}
