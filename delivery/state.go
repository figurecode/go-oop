package delivery

// State — структура статуса доставки и обработки сообщения.
type State struct {
	V   string
	Log func(mes string)
}

// Возможные значения перечисления DeliveryState.
const (
	StatePending   string = "pending"      // сообщение отправлено
	StateAck       string = "acknowledged" // сообщение получено
	StateProcessed string = "processed"    // сообщение обработано успешно
	StateCanceled  string = "canceled"     // обработка сообщения прервана
)

// IsValid проверяет валидность текущего значения типа State
func (s State) IsValid() bool {
	switch s.V {
	case StatePending, StateAck, StateProcessed, StateCanceled:
		return true
	default:
		return false
	}
}

// String возвращает строковое представление типа State
func (s State) String() string {
	return string(s.V)
}
