package easy

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
type DoubleMyLinkedList struct {
	val  int
	prev *DoubleMyLinkedList
	next *DoubleMyLinkedList
	head *DoubleMyLinkedList
	tail *DoubleMyLinkedList
}

/** Initialize your data structure here. */
func Constructor707() DoubleMyLinkedList {

	return DoubleMyLinkedList{
		val:  -1,
		prev: nil,
		next: nil,
		head: nil,
		tail: nil,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *DoubleMyLinkedList) Get707(index int) int {
	list := this.head
	i := 0
	for list != nil {
		if i == index {
			return list.val
		}
		list = list.next
		i++
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *DoubleMyLinkedList) AddAtHead707(val int) {
	if this.head == nil {
		this.val = val
		this.head = this
		this.tail = this
	} else {
		node := &DoubleMyLinkedList{
			val:  val,
			prev: nil,
			next: nil,
			head: nil,
			tail: nil,
		}
		node.next = this
		this.prev = node
		this.head = node
	}

}

/** Append a node of value val to the last element of the linked list. */
func (this *DoubleMyLinkedList) AddAtTail707(val int) {

	if this.tail == nil {
		this.val = val
		this.tail = this
		this.head = this
	} else {
		node := &DoubleMyLinkedList{
			val:  val,
			prev: nil,
			next: nil,
			head: nil,
			tail: nil,
		}

		node.prev = this.tail
		this.tail.next = node
		this.tail = node
	}

}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *DoubleMyLinkedList) AddAtIndex707(index int, val int) {

	if this.head == nil || index == 0 {
		this.AddAtHead707(val)
	}

	node := &DoubleMyLinkedList{
		val:  val,
		prev: nil,
		next: nil,
		head: nil,
		tail: nil,
	}
	list := this
	i := 1
	for list != nil {
		if i == index {

			if list.next == nil {
				this.tail = node
			}
			node.next = list.next
			list.next = node
			node.prev = list
			break

		}
		i++
		list = list.next
	}

}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *DoubleMyLinkedList) DeleteAtIndex707(index int) {

	if this.head == nil {
		return
	}
	if index == 0 {
		if this.head.next != nil {
			this.head.next.prev = nil
			this.head = this.head.next

		} else {
			this.head, this.next, this.prev, this.next = nil, nil, nil, nil
		}
		return
	}

	// 感觉思路有点问题，需要再想想

}
