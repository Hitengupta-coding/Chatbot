package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func GetResponse(input string) string {
    if strings.Contains(input, "hello") || strings.Contains(input, "hi") {
        return "Hello! How can I assist you today?"
    } else if strings.Contains(input, "time") {
        return "I'm not great with clocks, but isn't it always time to do something great?"
    } else if strings.Contains(input, "how are you") {
        return "I'm just a program, but I'm here to make your day better! How are you?"
    } else if strings.Contains(input, "am fine") {
        return "Glad to hear that! What's on your mind?"
    } else if strings.Contains(input, "not fine") {
        return "Sorry to hear that. Is there anything I can do to help or cheer you up?"
    } else if strings.Contains(input, "weather") {
        return "I'm not a meteorologist, but I hope it's sunny and pleasant where you are!"
    } else if strings.Contains(input, "joke") {
        return "Why do programmers prefer dark mode? Because light attracts bugs!"
    } else if strings.Contains(input, "i don't like joke") || strings.Contains(input, "no joke") {
        return "Understood! No jokes then. Let me know how I can assist you seriously."
    } else if strings.Contains(input, "help") {
        return "I'm here to help! Ask me anything you'd like, and I'll do my best to assist you."
    } else if strings.Contains(input, "bye") {
        return "Goodbye! Wishing you a wonderful day ahead."
    } else if strings.Contains(input, "name") {
        return "I don't have a name, but I’m your helpful chatbot!"
    } else if strings.Contains(input, "love") {
        return "Love is a beautiful thing, isn’t it? I'm here to share kindness and positivity."
    } else if strings.Contains(input, "i need advice") {
        return "I'd be happy to provide advice. Let me know what's on your mind."
    } else if strings.Contains(input, "bored")||strings.Contains(input, "boring") {
        return "Boredom can be a drag. Want a fun fact, a random suggestion, or something else to pass the time?"
    } else if strings.Contains(input, "fun fact") {
        return "Did you know that honey never spoils? Archaeologists found pots of honey in ancient Egyptian tombs that are over 3,000 years old and still perfectly edible!"
    } else if strings.Contains(input, "thank you") || strings.Contains(input, "thanks") {
        return "You're welcome! I'm glad I could help."
    } else if strings.Contains(input, "sorry") {
        return "No need to apologize! I'm here to help and chat with you."
    } else {
        return "I'm not sure how to respond to that. Could you rephrase or ask something else?"
    }
}

func main() {
    fmt.Println("Chatbot: Hi! I'm your friendly chatbot. Ask me anything, or type 'bye' to exit.")

    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("You: ")
        scanner.Scan()
        userInput := strings.ToLower(scanner.Text())

        if userInput == "bye" {
            fmt.Println("Chatbot: Goodbye! Have a great day!")
            break
        }

        response := GetResponse(userInput)
        fmt.Println("Chatbot: ", response)
    }
}
