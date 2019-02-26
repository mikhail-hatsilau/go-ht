# Welcom to Go!

 ## Common requirements

First of all you have to fork this repo. You can find «fork» button in the right corner.
When you press the button, you will have your own copy of the repository.

## Task requirements

Task can be reviewed and considered completed only when unit test are passed

## Process

You have to clone forked repo and create 3 branches for each task. You commit changes for a specific task to an appropriate branch.
There are 3 tasks and 3 separate folders for each of them. There are two files in each folder: main.go and main_test.go. main_test.go file contains unit tests for the task. So, you have to follow requirements below and launch unit tests to check you solution (how to do it will be written a bit later :) ). When you are ready and ALL related test cases are passed, you can push your branch and create a pull request which should be assigned on me (mikhail-hatsilau). I will check it and leave some comments if necessary. Don’t merge your branches without my approval!

P.S. You can create your packages without any problem. Only remember that required functions have to be defined in the main package!

## How to start unit tests

Go to the folder of concrete task and start command: go test -v.
When tests are fallen you can find a message in the console, which says the reason.
If all test cases are passed for the concrete task you are ready to commit it and create pull request.
Don’t touch units tests! Your code has two pass them. You don’t need to change anything ion *_test.go files.
If you see an issue and sure that tests are not correct, just contact to me and we will discuss it.

## Tasks description

### Task1

In the task 1 you have to create «Filter» function. It should accept int slice and predicate function. Predicate function should decide if item has to be included in the result and it returns bool. If true is returned from predicated function for a specific item, then item should be included in the result. Example:
`fmt.Println("Even", Filter([]int{1, 2, 3, 4, 5}, func(item, index int) bool { return item % 2 == 0 }))`
Result in the console: `[2, 4]`
Unit tests are included. Just use `go test -v` command inside the folder

### Task 2

In the task 2 you have to create «MapTo» function. It should accept int slice and a function, which converts item to the one. MapTo function should convert items in a slice according to the result of function passed as the second parameter.
Other words, function (passed as the second argument) should be called on each item and return new item and place in a new slice.
Example:
`fmt.Println(MapTo([]int{1, 2, 3, 4, 5, 10}, func(item, index int) int { return item * 2 }))`
Above example returns `[2, 4, 6, 8, 10, 20]`
Also, you have to write Convert function, with uses MapTo function and converts numbers to words:
Input: `[1, 2, 3, 4, 5]`
Outout: `[«one», «two», «three», «four», «five»]`
Only numbers from 1 to 9 should be covered. If any other number is passed, «unknown» string should be placed instead.
Example:
Input: `[0, 1, 2, 3, 10]`
Output: `[«unknown», «one», «two», «three», «unknown»]`
Hint: you can use go maps for matching number to string
Unit tests are included. Just use `go test -v` command inside the folder

### Task 3

In the task 3 you have to write console application for managing users list. It should accept there types of operation:
`add
list
findById
remove`
Users list should be stored in the json file. When you start your application and tries to perform some operations, existing file should be used or new one should be created if it does not exist.
Example of the json file (users.json):
`[{id: "1", email: «test@test.com», age: 31}, {id: "2", email: «test2@test.com», age: 41}]`
In the main.go file you can find a function called Perform(args Arguments, writer io.Writer) error.
You have to call this function from the main function and pass arguments from the console and os.Stdout stream. Perform function body you have to write by yourself :).
Arguments - is a `map[string]string` with the following fields:
`id, item, operation and fileName
Arguments should be passed via console flags:
`./main.go -operation «add» -item ‘{«id»: "1", «email»: «email@test.com», «age»: 23}’ -fileName «users.json»`
`-operation`, `-item`and `-fileName` are console flags. To parse them and build structure you can take a Look at «flag» package: https://golang.org/pkg/flag/.
Pay attention that `-fileName` flag should be provided every time with the name of file where you store users!

#### Getting list of items:
Application has to retrieve list from the users.json file and print it to the `io.Writer` stream. Use writer from the argument of Perform function to print the result! It is important for passing unit tests. It can be smth like `writer.Write(bytes)`
File content: `[{«id»: "1", «email»: «email@test.com», «age»: 23}]`
Command: `./main -operation «list» -fileName «users.json»` (main is bult go application. Binary file after go build command)
Output to the console: `[{«id»: "1", «email»: «email@test.com», «age»: 23}]`
If file is empty then nothing should be printed to the console.
**Errors:** 
1. If `-operation` flag is missing, then error `-operation` flag has to be specified» has to be returned from Perform function. Package `errors` can be used for creating errors (https://golang.org/pkg/errors/).
2. If `—fileName` flag is missing, then error «-fileName flag has to be specified» has to be returned from Perform function.
3. If operation can not be handled, for example «abc» operation, then «Operation abcd not allowed!» error has to be return from the Perform function
All cases are covered by unit tests. If you want to be sure your solution works correct, just start `go test -v` command in the task3 folder
#### Adding new item:
For adding new item to the array inside users.json file, application should provide the following cmd command:
`./main -operation «add» -item «{«id»: "1", «email»: «email@test.com», «age»: 23}» -fileName «users.json»`
`-item` - valid json object with the id, email and age fields
**Errors:**
1. All errors about operation and fileName flags mentioned above
2. If `-item` flag is not provided Perform function should return error «-item flag has to be specified»

#### Remove user
Application should allow to remove user with the following command:
`./main -operation «remove» -id "2" -fileName «users.json»`
If user with id `"2"`, for example, does not exist, Perform functions should print message to the `io.Writer` «Item with id 2 not found».
If user with specified id exists, it should be removed from the users.json file.
**Errors:**
1. All errors about operation and fileName flags mentioned above
2. If `-id` flag is not provided error «-id flag has to be specified» should be returned from Perform function

#### Find by id
Application should allow to find user by id with the following command:
`./main -operation «findById» -id "1" -fileName «users.json»`
If user with specified id does not exists in the users.json file, then empty string has to be written to  the `io.Writer`
If user exists, then json object should be written in `io.Writer`
**Errors:**
1. All errors about operation and fileName flags mentioned above
2. If `-id` flag is not provided error «-id flag has to be specified» should be returned from Perform function
All cases of the task 3 are covered by unit tests, So, you can check your solution during thee implementation. If you have any questions, just write in the slack.

### Useful info:
1. For opening and creating file use `os` package and `OpenFile` function https://golang.org/pkg/os/
2. To simply read file use `ioutil` package and `ReadAll` function https://golang.org/pkg/io/ioutil/
3. To convert json string to the object use `encoding/json` package and `Unmarshal` function: https://golang.org/pkg/encoding/json/
4. To convert json array or object to string use json package and `Marshal` function.
5. Go does not have throw operator and try catch statement. Instead it has multi return and allows to return error from a function: `func () (User, error) {}`
Take a look: https://medium.com/@hussachai/error-handling-in-go-a-quick-opinionated-guide-9199dd7c7f76
6. If you receive error in Perform function, just call panic function for exiting the execution and printing error

Note that flags and operations names should be the same as mentioned above or unit tests will never pass.
