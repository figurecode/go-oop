package embedded

import "fmt"

type Student struct {
	Person
	Group string
}

func NewStudent(name string, year int, group string) Student {
	return Student{
		Person: NewPerson(name, year),
		Group:  group,
	}
}

func (s Student) String() string {
	return fmt.Sprintf("%s, Группа: %s", s.Person, s.Group)
}

func (s *Student) Debug() {
	// Если несколько вложенных типов имеют одинаковые методы, происходит конфликт.
	// Обращаться к полям (или методам) вложенной структуры можно как с указанием
	// типа объекта (s.Person.Print()), так и без (s.Print()).
	// Для String-метода объекта Person требуется явно указывать тип,
	// так как метод переопределён типом Student.
	// Есть два правила разрешения конфликтов имён полей и методов:
	//    * Именованное поле (метод) структуры скрывает поле (метод) с тем же
	//      именем для вложенных структур. Имя верхнего уровня доминирует над именами
	//      более низких уровней. Например, вызов метода student.String() вызывает метод
	//      структуры Student, а не структуры Person.
	//    * Если имя поля (метода) встречается на том же уровне вложенности
	//      (дупликация имён), и оно использовано в коде, это ошибка.
	//      Если дупликация имён существует, но это имя не используется в коде,
	//      компилятор не выдаст ошибку. Например, если тип Student имел бы ещё одну
	//      вложенную структуру Faculty с методом Print(), то тогда при вызове метода
	//      необходимо указывать имя вложенного типа s.Person.Print() или s.Faculty.Print().

	// доступ к методам объекта Person
	s.Print()
	// или
	s.Person.Print()

	// доступ к полю 'Name' объекта Person
	s.Name = "Mark Smith"
	// или
	s.Person.Name = "Mark Smith"

	// вызовется метод String объекта Student
	fmt.Println(s)
	// вызовется метод String объекта Person
	fmt.Println(s.Person)
}
