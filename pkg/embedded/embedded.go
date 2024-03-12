package embedded

import (
	"embed"

	"github.com/cbroglie/mustache"
	openapi2beans_errors "github.com/techcobweb/openapi2beans/pkg/errors"
)

// Embed all the template files into the go executable, so there are no extra files
// we need to ship/install/locate on the target machine.
// We can access the "embedded" file system as if they are normal files.
//
//go:embed templates/*
var embeddedFileSystem embed.FS

// An instance of the ReadOnlyFileSystem interface, set once, used many times.
// It just delegates to teh embed.FS
var readOnlyFileSystem ReadOnlyFileSystem

type templates struct {
	JavaClassTemplate         *mustache.Template
	JavaEnumTemplate          *mustache.Template
	JavaAbstractClassTemplate *mustache.Template
}

const (
	JAVA_CLASS_TEMPLATE_FILEPATH                    = "templates/JavaClassTemplate.mustache"
	JAVA_ENUM_TEMPLATE_FILEPATH                     = "templates/JavaEnumTemplate.mustache"
	JAVA_ABSTRACT_CLASS_TEMPLATE_FILEPATH           = "templates/JavaAbstractClassTemplate.mustache"
)

var (
	templatesCache                    *templates = nil
)

func GetReadOnlyFileSystem() ReadOnlyFileSystem {
	if readOnlyFileSystem == nil {
		readOnlyFileSystem = NewReadOnlyFileSystem()
	}
	return readOnlyFileSystem
}

func GetJavaClassTemplate() (*mustache.Template, error) {
	var err error
	fs := GetReadOnlyFileSystem()
	// Note: The cache is set when we read the versions from the embedded file.
	templatesCache, err = readTemplatesFromEmbeddedFiles(fs, templatesCache)
	var template *mustache.Template
	if err == nil {
		template = templatesCache.JavaClassTemplate
	} else {
		err = openapi2beans_errors.NewError("Failed to read templates from embedded file. Reason is: %s", err.Error())
	}
	return template, err
}

func GetJavaEnumTemplate() (*mustache.Template, error) {
	var err error
	fs := GetReadOnlyFileSystem()
	// Note: The cache is set when we read the versions from the embedded file.
	templatesCache, err = readTemplatesFromEmbeddedFiles(fs, templatesCache)
	var template *mustache.Template
	if err == nil {
		template = templatesCache.JavaEnumTemplate
	} else {
		err = openapi2beans_errors.NewError("Failed to read templates from embedded file. Reason is: %s", err.Error())
	}
	return template, err
}

func GetJavaAbstractClassTemplate() (*mustache.Template, error) {
	var err error
	fs := GetReadOnlyFileSystem()
	// Note: The cache is set when we read the versions from the embedded file.
	templatesCache, err = readTemplatesFromEmbeddedFiles(fs, templatesCache)
	var template *mustache.Template
	if err == nil {
		template = templatesCache.JavaAbstractClassTemplate
	} else {
		err = openapi2beans_errors.NewError("Failed to read templates from embedded file. Reason is: %s", err.Error())
	}
	return template, err
}

// readVersionsFromEmbeddedFile - Reads a set of version data from an embedded property file, or returns
// a set of version data we already know about. So that the version data is only ever read once.
func readTemplatesFromEmbeddedFiles(fs ReadOnlyFileSystem, templatesAlreadyKnown *templates) (*templates, error) {
	var (
		err error
		bytes []byte
	)
	if templatesAlreadyKnown == nil {
		templatesAlreadyKnown = &templates {}
		bytes, err = fs.ReadFile(JAVA_CLASS_TEMPLATE_FILEPATH)
		if err == nil {
			templatesAlreadyKnown.JavaClassTemplate, err = mustache.ParseString(string(bytes))
			if err == nil {
				bytes, err = fs.ReadFile(JAVA_ENUM_TEMPLATE_FILEPATH)
				if err == nil {
					templatesAlreadyKnown.JavaEnumTemplate, err = mustache.ParseString(string(bytes))
					if err == nil {
						bytes, err = fs.ReadFile(JAVA_ABSTRACT_CLASS_TEMPLATE_FILEPATH)
						if err == nil {
							templatesAlreadyKnown.JavaAbstractClassTemplate, err = mustache.ParseString(string(bytes))

						}
					}
				}

			}
		}
	}
	return templatesAlreadyKnown, err
}
