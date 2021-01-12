package main

import "fmt"
import "strconv"
import "encoding/json"
import "strings"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution       = make(map[string]int, len(users))
	a, e, i, o, u, sum int
)

func GetResult([]string) string {
	//insert code here
	// access each name of users array - global
	for _, name := range users {
		// sum the coins according to occurence of each vowel and assign to map | check for >10
		a, e, i, o, u, sum = 0, 0, 0, 0, 0, 0
		a = strings.Count(name, "a") + strings.Count(name, "A")
		e = strings.Count(name, "e") + strings.Count(name, "E")
		i = strings.Count(name, "i") + strings.Count(name, "I")
		o = strings.Count(name, "o") + strings.Count(name, "O")
		u = strings.Count(name, "u") + strings.Count(name, "U")
		sum = a + e + 2*i + 3*o + 4*u
		if sum > 10 {
			sum = 10
		}
		distribution[name] = sum
	}

	fmt.Println(distribution)
	fmt.Println("Coins left:", coins)
	return strconv.Itoa(distribution["Matthew"]) + " " + strconv.Itoa(distribution["Sarah"]) + " " + strconv.Itoa(distribution["Augustus"]) + " " + strconv.Itoa(distribution["Heidi"]) + " " + strconv.Itoa(distribution["Emilie"]) + " " + strconv.Itoa(distribution["Peter"]) + " " + strconv.Itoa(distribution["Giana"]) + " " + strconv.Itoa(distribution["Adriano"]) + " " + strconv.Itoa(distribution["Aaron"]) + " " + strconv.Itoa(distribution["Elizabeth"])
}
