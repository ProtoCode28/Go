# TUTORIAL TO GO ;)

## WHAT'S GO?

“Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.” – https://golang.org [1].​

![](pictures/Anmerkung%202019-10-31%20153007.png)

## WHERE DOES GO COME FROM?

* Developed by google in 2007​

* Google needed a programming language:​
    * As fast as C​

    * As easy as Python​

    * Inheritance like in Java​

* Google servers were a mix of those three languages​

* Go is an offspring from C​

## WHY GO?​

* has a fast compiler ​

* huge standard library ​

* has many built-in tools (for testing, code coverage, etc.) ​

* supports many platforms: ​

    * Linux 2.6.23 or later ​

    * macOS 10.8 or later​

    * Windows XP SP2 or later ​

    * FreeBSD 9.3 or later ​
    
* Works on arm and x86 ​

* No interpreter is required for execution​

* Used by huge IT companies like:​

    * Google​

    * Twitter​

    * Uber​

    * And many more​
    
## HOW TO WRITE GO CODE​

* Official guidelines are:​

    * Keep all the code in a single workspace​

    * A workspace contains many version control repositories​

    * Each repository contains one or more packages​

    * A package has at least one Go source file per directory​

    * A package directory path determines its import path​

* Different from other programming languages since their projects usually have separate work spaces and those are closely tied to version control repositories​

## INTRODUCTION INTO CODING WITH GO​

![](pictures/Bild1.png)

Go programs are always made out of packages. Most of the packages are in the standard library. They´re basically code files that contain functions. With the import command you can import a package. In the example above, we import the "fmt" package which contains the Println function. 

## How to define Variables

![](pictures/Bild2.png)

If you define a variable outside of a function you need to declare it with *var*.
After that you name the variable. In our example we named the first variable *a*. The data type comes at the end of the declaration. In our example it's an integer. 
You can also define multiple variables of the same datatype and add the datatype at the end of the declaration. Look at the example in line 6. 

## How to define Variables PT.2

![](pictures/Bild3.png)

If you declare two variables in one declaration, you can also give them two values. You have to separate those by a comma like in line 5. 
Another possibility is to not declare the datatype. Go will assign a datatype automatically, based on the values, that were given to the variable. Check out line 6 of the example. 
If you define a variable within the function, you can just create a name and assign a value like in line 9. 

## For loop

![](pictures/Bild4.png)

The for loop is pretty similar to the for loop in c or c++. The only difference is, you can omit the (), but the {} are mandatory. 

## For loop PT.2

![](pictures/Bild5.png)

The initial statement and the post statement can be omitted like in line 7 in the example. 

## While loop

![](pictures/Bild6.png)

Some of you may have noticed it already in for loop PT.2. The for loop is also the while loop in go. You just write your while condition after the for statement. Make sure it terminates! 

## IF condition

![](pictures/Bild7.png)

The if statement is similar to if statements in C or C++, too. And once again the () are omitted but the {} are mandatory. 

## If condition pt.2

![](pictures/Bild8.png)

A cool thing you can do in go, is add a command within the if statement. It will be executed before the if condition is checked. Look at line 11 in the example. 

## Defer

![](pictures/Bild9.png)

Another cool feature is defer. Defer stalls function calls and puts them on the stack. The execution will be the very last one. If you defer more than one function the LIFO(last-in-first-out) principle will be used. Look at the output of our example and you will understand. 

![](pictures/Bild10.png)

## Switch case

![](pictures/Bild11.png)

The switch case function is pretty simple and similar to other programming languages. I decided to show it to you anyways because we will need it later for the Hands-On. 

## Arrays

![](pictures/Bild12.png)

The size of arrays can't be changed. You don't have to go with a for loop through the array to output it though. If you put the name of the array in the Println function, the entire array will be outputted. 

## Slices

![](pictures/Bild13.png)

While array have a fix size, slices are dynamic. The access of the elements in it is flexible. Coding with go, slices are used way more often than arrays. Slices are like references pointing on arrays. If you change an element in your slice, you change the array that's behind it. Check it out in this example. 

![](pictures/Bild14.png)

![](pictures/Bild15.png)

In the following example, the x is an array and the q is a slices. 
Do you see the difference? 

![](pictures/Bild16.png)

## Hands-On

In my Hands-On I'm coding the magic conch shell. 








