// The application should run from the command line, accept the GitHub username as an argument, fetch the user’s recent activity using the GitHub API, and display it in the terminal. The user should be able to:

// Provide the GitHub username as an argument when running the CLI.
// github-activity <username>
// Fetch the recent activity of the specified GitHub user using the GitHub API. You can use the following endpoint to fetch the user’s activity:
// # https://api.github.com/users/<username>/events
// # Example: https://api.github.com/users/kamranahmedse/events
// Display the fetched activity in the terminal.
// Output:
// - Pushed 3 commits to kamranahmedse/developer-roadmap
// - Opened a new issue in kamranahmedse/developer-roadmap
// - Starred kamranahmedse/developer-roadmap
// - ...

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

type Response struct {
	ID        string                 `json:"id"`
	Type      string                 `json:"type"`
	Actor     Actor                  `json:"actor"`
	Repo      Repo                   `json:"repo"`
	Payload   map[string]interface{} `json:"payload"`
	Public    bool                   `json:"public"`
	CreatedAt time.Time              `json:"created_at"`
}

type Actor struct {
	ID           int    `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarID   string `json:"gravatar_id"`
	URL          string `json:"url"`
	AvatarURL    string `json:"avatar_url"`
}

type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func readLine(prompt string) string {
	fmt.Println(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
func main() {
	// getting user's github usernamess
	prompt := "Enter User's Github username"
	userName := readLine(prompt)

	//formating the url based on the users github username
	url := fmt.Sprintf("https://api.github.com/users/%s/events", userName)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		fmt.Println("User not found")
		return
	}

	// closing the body using defer to close at the end

	//unmarshing the data comming into the
	decoder := json.NewDecoder(resp.Body)
	var responses []Response
	if err := decoder.Decode(&responses); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	if len(responses) == 0 {
		fmt.Println("No recent activity found for user:", userName)
		return
	}
	for i, event := range responses {
		fmt.Println("--------------------------------------------------")
		fmt.Printf("Event %d:\n", i+1)
		fmt.Printf("Event Type: %s\n", event.Type)
		fmt.Println("Event Details:")
		switch event.Type {
		case "PushEvent":
			commitsRaw, ok := event.Payload["commits"].([]interface{})
			if !ok {
				fmt.Println("No commits found or invalid format")
				continue
			}

			for _, c := range commitsRaw {
				commit, ok := c.(map[string]interface{})
				if !ok {
					continue
				}
				sha, _ := commit["sha"].(string)
				message, _ := commit["message"].(string)
				fmt.Printf(" - Commit: %s | Message: %s\n", sha, message)
			}

		case "CreateEvent":
			refType, _ := event.Payload["ref_type"].(string)
			ref, _ := event.Payload["ref"].(string)
			fmt.Printf(" - Created %s: %s\n", refType, ref)

		case "MemberEvent":
			action, _ := event.Payload["action"].(string)
			member, _ := event.Payload["member"].(map[string]interface{})
			memberLogin, _ := member["login"].(string)
			fmt.Printf(" - %s member: %s\n", action, memberLogin)

		default:
			fmt.Println(" - Event details:", event.Payload)
		}

		fmt.Println("Repository Name:", event.Repo.Name)
		fmt.Println("Created At:", event.CreatedAt.Format("2006-01-02 15:04:05"))

	}

}
