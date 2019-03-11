package main

type Counter map[interface{}]int

func (s *Counter) Add(key interface{}) {
	(*s)[key]++
}
func (s *Counter) Find(key interface{}) bool {
	_, ok := (*s)[key]
	return ok
}

func (s *Counter) Remove(key interface{}) {
	val, ok := (*s)[key]
	if ok == false {
		return
	} else if val == 1 {
		delete((*s), key)
		return
	}
	(*s)[key]--
}

func (s *Counter) Get(key interface{}) (int, bool) {
	val, ok := (*s)[key]
	return val, ok
}

/*
func main() {
	mp := make(Counter)
	mp.Add("a")
	mp.Add("b")
	mp.Add("a")
	mp.Add(1)
	mp.Add(1)
	mp.Add(1)
	mp.Add(1)
	fmt.Println(mp.Find("a"))
	fmt.Println(mp.Find("b"))
	fmt.Println(mp.Find("c"))
	fmt.Println(mp.Find(1))
	fmt.Println(mp.Get(1))
	fmt.Println(mp.Get("a"))
	fmt.Println(mp.Get("b"))
	fmt.Println(mp.Get("c"))
}
*/
