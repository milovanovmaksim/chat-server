package chat

// ValidationError представляет ошибку валидации входных данных.
// Удовлетворяет интерфейсу errors.Error.
type ValidationError struct {
	String string
}

func (v ValidationError) Error() string {
	return v.String
}
