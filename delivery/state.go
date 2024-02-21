package delivery

// State — статус доставки и обработки сообщения.
type State string

// Возможные значения перечисления DeliveryState.
const (
	StatePending   State = "pending"      // сообщение отправлено
	StateAck       State = "acknowledged" // сообщение получено
	StateProcessed State = "processed"    // сообщение обработано успешно
	StateCanceled  State = "canceled"     // обработка сообщения прервана
)

// IsValid проверяет валидность текущего значения типа State
func (s State) IsValid() bool {
	switch s {
	case StatePending, StateAck, StateProcessed, StateCanceled:
		return true
	default:
		return false
	}
}

// String возвращает строковое представление типа State
func (s State) String() string {
	return string(s)
}
