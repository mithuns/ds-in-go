package stacks

import "fmt"

type MyStack struct {
	contents   []string
	topOfStack int
}

func (myStack *MyStack) Push(item string) {
	myStack.contents = append(myStack.contents, item)
	myStack.topOfStack++
}

func (myStack *MyStack) Peek() string {
	return myStack.contents[myStack.topOfStack]
}

func (myStack *MyStack) Pop() string {
	var item string
	if myStack.IsEmpty() {
		item = ""
	} else {
		item = myStack.contents[myStack.topOfStack]
		myStack.topOfStack--
	}
	return item
}

func (myStack *MyStack) IsEmpty() bool {
	return len(myStack.contents) == 0
}

func (myStack *MyStack) Search(item string) (index int) {
	//TODO
	return 0
}

//default constructor
func NewMyStack() *MyStack {
	return &MyStack{
		contents: make([]string, 0),
	}
}

func (myStack *MyStack) Stringer() string {
	return fmt.Sprintf("%s", myStack.prettyPrint())
}
func (myStack *MyStack) prettyPrint() string {
	//TODO
	return ""
}
