package main

import "fmt"

func DeleteMap(m map[int]string, k int) {
	delete(m, k)
}

func main() {
	var map1 map[int]string = map[int]string{1: "mike", 2: "yorick"}
	fmt.Println(map1)

	map1[1] = "mike2"
	map1[2] = "yorick2"
	fmt.Println(map1)

	for k, v := range map1 {
		fmt.Println("k: ", k, " v:", v)
	}

	value, ok := map1[1]
	fmt.Println("value=", value, " ok=", ok)
	value2, ok2 := map1[3]
	fmt.Println("value2=", value2, " ok2=", ok2)

	delete(map1, 2)
	fmt.Println(map1)

	map1[2] = "tom2"
	fmt.Println(map1)
	DeleteMap(map1, 2)
	fmt.Println(map1)
}
