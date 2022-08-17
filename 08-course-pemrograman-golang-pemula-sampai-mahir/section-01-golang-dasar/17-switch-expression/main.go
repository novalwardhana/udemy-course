package main

import "fmt"

func main() {
	/* Switch */
	club := "chelsea"
	switch club {
	case "arsenal":
		{
			fmt.Println("This is arsenal")
		}
	case "chelsea":
		{
			fmt.Println("This is chelsea")
		}
	default:
		{
			fmt.Println("Other club")
		}
	}

	/* Switch with short statement */
	switch name := "Noval"; len(name) >= 5 {
	case true:
		{
			fmt.Println("Minimum length is fulfil")
		}
	case false:
		{
			fmt.Println("Name is too short")
		}
	}

	/* Switch wihout condition */
	var gpa float64 = 3.40
	switch {
	case gpa >= 3.50:
		{
			fmt.Println("Cumlaude")
		}
	case gpa > 2.75 && gpa < 3.5:
		{
			fmt.Println("Very satisfied")
		}
	case gpa >= 2.00:
		{
			fmt.Println("Not bad")
		}
	default:
		{
			fmt.Println("Bad")
		}
	}
}
