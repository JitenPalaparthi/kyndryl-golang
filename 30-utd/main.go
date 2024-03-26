package main

import "fmt"

func main() {

	var mymap1 MyMap
	mymap1 = make(MyMap, 10) // make can understand it

	mymap1["name"] = "Jiten"
	mymap1["age"] = 38
	mymap1["salary"] = 10000.34
	mymap1["married"] = true

	keys1 := mymap1.GetKeys()
	values1 := mymap1.GetValues()
	fmt.Println("Keys1:", keys1)
	fmt.Println("Values1:", values1)

	map2 := make(map[string]any, 5)
	map2["name"] = "Jiten"
	map2["age"] = 38
	map2["salary"] = 10000.34
	map2["married"] = true
	keys2 := MyMap(map2).GetKeys()
	values2 := MyMap(map2).GetValues()
	fmt.Println("Keys2:", keys2)
	fmt.Println("Values2:", values2)

}

type MyMap map[string]any // extending the existing map

func (m MyMap) GetKeys() []string {
	keys := make([]string, len(m))
	index := 0
	for k, _ := range m {
		keys[index] = k
		index++
	}
	return keys
}

func (m MyMap) GetValues() []any {
	values := make([]any, len(m))
	index := 0
	for _, v := range m {
		values[index] = v
		index++
	}
	return values
}
