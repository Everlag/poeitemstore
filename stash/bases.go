package stash

//go:generate go-bindata -pkg $GOPACKAGE -nometadata -o assets.go assets/

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// Base list acquired from
// https://poeapp.com/api/item-types.json
//
// The raw response was transformed with
// x as the parsed baseInfo.json
// transformed = x.map(v=> {
// 	v = Object.assign({}, v);
// 	v.root = v.key;
// 	delete v.key;
// 	v.flavors = v.value;
// 	delete v.value;
// 	return v})
// .map(v => {
// 	v.flavors = v.flavors.map(f => {
// 		f = Object.assign({}, f);
// 		f.flavor = f.key;
// 		delete f.key;
// 		f.bases = f.value;
// 		delete f.value;
// 		return f
// 	});
// 	return v;
// })
// JSON.stringify(transformed)

type flavor struct {
	Flavor string   `json:"flavor"`
	Bases  []string `json:"bases"`
}

type root struct {
	Root    string   `json:"root"`
	Flavors []flavor `json:"flavors"`
}

type baseLookup []root

// toInverter converts a baseLookup to a baseInverter
func (bl baseLookup) toInverter() baseInverter {

	basesToFlavors := make(map[string]string)
	flavorsToRoots := make(map[string]string)

	for _, root := range bl {
		for _, flavor := range root.Flavors {
			flavorsToRoots[flavor.Flavor] = root.Root

			for _, base := range flavor.Bases {
				basesToFlavors[base] = flavor.Flavor
			}
		}
	}

	return baseInverter{
		flavorsToRoots: flavorsToRoots,
		basesToFlavors: basesToFlavors,
	}

}

func getBaseLookup() (*baseLookup, error) {
	raw, err := Asset("assets/baseInfo.json")
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch baseInfo.json")
	}

	var bl baseLookup
	err = json.Unmarshal(raw, &bl)
	return &bl, err
}

type baseInverter struct {
	basesToFlavors map[string]string
	flavorsToRoots map[string]string
}

// typeLineToRootAndFlavor takes a given TypeLine on an item
// and converts it to a root and flavor.
//
// This even handles the super mangled typeLines that GGG sometimes includes...
func (inv baseInverter) typeLineToRootAndFlavor(typeLine string) (flavor,
	root string, ok bool) {

	// First we try a direct match
	flavor, ok = inv.basesToFlavors[typeLine]
	if ok {
		// Resolve the root and return
		root = inv.flavorsToRoots[flavor]
		return
	}

	// If direct match failed, we need an exhaustive substring search
	// This hurts performance but usually is fairly rare :|
	for propBase, propFlavor := range inv.basesToFlavors {
		if strings.Contains(typeLine, propBase) {
			flavor = propFlavor
			break
		}
	}
	if len(flavor) == 0 {
		ok = false
		return
	}

	// Resolve a correct flavor to base easily
	root = inv.flavorsToRoots[flavor]
	ok = true
	return

}

// Singleton inverter the package uses
var inverter baseInverter

// MatchTypeline returns a flavor and root for a discovered base
// or ok is false.
func MatchTypeline(typeline string) (flavor, root string, ok bool) {
	return inverter.typeLineToRootAndFlavor(typeline)
}

// Prep our singleton inverter
func init() {
	bl, err := getBaseLookup()
	if err != nil {
		panic(fmt.Sprintf("failed to deserialize baseLookup on stash init, err=%s", err))
	}

	inverter = bl.toInverter()
}
