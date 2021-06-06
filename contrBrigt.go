package main

import (
	"fmt"
	"image"
	"input"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/disintegration/imaging"
)

func main() {
	title("IMAGE CUTTER")
	fmt.Printf("")
	fmt.Println("cartella input; cartella output; contrasto; luminositá; saturation:")
	cose := input.String()
	start := time.Now()
	elements := strings.Split(cose, ";")
	_, numImages := getAllFileFromDir(elements[0])
	paths, names := genAllPath(elements[0])
	for i := 0; i < numImages; i++ {
		img := readingImages(paths[i])
		fmt.Println("")
		addBrightness := adjustBrightness(img, elements[3])
		addConstrast := adjustConstrast(addBrightness, elements[2])
		// done := adjustHue(addConstrast, elements[4])
		// done := adjustSaturation(addConstrast, elements[4])
		done := adjustGamma(addConstrast, elements[4])
		pathOutput := elements[1] + "/" + names[i]
		copyImg(done, pathOutput)
	}
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("il programma ha finito ed ha impiegato:", elapsed)
}

func copyImg(img image.Image, pathOutput string) {
	err := imaging.Save(img, pathOutput)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}

func adjustGamma(img image.Image, gamma string) image.Image {
	gammaI, _ := strconv.ParseFloat(gamma, 64)
	dstImage := imaging.AdjustGamma(img, gammaI)
	return dstImage
}

func adjustSaturation(img image.Image, sat string) image.Image {
	satI, _ := strconv.ParseFloat(sat, 64)
	dstImage := imaging.AdjustSaturation(img, satI)
	return dstImage
}

func adjustHue(img image.Image, hue string) image.Image {
	hueI, _ := strconv.ParseFloat(hue, 64)
	dstImage := imaging.AdjustHue(img, hueI)
	return dstImage
}

func adjustConstrast(img image.Image, brightness string) image.Image {
	brightnessI, _ := strconv.ParseFloat(brightness, 64)
	dstImage := imaging.AdjustContrast(img, brightnessI)
	return dstImage
}

func adjustBrightness(img image.Image, brightness string) image.Image {
	brightnessI, _ := strconv.ParseFloat(brightness, 64)
	dstImage := imaging.AdjustBrightness(img, brightnessI)
	return dstImage
}

func readingImages(paths string) image.Image {
	img, err := imaging.Open(paths)
	if err != nil {
		panic(err)
	}

	return img
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

func title(text string) {
	myFigure := figure.NewFigure(text, "", true)
	myFigure.Print()
	fmt.Println("")
	fmt.Println("")
}
