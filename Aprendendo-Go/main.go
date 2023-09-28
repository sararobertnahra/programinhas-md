package main

import (
	"fmt"
	"strings"
	"time"
	"sync"
)

const totalTickets int = 50
var conferenceName = "Go Conference"
var ticketsFaltando uint = 50
var reservas = make([]pessoaDados, 0)

type pessoaDados struct {
	nomePessoa string
	sobrenome string
	email string
	ticketsPessoa uint 
	isOptedInForNewsletter bool
}

var wg = sync.WaitGroup{

}

func main (){
	
	greetUsers()
	

	nomePessoa, sobrenome, email, ticketsPessoa := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(nomePessoa, sobrenome, email, ticketsPessoa, ticketsFaltando)
		
	if isValidName && isValidEmail && isValidTicketNumber {
			
		reservaTicket(ticketsPessoa, nomePessoa, sobrenome, email)

		wg.Add(1)
		go sendTicket(ticketsPessoa, nomePessoa, sobrenome, email)

		primeirosNomes := getPrimeirosNomes()
		fmt.Printf("Os primeiros nomes das reservas são: %v\n", primeirosNomes)
			
		if ticketsFaltando == 0 {
			fmt.Println("Tickets esgotados.")
		}
	} else {
		if !isValidName {
			fmt.Println ("Nome ou sobrenome muito curtos.")
		}
		if !isValidEmail {
			fmt.Println ("Email inválido, faltou @.")
		}
		if !isValidTicketNumber {
			fmt.Println ("Número de Tickets inválido!")
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking applicatin!\n", conferenceName)
	fmt.Printf ("Temos um total de %v tickets e %v tickets que ainda restam!\n", totalTickets, ticketsFaltando)
	fmt.Println("Pegue seu ticket aqui!")
}

func getPrimeirosNomes() []string {
	primeirosNomes := []string{}
	for _ , reserva := range reservas {
		primeirosNomes = append(primeirosNomes, reserva.nomePessoa)
	}
	return primeirosNomes	
}


func getUserInput() (string, string, string, uint){
	var nomePessoa string
	var ticketsPessoa uint
	var sobrenome string
	var email string
	// pedir o nome da pessoa
	fmt.Println("Digite o seu nome:")
	fmt.Scan(&nomePessoa)
	
	fmt.Println("Digite seu sobrenome: ")
	fmt.Scan(&sobrenome)
	
	fmt.Println("Digite seu email: ")
	fmt.Scan(&email)
	
	fmt.Println("Quantos tickets deseja comprar: ")
	fmt.Scan(&ticketsPessoa)

	return nomePessoa, sobrenome, email, ticketsPessoa
}

func reservaTicket(ticketsPessoa uint, nomePessoa string, sobrenome string, email string) {
	ticketsFaltando = ticketsFaltando - ticketsPessoa

	// criar um map para uma pessoa
	var pessoaDados = pessoaDados {
		nomePessoa: nomePessoa,
		sobrenome: sobrenome,
		email: email,
		ticketsPessoa: ticketsPessoa,
	}
	 
	reservas = append(reservas, pessoaDados)
	fmt.Printf("Lista de reservas é %v\n", reservas)
		 
	fmt.Printf("Obrigada %v %v por comprar %v tickets. Você receberá confirmação via email %v\n", nomePessoa, sobrenome, ticketsPessoa, email)
	fmt.Printf("%v tickets restantes para %v\n", ticketsFaltando, conferenceName)
}

func validateUserInput(nomePessoa string, sobrenome string, email string, ticketsPessoa uint, ticketsFaltando uint) (bool, bool, bool) {
	isValidName := len(nomePessoa) >= 2 && len(sobrenome) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := ticketsPessoa > 0 && ticketsPessoa <= ticketsFaltando 
	return isValidName, isValidEmail, isValidTicketNumber
}


func sendTicket(ticketsPessoa uint, nomePessoa string, sobrenome string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets para %v %v", ticketsPessoa, nomePessoa, sobrenome)
	fmt.Println("##########")
	fmt.Printf("Mandando os tickets:\n %v \npara o email %v\n", ticket, email)
	fmt.Println("##########")
	wg.Done()
}
