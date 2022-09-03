package main

import (
	"fmt"
	"math"
)

func main0() {
    q := new(Queue)
    for i := 1; i <= 5; i++ {
        q.Add(i)
    }
    for i := 1; i <= 5; i++ {
        fmt.Print(q.Remove(), " ")
    }
    fmt.Println()
}

func josephus(n int, k int) int {
    que := new(Queue)
    for i := 0; i < n; i++ {
        que.Add(i + 1);
    }

    for (que.Len() > 1) {
        for i := 0; i < k - 1; i++ {
            que.Add(que.Remove());
        }
        que.Remove();// Kth person executed.
    }
    return que.Front().(int);
}

// Testing code.
func main1()  {
    fmt.Println("Position :", josephus(11, 5));
}


func  CircularTour( arr [][2]int, n int) int{
    for i := 0; i < n; i++ {
        total := 0;
        found := true;
        for j := 0; j < n; j++ {
            total += (arr[(i + j) % n][0] - arr[(i + j) % n][1]);
            if (total < 0) {
                found = false;
                break;
            }
        }
        if (found){
            return i;
        }
    }
    return -1;
}

func  CircularTour2(arr [][2]int, n int)  int {
    que := new(Queue)
    nextPump := 0
    prevPump := 0
    count := 0
    petrol := 0

    for que.Len() != n {
        for petrol >= 0 && que.Len() != n {
            que.Add(nextPump)
            petrol += (arr[nextPump][0] - arr[nextPump][1])
            nextPump = (nextPump + 1) % n
        }
    
        for petrol < 0 && que.Len() > 0 {
            prevPump = que.Remove().(int)
            petrol -= (arr[prevPump][0] - arr[prevPump][1])
        }

        count += 1
        if count == n{
            return -1
        }
    }

    if petrol >= 0 {

        return que.Remove().(int)
    } else {

        return -1
    }
}

func main2(){
    tour := [][2]int{{8, 6}, {1, 4}, {7, 6}}
    fmt.Println("Circular Tour :" , CircularTour(tour, 3))
    fmt.Println("Circular Tour :" , CircularTour2(tour, 3))    
}

/*
Circular Tour : 2
*/

func convertXY(src int, dst int) int{
    que := new(Queue)
    visited := make(map[int]int)
    a := []int{src, 0}
    que.Add(a)
    ok := false

    for que.Len() != 0{
        node := que.Remove().([]int)
        visited[node[0]] = 1
        value := node[0]
        steps := node[1]
        if value == dst {
            return steps
        }
        _, ok =  visited[value*2]
        if value < dst && !ok{
            que.Add([]int{value*2, steps+1})
        }
        _, ok =  visited[value-1]
        if value > 0 && !ok{
            que.Add([]int{value-1, steps+1})
        }
    }
    return -1
}

func main3(){
    fmt.Println("Steps to convert 2 to 7 :", convertXY(2, 7))    
}

/*
Steps to convert 2 to 7 : 3
*/
func maxSlidingWindows( arr []int, k int) {
    size := len(arr)
    for i := 0; i < size - k + 1; i++ {
        max := arr[i];
        for j := 1; j < k; j++ {
            if( max < arr[i + j]){
                max = arr[i + j];
            }            
        }
        fmt.Print(max, " ");
    }
}

func maxSlidingWindows2(arr []int, k int){
    size := len(arr)
    que := new(Queue)
    i := 0
    for i < size {
       // Remove out of range elements
        if que.Len() > 0 && que.Front().(int) <= i - k {
            que.Remove()
        }
        // Remove smaller values at left.
        for que.Len() > 0 && arr[que.Back().(int)] <= arr[i] {
            que.RemoveBack()
        }

        que.Add(i)
        // window of size k
        if i >= (k - 1) {
            fmt.Print(arr[que.Front().(int)], " ")
        }

        i += 1
    }
}

func main4() {
    arr := []int{11, 2, 75, 92, 59, 90, 55}
    k := 3
    maxSlidingWindows(arr, k)
    fmt.Println()
    maxSlidingWindows2(arr, k)
    fmt.Println()
}

/*
Max Sliding Windows : 75 92 92 92 90 
*/

func minOfMaxSlidingWindows(arr []int, k int) {
    size := len(arr)
    que := new(Queue)
    minVal := math.MaxInt32
    i := 0
    for i < size {
        // Remove out of range elements
        if que.Len() > 0 && que.Front().(int) <= i - k {
            que.Remove()
        }
        // Remove smaller values at left.
        for que.Len() > 0 && arr[que.Back().(int)] <= arr[i] {
            que.RemoveBack()
        }
        que.Add(i)
        // window of size k
        if i >= (k - 1) && minVal > arr[que.Front().(int)] {
            minVal = arr[que.Front().(int)]
        }
        i += 1
    }
    fmt.Println("Min of max is : ", minVal)
}

func main5() {
    arr := []int{11, 2, 75, 92, 59, 90, 55}
    k := 3
    minOfMaxSlidingWindows(arr, k)
}

//Min of max is :  75

func maxOfMinSlidingWindows(arr []int, k int) {
    size := len(arr)
    que := new(Queue)
    maxVal := math.MinInt32
    i := 0
    for i < size {
        // Remove out of range elements
        if que.Len() > 0 && que.Front().(int) <= i - k {
            que.Remove()
        }
        // Remove smaller values at left.
        for que.Len() > 0 && arr[que.Back().(int)] >= arr[i] {
            que.RemoveBack()
        }
        que.Add(i)
        // window of size k
        if i >= (k - 1) && maxVal < arr[que.Front().(int)] {
            maxVal = arr[que.Front().(int)]
        }
        i += 1
    }
    fmt.Println("Max of min is : ", maxVal)
}

func main6() {
    arr := []int{11, 2, 75, 92, 59, 90, 55}
    k := 3
    maxOfMinSlidingWindows(arr, k)
}
/*
Max of min is :: 59
*/

func firstNegSlidingWindows(arr []int, k int) {
    size := len(arr)
    que := new(Queue)
    i := 0
    for i < size {
        // Remove out of range elements
        if que.Len() > 0 && que.Front().(int) <= i - k{
            que.Remove()
        }
        if arr[i] < 0{
            que.Add(i)
        }
        // window of size k
        if i >= (k - 1) {
            if que.Len() > 0{
                fmt.Print(arr[que.Front().(int)], " ")
            } else {
                fmt.Print("NAN ") 
            }

        }
        i += 1
    }        
}

func main7() {
    arr3 := []int{3, -2, -6, 10, -14, 50, 14, 21}
    firstNegSlidingWindows(arr3, 3)
}
/*
-2 -2 -6 -14 -14 NAN 
*/

func RottenFruitUtil(arr [][]int, maxCol int, maxRow int, currCol int, currRow int, traversed [][]int, day int) { // Range check
	if currCol < 0 || currCol >= maxCol || currRow < 0 || currRow >= maxRow {
		return
	}
	// Traversable and rot if not already rotten.
	if traversed[currCol][currRow] <= day || arr[currCol][currRow] == 0 {
		return
	}

    dir := [][]int { { -1, 0 }, { 1, 0 }, { 0, -1 }, { 0, 1 } };
	traversed[currCol][currRow] = day // Update rot time.
    for i := 0; i < 4; i++ {
        x := currCol + dir[i][0];
        y := currRow + dir[i][1];
	    RottenFruitUtil(arr, maxCol, maxRow, x, y, traversed, day+1)
    }
}

func RottenFruit(arr [][]int, maxCol int, maxRow int) int {
	traversed := make([][]int, maxRow)
	for i := range traversed {
		traversed[i] = make([]int, maxCol)
	}

	for i := 0; i < maxCol; i++ {
		for j := 0; j < maxRow; j++ {
			traversed[i][j] = 999999
		}
	}

	for i := 0; i < maxCol-1; i++ {
		for j := 0; j < maxRow-1; j++ {
			if arr[i][j] == 2 {
				RottenFruitUtil(arr, maxCol, maxRow, i, j, traversed, 0)
			}
		}
	}

	maxDay := 0
	for i := 0; i < maxCol-1; i++ {
		for j := 0; j < maxRow-1; j++ {
			if arr[i][j] == 1 {
				if traversed[i][j] == 999999 {
					return -1
				}
				if maxDay < traversed[i][j] {
					maxDay = traversed[i][j]
				}
			}
		}
	}
	return maxDay
}

type Fruit struct {
    x, y int ;
    day int ;
}

func  RottenFruit2(arr [][]int, maxCol int, maxRow int) int {
	traversed := make([][]bool, maxRow)
	for i := range traversed {
		traversed[i] = make([]bool, maxCol)
	}
    
    dir := [][]int { { -1, 0 }, { 1, 0 }, { 0, -1 }, { 0, 1 } };
    que := new(Queue)
    for i := 0; i < maxCol; i++ {
        for j := 0; j < maxRow; j++ {
            traversed[i][j] = false;
            if (arr[i][j] == 2) {
                que.Add(Fruit{x:i, y:j, day:0});
                traversed[i][j] = true;
            }
        }
    }
    max := 0
    var x, y, day int;
    var temp Fruit;
    for (!que.IsEmpty()) {
        temp = que.Remove().(Fruit);
        for i := 0; i < 4; i++ {
            x = temp.x + dir[i][0];
            y = temp.y + dir[i][1];
            day = temp.day + 1;
            if (x >= 0 && x < maxCol && y >= 0 && y < maxRow && arr[x][y] != 0 && traversed[x][y] == false) {
                que.Add(Fruit{x:x, y:y, day:day});
                if max < day {
                    max = day
                }
                traversed[x][y] = true;
            }
        }
    }
    for i := 0; i < maxCol; i++ {
        for j := 0; j < maxRow; j++ {
            if (arr[i][j] == 1 && traversed[i][j] == false) {
                return -1;
            }
        }
    }
    return max;
}

// Testing code.
func main21() {
	arr := [][]int{{1, 0, 1, 1, 0},{2, 1, 0, 1, 0},{0, 0, 0, 2, 1},{0, 2, 0, 0, 1},{1, 1, 0, 0, 1}}
	fmt.Println("RottenFruit :", RottenFruit(arr, 5, 5))
    fmt.Println("RottenFruit :", RottenFruit2(arr, 5, 5))
}
/*
RottenFruit : 3
*/


func StepsOfKnightUtil(size int, currCol int, currRow int, traversed [][]int, dist int) {
	// Range check
	if currCol < 0 || currCol >= size || currRow < 0 || currRow >= size {
		return
	}

	// Traverse if not already traversed.
	if traversed[currCol][currRow] <= dist {
		return
	}
	
    dir := [][]int {{-2,-1},{-2,1},{2,-1},{2,1},{-1,-2},{1,-2},{-1,2},{1,2 }};
	traversed[currCol][currRow] = dist

    for i := 0; i < 8; i++ {
        x := currCol + dir[i][0];
        y := currRow + dir[i][1];
		StepsOfKnightUtil(size, x, y, traversed, dist+1)
	}
}

func StepsOfKnight(size int, srcX int, srcY int, dstX int, dstY int) int {
	traversed := make([][]int, size)
	for i := range traversed {
		traversed[i] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			traversed[i][j] = 999999
		}
	}

	StepsOfKnightUtil(size, srcX-1, srcY-1, traversed, 0)
	retval := traversed[dstX-1][dstY-1]
	return retval
}

type Knight struct{
    x, y int;
    cost int;
}

func StepsOfKnight2( size int, srcX int, srcY int, dstX int, dstY int) int {
    traversed := make([][]int, size)
	for i := range traversed {
		traversed[i] = make([]int, size)
	}

    dir := [][]int {{-2,-1},{-2,1},{2,-1},{2,1},{-1,-2},{1,-2},{-1,2},{1,2 }};
    que := new(Queue)

    for i := 0; i < size; i++ {
        for j := 0; j < size; j++ {
            traversed[i][j] = 999999;
        }
    }
    que.Add(Knight{x:srcX - 1,y: srcY - 1,cost: 0});
    traversed[srcX - 1][srcY - 1] = 0;

    var x, y, cost int
    var temp Knight
    for (!que.IsEmpty()) {
        temp = que.Remove().(Knight);
        for i := 0; i < 8; i++ {
            x = temp.x + dir[i][0];
            y = temp.y + dir[i][1];
            cost = temp.cost + 1;
            if (x>=0 && x < size && y>=0 && y<size && traversed[x][y] > cost) {
                que.Add(Knight{x:x, y:y, cost:cost});
                traversed[x][y] = cost;
            }
        }
    }
    return traversed[dstX - 1][dstY - 1];
}

// Testing code.
func main22() {
	fmt.Println("StepsOfKnight :", StepsOfKnight(20, 10, 10, 20, 20))
    fmt.Println("StepsOfKnight :", StepsOfKnight2(20, 10, 10, 20, 20))
}
/*
StepsOfKnight : 8
*/

func DistNearestFillUtil(arr [][]int, maxCol int, maxRow int, currCol int, currRow int, traversed [][]int, dist int) { // Range check
	if currCol < 0 || currCol >= maxCol || currRow < 0 || currRow >= maxRow {
		return
	}
	// Traversable if their is a better distance.
	if traversed[currCol][currRow] <= dist {
		return
	}
	// Update distance.
	traversed[currCol][currRow] = dist
	dir := [][]int { { -1, 0 }, { 1, 0 }, { 0, -1 }, { 0, 1 } };
	
	// each line corresponding to 4 direction.
    for i := 0; i < 4; i++ {
        x := currCol + dir[i][0];
        y := currRow + dir[i][1];
	    DistNearestFillUtil(arr, maxCol, maxRow, x, y, traversed, dist+1)
	}
}

func DistNearestFill(arr [][]int, maxCol int, maxRow int) {
	traversed := make([][]int, maxRow)
	for i := range traversed {
		traversed[i] = make([]int, maxCol)
	}
	for i := 0; i < maxCol; i++ {
		for j := 0; j < maxRow; j++ {
			traversed[i][j] = 999999
		}
	}
	for i := 0; i < maxCol; i++ {
		for j := 0; j < maxRow; j++ {
			if arr[i][j] == 1 {
				DistNearestFillUtil(arr, maxCol, maxRow, i, j, traversed, 0)
			}
		}
	}

	for i := 0; i < maxCol; i++ {
		for j := 0; j < maxRow; j++ {
			fmt.Print(traversed[i][j], " ")
		}
		fmt.Println()
	}
}

type Node2 struct{
    x, y, dist int;
}

func DistNearestFill2( arr [][]int,  maxCol int,  maxRow int) {
    traversed := make([][]int, maxRow)
	for i := range traversed {
		traversed[i] = make([]int, maxCol)
	}
    dir := [][]int { { -1, 0 }, { 1, 0 }, { 0, -1 }, { 0, 1 } };
    que := new(Queue)
    for i := 0; i < maxCol; i++ {
        for j := 0; j < maxRow; j++ {
            traversed[i][j] = 999999;
            if (arr[i][j] == 1) {
                que.Add(Node2{x:i, y:j, dist:0});
                traversed[i][j] = 0;
            }
        }
    }
    var x, y, dist int;
    var temp Node2;
    for (!que.IsEmpty()) {
        temp = que.Remove().(Node2)
        for i := 0; i < 4; i++ {
            x = temp.x + dir[i][0];
            y = temp.y + dir[i][1];
            dist = temp.dist + 1;
            if (x>=0 && x<maxCol && y>=0 && y<maxRow && traversed[x][y] > dist) {
                que.Add(Node2{x:x, y:y, dist:dist});
                traversed[x][y] = dist;
            }
        }
    }
    for i := 0; i < maxCol; i++ {
        for j := 0; j < maxRow; j++ {
            fmt.Print(traversed[i][j], " ");
        }
        fmt.Println();
    }
}
// Testing code.
func main23() {
	arr := [][]int{{1, 0, 1, 1, 0},{1, 1, 0, 1, 0},{0, 0, 0, 0, 1},{0, 0, 0, 0, 1},{0, 0, 0, 0, 1}}

	DistNearestFill(arr, 5, 5)
    DistNearestFill2(arr, 5, 5)
}
/*
0 1 0 0 1
0 0 1 0 1
1 1 2 1 0
2 2 2 1 0
3 3 2 1 0
 */

func findLargestIslandUtil(arr [][]int, maxCol int, maxRow int, currCol int,
	currRow int, traversed [][]bool) int {
    dir := [][]int {{ -1, -1 }, { -1, 0 }, { -1, 1 }, { 0, -1 }, { 0, 1 }, { 1, -1 }, { 1, 0 }, { 1, 1 } }
    var x, y int
    var sum = 1
    for i := 0; i < 8; i++ {
        x = currCol + dir[i][0]
        y = currRow + dir[i][1]
        if (x >= 0 && x < maxCol && y >= 0 && y < maxRow && traversed[x][y] == false && arr[x][y] == 1) {
			traversed[x][y] = true
			sum += findLargestIslandUtil(arr, maxCol, maxRow, x, y, traversed)
		}
    }
	return sum

}

func findLargestIsland(arr [][]int, maxCol int, maxRow int) int {
	maxVal := 0
	currVal := 0

	traversed := make([][]bool, maxRow)
	for i := range traversed {
		traversed[i] = make([]bool, maxCol)
	}
	for i := 0; i < maxCol; i++ {
		for j := 0; j < maxRow; j++ {
			traversed[i][j] = false
		}
	}

	for i := 0; i < maxCol; i++ {
		for j := 0; j < maxRow; j++ {		
            if (arr[i][j] == 1) {
				traversed[i][j] = true
				currVal = findLargestIslandUtil(arr, maxCol, maxRow, i, j, traversed)
                if currVal > maxVal {
                    maxVal = currVal
                }
            }
		}
	}
	return maxVal
}

// Testing code.
func main24() {
	arr := [][]int{{1, 0, 1, 1, 0}, 
    {1, 0, 0, 1, 0},
    {0, 1, 1, 1, 1},
    {0, 1, 0, 0, 0},
    {1, 1, 0, 0, 1}}

	fmt.Println("Largest Island :", findLargestIsland(arr, 5, 5))
}
/*
Largest Island :  12
*/

func main(){
	main1()
	main2()
	main3()
	main4()
	main5()
	main6()
	main7()
    main21()
    main22()
    main23()
    main24()
}

type Queue struct {
	que []interface{}
}

func (q *Queue) Add(value interface{}) {
	q.que = append(q.que, value)
}

func (q *Queue) Remove() interface{} {
	n := len(q.que)
	value := q.que[0]
	q.que = q.que[1:n]
	return value
}

func (q *Queue) RemoveBack() interface{} {
	n := len(q.que)
	value := q.que[n-1]
	q.que = q.que[:n-1]
	return value
}

func (q *Queue) Front() interface{} {
	return q.que[0]
}

func (q *Queue) Back() interface{} {
	n := len(q.que)
	return q.que[n-1]
}

func (q *Queue) IsEmpty() bool {
	return len(q.que) == 0
}

func (q *Queue) Len() int {
	return len(q.que)
}

func (q Queue) Print() {
	fmt.Println(q.que)
}
