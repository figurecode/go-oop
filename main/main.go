package main

import "fmt"
import "delivery"

//Метод представляет собой функцию, привязанную к конкретному типу.
//Методы позволяют связывать поведение и данные типа в самом типе,
//обеспечивая инкапсуляцию.

// MyType Объявление типа
type MyType int

// Объявление метода типа MyType
func (m MyType) String() string {
	return fmt.Sprintf("MyType: %d", m)
}

func HandleMsgDeliveryState(status delivery.State) error {
	if !status.IsValid() {
		return fmt.Errorf("status: invalid")
	}

	// код обработки статуса

	return nil
}

func main() {
	var m MyType = 5

	s := m.String()

	fmt.Println(s)

	var state delivery.State = "new"

	err := HandleMsgDeliveryState(state)
	if err != nil {
		fmt.Println(err)
	}
}
