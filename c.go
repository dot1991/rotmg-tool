package main
import(
	"fmt"
	"net/http"
	"bufio"
	"os"
	"io/ioutil"
	"strings"
)
var(
	choose string //mainScreen
	isInfo = false
	//General uses
	username string
	email string
	password string

	//Less used
	newPassword string

	core_url = "https://realmofthemadgodhrd.appspot.com/"
	account_tag = "account/"
	char_tag = "char/"
	
	register string
	forgotPassword string 
	setName string
	sendVerifyMail string
	changePassword string
	banCheck string
)

//This module is for global defines

//Still i don't know why i need func. Not working if i don't use :/
func url_create(){
	register = core_url + account_tag + "register?isAgeVerified=1&newGUID="+email+"&newPassword="+password+"&name="+username
	forgotPassword = core_url + account_tag + "forgotPassword?guid="+email
	setName = core_url + account_tag + "setName?guid="+email+"&password"+password+"&name"+username
	sendVerifyMail = core_url + account_tag + "sendVerifyEmail?guid=" + email + "&password=" + password
	changePassword = core_url + account_tag + "changePassword?guid="+ email + "&password="+password+"&newPassword="+newPassword
	banCheck = core_url + char_tag + "list?guid="+email+"&password="+password
}

//Gloabal usage funcs

func send(s string)(string,error){ //Send request
	d,e := http.Get(s)
	if e != nil{
		fmt.Println(e)
	}
	defer d.Body.Close()
	b,e := ioutil.ReadAll(d.Body)
	c := string(b)
	if b != nil{
		//fmt.Println(c)
	}
	if e != nil{
		fmt.Println(e)
	}
	return c,e
}
func request(s string)(string){ //Same with "send" but you don't need extra code for data
	r,e := send(s)
	if e != nil {
		fmt.Println(e)
	}
	return r
}
func r()string{ //Line reader. You can use fmt.Scanf too.
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	return r.Text()
}
func c(s string)string{ //URL encoding. If you need more chars, you need to add them manually!
	s = strings.Replace(s,"%","%25",-1)
	s = strings.Replace(s,"@","%40",-1)
	s = strings.Replace(s,"?","%3F",-1)
	s = strings.Replace(s,"+","%2B",-1)
	s = strings.Replace(s,"#","%23",-1)
	s = strings.Replace(s,".","%2E",-1)
	s = strings.Replace(s,":","%3A",-1)
	s = strings.Replace(s,"=","%3D",-1)
	return s
}
//Script Usage
func Register(){
	fmt.Printf("Email: ")
	email = c(r())
	fmt.Printf("Password: ")
	password = c(r())
	fmt.Printf("Name: ")
	username = r()
	fmt.Println(request(register))	
}
func ForgotPassword(){
	fmt.Printf("Email: ")
	email = c(r())
	r,e := send(forgotPassword)
	if e != nil{
		fmt.Println(e)
	}
	url_create()
	fmt.Println(r)
}
func SetName(){
	fmt.Printf("Email: ")
	email = c(r())
	fmt.Printf("Password: ")
	password = c(r())
	fmt.Printf("Name: ")
	username = r()
	url_create()
	fmt.Println(request(setName))
}
func SendVerifyEmail(){
	fmt.Printf("Email: ")
	email = c(r())
	fmt.Printf("Password: ")
	password = c(r())
	url_create()
	fmt.Println(request(sendVerifyMail))
}
func ChangePassword(){
	fmt.Printf("Email: ")
	email = c(r())
	fmt.Printf("Password: ")
	password = c(r())
	fmt.Printf("New Password:")
	newPassword = c(r())
	url_create()
	fmt.Println(request(changePassword))
}
func isBanned() {
	fmt.Printf("Email: ")
	email = c(r())
	fmt.Printf("Password: ")
	password = c(r())
	url_create()
	x := request(banCheck)
	if strings.Contains(x,"Account is under maintenance"){
		fmt.Printf("Account is banned!")
	}else{
		fmt.Printf("Account is not banned!")
	}
}