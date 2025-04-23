package main

import (
	"encoding/json"
	"fmt"
	"gomaster/hw/mathslice"
	"gomaster/hw/toppackage/middlepackage/bottompackage/mathxxx"
	"math"
	"time"
)

const (
	one = iota*2 + 1 // укажите здесь формулу с iota
	three
	five
	seven
	nine
	eleven
)

func praktikum() {
	s := "Hello world!"
	myflag := true
	fmt.Println(s)
	fmt.Println(s[0])
	fmt.Println(string(s[0]))
	fmt.Println(myflag)
	pi := 3.1415
	fmt.Println(pi)
	pi, e := 3.14159, 2.71828
	fmt.Println(pi)
	fmt.Println(e)
	ver := "v0.0.1"
	var id int
	pi = 3.1415
	fmt.Println("ver =", ver, "id =", id, "pi =", pi)
	fmt.Println(one, three, five, seven, nine, eleven)
	day := "Sundayasd"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday":
		fmt.Println(day + " is a weekday")
		fallthrough
	case "Friday":
		fmt.Println("Start of the weekend")
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Unknown day")
	}

	fmt.Println("Введите год своего рождения:")
	var birthY int = 1900
	for birthY != 0 {
		birthY = 1900
		if _, err := fmt.Scan(&birthY); err != nil {
			fmt.Println("Неверный ввод, ошибка - !", err)
		}
		switch {
		case birthY < 1946 || birthY > 2025:
			fmt.Println("Введен некорректный год рождения!")
		case birthY < 1965:
			fmt.Println("Привет, бумер!")
		case birthY < 1981:
			fmt.Println("Привет, представитель X!")
		case birthY < 1997:
			fmt.Println("Привет, миллениал!")
		case birthY < 2013:
			fmt.Println("Привет, зумер!")
		case birthY < 2026:
			fmt.Println("Привет, альфа!")
		}
	}

	for i := 1; i <= 100; i++ {
		switch {
		case (i%15 == 0):
			fmt.Println("FizzBuzz")
		case (i%3 == 0):
			fmt.Println("Fizz")
		case (i%5 == 0):
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
func sliceAddReverse() {
	s := make([]int, 100)
	for i := range s {
		s[i] = i + 1
	}
	fmt.Println(s)

	s = append(s[:10], s[len(s)-10:]...)
	fmt.Println(s, len(s), cap(s))

	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		s[l], s[r] = s[r], s[l]
	}
	fmt.Println(s)
}
func mapTry() {
	const minP = 500
	prices := map[string]int{"хлеб": 60, "молоко": 100, "масло": 200, "колбаса": 500,
		"соль": 20, "огурцы": 200, "сыр": 600, "ветчина": 700,
		"буженина": 900, "помидоры": 250, "рыба": 300, "хамон": 1500}
	fmt.Print("Продкуты дороже ", minP, ":")
	for k, v := range prices {
		if v >= minP {
			fmt.Print(" ", k)
		}
	}
	fmt.Println(".")
	order := []string{"хлеб", "буженина", "сыр", "огурцы"}
	var sum int
	for _, product := range order {
		if v, ok := prices[product]; ok {
			sum += v
		} else {
			fmt.Println("Продукт ", product, " отсутствует!")
		}
	}
	fmt.Println("Итого: ", sum)
}

type ipair struct {
	i1 int
	i2 int
}

func allKSumPairs(A []int, k int) []ipair {
	res := make([]ipair, 0, len(A))
	remaindersMap := make(map[int][]int, len(A))
	for i := range A {
		key := k - A[i]
		remaindersMap[key] = append(remaindersMap[key], i)
	}
	for i, v := range A {
		if val, ok := remaindersMap[v]; ok {
			for _, index := range val {
				if index > i {
					res = append(res, ipair{i, index})
				}
			}
		}
	}
	return res
}

func removeDupsFromSlice(sa []string) []string {
	smap := make(map[string]int, len(sa))
	res := make([]string, 0, len(sa))
	for _, s := range sa {
		if _, ok := smap[s]; !ok {
			// 1st time met this key -> add to res and add to map
			res = append(res, s)
			smap[s] = 1
		}
	}
	return res
}

type Person struct {
	Name        string    `json:"Имя"`
	Email       string    `json:"Почта"`
	DateOfBirth time.Time `json:"-"`
}

const rawResp = `
{
    "header": {
        "code": 0,
        "message": ""
    },
    "data": [{
        "type": "user",
        "id": 100,
        "attributes": {
            "email": "bob@yandex.ru",
            "article_ids": [10, 11, 12]
        }
    }]
}
`

type (
	Response struct {
		Header RespHeader `json:"header"`
		Data   RespData   `json:"data,omitempty"`
	}

	RespHeader struct {
		Code    int    `json:"code"`
		Message string `json:"message,omitempty"`
	}

	RespData []RespDataItem

	RespDataItem struct {
		Type       string           `json:"type"`
		Id         int              `json:"id"`
		Attributes RespDataItemAttr `json:"attributes"`
	}

	RespDataItemAttr struct {
		Email      string `json:"email"`
		ArticleIds []int  `json:"article_ids"`
	}
)

func decodeRawJSON() {
	resp := Response{}
	if err := json.Unmarshal([]byte(rawResp), &resp); err != nil {
		fmt.Printf("JSON Unmarshall error(%v) occurred during process resp: %+v\n", err, resp)
	} else {
		fmt.Printf("%+v\n", resp)
	}
}

type figures int

const (
	square   figures = iota // квадрат
	circle                  // круг
	triangle                // равносторонний треугольник
)

func area(f figures) (func(float64) float64, bool) {
	switch f {
	case square:
		res := func(x float64) float64 {
			return x * x
		}
		return res, true
	case triangle:
		res := func(x float64) float64 {
			return x * x / 2
		}
		return res, true
	case circle:
		res := func(x float64) float64 {
			return math.Pi * x * x / 2
		}
		return res, true
	default:
		res := func(x float64) float64 {
			return 0.0
		}
		return res, false
	}
}

// defer train
var Global = 5

func tryDefer() {
	//defer to restore global
	defer func(OldGlobal int) {
		Global = OldGlobal
	}(Global)
	Global = 4
	fmt.Println(Global)
}

// try my slice pkg
func checkSlicePkgFuncs() {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("Slice s = %v sum = %d\n", s, mathslice.SumSlice(s))
	square := func(val int) int { return val * val }

	mathslice.MapSlice(s, square)
	fmt.Printf("Squared slice s = %v\n", s)

	fold := func(v1 int, v2 int) int {
		// just return the greater one
		if v1 > v2 {
			return v1
		}
		return v2
	}
	initVal := 32
	fmt.Printf("The greatest number out of slice(%v) and number %d is %d\n",
		s, initVal, mathslice.FoldSlice(s, fold, initVal))
}

func main() {
	// find all pairs with sum k
	k := 3
	A := []int{1, 2, 1, 2, 1, 2}
	res := allKSumPairs(A, k)
	fmt.Println(res)

	//remove dups from slice
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}
	input = removeDupsFromSlice(input)
	fmt.Println(input)

	//encode json
	u, err := json.Marshal(Person{Name: "Алекс", Email: "alex@yandex.ru", DateOfBirth: time.Now()})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", string(u))

	//decode json
	decodeRawJSON()

	//get figure area
	myFs := []figures{triangle, circle, square}
	for _, f := range myFs {
		ar, ok := area(f)
		if !ok {
			fmt.Println("Неизвестная фигура.")
		}
		fArea := ar(2.0)
		fmt.Println("Площадь фигуры = ", fArea)
	}

	//try defer
	tryDefer()
	fmt.Println(Global)

	// slice pkg func
	checkSlicePkgFuncs()

	// add new pkg mathxxx adn test it
	if sum := mathxxx.AddInts(3, 2); sum != 5 {
		panic(fmt.Sprintf("Sum must be equal 5; got %d", sum))
	} else {
		fmt.Println("Sum equal 5. Well done!")
	}

}
