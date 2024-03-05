package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertJavaClassCorrectlyRelatesToSchemaType(t *testing.T, schemaType *SchemaType, class *JavaClass) {
	assert.Equal(t, schemaType.name, class.Name)
	schemaPath := "#/components/schemas/" + schemaType.name
	propertiesVisited := 0

	for _, dataMember := range class.DataMembers {
		assert.NotNil(t, schemaType.properties[schemaPath+"/"+dataMember.Name])
		comparisonSchemaProperty := schemaType.properties[schemaPath+"/"+dataMember.Name]
		propertiesVisited += 1
		expectedName := comparisonSchemaProperty.name
		assert.Equal(t, expectedName, dataMember.Name)
		expectedType := getExpectedType(comparisonSchemaProperty)
		assert.Equal(t, dataMember.MemberType, expectedType)
	}
	assert.Equal(t, len(schemaType.properties), propertiesVisited)

	requiredPropertiesVisited := 0
	for _, requiredMember := range class.RequiredMembers {
		assert.NotNil(t, schemaType.properties[schemaPath+"/"+requiredMember.DataMember.Name])
		comparisonSchemaProperty := schemaType.properties[schemaPath+"/"+requiredMember.DataMember.Name]
		requiredPropertiesVisited += 1
		expectedName := comparisonSchemaProperty.name
		assert.Equal(t, expectedName, requiredMember.DataMember.Name)
		expectedType := getExpectedType(comparisonSchemaProperty)
		assert.Equal(t, requiredMember.DataMember.MemberType, expectedType)
		assert.True(t, comparisonSchemaProperty.IsSetInConstructor())
	}
	assert.Equal(t, numberOfRequiredProperties(schemaType.properties), requiredPropertiesVisited)
}

func numberOfRequiredProperties(properties map[string]*Property) int {
	counter := 0
	for _, prop := range properties {
		if prop.IsSetInConstructor() {
			counter += 1
		}
	}
	return counter
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

func assertJavaEnumRelatesToSchemaType(t *testing.T, schemaType *SchemaType, javaEnum JavaEnum) {
	assert.Equal(t, schemaType.name, javaEnum.Name)
	assert.Equal(t, schemaType.description, javaEnum.Description)
	for _, enumValue := range javaEnum.EnumValues {
		assert.NotNil(t, schemaType.ownProperty.possibleValues[enumValue])
	}
}

func TestTranslateSchemaTypesToJavaPackageReturnsPackageWithJavaClass(t *testing.T) {
	// Given...
	var schemaType *SchemaType
	name := "MyBean"
	ownProp := NewProperty(name, "#/components/schemas/MyBean", "", "object", nil, schemaType, Cardinality{min: 0, max: 1})
	schemaType = NewSchemaType(name, "", ownProp, nil)
	schemaTypeMap := make(map[string]*SchemaType)
	schemaTypeMap["#/components/schemas/MyBean"] = schemaType

	// When...
	javaPackage := translateSchemaTypesToJavaPackage(schemaTypeMap, TARGET_JAVA_PACKAGE)

	// Then...
	assert.Equal(t, "MyBean", javaPackage.Classes["MyBean"].Name)
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
	assert.Equal(t, "MyBean", javaPackage.Classes[schemaName].Name)
	assert.Equal(t, "MyRandomProperty", javaPackage.Classes[schemaName].DataMembers[0].Name)
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
	assert.Equal(t, "MyBean", javaPackage.Classes[schemaName].Name)
	assertJavaClassCorrectlyRelatesToSchemaType(t, schemaType, javaPackage.Classes[schemaName])
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
	assert.Equal(t, "MyBean", javaPackage.Classes[schemaName].Name)
	assertJavaClassCorrectlyRelatesToSchemaType(t, schemaType, javaPackage.Classes[schemaName])
}

func TestTranslateSchemaTypesToJavaPackageWithClassWithMixedArrayAndPrimitiveDataMembers(t *testing.T) {
	// Given...
	propName1 := "MyRandomProperty1"
	property1 := NewProperty(propName1, "#/components/schemas/MyBean/"+propName1, "", "string", nil, nil, Cardinality{min: 0, max: 100})
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
	assert.Equal(t, "MyBean", javaPackage.Classes[schemaName].Name)
	assertJavaClassCorrectlyRelatesToSchemaType(t, schemaType, javaPackage.Classes[schemaName])
}

func TestTranslateSchemaTypesToJavaPackageWithClassWithReferenceToOtherClass(t *testing.T) {
	// Given...
	schemaTypeMap := make(map[string]*SchemaType)
	var referencedSchemaType *SchemaType
	referencedSchemaName := "MyReferencedBean"
	referencedOwnProp := NewProperty(referencedSchemaName, "#/components/schemas/MyReferencedBean", "", "object", nil, referencedSchemaType, Cardinality{min: 0, max: 1})
	referencedSchemaType = NewSchemaType(referencedSchemaName, "", referencedOwnProp, nil)
	schemaTypeMap["#/components/schemas/MyReferencedBean"] = referencedSchemaType
	propName1 := "MyRandomProperty1"
	property1 := NewProperty(propName1, "#/components/schemas/MyBean/"+propName1, "", "object", nil, referencedSchemaType, Cardinality{min: 0, max: 1})
	properties := make(map[string]*Property)
	properties["#/components/schemas/MyBean/"+propName1] = property1
	var schemaType *SchemaType
	schemaName := "MyBean"
	ownProp := NewProperty(schemaName, "#/components/schemas/MyBean", "", "object", nil, schemaType, Cardinality{min: 0, max: 1})
	schemaType = NewSchemaType(schemaName, "", ownProp, properties)
	schemaTypeMap["#/components/schemas/MyBean"] = schemaType

	// When...
	javaPackage := translateSchemaTypesToJavaPackage(schemaTypeMap, TARGET_JAVA_PACKAGE)

	// Then...
	assert.Equal(t, "MyBean", javaPackage.Classes[schemaName].Name)
	assertJavaClassCorrectlyRelatesToSchemaType(t, schemaType, javaPackage.Classes[schemaName])
}

func TestTranslateSchemaTypesToJavaPackageWithClassWithArrayOfReferenceToClass(t *testing.T) {
	// Given...
	schemaTypeMap := make(map[string]*SchemaType)
	var referencedSchemaType *SchemaType
	referencedSchemaName := "MyReferencedBean"
	referencedOwnProp := NewProperty(referencedSchemaName, "#/components/schemas/MyReferencedBean", "", "object", nil, referencedSchemaType, Cardinality{min: 0, max: 100})
	referencedSchemaType = NewSchemaType(referencedSchemaName, "", referencedOwnProp, nil)
	schemaTypeMap["#/components/schemas/MyReferencedBean"] = referencedSchemaType
	propName1 := "MyRandomProperty1"
	property1 := NewProperty(propName1, "#/components/schemas/MyBean/"+propName1, "", "object", nil, referencedSchemaType, Cardinality{min: 0, max: 1})
	properties := make(map[string]*Property)
	properties["#/components/schemas/MyBean/"+propName1] = property1
	var schemaType *SchemaType
	schemaName := "MyBean"
	ownProp := NewProperty(schemaName, "#/components/schemas/MyBean", "", "object", nil, schemaType, Cardinality{min: 0, max: 1})
	schemaType = NewSchemaType(schemaName, "", ownProp, properties)
	schemaTypeMap["#/components/schemas/MyBean"] = schemaType

	// When...
	javaPackage := translateSchemaTypesToJavaPackage(schemaTypeMap, TARGET_JAVA_PACKAGE)

	// Then...
	assert.Equal(t, "MyBean", javaPackage.Classes[schemaName].Name)
	assertJavaClassCorrectlyRelatesToSchemaType(t, schemaType, javaPackage.Classes[schemaName])
}

func TestTranslateSchemaTypesToJavaPackageWithClassWithRequiredProperty(t *testing.T) {
	// Given...
	propName1 := "MyRandomProperty1"
	property1 := NewProperty(propName1, "#/components/schemas/MyBean/"+propName1, "", "string", nil, nil, Cardinality{min: 1, max: 1})
	properties := make(map[string]*Property)
	properties["#/components/schemas/MyBean/"+propName1] = property1
	schemaTypeMap := make(map[string]*SchemaType)
	var schemaType *SchemaType
	schemaName := "MyBean"
	ownProp := NewProperty(schemaName, "#/components/schemas/MyBean", "", "object", nil, schemaType, Cardinality{min: 0, max: 1})
	schemaType = NewSchemaType(schemaName, "", ownProp, properties)
	schemaTypeMap["#/components/schemas/MyBean"] = schemaType

	// When...
	javaPackage := translateSchemaTypesToJavaPackage(schemaTypeMap, TARGET_JAVA_PACKAGE)

	// Then...
	assert.Equal(t, "MyBean", javaPackage.Classes[schemaName].Name)
	assertJavaClassCorrectlyRelatesToSchemaType(t, schemaType, javaPackage.Classes[schemaName])
}

func TestTranslateSchemaTypesToJavaPackageWithClassWithEnum(t *testing.T) {
	// Given...
	possibleValues := map[string]string{
		"randValue1": "randValue1",
		"randValue2": "randValue2",
	}
	schemaTypeMap := make(map[string]*SchemaType)
	var schemaType *SchemaType
	schemaName := "MyEnum"
	ownProp := NewProperty(schemaName, "#/components/schemas/MyEnum", "", "string", possibleValues, schemaType, Cardinality{min: 0, max: 1})
	schemaType = NewSchemaType(schemaName, "", ownProp, nil)
	schemaTypeMap["#/components/schemas/MyEnum"] = schemaType

	// When...
	javaPackage := translateSchemaTypesToJavaPackage(schemaTypeMap, TARGET_JAVA_PACKAGE)

	// Then...
	assert.NotNil(t, javaPackage.Enums[schemaName])
	assertJavaEnumRelatesToSchemaType(t, schemaType, *javaPackage.Enums[schemaName])
}
