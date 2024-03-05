package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertJavaClassCorrectlyRelatesToSchemaType(t *testing.T, schemaType SchemaType, class JavaClass) {
	assert.Equal(t, schemaType.name, class.Name)
	schemaPath := "#/components/schemas/" + schemaType.name
	for _, dataMember := range class.DataMembers {
		comparisonSchemaProperty := schemaType.properties[schemaPath+"/"+dataMember.Name]
		expectedName := comparisonSchemaProperty.name
		assert.Equal(t, expectedName, dataMember.Name)
		expectedType := getExpectedType(comparisonSchemaProperty)
		assert.Equal(t, dataMember.MemberType, expectedType)
	}
}

func getExpectedType(schemaProp *Property) string {
	expectedType := ""
	if schemaProp.typeName == "string" {
		expectedType = "String"
	} else {
		expectedType = schemaProp.typeName
	}
	if schemaProp.cardinality.max > 1 {
		expectedType += "[]"
	}

	return expectedType
}

func TestTranslateSchemaTypesToJavaPackageReturnsPackageWithJavaClass(t *testing.T) {
	// Given...
	name := "MyBean"
	schemaType := NewSchemaType(name, "", nil, nil)
	schemaTypeMap := make(map[string]*SchemaType)
	schemaTypeMap["#/components/schemas/MyBean"] = schemaType

	// When...
	javaPackage := translateSchemaTypesToJavaPackage(schemaTypeMap, TARGET_JAVA_PACKAGE)

	// Then...
	assert.Equal(t, "MyBean", javaPackage.classes["MyBean"].Name)
}

func TestTranslateSchemaTypesToJavaPackageWithClassWithDataMember(t *testing.T) {
	// Given...
	propName1 := "MyRandomProperty"
	property := NewProperty(propName1, "#/components/schemas/MyBean/"+propName1, "", "string", nil, nil, Cardinality{min: 0, max: 1})
	properties := make(map[string]*Property)
	properties["#/components/schemas/MyBean/"+propName1] = property
	var schemaType *SchemaType
	schemaName := "MyBean"
	ownProp := NewProperty(schemaName, "#/components/schemas/MyBean", "", "object", nil, schemaType, Cardinality{min: 0, max: 1})
	schemaType = NewSchemaType(schemaName, "", ownProp, properties)
	schemaTypeMap := make(map[string]*SchemaType)
	schemaTypeMap["#/components/schemas/MyBean"] = schemaType

	// When...
	javaPackage := translateSchemaTypesToJavaPackage(schemaTypeMap, TARGET_JAVA_PACKAGE)

	// Then...
	assert.Equal(t, "MyBean", javaPackage.classes[schemaName].Name)
	assert.Equal(t, "MyRandomProperty", javaPackage.classes[schemaName].DataMembers[0].Name)
}

func TestTranslateSchemaTypesToJavaPackageWithClassWithMultipleDataMembers(t *testing.T) {
	// Given...
	propName1 := "MyRandomProperty1"
	property1 := NewProperty(propName1, "#/components/schemas/MyBean/"+propName1, "", "string", nil, nil, Cardinality{min: 0, max: 1})
	properties := make(map[string]*Property)
	properties["#/components/schemas/MyBean/"+propName1] = property1
	propName2 := "MyRandomProperty2"
	property2 := NewProperty(propName2, "#/components/schemas/MyBean/"+propName2, "", "string", nil, nil, Cardinality{min: 0, max: 1})
	properties["#/components/schemas/MyBean/"+propName2] = property2
	var schemaType *SchemaType
	schemaName := "MyBean"
	ownProp := NewProperty(schemaName, "#/components/schemas/MyBean", "", "object", nil, schemaType, Cardinality{min: 0, max: 1})
	schemaType = NewSchemaType(schemaName, "", ownProp, properties)
	schemaTypeMap := make(map[string]*SchemaType)
	schemaTypeMap["#/components/schemas/MyBean"] = schemaType

	// When...
	javaPackage := translateSchemaTypesToJavaPackage(schemaTypeMap, TARGET_JAVA_PACKAGE)

	// Then...
	assert.Equal(t, "MyBean", javaPackage.classes[schemaName].Name)
	assertJavaClassCorrectlyRelatesToSchemaType(t, *schemaType, *javaPackage.classes[schemaName])
}

func TestTranslateSchemaTypesToJavaPackageWithClassWithArrayDataMember(t *testing.T) {
	// Given...
	propName1 := "MyRandomProperty1"
	property1 := NewProperty(propName1, "#/components/schemas/MyBean/"+propName1, "", "string", nil, nil, Cardinality{min: 0, max: 100})
	properties := make(map[string]*Property)
	properties["#/components/schemas/MyBean/"+propName1] = property1
	var schemaType *SchemaType
	schemaName := "MyBean"
	ownProp := NewProperty(schemaName, "#/components/schemas/MyBean", "", "object", nil, schemaType, Cardinality{min: 0, max: 1})
	schemaType = NewSchemaType(schemaName, "", ownProp, properties)
	schemaTypeMap := make(map[string]*SchemaType)
	schemaTypeMap["#/components/schemas/MyBean"] = schemaType

	// When...
	javaPackage := translateSchemaTypesToJavaPackage(schemaTypeMap, TARGET_JAVA_PACKAGE)

	// Then...
	assert.Equal(t, "MyBean", javaPackage.classes[schemaName].Name)
	assertJavaClassCorrectlyRelatesToSchemaType(t, *schemaType, *javaPackage.classes[schemaName])
}

