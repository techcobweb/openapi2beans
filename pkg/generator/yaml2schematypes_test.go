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
	assert.Equal(t, "MyBeanName", schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetName(), "Wrong bean name read out of the yaml!")
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
	assert.Equal(t, "A simple example", schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetDescription(), "Wrong bean description read out of the yaml!")
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
	assert.NotEmpty(t, schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetProperties(), "Bean must have variable!")
	variable := schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetProperties()["#/components/schemas/MyBeanName/myStringVar"]
	assert.Equal(t, "myStringVar", variable.GetName(), "Wrong bean variable name read out of the yaml!")
	assert.Equal(t, "string", variable.GetType(), "Wrong bean variable type read out of the yaml!")
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
	variable := schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetProperties()["#/components/schemas/MyBeanName/myStringVar"]
	assert.Equal(t, "a test string", variable.GetDescription(), "Wrong bean variable description read out of the yaml!")
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
	assert.Equal(t, true, schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetProperties()["#/components/schemas/MyBeanName/myStringVar"].IsSetInConstructor(), "Wrong bean variable description read out of the yaml!")
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
	assert.Equal(t, false, schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetProperties()["#/components/schemas/MyBeanName/myStringVar"].IsSetInConstructor(), "Wrong bean variable description read out of the yaml!")
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
	assert.Equal(t, false, schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetProperties()["#/components/schemas/MyBeanName/myStringVar"].IsSetInConstructor(), "Wrong bean variable description read out of the yaml!")
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
	assert.Equal(t, true, schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetProperties()["#/components/schemas/MyBeanName/myStringVar"].IsSetInConstructor(), "Wrong bean variable description read out of the yaml!")
	assert.Equal(t, true, schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetProperties()["#/components/schemas/MyBeanName/myStringVar1"].IsSetInConstructor(), "Wrong bean variable description read out of the yaml!")
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
	assert.Equal(t, false, schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetProperties()["#/components/schemas/MyBeanName/myStringVar"].IsSetInConstructor(), "Wrong bean variable description read out of the yaml!")
	assert.Equal(t, false, schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetProperties()["#/components/schemas/MyBeanName/myStringVar1"].IsSetInConstructor(), "Wrong bean variable description read out of the yaml!")
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
	assert.Equal(t, false, schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetProperties()["#/components/schemas/MyBeanName/myStringVar"].IsSetInConstructor(), "Wrong bean variable isSetInConstructor read out of the yaml!")
	assert.Equal(t, true, schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetProperties()["#/components/schemas/MyBeanName/myStringVar1"].IsSetInConstructor(), "Wrong bean variable isSetInConstructor read out of the yaml!")
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
	assert.NotEmpty(t, schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetProperties(), "Bean must have variable!")
	assert.Equal(t, "mySecondStringVar", schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetProperties()["#/components/schemas/MyBeanName/mySecondStringVar"].GetName(), "Wrong bean variable name read out of the yaml!")
	assert.Equal(t, "string", schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetProperties()["#/components/schemas/MyBeanName/mySecondStringVar"].GetType(), "Wrong bean variable type read out of the yaml!")
	assert.Equal(t, "myStringVar", schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetProperties()["#/components/schemas/MyBeanName/myStringVar"].GetName(), "Wrong bean variable name read out of the yaml!")
	assert.Equal(t, "string", schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetProperties()["#/components/schemas/MyBeanName/myStringVar"].GetType(), "Wrong bean variable type read out of the yaml!")
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
	assert.NotEmpty(t, schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetProperties(), "Bean must have a variable!")
	assert.Equal(t, "myTestArray", schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetProperties()["#/components/schemas/MyBeanName/myTestArray"].GetName(), "Wrong bean variable name read out of the yaml!")
	assert.Equal(t, "string[]", schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetProperties()["#/components/schemas/MyBeanName/myTestArray"].GetType(), "Wrong bean variable type read out of the yaml!")
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
	assert.NotEmpty(t, schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetProperties(), "Bean must have a variable!")
	assert.Equal(t, "myTestArray", schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetProperties()["#/components/schemas/MyBeanName/myTestArray"].GetName(), "Wrong bean variable name read out of the yaml!")
	assert.Equal(t, "string[]", schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetProperties()["#/components/schemas/MyBeanName/myTestArray"].GetType(), "Wrong bean variable type read out of the yaml!")
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
	assert.NotEmpty(t, schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetProperties(), "Bean must have a variable!")
	assert.Equal(t, "nestedObject", schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetProperties()["#/components/schemas/MyBeanName/nestedObject"].GetName(), "Wrong bean variable name read out of the yaml!")
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
	assert.NotEmpty(t, schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetProperties(), "Bean must have a variable!")
	assert.Equal(t, "referencingObject", schemaTypes[SCHEMAS_PATH+"MyBeanName"].GetProperties()[SCHEMAS_PATH+"MyBeanName/referencingObject"].GetName(), "Wrong bean variable name read out of the yaml!")
}


