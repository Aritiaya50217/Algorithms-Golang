package main

import "fmt"

// make set initalize the set
func makeSet() *customSet {
	return &customSet{
		container: make(map[string]struct{}),
	}
}

type customSet struct {
	container map[string]struct{}
}

func (c *customSet) Exists(key string) bool {
	// check key
	_, exists := c.container[key]
	return exists
}

func (c *customSet) Add(key string) {
	// add key
	c.container[key] = struct{}{}
}

func (c *customSet) Remove(key string) error {
	_, exists := c.container[key]
	if !exists {
		return fmt.Errorf("Remove Error: Item doesn't exist in set")
	}
	delete(c.container, key)  // delete key in struct
	return nil
}

func (c *customSet) Size() int {
	return len(c.container)
}

func main() {
	customSet := makeSet()
	fmt.Printf("Add: A\n")

	fmt.Printf("%v\n", customSet)

	customSet.Add("A")
	fmt.Printf("%v\n", customSet)
	fmt.Printf("Add: B\n")

	fmt.Printf("%v\n", customSet)

	customSet.Add("B")

	fmt.Printf("%v\n", customSet)

	fmt.Printf("Size: %d\n", customSet.Size())
	fmt.Printf("A Exists?: %t\n", customSet.Exists("A"))  // check
	fmt.Printf("B Exists?: %t\n", customSet.Exists("B"))
	fmt.Printf("C Exists?: %t\n", customSet.Exists("C"))
	fmt.Printf("Remove: B\n")
	customSet.Remove("B")
	fmt.Printf("B Exists?: %t\n", customSet.Exists("B"))

	fmt.Printf("%v", customSet)

	// &{map[]} => app A => &{map[A:{}]} => add B => &{map[A:{} B:{}]}
	// &{map[A:{} B:{}]} => remove B => &{map[A:{}]}

}
