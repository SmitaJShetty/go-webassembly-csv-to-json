package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"syscall/js"
)

func main() {
	js.Global().Set("Process", fileProcessorFunc())
	<-make(chan bool)
}

//createConstructFromJSON put the json into a structure for later unmarshaling
func createConstructFromJSON(dataArray []byte) (map[string](map[string]interface{}), error) {
	var c map[string](map[string]interface{})
	err := json.Unmarshal(dataArray, &c)
	if err != nil {
		fmt.Println("error while unmarshalling:", err)
		return nil, err
	}
	return c, nil
}


func generateJSONFromCSV(fileContents string) (string, error) {
	//split up into header and body
	fileRows := strings.Split(fileContents, "\n")
	if len(fileRows) == 1 {
		return "", fmt.Errorf("no data in file")
	}

	header := fileRows[0]
	headerItems := strings.Split(header, `,`)
	body := fileRows[1:]

	//generate json from body
	var fileJSONArr bytes.Buffer
	fileJSONArr.WriteString("[")

	for _, row := range body {
		rowItems := strings.Split(row, `,`)
		rJSON := getRowJSON(rowItems, headerItems)
		fileJSONArr.WriteString(rJSON)
	}

	fileJSONArr.WriteString("]")
	return fileJSONArr.String(), nil
}

func getRowJSON(rowItems []string, headerItems []string) string {
	var rowJSON []string

	rowJSON = append(rowJSON, `{`)
	for i := 0; i < len(headerItems); i++ {
		s := `"` + headerItems[i] + `":"` + rowItems[i] + `",`
		rowJSON = append(rowJSON, s)
	}

	rowJSON = append(rowJSON, `},`)
	return strings.Join(rowJSON, ` `)
}

//createJSON marshalident the array
func createJSON(c string) ([]byte, error) {
	//var indentedJSON bytes.Buffer
	dataBytes := json.RawMessage(c)
	data, err := json.MarshalIndent(dataBytes, " 	", "\t")
	if err != nil {
		fmt.Println("err while indenting:", err)
		return nil, err
	}
	
	return data, nil
}

//Process processes file
func Process(fileContents string) ([]byte, error) {
	//generate json from csv
	jsonStr, jsonStrErr := generateJSONFromCSV(fileContents)
	if jsonStrErr != nil {
		return nil, jsonStrErr
	}

	fmt.Println("jsonstr:", jsonStr)
	createdJSON, createdJSONErr := createJSON(jsonStr)
	if createdJSONErr != nil {
		return nil, fmt.Errorf("err:%v", createdJSONErr)
	}

	return createdJSON, nil
}

func fileProcessorFunc() js.Func {
	fProcessFunc := js.FuncOf(
		func(this js.Value, args []js.Value) interface{} {
			if len(args) == 0 {
				return fmt.Errorf("Invalid number of arguments passed")
			}

			fileContents := args[0].String()
			data, err := Process(fileContents)
			if err != nil {
				return err
			}
			fmt.Println("formatted:", string(data))
			return string(data)
		})
	return fProcessFunc
}

/*
func test() js.Func {
	testFunc := js.FuncOf(
		func(this js.Value, args []js.Value) interface{} {
			if len(args) == 0 {
				return nil
			}

			fileContents := args[0].String()
			fileContentArray := strings.SplifileContents,"\n")
			arr:= generateResult(fileContentArray)
			return (arr)
		})
		return testFunc
}

func generateResult(arr []string)[]interface{}{
	var finalArr []interface{}
	for _, r:= range arr {
		finalArr = append(finalArr, r)
	}

	return finalArr
}

//createConstructArray load the other rows into array of the above struct
func createConstructArray(dataArray []string)([]ABC){
	var constrArr []ABC
	for _, row:= range dataArray{
		constrArr=append(constrArr, *createConstructItem(row))
	}
	return constrArr
}



func createConstructItem(row string) *ABC{
	rowItems:= strings.Split(row,",")

f len(rowItems)==0 {
		return nil
	}

	return &ABC{
		A: rowItems[0],
		B: rowItems[1],
		C: rowItems[2],
		D: rowItems[3],
		E: rowItems[5],
		F: rowItems[6],
		G: rowItems[6],
		H: rowItems[7],
	}
}


//ABC construct
type ABC struct {
	A string
	B string
	C string
	D string
	E string
	F string
	G string
	H string
}


//generateDynamicConstruct generate a dynamic struct from first row of data
func generateDynamicConstruct(header string)(ABC, error){
	headerItems:= strings.Split(header, ",")
	return ABC{
		A: headerItems[0],
		B: headerItems[1],
		C: headerItems[2],
		D: headerItems[3],
		E: headerItems[4],
		F: headerItems[5],
		G: headerItems[6],
		H: headerItems[7],
	}, nil
}


//getDataInJson gets data in json format
func getDataInJson(fileContents map[string](map[string]interface{})) (string, error) {
	data, err := json.Marshal(fileContents)
	if err != nil {
		return "", err
	}

	if data == nil {
		return "", fmt.Errorf("data was returned empty")
	}

	return string(data), nil
}

*/
