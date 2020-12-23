package utils

import (
	queue2 "bitcoin/queue"
	"fmt"
	"os"
	"strings"
)

var format_aToA = make(map[string]string)
var format_AToa = make(map[string]string)
var count int

func init() {
	format_aToA["a"] = "A"
	format_aToA["b"] = "B"
	format_aToA["c"] = "C"
	format_aToA["d"] = "D"
	format_aToA["e"] = "E"
	format_aToA["f"] = "F"
	format_aToA["g"] = "G"
	format_aToA["h"] = "H"
	format_aToA["i"] = "I"
	format_aToA["j"] = "J"
	format_aToA["k"] = "K"
	format_aToA["l"] = "L"
	format_aToA["m"] = "M"
	format_aToA["n"] = "N"
	format_aToA["o"] = "O"
	format_aToA["p"] = "P"
	format_aToA["q"] = "Q"
	format_aToA["r"] = "R"
	format_aToA["s"] = "S"
	format_aToA["t"] = "T"
	format_aToA["u"] = "U"
	format_aToA["v"] = "V"
	format_aToA["w"] = "W"
	format_aToA["x"] = "X"
	format_aToA["y"] = "Y"
	format_aToA["z"] = "Z"

	format_AToa["A"] = "a"
	format_AToa["B"] = "b"
	format_AToa["C"] = "c"
	format_AToa["D"] = "d"
	format_AToa["E"] = "e"
	format_AToa["F"] = "f"
	format_AToa["G"] = "g"
	format_AToa["H"] = "h"
	format_AToa["I"] = "i"
	format_AToa["J"] = "j"
	format_AToa["K"] = "k"
	format_AToa["L"] = "l"
	format_AToa["M"] = "m"
	format_AToa["N"] = "n"
	format_AToa["O"] = "o"
	format_AToa["P"] = "p"
	format_AToa["Q"] = "q"
	format_AToa["R"] = "r"
	format_AToa["S"] = "s"
	format_AToa["T"] = "t"
	format_AToa["U"] = "u"
	format_AToa["V"] = "v"
	format_AToa["W"] = "w"
	format_AToa["X"] = "x"
	format_AToa["Y"] = "y"
	format_AToa["Z"] = "z"

	count = 1
}

//非递归生成反序列化语句代码
func MyUnmarshal(data map[string]interface{}, entityName interface{}) {
	for key, value := range data {
		str := ""
		for i := 1; i < len(key); i++ {
			str += string(key[i])
		}
		fmt.Println(fmt.Sprintf("%v.%v%v = res[\"%v\"].(%T) ", entityName, format_aToA[string(key[0])], str, key, value))
	}
}

//递归生成反序列化语句代码,data为map类型的数据,
//entityName为N级实体 类型名,内部转换为string后使用,
//name为为N+1级实体 类型名,
func MyUnmarshals1(data map[string]interface{}, entityName string, name string, dataName string) {
	//for key, value := range data {
	//	str := ""
	//	for i := 1; i < len(key); i++ {
	//		str += string(key[i])
	//	}
	//
	//	valueType, ok := value.(map[string]interface{})
	//	if ok {
	//		str := ""
	//		for i := 1; i < len(key); i++ {
	//			str += string(key[i])
	//		}
	//		//str = name + "." + format_aToA[string(key[0])] + str
	//		//name1 := name
	//
	//		fmt.Println(fmt.Sprintf("%v.%v%v_ = %v_[\"%v\"].(%T)", entityName, format_aToA[string(key[0])], str, name, key, value))
	//		//entityName +=
	//		//fmt.Println(fmt.Sprintf("%v%v____ = %v[\"%v\"].(%T)", format_AToa[string(entityName.(string)[0])], str, entityName, key, valueType))
	//		//fmt.Println(fmt.Sprintf("%v%vBytes, ok := %v[\"%v\"].(%T)", format_AToa[string(entityName.(string)[0])], str, entityName, key, valueType))
	//		//fmt.Println("if !ok { return blockChainInfo }")
	//		//name += key
	//		MyUnmarshals(valueType, str,name,dataName)
	//	} else {
	//		if count == 1 {
	//			fmt.Println(fmt.Sprintf("%v.%v%v = %v[\"%v\"].(%T)", entityName, format_aToA[string(key[0])], str, dataName, key, value))
	//		}else {
	//			//str := "_" + strconv.Itoa(count)
	//			//dataName += str
	//			fmt.Println(fmt.Sprintf("%v.%v%v = %v[\"%v\"].(%T)", entityName, format_aToA[string(key[0])], str, dataName, key, value))
	//		}
	//	}
	//}

	//queue := queue.InitCircleQueue()
	//err := queue.Push(data)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//queue.Push(1)
	//for !queue.IsEmpty() {
	//	num, _ := queue.GetTail()
	//	for i := num.(int); i >= 0 ; i-- {
	//		var value interface{}
	//		var err error
	//		//出队一个元素
	//		if i == 0 {
	//			queue.Pull()
	//			break
	//		}else {
	//			value, err = queue.Pull()
	//			//fmt.Println("出队元素!")
	//			if err != nil {
	//				fmt.Println(err)
	//				return
	//			}
	//		}
	//
	//
	//		//断言出队元素
	//		mapPullValues, ok := value.(map[string]interface{})
	//		if !ok {
	//			fmt.Println("出队列断言map失败!")
	//			return
	//		}
	//
	//		count := 0 //计数器,用于统计当前所遍历的map里有多少个map
	//		//遍历出队元素(map)并进行断言; 若果出队map中还有是map类型的数据,则将次数据入队,否侧进行输出语句
	//		for key, value_range := range mapPullValues {
	//			mapValues, ok := value_range.(map[string]interface{})
	//			str := ""
	//			for i := 1; i < len(key); i++ {
	//				str += string(key[i])
	//			}
	//			//fmt.Println(mapValues)
	//			if ok {
	//				//入队
	//				//fmt.Println("入队的元素key:", key)
	//				count++
	//				err := queue.Push(mapValues)
	//				fmt.Println(fmt.Sprintf("%v.%v%v_ = %v[\"%v\"].(%T)", entityName, format_aToA[string(key[0])], str, dataName, key, value_range))
	//				if err != nil {
	//					fmt.Println("遍历元素入队出错:", err)
	//				}
	//				//fmt.Println(queue.Len())
	//			}else {
	//				fmt.Println(fmt.Sprintf("%v.%v%v = %v[\"%v\"].(%T)", entityName, format_aToA[string(key[0])], str, dataName, key, value_range))
	//			}
	//		}
	//		queue.Push(count)
	//
	//	}
	//}
}

func MyUnmarshals(data map[string]interface{}, entityName string, dataName string) {
	mapQueue := queue2.InitCircleQueue()
	entityQueue := queue2.InitCircleQueue()
	count := 0
	for {
		var key string
		var value interface{}
		var str string
		for key, value = range data {
			str = ""
			for i := 1; i < len(key); i++ {
				str += string(key[i])
			}
			_, ok := value.(map[string]interface{})
			if ok {
				mapQueue.Push(key)
			} else {

				if count == 0 {
					fmt.Println(fmt.Sprintf("%v.%v%v = %v[\"%v\"].(%T)", entityName, format_aToA[string(key[0])], str, dataName, key, value))
				} else {
					fmt.Println(fmt.Sprintf("%v.%v%v = %v_[\"%v\"].(%T)", entityName, format_aToA[string(key[0])], str, entityName, key, value))
				}
			}
		}

		outKey, err := mapQueue.Pull()
		if err != nil {
			break
		} else {
			str = ""
			for i := 1; i < len(outKey); i++ {
				str += string(outKey[i])
			}

			entityStr := entityName + "." + format_aToA[string(outKey[0])] + str
			entityQueue.Push(entityStr)
			if count != 0 {
				fmt.Printf("}\n")
			}

			fmt.Println(fmt.Sprintf("%v.%v%v_, ok == %v[\"%v\"].(%T)", entityName, format_aToA[string(outKey[0])], str, dataName, outKey, data))

			fmt.Printf("if ok {\n")
			count++
		}
		value = data[outKey]
		m, ok := value.(map[string]interface{})
		if ok {
			data = m
			entityName, err = entityQueue.Pull()
			if err != nil {
				return
			}
		}

	}
	fmt.Printf("}\n")
}

//非递归生成比特币客户端返回的结果data数据反序列化一级结构体,并自动写入文件中
func EntityOne(data map[string]interface{}, entityName string, path string) error {
	file, err := os.OpenFile(path, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// 查找文件末尾的偏移量
	n, err := file.Seek(0, 2)
	if err != nil {
		return err
	}

	entityStr := ""
	entityStr += fmt.Sprintf("\n\ntype %v struct {\n", entityName)
	//fmt.Printf("type %v struct {\n",entityName)
	for key, value := range data {
		str := ""
		for i := 1; i < len(key); i++ {
			str += string(key[i])
		}
		entityStr += fmt.Sprintf("\t%v%v %T\n", format_aToA[string(key[0])], str, value)
		//fmt.Println(fmt.Sprintf("%v%v %T", format_aToA[string(key[0])], str,value))
	}
	entityStr += "}"
	//fmt.Printf("}\n")
	// 从末尾的偏移量开始写入内容
	fmt.Println(entityStr)
	_, err = file.WriteAt([]byte(entityStr), n)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

//递归生成比特币客户端返回的结果data数据反序列化多级结构体,调用后须手动复制到相应的文件中
func Entitys(data map[string]interface{}, entityName string) {
	for key, value := range data {
		valueType, ok := value.(map[string]interface{})
		if ok {
			str := ""
			for i := 1; i < len(key); i++ {
				str += string(key[i])
			}
			str = fmt.Sprintf("%v", format_aToA[string(key[0])]) + str
			Entitys(valueType, str)
		}
	}
	fmt.Printf("type %v struct {\n", entityName)
	for key, value := range data {
		str := ""
		for i := 1; i < len(key); i++ {
			str += string(key[i])
		}
		valueType, ok := value.(map[string]interface{})
		if ok {
			fmt.Println(fmt.Sprintf("%v%v %v%v", format_aToA[string(key[0])], str, format_aToA[string(key[0])], str))

			//给结构体套用结构体的字段名后加一个下划线,类型依然为map,配合自动生成的反序列化语句使用,以便获取所有内容
			fmt.Println(fmt.Sprintf("%v%v_ %T", format_aToA[string(key[0])], str, valueType))
		} else {
			fmt.Println(fmt.Sprintf("%v%v %T", format_aToA[string(key[0])], str, value))
		}

	}
	fmt.Printf("}\n")
}


func GetCommand() {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON("help", paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))

	s := rpcResult.Data.Result.(string)
	slices := strings.Split(s, "\n")

	CommandSlice := make([]string, 0)

	for i := 0; i <len(slices); i++ {
		slice_ := strings.Split(slices[i], " ")

		CommandSlice = append(CommandSlice, slice_[0])

	}

	fmt.Println("const (")
	for j := 0; j < len(CommandSlice); j++ {
		if len(CommandSlice[j]) < 3 {
			continue
		}

		s := ""

		for i := 0; i < len(CommandSlice[j]); i++ {
			s += format_aToA[string(CommandSlice[j][i])]
		}

		fmt.Println(fmt.Sprintf("\t%v = \"%v\"", s, CommandSlice[j]))
	}
	fmt.Println(")\n")
}
