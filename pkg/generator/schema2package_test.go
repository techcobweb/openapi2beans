package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertJavaClassCorrectlyRelatesToSchemaType(t *testing.T, schemaType *SchemaType, class *JavaClass) {
	assert.Equal(t, schemaType.name, class.Name)
	schemaPath := "#/components/schemas/" + schemaType.name

	if len(schemaType.properties) > 0 {
		assert.Equal(t, len(schemaType.properties), len(class.DataMembers))
	}
	for _, dataMember := range class.DataMembers {
		comparisonSchemaProperty, exists := schemaType.properties[schemaPath+"/"+dataMember.Name]
		assert.True(t, exists)
		expectedName := comparisonSchemaProperty.name
		assert.Equal(t, expectedName, dataMember.Name)
		expectedType := getExpectedType(comparisonSchemaProperty)
		assert.Equal(t, dataMember.MemberType, expectedType)
		if dataMember.ConstantVal != "" {
			assert.True(t, comparisonSchemaProperty.IsConstant())
			posVal := possibleValuesToEnumValues(comparisonSchemaProperty.possibleValues)
			assert.Equal(t, 1, len(posVal))
			assert.Equal(t, posVal[0], dataMember.ConstantVal)
		}
	}

	requiredPropertiesVisited := 0
	for _, requiredMember := range class.RequiredMembers {
		comparisonSchemaProperty, exists := schemaType.properties[schemaPath+"/"+requiredMember.DataMember.Name]
		assert.True(t, exists)
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

func assertJavaEnumRelatesToSchemaType(t *testing.T, schemaType *SchemaType, javaEnum *JavaEnum) {
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
	class, classExists := javaPackage.Classes[schemaName]
	assert.True(t, classExists)
	assertJavaClassCorrectlyRelatesToSchemaType(t, schemaType, class)
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
	class, classExists := javaPackage.Classes[schemaName]
	assert.True(t, classExists)
	assertJavaClassCorrectlyRelatesToSchemaType(t, schemaType, class)
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
	class, classExists := javaPackage.Classes[schemaName]
	assert.True(t, classExists)
	assertJavaClassCorrectlyRelatesToSchemaType(t, schemaType, class)
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
	class, classExists := javaPackage.Classes[schemaName]
	assert.True(t, classExists)
	assertJavaClassCorrectlyRelatesToSchemaType(t, schemaType, class)
}

func TestTranslateSchemaTypesToJavaPackageWithClassWithArrayOfArray(t *testing.T) {
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
	class, classExists := javaPackage.Classes[schemaName]
	assert.True(t, classExists)
	assertJavaClassCorrectlyRelatesToSchemaType(t, schemaType, class)
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
	class, classExists := javaPackage.Classes[schemaName]
	assert.True(t, classExists)
	assertJavaClassCorrectlyRelatesToSchemaType(t, schemaType, class)
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
	class, classExists := javaPackage.Classes[schemaName]
	assert.True(t, classExists)
	assertJavaClassCorrectlyRelatesToSchemaType(t, schemaType, class)
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
	class, classExists := javaPackage.Classes[schemaName]
	assert.True(t, classExists)
	assert.Equal(t, schemaName, class.Name)
	assertJavaClassCorrectlyRelatesToSchemaType(t, schemaType, class)
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
	enum, enumExists := javaPackage.Enums[schemaName]
	assert.True(t, enumExists)
	assertJavaEnumRelatesToSchemaType(t, schemaType, enum)
}

func TestTranslateSchemaTypesToJavaPackageWithClassWithStringConstant(t *testing.T) {
	// Given...
	propName1 := "MyConstant"
	possibleValues := map[string]string{
		"constVal": "constVal",
	}
	property := NewProperty(propName1, "#/components/schemas/MyBean/"+propName1, "", "string", possibleValues, nil, Cardinality{min: 0, max: 1})
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
	class, classExists := javaPackage.Classes[schemaName]
	assert.True(t, classExists)
	assert.NotEmpty(t, class.DataMembers)
	assertJavaClassCorrectlyRelatesToSchemaType(t, schemaType, class)
}