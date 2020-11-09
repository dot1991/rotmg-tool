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
	choose int16 //mainScreen
	//General uses
	username string
	mail string
	password string

	//Less used
	newPassword string

	core_url = "https://realmofthemadgodhrd.appspot.com/"
	account_tag = "account/"
	char_tag = "char/"
	
	register = core_url + account_tag + "Register?isAgeVerified=1&newGUID="+mail+"&newPassword="+password+"&name="+username
	forgotPassword = core_url + account_tag + "forgotPassword?guid="+mail
	setName = core_url + account_tag + "setName?guid="+mail+"&password"+password+"&name"+username
	sendVerifyMail = core_url + account_tag + "sendVerifyEmail?guid="+mail+"&password="+password
	changePassword = core_url + account_tag + "changePassword?guid="+ mail + "&password="+password+"&newPassword="+newPassword
)

//This module is for global defines

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
func request(s string)(string,error){ //Same with "send" but you don't need extra code for data
	r,e := send(s)
	if e != nil {
		fmt.Println(e)
	}
	return r,e
}
func r()string{ //Line reader. You can use fmt.Scanf too.
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	return r.Text()
}
func c(s string)string{ //URL encoding. If you need more chars, you need to add them manually!
	c := s
	c = strings.Replace(s,"@","%40",99)
	s = strings.ReplaceAll(s,"?","%3F")
	//s = strings.ReplaceAll(s,"%","%25") Causing problem!
	s = strings.ReplaceAll(s,"+","%2B")
	s = strings.ReplaceAll(s,"#","%23")
	s = strings.ReplaceAll(s,".","%2E")
	s = strings.ReplaceAll(s,":","%3A")
	s = strings.ReplaceAll(s,"=","%3D")
	return c
}
//Script Usage
func Register(){
	fmt.Printf("Email: ")
	mail = r()
	fmt.Printf("Password: ")
	password = r()
	fmt.Printf("Name: ")
	username = r()
	r,e := send(register)
	if e != nil{
		fmt.Println(e)
	}
	fmt.Println(r)
}
func ForgotPassword(){
	fmt.Printf("Email: ")
	mail = r()
	r,e := send(forgotPassword)
	if e != nil{
		fmt.Println(e)
	}
	fmt.Println(r)
}
func SetName(){
	fmt.Printf("Email: ")
	mail = r()
	fmt.Printf("Password: ")
	password = r()
	fmt.Printf("Name: ")
	username = r()
	fmt.Println(request(setName))
}
func SendVerifyEmail(){
	fmt.Printf("Email: ")
	mail = r()
	fmt.Printf("Password: ")
	password = r()
	fmt.Println(request(sendVerifyMail))
}
func ChangePassword(){
	fmt.Printf("Email: ")
	mail = r()
	fmt.Printf("Password: ")
	password = r()
	fmt.Printf("New Password:")
	newPassword = r()
	fmt.Println(request(changePassword))
}