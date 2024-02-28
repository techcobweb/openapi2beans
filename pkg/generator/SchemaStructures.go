package generator

//
// SCHEMA TYPE
//
// SchemaType describes a schema part within swagger yaml that has the type of "object" or could be described as a class in Java
type SchemaType struct {
	name        string
	path        string
	description string
	properties  []Property
}

// Constructors
func NewSchemaType(name string, path string) *SchemaType {
	schemaType := SchemaType {
		name: name,
		path: path,
	}
	return &schemaType
}

// Getters
func (schemaType SchemaType) GetName() string {
	return schemaType.name
}

func (schemaType SchemaType) GetPath() string {
	return schemaType.path
}

func (schemaType SchemaType) GetDescription() string {
	return schemaType.description
}

func (schemaType SchemaType) GetProperties() []Property {
	return schemaType.properties
}

//
// PROPERTY
//
type Property struct {
	name           string
	path           string
	description    string
	typeName       string
	possibleValues []string
	resolvedType   SchemaType
	cardinality    Cardinality
}

//Constructors
func NewProperty(name string, path string, typeName string, cardinality Cardinality) *Property {
	property := Property {
		name: name,
		path: path,
		typeName: typeName,
		cardinality: cardinality,
	}
	return &property
}

// Getters
func (prop Property) GetName() string {
	return prop.name
}

func (prop Property) GetPath() string {
	return prop.path
}

func (prop Property) GetDescription() string {
	return prop.description
}

func (prop Property) GetType() string {
	return prop.typeName
}

func (prop Property) GetPossibleValues() []string {
	return prop.possibleValues
}

func (prop Property) GetResolvedType() SchemaType {
	return prop.resolvedType
}

func (prop Property) GetCardinality() Cardinality {
	return prop.cardinality
}

func (prop Property) IsSetInConstructor() bool {
	isSetInConstructor := false
	if prop.cardinality.min > 0 {
		isSetInConstructor = true
	}
	return isSetInConstructor
}

func (prop Property) IsArrayOrList() bool {
	isArrayOrList := false
	if prop.cardinality.max > 1 {
		isArrayOrList = true
	}
	return isArrayOrList
}

func (prop Property) IsEnum() bool {
	isEnum := false
	if len(prop.possibleValues) > 1 {
		isEnum = true
	}
	return isEnum
}

func (prop Property) IsConstant() bool {
	isConstant := false
	if len(prop.possibleValues) == 1 {
		isConstant = true
	}
	return isConstant
}

// Setters
func (prop *Property) SetResolvedType(resolvedType SchemaType) {
	prop.resolvedType = resolvedType
}

//
// CARDINALITY
//
type Cardinality struct {
	min int
	max int
}

// Getters
func (cardinality Cardinality) GetMin() int {
	return cardinality.min
}

func (cardinality Cardinality) GetMax() int {
	return cardinality.max
}