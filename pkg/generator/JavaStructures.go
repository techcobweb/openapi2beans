package generator

type JavaPackage struct {
	Name       string
	classes    map[string]*JavaClass
	enums      []JavaEnum
	interfaces []JavaInterface
}

type JavaClass struct {
	Name               string
	Description        string
	Includes           []string
	JavaPackage        *JavaPackage
	InheritedInterface *JavaInterface
	DataMembers        []*DataMember
}

func NewJavaClass(name string, description string, includes []string, javaPackage *JavaPackage, inheritedInterface *JavaInterface, dataMembers []*DataMember) *JavaClass {
	javaClass := JavaClass{
		Name:               name,
		Description:        description,
		Includes:           includes,
		JavaPackage:        javaPackage,
		InheritedInterface: inheritedInterface,
		DataMembers: dataMembers,
	}
	return &javaClass
}

type DataMember struct {
	Name        string
	MemberType  string
	Description string
}

type JavaEnum struct {
	name        string
	description string
	enumValues  []string
}

type JavaInterface struct {
	name              string
	description       string
	inheritingClasses []JavaClass
	javaPackage       JavaPackage
}
