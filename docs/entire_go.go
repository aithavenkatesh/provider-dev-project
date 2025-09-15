// golang_topics.go
// This file contains important Go (Golang) topics with example code.
// You can run each section's code by isolating it or commenting out other parts.

package main

import (
    "errors"  // For creating custom error messages
    "fmt"     // For formatted I/O (input/output)
    "math"    // For mathematical operations (e.g., Pi)
    "time"    // For time-related functions (e.g., Sleep)
    // "example.com/myproject/mypackage" // Uncomment this if you set up the 'Packages and Modules' example
)

// --- 1. Basic Structure & "Hello World!" ---
// Every executable Go program starts with a 'package main' and a 'main' function.
func helloWorld() {
    fmt.Println("\n--- 1. Basic Structure & 'Hello World!' ---")
    fmt.Println("Hello, Go!")
}

// --- 2. Variables & Data Types ---
// Variables store data. Go is statically typed, meaning types are checked at compile time.
func variablesAndDataTypes() {
    fmt.Println("\n--- 2. Variables & Data Types ---")
    // Using 'var' declaration
    var greeting string = "Hi there"
    var number int = 100
    var isComplete bool // Defaults to false

    fmt.Println("Greeting (var):", greeting)
    fmt.Println("Number (var):", number)
    fmt.Println("Is Complete (var):", isComplete)

    // Using short variable declaration (':=') - common inside functions
    product := "Laptop"
    price := 1200.50
    isOnSale := true

    fmt.Println("Product (:=):", product)
    fmt.Println("Price (:=):", price)
    fmt.Println("On Sale (:=):", isOnSale)
}

// --- 3. Constants ---
// Constants are values that cannot be changed after declaration.
func constants() {
    fmt.Println("\n--- 3. Constants ---")
    const PI = 3.14159        // Untyped constant
    const Language string = "Go" // Typed constant

    radius := 5.0
    area := PI * radius * radius

    fmt.Println("The area is:", area)
    fmt.Println("Programming Language:", Language)
}

// --- 4. Control Flow (If/Else, For, Switch) ---
// These constructs control the order of code execution.
func controlFlow() {
    fmt.Println("\n--- 4. Control Flow ---")

    // a) if-else Statements
    fmt.Println("\n--- a) if-else Statements ---")
    age := 18
    if age >= 18 {
        fmt.Println("You are an adult.")
    } else {
        fmt.Println("You are a minor.")
    }

    // 'if' with a short statement (variable 'num' is scoped to if/else)
    if num := 10; num%2 == 0 {
        fmt.Println(num, "is even.")
    } else {
        fmt.Println(num, "is odd.")
    }

    // b) for Loops
    fmt.Println("\n--- b) for Loops ---")
    // Standard for loop
    for i := 0; i < 3; i++ {
        fmt.Println("Count:", i)
    }

    // For loop as a 'while' loop
    sum := 1
    for sum < 10 {
        sum += sum // sum = 2, 4, 8, 16
    }
    fmt.Println("Sum (while-like):", sum)

    // For-range loop (for iterating over collections)
    fruits := []string{"apple", "banana"}
    for index, fruit := range fruits {
        fmt.Printf("Fruit at index %d is %s\n", index, fruit)
    }

    // c) switch Statements
    fmt.Println("\n--- c) switch Statements ---")
    day := "Wednesday"
    switch day {
    case "Monday", "Tuesday":
        fmt.Println("Start of the week.")
    case "Wednesday":
        fmt.Println("Mid-week!")
    default:
        fmt.Println("Weekend or other day.")
    }

    score := 85
    switch { // Switch without a tag (like if/else if chain)
    case score >= 90:
        fmt.Println("Grade A")
    case score >= 80:
        fmt.Println("Grade B")
    default:
        fmt.Println("Grade C or lower")
    }
}

// --- 5. Functions ---
// Functions are blocks of code to perform specific tasks.
// Simple function
func add(a, b int) int { // Parameters with same type can be grouped (a, b int)
    return a + b
}

// Function with multiple return values
func greet(name string, age int) (string, int) {
    return fmt.Sprintf("Hello, %s!", name), age
}

// Function with named return values and bare return
func calculate(x, y int) (sum int, diff int) {
    sum=x + y
    diff=x - y
    return // Returns the current values of sum and diff
}

// Variadic function: accepts a variable number of arguments
func sumAll(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func functions() {
    fmt.Println("\n--- 5. Functions ---")
    fmt.Println("Sum (5, 3):", add(5, 3))

    message, userAge := greet("Alice", 30)
    fmt.Printf("%s You are %d years old.\n", message, userAge)

    s, d := calculate(10, 4)
    fmt.Printf("Calc: Sum: %d, Difference: %d\n", s, d)

    fmt.Println("SumAll (1,2,3):", sumAll(1, 2, 3))
    fmt.Println("SumAll (1,2,3,4,5):", sumAll(1, 2, 3, 4, 5))
}

// --- 6. Data Structures (Arrays, Slices, Maps, Structs) ---
func dataStructures() {
    fmt.Println("\n--- 6. Data Structures ---")

    // a) Arrays: Fixed-size sequence of elements of the same type.
    fmt.Println("\n--- a) Arrays ---")
    var numbers [5]int // Declares an array of 5 ints, initialized to zeros
    numbers[0] = 10
    numbers[4] = 50
    fmt.Println("Array:", numbers)

    primes := [3]int{2, 3, 5} // Declares and initializes
    fmt.Println("Primes:", primes)

    // b) Slices: Go's dynamic arrays (more common than arrays).
    fmt.Println("\n--- b) Slices ---")
    var colors []string // Declares a nil slice
    fmt.Printf("Initial colors: %v, Length: %d, Capacity: %d\n", colors, len(colors), cap(colors))

    colors=append(colors, "Red", "Green", "Blue")
    fmt.Printf("Colors after append: %v, Length: %d, Capacity: %d\n", colors, len(colors), cap(colors))

    someColors := colors[1:3] // Slicing: elements from index 1 up to (but not including) 3
    fmt.Println("Some colors (sliced):", someColors)

    scores := make([]int, 3, 5) // Length 3, Capacity 5
    scores[0], scores[1], scores[2] = 90, 75, 88
    fmt.Printf("Scores (make): %v, Length: %d, Capacity: %d\n", scores, len(scores), cap(scores))

    // c) Maps: Unordered collections of key-value pairs.
    fmt.Println("\n--- c) Maps ---")
    ages := make(map[string]int) // Keys are strings, values are ints
    ages["Alice"] = 30
    ages["Bob"] = 24
    fmt.Println("Ages map:", ages)

    bobAge := ages["Bob"]
    fmt.Println("Bob's age:", bobAge)

    dianaAge, ok := ages["Diana"] // Check if key exists
    if ok {
        fmt.Println("Diana's age:", dianaAge)
    } else {
        fmt.Println("Diana not found.")
    }
    delete(ages, "Bob") // Delete an element
    fmt.Println("Ages after deleting Bob:", ages)

    // d) Structs: User-defined composite types that group named fields.
    fmt.Println("\n--- d) Structs ---")
    type Person struct {
        FirstName string
        LastName  string
        Age       int
        IsEmployed bool
    }

    p1 := Person{FirstName: "John", LastName: "Doe", Age: 30, IsEmployed: true}
    fmt.Println("Person 1:", p1)
    fmt.Println("Person 1 First Name:", p1.FirstName)
}

// --- 7. Pointers ---
// Pointers store the memory address of a value.
func increment(num *int) { // Takes a pointer to an int
    *num++ // Dereference num and increment the value it points to
}

func pointers() {
    fmt.Println("\n--- 7. Pointers ---")
    value := 10
    fmt.Println("Initial value:", value)

    ptr := &value // Get the address of 'value'
    fmt.Println("Value via pointer (*ptr):", *ptr)
    fmt.Println("Address of value (ptr):", ptr)

    *ptr = 20 // Change value through the pointer
    fmt.Println("New value (after *ptr = 20):", value)

    increment(&value) // Pass the address of 'value'
    fmt.Println("Value after increment function:", value)
}

// --- 8. Methods & Interfaces ---
// a) Methods: Functions associated with a specific type (receiver).
type Circle struct {
    Radius float64
}

// Area method with a value receiver
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

// Scale method with a pointer receiver (can modify the original struct)
func (c *Circle) Scale(factor float64) {
    c.Radius *= factor
}

// b) Interfaces: Define a set of method signatures. Types implicitly implement them.
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// printShapeDetails accepts any type that implements the Shape interface
func printShapeDetails(s Shape) {
    fmt.Printf("Type: %T, Area: %.2f, Perimeter: %.2f\n", s, s.Area(), s.Perimeter())
}

func methodsAndInterfaces() {
    fmt.Println("\n--- 8. Methods & Interfaces ---")

    fmt.Println("\n--- a) Methods ---")
    myCircle := Circle{Radius: 10}
    fmt.Printf("Initial Circle Area: %.2f\n", myCircle.Area())
    myCircle.Scale(2) // Modifies myCircle.Radius
    fmt.Printf("Scaled Circle Area: %.2f\n", myCircle.Area())

    fmt.Println("\n--- b) Interfaces ---")
    rect := Rectangle{Width: 10, Height: 5}
    circ := Circle{Radius: 7} // Note: Circle from above also fits the Shape interface
    printShapeDetails(rect)
    printShapeDetails(circ)

    shapes := []Shape{rect, circ} // Slice of interfaces
    fmt.Println("\nIterating through shapes:")
    for _, s := range shapes {
        printShapeDetails(s)
    }
}

// --- 9. Concurrency (Goroutines & Channels) ---
// Go's standout feature for concurrent programming.

// a) Goroutines: Lightweight, independently executing functions.
func say(s string) {
    for i := 0; i < 2; i++ {
        time.Sleep(50 * time.Millisecond)
        fmt.Println(s)
    }
}

// b) Channels: Typed conduits for goroutine communication and synchronization.
func worker(done chan bool, id int) {
    fmt.Printf("Worker %d: Working...\n", id)
    time.Sleep(time.Second) // Simulate work
    fmt.Printf("Worker %d: Work done.\n", id)
    done <- true // Signal completion
}

func concurrency() {
    fmt.Println("\n--- 9. Concurrency (Goroutines & Channels) ---")

    fmt.Println("\n--- a) Goroutines ---")
    go say("world") // Runs in a new goroutine
    say("hello")    // Runs in the main goroutine
    // Output order might be interleaved. Main goroutine might exit before 'world' finishes.
    time.Sleep(time.Millisecond * 200) // Give 'world' a chance to finish

    fmt.Println("\n--- b) Channels ---")
    done := make(chan bool, 1) // Buffered channel for signaling
    go worker(done, 1)
    <-done // Block until worker 1 sends true

    messages := make(chan string) // Unbuffered channel for data exchange
    go func() {
        messages <- "ping" // Send "ping"
        messages <- "pong"
    }()
    msg1 := <-messages // Receive "ping"
    msg2 := <-messages // Receive "pong"
    fmt.Println(msg1, msg2)
}

// --- 10. Error Handling & Defer ---
// Go handles errors by returning an 'error' type. 'defer' schedules cleanup.

// a) Error Handling: Functions return 'error' as the last return value.
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

// b) defer Statement: Defers execution of a function until the surrounding function returns.
func cleanup() {
    fmt.Println("  (Defer) Cleaning up resources.")
}

func processWithDefer() {
    defer cleanup() // This runs just before processWithDefer() exits
    fmt.Println("Starting processing...")
    fmt.Println("Processing completed.")
}

func errorHandlingAndDefer() {
    fmt.Println("\n--- 10. Error Handling & Defer ---")

    fmt.Println("\n--- a) Error Handling ---")
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result of 10/2:", result)
    }

    result, err=divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result of 10/0:", result)
    }

    fmt.Println("\n--- b) defer Statement ---")
    processWithDefer()
    fmt.Println("Main function continues after processWithDefer.")
}

// --- 11. Packages and Modules ---
// Programs are organized into packages. Modules manage dependencies.
// To fully run this, you'd need to set up the file structure described in the original response.
// For this single file, we'll just demonstrate importing a standard package.

// func packagesAndModules() {
// 	fmt.Println("\n--- 11. Packages and Modules ---")
// 	// Example of importing a standard package (fmt and others already imported)
// 	fmt.Println("Using 'fmt' package for printing.")

// 	// To use your own package 'mypackage' from 'example.com/myproject/mypackage':
// 	// 1. Create a directory structure:
// 	//    your_project_root/
// 	//    ├── go.mod (module example.com/myproject)
// 	//    ├── main.go (this file)
// 	//    └── mypackage/
// 	//        └── greetings.go
// 	// 2. In mypackage/greetings.go:
// 	//    package mypackage
// 	//    import "fmt"
// 	//    func Hello(name string) string { return fmt.Sprintf("Hello, %s!", name) }
// 	//    func Goodbye(name string) string { return fmt.Sprintf("Goodbye, %s!", name) }
// 	// 3. Uncomment the import for "example.com/myproject/mypackage" at the top.
// 	// 4. Uncomment the lines below.
// 	//
// 	// fmt.Println(mypackage.Hello("Charlie"))
// 	// fmt.Println(mypackage.Goodbye("Diana"))
// }

// main function to call all topic examples
func main() {
    fmt.Println("--- Starting Go Topics Overview ---")

    helloWorld()
    variablesAndDataTypes()
    constants()
    controlFlow()
    functions()
    dataStructures()
    pointers()
    methodsAndInterfaces()
    concurrency()
    errorHandlingAndDefer()
    // packagesAndModules() // Uncomment if you set up the package example

    fmt.Println("\n--- Go Topics Overview Finished ---")
}