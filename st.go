package tools
type DBkeys[T int|float32|float64|string|bool] struct{
	Key []byte
	Value T
	Langth int
}
type DBarray[T int|float32|float64|string|bool|byte] struct{
	Key []byte
	Value []T
	Langth int
}
func (s *DBkeys[T]) setvalue(value T) bool{
	s.Value=value
	return true
}
func (s *DBarray[T]) setvalue(value []T) bool{
	s.Value=value
	s.Langth=len(value)
	return true
}
func NewKey[T int|float32|float64|string|bool](key string,value T) (*DBkeys[T],bool){
	var res DBkeys[T]
	res.Key=[]byte(key)
	res.setvalue(value)
	return &res,true
}
func Newarray[T int|float32|float64|string|bool|byte](key string,value []T) (*DBarray[T],bool){
	var res DBarray[T]
	res.Key=[]byte(key)
	res.setvalue(value)
	return &res,true
}