package generic

import (
	"testing"

	"sbl.systems/go/synwork/plugin-sdk/tunit"
)

func TestGenericJoin01(t *testing.T) {
	_defs := `
	method "join" "dum" "join01" {
		left = $method.do1.left
		right = $method.do2.right
		where {
			left_field = "la"
			right_field = "ra"
		}
	}
	`
	mm := tunit.MethodMock{
		ProcessorDef: Opts.Provider,
		InstanceMock: struct{}{},
		ExecFunc:     generic_join,
		References: map[string]interface{}{
			"method": map[string]interface{}{
				"do1": map[string]interface{}{
					"left": []interface{}{
						map[string]interface{}{
							"la": 1,
							"lb": "01",
						},
						map[string]interface{}{
							"la": 3,
							"lb": "03",
						},
						map[string]interface{}{
							"la": 4,
							"lb": "04",
						},
						map[string]interface{}{
							"la": 5,
							"lb": "05",
						},
					},
				},
				"do2": map[string]interface{}{
					"right": []interface{}{
						map[string]interface{}{
							"ra": 0,
							"rb": "01",
						},
						map[string]interface{}{
							"ra": 3,
							"rb": "03",
						},
						map[string]interface{}{
							"ra": 4,
							"rb": "04",
						},
						map[string]interface{}{
							"ra": 7,
							"rb": "05",
						},
					},
				},
			},
		},
	}
	result := tunit.CallMockMethod(t, mm, _defs)
	join := result["join"].([]interface{})
	if len(join) != 2 {
		t.Fatal()
	}
}
