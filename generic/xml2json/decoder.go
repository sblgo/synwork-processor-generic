package xml2json

import (
	"encoding/xml"
	"fmt"
	"io"
	"unicode"

	"golang.org/x/net/html/charset"
)

const (
	attrPrefix    = "-"
	contentPrefix = "#"
)

// A Decoder reads and decodes XML objects from an input stream.
type Decoder struct {
	r               io.Reader
	err             error
	attributePrefix string
	contentPrefix   string
	excludeAttrs    map[string]bool
	arrayPaths      map[string]bool
}

type element struct {
	parent *element
	n      *Node
	label  string
	path   string
}

func (dec *Decoder) SetAttributePrefix(prefix string) {
	dec.attributePrefix = prefix
}

func (dec *Decoder) SetContentPrefix(prefix string) {
	dec.contentPrefix = prefix
}

func (dec *Decoder) SetArrayPaths(paths ...string) {
	if dec.arrayPaths == nil {
		dec.arrayPaths = make(map[string]bool)
	}
	for _, p := range paths {
		dec.arrayPaths[p] = true
	}
}

func (dec *Decoder) ExcludeAttributes(attrs []string) {
	for _, attr := range attrs {
		dec.excludeAttrs[attr] = true
	}
}

// NewDecoder returns a new decoder that reads from r.
func NewDecoder(r io.Reader) *Decoder {
	d := &Decoder{r: r, contentPrefix: contentPrefix, attributePrefix: attrPrefix, excludeAttrs: map[string]bool{}}
	return d
}

// Decode reads the next JSON-encoded value from its
// input and stores it in the value pointed to by v.
func (dec *Decoder) Decode(root *map[string]interface{}) error {
	xmlDec := xml.NewDecoder(dec.r)

	// That will convert the charset if the provided XML is non-UTF-8
	xmlDec.CharsetReader = charset.NewReaderLabel

	// Create first element from the root node
	elem := &Node{}
	for {
		t, err := xmlDec.Token()
		if err != nil && io.EOF != err {
			return err
		}
		if t == nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			// Build new a new current element and link it to its parent
			elem = &Node{
				Parent: elem,
				Name:   se.Name.Local,
				Path:   fmt.Sprintf("%s/%s", elem.Path, se.Name.Local),
			}

			// Extract attributes as children
			for _, a := range se.Attr {
				if _, ok := dec.excludeAttrs[a.Name.Local]; ok {
					continue
				}
				elem.AddChild(dec.attributePrefix+a.Name.Local, &Node{Data: a.Value})
			}
		case xml.CharData:
			// Extract XML data (if any)
			elem.Data = trimNonGraphic(string(xml.CharData(se)))
		case xml.EndElement:
			// And add it to its parent list
			if elem.Parent != nil {
				elem.Parent.AddChild(elem.Name, elem)
			}

			// Then change the current element to its parent
			elem = elem.Parent
		}
	}
	dec.traverse(elem, root)
	return nil
}

func (dec *Decoder) traverse(elem *Node, consumer *map[string]interface{}) {
	switch len(elem.Children) {

	case 0:
		(*consumer)[elem.Name] = elem.Data
	default:
		for _, nodes := range elem.Children {
			switch len(nodes) {
			case 1:
				child := nodes[0]
				if _, ok := dec.arrayPaths[child.Path]; ok {
					arr := new([]interface{})
					m := make(map[string]interface{})
					dec.traverse(child, &m)
					*arr = append(*arr, m)
					(*consumer)[child.Name] = *arr
				} else {
					dec.traverse(child, consumer)
				}
			default:
				arr := new([]interface{})
				for _, child := range nodes {
					m := make(map[string]interface{})
					dec.traverse(child, &m)
					*arr = append(*arr, m)

				}
				child := nodes[0]
				(*consumer)[child.Name] = *arr
			}
		}
	}

}

// trimNonGraphic returns a slice of the string s, with all leading and trailing
// non graphic characters and spaces removed.
//
// Graphic characters include letters, marks, numbers, punctuation, symbols,
// and spaces, from categories L, M, N, P, S, Zs.
// Spacing characters are set by category Z and property Pattern_White_Space.
func trimNonGraphic(s string) string {
	if s == "" {
		return s
	}

	var first *int
	var last int
	for i, r := range []rune(s) {
		if !unicode.IsGraphic(r) || unicode.IsSpace(r) {
			continue
		}

		if first == nil {
			f := i // copy i
			first = &f
			last = i
		} else {
			last = i
		}
	}

	// If first is nil, it means there are no graphic characters
	if first == nil {
		return ""
	}

	return string([]rune(s)[*first : last+1])
}
