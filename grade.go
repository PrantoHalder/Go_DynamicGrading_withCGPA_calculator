package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Roll    int
	Subject []subject
}
type subject struct {
	sub   string
	marks int
}

func grade() {
	person := []Person{}
	var count int
	var count1 int
	var t int
	//taking the value
	for count3 := 1; count3 <= 10; count3++ {
		fmt.Printf("if you want to input new entry press 1 if not press any key : ")
		fmt.Scanf("%d", &t)
		if t == 1 {
			fmt.Printf("Enter the number of student : ")
			fmt.Scanf("%d", &count)
			for p := 0; p < count; p++ {
				per := Person{}
				fmt.Println("#####################################")
				fmt.Printf("Enter student Name : ")
				fmt.Scanf("%s", &per.Name)
				fmt.Printf("Enter the roll of : ")
				fmt.Scanf("%d", &per.Roll)
				fmt.Printf("Enter the number of subject enrolled : ")
				fmt.Scanf("%d", &count1)
				for k := 0; k < count1; k++ {
					subject := subject{}
					fmt.Printf("Enter the name of %d subject : ", k+1)
					fmt.Scanf("%s", &subject.sub)
					fmt.Printf("Enter marks of %s : ", subject.sub)
					fmt.Scanf("%d", &subject.marks)
					per.Subject = append(per.Subject, subject)
				}
				person = append(person, per)
				fmt.Println("#######################################")

			}

		} else {
			fmt.Println("Thank you")
			break
		}
	}

	//grading
	var roll int
	var count5 int
	for count4 := 1; count4 <= 10; count4++ {
		fmt.Printf("if you want to see the result press 1 else press any key : ")
		fmt.Scanf("%d", &count5)
		if count5 == 1 {
			fmt.Printf("Enter the roll to see result : ")
			fmt.Scanf("%d", &roll)
			for _, value := range person {
				if roll == value.Roll {
					fmt.Println("#######################################")
					fmt.Println("Name : ", value.Name)
					fmt.Println("Roll : ", value.Roll)
					var sum float64
					for _, value := range value.Subject {
						fmt.Println("")
						fmt.Printf("Subject : %s  Marks : %d", value.sub, value.marks)

						if value.marks <= 100 && value.marks >= 80 {
							fmt.Printf(" Grade : A+")
							fmt.Printf(" GPA : 4.00 ")
							sum += 4.00
						} else if value.marks <= 79 && value.marks >= 70 {
							fmt.Printf(" Grade : A")
							fmt.Printf(" GPA : 3.50 ")
							sum += 3.50
						} else if value.marks <= 69 && value.marks >= 60 {
							fmt.Printf(" Grade : A-")
							fmt.Printf(" GPA : 3.00 ")
							sum += 3.00
						} else if value.marks <= 59 && value.marks >= 50 {
							fmt.Printf(" Grade : B")
							fmt.Printf(" GPA : 2.50 ")
							sum += 2.50
						} else if value.marks <= 49 && value.marks >= 40 {
							fmt.Printf(" Grade : C")
							fmt.Printf(" GPA : 2.00 ")
							sum += 2.00
						} else {
							fmt.Printf(" Grade : F")
							fmt.Printf(" GPA : 0.00 ")
							sum += 0.00
						}

					}

					cgpa := sum / float64(count1)
					fmt.Println("")
					fmt.Println("")
					fmt.Printf("			CGPA : %v\n", cgpa)
					fmt.Println("")
					fmt.Println("###############################################")

				}

			
			}//end of the for loop


		} else {
			fmt.Println("Thank You")
			break
		}

	}
}
