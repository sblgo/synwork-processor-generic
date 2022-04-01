package generic

import (
	"context"

	"sbl.systems/go/synwork/plugin-sdk/plugin"
	"sbl.systems/go/synwork/plugin-sdk/schema"
)

var Opts = plugin.PluginOptions{
	Provider: func() schema.Processor {
		return schema.Processor{
			Schema: map[string]*schema.Schema{},
			MethodMap: map[string]*schema.Method{
				"create":     Method_generic_create,
				"write_json": Method_generic_write_json,
				"read-xml":   Method_generic_read_xml,
				"join":       Method_generic_join,
			},
			InitFunc: func(ctx context.Context, od *schema.ObjectData, i interface{}) (interface{}, error) { return nil, nil },
		}
	},
}
