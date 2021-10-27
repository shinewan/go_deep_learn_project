package queue

/**查看接口变量
	interface{} 表示任何类型
	Type Assertion
	Type Switch
 */

type Queue []interface{}
//type Queue []int  这就是控制类型为：int

func (q *Queue) Push(v interface{}){
	*q = append(*q,v)
}
func (q *Queue) Pop() interface{}{
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool{
	return len(*q) ==0
}
