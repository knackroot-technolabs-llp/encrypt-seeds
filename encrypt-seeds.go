package main;

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
	"gitlab.com/david_mbuvi/go_asterisks"
)

func encrypt(plaintext string, key string) string {
	ciphertext := ""
	encryptKey, _ := strconv.Atoi(key)
	encryptKey = encryptKey % 26
	for _, char := range plaintext {
		if char >= 'A' && char <= 'Z' {
			ciphertext += string((int(char-'A')+encryptKey)%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			ciphertext += string((int(char-'a')+encryptKey)%26 + 'a')
		} else {
			ciphertext += string(char)
		}
	}
	if decrypt(ciphertext, key) == plaintext {
		return ciphertext
	}
	return "";
}

func decrypt(ciphertext string, key string) string {
	plaintext := ""
	decryptKey, _ := strconv.Atoi(key)
	decryptKey = decryptKey % 26
	for _, char := range ciphertext {
		if char >= 'A' && char <= 'Z' {
			plaintext += string((int(char-'A')-decryptKey+26)%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			plaintext += string((int(char-'a')-decryptKey+26)%26 + 'a')
		} else {
			plaintext += string(char)
		}
	}
	return plaintext
}

func returnErrorMsg(msg string) {
	fmt.Println(msg)
	os.Exit(0)
}

func main() {
	app := &cli.App{
		Name:  "Encrypt your seed phrase",
		Usage: "Encrypt seed phrase of 12 or 24 words",
		Commands: []*cli.Command{
			{
				Name:  "encrypt",
				Usage: "Encrypt a seed-phrase",
				Action: func(c *cli.Context) error {
					scanner := bufio.NewScanner(os.Stdin)
					fmt.Print("How many words are there in your seed phrase - 12 or 24 : ")
					scanner.Scan()
					words, err := strconv.Atoi(scanner.Text())
					if err != nil {
						returnErrorMsg("Please enter 12 or 24 in numbers")
					}
					if words == 12 || words == 24 {

						encryptedWords := make([]string, words)
						plainWords := make([]string, words)
						fmt.Print("Enter your passcode of 4 to 8 digits : ")
						passcode, _ := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
						_, err := strconv.Atoi(string(passcode))
						if err != nil || !(len(passcode) >= 4 && len(passcode) <= 8) {
							returnErrorMsg("Passcode should be of 4 to 8 digits and should be numbers")
						}
						fmt.Print("Confirm your passcode : ")
						confirmPasscode, err := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
						
						if string(passcode) != string(confirmPasscode) {
							returnErrorMsg("Passwords are not same. Please enter same password.")
						}

						fmt.Println("Remember your passcode to decrypt your seed phrase..")

						fmt.Println("Enter words one by one: ")

						for i := 0; i < words; i++ {
							fmt.Print(i+1, ": ")
							scanner.Scan()
							word := scanner.Text()
							if len(word) != 0 {
								encrypted := encrypt(word, string(passcode))
								encryptedWords[i] = encrypted
								plainWords[i] = word
							} else {
								returnErrorMsg("Seed word cannot be empty.")
							}
						}

						encryptedPhrase := strings.Join(encryptedWords, " ")

						plaintext := strings.Join(plainWords, " ")

						fmt.Printf("\n\n#####################################################################\n")
						fmt.Println("Please store below encrypted seed phrase somewhere....")
						fmt.Printf("#####################################################################\n\n")

						fmt.Println("Your original seed phrase is : ")
						fmt.Printf("%s\n\n", plaintext)

						fmt.Println("Your encrypted seed phrase is : ")
						fmt.Printf("%s\n", encryptedPhrase)
						return nil
					} else {
						returnErrorMsg("Please select 12 or 24 only")
						return nil
					}
				},
			},
			{
				Name:  "decrypt",
				Usage: "Decrypt a seed-phrase",
				Action: func(c *cli.Context) error {
					scanner := bufio.NewScanner(os.Stdin)
					fmt.Print("How many words are there in your seed phrase - 12 or 24 : ")
					scanner.Scan()
					words, err := strconv.Atoi(scanner.Text())
					if err != nil {
						returnErrorMsg("It should only be 12 or 24")
					}
					if words == 12 || words == 24 {

						decryptedWords := make([]string, words)
						fmt.Print("Enter your passcode of 4 to 8 digits : ")
						passcode, _ := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
						_, er := strconv.Atoi(string(passcode))
						if er != nil || !(len(passcode) >= 4 && len(passcode) <= 8) {
							returnErrorMsg("Passcode should be of 4 to 8 digits and should be numbers")
						}
						fmt.Println("Enter words one by one : ")
						for i := 0; i < words; i++ {
							fmt.Print(i+1, ": ")
							scanner.Scan()
							word := scanner.Text()
							if len(word) != 0 {
								decryptedWords[i] = decrypt(word, string(passcode))
							} else {
								returnErrorMsg("Seed word cannot be empty.")
							}
						}

						decryptedPhrase := strings.Join(decryptedWords, " ")

						fmt.Printf("\nDecrypted Phrase: %s\n", decryptedPhrase)
						return nil
					} else {
						returnErrorMsg("Please select 12 or 24 only")
						return nil
					}
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
