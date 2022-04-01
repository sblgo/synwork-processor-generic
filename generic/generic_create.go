package generic

import (
	"context"
	"fmt"

	"sbl.systems/go/synwork/plugin-sdk/schema"
)

var Method_generic_create = &schema.Method{
	Schema: map[string]*schema.Schema{
		"attribute": {
			Type:     schema.TypeList,
			Required: true,
			Elem: map[string]*schema.Schema{
				"name": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "the name of the new attribute in the result",
				},
				"data": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "reference to an result of a other processor",
				},
			},
		},
	},
	Result: map[string]*schema.Schema{
		"result": {
			Type:     schema.TypeGeneric,
			Required: true,
		},
	},
	Description: `create a new structure with new attributes
"result" : {
	"attr01": {

	},
	"attr02": {

	},
	"attr03": {

	},
	etc...
}
	
	`,
	ExecFunc: generic_create,
}

func generic_create(ctx context.Context, data *schema.MethodData, client interface{}) error {
	uData := data.GetConfig("attribute")
	if uData == nil {
		return fmt.Errorf("missing configuration")
	}
	arrData := uData.([]interface{})
	resultMap := map[string]interface{}{}
	for _, i := range arrData {
		mapItem := i.(map[string]interface{})
		name := mapItem["name"].(string)
		result := mapItem["data"]
		resultMap[name] = result
	}
	data.SetResult("result", resultMap)
	return nil
}
