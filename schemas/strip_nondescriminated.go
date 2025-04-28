package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	f, err := os.ReadFile("./api.github.com.json")
	if err != nil {
		fmt.Println("error reading", err)
	}
	var v map[string]any
	err = json.Unmarshal(f, &v)
	if err != nil {
		fmt.Println("error unmarshalling", err)
	}
	// fixup oneOf and allOf from components that don't have a discriminator
	componentsEl := v["components"]
	if components, ok := componentsEl.(map[string]any); ok {
		transformUndiscriminated(components)
	}
	// fixup oneOf and allOf from paths that don't have a discriminator
	pathsEl := v["paths"]
	if paths, ok := pathsEl.(map[string]any); ok {
		transformUndiscriminated(paths)
	}
	// remove required elements. Required oneOf elements without discriminator are a problem.
	// ideally this would remove only if it's a oneOf
	pruneRequired(v)
	b, err := json.Marshal(v)
	fmt.Print(string(b))
}

// XXX should track the path it traversed and output what has been modified
//
//		objectType: {"type": "object"} | {"$ref": ""} // assumes refs are only to objects
//		mixedStringNumber: [{"type": "string"}, {"type": "number"|"integer"}...]
//		mixedObject: [objectType, objectType...]
//		|case  | replacement | comment
//	 |---   |---          |---
//		|oneOf|allOf: mixedStringNumber       | {"type": "string"} | remove integer/number
//		|oneOf|allOf: mixedObject             | {"type": "object", "additionalProperties": true}
//		|oneOf|allOf: [{"oneOf"} | {"allOf"}] | {"type": "object", "additionalProperties": true}
//		|oneOf|allOf: [{type: "object",...}] // all objects or $refs
//		  replace with object // I think this is only in the spec once
func transformUndiscriminated(obj map[string]any) {
	for k := range obj {
		for _, keyword := range []string{"oneOf", "anyOf"} {
			if k == keyword {
				if _, ok := obj["discriminator"]; !ok {
					switch {
					// mixed string and object types with no discriminator: delete and rely on additionalData
					//	oneOf|allOf [type: "string", type: "object"]
					case hasType("string", obj[k]) && hasType("object", obj[k]):
						delete(obj, k)
					// mixed string and integer/number types with no discriminator: use string or delete
					case hasType("string", obj[k]) && (hasType("integer", obj[k]) || hasType("number", obj[k])):
						if multiple := obj[k].([]map[string]any); ok {
							if len(multiple) == 2 {
								keepOnlyType("string", multiple)
							} else {
								delete(obj, k)
							}
						}
					// mixed string and number types with no discriminator: use string or delete
					//	oneOf|allOf [type: "string", type: "number"]
					case hasType("string", obj[k]) && hasType("number", obj[k]):
						if multiple := obj[k].([]map[string]any); ok {
							if len(multiple) == 2 {
								keepOnlyType("string", multiple)
							} else {
								delete(obj, k)
							}
						}
					}
					// all $refs, assume all object and replace with generic object type with additonalproperties true
				}
			} else {
				e := obj[i]
				switch e.(type) {
				case map[string]any:
					transformUndiscriminated(e.(map[string]any))
				}
			}
		}
	}
}

func hasType(t string, o any) bool {
	if obj, ok := o.([]map[string]any); ok {
		for _, e := range obj {
			if e["type"] == t {
				return true
			}
		}
	}
	return false
}

func keepOnlyType(t string, o any) {
	if slice, ok := o.([]map[string]any); ok {
		for i, e := range slice {
			if e["type"] == t {
				slice = slice[i : i+1]
			}
		}
	}
}

func pruneRequired(obj map[string]any) {
	for i := range obj {
		if i == "required" {
			delete(obj, i)
		} else {
			e := obj[i]
			switch e.(type) {
			case map[string]any:
				pruneRequired(e.(map[string]any))
			}
		}
	}
}
