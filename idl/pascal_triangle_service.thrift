namespace go pascal   # 决定了golang生成代码的package名字

# 杨辉三角
service PascalTriangle {
    // 输入：给定一个非负整数 numRows
    // 输出：生成「杨辉三角」的前 numRows 行
    list<list<i32>> genPascalTriangle(1: i32 numRows)
}