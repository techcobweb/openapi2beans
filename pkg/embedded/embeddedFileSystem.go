package embedded

import (
	"embed"
	openapi2beans_errors "github.com/techcobweb/openapi2beans/pkg/errors"
)

type ReadOnlyFileSystem interface {
	ReadFile(filePath string) ([]byte, error)
}

type EmbeddedFileSystem struct {
	embeddedFileSystem embed.FS
}

func NewReadOnlyFileSystem() ReadOnlyFileSystem {
	result := EmbeddedFileSystem{
		embeddedFileSystem: embeddedFileSystem,
	}
	return &result
}

//------------------------------------------------------------------------------------
// Interface methods...
//------------------------------------------------------------------------------------

// The only thing which this class actually supports.
func (fs *EmbeddedFileSystem) ReadFile(filePath string) ([]byte, error) {

	bytes, err := fs.embeddedFileSystem.ReadFile(filePath)
	if err != nil {
		openapi2beans_errors.NewError("Error: unable to read embedded file system. Reason is: %s", err.Error())
	}
	return bytes, err
}
