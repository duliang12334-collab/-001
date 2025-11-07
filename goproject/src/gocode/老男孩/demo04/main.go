package main

import "fmt"

func main() {
	//定义一个数组并赋值
	var name = [3]string{"张三", "李四", "王五"}
	fmt.Println(name)
	//定义一个切片
	var ages = []int{32, 33, 34}
	fmt.Println(ages)
	// 数组 切片都是按索引顺序储存，而map没有

	fmt.Println("----------------------------------")

	//定义一个map 映射 键值对
	var info = make(map[string]string)
	info["name"] = "费六"
	info["age"] = "88"
	fmt.Println(info)
	info2 := info           //进行了引用拷贝
	info["gender"] = "male" //info变 info2也跟着变，指向了同一个值
	fmt.Println(info2)      // 结果 map[age:88 gender:male name:费六]

	fmt.Println("----------------------------------")

	//map 的增删改查
	//1.查询
	var info3 = make(map[string]string)
	info3["name"] = "解七"
	info3["age"] = "99"
	// 变量名 + key 查询年龄
	fmt.Println(info3["age"])
	// 通过exist 判断键值是否存在
	value, exist := info["gender2"] //exist的：意思是否存在
	fmt.Println(value, exist)
	if exist {
		fmt.Println("性别存在")
	} else {
		fmt.Println("性别不存在")
	}
	//2.添加操作
	//如果原map中没有相同键值，就直接变量名添加，上info3没有gender2，直接添加，例如：
	info3["gender2"] = "male"
	info3["name2"] = "狗蛋"
	info3["age2"] = "10"
	info3["name3"] = "翠花"
	info3["age3"] = "20"
	fmt.Println(info3)

	//3.修改键值对
	//如果原map中有相同键值，可在原键值基础上直接进行修改，原年龄99 改为100，如下：
	info3["age"] = "100"
	fmt.Println(info3)

	//4.删除键值对
	//利用内置函数 delete+map对象，以及对应的 key
	delete(info3, "gender2")
	fmt.Println(info3)

	fmt.Println("----------------------------------")

	//map的遍历
	// 利用range针对元素个数进行循环，有多少个元素就循环多少次
	for k, v := range info3 {
		fmt.Println(k, v)
	}

	fmt.Println("----------------------------------")

	//map的嵌套
	//1. 切片
	var x = []int{1, 2, 3}
	fmt.Println(x)
	// 取值
	fmt.Println(x[1])

	//2.切片里面放map
	stu01 := map[string]string{"name": "rain", "age": "18"}
	stu02 := map[string]string{"name": "kevin", "age": "19"}
	stu03 := map[string]string{"name": "tony", "age": "20"}
	var stus = []map[string]string{stu01, stu02, stu03}
	fmt.Println(stus)
	//取值
	fmt.Println(stus[0]) // map[age:18 name:rain]
	//取name 键值
	fmt.Println(stus[0]["name"]) //rain

	fmt.Println("----------------------------------")

	//3. map 嵌套 map
	var date = make(map[string]map[string]string)
	date["niuzi"] = map[string]string{"age": "22", "gender": "male"}
	date["niuniu"] = map[string]string{"age": "32", "gender": "male"}
	fmt.Println(date["niuzi"]["age"])

}
