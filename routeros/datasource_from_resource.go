package routeros

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dsEntriesKey is the attribute every derived data source returns its results
// under. A fixed name is used rather than a per-menu plural: the set is
// generated, and "entries" is predictable across all of them in a way that
// invented plurals are not.
const dsEntriesKey = "entries"

func isMetaAttribute(name string) bool {
	return strings.HasPrefix(name, "___") && strings.HasSuffix(name, "___")
}

// computedSchema mirrors an attribute as read-only. Validators, defaults and
// diff suppression describe how a value is written and mean nothing on a data
// source, so only the shape and the description are carried over.
func computedSchema(in *schema.Schema) *schema.Schema {
	out := &schema.Schema{
		Type:        in.Type,
		Computed:    true,
		Description: in.Description,
		Sensitive:   in.Sensitive,
	}

	switch elem := in.Elem.(type) {
	case *schema.Schema:
		// The element of a list or set of primitives carries the type and
		// nothing else; the SDK rejects a schema with anything more set.
		out.Elem = &schema.Schema{Type: elem.Type}
	case *schema.Resource:
		nested := map[string]*schema.Schema{}
		for name, attr := range elem.Schema {
			if isMetaAttribute(name) {
				continue
			}
			nested[name] = computedSchema(attr)
		}
		out.Elem = &schema.Resource{Schema: nested}
	}

	return out
}

// datasourceFromResource builds a filterable, read-only data source over the
// same RouterOS menu as the resource it is given.
//
// Only menus that hold a list of entries are passed here; a settings menu has
// nothing to filter and returns a single object. Resources declaring
// MetaTransformSet are also excluded, because the data source serializer
// applies drift renames but not a transform set, so their renamed attributes
// would be dropped on read.
func datasourceFromResource(res *schema.Resource) *schema.Resource {
	pathAttr, ok := res.Schema[MetaResourcePath]
	if !ok {
		return nil
	}
	path, ok := pathAttr.Default.(string)
	if !ok {
		return nil
	}

	entry := map[string]*schema.Schema{
		// RouterOS returns .id for every entry in a list menu.
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The RouterOS internal id of the entry.",
		},
	}
	for name, attr := range res.Schema {
		if isMetaAttribute(name) || name == "id" {
			continue
		}
		entry[name] = computedSchema(attr)
	}

	s := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath(path),
		MetaId:           PropId(Id),

		KeyFilter: PropFilterRw,
		dsEntriesKey: {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "Entries read from `" + path + "`, narrowed by `filter`.",
			Elem:        &schema.Resource{Schema: entry},
		},
	}

	// The serializer reads this from the top level of the data source schema,
	// so a resource that hides fields must hide the same ones here.
	if sf, ok := res.Schema[MetaSkipFields]; ok {
		s[MetaSkipFields] = &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Default:     sf.Default,
			Description: sf.Description,
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				return true
			},
		}
	}

	return &schema.Resource{
		Schema: s,
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			items, err := ReadItemsFiltered(
				buildReadFilter(d.Get(KeyFilter).(map[string]interface{})), path, m.(Client))
			if err != nil {
				return diag.FromErr(err)
			}

			return MikrotikResourceDataToTerraformDatasource(items, dsEntriesKey, s, d)
		},
	}
}

// registerDerivedDatasources adds a data source for every list menu in
// derivedDatasourceMenus that does not already have a hand-written one. Doing
// this from the resource map rather than from generated files keeps the two in
// step: a new resource on a listed menu gets its data source automatically.
func registerDerivedDatasources(p *schema.Provider) {
	for name, res := range p.ResourcesMap {
		if _, taken := p.DataSourcesMap[name]; taken {
			continue
		}
		pathAttr, ok := res.Schema[MetaResourcePath]
		if !ok {
			continue
		}
		path, ok := pathAttr.Default.(string)
		if !ok {
			continue
		}
		if _, eligible := derivedDatasourceMenus[path]; !eligible {
			continue
		}
		if ds := datasourceFromResource(res); ds != nil {
			p.DataSourcesMap[name] = ds
		}
	}
}
