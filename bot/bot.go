package bot

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"../config"
	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}
	goBot.AddHandler(messageHandler)
	BotID = u.ID

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
	fmt.Println("!ascii / !ascii2 / !ascii3")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.ID == BotID {

			return
		}

		split := strings.Split(m.Content, " ")

		for _, str := range split {
			fmt.Println("arg: ", str)
		}

		fmt.Println("commande: ", split[1:])
		if len(split) > 1 {
			if split[0] == "!ascii" {
				myString := strings.Join(split[1:], " ")
				//_, _ = s.ChannelMessageSend(m.ChannelID, "```"+result+"```")
				file, _ := os.Open("standard.txt") // indique quelle fichier regarder
				fileVal := ScanFile(file)
				//for _, v := range myString {
				_, _ = s.ChannelMessageSend(m.ChannelID, "```"+printLetter(string(myString), fileVal)+"```") //fait jusqu'a la fin de l'arg renseigner
				//	}

			}
			if split[0] == "!ascii1" {
				myString := strings.Join(split[1:], " ")
				//_, _ = s.ChannelMessageSend(m.ChannelID, "```"+result+"```")
				file, _ := os.Open("standard.txt") // indique quelle fichier regarder
				fileVal := ScanFile(file)
				//for _, v := range myString {
				_, _ = s.ChannelMessageSend(m.ChannelID, "```"+printLetter(string(myString), fileVal)+"```") //fait jusqu'a la fin de l'arg renseigner
				//	}

			}
			if split[0] == "!ascii2" {
				myString := strings.Join(split[1:], " ")
				//_, _ = s.ChannelMessageSend(m.ChannelID, "```"+result+"```")
				file, _ := os.Open("thinkertoy.txt") // indique quelle fichier regarder
				fileVal := ScanFile(file)
				//for _, v := range myString {
				_, _ = s.ChannelMessageSend(m.ChannelID, "```"+printLetter(string(myString), fileVal)+"```") //fait jusqu'a la fin de l'arg renseigner

				//}
			}

			if split[0] == "!ascii3" {

				myString := strings.Join(split[1:], " ")
				//_, _ = s.ChannelMessageSend(m.ChannelID, "```"+result+"```")
				file, _ := os.Open("shadow.txt") // indique quelle fichier regarder
				fileVal := ScanFile(file)
				//for _, v := range myString {
				_, _ = s.ChannelMessageSend(m.ChannelID, "```"+printLetter(string(myString), fileVal)+"```") //fait jusqu'a la fin de l'arg renseigner

				//}
			}
		} else {
			s.ChannelMessageSend(m.ChannelID, "il n'y a pas de message")
		}

	}
}
func printLetter(s string, fileVal []string) string { // determine la ligne et ecrit la ligne
	k := ""

	for i := 1; i <= 8; i++ { // permet d'ecrire la ligne 1, puis 2 ...
		for _, arg := range s {
			indexBase := int(rune(arg)-32) * 9 // trouve quelle est la ligne
			//	fmt.Println(indexBase + i)
			k += fileVal[indexBase+i] // ecrit la ligne voulue

		}
		k += "\n" // permet un simple retour a la ligne

	}
	return k
}
func ScanFile(arg *os.File) []string {
	var fileValue []string
	scanner := bufio.NewScanner(arg)
	for scanner.Scan() {
		lines := scanner.Text()
		fileValue = append(fileValue, lines)
	}
	return fileValue
}
