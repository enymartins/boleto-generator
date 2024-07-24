package main

import (
	"html/template"
	"log"
	"os"
)

type BoletoData struct {
	Identificacao    string
	CPFCNPJ          string
	Endereco         string
	CidadeUF         string
	DigitableLine    string
	Value            string
	CodigoBancoComDv string
	Agencia          string
}

func main() {
	tmpl, err := template.ParseFiles("boleto.html")
	if err != nil {
		log.Fatal("Erro ao ler o template HTML:", err)
	}

	p := BoletoData{
		Identificacao:    "Malga Ltda",
		CPFCNPJ:          "12.345.678/0001-99",
		Endereco:         "Rua Exemplo, 123",
		CidadeUF:         "Cidade/UF",
		DigitableLine:    "34191.79001 01043.510047 91020.150008 5 97860026000",
		Value:            "R$ 260,00",
		CodigoBancoComDv: "341-7",
		Agencia:          "0000",
	}

	outFile, err := os.Create("output.html")
	if err != nil {
		log.Fatal("Erro ao criar o arquivo HTML:", err)
	}
	defer outFile.Close()

	err = tmpl.ExecuteTemplate(outFile, "boleto.html", p)
	if err != nil {
		log.Fatal("Erro ao renderizar o template:", err)
	}

	log.Println("HTML gerado e salvo como output.html")
}
