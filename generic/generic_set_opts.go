package generic

import "sbl.systems/go/synwork/plugin-sdk/schema"

var Method_generic_set_intersect = &schema.Method{
	Schema: Method_generic_join.Schema,
	Result: map[string]*schema.Schema{
		"intersection": LeftRightSchema,
	},
	Description: `compares each element of the left side list with each element on the right side list and 
	gives a result list back. The enttities have two fields: "left" and "right"
	`,
	ExecFunc: BuildSetFunc("intersection", func(c *SetCommon) error {
		cache := map[JoinKey]interface{}{}
		for _, li := range c.LeftData {
			lim := li.(map[string]interface{})
			lk := c.LeftKeyBuilder.Build(lim)
			cache[lk] = li
		}
		for _, ri := range c.RightData {
			lim := ri.(map[string]interface{})
			rk := c.RightKeyBuilder.Build(lim)
			if li, ok := cache[rk]; ok {
				c.Result.AddLeftRight(li, ri)
			}
		}
		return nil
	}),
}

var Method_generic_set_relative_complement = &schema.Method{
	Schema: Method_generic_join.Schema,
	Result: map[string]*schema.Schema{
		"complement": LeftRightSchema,
	},
	Description: `compares each element of the left side list with each element on the right side list and 
	gives a result list back. The enttities have two fields: "left" and "right". right is always empty
	`,
	ExecFunc: BuildSetFunc("complement", func(c *SetCommon) error {
		cache := map[JoinKey]interface{}{}
		for _, li := range c.LeftData {
			lim := li.(map[string]interface{})
			lk := c.LeftKeyBuilder.Build(lim)
			cache[lk] = li
		}
		for _, ri := range c.RightData {
			lim := ri.(map[string]interface{})
			rk := c.RightKeyBuilder.Build(lim)
			delete(cache, rk)
		}
		for _, v := range cache {
			c.Result.AddLeftRight(v, nil)
		}
		return nil
	}),
}

var Method_generic_set_symmetric_difference = &schema.Method{
	Schema: Method_generic_join.Schema,
	Result: map[string]*schema.Schema{
		"complement": LeftRightSchema,
	},
	Description: `compares each element of the left side list with each element on the right side list and 
	gives a result list back. The enttities have two fields: "left" and "right". 
	`,
	ExecFunc: BuildSetFunc("complement", func(c *SetCommon) error {
		cacheLeft := map[JoinKey]interface{}{}
		cacheRight := map[JoinKey]interface{}{}
		for _, li := range c.LeftData {
			lim := li.(map[string]interface{})
			lk := c.LeftKeyBuilder.Build(lim)
			cacheLeft[lk] = li
		}
		for _, ri := range c.RightData {
			lim := ri.(map[string]interface{})
			rk := c.RightKeyBuilder.Build(lim)
			cacheRight[rk] = ri
			if _, ok := cacheLeft[rk]; !ok {
				c.Result.AddLeftRight(nil, ri)
			}
		}
		for k, li := range cacheLeft {
			if _, ok := cacheRight[k]; !ok {
				c.Result.AddLeftRight(li, nil)
			}
		}
		return nil
	}),
}
