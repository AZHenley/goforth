# goforth

An interpreter for a small dialect of Forth written in Go.

Fibonacci sequence:
```
: fib over over + ;  ( duplicates top two values and adds them )
```

Count down from 10:
```
10 @ lbl  ( start at 10 and define label )
dup . 10 emit  ( print current value and newline )
1 - dup if  ( subtract 1 and check if not zero )
lbl goto then  ( loop )
```

Duplicate the top 4 elements of the stack:
```
: 4dup cross cross cross dup 
back dup rot rot 
back dup cross rot rot 
back back dup cross swap cross rot rot 
back back ;
```

---

Based on the Forth programming language, this is a stack-based language with space-separated commands called _words_. Most words manipulate the stack in some way. Integers are pushed to the top of the stack. Running `1 2 +` will result in a stack with `3` on top.

This dialect only supports integer values and does not have variables. Instead it has a second stack that you can push and pop from (using `cross` and `back`). Additionally, it supports labels and goto rather than loops. For control flow, zero evaluates to false and any non-zero value evaluates to true. Referencing a label pushes an integer on the top of the stack where execution will continue. Words and labels must be unique. Comments beging with `(` and end with `)` with a space before and after each.

The table below describes all of the built-in words. The stacks are shown bottom to top (i.e., the right most element is the top).

| Word       | Stack effect          | Description |
| ------------- |:-------------:| -----:|
| dup      | ( 1 -- 1 1 ) | Duplicates the top element. |
| drop     | ( 1 -- )  | Pops the top element. |
| swap | ( 1 2 -- 2 1 )    | Swaps the top two elements. |
| over      | ( 1 2 -- 1 2 1 ) | Duplicates the second element and puts it on top. |
| rot     | ( 1 2 3 -- 2 3 1 )  | Rotates the top three elements. |
| cross      | ( 1 , -- , 1 ) | Pops the top element and pushes it to the secondary stack. |
| back     | ( , 1 -- 1 , )  | Pops the top element of the second stack and pushes it to stack. |
| + | ( 1 2 -- 3 )  | Pops the top two elements and pushes the sum. |
| - | ( 1 2 -- -1 )  | Pops the top two elements and pushes the difference. |
| * | ( 1 2 -- 2 )  | Pops the top two elements and pushes the product. |
| / | (1 2 -- 0 )  | Pops the top two elements and pushes the quotient. |
| mod | ( 1 2 -- 1 ) | Pops the top two elements and pushes the remainder. |
| >     | ( 1 2 -- 0 )  | Compares if the second element is greater than the top element. |
| < | ( 1 2 -- 1 ) | Compares if the second element is less than the top element. |
| if | ( 1 -- ) | If the top is non-zero, continue execution at the next word. |
| else | no effect | Optional after an `If`, continue executing if the `If` was false.  |
| then | no effect | Ends an if/if-else block. |
| : | no effect | Start of a word definition. The next token is the name followed by the body. |
| ; | no effect | End of a word definition. |
| @ | no effect | Label definition. The next token is the name. |
| . | ( 1 -- ) | Prints the top. |
| emit | ( 1 -- ) | Prints the top as an ASCII character. |
| key | ( -- 1 ) | Gets keyboard input and pushes the ASCII value. |

---

## What is the point?

Just exploring how Forth works. It isn't an optimal implementation in any way.
