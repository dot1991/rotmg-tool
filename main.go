package main

import(
	"fmt"
)

//Main is here
func main(){
	mainScreen()
}

func mainScreen(){
	fmt.Printf("************\n")
	fmt.Printf("*RotMG Tool*\n")
	fmt.Printf("************\n")
	fmt.Printf("1)Register\n")
	fmt.Printf("2)Forgot Password\n")
	fmt.Printf("3)Set Name\n")
	fmt.Printf("4)Send Verify Email\n")
	fmt.Printf("5)Change Password\n")
	fmt.Printf("6)Ban Checker\n")
	fmt.Printf("$")
	choose = r()
	if choose == "1" {
		Register()
	}
	if choose == "2" {
		ForgotPassword()
	}
	if choose == "3" {
		SetName()
	}
	if choose == "4" {
		SendVerifyEmail()
	}
	if choose == "5" {
		ChangePassword()
	}
	if choose == "6" {
		isBanned()
	}
}
