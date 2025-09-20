package prices

import (
	"fmt"
	"go-learning/main/09-price-calculator/conversion"
)

type IOManager interface {
	ReadLines() ([]string, error)
	WriteResult(data interface{}) error
}

type TaxIncludedPriceJob struct {
	IOManager         IOManager         `json:"-"`
	TaxRate           float64           `json:"tax_rate"`
	InputPrices       []float64         `json:"input_prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
}

func NewTaxIncludedPriceJob(io IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:         io,
		TaxRate:           taxRate,
		TaxIncludedPrices: make(map[string]string, 4),
	}
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errChan chan error) {
	err := job.loadData()
	if err != nil {
		errChan <- err
		return
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		result[format(price)] = format(price * (1 + job.TaxRate))
	}

	job.TaxIncludedPrices = result

	job.IOManager.WriteResult(job)

	doneChan <- true
}

func (job *TaxIncludedPriceJob) loadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		return err
	}

	job.InputPrices = prices

	return nil
}

func format(num float64) string {
	return fmt.Sprintf("%.2f", num)
}
