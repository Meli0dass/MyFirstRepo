package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"
)

const IssueBaseUrl = "https://api.github.com/repos/Meli0dass/myFirstRepo/issues"

type Issue struct {
	Number    int       `json:"number,omitempty"`
	HTMLURL   string    `json:"html_url,omitempty"`
	Title     string    `json:"title,omitempty"`
	State     string    `json:"state,omitempty"`
	User      *User     `json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Body      string    `json:"body,omitempty"`
}

type User struct {
	Login   string `json:"login"`
	HTMLURL string `json:"html_url"`
}

type IssueGetResult struct {
	Issues []Issue
}

const Authorization = "Bearer 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"

func IssueGet(terms []string) (*IssueGetResult, error) {
	queryExpr := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssueBaseUrl + "?q=" + queryExpr)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssueGetResult
	if err := json.NewDecoder(resp.Body).Decode(&result.Issues); err != nil {
		resp.Body.Close()
		return nil, err
	}
	defer resp.Body.Close()
	return &result, nil
}

func IssueCreate(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("invalid number of args: want 2, got %d", len(args))
	}
	reqBody, err := json.Marshal(Issue{Title: args[0], Body: args[1]})
	request, err := http.NewRequest("POST", IssueBaseUrl, strings.NewReader(string(reqBody)))
	if err != nil {
		return err
	}
	request.Header.Set("Authorization", Authorization)
	resp, err := http.DefaultClient.Do(request)
	if resp.StatusCode != http.StatusCreated {
		resp.Body.Close()
		return fmt.Errorf("create issue failed: %s", resp.Status)
	}

	return nil
}

func IssueClose() {

}

func IssueUpdate() {

}

const (
	SEARCH = "search"
	CREATE = "create"
	DELETE = "delete"
	UPDATE = "update"
)

func main() {
	action := flag.String("action", "search", "options: search, create, delete, update")
	flag.Parse()
	args := flag.Args()
	switch *action {
	case SEARCH:
		result, err := IssueGet(args)
		if err != nil {
			log.Fatal(err)
		}
		for _, item := range result.Issues {
			fmt.Printf("#%-5d %9.9s %.55s %.100s\n", item.Number, item.User.Login, item.Title, item.Body)
		}
	case CREATE:
		err := IssueCreate(args)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("create issue success")
	case UPDATE:
		cmd := exec.Command("vi", "test.txt")
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout

		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("please choose one action")
	}
}
