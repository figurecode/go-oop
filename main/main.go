package main

import (
	"circular"
	"delivery"
	"embedded"
	"fmt"
	"loggerLevel"
	"time"
	"timer"
)

//Метод представляет собой функцию, привязанную к конкретному типу.
//Методы позволяют связывать поведение и данные типа в самом типе,
//обеспечивая инкапсуляцию.

// MyType Объявление типа
type MyType int

// Объявление метода типа MyType
func (m MyType) String() string {
	return fmt.Sprintf("MyType: %d", m)
}

func HandleMsgDeliveryState(status *delivery.State) {
	if !status.IsValid() {
		status.Log("status: [" + status.V + "] invalid")

		return
	}

	status.Log("status: [" + status.V + "]  OK")

	// код обработки статуса
}

// HandleBuffer передача метода как аргумента функции
func HandleBuffer(num float64, add func(float64)) {
	add(num)
}

func main() {
	var m MyType = 5

	s := m.String()

	fmt.Println(s)

	// ***************************
	var state = delivery.State{
		V:   "new",
		Log: func(m string) { fmt.Println("LOG: ", m) },
	}

	HandleMsgDeliveryState(&state)

	state.V = "pending"

	HandleMsgDeliveryState(&state)

	// ***************************
	buf := circular.NewCircularBuffer(4)
	for i := 0; i < 6; i++ {
		if i > 0 {
			fmt.Println("Добавляем в буфер: ", i)

			buf.AddValue(float64(i))
		}

		fmt.Printf("[%d]: %v\n", buf.GetCurrentSize(), buf.GetValues())
	}

	fmt.Println("Буфер: ", buf)

	for i := 4; i < 8; i++ {
		HandleBuffer(float64(i), buf.AddValue)
	}

	fmt.Println("Буфер: ", buf)

	// ***************************

	sw := timer.Stopwatch{}
	sw.Start()

	time.Sleep(1 * time.Second)
	sw.SaveSplit()

	time.Sleep(500 * time.Millisecond)
	sw.SaveSplit()

	time.Sleep(300 * time.Millisecond)
	sw.SaveSplit()

	fmt.Println(sw.GetResults())

	// ***************************

	student := embedded.NewStudent(&embedded.Person{Name: "John Doe", Year: 1980}, "701")
	student.Print()

	fmt.Println(student)
	fmt.Println(student.Name, student.Year, student.Group)

	embedded.ChangeName(student, "Teodor")

	fmt.Println(student)

	// ***************************

	extBuffer := circular.NewExtCircularBuffer(5)
	extBuffer.AddValues(1, 2, 3, 4, 5)
	fmt.Printf("[%d]: %v\n", extBuffer.GetCurrentSize(), extBuffer.GetValues())

	// ***************************

	logger := loggerLevel.NewLogExtended()
	logger.SetLogLevel(loggerLevel.LogLevelWarning)
	logger.Infoln("Не должно напечататься")
	logger.Warningln("Hello")
	logger.Errorln("World")

	// ***************************

}
