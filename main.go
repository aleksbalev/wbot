package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/joho/godotenv"
	"github.com/ktrysmt/go-bitbucket"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func main() {
	godotenv.Load(".env")

	token := os.Getenv("SLACK_AUTH_TOKEN")
	appToken := os.Getenv("SLACK_APP_TOKEN")
	bbKey := os.Getenv("BB_KEY")
	bbSecret := os.Getenv("BB_SECRET")

	client := slack.New(token, slack.OptionDebug(true), slack.OptionAppLevelToken(appToken))
	bbClient := bitbucket.NewOAuthClientCredentials(bbKey, bbSecret)

	user, err := bbClient.User.Profile()
	if err != nil {
		fmt.Println("Not logged in or token is invalid:", err)
		return
	}

	fmt.Printf("Profile %v", user.Nickname)

	repoFileOptions := bitbucket.RepositoryFilesOptions{
		Owner:    "AleksBL",
		RepoSlug: "wbot-test",
		Ref:      "main",
		Path:     "CHANGELOG.md",
	}

	fileContent, err := bbClient.Repositories.Repository.GetFileContent(&repoFileOptions)
	if err != nil {
		fmt.Printf("Error while getting repo: %s", err)
	}

  transformedContent, notFoundText := transformChangelog(fileContent, nil)
  if notFoundText != nil {
	fmt.Printf("Text transform error: %s", notFoundText)
  }

  fmt.Printf("Transformed changelog: %s", transformedContent)

	socketClient := socketmode.New(
		client,
		socketmode.OptionDebug(true),
		socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func(ctx context.Context, client *slack.Client, socketClient *socketmode.Client, bbClient *bitbucket.Client) {
		for {
			select {
			case <-ctx.Done():
				log.Println("Shutting down socketmode listener")
				return
			case event := <-socketClient.Events:
				switch event.Type {
				case socketmode.EventTypeSlashCommand:
					command, ok := event.Data.(slack.SlashCommand)
					if !ok {
						log.Printf("Could not type cast the message to a SlashCommand: %v\n", command)
					}

					payload, err := handleSlashCommand(command, client, bbClient)
					if err != nil {
						log.Fatal(err)
					}

					socketClient.Ack(*event.Request, payload)
				}
			}
		}
	}(ctx, client, socketClient, bbClient)

	// socketClient.Run()
}

func handleSlashCommand(command slack.SlashCommand, client *slack.Client, bbClient *bitbucket.Client) (interface{}, error) {
	switch command.Command {
	case "/release-log":
		return nil, handleReleaseLogCommand(command, client, bbClient)
	}
	return nil, nil
}

func handleReleaseLogCommand(command slack.SlashCommand, client *slack.Client, bbClient *bitbucket.Client) error {
	attachment := slack.Attachment{}

	attachment.Text = fmt.Sprintf("Repository name: %s", bbClient.Repositories.Repository.Name)

	_, _, err := client.PostMessage(command.ChannelID, slack.MsgOptionAttachments(attachment))
	if err != nil {
		return err
	}

	//  params := slack.FileUploadParameters{
	//   Channels: []string{os.Getenv("SLACK_CHANNEL_ID")},
	//   File: "./files/CHANGELOG.md",
	//  }
	//  _, err := client.UploadFile(params)
	//  if err != nil {
	//    return err
	//  }

	return nil
}

func currentTime() string {
	loc, err := time.LoadLocation("CET")
	if err != nil {
		log.Fatal("Can't load location: ", err)
	}

	now := time.Now()
	czechia := now.In(loc).Format("02.01.2006 15:04:05")

	return czechia
}

func transformChangelog(content []byte, version *string) (transformedContent *string, notFoundText *string) {
	var startRegex regexp.Regexp
  var matches []string

	if version == nil {
		startRegex = *regexp.MustCompile(`### Unreleased - [\d-]+(.*?)###`)
	} else {
		startRegex = *regexp.MustCompile(fmt.Sprintf(`(?s)(%s.*?)(### .*?|\z)`, regexp.QuoteMeta(*version)))
	}

  matches = startRegex.FindStringSubmatch(string(content[:]))
  if len(matches) > 1 {
    notFoundText := "Used version not found"

    return nil, &notFoundText 
  }

  fmt.Printf("matches: %s", matches[1])
  cuttedContent := matches[1]

	return &cuttedContent, nil
}

func trimLeadingWhitespace(s string) string {
	re := regexp.MustCompile(`(?m)^\s+`)
	return re.ReplaceAllString(s, "")
}
