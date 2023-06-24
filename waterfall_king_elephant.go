package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Print the welcome message.
	fmt.Println("Welcome to Health In The City!")

	// Ask for user input to select an option.
	var option int
	fmt.Println("What would you like to do? (Enter option number)")
	fmt.Println("1. Check your health")
	fmt.Println("2. Search for a gym")
	fmt.Println("3. Find a nutrition plan")
	fmt.Println("4. Find a doctor")
	fmt.Println("5. Exit")
	fmt.Scanf("%d\n", &option)

	// Switch based on user input.
	switch option {
	case 1:
		checkHealth()
	case 2:
		searchGym()
	case 3:
		findNutritionPlan()
	case 4:
		findDoctor()
	case 5:
		fmt.Println("Goodbye!")
		os.Exit(0)
	default:
		fmt.Println("Invalid option! Goodbye.")
		os.Exit(1)
	}

}

// checkHealth prints out some health tips and information.
func checkHealth() {
	fmt.Println("Checking your health...")
	fmt.Println("\nHere are some tips to help you stay healthy!")
	fmt.Println("1. Eat a balanced diet with lots of fruits and vegetables.")
	fmt.Println("2. Exercise at least 3 times a week.")
	fmt.Println("3. Get enough rest and sleep.")
	fmt.Println("4. Drink plenty of water.")
	fmt.Println("5. Avoid smoking and drugs.")
	fmt.Println("\nThanks for using Health In The City!")
	time.Sleep(5 * time.Second)
	os.Exit(0)
}

// searchGym prints out some gym information.
func searchGym() {
	fmt.Println("Searching for a gym...")
	fmt.Println("\nHere are some gyms near you:")
	fmt.Println("1. Fitness First")
	fmt.Println("2. FitX")
	fmt.Println("3. Anytime Fitness")
	fmt.Println("4. CrossFit")
	fmt.Println("5. Gold's Gym")
	fmt.Println("\nThanks for using Health In The City!")
	time.Sleep(5 * time.Second)
	os.Exit(0)
}

// findNutritionPlan prints out some nutrition information.
func findNutritionPlan() {
	fmt.Println("Finding a nutrition plan...")
	fmt.Println("\nHere are some nutrition plans you can try:")
	fmt.Println("1. Whole food plant-based diet")
	fmt.Println("2. Mediterranean diet")
	fmt.Println("3. Ketogenic diet")
	fmt.Println("4. Intermittent fasting")
	fmt.Println("5. Low-carbohydrate diet")
	fmt.Println("\nThanks for using Health In The City!")
	time.Sleep(5 * time.Second)
	os.Exit(0)
}

// findDoctor prints out some doctor information.
func findDoctor() {
	fmt.Println("Finding a doctor...")
	fmt.Println("\nHere are some doctors near you:")
	fmt.Println("1. Dr. Smith")
	fmt.Println("2. Dr. Jones")
	fmt.Println("3. Dr. Patel")
	fmt.Println("4. Dr. Williams")
	fmt.Println("5. Dr. Johnson")
	fmt.Println("\nThanks for using Health In The City!")
	time.Sleep(5 * time.Second)
	os.Exit(0)
}