package main

import "fmt"

func main() {

	var mymap1 map[string]string // only defining map but not instantiating a map
	//var mymap2 map[string]interface{} // map[string]any
	if mymap1 == nil {
		mymap1 = make(map[string]string, 10)
	}
	// make is only used to create slice, map and channel
	mymap1["522000"] = "Guntur-1"
	mymap1["522001"] = "Guntur-2"
	mymap1["560086"] = "Yeshvantpur"
	mymap1["560096"] = "RajajiNagar"

	delete(mymap1, "522000")

	fmt.Println(mymap1)

	v := mymap1["522000"]
	println(v)
	v = mymap1["2342342"] // ""
	println(v)

	mymap2 := make(map[string]int)
	mymap2["banana"] = 10
	mymap2["apple"] = 5
	mymap2["rambutan"] = 0
	mymap2["mango"] = 7

	v1, ok := mymap2["banana"]
	if !ok {
		println("no map with the key", ok)
	} else {
		println(v1, ok)
	}
	v1, ok = mymap2["rambutan"]
	if !ok {
		println("no map with the key", ok)
	} else {
		println(v1, ok)
	}
	v1, ok = mymap2["peach"]
	if !ok {
		println("no map with the key", ok)
	} else {
		println(v1, ok)
	}

	for k, v := range mymap2 {
		fmt.Println("Key:", k, "Value:", v)
	}

}

// map : key-value pair
// keys must be unique
// Bucket: An array of buckets
// 		each bucket contains multiple key-values, it could be a linked list or a slice of elements
// hash function: it caliculates a hash value from the key that is provided. Hash Function returns a value
// map key-values are not orders while retriving
// maps is not thread safe
