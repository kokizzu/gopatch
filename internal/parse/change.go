package parse

import (
	"github.com/uber-go/gopatch/internal/parse/section"
)

// Parses the change at index i.
func (p *parser) parseChange(i int, c *section.Change) (_ *Change, err error) {
	change := Change{Name: c.Name}

	change.Meta, err = p.parseMeta(i, c)
	if err != nil {
		return nil, err
	}

	// TODO: Parse the rest of the change.

	return &change, nil
}
