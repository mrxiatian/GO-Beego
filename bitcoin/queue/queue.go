package queue

import (
	"errors"
	"fmt"
)

//实现链式(单链表)存储结构的环形队列,极其相关操作
type CircleQueue struct {
	data string					//入队数据
	next *CircleQueue	//下一个元素
	tail *CircleQueue	//队尾元素
	//用头节点端作为队头
}

//初始化一个链式结构的环形队列
func InitCircleQueue() *CircleQueue {
	CircleQueue := &CircleQueue{
		data: "",
		next: nil,
		tail: nil,
	}
	CircleQueue.next = CircleQueue
	CircleQueue.tail = CircleQueue

	return CircleQueue
}

//新建一个单链表节点
func (Q *CircleQueue) New(data string) *CircleQueue {
	return &CircleQueue{
		data: data,
		next: nil,
		tail: nil,
	}
}

//入队操作
func (Q *CircleQueue) Push(data string) error {
	//判断队列是否已满
	if Q.IsFull() {
		return errors.New("队列已满,请Pull后重试!")
	}

	//将数据入队
	newNode := Q.New(data)

	Q.tail.next = newNode
	newNode.next = Q
	Q.tail = Q.tail.next
	return nil
}

//获得队尾元素数据
func (Q *CircleQueue) GetTail() (string,error) {
	//判断队列是否为空
	if Q.IsEmpty() {
		return "", errors.New("当前队列为空,请Push后重试!")
	}

	//返回数据
	return Q.tail.data, nil
}

//获得队首元素数据
func (Q *CircleQueue) GetFirst() (string,error) {
	//判断队列是否为空
	if Q.next == Q {
		return "", errors.New("当前队列为空,请Push后重试!")
	}

	//返回数据
	return Q.next.data, nil
}

//出队列操作
func (Q *CircleQueue) Pull() (string,error) {
	//判断队列是否为空
	if Q.IsEmpty() {
		return "", errors.New("当前顺序结构环形队列为空,请Push后重试!")
	}
	//数据出队列
	data := Q.next.data
	Q.next = Q.next.next
	if Q.next == Q {
		Q.tail = Q
	}

	return data, nil
}

//判断队列是否已满
func (Q *CircleQueue) IsFull() bool {
	//初始化后没有添加元素进队列中的时候Q.next为空
	if Q.next == nil {
		return false
	}

	//计数器
	count := 0

	//定义一个中间变量,辅助计数
	temp := Q

	//添加数据后最后一个队列节点的next会指向头节点
	for temp.next != Q {
		count++
		temp = temp.next
	}

	//其实链式结构不用定义最大存储容量,这里为了测试定义为20
	return count == 200
}

//判断队列是否为空
func (Q *CircleQueue) IsEmpty() bool {
	//有可能没有进行添加数据也判断是否为空
	return Q.next == Q || Q.tail == Q
}

//获取队列中当前的长度
func (Q *CircleQueue) Len() int {
	//初始化后没有添加元素进队列中的时候Q.next为空
	if Q.next == nil {
		return 0
	}
	//计数器
	count := 0

	//定义一个中间变量,辅助计数
	temp := Q

	//添加数据后最后一个队列节点的next会指向头节点
	for temp.next != Q {
		count++
		temp = temp.next
	}
	return count
}

//显示队列
func (Q *CircleQueue) Show() {
	//长度为零说明队列中没有元素
	if Q.Len() == 0 {
		return
	}

	//定义一个中间变量,帮忙遍历队列
	temp := Q

	for temp.next != Q {
		if temp.next.next != Q {
			fmt.Printf("%d-->",temp.next.data)
		}else {
			fmt.Printf("%d\n",temp.next.data)
		}
		temp = temp.next
	}
}
