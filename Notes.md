### Introduction

- Every Go Code needs to have a main function. 
- main.go and package main are conventions. We can choose whatever we want.
- first line should always have package definition. 
- 

### Basics 
#### Variables
- We are not allowed to declare variables without storing something in them in go. 
- We are not allowed to declare variables without using them. 
- Normal int is int64 if 64bit computer else it is 32bit. 

#### Pointers
- & to get reference of variable. 
- * to pass a pointer. 


#### Types and Struct 
- Go initializes default values. 
- Capital functions and variables are public 
- LowerCase function and variables are private. 
- Structs with functions are called as receivers.


#### Datastructures 
- We use make keyword to create a map. 
- Maps are immutable. 
- Maps are not ordered. 
- We can use interface{} to store different types if it is not known to us.
- We don't use arrays in go. We mostly use slices. 
- In Go 
    - A string is slice of runes. 
    - A rune is actually a byte.
    - A string is immutable. So when reassigning the existing object is destroyed.  

#### Decision making statements 
- if else
- switch case default

#### Looping 
- Only for loop is there in go. 

#### Interfaces 
- All structs to implement the function in interface. 

