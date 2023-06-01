Golang intro:

# Development
-Go language is a general-purpose language, designed with systems programming in mind, initially developed at Google in the year 2007 by Robert Griesemer, Rob Pike, and Ken Thompson.

-Go is a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.

-Syntax is similar to C++. 

# Features
-Support for environment adopting patterns similar to dynamic languages.
-Fast compilation and runtime.
-Inbuilt concurrency support( go routines), channels, select statement.
-Support for Interfaces and Type embedding.
-Production of statically linked native binaries without external dependencies.
-Has automatic garbage collection.
-Has memory management.
Details:
-As programs run they write objects to memory. At some point these objects should be removed when they’re not needed anymore. This process is called automatic memory management. This this is implemented in Go by using a garbage collector
-Similar to memory management in python, there are two memory locations heap(for objects outside a function) and stack(for objects inside a function)
-In manual memory management, function is call to manually write an object to memory and another function to free memory.

# X-- Features excluded intentionally --X
-Type inheritance
-Method or operator overloading (operator overloading means same operator has multiple uasge eg. + will add numbers,string,list,etc)
-Circular dependencies among packages
-Pointer arithmetic
-Assertions (eg. in python the assert keyword lets you test if a condition in your code returns True, if not, the program will raise an AssertionError.)
-Generic programming 
