package generator

type JavaPackage struct {
	name       string
	classes    []JavaClass
	enums      []JavaEnum
	interfaces []JavaInterface
}

type JavaClass struct {
	name               string
	includes           []string
	javaPackage        JavaPackage
	inheritedInterface JavaInterface
}

type DataMember struct {
	name       string
	memberType string
}

type JavaEnum struct {
	name       string
	enumValues []EnumValue
}

type EnumValue struct {
	name  string
	value string
}

type JavaInterface struct {
	name              string
	inheritingClasses []JavaClass
	javaPackage       JavaPackage
}
