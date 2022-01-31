package main

import "fmt"

func main() {

	user := new(UserServiceImpl)
	number := 5
	user.play(&number)
}


type UserService interface{
	play()
}

type UserServiceImpl struct{
	UserService
}

func (userServiceImpl *UserServiceImpl) play(number *int){
	fmt.Print("Play", number)
}