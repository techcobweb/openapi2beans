package generator

import (
	"regexp"
	"sort"
	"strings"
)

func translateSchemaTypesToJavaPackage(schemaTypes map[string]*SchemaType, packageName string) (javaPackage *JavaPackage){
	javaPackage = NewJavaPackage(packageName)
	for _, schemaType := range schemaTypes {
		description := strings.Split(schemaType.description, "\n")
		if len(description) == 1 && description[0] == "" {
			description = nil
		} else if len(description) > 1 {
			description = description[:len(description)-2]
		}
		
		if schemaType.ownProperty.IsEnum() {
			enumValues := possibleValuesToEnumValues(schemaType.ownProperty.possibleValues)
			
			javaEnum := NewJavaEnum(convertToCamelCase(schemaType.ownProperty.name), description, enumValues, javaPackage)

			javaPackage.Enums[convertToCamelCase(schemaType.ownProperty.name)] = javaEnum
		} else {
			dataMembers, requiredMembers, constantDataMembers := retrieveDataMembersFromSchemaType(schemaType)
			
			javaClass := NewJavaClass(convertToCamelCase(schemaType.name), description, javaPackage, nil, dataMembers, requiredMembers, constantDataMembers)
			javaPackage.Classes[convertToCamelCase(schemaType.name)] = javaClass
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

func retrieveDataMembersFromSchemaType(schemaType *SchemaType) (dataMembers []*DataMember, requiredMembers []*RequiredMember, constantDataMembers []*DataMember){
	for _, property := range schemaType.properties {
		var constVal string
		name := property.name
		description := strings.Split(property.description, "\n")
		if len(description) == 1 && description[0] == "" {
			description = nil
		} else if len(description) > 1 {
			description = description[:len(description)-2]
		}
		if property.IsConstant() {
			posVal := possibleValuesToEnumValues(property.GetPossibleValues())
			name = convertToConstName(name)
			constVal = convertConstValueToJavaReadable(posVal[0], property.typeName)

			constDataMember := DataMember {
				Name: name,
				CamelCaseName: convertToCamelCase(name),
				MemberType: propertyToJavaType(property),
				Description: description,
				ConstantVal: constVal,
			}

			constantDataMembers = append(constantDataMembers, &constDataMember)

		} else {

			dataMember := DataMember {
				Name: name,
				CamelCaseName: convertToCamelCase(name),
				MemberType: propertyToJavaType(property),
				Description: description,
				ConstantVal: constVal,
			}
			dataMembers = append(dataMembers, &dataMember)
				
			if property.IsSetInConstructor() {
				requiredMember := RequiredMember {
					DataMember: &dataMember,
				}
				requiredMembers = append(requiredMembers, &requiredMember)
			}
		}
		
	}
	sort.SliceStable(dataMembers, func (i int, j int) bool {return isDataMemberLessThanComparison(dataMembers[i], dataMembers[j])})
	sort.SliceStable(requiredMembers, func (i int, j int) bool {return isDataMemberLessThanComparison(requiredMembers[i].DataMember, requiredMembers[j].DataMember)})
	sort.SliceStable(constantDataMembers, func (i int, j int) bool {return isDataMemberLessThanComparison(constantDataMembers[i], constantDataMembers[j])})
	if requiredMembers != nil {
		requiredMembers[0].IsFirst = true
	}
	return dataMembers, requiredMembers, constantDataMembers
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
		}else if property.typeName == "" {
			javaType = "Object"
		} else {
			javaType = property.typeName
		}
	}

	if property.IsCollection() {
		dimensions := property.cardinality.max / 128
		for range dimensions {
			javaType += "[]"
		}
	}

	return javaType
}

func convertToCamelCase(name string) string {
	initialLetter := name[0]
	camelCaseName := strings.ToUpper(string(initialLetter)) + name[1:]
	return camelCaseName
}

func convertToConstName(name string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")

	constName := matchFirstCap.ReplaceAllString(name, "${1}_${2}")
    constName  = matchAllCap.ReplaceAllString(constName, "${1}_${2}")

    return strings.ToUpper(constName)
}

func convertConstValueToJavaReadable(constVal string, constType string) string {
	if constType == "string" {
		constVal = "\"" + constVal + "\""
	}
	return constVal
}

func isDataMemberLessThanComparison(dataMember *DataMember, comparisonMember *DataMember) bool {
	less := true
	switch memberType := dataMember.MemberType; {
	case strings.Contains(memberType, "boolean"):
		switch comparisonMemberTpye := comparisonMember.MemberType; {
		case strings.Contains(comparisonMemberTpye, "boolean"):
			less = dataMember.Name > comparisonMember.Name
		default:
			less = true
		}
	case strings.Contains(memberType, "int"):
		switch comparisonMember.MemberType {
		case "boolean": 
			less = false
		case "int":
			less = dataMember.Name > comparisonMember.Name
		default:
			less = true
		}
	case strings.Contains(memberType, "double"):
		switch comparisonMemberType := comparisonMember.MemberType; {
		case strings.Contains(comparisonMemberType, "boolean"), strings.Contains(comparisonMemberType, "int"): 
			less = false
		case strings.Contains(comparisonMemberType, "double"):
			less = dataMember.Name > comparisonMember.Name
		default:
			less = true
		}
	case strings.Contains(memberType, "String"):
		switch comparisonMemberType := comparisonMember.MemberType; {
		case strings.Contains(comparisonMemberType, "boolean"), strings.Contains(comparisonMemberType, "int"), strings.Contains(comparisonMemberType, "double"): 
			less = false
		case strings.Contains(comparisonMemberType, "String"):
			less = dataMember.Name > comparisonMember.Name
		default:
			less = true
		}
	default:
		if dataMember.MemberType == comparisonMember.MemberType {
			less = dataMember.Name > comparisonMember.Name
		} else {
			less = dataMember.MemberType > comparisonMember.MemberType
		}
	}
	return less
}