package generator

func translateSchemaTypesToJavaPackage(schemaTypes map[string]*SchemaType, packageName string) (javaPackage *JavaPackage){
	javaPackage = NewJavaPackage(packageName)
	for _, schemaType := range schemaTypes {
		if schemaType.ownProperty.IsEnum() {
			enumValues := possibleValuesToEnumValues(schemaType.ownProperty.possibleValues)
			
			javaEnum := NewJavaEnum(schemaType.ownProperty.name, schemaType.ownProperty.description, enumValues, javaPackage)

			javaPackage.Enums[schemaType.ownProperty.name] = javaEnum
		} else {
			dataMembers, requiredMembers := retrieveDataMembersFromSchemaType(schemaType)
			
			javaClass := NewJavaClass(schemaType.name, schemaType.description, javaPackage, nil, dataMembers, requiredMembers)
			javaPackage.Classes[schemaType.name] = javaClass
		}
	}
	return javaPackage
}

func possibleValuesToEnumValues(possibleValues map[string]string) (enumValues []string) {
	for _, value := range possibleValues {
		enumValues = append(enumValues, value)
	}
	return enumValues
}

func retrieveDataMembersFromSchemaType(schemaType *SchemaType) (dataMembers []*DataMember, requiredMembers []*RequiredMember){
	for _, property := range schemaType.properties {
		var constVal string
		if property.IsConstant() {
			posVal := possibleValuesToEnumValues(property.GetPossibleValues())
			constVal = posVal[0]
		}
		dataMember := DataMember {
			Name: property.name,
			MemberType: propertyToJavaType(property),
			Description: property.description,
			ConstantVal: constVal,
		}
		dataMembers = append(dataMembers, &dataMember)
		
		if property.IsSetInConstructor() {
			requiredMember := RequiredMember {
				DataMember: &dataMember,
				IsFirst: len(requiredMembers) == 0,
			}
			requiredMembers = append(requiredMembers, &requiredMember)
		}
	}
	return dataMembers, requiredMembers
}

func propertyToJavaType(property *Property) string {
	javaType := ""
	if property.IsReferencing() || property.typeName == "object" || property.IsEnum() {
		javaType = property.resolvedType.name
	} else {
		if property.typeName == "string" {
			javaType = "String"
		} else if property.typeName == "integer" {
			javaType = "int"
		} else if property.typeName == "number" {
			javaType = "double"
		} else {
			javaType = property.typeName
		}
	}

	if property.IsCollection() {
		javaType += "[]"
	}

	return javaType
}