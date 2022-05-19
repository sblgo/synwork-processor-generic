package generic

import (
	"context"
	"fmt"
	"strings"

	"sbl.systems/go/synwork/plugin-sdk/schema"
)

type (
	SetCommon struct {
		LeftData, RightData             []interface{}
		LeftKeyBuilder, RightKeyBuilder *KeyBuilder
		Result                          *Collector
	}
	JoinKeyEntry struct {
		Name  string
		Value interface{}
	}
	JoinKey struct {
		Key string
	}
	KeyEntryBuilder func(m map[string]interface{}) JoinKeyEntry

	KeyBuilder struct {
		KeyFieldBuilder []KeyEntryBuilder
	}

	Collector struct {
		data []interface{}
	}
)

func NewSetProcessor(data *schema.MethodData) (*SetCommon, error) {
	proc := &SetCommon{
		LeftData:        data.GetConfig("left").([]interface{}),
		RightData:       data.GetConfig("right").([]interface{}),
		Result:          &Collector{},
		LeftKeyBuilder:  new(KeyBuilder),
		RightKeyBuilder: new(KeyBuilder),
	}
	wheres := data.GetConfig("where").([]interface{})
	for _, w := range wheres {
		where := w.(map[string]interface{})
		proc.LeftKeyBuilder.Add(newKeyFieldBuilder(where["left_field"].(string)))
		proc.RightKeyBuilder.Add(newKeyFieldBuilder(where["right_field"].(string)))
	}

	return proc, nil
}

func (c *Collector) Add(v interface{}) {
	c.data = append(c.data, v)
}

func (c *Collector) Data() interface{} {
	return c.data
}

func (c *Collector) AddLeftRight(left, right interface{}) {
	m := map[string]interface{}{}
	if left != nil {
		m["left"] = left
	}
	if right != nil {
		m["right"] = right
	}
	c.Add(m)
}

func (kb *KeyBuilder) Build(m map[string]interface{}) JoinKey {
	key := ""
	for _, k := range kb.KeyFieldBuilder {
		field := k(m)
		key = fmt.Sprintf("%s<%v>", key, field.Value)
	}
	return JoinKey{Key: key}
}

func (kb *KeyBuilder) Add(b KeyEntryBuilder) {
	if kb.KeyFieldBuilder == nil {
		kb.KeyFieldBuilder = []KeyEntryBuilder{}
	}
	kb.KeyFieldBuilder = append(kb.KeyFieldBuilder, b)
}

func newKeyFieldBuilder(path string) KeyEntryBuilder {
	pathParts := strings.Split(strings.Trim(path, "/"), "/")
	return func(m map[string]interface{}) JoinKeyEntry {
		value, _ := schema.GetValueMap(m, pathParts)
		return JoinKeyEntry{
			Name:  path,
			Value: value,
		}
	}
}

func BuildSetFunc(resultName string, setFunc func(c *SetCommon) error) func(ctx context.Context, data *schema.MethodData, client interface{}) error {

	return func(ctx context.Context, data *schema.MethodData, client interface{}) error {
		if proc, err := NewSetProcessor(data); err != nil {
			return err
		} else if err := setFunc(proc); err != nil {
			return err
		} else {
			data.SetResult(resultName, proc.Result.Data())
			return nil
		}
	}
}
