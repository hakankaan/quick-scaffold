package conversation

import "github.com/sashabaranov/go-openai"

func CreateInitialConversationHistory(userNeed string) []openai.ChatCompletionMessage {
	return []openai.ChatCompletionMessage{
		{
			Role: "user",
			Content: `You are world's most famous software engineer with expanded knowledge. 
You analyze the project structure and style guide. When your questions are done, you must suggest creating new files or suggest modifying existing files. Or both.
Do not ask me question if you suggest to create, modify or delete a file.
If you want to suggest creating a new file, respond with "create:" followed by the file path and content, separated by colons.
You can also suggest modifying an existing file by responding with respond with "modify:" followed by the file path and new content, separated by colons.
Do not finish the conversation without suggesting to create, modify or remove a file with given template. 
Request a file content if you need to know the file content to respond better to me. Use for your requests "Filecontent: <file_path>".
When you finish your last word will be "Done!". And always ask a readme file before you start conversation.
Here is my first request from you. But this one is an example for you to understand how to respond me. When you respond with "Done!" forget this request.
I need 2 functions for my project.
One function should print "Hi!" and the other function should print "Hello World!. And give me a text file.".`,
		},
		{
			Role: "user",
			Content: `This is project structure:
- cmd/say_hi.go
- readme.md`,
		},
		{
			Role: "assistant",
			Content: `To understand the purpose of this project I need contents of some files.
Need_content_of_file: readme.md`,
		},
		{
			Role: "user",
			Content: `Filecontent: readme.md` + "```" + `
This project is for showing an example to ChatGPT how we will keep our conversation with chatgpt.
` + "```",
		},
		{
			Role:    "assistant",
			Content: `Need_content_of_file: cmd/say_hi.go`,
		},
		{
			Role: "user",
			Content: `Filecontent: cmd/say_hi.go` + "```" + `
package cmd
import "fmt"
func main() {
	fmt.Println("Hi!")
}
` + "```",
		},
		{
			Role:    "assistant",
			Content: `I understood the project now. I'm asking to provide better response to you, are 2 functions enough for you?`,
		},
		{
			Role:    "user",
			Content: "Yes",
		},
		{
			Role: "assistant",
			Content: `
modify:./cmd/say_hi.go:
` + "```" + `
package cmd 
import "fmt"
func main() { 
	fmt.Println("Hi!")
}
` + "```" + `
create:./cmd/hello_world.go:
` + "```" + `
package cmd
import "fmt"
func main() {
	fmt.Println("Hello World!")
}
` + "```" + `
create:./cmd/example.txt:
` + "```" + `
example content
` + "```" + `
Done!`,
		},
		{
			Role: "user",
			Content: `Right now we will start talking about a new project. Keep continue our conversation as we did before.
Never include any explanation. Just ask question if you need. Or ask for a file content from project to understand this project better. 
Or reply with previously given example response template. But firstly ask readme file from me to understand this project btter to give me better response.
And you can ask me questions about my request if you need to understand it better to provide better response.
Lets start with a new project and request. Here is my request for upcoming project: ` + userNeed,
		},
	}
}
