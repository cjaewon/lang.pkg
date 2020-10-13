package cmd

import "lang.pkg/ent"

// cmd = commands list, not a conventional concept in golang

// Bootstrap : Bootstrap All Commands
func Bootstrap(client *ent.Client) {
	Help{client: client}.Init()
	User{client: client}.Init()
	Book{client: client}.Init()
	Voca{client: client}.Init()
}
