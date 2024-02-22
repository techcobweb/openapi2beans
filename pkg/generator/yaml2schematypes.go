package generator

import "gopkg.in/yaml.v2"

const (
	OPENAPI_YAML_KEYWORD_COMPONENTS  = "components"
	OPENAPI_YAML_KEYWORD_SCHEMAS     = "schemas"
	OPENAPI_YAML_KEYWORD_DESCRIPTION = "description"
	OPENAPI_YAML_KEYWORD_PROPERTIES  = "properties"
	OPENAPI_YAML_KEYWORD_TYPE        = "type"
	OPENAPI_YAML_KEYWORD_REQUIRED    = "required"
	OPENAPI_YAML_KEYWORD_ITEMS       = "items"
	OPENAPI_YAML_KEYWORD_ALLOF       = "allOf"
	OPENAPI_YAML_KEYWORD_REF         = "$ref"
)

func getSchemaTypesFromYaml(apiyaml []byte, packageName string) (schemaTypes []SchemaType, err error) {
	var schemasMap map[interface{}]interface{}
	entireYamlMap := make(map[string]interface{})

	err = yaml.Unmarshal(apiyaml, &entireYamlMap)

	if err == nil {
		schemasMap, err = retrieveSchemasMapFromEntireYamlMap(entireYamlMap)

		if err == nil {
			var parsedSchemas map[string]SchemaType
			var referencingStructures map[string]SchemaType
			parsedSchemas, referencingStructures, err = retrieveStructuresFromMap(schemasMap, "#/components/schemas")
			parsedSchemas = resolveReferences(parsedSchemas, referencingStructures)

			for _, structure := range parsedSchemas {
				if structure.GetType() == "object" {
					bean := Bean{
						object:      structure.(Object),
						beanPackage: packageName,
					}
					beans = append(beans, bean)
				}
			}
		}
	}

	return schemaTypes, err
}