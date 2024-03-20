package generator

import (
	"log"
	"strings"

	openapi2beans_errors "github.com/techcobweb/openapi2beans/pkg/errors"
	"gopkg.in/yaml.v3"
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
var (
	schemaTypes = make(map[string]*SchemaType)
	properties = make(map[string]*Property)
	errList []error
	arrayDimensions = 0
)

func getSchemaTypesFromYaml(apiyaml []byte) (map[string]*SchemaType, []error, error) {
	var schemasMap map[string]interface{}
	var fatalErr error
	entireYamlMap := make(map[string]interface{})

	fatalErr = yaml.Unmarshal(apiyaml, &entireYamlMap)

	if fatalErr == nil {
		schemasMap, fatalErr = retrieveSchemasMapFromEntireYamlMap(entireYamlMap)

		if fatalErr == nil {
			retrieveSchemaComponentsFromMap(schemasMap, "#/components/schemas")
			resolveReferences()
		}
	}

	return schemaTypes, errList, fatalErr
}

func retrieveSchemasMapFromEntireYamlMap(entireYamlMap map[string]interface{}) (map[string]interface{}, error) {
	var err error
	schemasMap := make(map[string]interface{})

	components, isComponentsPresent := entireYamlMap[OPENAPI_YAML_KEYWORD_COMPONENTS]
	if isComponentsPresent {
		componentsMap := components.(map[string]interface{})
		schemas, isSchemasPresent := componentsMap[OPENAPI_YAML_KEYWORD_SCHEMAS]
		if isSchemasPresent {
			schemasMap = schemas.(map[string]interface{})
		} else {
			err = openapi2beans_errors.NewError("RetrieveSchemasMapFromEntireYamlMap: Failed to find schemas within %v", entireYamlMap)
		}
	} else {
		err = openapi2beans_errors.NewError("RetrieveSchemasMapFromEntireYamlMap: Failed to find components within %v", entireYamlMap)
	}
	return schemasMap, err
}

func retrieveSchemaComponentsFromMap(inputMap map[string]interface{}, parentPath string) {
	var err error
	for subMapKey, subMapObj := range inputMap {
		log.Printf("RetrieveSchemaTypesFromMap: %v\n", subMapObj)

		subMap := subMapObj.(map[string]interface{})
		apiSchemaPartPath := parentPath + "/" + subMapKey
		varName := subMapKey

		var typeName string
		var cardinality Cardinality
		description := retrieveDescription(subMap)
		typeName, cardinality, err = retrieveVarType(subMap, apiSchemaPartPath)
		arrayDimensions = 0
		possibleValues := retrievePossibleValues(subMap)

		if err == nil {
			property := NewProperty(subMapKey, apiSchemaPartPath, description, typeName, possibleValues, nil, cardinality)
			assignPropertyToSchemaType(parentPath, apiSchemaPartPath, property)
			
			if typeName == "object" {
				err = assignSchemaTypeToSchemaTypesMap(subMap, apiSchemaPartPath, varName, description, property)
			} else if property.IsEnum() {
				enumSchemaType := NewSchemaType(varName, description, property, nil)
				property.SetResolvedType(enumSchemaType)
				schemaTypes[apiSchemaPartPath] = enumSchemaType
			} else if property.IsConstant() {
				property.name = strings.ToUpper(property.name)
			}

			if err == nil {
				properties[apiSchemaPartPath] = property
			}
		}
		if err != nil {
			errList = append(errList, err)
		}
	}
}

func resolveReferences() {
	for _, property := range properties {
		if property.IsReferencing() {
			referencingPath := strings.Split(property.GetType(), ":")[1]
			referencedProp, isRefPropPresent := properties[referencingPath]
			if isRefPropPresent {
				property.Resolve(referencedProp)
			} else {
				err := openapi2beans_errors.NewError("ResolveReferences: Failed to find referenced property for %v\n", property)
				if err != nil {
					errList = append(errList, err)
				}
			}
		}
	}
}

func retrieveNestedProperties(subMap map[string]interface{}, yamlPath string) (err error) {
	var schemaPropertiesMap map[string]interface{}

	propertiesObj, isPropertyPresent := subMap[OPENAPI_YAML_KEYWORD_PROPERTIES]
	if isPropertyPresent {
		schemaPropertiesMap = propertiesObj.(map[string]interface{})
		retrieveSchemaComponentsFromMap(schemaPropertiesMap, yamlPath)
	}

	return err
}

func retrieveVarType(variableMap map[string]interface{}, apiSchemaPartPath string) (varType string, cardinality Cardinality, err error) {
	maxCardinality := 0
	varTypeObj, isTypePresent := variableMap[OPENAPI_YAML_KEYWORD_TYPE]
	refObj, isRefPresent := variableMap[OPENAPI_YAML_KEYWORD_REF]

	if isTypePresent {
		varType = varTypeObj.(string)
		if varType == "array" {
			varType, err = retrieveArrayType(variableMap, apiSchemaPartPath)
			maxCardinality = 128 * arrayDimensions
		} else {
			maxCardinality = 1
		}
		cardinality = Cardinality {min: 0, max: maxCardinality}
	} else if isRefPresent {
		varType = "$ref:" + refObj.(string)
	} else {
		err = openapi2beans_errors.NewError("RetrieveVarType: Failed to find required type for %v\n", apiSchemaPartPath)
	}

	return varType, cardinality, err
}

func retrieveArrayType(varMap map[string]interface{}, schemaPartPath string) (arrayType string, err error) {
	arrayDimensions += 1
	itemsObj, isItemsPresent := varMap[OPENAPI_YAML_KEYWORD_ITEMS]
	if isItemsPresent {
		itemsMap := itemsObj.(map[string]interface{})

		allOfObj, isAllOfPresent := itemsMap[OPENAPI_YAML_KEYWORD_ALLOF]
		if isAllOfPresent {
			allOfSlice := allOfObj.([]interface{})
			itemsMap = allOfSlice[0].(map[string]interface{})
		}
		arrayType, _, err = retrieveVarType(itemsMap, schemaPartPath)
		
	} else {
		err = openapi2beans_errors.NewError("RetrieveArrayType: Failed to find required items section for %v\n", schemaPartPath)
	}

	return arrayType, err
}

func retrieveDescription(subMap map[string]interface{}) (description string) {
	descriptionObj, isDescriptionPresent := subMap[OPENAPI_YAML_KEYWORD_DESCRIPTION]
	if isDescriptionPresent {
		description = descriptionObj.(string)
	}
	return description
}

func resolvePropertiesMinCardinalities(schemaTypeMap map[string]interface{}, schemaTypeProps map[string]*Property, schemaTypePath string) {
	requiredMapObj, isRequiredPresent := schemaTypeMap[OPENAPI_YAML_KEYWORD_REQUIRED]
	if isRequiredPresent {
		requiredMap := requiredMapObj.([]interface{})
		for _, required := range requiredMap {
			property, isPropertyNamePresent := schemaTypeProps[schemaTypePath + "/" + required.(string)]
			if isPropertyNamePresent {
				property.cardinality.min = 1
			}
		}
	}
}

func retrievePossibleValues(varMap map[string]interface{}) (possibleValues map[string]string) {
	possibleValues = make(map[string]string)
	enumObj, isEnumPresent := varMap[OPENAPI_YAML_KEYWORD_ENUM]
	if isEnumPresent {
		enums := enumObj.([]interface{})
		for _, enum := range enums {
			enumName := enum.(string)
			possibleValues[enumName] = enumName
		}
	}
	return possibleValues
}

func assignSchemaTypeToSchemaTypesMap(schemaTypeMap map[string]interface{}, apiSchemaPartPath string, varName string, description string, ownProperty *Property) error {
	resolvedType := NewSchemaType(varName, description, ownProperty, nil)

	ownProperty.SetResolvedType(resolvedType)

	schemaTypes[apiSchemaPartPath] = resolvedType

	err := retrieveNestedProperties(schemaTypeMap, apiSchemaPartPath)

	if err == nil {
		resolvePropertiesMinCardinalities(schemaTypeMap, resolvedType.properties, apiSchemaPartPath)
	}
	return err
}

func assignPropertyToSchemaType(parentPath string, apiSchemaPartPath string, property *Property) {
	localMap := schemaTypes
	schemaType, isPropertyPartOfSchemaType := localMap[parentPath]
	if isPropertyPartOfSchemaType {
		schemaType.properties[apiSchemaPartPath] = property
	}
}