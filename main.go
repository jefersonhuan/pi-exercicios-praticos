package main

import (
	"fmt"
	"github.com/anthonynsimon/bild/blend"
	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/segment"
	"image"
	"log"
)

const sourceDir = "imagens/"
const resDir = "resultados/"

func saveImage(img image.Image, filename string) {
	err := imgio.Save(resDir+filename, img, imgio.PNGEncoder())
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Arquivo processado em:", filename)
	}
}

func ex01() image.Image {
	img, err := imgio.Open(sourceDir + "Exercicio1.png")

	if err != nil {
		log.Fatal(err)
	}

	gray := effect.Grayscale(img)
	eroded := effect.Dilate(gray, 12)

	saveImage(eroded, "ex01.png")

	return eroded
}

func ex02(img image.Image) {
	dilated := effect.Erode(img, 12)

	saveImage(dilated, "ex02.png")
}

func ex03() {
	coins, err := imgio.Open(sourceDir + "moedas.jpg")

	if err != nil {
		log.Fatal(err)
	}

	inverted := effect.Invert(coins)
	segmented := segment.Threshold(inverted, 80)

	eroded := effect.Erode(segmented, 1)
	dilated := effect.Dilate(eroded, 6)

	saveImage(dilated, "ex03-segmentado.png")

	final := blend.Add(coins, dilated)

	saveImage(final, "ex03.png")
}

func main() {
	fmt.Println("Executando Exercício 01...")
	eroded := ex01()

	fmt.Println("Executando Exercício 02...\nUtilizando saída do exercício anterior")
	ex02(eroded)

	fmt.Println("Executando Exercício 05...")
	ex03()

	fmt.Println("Para 'filtrar' a imagem segmentada foi preciso aplicar erosão + dilatação, o que alterou o tamanho dos elementos com relação à imagem original")
}
