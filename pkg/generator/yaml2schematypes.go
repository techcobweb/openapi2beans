package generator

import (
	"log"
	"maps"
	"strings"

	openapi2beans_errors "github.com/techcobweb/openapi2beans/pkg/errors"
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
	OPENAPI_YAML_KEYWORD_ENUM		 = "enum"
)

func getSchemaTypesFromYaml(apiyaml []byte) (parsedSchemaTypes map[string]*SchemaType, err error) {
	var schemasMap map[interface{}]interface{}
	entireYamlMap := make(map[string]interface{})

	err = yaml.Unmarshal(apiyaml, &entireYamlMap)

	if err == nil {
		schemasMap, err = retrieveSchemasMapFromEntireYamlMap(entireYamlMap)

		if err == nil {
			var properties map[string]*Property
			parsedSchemaTypes, properties, err = retrieveSchemaTypesFromMap(schemasMap, "#/components/schemas")
			resolveReferences(properties)
		}
	}

	return parsedSchemaTypes, err
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
			err = openapi2beans_errors.NewError("RetrieveSchemasMapFromEntireYamlMap: Failed to find schemas within %v", entireYamlMap)
		}
	} else {
		err = openapi2beans_errors.NewError("RetrieveSchemasMapFromEntireYamlMap: Failed to find components within %v", entireYamlMap)
	}
	return schemasMap, err
}

func retrieveSchemaTypesFromMap(inputMap map[interface{}]interface{}, yamlPath string) (schemaTypes map[string]*SchemaType, properties map[string]*Property, err error) {
	schemaTypes = make(map[string]*SchemaType)
	properties = make(map[string]*Property)

	for subMapKey, subMapObj := range inputMap {
		log.Printf("RetrieveSchemaTypesFromMap: %v\n", subMapObj)

		subMap := subMapObj.(map[interface{}]interface{})
		apiSchemaPartPath := yamlPath + "/" + subMapKey.(string)
		varName := subMapKey.(string)

		var typeName string
		var cardinality Cardinality
		var possibleValues map[string]string
		description := retrieveDescription(subMap)
		typeName, cardinality, err = retrieveVarType(subMap, apiSchemaPartPath)
		possibleValues = retrievePossibleValues(subMap)

		if err == nil {
			property := NewProperty(subMapKey.(string), apiSchemaPartPath, description, typeName, possibleValues, nil, cardinality)
			
			if typeName == "object" {
				var nestedProperties map[string]*Property
				var nestedSchemaTypes map[string]*SchemaType

				nestedSchemaTypes, nestedProperties, err = retrieveNestedProperties(subMap, apiSchemaPartPath)

				if err == nil {
					resolvedType := NewSchemaType(varName, description, property, nestedProperties)
					maps.Copy(properties, nestedProperties)
					maps.Copy(schemaTypes, nestedSchemaTypes)
					property.SetResolvedType(resolvedType)

					schemaTypes[apiSchemaPartPath] = resolvedType
				}
			} else if property.IsEnum() {
				enumSchemaType := NewSchemaType(varName, description, property, nil)
				property.SetResolvedType(enumSchemaType)
				schemaTypes[apiSchemaPartPath] = enumSchemaType
			}

			if property.IsConstant() {
				property.name = strings.ToUpper(property.name)
			}

			if err == nil {
				properties[apiSchemaPartPath] = property
			}
		}
	}

	return schemaTypes, properties, err
}

func resolveReferences(properties map[string]*Property) error {
	var err error
	log.Printf("ResolveReferences: resolving references in property map: %v\n", properties)
	for _, property := range properties {
		if property.IsReferencing() {
			log.Printf("ResolveReferences: attempting to resolve: %v\n", property)
			referencingPath := strings.Split(property.GetType(), ":")[1]
			referencedProp, isRefPropPresent := properties[referencingPath]
			if isRefPropPresent {
				log.Printf("ResolvingReferences: resolving %v with %v\n", property, referencedProp)
				property.Resolve(referencedProp)
			} else {
				err = openapi2beans_errors.NewError("ResolveReferences: Failed to find referenced property for %v\n", property)
			}
		}
	}
	return err
}

func retrieveNestedProperties(subMap map[interface{}]interface{}, yamlPath string) (schemaTypes map[string]*SchemaType, properties map[string]*Property, err error) {
	var schemaPropertiesMap map[interface{}]interface{}

	propertiesObj, isPropertyPresent := subMap[OPENAPI_YAML_KEYWORD_PROPERTIES]
	if isPropertyPresent {
		schemaPropertiesMap = propertiesObj.(map[interface{}]interface{})
		schemaTypes, properties, err = retrieveSchemaTypesFromMap(schemaPropertiesMap, yamlPath)
	}

	return schemaTypes, properties, err
}

func retrieveVarType(variableMap map[interface{}]interface{}, apiSchemaPartPath string) (varType string, cardinality Cardinality, err error) {
	maxCardinality := 0
	varTypeObj, isTypePresent := variableMap[OPENAPI_YAML_KEYWORD_TYPE]
	refObj, isRefPresent := variableMap[OPENAPI_YAML_KEYWORD_REF]

	if isTypePresent {
		varType = varTypeObj.(string)
		if varType == "array" {
			var returnCardinality int
			maxCardinality = 128
			varType, returnCardinality, err = retrieveArrayType(variableMap, apiSchemaPartPath)
			maxCardinality += returnCardinality
		} else {
			maxCardinality = 1
		}
		cardinality = Cardinality {min: getMinCardinality(variableMap), max: maxCardinality}
	} else if isRefPresent {
		varType = "$ref:" + refObj.(string)
	} else {
		err = openapi2beans_errors.NewError("RetrieveVarType: Failed to find required type for %v\n", apiSchemaPartPath)
	}

	return varType, cardinality, err
}

func retrieveArrayType(varMap map[interface{}]interface{}, schemaPartPath string) (arrayType string, maxCardinality int, err error) {

	itemsObj, isItemsPresent := varMap[OPENAPI_YAML_KEYWORD_ITEMS]
	if isItemsPresent {
		itemsMap := itemsObj.(map[interface{}]interface{})

		allOfObj, isAllOfPresent := itemsMap[OPENAPI_YAML_KEYWORD_ALLOF]
		if isAllOfPresent {
			allOfSlice := allOfObj.([]interface{})
			itemsMap = allOfSlice[0].(map[interface{}]interface{})
		}
		var cardinality Cardinality
		arrayType, cardinality, err = retrieveVarType(itemsMap, schemaPartPath)
		if cardinality.max > 1 {
			maxCardinality += 128
		}
		
	} else {
		err = openapi2beans_errors.NewError("RetrieveArrayType: Failed to find required items section for %v\n", schemaPartPath)
	}

	return arrayType, maxCardinality, err
}

func retrieveDescription(subMap map[interface{}]interface{}) (description string) {
	descriptionObj, isDescriptionPresent := subMap[OPENAPI_YAML_KEYWORD_DESCRIPTION]
	if isDescriptionPresent {
		description = descriptionObj.(string)
	}
	return description
}

func getMinCardinality(varMap map[interface{}]interface{}) (minCardinality int) {
	minCardinality = 0
	requiredObj, isRequiredPresent := varMap[OPENAPI_YAML_KEYWORD_REQUIRED]
	if isRequiredPresent {
		if requiredObj.(bool) {
			minCardinality = 1
		}
	}
	return minCardinality
}

func retrievePossibleValues(varMap map[interface{}]interface{}) (possibleValues map[string]string) {
	possibleValues = make(map[string]string)
	enumObj, isEnumPresent := varMap[OPENAPI_YAML_KEYWORD_ENUM]
	if isEnumPresent {
		enums := enumObj.([]interface{})
		for _, enum := range enums {
			enumName := enum.(string)
			possibleValues[enumName] = enumName
		}
	}
	return
}