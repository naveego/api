package pipeline

import (
	"hash/crc32"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestShaper(t *testing.T) {

	testCases := []struct {
		name                  string
		keyNames              []string
		expectedKeyHashString string
		data                  map[string]interface{}
		expectedShape         []string
		expectedHashString    string
		expectedError         string
	}{
		{
			"Given a simple data structure",
			[]string{"id"},
			"id",
			map[string]interface{}{"id": 1, "name": "John", "active": true},
			[]string{"active:bool", "id:number", "name:string"},
			"active:bool,id:number,name:string",
			"",
		},
		{
			"Given a simple data structure with null property value",
			[]string{"id"},
			"id",
			map[string]interface{}{"id": 1, "name": nil, "active": true},
			[]string{"active:bool", "id:number", "name:unknown"},
			"active:bool,id:number,name:unknown",
			"",
		},
		{
			"Given a simple data structure with empty string property value",
			[]string{"id"},
			"id",
			map[string]interface{}{"id": 1, "name": "", "active": true},
			[]string{"active:bool", "id:number", "name:string"},
			"active:bool,id:number,name:string",
			"",
		},
		{
			"Given a simple data structure with mixed case property names",
			[]string{"id"},
			"id",
			map[string]interface{}{"id": 1, "Name": "John", "active": true},
			[]string{"active:bool", "id:number", "Name:string"},
			"active:bool,id:number,name:string",
			"",
		},
		{
			"Given a simple data structure with spaces in property names",
			[]string{"id"},
			"id",
			map[string]interface{}{"id": 1, "First Name": "John", "active": true},
			[]string{"active:bool", "First Name:string", "id:number"},
			"active:bool,first name:string,id:number",
			"",
		},
		{
			"Given a simple data structure with # in property names",
			[]string{"id"},
			"id",
			map[string]interface{}{"Job #": 1, "First Name": "John", "active": true},
			[]string{"active:bool", "First Name:string", "Job #:number"},
			"active:bool,first name:string,job #:number",
			"",
		},
		{
			"Given a simple data structure with - in property names",
			[]string{"id"},
			"id",
			map[string]interface{}{"id": 1, "First - Name": "John", "active": true},
			[]string{"active:bool", "First - Name:string", "id:number"},
			"active:bool,first - name:string,id:number",
			"",
		},
		{
			"Given a simple data structure with [ ] in property names",
			[]string{"id"},
			"id",
			map[string]interface{}{"id": 1, "First [Name]": "John", "active": true},
			[]string{"active:bool", "First [Name]:string", "id:number"},
			"active:bool,first [name]:string,id:number",
			"",
		},
		{
			"Given a simple data structure with / in property names",
			[]string{"id"},
			"id",
			map[string]interface{}{"id": 1, "First /Name": "John", "active": true},
			[]string{"active:bool", "First /Name:string", "id:number"},
			"active:bool,first /name:string,id:number",
			"",
		},
		{
			"Given a simple data structure long complex property name",
			[]string{"id"},
			"id",
			map[string]interface{}{"id": 1, "Intermediate - Start Date (Estimated) [Well Type]": "John", "active": true},
			[]string{"active:bool", "id:number", "Intermediate - Start Date (Estimated) [Well Type]:string"},
			"active:bool,id:number,intermediate - start date (estimated) [well type]:string",
			"",
		},
		{
			"Given a simple data structure with ? in property names",
			[]string{"id"},
			"id",
			map[string]interface{}{"id": 1, "First Name?": "John", "active": true},
			[]string{"active:bool", "First Name?:string", "id:number"},
			"active:bool,first name?:string,id:number",
			"",
		},
		{
			"Given a simple data structure with ( ) in property names",
			[]string{"id"},
			"id",
			map[string]interface{}{"id": 1, "First (Name)": "John", "active": true},
			[]string{"active:bool", "First (Name):string", "id:number"},
			"active:bool,first (name):string,id:number",
			"",
		},
		{
			"Given a complex data structure",
			[]string{"id"},
			"id",
			map[string]interface{}{"id": 1, "Name": "John", "active": true, "company": map[string]interface{}{"name": "test", "address": "123 Main street"}},
			[]string{"active:bool", "company:object", "company.address:string", "company.name:string", "id:number", "Name:string"},
			"active:bool,company:object,company.address:string,company.name:string,id:number,name:string",
			"",
		},
		{
			"Given a complex data structure, with multiple nested objects",
			[]string{"id"},
			"id",
			map[string]interface{}{"id": 1, "Name": "John", "active": true, "company": map[string]interface{}{"name": "test", "address": "123 Main street"}, "parents": map[string]interface{}{"mom": "test", "momsAge": 45, "dad": "123 Main street", "dadsAge": 46}},
			[]string{"active:bool", "company:object", "company.address:string", "company.name:string", "id:number", "Name:string", "parents:object", "parents.dad:string", "parents.dadsAge:number", "parents.mom:string", "parents.momsAge:number"},
			"active:bool,company:object,company.address:string,company.name:string,id:number,name:string,parents:object,parents.dad:string,parents.dadsage:number,parents.mom:string,parents.momsage:number",
			"",
		},
		{
			"Given a data structure that contains a date string that conforms with RFC 3339 and has an offset",
			[]string{"id"},
			"id",
			map[string]interface{}{"id": 1, "Name": "John", "dateOfBirth": "1981-01-01T12:30:00-07:00"},
			[]string{"dateOfBirth:date", "id:number", "Name:string"},
			"dateofbirth:date,id:number,name:string",
			"",
		},
		{
			"Given a data structure that contains a date string that conforms with RFC 3339 and uses Z",
			[]string{"id"},
			"id",
			map[string]interface{}{"id": 1, "Name": "John", "dateOfBirth": "1981-01-01T12:30:00Z"},
			[]string{"dateOfBirth:date", "id:number", "Name:string"},
			"dateofbirth:date,id:number,name:string",
			"",
		},
		{
			"Given a data structure that contains a date string that conforms with RFC 3339 and has nano seconds with offset",
			[]string{"id"},
			"id",
			map[string]interface{}{"id": 1, "Name": "John", "dateOfBirth": "1981-01-01T12:30:00.323-05:00"},
			[]string{"dateOfBirth:date", "id:number", "Name:string"},
			"dateofbirth:date,id:number,name:string",
			"",
		},
		{
			"Given a data structure that contains a date string that conforms with RFC 3339 and has nano seconds with Z",
			[]string{"id"},
			"id",
			map[string]interface{}{"id": 1, "Name": "John", "dateOfBirth": "1981-01-01T12:30:00.323Z"},
			[]string{"dateOfBirth:date", "id:number", "Name:string"},
			"dateofbirth:date,id:number,name:string",
			"",
		},
		{
			"Given a field name containing a : should return error",
			[]string{"id"},
			"id",
			map[string]interface{}{"id": 1, "Name": "John", "date:of:birth": "1981-01-01T12:30:00.323Z"},
			nil,
			"",
			"Invalid character found in property 'date:of:birth'.",
		},
		{
			"Given a field name containing a , should return error",
			[]string{"id"},
			"id",
			map[string]interface{}{"id": 1, "Name": "John", "date,birth": "1981-01-01T12:30:00.323Z"},
			nil,
			"",
			"Invalid character found in property 'date,birth'.",
		},
	}

	for _, testCase := range testCases {
		Convey(testCase.name, t, func() {

			shape, err := NewShaper().GetShape(testCase.keyNames, testCase.data)
			if err != nil && testCase.expectedError == "" {
				t.Fatal("Did not expect error: ", err)
			}

			if testCase.expectedError != "" {
				if err == nil {
					t.Fatal("Expected error but did not get one")
				}

				Convey("Should return error message", func() {
					So(err.Error(), ShouldEqual, testCase.expectedError)
				})
				return
			}

			expectedKeyHash := doCrc([]byte(testCase.expectedKeyHashString))
			expectedHash := doCrc([]byte(testCase.expectedHashString))

			Convey("Should set the key names property", func() {
				So(shape.KeyNames, ShouldResemble, testCase.keyNames)
			})

			Convey("Should generate the correct shape properties", func() {
				So(shape.Properties, ShouldResemble, testCase.expectedShape)
			})

			Convey("Should generate the correct key names hash", func() {
				So(shape.KeyNamesHash, ShouldEqual, expectedKeyHash)
			})

			Convey("Should generate the correct property hash", func() {
				So(shape.PropertyHash, ShouldEqual, expectedHash)
			})

		})
	}
}

func doCrc(data []byte) uint32 {
	crc := crc32.New(castagnoliTable)
	crc.Write(data)
	return crc.Sum32()
}
