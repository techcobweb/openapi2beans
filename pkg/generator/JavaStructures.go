package generator

type JavaPackage struct {
	Name       string
	Classes    map[string]*JavaClass
	Enums      map[string]*JavaEnum
	Interfaces map[string]*JavaInterface
}

func NewJavaPackage(name string) *JavaPackage {
	javaPackage := JavaPackage{
		Name:       name,
		Classes:    make(map[string]*JavaClass),
		Enums:      make(map[string]*JavaEnum),
		Interfaces: make(map[string]*JavaInterface),
	}
	return &javaPackage
}

type JavaClass struct {
	Name               string
	Description        string
	JavaPackage        *JavaPackage
	InheritedInterface *JavaInterface
	DataMembers        []*DataMember
	RequiredMembers    []*RequiredMember
}

func NewJavaClass(name string, description string, javaPackage *JavaPackage, inheritedInterface *JavaInterface, dataMembers []*DataMember, requiredMembers []*RequiredMember) *JavaClass {
	javaClass := JavaClass{
		Name:               name,
		Description:        description,
		JavaPackage:        javaPackage,
		InheritedInterface: inheritedInterface,
		DataMembers:        dataMembers,
		RequiredMembers:    requiredMembers,
	}
	return &javaClass
}

type DataMember struct {
	Name        string
	MemberType  string
	Description string
	Required    bool
}

type RequiredMember struct {
	IsFirst    bool
	DataMember *DataMember
}

type JavaEnum struct {
	Name        string
	Description string
	EnumValues  []string
	JavaPackage *JavaPackage
}

func NewJavaEnum(name string, description string, enumValues []string, javaPackage *JavaPackage) *JavaEnum {
	javaEnum := JavaEnum{
		Name:        name,
		Description: description,
		EnumValues:  enumValues,
		JavaPackage: javaPackage,
	}
	return &javaEnum
}

type JavaInterface struct {
	Name        string
	Description string
	JavaPackage *JavaPackage
}
