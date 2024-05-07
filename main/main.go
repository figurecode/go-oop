package main

import (
	"bufio"
	"circular"
	"delivery"
	"embedded"
	"errors"
	"fmt"
	"loggerLevel"
	"os"
	"strings"
	"time"
	"timeError"
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

func HandleMsgDeliveryState(status *delivery.State) error {
	if !status.IsValid() {
		status.Log("status: [" + status.V + "] invalid")

		return timeError.NewTimeError(fmt.Sprintf("status \"%v\" invalid", status.V))
	}

	status.Log("status: [" + status.V + "]  OK")

	// код обработки статуса

	return nil
}

// HandleBuffer передача метода как аргумента функции
func HandleBuffer(num float64, add func(float64)) {
	add(num)
}

// Функция должна находить все возможные ошибки за один вызов.
// Результат должен содержать слайс ошибок, по которым строка не прошла проверку, или быть nil.

type SliceError []error

func (errs SliceError) Error() string {
	var out string
	for _, err := range errs {
		out += err.Error() + ";"
	}

	return strings.TrimRight(out, ";")
}

func MyCheck(input string) error {
	var (
		err      SliceError
		spaces   int
		hasDigit bool
	)
	// Длина должна быть меньше 20 символов
	if len([]rune(input)) >= 20 {
		err = SliceError{errors.New("line is to long")}
	}

	for _, ch := range input {
		if ch == ' ' {
			spaces++
		} else if ch >= '0' && ch <= '9' {
			hasDigit = true
		}
	}

	// Строка не должна содержать цифр
	if hasDigit {
		err = append(err, errors.New("found numbers"))
	}

	// Строка должна иметь два пробела
	if spaces != 2 {
		err = append(err, errors.New("no two spaces"))
	}

	if len(err) == 0 {
		return nil
	}

	return err
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

	if err := HandleMsgDeliveryState(&state); err != nil {
		fmt.Println(err)
	}

	state.V = "pending"

	if err := HandleMsgDeliveryState(&state); err != nil {
		fmt.Println(err)
	}

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

	// Пример обработки ошибок. Ошибки складываются в слайс.
	for {
		fmt.Printf("Укажите строку (q для выхода): ")
		reader := bufio.NewReader(os.Stdin)
		ret, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		ret = strings.TrimRight(ret, "\n")
		if ret == `q` {
			break
		}
		if err = MyCheck(ret); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(`Строка прошла проверку`)
		}
	}
}
