package generator

import (
	"errors"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

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
			log.Printf("Failed to find required type within %v", entireYamlMap)
			err = errors.New("failed to find schemas within components section of given yaml")
		}
	} else {
		log.Printf("Failed to find components within %v", entireYamlMap)
		err = errors.New("failed to find components within given yaml")
	}
	return schemasMap, err
}

func retrieveStructuresFromMap(inputMap map[interface{}]interface{}, yamlPath string) (structures map[string]SchemaPart, referencingStructures map[string]SchemaPart, err error) {
	structures = make(map[string]SchemaPart)
	referencingStructures = make(map[string]SchemaPart)

	for subMapKey, subMapObj := range inputMap {
		log.Printf("%v\n", subMapObj)
		subMap := subMapObj.(map[interface{}]interface{})
		var varType string

		description := retrieveDescription(subMap)
		varType, err := retrieveVarType(subMap, subMapKey.(string))

		if err != nil {
			log.Printf("Failed to work out the type of the variable.\n")
		} else {

			if varType == "object" {
				var variables map[string]SchemaPart
				newPath := yamlPath + "/" + subMapKey.(string)
				variables, err = retrieveVariables(subMap, newPath)

				if err == nil {
					object := Object{
						varName:     subMapKey.(string),
						description: description,
						varTypeName: varType,
						variables:   variables,
					}
					objectPath := yamlPath + "/" + object.varName
					structures[objectPath] = object
				}
			} else {
				if varType == "array" {
					varType, err = retrieveArrayType(subMap, subMapKey.(string))
				}

				if err == nil {
					isSetInConstructor := isSetInConstructor(subMap)

					variable := Variable{
						varName:            subMapKey.(string),
						varDescription:     description,
						varTypeName:        varType,
						isSetInConstructor: isSetInConstructor,
					}

					varPath := yamlPath + "/" + variable.varName
					if strings.Split(varType, ":")[0] == "$ref" {
						referencingStructures[varPath] = variable
					}
					structures[varPath] = variable
				}
			}
		}
	}

	return structures, referencingStructures, err
}

// Recieves 2 maps of Structures with the key a reference to it's reference path. The aim is to add the reference to an Object or array.
//
//	So Object {
//		varType: object
//	 variables {
//	   Variable {
//	     varName: nestedObject
//			varType: $ref:#components/schemas/NestedObject
//	  }
//
// }
//
//	To Object {
//		varType: object
//	 variables {
//	   Variable {
//	     varName: nestedObject
//			varType: object
//			  properties:
//			    randomVar
//				  type: string
//	  }
//
// }
// can we do that by saying that variables are pointers? and then changing the variable that is being pointed to?
// so for each structure in referencing structures the var type needs to change. (I have had the varType hold the reference)
func resolveReferences(parsedStructures map[string]SchemaPart, referencingStructures map[string]SchemaPart) map[string]SchemaPart {
	for refPath, schema := range referencingStructures {
		croppedReference := strings.Split(schema.GetType(), ":")[1]
		referencedStructure := parsedStructures[croppedReference]
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

func retrieveVarType(variableMap map[interface{}]interface{}, varMapKeyString string) (string, error) {
	var varType string
	var err error
	varTypeObj, isTypePresent := variableMap[OPENAPI_YAML_KEYWORD_TYPE]
	refObj, isRefPresent := variableMap[OPENAPI_YAML_KEYWORD_REF]

	if isTypePresent {
		varType = getJavaReadableType(varTypeObj.(string))
	} else if isRefPresent {
		varType = "$ref:" + refObj.(string)
	} else {
		log.Printf("Failed to find required type for %v", variableMap)
		err = errors.New("failed to find required type for " + varMapKeyString)
	}
	return varType, err
}

func retrieveArrayType(subMap map[interface{}]interface{}, subMapKeyString string) (string, error) {
	var arrayType string
	var err error
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
		} else {
			log.Printf("Failed to find required type within items section for %v", subMap)
			err = errors.New("failed to find required type within items section for " + subMapKeyString)
		}
	} else {
		log.Printf("Failed to find required items section for %v", subMap)
		err = errors.New("failed to find required items section for " + subMapKeyString)
	}

	return arrayType, err
}

func retrieveDescription(subMap map[interface{}]interface{}) string {
	description := ""
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
func getJavaReadableType(yamlReadableType string) string {
	var javaReadableType string
	if yamlReadableType == "string" {
		javaReadableType = "String"
	} else {
		javaReadableType = yamlReadableType
	}

	return javaReadableType
}
