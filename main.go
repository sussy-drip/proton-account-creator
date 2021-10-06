package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/tebeka/selenium"
	"github.com/thanhpk/randstr"
)

const loginID = "#username"
const passwordID = "#password"
const passwordIDRepeat = "#repeat-password"
const letters = "aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ12345678910$#^*"

var email = flag.String("email", "", "recovery email to use for your acc")

func sleep(ms int) {
	time.Sleep(time.Millisecond * time.Duration(ms))
}

func nakedTextEntry(cmd string) {
	asArr := strings.Split(cmd, "")
	for _, v := range asArr {
		if v == " " {
			robotgo.KeyTap("space")
		} else if v == "@" {
			robotgo.KeyDown("shift")
			robotgo.KeyTap("2")
			robotgo.KeyUp("shift")
		} else {
			robotgo.KeyTap(v)
		}
	}
}

// func getProtonCode() string {
// 	svc := getMailService()
// 	call, err := svc.Users.Messages.List("me").Do()
// 	if err != nil {
// 		log.Println(err.Error())
// 	}

// 	if len(call.Messages) > 0 {
// 		log.Println("got", len(call.Messages), "messages")
// 		for _, c := range call.Messages {
// 			msg, err := svc.Users.Messages.Get("me", c.Id).Do()
// 			if err != nil {
// 				log.Println(err.Error())
// 			}
// 			if strings.Contains(strings.ToLower(msg.Snippet), "proton verification code") {
// 				arr := strings.Split(msg.Snippet, ": ")
// 				svc.Users.Messages.Delete("me", c.Id)
// 				return arr[1]
// 			}
// 		}
// 	}
// 	return ""
// }

func main() {
	flag.Parse()
	opts := []selenium.ServiceOption{
		selenium.GeckoDriver("./geckodriver.exe"),
		selenium.Output(os.Stderr),
	}
	selenium.SetDebug(false)
	_, err := selenium.NewSeleniumService("./selenium-server.jar", 8080, opts...)
	if err != nil {
		log.Println("failed to start service")
		panic(err)
	}
	//defer service.Stop()

	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, "http://localhost:8080/wd/hub")
	if err != nil {
		log.Println("failed to setup remote")
		panic(err.Error())
	}
	//defer wd.Quit()

	if err = wd.Get("https://account.protonmail.com/signup?language=en"); err != nil {
		log.Println("failed to navigate to login page")
		panic(err.Error())
	}

	sleep(5000)

	uname, pword := getUsernamePassword()
	elem, err := wd.FindElement(selenium.ByCSSSelector, loginID)
	if err != nil {
		log.Println("failed to find login element")
		panic(err.Error())
	}
	elem.Clear()
	elem.Click()

	nakedTextEntry(uname)

	elem, err = wd.FindElement(selenium.ByCSSSelector, passwordID)
	if err != nil {
		log.Println("failed to find password element")
		panic(err.Error())
	}

	elem.Clear()
	elem.Click()

	err = elem.SendKeys(pword)
	if err != nil {
		log.Println("failed to send text")
		panic(err.Error())
	}

	elem, err = wd.FindElement(selenium.ByCSSSelector, passwordIDRepeat)
	if err != nil {
		log.Println("failed to find password element")
		panic(err.Error())
	}

	elem.Clear()
	elem.Click()

	err = elem.SendKeys(pword)
	if err != nil {
		log.Println("failed to send text")
		panic(err.Error())
	}

	btns, _ := wd.FindElements(selenium.ByCSSSelector, "button[type=submit]")
	switch len(btns) {
	case 0:
		log.Println("found no buttons with selector :(")
	case 1:
		btns[0].Click()
		log.Println("found 1 btn!")
	default:
		log.Println("found many buttons :(")
	}
	sleep(2000)
	nakedTextEntry(*email)

	btns, _ = wd.FindElements(selenium.ByCSSSelector, "button[type=submit]")
	switch len(btns) {
	case 0:
		log.Println("found no buttons with selector :(")
	case 1:
		btns[0].Click()
		log.Println("found 1 btn!")
	default:
		log.Println("found many buttons :(")
	}
	btns, _ = wd.FindElements(selenium.ByCSSSelector, "button[type=button]")
	for _, btn := range btns {
		planDesc, _ := btn.GetAttribute("aria-describedby")
		for i := 0; i < 4; i++ {
			log.Println("--------------------------------------------------------------------------------")
		}
		log.Println(planDesc)
		for i := 0; i < 4; i++ {
			log.Println("--------------------------------------------------------------------------------")
		}
		if strings.ContainsAny(planDesc, "Free") {
			sleep(500)
			btn.Click()
			sleep(500)
			break
		}
	}
	// btn, _ := wd.FindElement(selenium.ByCSSSelector, "#label_1")
	// btn.Click()
	// sleep(500)
	// robotgo.KeyTap("enter")
	// sleep(10000)
	// code := ""
	// if code == "" {
	// 	code = getProtonCode()
	// 	sleep(1000)
	// }
	// nakedTextEntry(code)
	// sleep(500)
	// robotgo.KeyTap("enter")
	for i := 0; i < 4; i++ {
		log.Println("--------------------------------------------------------------------------------")
	}
	log.Println("username:", uname)
	log.Println("password:", pword)
	log.Println("please complete the captcha!")
	for i := 0; i < 4; i++ {
		log.Println("--------------------------------------------------------------------------------")
	}
	os.Exit(0)
}

func getUsernamePassword() (string, string) {
	letterArr := strings.Split(letters, ",")
	log.Println(letterArr)
	uname := randstr.String(18)
	pword := randstr.String(16, letters)
	log.Println("using", uname, "/", pword)
	return uname, pword
}
