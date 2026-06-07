// Embedding lets one struct "inherit" fields and methods from another without true inheritance
package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Speak() string {
	return a.Name + " makes a sound"
}

// Dog embeds Animal — it gets Name and Speak() for free
type Dog struct {
	Animal      // no field name — this is embedding
	Breed string
}

// Dog overrides Speak to give its own version
func (d Dog) Speak() string {
	return d.Name + " barks!"
}

type Cat struct {
	Animal
}

// Cat doesn't override Speak — it uses Animal's version

type Worker struct {
	Name string
}

func (w Worker) Work() string {
	return w.Name + " is working"
}

// Manager embeds Worker and adds its own field
type Manager struct {
	Worker
	Department string
}

func main() {
	d := Dog{Animal: Animal{Name: "Rex"}, Breed: "Husky"}
	fmt.Println(d.Speak())           // Dog's Speak
	fmt.Println(d.Animal.Speak())    // Animal's Speak, accessed explicitly
	fmt.Println("Breed:", d.Breed)
	fmt.Println("Name (promoted):", d.Name) // Name promoted from Animal

	c := Cat{Animal: Animal{Name: "Luna"}}
	fmt.Println(c.Speak()) // Uses Animal's Speak since Cat doesn't override

	m := Manager{Worker: Worker{Name: "Alice"}, Department: "Engineering"}
	fmt.Println(m.Work()) // promoted from Worker
	fmt.Println("Dept:", m.Department)
}
