// 包含头文件
// include "shared.thrift"

namespace cpp tutorial	// 指定目标语言

// 定义变量
const i32 INT32CONSTANT = 9853

// 定义结构体
struct Work {
	1: i32 num1 = 0,
    2: i32 num2,
	3: i32 op,
	4: optional string comment,
}

// 定义服务
service Calculator {
	/**
	* A method definition looks like C code. It has a return type, arguments,
	* and optionally a list of exceptions that it may throw. Note that argument
	* lists and exception lists are specified using the exact same syntax as
	* field lists in struct or exception definitions.
	*/
	void ping(),

	/**
	* 加法服务
	*/
	i32 add(1:i32 num1, 2:i32 num2),

	i32 calculate(1:i32 logid, 2:Work w),

	/**
	* This method has a oneway modifier. That means the client only makes
	* a request and does not listen for any response at all. Oneway methods
	* must be void.
	*/
	oneway void zip()
}
