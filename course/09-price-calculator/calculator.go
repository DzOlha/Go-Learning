package _9_price_calculator

import (
	"fmt"
	"go-learning/main/09-price-calculator/filemanager"
	prices2 "go-learning/main/09-price-calculator/prices"
)

var filePath string = "course/09-price-calculator/prices/"

func App() {
	var taxRates = []float64{0, 0.07, 0.1, 0.15}
	dones := make([]chan bool, len(taxRates))
	errorsChans := make([]chan error, len(taxRates))

	inputFile := filePath + "prices.txt"
	//cmd := cmdmanager.New()

	for indexRate, rate := range taxRates {
		dones[indexRate] = make(chan bool)
		errorsChans[indexRate] = make(chan error)

		fm := filemanager.New(inputFile, getOutputPath(indexRate))
		job := prices2.NewTaxIncludedPriceJob(fm, rate)
		go job.Process(dones[indexRate], errorsChans[indexRate])
		//if err != nil {
		//	fmt.Println("Could not process job")
		//	fmt.Println(err)
		//}
	}

	for index := range taxRates {
		select { // wait for one or more channels to emit the value
		case err := <-errorsChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-dones[index]:
			fmt.Println("Done!")
		}
	}
}

func getOutputPath(num int) string {
	return filePath + fmt.Sprintf("result_%v.json", num)
}
