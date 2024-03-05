package generator

import (
	"testing"

	"github.com/galasa-dev/cli/pkg/files"
	"github.com/stretchr/testify/assert"
)

const (
	TARGET_JAVA_PACKAGE = "generated"
)

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
		assert.Contains(t, generatedFile, dataMember.MemberType + " " + dataMember.Name)
		assert.Contains(t, generatedFile, "// " + dataMember.Description)
		assert.Contains(t, generatedFile, "public " + dataMember.MemberType + " Get" + dataMember.Name + "() {")
		assert.Contains(t, generatedFile, "this." + dataMember.Name + " = " + dataMember.Name)
		assert.Contains(t, generatedFile, "public void Set" + dataMember.Name + "(" + dataMember.MemberType + " " + dataMember.Name + ") {")
		assert.Contains(t, generatedFile, "this." + dataMember.Name + " = " + dataMember.Name)
	}
}

func assertEnumFileGeneratedOk(t *testing.T, generatedFile string, javaEnum *JavaEnum) {
	assert.Contains(t, generatedFile, "package "+ TARGET_JAVA_PACKAGE)
	assert.Contains(t, generatedFile, "public enum " + javaEnum.Name)
	for _, value := range javaEnum.EnumValues {
		assert.Contains(t, generatedFile, value)
	}
}

func TestPackageStructParsesToTemplate(t *testing.T) {
	// Given...
	className := "MyBean"
	var javaPackage JavaPackage
	javaPackage.Name = TARGET_JAVA_PACKAGE
	class := NewJavaClass(className, "", nil, &javaPackage, nil, nil, nil)
	mockFileSystem := files.NewMockFileSystem()
	storeFilepath := "generated"
	generatedCodeFilePath := storeFilepath + "/" + className + ".java"

	// When...
	err := createJavaClassFile(class, mockFileSystem, storeFilepath)

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
		Description: "random member for test purposes",
		MemberType: "String",
	}
	dataMembers := []*DataMember{}
	dataMembers = append(dataMembers, &dataMember)
	class := NewJavaClass(className, "", nil, &javaPackage, nil, dataMembers, nil)
	mockFileSystem := files.NewMockFileSystem()
	storeFilepath := "generated"
	generatedCodeFilePath := storeFilepath + "/" + className + ".java"

	// When...
	err := createJavaClassFile(class, mockFileSystem, storeFilepath)

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
		Description: "random member for test purposes",
		MemberType: "String",
	}
	memberName2 := "RandMember2"
	dataMember2 := DataMember {
		Name: memberName2,
		Description: "random member for test purposes",
		MemberType: "String",
	}
	dataMembers := []*DataMember{}
	dataMembers = append(dataMembers, &dataMember1, &dataMember2)
	class := NewJavaClass(className, "", nil, &javaPackage, nil, dataMembers, nil)
	mockFileSystem := files.NewMockFileSystem()
	storeFilepath := "generated"
	generatedCodeFilePath := storeFilepath + "/" + className + ".java"

	// When...
	err := createJavaClassFile(class, mockFileSystem, storeFilepath)

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
		Description: "random member for test purposes",
		MemberType: "String[]",
	}
	dataMembers := []*DataMember{}
	dataMembers = append(dataMembers, &dataMember1)
	class := NewJavaClass(className, "", nil, &javaPackage, nil, dataMembers, nil)
	mockFileSystem := files.NewMockFileSystem()
	storeFilepath := "generated"
	generatedCodeFilePath := storeFilepath + "/" + className + ".java"

	// When...
	err := createJavaClassFile(class, mockFileSystem, storeFilepath)

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
		Description: "random member for test purposes",
		MemberType: "String[]",
	}
	memberName2 := "RandMember2"
	dataMember2 := DataMember {
		Name: memberName2,
		Description: "random member for test purposes",
		MemberType: "String",
	}
	dataMembers := []*DataMember{}
	dataMembers = append(dataMembers, &dataMember1, &dataMember2)
	class := NewJavaClass(className, "", nil, &javaPackage, nil, dataMembers, nil)
	mockFileSystem := files.NewMockFileSystem()
	storeFilepath := "generated"
	generatedCodeFilePath := storeFilepath + "/" + className + ".java"

	// When...
	err := createJavaClassFile(class, mockFileSystem, storeFilepath)

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
		Description: "random member for test purposes",
		MemberType: "ReferencedClass",
	}
	dataMembers := []*DataMember{}
	dataMembers = append(dataMembers, &dataMember1)
	class := NewJavaClass(className, "", nil, &javaPackage, nil, dataMembers, nil)
	mockFileSystem := files.NewMockFileSystem()
	storeFilepath := "generated"
	generatedCodeFilePath := storeFilepath + "/" + className + ".java"

	// When...
	err := createJavaClassFile(class, mockFileSystem, storeFilepath)

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
		Description: "random member for test purposes",
		MemberType: "ReferencedClass[]",
	}
	dataMembers := []*DataMember{}
	dataMembers = append(dataMembers, &dataMember1)
	class := NewJavaClass(className, "", nil, &javaPackage, nil, dataMembers, nil)
	mockFileSystem := files.NewMockFileSystem()
	storeFilepath := "generated"
	generatedCodeFilePath := storeFilepath + "/" + className + ".java"

	// When...
	err := createJavaClassFile(class, mockFileSystem, storeFilepath)

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
		Description: "random member for test purposes",
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
	class := NewJavaClass(className, "", nil, &javaPackage, nil, dataMembers, requiredMembers)
	mockFileSystem := files.NewMockFileSystem()
	storeFilepath := "generated"
	generatedCodeFilePath := storeFilepath + "/" + className + ".java"

	// When...
	err := createJavaClassFile(class, mockFileSystem, storeFilepath)

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
	enumDesc := "test enum"
	var javaPackage JavaPackage
	javaPackage.Name = TARGET_JAVA_PACKAGE
	javaEnum := JavaEnum {
		Name: enumName,
		Description: enumDesc,
		JavaPackage: &javaPackage,
	}
	mockFileSystem := files.NewMockFileSystem()
	storeFilepath := "generated"
	generatedCodeFilePath := storeFilepath + "/" + enumName + ".java"

	// When...
	err := createJavaEnumFile(&javaEnum, mockFileSystem, storeFilepath)

	// Then...
	assert.Nil(t, err)
	generatedFile := openGeneratedFile(t, mockFileSystem, generatedCodeFilePath)
	assertEnumFileGeneratedOk(t, generatedFile, &javaEnum)
}