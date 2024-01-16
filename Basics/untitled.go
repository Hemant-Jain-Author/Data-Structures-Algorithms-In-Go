
func IncrementPassByValue(x int) {
	x++
}

func IncrementPassByPointer(ptr *int) {
	(*ptr)++
}

func main1() {
	i := 10
	fmt.Println("Value of i before increment is :  ", i)
	IncrementPassByValue(i)
	fmt.Println("Value of i after increment is :  ", i)

	fmt.Println("Value of i before increment is :  ", i)
	IncrementPassByPointer(&i)
	fmt.Println("Value of i after increment is :  ", i)
}

type coord struct {
	x int
	y int
}



func ArrayExample() {
	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	PrintSlice(arr[:])
}

func SliceExample() {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	PrintSlice(s)
}

func PrintSlice(data []int) {
	count := len(data)
	fmt.Print("Values stored are : ")
	for i := 0; i < count; i++ {
		fmt.Print(" ", data[i])
	}
	fmt.Println()
}


func VariableExample() {
	var1 := 100
	var2 := 200
	var3 := var1 + var2
	fmt.Println("Adding ", var1, " and ", var2, " will give ", var3)
}

func VectorExample() {
	var data [10]int
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	PrintSlice(data[:])
}

func twoDArrayExample() {
	var data [4][2]int
	count := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			count++
			data[i][j] = count
		}
	}
	print2DArray(data, 4, 2)
}

func print2DArray(data [4][2]int, R int, C int) {
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			fmt.Println(" ", data[i][j])
		}
	}
}



func main6() {
	point := &coord{10, 10}
	fmt.Println("X axis coord value is  ", point.x)
	fmt.Println("Y axis coord value is  ", point.y)
}

func (s student) GetName() string {
	return s.name
}

func (s student) GetRollNo() int {
	return s.rollNo
}

type student struct {
	rollNo int
	name   string
}

func main7() {
	stud := student{1, "johny"}
	fmt.Println(stud)
	fmt.Println("Student name ::", stud.name) // Accessing inner fields.

	pstud := &stud
	fmt.Println("Student name ::", pstud.name) // Accessing inner fields.

	fmt.Println(student{rollNo: 2, name: "Ann"}) // Named initialization.
	fmt.Println(student{name: "Ann", rollNo: 2}) // Order does not matter.
	fmt.Println(student{name: "Alice"})          // Default initialization of rollNo.
}

/*
{1 johny}
Student name :: johny
Student name :: johny
{2 Ann}
{2 Ann}
{0 Alice}
*/


func Sum(num1 int, num2 int) int {
	result := num1 + num2
	return result
}

func main8() {
	// calling a function to calculate sum
	result := Sum(10, 20)
	fmt.Println("Sum is : ", result)
}

func main9() {
	// local variable definition
	x := 10
	y := 20
	// calling a function to find sum
	result := Sum(x, y)
	fmt.Println("Sum is : ", result)
}
