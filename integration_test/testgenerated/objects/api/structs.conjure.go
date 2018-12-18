// This file was generated by Conjure and should not be manually edited.

package api

import (
	"encoding/json"

	"github.com/palantir/conjure-go/conjure/types/conjuretype"
)

type Basic struct {
	Data string `json:"data" yaml:"data,omitempty"`
}

type Collections struct {
	MapVar   map[string][]int   `json:"mapVar" yaml:"mapVar,omitempty"`
	ListVar  []string           `json:"listVar" yaml:"listVar,omitempty"`
	MultiDim [][]map[string]int `json:"multiDim" yaml:"multiDim,omitempty"`
}

func (o Collections) MarshalJSON() ([]byte, error) {
	if o.MapVar == nil {
		o.MapVar = make(map[string][]int, 0)
	}
	if o.ListVar == nil {
		o.ListVar = make([]string, 0)
	}
	if o.MultiDim == nil {
		o.MultiDim = make([][]map[string]int, 0)
	}
	type CollectionsAlias Collections
	return json.Marshal(CollectionsAlias(o))
}

func (o *Collections) UnmarshalJSON(data []byte) error {
	type CollectionsAlias Collections
	var rawCollections CollectionsAlias
	if err := json.Unmarshal(data, &rawCollections); err != nil {
		return err
	}
	if rawCollections.MapVar == nil {
		rawCollections.MapVar = make(map[string][]int, 0)
	}
	if rawCollections.ListVar == nil {
		rawCollections.ListVar = make([]string, 0)
	}
	if rawCollections.MultiDim == nil {
		rawCollections.MultiDim = make([][]map[string]int, 0)
	}
	*o = Collections(rawCollections)
	return nil
}

func (o Collections) MarshalYAML() (interface{}, error) {
	if o.MapVar == nil {
		o.MapVar = make(map[string][]int, 0)
	}
	if o.ListVar == nil {
		o.ListVar = make([]string, 0)
	}
	if o.MultiDim == nil {
		o.MultiDim = make([][]map[string]int, 0)
	}
	type CollectionsAlias Collections
	return CollectionsAlias(o), nil
}

func (o *Collections) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type CollectionsAlias Collections
	var rawCollections CollectionsAlias
	if err := unmarshal(&rawCollections); err != nil {
		return err
	}
	if rawCollections.MapVar == nil {
		rawCollections.MapVar = make(map[string][]int, 0)
	}
	if rawCollections.ListVar == nil {
		rawCollections.ListVar = make([]string, 0)
	}
	if rawCollections.MultiDim == nil {
		rawCollections.MultiDim = make([][]map[string]int, 0)
	}
	*o = Collections(rawCollections)
	return nil
}

type Compound struct {
	Obj Collections `json:"obj" yaml:"obj,omitempty"`
}

type ExampleUuid struct {
	Uid conjuretype.UUID `json:"uid" yaml:"uid,omitempty"`
}
