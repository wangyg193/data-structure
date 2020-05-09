package datastruct

type Queue struct {
	*SingleList
	Length int // 队列的长度
}

// 创建队列结构实例
func NewQueue() *Queue {
	return &Queue{NewSingleList(), 0}
}

// Push入队列
func (this *Queue) Push(data ElementObject) {
	this.InsertData(1, data)
	this.Length += 1
}

// Pop出队列
func (this *Queue) Pop() (ElementObject, error) {
	data, err := this.GetLastSingleList().DataValue()
	if err != nil {
		return nil, err
	}
	if errDel := this.DeleteLast(); errDel != nil {
		return nil, errDel
	}
	this.Length -= 1
	return data, err
}

// 获取队列长度
func (this *Queue) GetLen() int {
	return this.Length
}

// 遍历队列, 打印数据
func (this *Queue) Traverse() {
	this.SingleList.Traverse()
}