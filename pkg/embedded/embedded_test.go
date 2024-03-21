package embedded

import (
	"testing"

	"github.com/cbroglie/mustache"
	"github.com/stretchr/testify/assert"
)

type MockReadOnlyFileSystem struct {
	files map[string]string
}

func NewMockReadOnlyFileSystem() *MockReadOnlyFileSystem {
	fs := MockReadOnlyFileSystem{
		files: make(map[string]string, 0),
	}
	return &fs
}

// WriteFile - This function is not on the ReadOnlyFileSystem interface, but does allow unit tests
// to add data files to the mock file system, so the code under test can read it back.
func (fs *MockReadOnlyFileSystem) WriteFile(filePath string, content string) {
	fs.files[filePath] = content
}

func (fs *MockReadOnlyFileSystem) ReadFile(filePath string) ([]byte, error) {
	content := fs.files[filePath]
	return []byte(content), nil
}

func TestGetJavaClassReturnsTemplate(t *testing.T) {
	// Given...
	var (
		err error
		template *mustache.Template
		rendered string
	)

	// When...
	template, err = GetJavaClassTemplate()

	// Then...
	assert.Nil(t, err)
	assert.NotNil(t, template)
	rendered, err = template.Render("")
	assert.Nil(t, err)
	assert.Contains(t, rendered, "package")
	assert.Contains(t, rendered, "public class")
}

func TestGetJavaEnumReturnsTemplate(t *testing.T) {
	// Given...
	var (
		err error
		template *mustache.Template
		rendered string
	)

	// When...
	template, err = GetJavaEnumTemplate()

	// Then...
	assert.Nil(t, err)
	assert.NotNil(t, template)
	rendered, err = template.Render("")
	assert.Nil(t, err)
	assert.Contains(t, rendered, "package")
	assert.Contains(t, rendered, "public enum")
}

func TestCanParseTemplatesFromEmbeddedFS(t *testing.T) {
	// Given...
	fs := NewMockReadOnlyFileSystem()
	javaClassTemplateContent := "class template"
	fs.WriteFile(JAVA_CLASS_TEMPLATE_FILEPATH, javaClassTemplateContent)
	javaEnumTemplateContent := "enum content"
	fs.WriteFile(JAVA_ENUM_TEMPLATE_FILEPATH, javaEnumTemplateContent)
	javaAbstractClassTemplateContent := "abstract class template"
	fs.WriteFile(JAVA_ABSTRACT_CLASS_TEMPLATE_FILEPATH, javaAbstractClassTemplateContent)

	// When...
	templates, err := readTemplatesFromEmbeddedFiles(fs, nil)

	// Then...
	assert.Nil(t, err)
	assert.NotNil(t, templates)
	var renderResult string
	// class
	assert.NotNil(t, templates.JavaClassTemplate)
	renderResult, err = templates.JavaClassTemplate.Render("")
	assert.Nil(t, err)
	assert.Equal(t, javaClassTemplateContent, renderResult)
	// enum
	assert.NotNil(t, templates.JavaEnumTemplate)
	renderResult, err = templates.JavaEnumTemplate.Render("")
	assert.Nil(t, err)
	assert.Equal(t, javaEnumTemplateContent, renderResult)
	// abstract class
	assert.NotNil(t, templates.JavaAbstractClassTemplate)
	renderResult, err = templates.JavaAbstractClassTemplate.Render("")
	assert.Nil(t, err)
	assert.Equal(t, javaAbstractClassTemplateContent, renderResult)
}

func TestDoesntRereadTemplatesWhenTemplatesAlreadyKnown(t *testing.T) {
	// Given...
	fs := NewMockReadOnlyFileSystem()
	javaClassTemplateContent := "class template"
	fs.WriteFile(JAVA_CLASS_TEMPLATE_FILEPATH, javaClassTemplateContent)
	javaEnumTemplateContent := "enum content"
	fs.WriteFile(JAVA_ENUM_TEMPLATE_FILEPATH, javaEnumTemplateContent)
	javaAbstractClassTemplateContent := "abstract class template"
	fs.WriteFile(JAVA_ABSTRACT_CLASS_TEMPLATE_FILEPATH, javaAbstractClassTemplateContent)


	// When...
	expectedClassString := "expected class string"
	expectedClassTemplate, err := mustache.ParseString(expectedClassString)
	assert.Nil(t, err)
	expectedEnumString := "expected enum string"
	expectedEnumTemplate, err := mustache.ParseString(expectedEnumString)
	assert.Nil(t, err)
	expectedAbstractClassString := "expected abstract class string"
	expectedAbstractClassTemplate, err := mustache.ParseString(expectedAbstractClassString)
	assert.Nil(t, err)

	alreadyKnownTemplates := templates{
		JavaClassTemplate: expectedClassTemplate,
		JavaEnumTemplate: expectedEnumTemplate,
		JavaAbstractClassTemplate: expectedAbstractClassTemplate,
	}

	templates, err := readTemplatesFromEmbeddedFiles(fs, &alreadyKnownTemplates)

	// Then...
	assert.Nil(t, err)
	assert.NotNil(t, templates)
	var renderResult string
	// class
	assert.NotNil(t, templates.JavaClassTemplate)
	renderResult, err = templates.JavaClassTemplate.Render("")
	assert.Nil(t, err)
	assert.Equal(t, expectedClassString, renderResult)
	// enum
	assert.NotNil(t, templates.JavaEnumTemplate)
	renderResult, err = templates.JavaEnumTemplate.Render("")
	assert.Nil(t, err)
	assert.Equal(t, expectedEnumString, renderResult)
	// abstract class
	assert.NotNil(t, templates.JavaAbstractClassTemplate)
	renderResult, err = templates.JavaAbstractClassTemplate.Render("")
	assert.Nil(t, err)
	assert.Equal(t, expectedAbstractClassString, renderResult)
}