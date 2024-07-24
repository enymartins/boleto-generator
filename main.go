// package main

// import (
// 	"fmt"
// 	"log"
	

// 	"github.com/jung-kurt/gofpdf"
// )

// func main() {
// 	// Dados do boleto
// 	digitableLine := "34191.79001 01043.510047 91020.150008 5 97860026000"
// 	beneficiary := "Microhouse Informática S/C Ltda"
// 	dueDate := "23/07/2024"
// 	value := "R$ 260,00"
// 	agency := "0049"
// 	account := "10201-5"
// 	documentDate := "23/07/2024"
// 	documentNumber := "DF 00113"
// 	instructions := "Taxa de visita de suporte\nApós o vencimento R$ 0,80 ao dia"

// 	// Criar novo PDF
// 	pdf := gofpdf.New("P", "mm", "A4", "")
// 	pdf.AddPage()

// 	// Adicionar título
// 	pdf.SetFont("Arial", "B", 16)
// 	pdf.CellFormat(190, 10, "Boleto Bancário", "0", 1, "C", false, 0, "")

// 	// Adicionar linha digitável
// 	pdf.SetFont("Arial", "B", 12)
// 	pdf.CellFormat(190, 10, "Linha Digitável", "0", 1, "L", false, 0, "")
// 	pdf.SetFont("Arial", "", 12)
// 	pdf.CellFormat(190, 10, digitableLine, "0", 1, "L", false, 0, "")

// 	// Adicionar informações do beneficiário
// 	pdf.Ln(10)
// 	pdf.SetFont("Arial", "B", 12)
// 	pdf.CellFormat(190, 10, "Beneficiário", "0", 1, "L", false, 0, "")
// 	pdf.SetFont("Arial", "", 12)
// 	pdf.CellFormat(190, 10, beneficiary, "0", 1, "L", false, 0, "")

// 	// Adicionar data de vencimento
// 	pdf.Ln(10)
// 	pdf.SetFont("Arial", "B", 12)
// 	pdf.CellFormat(190, 10, "Data de Vencimento", "0", 1, "L", false, 0, "")
// 	pdf.SetFont("Arial", "", 12)
// 	pdf.CellFormat(190, 10, dueDate, "0", 1, "L", false, 0, "")

// 	// Adicionar valor
// 	pdf.Ln(10)
// 	pdf.SetFont("Arial", "B", 12)
// 	pdf.CellFormat(190, 10, "Valor", "0", 1, "L", false, 0, "")
// 	pdf.SetFont("Arial", "", 12)
// 	pdf.CellFormat(190, 10, value, "0", 1, "L", false, 0, "")

// 	// Adicionar agência e conta
// 	pdf.Ln(10)
// 	pdf.SetFont("Arial", "B", 12)
// 	pdf.CellFormat(190, 10, "Agência / Código do Cedente", "0", 1, "L", false, 0, "")
// 	pdf.SetFont("Arial", "", 12)
// 	pdf.CellFormat(190, 10, fmt.Sprintf("%s / %s", agency, account), "0", 1, "L", false, 0, "")

// 	// Adicionar data e número do documento
// 	pdf.Ln(10)
// 	pdf.SetFont("Arial", "B", 12)
// 	pdf.CellFormat(190, 10, "Data do Documento / Nº do Documento", "0", 1, "L", false, 0, "")
// 	pdf.SetFont("Arial", "", 12)
// 	pdf.CellFormat(190, 10, fmt.Sprintf("%s / %s", documentDate, documentNumber), "0", 1, "L", false, 0, "")

// 	// Adicionar instruções
// 	pdf.Ln(10)
// 	pdf.SetFont("Arial", "B", 12)
// 	pdf.CellFormat(190, 10, "Instruções", "0", 1, "L", false, 0, "")
// 	pdf.SetFont("Arial", "", 12)
// 	pdf.MultiCell(190, 10, instructions, "0", "L", false)

// 	// Salvar PDF
// 	fileName := "boleto.pdf"
// 	err := pdf.OutputFileAndClose(fileName)
// 	if err != nil {
// 		log.Fatalf("Erro ao salvar PDF: %v", err)
// 	}

// 	fmt.Println("PDF gerado com sucesso!")

// 	// Abrir PDF automaticamente
// }
