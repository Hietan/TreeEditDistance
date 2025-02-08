package parser

import (
	"encoding/json"
	"github.com/Hietan/TreeEditDistance/internal/model"
	"os"
)

type jsonNode[T any] struct {
	Value    T              `json:"value"`
	Children []*jsonNode[T] `json:"children,omitempty"`
}

func convertJSONNode[T any](j *jsonNode[T], parent *model.Node[T]) *model.Node[T] {
	node := model.NewNode(j.Value)
	node.SetParent(parent)

	for _, childJSON := range j.Children {
		childNode := convertJSONNode[T](childJSON, node)
		node.AddChild(childNode)
	}

	return node
}

func parseTree[T any](data []byte) (*model.Tree[T], error) {
	var rootJSON jsonNode[T]
	if err := json.Unmarshal(data, &rootJSON); err != nil {
		return nil, err
	}
	rootNode := convertJSONNode(&rootJSON, nil)
	return model.NewTreeFromNode(rootNode), nil
}

func LoadTreeFromFile[T any](filePath string) *model.Tree[T] {
	data, err := os.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	tree, err := parseTree[T](data)

	if err != nil {
		panic(err)
	}

	return tree
}
