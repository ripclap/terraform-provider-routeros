//go:build ignore
// +build ignore

// Prints, as JSON, every resource this provider registers: its RouterOS menu
// path and the fields its schema declares. Intended to be diffed against what a
// device reports, to find menus that are not modelled and fields that are.
package main

import (
	"encoding/json"
	"os"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ripclap/terraform-provider-routeros/routeros"
)

type resourceInfo struct {
	Path     string   `json:"path"`
	Fields   []string `json:"fields"`
	ReadOnly []string `json:"read_only"`
}

// Walks nested blocks too. A block like `input { accept_nlri = ... }` is the
// provider's way of modelling the device's flat `input.accept-nlri`, so both the
// dotted and the underscored spelling are recorded; otherwise every nested
// attribute looks like a field the provider is missing.
func describe(s map[string]*schema.Schema, prefix string, fields, ro *[]string) {
	for name, sch := range s {
		if strings.HasPrefix(name, "___") {
			continue
		}
		full := name
		if prefix != "" {
			full = prefix + "_" + name
		}
		*fields = append(*fields, full)
		if sch.Computed && !sch.Optional {
			*ro = append(*ro, full)
		}
		if nested, ok := sch.Elem.(*schema.Resource); ok && nested != nil {
			describe(nested.Schema, full, fields, ro)
		}
	}
}

func describeTop(s map[string]*schema.Schema) ([]string, []string) {
	var fields, ro []string
	describe(s, "", &fields, &ro)
	sort.Strings(fields)
	sort.Strings(ro)
	return fields, ro
}

func main() {
	p := routeros.Provider()
	out := map[string]resourceInfo{}

	for name, res := range p.ResourcesMap {
		path := ""
		if f, ok := res.Schema[routeros.MetaResourcePath]; ok && f.Default != nil {
			path, _ = f.Default.(string)
		}
		fields, ro := describeTop(res.Schema)
		out[name] = resourceInfo{Path: path, Fields: fields, ReadOnly: ro}
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(out); err != nil {
		panic(err)
	}
}
