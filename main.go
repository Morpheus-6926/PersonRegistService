package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name string
	Age  uint
	CPF  string
}

var registers = make(map[int]Person)
var option = -1

func main() {
	mainMenu()
}
func mainMenu() {
	fmt.Println("############################################")
	fmt.Println("Welcome to the Person Registration Service")
	fmt.Println("Please Select an option:\n1-Search register\n2-Register person\n3-Alter register\n4-Delete register")
	for {
		fmt.Scan(&option)
		switch option {
		case 1:
			search()
		case 2:
			register()
		case 3:
		case 4:
		default:
			fmt.Println("Choose a valid option")
		}
	}
}
func register() {
	//F1:
	for {
		var register = Person{}
		fmt.Println("############################################")
		fmt.Println("Person Registration")
		fmt.Println("Enter the name:")
		fmt.Scan(&register.Name)
		fmt.Println("Enter the age:")
		fmt.Scan(&register.Age)
		fmt.Println("Enter the CPF(11 digits, only numbers)")
		fmt.Scan(&register.CPF)
	F1:
		for {
			if _, err := strconv.Atoi(register.CPF); err == nil {
				break F1
			} else {
				println("invalid CPF, try again")
			}
		}

		fmt.Println("############################################")
		fmt.Println("Confirm  Registation")
		fmt.Printf("Name: %v\n", register.Name)
		fmt.Printf("Age: %v\n", register.Age)
		fmt.Printf("CPF: %v\n", formatCPF(register.CPF))
		fmt.Println("1-Yes\n2-No")
	F2:
		for {
			fmt.Scan(&option)
			switch option {
			case 1:

				if i, err := strconv.Atoi(register.CPF); err == nil {
					register.CPF = formatCPF(register.CPF)
					registers[i] = register
				}
			case 2:

			default:
				fmt.Println("Choose a valid option")
			}
			//F3:
			for {
				fmt.Println("1-back to main menu\n2-make another registration")
				fmt.Scan(&option)
				switch option {
				case 1:
					mainMenu()
				case 2:
					break F2
				default:
					fmt.Println("Choose a valid option")
				}
			}
		}
	}

}
func search() {
	var serchedCpf string
	
		fmt.Println("############################################")
		fmt.Println("Register Seach")
		for {
		fmt.Println("enter a CPF(11 digits, only numbers) to find a register or enter \"all\" to show all registers")
		fmt.Scan(&serchedCpf)
		if serchedCpf == "all" {
			for i, _ := range registers {
				showRegister(i)
			}

		} else {
			if i, err := strconv.Atoi(serchedCpf); err == nil && len(serchedCpf) == 11 {
				showRegister(i)
			F2:
				for {
					fmt.Println("1-back to main menu\n2-make another search")
					fmt.Scan(&option)
					switch option {
					case 1:
						mainMenu()
					case 2:
						break F2
					default:
						fmt.Println("Choose a valid option")
					}
				}
			} else {
				fmt.Println("You entered a invalid CPF")
			}
		}
		F3:
			for {
				fmt.Println("1-back to main menu\n2-make another search")
				fmt.Scan(&option)
				switch option {
				case 1:
					mainMenu()
				case 2:
					break F3
				default:
					fmt.Println("Choose a valid option")
				}
			}
	}

}

/*func delete() {

}
func alter() {

}*/
func showRegister(_cpf int) {
	fmt.Println("#####################")
	register := registers[_cpf]
	fmt.Printf("Name: %v\n", register.Name)
	fmt.Printf("Age: %v\n", register.Age)
	fmt.Printf("CPF: %v\n", register.CPF)
}
func formatCPF(cpf string) string {
	var n string
	for i, _ := range cpf {
		n = n + cpf[i:i+1]
		if i == 2 || i == 5 {
			n = n + "."
		} else if i == 8 {
			n = n + "-"
		}
	}
	return n
}
