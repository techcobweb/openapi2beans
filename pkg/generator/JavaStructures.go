package generator

type JavaPackage struct {
	Name       string
	Classes    map[string]*JavaClass
	Enums      map[string]*JavaEnum
	Interfaces map[string]*JavaInterface
}

type JavaClass struct {
	Name               string
	Description        string
	Includes           []string
	JavaPackage        *JavaPackage
	InheritedInterface *JavaInterface
	DataMembers        []*DataMember
	RequiredMembers    []*RequiredMember
}

func NewJavaClass(name string, description string, includes []string, javaPackage *JavaPackage, inheritedInterface *JavaInterface, dataMembers []*DataMember, requiredMembers []*RequiredMember) *JavaClass {
	javaClass := JavaClass{
		Name:               name,
		Description:        description,
		Includes:           includes,
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

type JavaInterface struct {
	Name              string
	Description       string
	InheritingClasses []JavaClass
	JavaPackage       *JavaPackage
}
