package main

import (
	"github.com/charmbracelet/huh"
	"fmt"
	"log"
	"errors"
)

// test with new git config
var (
    category string
    severityLevel int
	description string
	experience bool
)

func main() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose your category").
				Options(
					huh.NewOption("Ideas", "ideia"),
					huh.NewOption("Improvements", "melhoria"),
					huh.NewOption("Compliments", "elogio"),
				).
				Value(&category), // store the category via pointer
	
		),
		huh.NewGroup(
			huh.NewInput().
				Title("What's your it?").
				Value(&description).
				// Validating fields is easy. The form will mark erroneous fields
				// and display error messages accordingly.
				Validate(func(str string) error {
					if str == "" {
						return errors.New("Sorry, we don’t store empty.")
					}
					return nil
				}),
	
			// Option values in selects and multi selects can be any type you
			// want. We’ve been recording strings above, but here we’ll store
			// answers as integers. Note the generic "[int]" directive below.
			huh.NewSelect[int]().
				Title("How fast do you want to implement? Think about its severity").
				Options(
					huh.NewOption("Low", 1),
					huh.NewOption("Medium", 2),
					huh.NewOption("High", 3),
					huh.NewOption("Critic", 4),
				).
				Value(&severityLevel),

			huh.NewConfirm().
				Title("Do you like the experience of the script?").
				Value(&experience),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	if !experience {
		fmt.Println("What? You didn’t like the experience?!")
	} else {
		fmt.Printf("Categoria: %s, Descrição: %s, Dificuldade: %d/4\n", category, description, severityLevel)
	}
}
