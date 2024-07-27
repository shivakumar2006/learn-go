package main

// func main() {
// 	var array = [...]int {1, 2, 3, 4, 5, 69, 73};
// 	var array2 = [...]int {3, 51, 69, 69, 69, 73};
// 	var array3 = [...]int {5, 4, 3, 2, 1, 67, 68, 69, 70, 76};

// 	array[4] = 24;
// 	array2[0] = -6;
// 	array3[9] = -76;

// 	fmt.Println(array[1:6]);
// 	fmt.Println("Full Array: ", array);
// 	fmt.Println("Length of an Array: ", len(array), array);

// 	fmt.Println(array2[0:3]);
// 	fmt.Println("Full Array2: ", array2);

// 	fmt.Println(array3[2:7])
// 	fmt.Println("Full Array3: ", array3);
// 	fmt.Println("Length of an Array: ", len(array3), array3);
// }

//MULTIDIMENSIONAL ARRAYS => {}

// func main() {
// 	var array = [4][4]int { {1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}};
// 	fmt.Println("MultiDimensional Arrays: ", array);
// 	fmt.Println("Length of an Array: ", len(array));
// 	fmt.Println(array[2][3]);
// 	fmt.Println(array[3][2]);
// }

//MAP => {}

// func main() {
// empSal := make(map[string]int);  //Declaration
// empSal = map[string]int {		 // Initialization
// 	"Sachin": 30000,
// 	"Sudhanshu": 32000,
// 	"amit": 45000,
// }

// fmt.Println(empSal);

// empSal := map[string]int {		//Initialization and Declaration in same line.....
// 	"Sachin": 30000,
// 	"Sudhanshu": 32000,
//  	"amit": 45000,
// }

// fmt.Println(empSal);

//ADD NEW EMPLOYEE...... with the help of MAP => {}
// empSal := map[string]int {
// 	"Sachin": 30000,
// 	"Sudhanshu": 32000,
// 	"Amit": 45000,
// }

// empSal["Kashish"] = 23000;

// fmt.Println(empSal);

//Delete any Employee....with MAP => {}
// empSal := map[string]int {
// 	"Sachin": 30000,
// 	"Sudhanshu": 32000,
// 	"Amit": 45000,
// }

// empSal["Kashish"] = 23000;

// fmt.Println(empSal);
// delete(empSal, "Amit"); // THZis is a Delete Function in MAP...........!
// fmt.Println(empSal);

//Find Availability of KEY....
// 	empSal := map[string]int {
// 		"Sachin": 30000,
// 		"Sudhanshu": 32000,
// 		"Amit": 45000,
// 	}

// 	empSal["Kashish"] = 23000;

// 	delete(empSal, "Amit");
// 	fmt.Println(empSal);

// 	Amit, ok := empSal["Amit"];		// Check Availability...!
// 	fmt.Println(Amit, ok);

// 	Kashish, Check := empSal["Kashish"];	// Check Availability...!
// 	fmt.Println(Kashish, Check);
// }

// SLICING VS ARRAY.......!
// func main() {
// 	//Arrays
// 	// var array = [5]int {1, 2, 3, 4, 5};
// 	// var array2 = [...]int {1, 3, 2, 4, 6, 5, 7, 9};
// 	// THESE ARE THE TYPE OF ARRAY...

// 	//Slicing
// 	// var array = []int {1, 2, 3, 4, 5}; // THIS IS SLICING, it's Syntax Slightly different from the Arrays and work same as ARRAY.

// 	// fmt.Println(array);

// 	//CHECK CAPACITY....!!
// 	// var array = []int {1, 2, 3, 4, 5};
// 	// var array2 = make([]int, 1, 6);

// 	// fmt.Println("Capacity of array: ", cap(array));
// 	// fmt.Println("Length of array: ", len(array));
// 	// fmt.Println("Capacity of array2: ", cap(array2));
// 	// fmt.Println("Length of array2: ", len(array2));

// 	//ADD NEW ELEMENT IN ARRAYS USING APPEND...
// 	// var array = []int {1, 2, 3, 4, 5};
// 	// var array2 []int = append(array, 500);
// 	// var array3 []int = append(array, array2...);

// 	// fmt.Println("Array 3 = ", array3);
// }

// STRUCT......=>>>>>{}..!
// type employee struct {
// 	EmpId int;
// 	empName string;
// 	empMobile []string;
// }

// func main() {
// 	emp := employee{
// 		EmpId: 69,
// 		empName: "Amit",
// 		empMobile: []string {"9540611408", "1140895406"},
// 	}

// 	fmt.Println(emp);
// }

//LOOPS......________..!!!!
// func main() {
// 	for i := 1; i <= 8; i++ {
// 		fmt.Printf("value of i - %v \n", i);
// 	}
// } //FOR LOOP ARE THE ONLY LOOP WHICH ARE PRESENT IN GO.

//IF-ELSE.....
// func main() {
// 	for i := 0; i<=10; i++ {
// 		if i % 2 == 0 {
// 			fmt.Printf("even number: %v \n", i);
// 		} else {
// 			fmt.Printf("odd number: %v \n", i);
// 		}
// 	}
// }

// SWITCH.........__!!
// func main() {
// 	switch 6 {
// 	case 4:
// 		fmt.Println("Shiva");
// 	case 3:
// 		fmt.Println("Kumar");
// 	case 6:
// 		fmt.Println("Shiva Kumar");
// 	}
// }

// DEFER....._____!!!!
// func main() {
// 	// defer fmt.Println("i am");
// 	// defer fmt.Println("Shiva");
// 	// defer fmt.Println("Kumar");

// 	//  fmt.Println("i am");
// 	// defer fmt.Println("Shiva");
// 	//  fmt.Println("Kumar");
//

// POINTERS.......
// func main() {
// 	x := 99;
// 	var y *int = &x;

// 	fmt.Println(x);
// 	fmt.Println(y); // this give an address of a value where it can store
// 	fmt.Println(*y); // this give a value of that specific variable.
// }

//FUNCTION.......!!!
// func main() {
// 	fmt.Println(show());
// 	fmt.Println(number());
// 	fullName();
// 	summation(52, 48);
// 	add(2, 5, 8, 3, 5);

// 	mul, add := calculate(5, 4);
// 	fmt.Println(mul);
// 	fmt.Println(add);
// }

// func show() string {
// 	return "Shiva";
// }

// func number() int {
// 	return 45;
// }

// func fullName() {
// 	fmt.Println("Shiva Kumar");
// }

// func summation(x int, y int) {
// 	fmt.Println(x + y);
// }

// func add(values ... int) {
// 	sum := 0;
// 	for _, n := range values {
// 		sum += n;
// 	}
// 	fmt.Println(sum);
// }

// //When we want to perform 2 or more operations in on efunction =>
// func calculate(x int, y int) (int, int) {
// 	return (x * y), (x + y);
// }

//Test method in Functions..... Also known as Anonymus Function...
// type Employee struct {
// 	EmpName string;
// }

// func main() {
// 	emp := Employee{"Amit"};
// 	emp.test();
// }

// func (emp Employee) test() {
// 	fmt.Println(emp.EmpName);
// }

//Declaration and Call in same function....
// func main() {
// 	// f := func() {
// 	// 	fmt.Println("Shiva Kumar");
// 	// }

// 	// f();

// }

//INTERFACE IN GO.......
// func main() {
// 	var object Vehicle = Bike {
// 		Name: "Ninja ZX10R",
// 		Color: "Green",
// 		Price: 200000,
// 	}

// 	object.showDetails();
// 	fmt.Println(object.showName())

// }

// type Vehicle interface {
// 	showDetails();
// 	showName() string;
// }
// type Bike struct {
// 	Name, Color string;
// 	Price float64;
// }

// func (bike Bike) showDetails() {
// 	fmt.Println("Bike Name: ", bike.Name);
// 	fmt.Println("Bike Color ", bike.Color);
// 	fmt.Println("Bike Price ", bike.Price);
// }

// func (bike Bike) showName() string {
// 	return bike.Name;
// }


