package main

import (
	"container/list"
	"fmt"
)

func main() {
	//var linkList list.List
	//
	//// 在链表末尾添加元素 "go"
	//linkList.PushBack("go")
	//
	//// 在链表开头添加元素 "java"
	//linkList.PushFront("java")
	//
	//// 再次在链表末尾添加元素 "python"
	//linkList.PushBack("python")
	//
	//// 打印链表长度
	//fmt.Println("链表长度:", linkList.Len())
	//
	//// 遍历并打印链表元素
	//for i := linkList.Front(); i != nil; i = i.Next() {
	//	fmt.Println(i.Value)
	//}
	//
	//// 从后向前遍历并打印链表元素
	//for i := linkList.Back(); i != nil; i = i.Prev() {
	//	fmt.Println(i.Value)
	//}
	//
	//// 获取并打印链表的第一个和最后一个元素
	//fmt.Println("第一个元素:", linkList.Front().Value)
	//fmt.Println("最后一个元素:", linkList.Back().Value)

	//ll := list.New()
	//ll.PushBack(1)
	//ll.PushBack(2)
	//ll.PushBack(3)
	//
	//fmt.Println("Original list:")
	//printList(ll)
	//
	//moveLastToFront(ll)

	//l1 := list.New()
	//l1.PushBack(1)
	//l1.PushBack(2)
	//l1.PushBack(3)
	//
	//l2 := list.New()
	//l2.PushBack(4)
	//l2.PushBack(5)
	//l2.PushBack(6)
	//
	//LinkListsDirectly(l1, l2)

	l1 := list.New()
	l1.PushBack(1)
	l1.PushBack(2)
	l1.PushBack(3)

	res := reverseList(l1)

	printList(res)
}

func operation() {
	// 初始化一个链表
	courseList := list.New()
	// 填入的是值，PushBack 会帮助你封装成一个节点
	courseList.PushBack("java")
}

func moveLastToFront(ll *list.List) {
	if ll.Len() > 1 {
		lastElement := ll.Back()
		ll.MoveToFront(lastElement)
	}
}

func printList(ll *list.List) {
	for e := ll.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func LinkListsDirectly(l1, l2 *list.List) {
	if l2.Len() > 0 {
		l1.PushBackList(l2)
	}
}

func removeElements(ll *list.List, value interface{}) {
	for e := ll.Front(); e != nil; {
		next := e.Next() // 保存下一个元素的指针
		if e.Value == value {
			ll.Remove(e) // 删除当前元素
		}
		e = next // 移动到下一个元素
	}
}

func updateElements(ll *list.List, oldValue, newValue interface{}) {
	for e := ll.Front(); e != nil; e = e.Next() {
		if e.Value == oldValue {
			e.Value = newValue
		}
	}
}

func reverseList(ll *list.List) *list.List {
	reversed := list.New()
	for e := ll.Front(); e != nil; {
		next := e.Next()            // 保存下一个元素的引用
		ll.Remove(e)                // 从原链表中移除当前元素
		reversed.PushFront(e.Value) // 将元素添加到新链表的前端
		e = next
	}
	return reversed
}
