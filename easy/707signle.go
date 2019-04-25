package easy

import "fmt"

/*
设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next 是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。

在链表类中实现这些功能：

get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。


示例：

MyLinkedList linkedList = new MyLinkedList();
linkedList.addAtHead(1);
linkedList.addAtTail(3);
linkedList.addAtIndex(1,2);   //链表变为1-> 2-> 3
linkedList.get(1);            //返回2
linkedList.deleteAtIndex(1);  //现在链表是1-> 3
linkedList.get(1);            //返回3
*/
type MyLinkedList struct {
	val  int
	next *MyLinkedList
	use  bool
}

/** Initialize your data structure here. */
func Constructor707S() MyLinkedList {

	return MyLinkedList{
		val:  -1,
		next: nil,
		use:  false,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {

	node := this
	i := 0
	for node != nil {
		if i == index {
			return node.val
		}
		i++
		node = node.next
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {

	if !this.use {
		this.val = val
		this.next = nil
		this.use = true
		return
	}
	node := &MyLinkedList{
		val:  this.val,
		next: this.next,
	}
	this.val = val
	this.next = node

}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {

	node := this

	for node.next != nil {
		node = node.next
	}

	node.next = &MyLinkedList{
		val:  val,
		next: nil,
	}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {

	if index == 0 {
		this.AddAtHead(val)
	}
	node := this
	i := 1
	for node != nil {
		if i == index {
			tmpV := &MyLinkedList{
				val:  val,
				next: node.next,
			}
			node.next = tmpV
		}
		i++
		node = node.next
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {

	if index == 0 {
		if this.next != nil {
			this.val = this.next.val
			this.next = this.next.next
		} else {
			this.val = -1
			this.next = nil
			this.use = false
		}
		return
	}
	i := 1
	node := this
	for node.next != nil {
		if i == index {
			node.next = node.next.next
			break
		}
		i++
		node = node.next
	}
}

func (this *MyLinkedList) P() {

	node := this
	for node != nil {
		fmt.Printf("%d,", node.val)
		node = node.next
	}
	fmt.Println("")
}
