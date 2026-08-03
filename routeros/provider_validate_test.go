package routeros

import "testing"

// TestProviderInternalValidate runs the SDK's own schema checks over every
// registered resource and data source. It needs no device, and it catches the
// class of mistake that otherwise surfaces only when a user writes a
// configuration against the offending attribute.
func TestProviderInternalValidate(t *testing.T) {
	if err := NewProvider().InternalValidate(); err != nil {
		t.Fatalf("provider schema is invalid: %v", err)
	}
}

// TestDerivedDatasourcesRegistered guards the derivation in
// registerDerivedDatasources: every menu listed in derivedDatasourceMenus that
// has a resource should end up with a data source of the same name.
func TestDerivedDatasourcesRegistered(t *testing.T) {
	p := NewProvider()

	for name, res := range p.ResourcesMap {
		path, ok := res.Schema[MetaResourcePath].Default.(string)
		if !ok {
			continue
		}
		if _, eligible := derivedDatasourceMenus[path]; !eligible {
			continue
		}
		ds, ok := p.DataSourcesMap[name]
		if !ok {
			t.Errorf("%s covers %s but has no data source", name, path)
			continue
		}
		if _, ok := ds.Schema[dsEntriesKey]; !ok {
			t.Errorf("data source %s has no %q attribute", name, dsEntriesKey)
		}
	}
}
