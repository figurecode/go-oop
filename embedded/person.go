package embedded

import "fmt"

type Person struct {
	Name string
	Year int
}

func NewPerson(name string, year int) Person {
	return Person{
		Name: name,
		Year: year,
	}
}

func (p Person) String() string {
	return fmt.Sprintf("Имя: %s, Год рождения: %d", p.Name, p.Year)
}

func (p Person) Print() {
	fmt.Println(p)
}
