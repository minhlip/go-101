package lesson

import "fmt"

// (1) In `Go`, `variables` are explicitly declared and used by the compiler to e.g. check type correctness of function calls.
func Variables() {

	// (2) `var` declares 1 or more variables.
	var a = "initial"
	fmt.Println(a)

	// (3) You can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// (4) Go will infer the type of initialized variables.
	var d = true
	fmt.Println(d)

	// (5) The `zero` value of a variable is `0` for `int`, `false` for `bool`, and `""` for `string`.
	var e int
	fmt.Println(e)

	var f string
	fmt.Println(f)

	// (6) The `:=` syntax is shorthand for declaring and initializing a variable, e.g. for `var f string = "apple"` in this case.
	g := "apple"
	fmt.Println(g)
}

/*
(1): Trong go, biến phải được khai báo rõ ràng, Compiler sẽ kiểm tra kiểu dữ liệu khi gọi hàm
(2): `var` khai báo 1 hoặc nhiều biến.
(3): Có thể khai báo nhiều biến cùng lúc.
(4): Type inference (tự suy luận kiểu): Go sẽ suy luận kiểu dữ liệu của biến đã khởi tạo từ giá trị khởi tạo. `d` được suy luận là `bool` vì giá trị là true
(5): Zero-values (Giá trị mặc định): Zero value là giá trị mặc định khi khai báo không gán giá trị:
- `int`: 0
- `bool`: false
- `string`: ""
(6): Short variable declaration (khai báo biến tắt): `:=` là cách viết tắt ngắn gọn để khai báo và khởi tạo biến, ở ví dụ trên thì tương đương `var g string = "apple"`. Chỉ được dùng trong function, không được dùng ở package level.
*/

/*`Variables` function output:
initial
1 2
true
0
""
apple
*/