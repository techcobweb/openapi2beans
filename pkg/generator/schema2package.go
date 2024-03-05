package generator

func translateSchemaTypesToJavaPackage(schemaTypes map[string]*SchemaType, packageName string) (javaPackage JavaPackage){
	javaPackage.Name = packageName
	javaPackage.classes = make(map[string]*JavaClass)
	for _, schemaType := range schemaTypes {
		dataMembers := []*DataMember{}
		requiredMembers := []*RequiredMember{}
		for _, property := range schemaType.properties {
			dataMember := DataMember {
				Name: property.name,
				MemberType: propertyToJavaType(property),
				Description: property.description,
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
		javaClass := NewJavaClass(schemaType.name, schemaType.description, nil, &javaPackage, nil, dataMembers, requiredMembers)
		javaPackage.classes[schemaType.name] = javaClass
	}
	return javaPackage
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