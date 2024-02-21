package generator

type Bean struct {
	// Would call it 'package' but that's a keyword in go!
	beanPackage string
	object      Object
}

type SchemaPart interface {
	GetType() string
	GetName() string
	GetDescription() string
}

type Variable struct {
	varDescription string
	// Would call it 'type' but that's a keyword in go!
	varTypeName        string
	varName            string
	isSetInConstructor bool
}

func (variable Variable) GetType() string {
	return variable.varTypeName
}

func (variable Variable) GetName() string {
	return variable.varName
}

func (variable Variable) GetDescription() string {
	return variable.varDescription
}

func (variable Variable) IsSetInConstructor() bool {
	return variable.isSetInConstructor
}

func (variable *Variable) SetType(varType string) {
	variable.varTypeName = varType
}

type Object struct {
	description string
	// Would call it 'type' but that's a keyword in go!
	varTypeName string
	varName     string
	variables   map[string]SchemaPart
}

func (obj Object) GetType() string {
	return obj.varTypeName
}

func (obj Object) GetName() string {
	return obj.varName
}

func (obj Object) GetDescription() string {
	return obj.description
}