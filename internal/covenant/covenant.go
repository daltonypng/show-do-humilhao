package covenant

type Database interface {
	Create(destiny interface{}) error
	Read(destiny interface{}, conditionals ...interface{}) error
	Update(destiny interface{}) error
	Delete(destiny interface{}, conditionals ...interface{}) error
}
