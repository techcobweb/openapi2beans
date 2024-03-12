package generator

func translateSchemaTypesToJavaPackage(schemaTypes map[string]*SchemaType, packageName string) (javaPackage JavaPackage){
	javaPackage.Name = packageName
	javaPackage.Classes = make(map[string]*JavaClass)
	javaPackage.Enums = make(map[string]*JavaEnum)
	for _, schemaType := range schemaTypes {
		if schemaType.ownProperty.IsEnum() {
			enumValues := possibleValuesToEnumValues(schemaType.ownProperty.possibleValues)
			javaEnum := JavaEnum {
				Name: schemaType.ownProperty.name,
				Description: schemaType.ownProperty.description,
				EnumValues: enumValues,
			}
			javaPackage.Enums[schemaType.ownProperty.name] = &javaEnum
		} else {
			dataMembers := []*DataMember{}
			requiredMembers := []*RequiredMember{}
			for _, property := range schemaType.properties {
				dataMember := DataMember {
					Name: property.name,
					MemberType: propertyToJavaType(property),
					Description: property.description,
				}
				dataMembers = append(dataMembers, &dataMember)
				if property.IsConstant() {
					// DO STUFF HERE
				} else {
					if property.IsSetInConstructor() {
						requiredMember := RequiredMember {
							DataMember: &dataMember,
							IsFirst: len(requiredMembers) == 0,
						}
						requiredMembers = append(requiredMembers, &requiredMember)
					}
				}
			}
			javaClass := NewJavaClass(schemaType.name, schemaType.description, &javaPackage, nil, dataMembers, requiredMembers)
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

func propertyToJavaType(property *Property) string {
	javaType := ""
	if property.IsReferencing() {
		javaType = property.resolvedType.name
	} else {
		if property.typeName == "string" {
			javaType = "String"
		} else {
			javaType = property.typeName
		}
	}

	if property.IsCollection() {
		javaType += "[]"
	}

	return javaType
}