package generic

import (
	"context"

	"sbl.systems/go/synwork/plugin-sdk/schema"
)

// left = $method.do1.left
// right = $method.do2.right
// where {
// 	left_field = "la"
// 	right_field = "ra"
// }

var LeftRightSchema = &schema.Schema{
	Type:     schema.TypeList,
	Required: true,
	Elem: map[string]*schema.Schema{
		"left": {
			Type:        schema.TypeGeneric,
			Required:    true,
			Description: "entry of the left side ",
		},
		"right": {
			Type:        schema.TypeGeneric,
			Required:    true,
			Description: "entry of the right side ",
		},
	},
}

var Method_generic_join = &schema.Method{
	Schema: map[string]*schema.Schema{
		"left": {
			Type:        schema.TypeList,
			Required:    true,
			ElemType:    schema.TypeGeneric,
			Description: "reference to one element to join",
		},
		"right": {
			Type:        schema.TypeList,
			Required:    true,
			ElemType:    schema.TypeGeneric,
			Description: "reference to other element to join",
		},
		"where": {
			Type:        schema.TypeList,
			Required:    true,
			Description: "condition for joining",
			Elem: map[string]*schema.Schema{
				"left_field": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "path in the left side to field",
				},
				"right_field": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "path in the right side to field",
				},
			},
		},
	},
	Result: map[string]*schema.Schema{
		"join": LeftRightSchema,
	},
	Description: `compares each element of the left side list with each element on the right side list and 
	gives a result list back. The enttities have two fields: "left" and "right"
	`,
	ExecFunc: generic_join,
}

func generic_join(ctx context.Context, data *schema.MethodData, client interface{}) error {
	left := data.GetConfig("left").([]interface{})
	right := data.GetConfig("right").([]interface{})
	result := []interface{}{}
	leftKeyBuilder := new(KeyBuilder)
	rightKeyBuilder := new(KeyBuilder)
	wheres := data.GetConfig("where").([]interface{})
	for _, w := range wheres {
		where := w.(map[string]interface{})
		leftKeyBuilder.Add(newKeyFieldBuilder(where["left_field"].(string)))
		rightKeyBuilder.Add(newKeyFieldBuilder(where["right_field"].(string)))
	}
	cache := map[JoinKey]interface{}{}
	for _, li := range left {
		lim := li.(map[string]interface{})
		lk := leftKeyBuilder.Build(lim)
		cache[lk] = li
	}

	for _, ri := range right {
		lim := ri.(map[string]interface{})
		rk := rightKeyBuilder.Build(lim)
		if li, ok := cache[rk]; ok {
			result = append(result, map[string]interface{}{
				"left":  li,
				"right": ri,
			})
		}
	}

	data.SetResult("join", result)
	return nil
}
