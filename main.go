package tools
// simple to reverse string to string Reverse('hello')->olleh
func Reverse(value string) string{
	news:=""
	for i := 1; i <= len(value); i++ {
		news+=value[len(value)-i:len(value)-i+1]
	}
	return news
}
// sort array example:Sort(yourarray)->(sort from high to low)
func Sort[T int | float32 | float64] (value []T) []T{
	max:=value[0]
	var newarr []T
	for _,v:=range value{
		if v>max{
			max=v
		}
	}
	newarr=append(newarr,max)
	for len(newarr)<len(value){
		max=0
		for _,v:=range value{
			if v>max && v<newarr[len(newarr)-1]{
				max=v
			}
		}
		newarr = append(newarr, max)
	}
	return newarr
	
}
// sort Two-dimensional array
func Sortplus[T int | float32 | float64] (value [][]T,post int) [][]T{
	max:=value[0]
	var newarr [][]T
	
	for _,v:=range value{
		if v[post]>max[post]{
			max=v
		}
	}
	newarr = append(newarr, max)
	for len(newarr)<len(value){
		for _,v:=range value{
			if v[post]>max[post] && v[post]<newarr[len(newarr)-1][post]{
				max=v
			}
		}
		newarr = append(newarr, max)
	}
	return newarr
}
// find max and min from your array.example:max,min:=MaxandMin(your arrary)
func MaxandMin[T int | float32 | float64](value []T) (T,T){
	max:=value[0]
	min:=value[0]
	for _,v:=range value{
		if v>max{
			max=v
		}else if v<min{
			min=v
		}
	}
	return max,min
}

