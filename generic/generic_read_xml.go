package generic

import (
	"bytes"
	"context"
	"io/ioutil"
	"strings"

	"sbl.system/synwork/generic/generic/xml2json"
	"sbl.systems/go/synwork/plugin-sdk/schema"
)

var Method_generic_read_xml = &schema.Method{
	Schema: map[string]*schema.Schema{
		"file_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "the name of the file which should be read",
		},
		"arrays": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "a semicolon seperated list of elements, which should be represented as array",
			DefaultValue: "",
		},
	},
	Result: map[string]*schema.Schema{
		"xml": {
			Type:     schema.TypeGeneric,
			Required: true,
		},
	},
	Description: `create a new json structure with new attributes
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
	ExecFunc: generic_read_xml,
}

func generic_read_xml(ctx context.Context, data *schema.MethodData, client interface{}) error {
	uData := data.GetConfig("file_name").(string)
	b, err := ioutil.ReadFile(uData)
	if err != nil {
		return err
	}
	arrays := data.GetConfig("arrays").(string)
	dec := xml2json.NewDecoder(bytes.NewReader(b))
	dec.SetArrayPaths(strings.Split(arrays, ";")...)
	resultMap := map[string]interface{}{}
	err = dec.Decode(&resultMap)
	if err != nil {
		return err
	}
	// xmlDec := xml.NewDecoder(bytes.NewReader(b))

	// for _, i := range arrData {
	// 	mapItem := i.(map[string]interface{})
	// 	name := mapItem["name"].(string)
	// 	result := mapItem["data"]
	// 	resultMap[name] = result
	// }
	data.SetResult("xml", resultMap)
	return nil
}
