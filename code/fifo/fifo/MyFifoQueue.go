/*
@Time :  2020-07-30 15:45
@Author : Caden
@File : MyFifoQueue
@Software: GoLand
*/
package fifo

type MyFifoQueue struct {
	head int
	tail int
	data []interface{}
}

func NewMyFifoQueue() (queue *MyFifoQueue) {
	return &MyFifoQueue{
		data: make([]interface{}, 0),
		head: 0,
		tail: 0,
	}
}
func NewMyFifoQueueForNum(num int) (queue *MyFifoQueue) {
	return &MyFifoQueue{
		data: make([]interface{}, num),
		tail: 0,
		head: 0,
	}
}

func (self *MyFifoQueue) Enqueue(item interface{}) {
	i := append(self.data, item)
	self.data = i
}
func (self *MyFifoQueue) Dequeue() interface{} {
	i := self.data[0]
	self.data = self.data[1:len(self.data)]
	return i
}
