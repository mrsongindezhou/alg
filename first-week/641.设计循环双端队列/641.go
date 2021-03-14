package main

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/14 11:19 上午
 * @description：
 * @modified By：
 * @version    ：$
 */
type MyCircularDeque struct {
	Head *LinkedNode
	Tail *LinkedNode
	Size int // 队列大小
	Length int // 当前队列长度
}

type LinkedNode struct {
	Val  int
	Pre  *LinkedNode
	Next *LinkedNode
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	head := &LinkedNode{}
	tail := &LinkedNode{}
	head.Next = tail
	tail.Pre = head
	mcd := MyCircularDeque{
		Head: head,
		Tail: tail,
		Size: k,
		Length: 0,
	}
	return mcd
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	// 判断队列是否已满
	if this.Length >= this.Size {
		return false
	}
	// 生成新节点
	newNode := &LinkedNode{
		Val: value,
		Next:this.Head.Next,
		Pre: this.Head,
	}
	// 插入元素
	this.Head.Next.Pre = newNode
	this.Head.Next = newNode
	this.Length++
	return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	// 判断队列是否已满
	if this.Length >= this.Size {
		return false
	}
	// 生成新节点
	newNode := &LinkedNode{
		Val: value,
		Next:this.Tail,
		Pre: this.Tail.Pre,
	}
	// 插入元素
	this.Tail.Pre.Next = newNode
	this.Tail.Pre = newNode
	this.Length++
	return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.Length == 0 {
		return false
	}
	this.Head.Next.Next.Pre = this.Head
	this.Head.Next = this.Head.Next.Next
	this.Length--
	return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.Length == 0 {
		return false
	}
	this.Tail.Pre.Pre.Next = this.Tail
	this.Tail.Pre = this.Tail.Pre.Pre
	this.Length--
	return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.Length == 0 {
		return -1
	}
	return this.Head.Next.Val
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.Length == 0 {
		return -1
	}
	return this.Tail.Pre.Val
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	if this.Length == 0 {
		return true
	}
	return false
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	if this.Length == this.Size {
		return true
	}
	return false
}