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
