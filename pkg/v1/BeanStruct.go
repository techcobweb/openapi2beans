package v1_generator

// TODO: Instead of type Object, Variable, we should be using *Object and *Variable.
// What we have here will pass by value rather than by reference...
// Possibly not overly important for performance in this application,
// but just so you know it's inefficient.

type Bean struct {
	// Would call it 'package' but that's a keyword in go!
	BeanPackage string
	Name        string
	Description string
	Variables   map[string]SchemaPart
}

type Cardinality struct {
	min int
	max int
}

type SchemaPart interface {
	// For arrays/lists, this is the type of each element.
	GetType() string

	GetName() string
	GetDescription() string

	GetVariables() map[string]SchemaPart

	IsSetInConstructor() bool

	IsArrayOrList() bool
}

type VariableSchema struct {
	varDescription string

	// Would call it 'type' but that's a keyword in go!
	// For arrays/lists, this is the type of the variable/property/data member
	varTypeName string

	varName string

	// Is the variable multiple values ? ie: An array or list ?
	// Also determines whether variable is set in constructor i.e. if min cardinality > 0
	cardinality Cardinality
}

func (variable VariableSchema) GetType() string {
	return variable.varTypeName
}

func (variable VariableSchema) GetName() string {
	return variable.varName
}

func (variable VariableSchema) GetDescription() string {
	return variable.varDescription
}

func (variable VariableSchema) IsArrayOrList() bool {
	isArrayOrList := false
	if variable.cardinality.max > 1 {
		isArrayOrList = true
	}
	return isArrayOrList
}

func (variable VariableSchema) IsSetInConstructor() bool {
	isSetInConstructor := false
	if variable.cardinality.min > 0 {
		isSetInConstructor = true
	}
	return isSetInConstructor
}

func (variable VariableSchema) GetVariables() map[string]SchemaPart {
	return nil
}

func (variable *VariableSchema) SetType(varType string) {
	variable.varTypeName = varType
}

type ObjectSchema struct {
	description string
	varName     string
	variables   map[string]SchemaPart
}

func (obj ObjectSchema) GetType() string {
	return "object"
}

func (obj ObjectSchema) GetName() string {
	return obj.varName
}

func (obj ObjectSchema) GetDescription() string {
	return obj.description
}

func (obj ObjectSchema) IsSetInConstructor() bool {
	return false
}

func (obj ObjectSchema) GetVariables() map[string]SchemaPart {
	return obj.variables
}

func (obj ObjectSchema) IsArrayOrList() bool {
	return false
}