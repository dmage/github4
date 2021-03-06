package main

import (
	"context"
	"log"
	"os"

	"github.com/dmage/github4"
	"golang.org/x/oauth2"
)

func main() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	githubClient := github4.Client{Client: httpClient}

	var r struct {
		Repository github4.Repository
	}
	err := githubClient.Do(`
		query($owner: String!, $name: String!, $prnumber: Int!) {
			repository(owner: $owner, name: $name) {
				description
				pullRequest(number: $prnumber) {
					author {
						__typename
						login
						... on User {
							company
						}
					}
					timeline(first: 10) {
						nodes {
							__typename
						}
					}
				}
			}
		}
	`, map[string]interface{}{
		"owner":    "openshift",
		"name":     "origin",
		"prnumber": 14521,
	}, &r)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#+v", r)
	log.Println(*r.Repository.Description())
	log.Printf("%#+v", r.Repository.PullRequest().Author())
	log.Println(r.Repository.PullRequest().Author().Login())
	for _, item := range r.Repository.PullRequest().Timeline().Nodes() {
		log.Printf("%T", item)
	}
}
