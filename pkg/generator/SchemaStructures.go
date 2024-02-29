package generator

import "strings"

// SCHEMA TYPE
//
// SchemaType describes a schema part within swagger yaml that has the type of "object" or could be described as a class in Java
type SchemaType struct {
	name        string
	description string
	ownProperty *Property
	properties  map[string]*Property
}

// Constructors
func NewSchemaType(name string, description string, ownProperty *Property, properties map[string]*Property) *SchemaType {
	schemaType := SchemaType{
		name:        name,
		description: description,
		ownProperty: ownProperty,
	}
	schemaType.properties = make(map[string]*Property)
	schemaType.SetProperties(properties)
	return &schemaType
}

// Getters
func (schemaType SchemaType) GetName() string {
	return schemaType.name
}

func (schemaType SchemaType) GetDescription() string {
	return schemaType.description
}

func (schemaType SchemaType) GetProperties() map[string]*Property {
	return schemaType.properties
}

// Setters
func (schemaType *SchemaType) SetProperties(properties map[string]*Property) {
	if properties != nil {
		schemaTypePath := schemaType.ownProperty.path
		splitSchemaTypePath := strings.Split(schemaTypePath, "/")
		for _, property := range properties {
			match := true
			splitPropertyPath := strings.Split(property.GetPath(), "/")
			if len(splitPropertyPath) - 1 == len(splitSchemaTypePath) {
				for pos, element := range splitPropertyPath[:len(splitPropertyPath)-1] {
					if element != splitSchemaTypePath[pos] {
						match = false
					}
				}
				if match {
					schemaType.properties[property.path] = property
				}
			}
		}
	}
}

// PROPERTY
type Property struct {
	name           string
	path           string
	description    string
	typeName       string
	possibleValues []string
	resolvedType   *SchemaType
	cardinality    Cardinality
}

// Constructors
func NewProperty(name string, path string, description string, typeName string, possibleValues []string, resolvedType *SchemaType, cardinality Cardinality) *Property {
	property := Property{
		name:           name,
		path:           path,
		description:    description,
		typeName:       typeName,
		possibleValues: possibleValues,
		resolvedType:   resolvedType,
		cardinality:    cardinality,
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

func (prop Property) GetResolvedType() *SchemaType {
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

func (prop Property) IsReferencing() bool {
	isReferencing := false
	if strings.HasPrefix(prop.typeName, "$ref:") {
		isReferencing = true
	}
	return isReferencing
}

// Setters
func (prop *Property) SetResolvedType(resolvedType *SchemaType) {
	prop.resolvedType = resolvedType
}

func (prop *Property) Resolve(resolvingProperty *Property) {
	prop.description = resolvingProperty.GetDescription()
	prop.typeName = resolvingProperty.GetType()
	prop.possibleValues = resolvingProperty.GetPossibleValues()
	prop.resolvedType = resolvingProperty.GetResolvedType()
	prop.cardinality = resolvingProperty.GetCardinality()
}

// CARDINALITY
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
