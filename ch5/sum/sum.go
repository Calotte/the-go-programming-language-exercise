package sum

func Sum(vals ...int)int{
	total:=0
	for _,val:=range vals{
		total+=val
	}
	return total
}
