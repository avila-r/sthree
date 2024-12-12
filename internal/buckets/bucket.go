package buckets

import "github.com/avila-r/sthree/internal/objects"

func (m *Module) Bucket(name string) *objects.Module {
	return &objects.Module{
		Bucket: name,
		Sdk:    m.Sdk,
	}
}

func (m *Module) In(name string) *objects.Module {
	return m.Bucket(name)
}
