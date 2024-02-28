package data

import "goLearning/protobuf/basic1/proto"

var LanauageCourses = []*proto.CourseInfo{
	{
		Cid:     1,
		Cname:   "Go",
		Teacher: "张三",
	}, {
		Cid:     2,
		Cname:   "Python",
		Teacher: "李四",
	}, {
		Cid:     3,
		Cname:   "Java",
		Teacher: "王五",
	},
}

var ComputerCourses = []*proto.CourseInfo{
	{
		Cid:     4,
		Cname:   "计算机基础",
		Teacher: "赵六",
	}, {
		Cid:     5,
		Cname:   "数据结构",
		Teacher: "孙七",
	}, {
		Cid:     6,
		Cname:   "算法",
		Teacher: "周八",
	},
}
