package generator

import (
	"testing"

	"github.com/cbroglie/mustache"
	"github.com/galasa-dev/cli/pkg/files"
	"github.com/stretchr/testify/assert"
	"github.com/techcobweb/openapi2beans/pkg/embedded"
)

const (
	TARGET_JAVA_PACKAGE = "generated"
)

func getEmbeddedClassTemplate(t *testing.T) *mustache.Template{
	classTemplate, err := embedded.GetJavaClassTemplate()
	assert.Nil(t, err)
	return classTemplate
}

func getEmbeddedEnumTemplate(t *testing.T) *mustache.Template{
	enumTemplate, err := embedded.GetJavaEnumTemplate()
	assert.Nil(t, err)
	return enumTemplate
}

func openGeneratedFile(t *testing.T, mockFileSystem files.FileSystem, generatedCodeFilepath string) string{
	exists, err := mockFileSystem.Exists(generatedCodeFilepath)
	assert.Nil(t, err)
	assert.True(t, exists)
	generatedFile, err := mockFileSystem.ReadTextFile(generatedCodeFilepath)
	assert.Nil(t, err)
	return generatedFile
}

func assertClassFileGeneratedOk(t *testing.T, generatedFile string, className string) {
	assert.Contains(t, generatedFile, "package "+ TARGET_JAVA_PACKAGE)
	assert.Contains(t, generatedFile, "public class "+ className)
	assert.Contains(t, generatedFile, "public "+ className +" (")
}

func assertVariablesGeneratedOk(t *testing.T, generatedFile string, dataMembers []*DataMember) {
	for _, dataMember := range dataMembers {
		assert.Contains(t, generatedFile, "private " + dataMember.MemberType + " " + dataMember.Name)
		for _, line := range dataMember.Description {
			assert.Contains(t, generatedFile, "// " + line)
		}
		assert.Contains(t, generatedFile, "public " + dataMember.MemberType + " Get" + dataMember.CamelCaseName + "() {")
		assert.Contains(t, generatedFile, "this." + dataMember.Name + " = " + dataMember.Name)
		assert.Contains(t, generatedFile, "public void Set" + dataMember.CamelCaseName + "(" + dataMember.MemberType + " " + dataMember.Name + ") {")
		assert.Contains(t, generatedFile, "this." + dataMember.Name + " = " + dataMember.Name)
	}
}

func assertConstantsGeneratedOk(t *testing.T, generatedFile string, constDataMembers []*DataMember) {
	for _, constDataMember := range constDataMembers{
		assert.Contains(t, generatedFile, "public static final " + constDataMember.MemberType + " " + constDataMember.Name + " = " + constDataMember.ConstantVal)
		for _, line := range constDataMember.Description {
			assert.Contains(t, generatedFile, "// " + line)
		}
	}
}

func assertEnumFileGeneratedOk(t *testing.T, generatedFile string, javaEnum *JavaEnum) {
	assert.Contains(t, generatedFile, "package "+ TARGET_JAVA_PACKAGE)
	assert.Contains(t, generatedFile, "public enum " + javaEnum.Name)
	for _, value := range javaEnum.EnumValues {
		assert.Contains(t, generatedFile, value + ",")
	}
}

func TestPackageStructParsesToTemplate(t *testing.T) {
	// Given...
	className := "MyBean"
	var javaPackage JavaPackage
	javaPackage.Name = TARGET_JAVA_PACKAGE
	class := NewJavaClass(className, []string{}, &javaPackage, nil, nil, nil, nil)
	mockFileSystem := files.NewMockFileSystem()
	storeFilepath := "generated"
	generatedCodeFilePath := storeFilepath + "/" + className + ".java"

	// When...
	err := createJavaClassFile(class, mockFileSystem, getEmbeddedClassTemplate(t), storeFilepath)

	// Then...
	assert.Nil(t, err)
	generatedFile := openGeneratedFile(t, mockFileSystem, generatedCodeFilePath)
	assertClassFileGeneratedOk(t, generatedFile, className)
}

func TestPackageStructParsesToTemplateWithClassWithMember(t *testing.T) {
	// Given...
	className := "MyBean"
	var javaPackage JavaPackage
	javaPackage.Name = TARGET_JAVA_PACKAGE
	memberName := "RandMember"
	dataMember := DataMember {
		Name: memberName,
		Description: []string{"random member for test purposes"},
		MemberType: "String",
	}
	dataMembers := []*DataMember{&dataMember}
	class := NewJavaClass(className, []string{}, &javaPackage, nil, dataMembers, nil, nil)
	mockFileSystem := files.NewMockFileSystem()
	storeFilepath := "generated"
	generatedCodeFilePath := storeFilepath + "/" + className + ".java"

	// When...
	err := createJavaClassFile(class, mockFileSystem, getEmbeddedClassTemplate(t), storeFilepath)

	// Then...
	assert.Nil(t, err)
	generatedFile := openGeneratedFile(t, mockFileSystem, generatedCodeFilePath)
	assertClassFileGeneratedOk(t, generatedFile, className)
	assertVariablesGeneratedOk(t, generatedFile, dataMembers)
}

func TestPackageStructParsesToTemplateWithClassWithMultipleMembers(t *testing.T) {
	// Given...
	className := "MyBean"
	var javaPackage JavaPackage
	javaPackage.Name = TARGET_JAVA_PACKAGE
	memberName1 := "RandMember1"
	dataMember1 := DataMember {
		Name: memberName1,
		Description: []string{"random member for test purposes"},
		MemberType: "String",
	}
	memberName2 := "RandMember2"
	dataMember2 := DataMember {
		Name: memberName2,
		Description: []string{"random member for test purposes"},
		MemberType: "String",
	}
	dataMembers := []*DataMember{&dataMember1, &dataMember2}
	class := NewJavaClass(className, []string{}, &javaPackage, nil, dataMembers, nil, nil)
	mockFileSystem := files.NewMockFileSystem()
	storeFilepath := "generated"
	generatedCodeFilePath := storeFilepath + "/" + className + ".java"

	// When...
	err := createJavaClassFile(class, mockFileSystem, getEmbeddedClassTemplate(t), storeFilepath)

	// Then...
	assert.Nil(t, err)
	generatedFile := openGeneratedFile(t, mockFileSystem, generatedCodeFilePath)
	assertClassFileGeneratedOk(t, generatedFile, className)
	assertVariablesGeneratedOk(t, generatedFile, dataMembers)
}

func TestPackageStructParsesToTemplateWithClassWithArrayDataMember(t *testing.T) {
	// Given...
	className := "MyBean"
	var javaPackage JavaPackage
	javaPackage.Name = TARGET_JAVA_PACKAGE
	memberName1 := "RandMember1"
	dataMember1 := DataMember {
		Name: memberName1,
		Description: []string{"random member for test purposes"},
		MemberType: "String[]",
	}
	dataMembers := []*DataMember{&dataMember1}
	class := NewJavaClass(className, []string{}, &javaPackage, nil, dataMembers, nil, nil)
	mockFileSystem := files.NewMockFileSystem()
	storeFilepath := "generated"
	generatedCodeFilePath := storeFilepath + "/" + className + ".java"

	// When...
	err := createJavaClassFile(class, mockFileSystem, getEmbeddedClassTemplate(t), storeFilepath)

	// Then...
	assert.Nil(t, err)
	generatedFile := openGeneratedFile(t, mockFileSystem, generatedCodeFilePath)
	assertClassFileGeneratedOk(t, generatedFile, className)
	assertVariablesGeneratedOk(t, generatedFile, dataMembers)
}

func TestPackageStructParsesToTemplateWithClassWithMultiDimensionalArrayDataMember(t *testing.T) {
	// Given...
	className := "MyBean"
	var javaPackage JavaPackage
	javaPackage.Name = TARGET_JAVA_PACKAGE
	memberName1 := "RandMember1"
	dataMember1 := DataMember {
		Name: memberName1,
		Description: []string{"random member for test purposes"},
		MemberType: "String[][]",
	}
	dataMembers := []*DataMember{&dataMember1}
	class := NewJavaClass(className, []string{}, &javaPackage, nil, dataMembers, nil, nil)
	mockFileSystem := files.NewMockFileSystem()
	storeFilepath := "generated"
	generatedCodeFilePath := storeFilepath + "/" + className + ".java"

	// When...
	err := createJavaClassFile(class, mockFileSystem, getEmbeddedClassTemplate(t), storeFilepath)

	// Then...
	assert.Nil(t, err)
	generatedFile := openGeneratedFile(t, mockFileSystem, generatedCodeFilePath)
	assertClassFileGeneratedOk(t, generatedFile, className)
	assertVariablesGeneratedOk(t, generatedFile, dataMembers)
}

func TestPackageStructParsesToTemplateWithClassWithMixedArrayAndPrimitiveDataMembers(t *testing.T) {
	// Given...
	className := "MyBean"
	var javaPackage JavaPackage
	javaPackage.Name = TARGET_JAVA_PACKAGE
	memberName1 := "RandMember1"
	dataMember1 := DataMember {
		Name: memberName1,
		Description: []string{"random member for test purposes"},
		MemberType: "String[]",
	}
	memberName2 := "RandMember2"
	dataMember2 := DataMember {
		Name: memberName2,
		Description: []string{"random member for test purposes"},
		MemberType: "String",
	}
	dataMembers := []*DataMember{&dataMember1, &dataMember2}
	class := NewJavaClass(className, []string{}, &javaPackage, nil, dataMembers, nil, nil)
	mockFileSystem := files.NewMockFileSystem()
	storeFilepath := "generated"
	generatedCodeFilePath := storeFilepath + "/" + className + ".java"

	// When...
	err := createJavaClassFile(class, mockFileSystem, getEmbeddedClassTemplate(t), storeFilepath)

	// Then...
	assert.Nil(t, err)
	generatedFile := openGeneratedFile(t, mockFileSystem, generatedCodeFilePath)
	assertClassFileGeneratedOk(t, generatedFile, className)
	assertVariablesGeneratedOk(t, generatedFile, dataMembers)
}

func TestPackageStructParsesToTemplateWithClassWithReferencedClassType(t *testing.T) {
	// Given...
	className := "MyBean"
	var javaPackage JavaPackage
	javaPackage.Name = TARGET_JAVA_PACKAGE
	memberName1 := "RandMember1"
	dataMember1 := DataMember {
		Name: memberName1,
		Description: []string{"random member for test purposes"},
		MemberType: "ReferencedClass",
	}
	dataMembers := []*DataMember{&dataMember1}
	class := NewJavaClass(className, []string{}, &javaPackage, nil, dataMembers, nil, nil)
	mockFileSystem := files.NewMockFileSystem()
	storeFilepath := "generated"
	generatedCodeFilePath := storeFilepath + "/" + className + ".java"

	// When...
	err := createJavaClassFile(class, mockFileSystem, getEmbeddedClassTemplate(t), storeFilepath)

	// Then...
	assert.Nil(t, err)
	generatedFile := openGeneratedFile(t, mockFileSystem, generatedCodeFilePath)
	assertClassFileGeneratedOk(t, generatedFile, className)
	assertVariablesGeneratedOk(t, generatedFile, dataMembers)
}

func TestPackageStructParsesToTemplateWithClassWithArrayOfReferencedClassType(t *testing.T) {
	// Given...
	className := "MyBean"
	var javaPackage JavaPackage
	javaPackage.Name = TARGET_JAVA_PACKAGE
	memberName1 := "RandMember1"
	dataMember1 := DataMember {
		Name: memberName1,
		Description: []string{"random member for test purposes"},
		MemberType: "ReferencedClass[]",
	}
	dataMembers := []*DataMember{&dataMember1}
	class := NewJavaClass(className, []string{}, &javaPackage, nil, dataMembers, nil, nil)
	mockFileSystem := files.NewMockFileSystem()
	storeFilepath := "generated"
	generatedCodeFilePath := storeFilepath + "/" + className + ".java"

	// When...
	err := createJavaClassFile(class, mockFileSystem, getEmbeddedClassTemplate(t), storeFilepath)

	// Then...
	assert.Nil(t, err)
	generatedFile := openGeneratedFile(t, mockFileSystem, generatedCodeFilePath)
	assertClassFileGeneratedOk(t, generatedFile, className)
	assertVariablesGeneratedOk(t, generatedFile, dataMembers)
}

func TestPackageStructParsesToTemplateWithClassWithRequiredProperty(t *testing.T) {
	// Given...
	className := "MyBean"
	var javaPackage JavaPackage
	javaPackage.Name = TARGET_JAVA_PACKAGE
	memberName1 := "RandMember1"
	dataMember1 := DataMember {
		Name: memberName1,
		Description: []string{"random member for test purposes"},
		MemberType: "String",
	}
	requiredMember1 := RequiredMember {
		IsFirst: true,
		DataMember: &dataMember1,
	}
	dataMembers := []*DataMember{}
	dataMembers = append(dataMembers, &dataMember1)
	requiredMembers := []*RequiredMember{}
	requiredMembers = append(requiredMembers, &requiredMember1)
	class := NewJavaClass(className, []string{}, &javaPackage, nil, dataMembers, requiredMembers, nil)
	mockFileSystem := files.NewMockFileSystem()
	storeFilepath := "generated"
	generatedCodeFilePath := storeFilepath + "/" + className + ".java"

	// When...
	err := createJavaClassFile(class, mockFileSystem, getEmbeddedClassTemplate(t), storeFilepath)

	// Then...
	assert.Nil(t, err)
	generatedFile := openGeneratedFile(t, mockFileSystem, generatedCodeFilePath)
	assertClassFileGeneratedOk(t, generatedFile, className)
	assertVariablesGeneratedOk(t, generatedFile, dataMembers)
	constructor := `public MyBean (String RandMember1) {
        this.RandMember1 = RandMember1;
    }`
	assert.Contains(t, generatedFile, constructor)
}

func TestPackageStructParsesToJavaEnumTemplate(t *testing.T) {
	// Given...
	enumName := "MyEnum"
	enumDesc := []string{"test enum"}
	var javaPackage JavaPackage
	javaPackage.Name = TARGET_JAVA_PACKAGE
	javaEnum := JavaEnum {
		Name: enumName,
		Description: enumDesc,
		EnumValues: []string{"randVal1", "randVal2"},
		JavaPackage: &javaPackage,
	}
	mockFileSystem := files.NewMockFileSystem()
	storeFilepath := "generated"
	generatedCodeFilePath := storeFilepath + "/" + enumName + ".java"

	// When...
	err := createJavaEnumFile(&javaEnum, mockFileSystem, getEmbeddedEnumTemplate(t), storeFilepath)

	// Then...
	assert.Nil(t, err)
	generatedFile := openGeneratedFile(t, mockFileSystem, generatedCodeFilePath)
	assertEnumFileGeneratedOk(t, generatedFile, &javaEnum)
}

func TestPackageStructWithClassWithReferenceToEnumParsesCorrectly(t *testing.T) {
	// Given...
	className := "MyBean"
	classDesc := []string{"test class"}
	enumName := "MyEnum"
	enumDesc := []string{"test enum"}
	javaPackage := NewJavaPackage(TARGET_JAVA_PACKAGE)
	javaEnum := JavaEnum {
		Name: enumName,
		Description: enumDesc,
		EnumValues: []string{"randVal1", "randVal2"},
		JavaPackage: javaPackage,
	}
	dataMember := &DataMember {
		Name: "enumMember",
		Description: enumDesc,
		MemberType: enumName,
		Required: true,
	}
	dataMembers := []*DataMember{dataMember}
	class := NewJavaClass(className, classDesc, javaPackage, nil, dataMembers, nil, nil)
	javaPackage.Classes[className] = class
	javaPackage.Enums[enumName] = &javaEnum
	mockFileSystem := files.NewMockFileSystem()
	storeFilepath := "generated"
	generatedEnumPath := storeFilepath + "/" + enumName + ".java"
	generatedClassPath := storeFilepath + "/" + className + ".java"

	// When...
	convertJavaPackageToJavaFiles(javaPackage, mockFileSystem, storeFilepath)

	// Then...
	generatedEnumFile := openGeneratedFile(t, mockFileSystem, generatedEnumPath)
	assertEnumFileGeneratedOk(t, generatedEnumFile, &javaEnum)
	generatedClassFile := openGeneratedFile(t, mockFileSystem, generatedClassPath)
	assertClassFileGeneratedOk(t, generatedClassFile, className)
	assertVariablesGeneratedOk(t, generatedClassFile, class.DataMembers)
}

func TestPackageStructParsesToTemplateWithClassWithConstantMember(t *testing.T) {
	// Given...
	className := "MyBean"
	var javaPackage JavaPackage
	javaPackage.Name = TARGET_JAVA_PACKAGE
	memberName := "RAND_MEMBER"
	constDataMember := DataMember {
		Name: memberName,
		Description: []string{"random constant member for test purposes"},
		MemberType: "String",
		ConstantVal: "random string thing",
	}
	constDataMembers := []*DataMember{&constDataMember}
	class := NewJavaClass(className, []string{}, &javaPackage, nil, nil, nil, constDataMembers)
	mockFileSystem := files.NewMockFileSystem()
	storeFilepath := "generated"
	generatedCodeFilePath := storeFilepath + "/" + className + ".java"

	// When...
	err := createJavaClassFile(class, mockFileSystem, getEmbeddedClassTemplate(t), storeFilepath)

	// Then...
	assert.Nil(t, err)
	generatedFile := openGeneratedFile(t, mockFileSystem, generatedCodeFilePath)
	assertClassFileGeneratedOk(t, generatedFile, className)
	assertConstantsGeneratedOk(t, generatedFile, constDataMembers)
}