package main

import (
	"bytes"
	"encoding/base64"
	"html/template"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/twooffive"
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
	CodigoBarras     string
}

func barcodeToImage(bc barcode.Barcode, width, height int) (image.Image, error) {
	if width < 440 {
		width = 440
	}
	barcodeImg, err := barcode.Scale(bc, width, height)
	if err != nil {
		return nil, err
	}
	return barcodeImg, nil
}

func imageToBase64(img image.Image) (string, error) {
	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func main() {
	tmpl, err := template.ParseFiles("boleto.html")
	if err != nil {
		log.Fatal("Erro ao ler o template HTML:", err)
	}

	barcodeContent := "34196977800000000011094546464620563383058000"
	barcode, err := twooffive.Encode(barcodeContent, true)
	if err != nil {
		log.Fatal(err)
	}

	newWidth := 440
	newHeight := 13
	img, err := barcodeToImage(barcode, newWidth, newHeight)
	if err != nil {
		log.Fatal(err)
	}

	barcodeBase64, err := imageToBase64(img)
	if err != nil {
		log.Fatal("Erro ao converter a imagem para base64:", err)
	}

	outFile, err := os.Create("img/codigo_de_barras.png")
	if err != nil {
		log.Fatal("Erro ao criar o arquivo PNG:", err)
	}
	defer outFile.Close()

	err = png.Encode(outFile, img)
	if err != nil {
		log.Fatal("Erro ao salvar o código de barras como PNG:", err)
	}

	p := BoletoData{
		Identificacao:    "Malga Ltda",
		CPFCNPJ:          "12.345.678/0001-99",
		Endereco:         "Rua Exemplo, 123",
		CidadeUF:         "Cidade/UF",
		DigitableLine:    "34191098346591006056833830580008999470000000100",
		Value:            "R$ 100,00",
		CodigoBancoComDv: "341-7",
		Agencia:          "0000",
		CodigoBarras:     barcodeBase64,
	}

	htmlFile, err := os.Create("output.html")
	if err != nil {
		log.Fatal("Erro ao criar o arquivo HTML:", err)
	}
	defer htmlFile.Close()

	err = tmpl.ExecuteTemplate(htmlFile, "boleto.html", p)
	if err != nil {
		log.Fatal("Erro ao renderizar o template:", err)
	}

	log.Println("HTML gerado e salvo como output.html")
}
