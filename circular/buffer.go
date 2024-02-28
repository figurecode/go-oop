package circular

// Обратите внимание, у типа CircularBuffer есть методы с получателями и по
// значению, и по указателю. Такое смешивание разных типов методов допускается
// стандартом языка, но не принято в Go-сообществе. Придерживайтесь соглашений,
// которым следует ваша команда разработчиков. Если у объекта все методы только с
// получателем по значению и все поля неэкспортируемые, можно сказать, что этот
// объект неизменяем (immutable). И наоборот, объект изменяем (mutable), если все
// методы с получателем по указателю или одно из полей экспортируемое. Это сильно
// облегчает жизнь пользователям вашего типа, если вы разрабатываете библиотеку с
// десятками тысяч ⭐️ на GitHub.

// CircularBuffer реализует структуру данных [кольцевой буфер](https://ru.wikipedia.org/wiki/Кольцевой_буфер) для значений float64.
type CircularBuffer struct {
	values  []float64 // текущее значение буфера
	headIdx int       // индекс головы (первый непустой элемент)
	tailIdx int       // индекс хвоста (первый пустой элемент)
}

// GetCurrentSize возвращает текущую длину буфера
func (b CircularBuffer) GetCurrentSize() int {
	if b.tailIdx < b.headIdx {
		return b.tailIdx + cap(b.values) - b.headIdx
	}

	return b.tailIdx - b.headIdx
}

// GetValues возвращает слайс текуших значений буфера, сохраняя порядок записи
func (b CircularBuffer) GetValues() (retValues []float64) {
	for i := b.headIdx; i != b.tailIdx; i = (i + 1) % cap(b.values) {
		retValues = append(retValues, b.values[i])
	}

	return
}

// AddValue добавляет новое значение в буфер
func (b *CircularBuffer) AddValue(v float64) {
	b.values[b.tailIdx] = v

	b.tailIdx = (b.tailIdx + 1) % cap(b.values)
	if b.tailIdx == b.headIdx {
		b.headIdx = (b.headIdx + 1) % cap(b.values)
	}
}

// NewCircularBuffer - конструктор типа CircularBuffer
func NewCircularBuffer(size int) CircularBuffer {
	return CircularBuffer{values: make([]float64, size+1)}
}
