package generator

import (
	"testing"
	"github.com/galasa-dev/cli/pkg/files"
	"github.com/stretchr/testify/assert"
)

const TARGET_JAVA_PACKAGE = "generated"

func assertClassFileGeneratedOk(t *testing.T, mockFileSystem files.FileSystem, generatedCodeFilepath string, className string) string {
	exists, err := mockFileSystem.Exists(generatedCodeFilepath)
	assert.Nil(t, err)
	assert.True(t, exists)
	generatedFile, err := mockFileSystem.ReadTextFile(generatedCodeFilepath)
	assert.Nil(t, err)
	assert.Contains(t, generatedFile, "package "+ TARGET_JAVA_PACKAGE)
	assert.Contains(t, generatedFile, "public class "+ className)
	assert.Contains(t, generatedFile, "public "+ className +" (")
	return generatedFile
}

func assertVariablesGeneratedOk(t *testing.T, generatedFile string, dataMembers []*DataMember) {
	for _, dataMember := range dataMembers {
		assert.Contains(t, generatedFile, dataMember.MemberType + " " + dataMember.Name)
		assert.Contains(t, generatedFile, "// " + dataMember.Description)
	}
}

func TestPackageStructParsesToTemplate(t *testing.T) {
	// Given...
	className := "MyBean"
	var javaPackage JavaPackage
	javaPackage.Name = TARGET_JAVA_PACKAGE
	class := NewJavaClass(className, "", nil, &javaPackage, nil, nil)
	mockFileSystem := files.NewMockFileSystem()
	storeFilepath := "generated"
	generatedCodeFilePath := storeFilepath + "/" + className + ".java"

	// When...
	err := createJavaClassFile(*class, mockFileSystem, storeFilepath)

	// Then...
	assert.Nil(t, err)
	assertClassFileGeneratedOk(t, mockFileSystem, generatedCodeFilePath, className)
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
	class := NewJavaClass(className, "", nil, &javaPackage, nil, dataMembers)
	mockFileSystem := files.NewMockFileSystem()
	storeFilepath := "generated"
	generatedCodeFilePath := storeFilepath + "/" + className + ".java"

	// When...
	err := createJavaClassFile(*class, mockFileSystem, storeFilepath)

	// Then...
	assert.Nil(t, err)
	generatedFile := assertClassFileGeneratedOk(t, mockFileSystem, generatedCodeFilePath, className)
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
	class := NewJavaClass(className, "", nil, &javaPackage, nil, dataMembers)
	mockFileSystem := files.NewMockFileSystem()
	storeFilepath := "generated"
	generatedCodeFilePath := storeFilepath + "/" + className + ".java"

	// When...
	err := createJavaClassFile(*class, mockFileSystem, storeFilepath)

	// Then...
	assert.Nil(t, err)
	generatedFile := assertClassFileGeneratedOk(t, mockFileSystem, generatedCodeFilePath, className)
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
	memberName2 := "RandMember2"
	dataMember2 := DataMember {
		Name: memberName2,
		Description: "random member for test purposes",
		MemberType: "String",
	}
	dataMembers := []*DataMember{}
	dataMembers = append(dataMembers, &dataMember1, &dataMember2)
	class := NewJavaClass(className, "", nil, &javaPackage, nil, dataMembers)
	mockFileSystem := files.NewMockFileSystem()
	storeFilepath := "generated"
	generatedCodeFilePath := storeFilepath + "/" + className + ".java"

	// When...
	err := createJavaClassFile(*class, mockFileSystem, storeFilepath)

	// Then...
	assert.Nil(t, err)
	generatedFile := assertClassFileGeneratedOk(t, mockFileSystem, generatedCodeFilePath, className)
	assertVariablesGeneratedOk(t, generatedFile, dataMembers)
}