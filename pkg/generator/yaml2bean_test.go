package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBeanFromYamlReturns1BeanOK(t *testing.T) {
	// Given...
	apiYaml := `openapi: 3.0.3
components:
  schemas:
    MyBeanName:
      type: object
`
	// When...
	beans, err := getBeansFromYaml([]byte(apiYaml), TARGET_JAVA_PACKAGE)

	// Then...
	assert.Nil(t, err)
	assert.Equal(t, 1, len(beans))
}

func TestGetBeanFromYamlReturnsBeanWithName(t *testing.T) {
	// Given...
	apiYaml := `openapi: 3.0.3
components:
  schemas:
    MyBeanName:
      type: object
`

	// When...
	beans, err := getBeansFromYaml([]byte(apiYaml), TARGET_JAVA_PACKAGE)

	// Then...
	assert.Nil(t, err)
	assert.Equal(t, "MyBeanName", beans[0].object.varName, "Wrong bean name read out of the yaml!")
}

func TestGetBeanFromYamlParsesDescription(t *testing.T) {
	// Given...
	apiYaml := `openapi: 3.0.3
components:
  schemas:
    MyBeanName:
      type: object
      description: A simple example
`
	// When...
	beans, err := getBeansFromYaml([]byte(apiYaml), TARGET_JAVA_PACKAGE)

	// Then...
	assert.Nil(t, err)
	assert.Equal(t, "A simple example", beans[0].object.description, "Wrong bean description read out of the yaml!")
}

func TestGetBeanFromYamlParsesSingleStringVariable(t *testing.T) {
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
	beans, err := getBeansFromYaml([]byte(apiYaml), TARGET_JAVA_PACKAGE)


	// Then...
	assert.Nil(t, err)
	assert.NotEmpty(t, beans[0].object.variables, "Bean must have variable!")
	variable := beans[0].object.variables["#/components/schemas/MyBeanName/myStringVar"]
	assert.Equal(t, "myStringVar", variable.GetName(), "Wrong bean variable name read out of the yaml!")
	assert.Equal(t, "String", variable.GetType(), "Wrong bean variable type read out of the yaml!")
}

func TestGetBeanFromYamlParsesSingleStringVariableWithDescription(t *testing.T) {
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
	beans, err := getBeansFromYaml([]byte(apiYaml), TARGET_JAVA_PACKAGE)

	// Then...
	assert.Nil(t, err)
	variable := beans[0].object.variables["#/components/schemas/MyBeanName/myStringVar"]
	assert.Equal(t, "a test string", variable.GetDescription(), "Wrong bean variable description read out of the yaml!")
}

func TestGetBeanFromYamlParsesSingleStringVariableWithTrueRequiredField(t *testing.T) {
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
	beans, err := getBeansFromYaml([]byte(apiYaml), TARGET_JAVA_PACKAGE)

	// Then...
	assert.Nil(t, err)
	assert.Equal(t, true, beans[0].object.variables["#/components/schemas/MyBeanName/myStringVar"].(Variable).IsSetInConstructor(), "Wrong bean variable description read out of the yaml!")
}

func TestGetBeanFromYamlParsesSingleStringVariableWithFalseRequiredField(t *testing.T) {
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
	beans, err := getBeansFromYaml([]byte(apiYaml), TARGET_JAVA_PACKAGE)

	// Then...
	assert.Nil(t, err)
	assert.Equal(t, false, beans[0].object.variables["#/components/schemas/MyBeanName/myStringVar"].(Variable).IsSetInConstructor(), "Wrong bean variable description read out of the yaml!")
}

func TestGetBeanFromYamlParsesSingleStringVariableWithNoRequiredFieldReturnsFalse(t *testing.T) {
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
	beans, err := getBeansFromYaml([]byte(apiYaml), TARGET_JAVA_PACKAGE)

	// Then...
	assert.Nil(t, err)
	assert.Equal(t, false, beans[0].object.variables["#/components/schemas/MyBeanName/myStringVar"].(Variable).IsSetInConstructor(), "Wrong bean variable description read out of the yaml!")
}

func TestGetBeanFromYamlParsesMultipleStringVariableWithTrueRequiredFields(t *testing.T) {
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
	beans, err := getBeansFromYaml([]byte(apiYaml), TARGET_JAVA_PACKAGE)

	// Then...
	assert.Nil(t, err)
	assert.Equal(t, true, beans[0].object.variables["#/components/schemas/MyBeanName/myStringVar"].(Variable).isSetInConstructor, "Wrong bean variable description read out of the yaml!")
	assert.Equal(t, true, beans[0].object.variables["#/components/schemas/MyBeanName/myStringVar1"].(Variable).isSetInConstructor, "Wrong bean variable description read out of the yaml!")
}

func TestGetBeanFromYamlParsesMultipleStringVariablesWithFalseRequiredFields(t *testing.T) {
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
	beans, err := getBeansFromYaml([]byte(apiYaml), TARGET_JAVA_PACKAGE)

	// Then...
	assert.Nil(t, err)
	assert.Equal(t, false, beans[0].object.variables["#/components/schemas/MyBeanName/myStringVar"].(Variable).isSetInConstructor, "Wrong bean variable description read out of the yaml!")
	assert.Equal(t, false, beans[0].object.variables["#/components/schemas/MyBeanName/myStringVar1"].(Variable).isSetInConstructor, "Wrong bean variable description read out of the yaml!")
}

func TestGetBeanFromYamlParsesMultipleStringVariablesWithMixedRequiredFields(t *testing.T) {
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
	beans, err := getBeansFromYaml([]byte(apiYaml), TARGET_JAVA_PACKAGE)

	// Then...
	assert.Nil(t, err)
	assert.Equal(t, false, beans[0].object.variables["#/components/schemas/MyBeanName/myStringVar"].(Variable).isSetInConstructor, "Wrong bean variable isSetInConstructor read out of the yaml!")
	assert.Equal(t, true, beans[0].object.variables["#/components/schemas/MyBeanName/myStringVar1"].(Variable).isSetInConstructor, "Wrong bean variable isSetInConstructor read out of the yaml!")
}

func TestGetBeanFromYamlParsesMultipleStringVariables(t *testing.T) {
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
	beans, err := getBeansFromYaml([]byte(apiYaml), TARGET_JAVA_PACKAGE)

	// Then...
	assert.Nil(t, err)
	assert.NotEmpty(t, beans[0].object.variables, "Bean must have variable!")
	assert.Equal(t, "mySecondStringVar", beans[0].object.variables["#/components/schemas/MyBeanName/mySecondStringVar"].(Variable).varName, "Wrong bean variable name read out of the yaml!")
	assert.Equal(t, "String", beans[0].object.variables["#/components/schemas/MyBeanName/mySecondStringVar"].(Variable).varTypeName, "Wrong bean variable type read out of the yaml!")
	assert.Equal(t, "myStringVar", beans[0].object.variables["#/components/schemas/MyBeanName/myStringVar"].(Variable).varName, "Wrong bean variable name read out of the yaml!")
	assert.Equal(t, "String", beans[0].object.variables["#/components/schemas/MyBeanName/myStringVar"].(Variable).varTypeName, "Wrong bean variable type read out of the yaml!")
}

func TestGetBeanFromYamlParsesObjectWithArray(t *testing.T) {
	// Given...
	apiYaml := `openapi: 3.0.3
components:
  schemas:
    myBeanName:
      type: object
      properties:
        myTestArray:
          type: array
          items:
            type: string
`
	// When...
	beans, err := getBeansFromYaml([]byte(apiYaml), TARGET_JAVA_PACKAGE)

	// Then...
	assert.Nil(t, err)
	assert.NotEmpty(t, beans[0].object.variables, "Bean must have a variable!")
	assert.Equal(t, "myTestArray", beans[0].object.variables["#/components/schemas/myBeanName/myTestArray"].(Variable).varName, "Wrong bean variable name read out of the yaml!")
	assert.Equal(t, "String[]", beans[0].object.variables["#/components/schemas/myBeanName/myTestArray"].(Variable).varTypeName, "Wrong bean variable type read out of the yaml!")
}

func TestGetBeanFromYamlParsesObjectWithArrayContainingAllOfPart(t *testing.T) {
	// Given...
	apiYaml := `openapi: 3.0.3
components:
  schemas:
    myBeanName:
      type: object
      properties:
        myTestArray:
          type: array
          items:
            allOf:
            - type: string
`
	// When...
	beans, err := getBeansFromYaml([]byte(apiYaml), TARGET_JAVA_PACKAGE)

	// Then...
	assert.Nil(t, err)
	assert.NotEmpty(t, beans[0].object.variables, "Bean must have a variable!")
	assert.Equal(t, "myTestArray", beans[0].object.variables["#/components/schemas/myBeanName/myTestArray"].(Variable).varName, "Wrong bean variable name read out of the yaml!")
	assert.Equal(t, "String[]", beans[0].object.variables["#/components/schemas/myBeanName/myTestArray"].(Variable).varTypeName, "Wrong bean variable type read out of the yaml!")
}

func TestGetBeanFromYamlParsesNestedObjects(t *testing.T) {
	// Given..
	apiYaml := `openapi: 3.0.3
components:
  schemas:
    myBeanName:
      type: object
      properties:
        nestedObject:
          type: object
          properties:
            randomString:
              type: string
`

	// When...
	beans, err := getBeansFromYaml([]byte(apiYaml), TARGET_JAVA_PACKAGE)

	// Then...
	assert.Nil(t, err)
	assert.NotEmpty(t, beans[0].object.variables, "Bean must have a variable!")
	assert.Equal(t, "nestedObject", beans[0].object.variables["#/components/schemas/myBeanName/nestedObject"].(Object).varName, "Wrong bean variable name read out of the yaml!")
}

func TestGetBeanFromYamlParsesReferenceToObject(t *testing.T) {
	// Given..
	apiYaml := `openapi: 3.0.3
components:
  schemas:
    myBeanName:
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
	beans, err := getBeansFromYaml([]byte(apiYaml), TARGET_JAVA_PACKAGE)

	// Then...
	assert.Nil(t, err)
	assert.NotEmpty(t, beans[0].object.variables, "Bean must have a variable!")
	// CURRENTLY REF IS NOT BEING UNMARSHALLED.
	assert.Equal(t, "referencingObject", beans[0].object.variables["#/components/schemas/myBeanName/referencingObject"].(Object).varName, "Wrong bean variable name read out of the yaml!")
}

