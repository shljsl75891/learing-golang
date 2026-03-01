# Learning Go Programming language through [BootDev](https://boot.dev) platform

## Datatypes

- `bool`
- `int`
- `byte` = Just like `buffer` in nodejs. It can contains 8 bit of data
- `float64`
- `string`

## CPU Performance and Memory Management

- Go compiles faster than other compiled languages like C, Rust, C++ but its execution speed is slower than them due to go runtime. Whenever a go file is compiled, a small runtime is also always added into every go executable. This runtime is responsible for memory management, goroutine scheduling and other features of the language. But, it is much faster then interpreted languages like Python and JavaScript.

![](/assets/2026-02-28-10-53-16.png)

- In terms of execution speed, it is faster than interpreted languages like Python and JavaScript but slower than compiled languages like C, C++ and Rust due to the presence of go runtime (automatic memory management e.g.) which adds some overhead to the execution.

![](/assets/2026-02-28-11-28-15.png)

- In terms of memory usage as well, it is performant then Java, C# as it doesn't need a virtual machine to run but have automatic garbage collection like those, but is less performant then C, C++ and Rust as they don't have a runtime and automatic garbage collection but needs manual memory management. It also does allocate more data in stack more often than heap which is faster to access and manage.

![](/assets/2026-02-28-10-55-46.png)

#### Memory Layout of programs

![](/assets/2026-02-28-11-06-19.png)

1. **Stack**: It is a region of memory that stores local variables, function parameters, return addresses. It is typically faster to access than the heap because of its contiguous memory allocation. It has limited size and can overflow if too much data is stored on it (like deep recursion), which is known as a stack overflow. It is automatically allocated when functions are called and deallocated when they return.

2. **Heap**: It is a region of memory that is used for dynamic memory allocation. It is managed by the runtime and can grow or shrink as needed. The heap is used for storing data that needs to persist beyond the scope of a function, such as objects or data structures. It is typically slower to access than the stack because of its non-contiguous memory allocation. It can also lead to memory leaks if not managed properly, and it required GC overhead to manage memory.
3. **Data Segment**: It is a region of memory that stores global variables and static variables. It is divided into two parts: the initialized data segment, which contains variables that are initialized with a value, and the uninitialized data segment (also known as the BSS segment), which contains variables that are not initialized. Its lifecycle is the entire duration of the program.

4. **Text Segment**: It is a region of memory that contains the executable code of the program. It is typically read-only and is shared among all instances of the program. Its lifecycle is the entire duration of the program.

## Constants in Go

Constants are the variables which are immutable while the program is running. The walrus operator `:=` cannot be used to declare constants. Go has kept this syntactical difference to make it clear that the value is a constant and its value must be known at the compile time, but the `var` are mutable and their values can be changed while runtime. We can use constants for computed values as long as they can be known to compiler at compile time.

![](/assets/2026-02-28-11-22-46.png)

## Printing in Go

Go follows C style of printing using `fmt` package along with formatting verbs like `%v`, `%d`, `%f` etc.

## Defer keyword in Go

`defer` keyword is a special feature in Go lang to defer a function's execution until the enclosing function returns. All deferred functions are executed in `LIFO`. Think of like all deferred functions goes into stack, and each function is popped and executed one by one before enclosing function returns itself. It is best way to run a piece of code just before each return statement by writting it once like closing db connections, or cleaning up resources.

## Datastructures in GO

#### Structs

Structs are custom data types that allow you to group together related data. They are similar to objects in JS or dictionaries in Python. Structs can also contains other structs (known as nested structs) as fields for creating complex data structures.

```go
type message struct { // Uppercase Name = Exported (Public), Lowercase Name = Unexported (Private)
    text: string
    sender: user
    recipient: user
}

type user struct {
    name string
    age int
}

type car struct {
  brand string
  model string
  doors int
  mileage int
  // wheel is a field containing an anonymous struct = Always prefer named struct. Use this only when you need it just once, and not reusable anywhere else in the codebase.
 wheel struct {
    radius int
    material string
  }
}

type car struct {
  brand string
  model string
}

type truck struct {
    // embedded struct for just data only inheritance (actually composition). the fields of embedded struct can be accessed directly from the top level unlike nested struct
    car
    bedSize int
}

// Methods on structs
func (c car) upgradeModel() {
    c.model = "New Model" // This will not change original model, as the whole struct is passed by value (completely new copy of struct is passed)
}

func (c *car) upgradeModel() {
    c.model = "New Model" // This will change original model, as the struct is now passed by reference (pointer to original struct is passed)
}
```
