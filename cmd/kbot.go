/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var (
	// TeleToken bot
	TeleToken = os.Getenv("TELE_TOKEN")
)

// kbotCmd represents the kbot command
var kbotCmd = &cobra.Command{
	Use:   "kbot",
	Aliases: []string{"start"},
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("kbot %s started", appVersion)
		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("Please check TELE_TOKEN env variable. %s", err)
			return
		}

		kbot.Handle(telebot.OnText, func(m telebot.Context) error {
			log.Print(m.Message().Payload, m.Text())

			return m.Send(fmt.Sprintf("Hi "+ m.Sender().FirstName +", I'm Kbot %s! I will help you to generate a strong password. Just type: /generate", appVersion))
		})

		kbot.Handle("/generate", func(m telebot.Context) error {
			payload := m.Message().Payload
			password := ""

			if payload == "" {
				// log.Print(generatePassword(10))
				return m.Send(generatePassword(10));
			} else {
				length, err := strconv.Atoi(payload)

				if err != nil{
					//executes if there is any error
					log.Println(err)
					password = "Invalid password length"
					// return m.Send(err);
				} else {
					//executes if there is NO error
					if length != 0 {
						password = generatePassword(length)
					} else {
						password = generatePassword(10)
					}
				}
			}

			return m.Send(password);
		})

		kbot.Handle("/help", func(m telebot.Context) error {
			return m.Send("Use /generate N where N - is a password length")
		})

		kbot.Start()
	},
}

const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&*()"

func generatePassword(length int) string {
	src := rand.NewSource(time.Now().UnixNano())
  	rand := rand.New(src)
	var password []byte
	
	for i := 0; i < length; i++ {
	  password = append(password, chars[rand.Intn(len(chars))]) 
	}
	
	return string(password)
  }

func init() {
	rootCmd.AddCommand(kbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// go build -ldflags "-X="github.com/autonibit/kbot/cmd.appVersion=v1.0.2