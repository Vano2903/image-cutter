// TODO input cartella di input x
// TODO input cartella di output x
// TODO input top left immagini x
// TODO input dimension crop immagini x
// TODO input dimensioni per il resize dell'immagine x

// TODO prendere tutte le immagini all'interno della cartella x
// TODO croppare le immagini x
// TODO ridimensionarle x
// TODO salvare le immagini nella cartella di output x

package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"input"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
)

func main() {
	title("IMAGE CUTTER")
	fmt.Printf("")
	fmt.Println("cartella input; cartella output; top; left; lunghezza; altezza; lunghezza finale:")
	cose := input.String()
	start := time.Now()
	elements := strings.Split(cose, ";")
	_, numImages := getAllFileFromDir(elements[0])
	paths, names := genAllPath(elements[0])
	var image image.Image
	for i := 0; i < numImages; i++ {
		image = readingImages(paths[i])
		fmt.Println("")

		img := cropImages(image, elements[2], elements[3], elements[4], elements[5])
		pathOutput := elements[1] + "/" + names[i]
		resizeImg(img, elements[6], pathOutput, names[i])
		fmt.Println("immagini rimanenti: ", numImages-i)
	}
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("il programma ha finito ed ha impiegato:", elapsed)
}

func resizeImg(img image.Image, size string, path string, name string) {
	sizeI, _ := strconv.ParseUint(size, 10, 32)
	m := resize.Resize(uint(sizeI), 0, img, resize.NearestNeighbor)
	outFiles(m, path, name)
}

func outFiles(finishedImg image.Image, path string, name string) {
	out, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	jpeg.Encode(out, finishedImg, nil)
	fmt.Println(name, "croppata e salvata")
}

func cropImages(img image.Image, top string, left string, width string, height string) image.Image {
	topI, _ := strconv.Atoi(top)
	leftI, _ := strconv.Atoi(left)
	widthI, _ := strconv.Atoi(width)
	heightI, _ := strconv.Atoi(height)
	croppedImg, err := cutter.Crop(img, cutter.Config{
		Width:  widthI,
		Height: heightI,
		Mode:   cutter.TopLeft,
		Anchor: image.Point{topI, leftI},
	})
	if err != nil {
		panic(err)
	}
	return croppedImg
}

func getAllFileFromDir(dirIn string) ([]string, int) {
	files, err := ioutil.ReadDir(dirIn)
	var names []string
	if err != nil {
		log.Fatal(err)
	}
	num := 0
	for _, file := range files {
		fmt.Println(file.Name())
		names = append(names, file.Name())
		num++
	}
	return names, num
}

func genAllPath(dirIn string) ([]string, []string) {
	names, _ := getAllFileFromDir(dirIn)
	var paths []string
	checkLastCharDir := dirIn[len(dirIn)-1:]
	if checkLastCharDir == "/" || checkLastCharDir == "\\" {
		strings.TrimSuffix(dirIn, "/")
		strings.TrimSuffix(dirIn, "\\")
	}
	for _, file := range names {
		completePath := dirIn + "/" + file
		paths = append(paths, completePath)
	}
	return paths, names
}

func readingImages(paths string) image.Image {
	existingImageFile, err := os.Open(paths)
	if err != nil {
		panic(err)
	}
	defer existingImageFile.Close()

	imageData, _, err := image.Decode(existingImageFile)
	if err != nil {
		panic(err)
	}

	return imageData
}

func title(text string) {
	myFigure := figure.NewFigure(text, "", true)
	myFigure.Print()
	fmt.Println("")
	fmt.Println("")
}
