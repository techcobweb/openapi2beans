package generator

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

func NewError(template string, params ... interface{}) error {
	msg := fmt.Sprintf(template, params...)
	log.Print(msg)
	err := errors.New(msg)
	return err
}

func getBeansFromYaml(apiyaml []byte, packageName string) ([]Bean, error) {
	var beans []Bean
	var schemasMap map[interface{}]interface{}
	entireYamlMap := make(map[string]interface{})

	err := yaml.Unmarshal(apiyaml, &entireYamlMap)

	if err == nil {
		schemasMap, err = retrieveSchemasMapFromEntireYamlMap(entireYamlMap)

		if err == nil {
			var parsedSchemas map[string]SchemaPart
			var referencingStructures map[string]SchemaPart
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

	return beans, err
}

func retrieveSchemasMapFromEntireYamlMap(entireYamlMap map[string]interface{}) (map[interface{}]interface{}, error) {
	var err error
	schemasMap := make(map[interface{}]interface{})

	components, isComponentsPresent := entireYamlMap[OPENAPI_YAML_KEYWORD_COMPONENTS]
	if isComponentsPresent {
		componentsMap := components.(map[interface{}]interface{})
		schemas, isSchemasPresent := componentsMap[OPENAPI_YAML_KEYWORD_SCHEMAS]
		if isSchemasPresent {
			schemasMap = schemas.(map[interface{}]interface{})
		} else {
			err = NewError("Failed to find schemas within %v", entireYamlMap)
		}
	} else {
		err = NewError("Failed to find components within %v", entireYamlMap)
	}
	return schemasMap, err
}

func retrieveStructuresFromMap(inputMap map[interface{}]interface{}, yamlPath string) (schemaParts map[string]SchemaPart, referencingSchemaParts map[string]SchemaPart, err error) {
	schemaParts = make(map[string]SchemaPart)
	referencingSchemaParts = make(map[string]SchemaPart)

	for subMapKey, subMapObj := range inputMap {
		log.Printf("%v\n", subMapObj)
		subMap := subMapObj.(map[interface{}]interface{})
		var varType string
		apiSchemaPartPath := yamlPath + "/" + subMapKey.(string)
		
		description := retrieveDescription(subMap)
		varType, err := retrieveVarType(subMap, apiSchemaPartPath)

		if err != nil {
			// do something
		} else {
			if varType == "object" {
				var variables map[string]SchemaPart
				
				variables, err = retrieveVariables(subMap, apiSchemaPartPath)
				
				if err == nil {
					object := Object {
						varName: subMapKey.(string),
						description: description,
						varTypeName: varType,
						variables: variables,
					}
					schemaParts[apiSchemaPartPath] = object
				}
			} else {
				isSetInConstructor := isSetInConstructor(subMap)
	
				variable := Variable {
					varName: subMapKey.(string),
					varDescription: description,
					varTypeName: varType,
					isSetInConstructor: isSetInConstructor,
				}

				if strings.Split(varType, ":")[0] == "$ref" {
					referencingSchemaParts[apiSchemaPartPath] = variable
				}
				schemaParts[apiSchemaPartPath] = variable
			}
		}
	}

	return schemaParts, referencingSchemaParts, err
}

func resolveReferences(parsedStructures map[string]SchemaPart, referencingSchemaParts map[string]SchemaPart) (map[string]SchemaPart) {
	for refPath, schema := range referencingSchemaParts {
		croppedReferencePath := strings.Split(schema.GetType(), ":")[1]
		referencedStructure := parsedStructures[croppedReferencePath]
		parsedStructures[refPath] = referencedStructure
	}
	return parsedStructures
}

func retrieveVariables(subMap map[interface{}]interface{}, yamlPath string) (map[string]SchemaPart, error) {
	var properties map[interface{}]interface{}
	var variables map[string]SchemaPart
	var err error
	propertiesObj, isPropertyPresent := subMap[OPENAPI_YAML_KEYWORD_PROPERTIES]
	if isPropertyPresent {
		properties = propertiesObj.(map[interface{}]interface{})
		variables, _, err = retrieveStructuresFromMap(properties, yamlPath)
	}
	return variables, err
}

func retrieveVarType(variableMap map[interface{}]interface{}, apiSchemaPartPath string) (string, error) {
	var varType string
	var err error
	varTypeObj, isTypePresent := variableMap[OPENAPI_YAML_KEYWORD_TYPE]
	refObj, isRefPresent := variableMap[OPENAPI_YAML_KEYWORD_REF]

	if isTypePresent {
		varType = getJavaReadableType(varTypeObj.(string))
		if varType == "array" {
			varType, err = retrieveArrayType(variableMap, apiSchemaPartPath)
		}
	} else if isRefPresent {
		varType = "$ref:" + refObj.(string)
	} else {
		err = NewError("Failed to find required type for %v\n", apiSchemaPartPath)
	}
	
	return varType, err
}

func retrieveArrayType(subMap map[interface{}]interface{}, schemaPartPath string) (arrayType string, err error) {

	itemsObj, isItemsPresent := subMap[OPENAPI_YAML_KEYWORD_ITEMS]
	if isItemsPresent {
		itemsMap := itemsObj.(map[interface{}]interface{})

		allOfObj, isAllOfPresent := itemsMap[OPENAPI_YAML_KEYWORD_ALLOF]
		if isAllOfPresent {
			allOfSlice := allOfObj.([]interface{})
			itemsMap = allOfSlice[0].(map[interface{}]interface{})
		}

		arrayTypeObj, isArrayTypePresent := itemsMap[OPENAPI_YAML_KEYWORD_TYPE]

		if isArrayTypePresent {
			arrayType = getJavaReadableType(arrayTypeObj.(string)) + "[]"
		}else {
			err = NewError("Failed to find required type within items section for %v\n", schemaPartPath)
		}
	} else {
		err = NewError("Failed to find required items section for %v\n", schemaPartPath)
	}

	return arrayType, err
}

func retrieveDescription(subMap map[interface{}]interface{}) (description string) {
	descriptionObj, isDescriptionPresent := subMap[OPENAPI_YAML_KEYWORD_DESCRIPTION]
	if isDescriptionPresent {
		description = descriptionObj.(string)
	}
	return description
}

func isSetInConstructor(subMap map[interface{}]interface{}) bool {
	isSetInConstructor := false
	requiredObj, isRequiredPresent := subMap[OPENAPI_YAML_KEYWORD_REQUIRED]
	if isRequiredPresent {
		isSetInConstructor = requiredObj.(bool)
	}
	return isSetInConstructor
}

// To be expanded on if necessary
// Now to be moved to schemtype2javastruct transformer
func getJavaReadableType(yamlReadableType string) (javaReadableType string) {
	if yamlReadableType == "string" {
		javaReadableType = "String"
	} else {
		javaReadableType = yamlReadableType
	}

	return javaReadableType
}
