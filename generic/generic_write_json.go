package generic

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"sbl.systems/go/synwork/plugin-sdk/schema"
)

var Method_generic_write_json = &schema.Method{
	Schema: map[string]*schema.Schema{
		"file_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "the name of the file",
		},
		"data": {
			Type:        schema.TypeGeneric,
			Required:    true,
			Description: "reference to an result of a other processor",
		},
		"indent": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "",
			DefaultValue: "",
		},
		"prefix": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "",
			DefaultValue: "",
		},
	},
	Result: map[string]*schema.Schema{
		"result": {
			Type:     schema.TypeGeneric,
			Required: true,
		},
	},
	Description: `write the result of a other processer `,
	ExecFunc:    generic_write_json,
}

func generic_write_json(ctx context.Context, data *schema.MethodData, client interface{}) error {
	uData := data.GetConfig("data")
	if uData == nil {
		return fmt.Errorf("missing configuration")
	}
	fileName := data.GetConfig("file_name").(string)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return fmt.Errorf("error opening the file %s. %s", fileName, err.Error())
	}
	prefix := data.GetConfig("prefix").(string)
	indent := data.GetConfig("indent").(string)
	defer file.Close()
	enc := json.NewEncoder(file)
	enc.SetIndent(prefix, indent)
	err = enc.Encode(uData)
	return err
}
