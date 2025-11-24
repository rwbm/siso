package siso

import (
	"fmt"
	"sort"
	"strings"
)

type IsoMessage struct {
	Header Field
	fields map[int]Field
}

func (i *IsoMessage) WithField(f Field) *IsoMessage {
	// bitmaps aren't set manually
	if f.ID == FieldPrimaryBitmap || f.ID == FieldSecondaryBitmap {
		return i
	}

	i.ensureMap()
	i.fields[f.ID] = f
	return i
}

func (i *IsoMessage) Remove(id int) {
	if i.fields != nil {
		delete(i.fields, id)
	}
}

func (i *IsoMessage) Field(id int) (f *Field, ok bool) {
	if i.fields != nil {
		value, ok := i.fields[id]
		return &value, ok
	}
	return nil, false
}

func (i *IsoMessage) Contains(id int) bool {
	if i.fields != nil {
		_, ok := i.fields[id]
		return ok
	}
	return false
}

func (i *IsoMessage) MTI() (f *Field, ok bool) {
	if i.fields != nil {
		value, ok := i.fields[FieldMessageTypeIndicator]
		return &value, ok
	}
	return nil, false
}

func (i *IsoMessage) PrimaryBitmap() (f *Field, ok bool) {
	if i.fields != nil {
		value, ok := i.fields[FieldPrimaryBitmap]
		return &value, ok
	}
	return nil, false
}

func (i *IsoMessage) SecondaryBitmap() (f *Field, ok bool) {
	if i.fields != nil {
		value, ok := i.fields[FieldSecondaryBitmap]
		return &value, ok
	}
	return nil, false
}

// StringXml returns a string with an XML representation of the message.
func (i *IsoMessage) StringXml() string {
	// extract ids and sort it
	keys := make([]int, 0, len(i.fields))
	for k := range i.fields {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	sb := strings.Builder{}
	if mti, exists := i.Field(FieldMessageTypeIndicator); exists {
		sb.WriteString(fmt.Sprintf("<iso mti=\"%s\">\n", mti.value))
	} else {
		sb.WriteString("<iso>\n")
	}

	for _, k := range keys {
		// skip special fields
		if k == FieldMessageTypeIndicator || k == FieldPrimaryBitmap || k == FieldSecondaryBitmap {
			continue
		}
		sb.WriteString(fmt.Sprintf("   <field id=\"%d\">%s</field>\n", i.fields[k].ID, i.fields[k].value))
	}

	sb.WriteString("<iso>")
	return sb.String()
}

func (i *IsoMessage) ensureMap() {
	if i.fields == nil {
		i.fields = make(map[int]Field)
	}
}

func (i *IsoMessage) refreshBitmaps() {
	// TODO:
}
