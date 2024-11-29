package chat



type ValidationError struct {
	String string
}


func (v ValidationError) Error() string {
	return v.String
}