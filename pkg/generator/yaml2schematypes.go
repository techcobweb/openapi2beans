package generator

// import "gopkg.in/yaml.v2"

// const (
// 	OPENAPI_YAML_KEYWORD_COMPONENTS  = "components"
// 	OPENAPI_YAML_KEYWORD_SCHEMAS     = "schemas"
// 	OPENAPI_YAML_KEYWORD_DESCRIPTION = "description"
// 	OPENAPI_YAML_KEYWORD_PROPERTIES  = "properties"
// 	OPENAPI_YAML_KEYWORD_TYPE        = "type"
// 	OPENAPI_YAML_KEYWORD_REQUIRED    = "required"
// 	OPENAPI_YAML_KEYWORD_ITEMS       = "items"
// 	OPENAPI_YAML_KEYWORD_ALLOF       = "allOf"
// 	OPENAPI_YAML_KEYWORD_REF         = "$ref"
// )

// func getSchemaTypesFromYaml(apiyaml []byte, packageName string) (schemaTypes []SchemaType, err error) {
// 	var schemasMap map[interface{}]interface{}
// 	entireYamlMap := make(map[string]interface{})

// 	err = yaml.Unmarshal(apiyaml, &entireYamlMap)

// 	if err == nil {
// 		schemasMap, err = retrieveSchemasMapFromEntireYamlMap(entireYamlMap)

// 		if err == nil {
// 			var parsedSchemas map[string]SchemaType
// 			var referencingStructures map[string]SchemaType
// 			parsedSchemas, referencingStructures, err = retrieveSchemaTypesFromMap(schemasMap, "#/components/schemas")
// 			parsedSchemas = resolveReferences(parsedSchemas, referencingStructures)

// 			for _, structure := range parsedSchemas {
// 				if structure.GetType() == "object" {
// 					bean := Bean{
// 						object:      structure.(Object),
// 						beanPackage: packageName,
// 					}
// 					beans = append(beans, bean)
// 				}
// 			}
// 		}
// 	}

// 	return schemaTypes, err
// }

// func retrieveSchemasMapFromEntireYamlMap(entireYamlMap map[string]interface{}) (map[interface{}]interface{}, error) {
// 	var err error
// 	schemasMap := make(map[interface{}]interface{})

// 	components, isComponentsPresent := entireYamlMap[OPENAPI_YAML_KEYWORD_COMPONENTS]
// 	if isComponentsPresent {
// 		componentsMap := components.(map[interface{}]interface{})
// 		schemas, isSchemasPresent := componentsMap[OPENAPI_YAML_KEYWORD_SCHEMAS]
// 		if isSchemasPresent {
// 			schemasMap = schemas.(map[interface{}]interface{})
// 		} else {
// 			err = NewError("Failed to find schemas within %v", entireYamlMap)
// 		}
// 	} else {
// 		err = NewError("Failed to find components within %v", entireYamlMap)
// 	}
// 	return schemasMap, err
// }