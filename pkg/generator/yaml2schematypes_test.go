package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const SCHEMAS_PATH = "#/components/schemas/"

func TestGetSchemaTypesFromYamlReturns1BeanOK(t *testing.T) {
	// Given...
	apiYaml := `openapi: 3.0.3
components:
  schemas:
    MyBeanName:
      type: object
`
	// When...
	schemaTypes, err := getSchemaTypesFromYaml([]byte(apiYaml))

	// Then...
	assert.Nil(t, err)
	assert.Equal(t, 1, len(schemaTypes))
}

func TestGetSchemaTypesFromYamlReturnsBeanWithName(t *testing.T) {
	// Given...
	apiYaml := `openapi: 3.0.3
components:
  schemas:
    MyBeanName:
      type: object
`

	// When...
	schemaTypes, err := getSchemaTypesFromYaml([]byte(apiYaml))

	// Then...
	assert.Nil(t, err)
	schemaType, schemaTypeExists := schemaTypes[SCHEMAS_PATH+"MyBeanName"]
	assert.True(t, schemaTypeExists)
	assert.Equal(t, "MyBeanName", schemaType.GetName(), "Wrong bean name read out of the yaml!")
}

func TestGetSchemaTypesFromYamlParsesDescription(t *testing.T) {
	// Given...
	apiYaml := `openapi: 3.0.3
components:
  schemas:
    MyBeanName:
      type: object
      description: A simple example
`
	// When...
	schemaTypes, err := getSchemaTypesFromYaml([]byte(apiYaml))

	// Then...
	assert.Nil(t, err)
	schemaType, schemaTypeExists := schemaTypes[SCHEMAS_PATH+"MyBeanName"]
	assert.True(t, schemaTypeExists)
	assert.Equal(t, "A simple example", schemaType.GetDescription(), "Wrong bean description read out of the yaml!")
}

func TestGetSchemaTypesFromYamlParsesSingleStringVariable(t *testing.T) {
	// Given...
	apiYaml := `openapi: 3.0.3
components:
  schemas:
    MyBeanName:
      type: object
      description: A simple example
      properties:
        myStringVar:
          type: string
`
	// When...
	schemaTypes, err := getSchemaTypesFromYaml([]byte(apiYaml))

	// Then...
	assert.Nil(t, err)
	schemaType, schemaTypeExists := schemaTypes[SCHEMAS_PATH+"MyBeanName"]
	assert.True(t, schemaTypeExists)
	assert.NotEmpty(t, schemaType.GetProperties(), "Bean must have variable!")
	property, propertyExists := schemaType.GetProperties()["#/components/schemas/MyBeanName/myStringVar"]
	assert.True(t, propertyExists)
	assert.Equal(t, "myStringVar", property.GetName(), "Wrong bean variable name read out of the yaml!")
	assert.Equal(t, "string", property.GetType(), "Wrong bean variable type read out of the yaml!")
}

func TestGetSchemaTypesFromYamlParsesSingleStringVariableWithDescription(t *testing.T) {
	// Given...
	apiYaml := `openapi: 3.0.3
components:
  schemas:
    MyBeanName:
      type: object
      description: A simple example
      properties:
        myStringVar:
          type: string
          description: a test string
`
	// When...
	schemaTypes, err := getSchemaTypesFromYaml([]byte(apiYaml))

	// Then...
	assert.Nil(t, err)
	schemaType, schemaTypeExists := schemaTypes[SCHEMAS_PATH+"MyBeanName"]
	assert.True(t, schemaTypeExists)
	property, propertyExists := schemaType.GetProperties()["#/components/schemas/MyBeanName/myStringVar"]
	assert.True(t, propertyExists)
	assert.Equal(t, "a test string", property.GetDescription(), "Wrong bean variable description read out of the yaml!")
}

func TestGetSchemaTypesFromYamlParsesSingleStringVariableWithTrueRequiredField(t *testing.T) {
	// Given...
	apiYaml := `openapi: 3.0.3
components:
  schemas:
    MyBeanName:
      type: object
      description: A simple example
      properties:
        myStringVar:
          type: string
          description: a test string
          required: true
`
	// When...
	schemaTypes, err := getSchemaTypesFromYaml([]byte(apiYaml))

	// Then...
	assert.Nil(t, err)
	schemaType, schemaTypeExists := schemaTypes[SCHEMAS_PATH+"MyBeanName"]
	assert.True(t, schemaTypeExists)
	assert.NotEmpty(t, schemaType.GetProperties(), "Bean must have variable!")
	property, propertyExists := schemaType.GetProperties()["#/components/schemas/MyBeanName/myStringVar"]
	assert.True(t, propertyExists)
	assert.Equal(t, true, property.IsSetInConstructor(), "Wrong bean variable required status read out of the yaml!")
}

func TestGetSchemaTypesFromYamlParsesSingleStringVariableWithFalseRequiredField(t *testing.T) {
	// Given...
	apiYaml := `openapi: 3.0.3
components:
  schemas:
    MyBeanName:
      type: object
      description: A simple example
      properties:
        myStringVar:
          type: string
          description: a test string
          required: false
`
	// When...
	schemaTypes, err := getSchemaTypesFromYaml([]byte(apiYaml))

	// Then...
	assert.Nil(t, err)
	schemaType, schemaTypeExists := schemaTypes[SCHEMAS_PATH+"MyBeanName"]
	assert.True(t, schemaTypeExists)
	assert.NotEmpty(t, schemaType.GetProperties(), "Bean must have variable!")
	property, propertyExists := schemaType.GetProperties()["#/components/schemas/MyBeanName/myStringVar"]
	assert.True(t, propertyExists)
	assert.Equal(t, false, property.IsSetInConstructor(), "Wrong bean variable required status read out of the yaml!")
}

func TestGetSchemaTypesFromYamlParsesSingleStringVariableWithNoRequiredFieldReturnsFalse(t *testing.T) {
	// Given...
	apiYaml := `openapi: 3.0.3
components:
  schemas:
    MyBeanName:
      type: object
      description: A simple example
      properties:
        myStringVar:
          type: string
          description: a test string
`
	// When...
	schemaTypes, err := getSchemaTypesFromYaml([]byte(apiYaml))

	// Then...
	assert.Nil(t, err)
	schemaType, schemaTypeExists := schemaTypes[SCHEMAS_PATH+"MyBeanName"]
	assert.True(t, schemaTypeExists)
	assert.NotEmpty(t, schemaType.GetProperties(), "Bean must have variable!")
	property, propertyExists := schemaType.GetProperties()["#/components/schemas/MyBeanName/myStringVar"]
	assert.True(t, propertyExists)
	assert.Equal(t, false, property.IsSetInConstructor(), "Wrong bean variable required status read out of the yaml!")
}

func TestGetSchemaTypesFromYamlParsesMultipleStringVariableWithTrueRequiredFields(t *testing.T) {
	// Given...
	apiYaml := `openapi: 3.0.3
components:
  schemas:
    MyBeanName:
      type: object
      description: A simple example
      properties:
        myStringVar:
          type: string
          description: a test string
          required: true
        myStringVar1:
          type: string
          description: a test string
          required: true
`
	// When...
	schemaTypes, err := getSchemaTypesFromYaml([]byte(apiYaml))

	// Then...
	assert.Nil(t, err)
	schemaType, schemaTypeExists := schemaTypes[SCHEMAS_PATH+"MyBeanName"]
	assert.True(t, schemaTypeExists)
	assert.NotEmpty(t, schemaType.GetProperties(), "Bean must have variable!")
	property1, propertyExists := schemaType.GetProperties()["#/components/schemas/MyBeanName/myStringVar"]
	assert.True(t, propertyExists)
	property2, propertyExists := schemaType.GetProperties()["#/components/schemas/MyBeanName/myStringVar1"]
	assert.True(t, propertyExists)
	assert.Equal(t, true, property1.IsSetInConstructor(), "Wrong bean variable required status read out of the yaml!")
	assert.Equal(t, true, property2.IsSetInConstructor(), "Wrong bean variable required status read out of the yaml!")
}

func TestGetSchemaTypesFromYamlParsesMultipleStringVariablesWithFalseRequiredFields(t *testing.T) {
	// Given...
	apiYaml := `openapi: 3.0.3
components:
  schemas:
    MyBeanName:
      type: object
      description: A simple example
      properties:
        myStringVar:
          type: string
          description: a test string
          required: false
        myStringVar1:
          type: string
          description: a test string in addition to the other
          required: false
`
	// When...
	schemaTypes, err := getSchemaTypesFromYaml([]byte(apiYaml))

	// Then...
	assert.Nil(t, err)
	schemaType, schemaTypeExists := schemaTypes[SCHEMAS_PATH+"MyBeanName"]
	assert.True(t, schemaTypeExists)
	assert.NotEmpty(t, schemaType.GetProperties(), "Bean must have variable!")
	property1, propertyExists := schemaType.GetProperties()["#/components/schemas/MyBeanName/myStringVar"]
	assert.True(t, propertyExists)
	property2, propertyExists := schemaType.GetProperties()["#/components/schemas/MyBeanName/myStringVar1"]
	assert.True(t, propertyExists)
	assert.Equal(t, false, property1.IsSetInConstructor(), "Wrong bean variable required status read out of the yaml!")
	assert.Equal(t, false, property2.IsSetInConstructor(), "Wrong bean variable required status read out of the yaml!")
}

func TestGetSchemaTypesFromYamlParsesMultipleStringVariablesWithMixedRequiredFields(t *testing.T) {
	// Given...
	apiYaml := `openapi: 3.0.3
components:
  schemas:
    MyBeanName:
      type: object
      description: A simple example
      properties:
        myStringVar:
          type: string
          description: a test string
          required: false
        myStringVar1:
          type: string
          description: a test string in addition to the other
          required: true
`
	// When...
	schemaTypes, err := getSchemaTypesFromYaml([]byte(apiYaml))

	// Then...
	assert.Nil(t, err)
	schemaType, schemaTypeExists := schemaTypes[SCHEMAS_PATH+"MyBeanName"]
	assert.True(t, schemaTypeExists)
	assert.NotEmpty(t, schemaType.GetProperties(), "Bean must have variable!")
	property1, propertyExists := schemaType.GetProperties()["#/components/schemas/MyBeanName/myStringVar"]
	assert.True(t, propertyExists)
	property2, propertyExists := schemaType.GetProperties()["#/components/schemas/MyBeanName/myStringVar1"]
	assert.True(t, propertyExists)
	assert.Equal(t, false, property1.IsSetInConstructor(), "Wrong bean variable required status read out of the yaml!")
	assert.Equal(t, true, property2.IsSetInConstructor(), "Wrong bean variable required status read out of the yaml!")
}

func TestGetSchemaTypesFromYamlParsesMultipleStringVariables(t *testing.T) {
	// Given...
	apiYaml := `openapi: 3.0.3
components:
  schemas:
    MyBeanName:
      type: object
      description: A simple example
      properties:
        myStringVar:
          type: string
          description: a test string
        mySecondStringVar:
          type: string
          description: a second test string
`
	// When...
	schemaTypes, err := getSchemaTypesFromYaml([]byte(apiYaml))

	// Then...
	assert.Nil(t, err)
	schemaType, schemaTypeExists := schemaTypes[SCHEMAS_PATH+"MyBeanName"]
	assert.True(t, schemaTypeExists)
	assert.NotEmpty(t, schemaType.GetProperties(), "Bean must have variable!")
	property1, propertyExists := schemaType.GetProperties()["#/components/schemas/MyBeanName/myStringVar"]
	assert.True(t, propertyExists)
	property2, propertyExists := schemaType.GetProperties()["#/components/schemas/MyBeanName/mySecondStringVar"]
	assert.True(t, propertyExists)
	assert.Equal(t, "myStringVar", property1.GetName(), "Wrong bean variable name read out of the yaml!")
	assert.Equal(t, "string", property1.GetType(), "Wrong bean variable type read out of the yaml!")
	assert.Equal(t, "mySecondStringVar", property2.GetName(), "Wrong bean variable name read out of the yaml!")
	assert.Equal(t, "string", property2.GetType(), "Wrong bean variable type read out of the yaml!")
}

func TestGetSchemaTypesFromYamlParsesObjectWithArray(t *testing.T) {
	// Given...
	apiYaml := `openapi: 3.0.3
components:
  schemas:
    MyBeanName:
      type: object
      properties:
        myTestArray:
          type: array
          items:
            type: string
`
	// When...
	schemaTypes, err := getSchemaTypesFromYaml([]byte(apiYaml))

	// Then...
	assert.Nil(t, err)
	schemaType, schemaTypeExists := schemaTypes[SCHEMAS_PATH+"MyBeanName"]
	assert.True(t, schemaTypeExists)
	assert.NotEmpty(t, schemaType.GetProperties(), "Bean must have variable!")
	property1, propertyExists := schemaType.GetProperties()["#/components/schemas/MyBeanName/myTestArray"]
	assert.True(t, propertyExists)
	assert.Equal(t, "myTestArray", property1.GetName(), "Wrong bean variable name read out of the yaml!")
	assert.Equal(t, "string", property1.GetType(), "Wrong bean variable type read out of the yaml!")
	assert.Equal(t, true, property1.IsCollection(), "Wrong bean variable cardinality read out of the yaml!")
}

func TestGetSchemaTypesFromYamlParsesObjectWithArrayContainingAllOfPart(t *testing.T) {
	// Given...
	apiYaml := `openapi: 3.0.3
components:
  schemas:
    MyBeanName:
      type: object
      properties:
        myTestArray:
          type: array
          items:
            allOf:
            - type: string
`
	// When...
	schemaTypes, err := getSchemaTypesFromYaml([]byte(apiYaml))

	// Then...
	assert.Nil(t, err)
	schemaType, schemaTypeExists := schemaTypes[SCHEMAS_PATH+"MyBeanName"]
	assert.True(t, schemaTypeExists)
	assert.NotEmpty(t, schemaType.GetProperties(), "Bean must have variable!")
	property1, propertyExists := schemaType.GetProperties()["#/components/schemas/MyBeanName/myTestArray"]
	assert.True(t, propertyExists)
	assert.Equal(t, "myTestArray", property1.GetName(), "Wrong bean variable name read out of the yaml!")
	assert.Equal(t, "string", property1.GetType(), "Wrong bean variable type read out of the yaml!")
	assert.Equal(t, true, property1.IsCollection(), "Wrong bean variable cardinality read out of the yaml!")
}

func TestGetSchemaTypesFromYamlParsesObjectWithArrayContainingArray(t *testing.T) {
	// Given...
	apiYaml := `openapi: 3.0.3
components:
  schemas:
    MyBeanName:
      type: object
      properties:
        myTestArray:
          type: array
          items:
            type: array
            items:
              type: string
`
	// When...
	schemaTypes, err := getSchemaTypesFromYaml([]byte(apiYaml))

	// Then...
	assert.Nil(t, err)
	schemaType, schemaTypeExists := schemaTypes[SCHEMAS_PATH+"MyBeanName"]
	assert.True(t, schemaTypeExists)
	assert.NotEmpty(t, schemaType.GetProperties(), "Bean must have variable!")
	property1, propertyExists := schemaType.GetProperties()["#/components/schemas/MyBeanName/myTestArray"]
	assert.True(t, propertyExists)
	assert.Equal(t, "myTestArray", property1.GetName(), "Wrong bean variable name read out of the yaml!")
	assert.Equal(t, "string", property1.GetType(), "Wrong bean variable type read out of the yaml!")
	assert.Equal(t, true, property1.IsCollection(), "Wrong bean variable cardinality read out of the yaml!")
	assert.Equal(t, 2, property1.cardinality.GetDimensions(), "Wrong array dimension read out of the yaml!")
}

func TestGetSchemaTypesFromYamlParsesNestedObjects(t *testing.T) {
	// Given..
	apiYaml := `openapi: 3.0.3
components:
  schemas:
    MyBeanName:
      type: object
      properties:
        nestedObject:
          type: object
          properties:
            randomString:
              type: string
`

	// When...
	schemaTypes, err := getSchemaTypesFromYaml([]byte(apiYaml))

	// Then...
	assert.Nil(t, err)
	schemaType, schemaTypeExists := schemaTypes[SCHEMAS_PATH+"MyBeanName"]
	assert.True(t, schemaTypeExists)
	assert.NotEmpty(t, schemaType.GetProperties(), "Bean must have variable!")
	property1, propertyExists := schemaType.GetProperties()["#/components/schemas/MyBeanName/nestedObject"]
	assert.True(t, propertyExists)
	property2, propertyExists := property1.resolvedType.GetProperties()["#/components/schemas/MyBeanName/nestedObject/randomString"]
	assert.True(t, propertyExists)
	assert.Equal(t, "nestedObject", property1.GetName(), "Wrong bean variable name read out of the yaml!")
	assert.Equal(t, "randomString", property2.GetName(), "Wrong bean variable name read out of the yaml!")
}

func TestGetSchemaTypesFromYamlParsesReferenceToObject(t *testing.T) {
	// Given..
	apiYaml := `openapi: 3.0.3
components:
  schemas:
    MyBeanName:
      type: object
      properties:
        referencingObject:
          $ref: '#/components/schemas/ReferencedObject'
    ReferencedObject:
      type: object
      properties:
        randomString:
          type: string
`

	// When...
	schemaTypes, err := getSchemaTypesFromYaml([]byte(apiYaml))

	// Then...
	assert.Nil(t, err)
	schemaType, schemaTypeExists := schemaTypes[SCHEMAS_PATH+"MyBeanName"]
	assert.True(t, schemaTypeExists)
	assert.NotEmpty(t, schemaType.GetProperties(), "Bean must have variable!")
	property1, propertyExists := schemaType.GetProperties()["#/components/schemas/MyBeanName/referencingObject"]
	assert.True(t, propertyExists)
	assert.Equal(t, "referencingObject", property1.GetName(), "Wrong bean variable name read out of the yaml!")
	assert.Equal(t, "ReferencedObject", property1.GetType())
}

func TestGetSchemaTypesFromYamlParsesObjectWithArrayContainingTypeRefToObject(t *testing.T) {
	// Given...
	apiYaml := `openapi: 3.0.3
components:
  schemas:
    MyBeanName:
      type: object
      properties:
        myTestArray:
          type: array
          items:
            $ref: '#/components/schemas/ReferencedObject'
    ReferencedObject:
      type: object
      properties:
        randomString:
          type: string
`
	// When...
	schemaTypes, err := getSchemaTypesFromYaml([]byte(apiYaml))

	// Then...
	assert.Nil(t, err)
	schemaType, schemaTypeExists := schemaTypes[SCHEMAS_PATH+"MyBeanName"]
	assert.True(t, schemaTypeExists)
	assert.NotEmpty(t, schemaType.GetProperties(), "Bean must have variable!")
	property1, propertyExists := schemaType.GetProperties()["#/components/schemas/MyBeanName/myTestArray"]
	assert.True(t, propertyExists)
	assert.Equal(t, "myTestArray", property1.GetName(), "Wrong bean variable name read out of the yaml!")
	assert.Equal(t, "ReferencedObject", property1.GetType(), "Wrong bean variable type read out of the yaml!")
	assert.Equal(t, true, property1.IsCollection(), "Wrong bean variable cardinality read out of the yaml!")
}

func TestGetSchemaTypesFromYamlParsesObjectWithArrayContainingAllOfRefToObject(t *testing.T) {
	// Given...
	apiYaml := `openapi: 3.0.3
components:
  schemas:
    MyBeanName:
      type: object
      properties:
        myTestArray:
          type: array
          items:
            allOf:
            - $ref: '#/components/schemas/ReferencedObject'
    ReferencedObject:
      type: object
      properties:
        randomString:
          type: string
`
	// When...
	schemaTypes, err := getSchemaTypesFromYaml([]byte(apiYaml))

	// Then...
	assert.Nil(t, err)
	schemaType, schemaTypeExists := schemaTypes[SCHEMAS_PATH+"MyBeanName"]
	assert.True(t, schemaTypeExists)
	assert.NotEmpty(t, schemaType.GetProperties(), "Bean must have variable!")
	property1, propertyExists := schemaType.GetProperties()["#/components/schemas/MyBeanName/myTestArray"]
	assert.True(t, propertyExists)
	assert.Equal(t, "myTestArray", property1.GetName(), "Wrong bean variable name read out of the yaml!")
	assert.Equal(t, "ReferencedObject", property1.GetType(), "Wrong bean variable type read out of the yaml!")
	assert.Equal(t, true, property1.IsCollection(), "Wrong bean variable cardinality read out of the yaml!")
}

func TestGetSchemaTypesFromYamlParsesEnum(t *testing.T) {
	// Given..
	apiYaml := `openapi: 3.0.3
components:
  schemas:
    MyBeanName:
      type: object
      properties:
        MyEnum:
          type: string
          enum: [randValue1, randValue2]
`

	// When...
	schemaTypes, err := getSchemaTypesFromYaml([]byte(apiYaml))
	schemaPath := SCHEMAS_PATH+"MyBeanName"
	propertyPath := schemaPath + "/MyEnum"

	// Then...
	assert.Nil(t, err)
	schemaType, schemaTypeExists := schemaTypes[schemaPath]
	assert.True(t, schemaTypeExists)
	assert.NotEmpty(t, schemaType.GetProperties(), "Bean must have variable!")
	property1, propertyExists := schemaType.GetProperties()[propertyPath]
	assert.True(t, propertyExists)
	assert.Equal(t, true, property1.IsEnum())
	posValue1, posValueExists := property1.GetPossibleValues()["randValue1"]
	assert.True(t, posValueExists)
	assert.Equal(t, "randValue1", posValue1)
	posValue2, posValueExists := property1.GetPossibleValues()["randValue2"]
	assert.True(t, posValueExists)
	assert.Equal(t, "randValue2", posValue2)
	enumSchemaType, enumSchemaTypeExists := schemaTypes[propertyPath]
	assert.Equal(t, true, enumSchemaTypeExists)
	assert.Equal(t, "MyEnum", enumSchemaType.name)
	assert.Equal(t, "MyEnum", enumSchemaType.ownProperty.name)
}

func TestGetSchemaTypesFromYamlParsesEnumAsConstant(t *testing.T) {
	// Given..
	apiYaml := `openapi: 3.0.3
components:
  schemas:
    MyBeanName:
      type: object
      properties:
        MyConstant:
          type: string
          enum: [randValue1]
`
	
		// When...
		schemaTypes, err := getSchemaTypesFromYaml([]byte(apiYaml))
		schemaPath := SCHEMAS_PATH+"MyBeanName"
		propertyPath := schemaPath + "/MyConstant"
	
		// Then...
		assert.Nil(t, err)
		schemaType, schemaTypeExists := schemaTypes[schemaPath]
		assert.True(t, schemaTypeExists)
		assert.NotEmpty(t, schemaType.GetProperties(), "Bean must have variable!")
		property1, propertyExists := schemaType.GetProperties()[propertyPath]
		assert.True(t, propertyExists)
		assert.Equal(t, true, property1.IsConstant())
		posValue, posValueExists := property1.GetPossibleValues()["randValue1"]
		assert.True(t, posValueExists)
		assert.Equal(t, "randValue1", posValue)
}