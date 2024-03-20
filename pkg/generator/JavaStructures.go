package generator

type JavaPackage struct {
	Name            string
	Classes         map[string]*JavaClass
	Enums           map[string]*JavaEnum
	AbstractClasses map[string]*JavaAbstractClass
}

func NewJavaPackage(name string) *JavaPackage {
	javaPackage := JavaPackage{
		Name:            name,
		Classes:         make(map[string]*JavaClass),
		Enums:           make(map[string]*JavaEnum),
		AbstractClasses: make(map[string]*JavaAbstractClass),
	}
	return &javaPackage
}

type JavaClass struct {
	Name                string
	Description         []string
	JavaPackage         *JavaPackage
	InheritedClass      *JavaAbstractClass
	DataMembers         []*DataMember
	RequiredMembers     []*RequiredMember
	ConstantDataMembers []*DataMember
}

func NewJavaClass(name string, description []string, javaPackage *JavaPackage, inheritedClass *JavaAbstractClass, dataMembers []*DataMember, requiredMembers []*RequiredMember, constantDataMembers []*DataMember) *JavaClass {
	javaClass := JavaClass{
		Name:                name,
		Description:         description,
		JavaPackage:         javaPackage,
		InheritedClass:      inheritedClass,
		DataMembers:         dataMembers,
		RequiredMembers:     requiredMembers,
		ConstantDataMembers: constantDataMembers,
	}
	return &javaClass
}

type DataMember struct {
	Name          string
	CamelCaseName string
	MemberType    string
	Description   []string
	Required      bool
	ConstantVal   string
}

func (dataMember DataMember) IsConstant() bool {
	return dataMember.ConstantVal != ""
}

type RequiredMember struct {
	IsFirst    bool
	DataMember *DataMember
}

type JavaEnum struct {
	Name        string
	Description []string
	EnumValues  []string
	JavaPackage *JavaPackage
}

func NewJavaEnum(name string, description []string, enumValues []string, javaPackage *JavaPackage) *JavaEnum {
	javaEnum := JavaEnum{
		Name:        name,
		Description: description,
		EnumValues:  enumValues,
		JavaPackage: javaPackage,
	}
	return &javaEnum
}

type JavaAbstractClass struct {
	Name        string
	Description []string
	JavaPackage *JavaPackage
}
