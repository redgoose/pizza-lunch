package pdf

import (
	"strconv"
	"time"

	"github.com/go-pdf/fpdf"
	"github.com/redgoose/pizza-day/order"
)

func BuildPDF(roomNumbers []string, orderTotalsByRoom map[string]*order.OrderTotal, orderTotals order.OrderTotal) {
	pdf := fpdf.New("P", "mm", "A4", "")

	orderTotalsTable := func() {
		pdf.AddPage()
		pdf.SetFont("Arial", "B", 16)
		pdf.Cell(50, 15, "Order Totals")
		pdf.Ln(15)

		header := []string{
			"Pizza type",
			"Slices ordered",
			"Pizzas to order",
		}

		// Colors, line width and bold font
		pdf.SetFillColor(255, 0, 0)
		pdf.SetTextColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		pdf.SetLineWidth(.3)
		pdf.SetFont("", "B", 14)

		// 	Header
		w := []float64{35, 45, 45}

		wSum := 0.0
		for _, v := range w {
			wSum += v
		}

		cellHeight := 10.0

		for j, str := range header {
			pdf.CellFormat(w[j], cellHeight, str, "1", 0, "C", true, 0, "")
		}
		pdf.Ln(-1)

		// Color and font restoration
		pdf.SetFillColor(224, 235, 255)
		pdf.SetTextColor(0, 0, 0)
		pdf.SetFont("", "", 0)

		// 	Data
		fill := false

		cheeseTxt := strconv.Itoa(orderTotals.CheeseSlices)
		if orderTotals.ExtraCheeseSlices > 0 {
			cheeseTxt += " (" + strconv.Itoa(orderTotals.ExtraCheeseSlices) + " extra)"
		}

		pdf.CellFormat(w[0], cellHeight, "Cheese", "LR", 0, "C", fill, 0, "")
		pdf.CellFormat(w[1], cellHeight, cheeseTxt, "LR", 0, "C", fill, 0, "")
		pdf.CellFormat(w[2], cellHeight, strconv.Itoa(orderTotals.CheesePizzas), "LR", 0, "C", fill, 0, "")
		pdf.Ln(-1)
		fill = !fill

		pdf.CellFormat(w[0], cellHeight, "Pepperoni", "LR", 0, "C", fill, 0, "")
		pdf.CellFormat(w[1], cellHeight, strconv.Itoa(orderTotals.PepperoniSlices), "LR", 0, "C", fill, 0, "")
		pdf.CellFormat(w[2], cellHeight, strconv.Itoa(orderTotals.PepperoniPizzas), "LR", 0, "C", fill, 0, "")
		pdf.Ln(-1)
		fill = !fill

		pdf.CellFormat(w[0], cellHeight, "GF cheese", "LR", 0, "C", fill, 0, "")
		pdf.CellFormat(w[1], cellHeight, strconv.Itoa(orderTotals.GlutenFreeCheeseSlices), "LR", 0, "C", fill, 0, "")
		pdf.CellFormat(w[2], cellHeight, strconv.Itoa(orderTotals.GlutenFreeCheesePizzas), "LR", 0, "C", fill, 0, "")
		pdf.Ln(-1)
		fill = !fill

		pdf.CellFormat(w[0], cellHeight, "DF cheese", "LR", 0, "C", fill, 0, "")
		pdf.CellFormat(w[1], cellHeight, strconv.Itoa(orderTotals.DairyFreeCheeseSlices), "LR", 0, "C", fill, 0, "")
		pdf.CellFormat(w[2], cellHeight, strconv.Itoa(orderTotals.DairyFreeCheesePizzas), "LR", 0, "C", fill, 0, "")
		pdf.Ln(-1)

		pdf.CellFormat(wSum, 0, "", "T", 0, "", false, 0, "")
		pdf.Ln(10)

		// drinks table
		pdf.SetFillColor(255, 0, 0)
		pdf.SetTextColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		pdf.SetLineWidth(.3)
		pdf.SetFont("", "B", 14)
		pdf.CellFormat(45, cellHeight, "Drinks ordered", "1", 0, "C", true, 0, "")
		pdf.Ln(-1)

		pdf.SetTextColor(0, 0, 0)
		pdf.SetFont("", "", 0)
		pdf.CellFormat(45, cellHeight, strconv.Itoa(orderTotals.Drinks), "LR", 0, "C", false, 0, "")
		pdf.Ln(-1)

		pdf.CellFormat(45, 0, "", "T", 0, "", false, 0, "")
	}

	orderTotalsByRoomTable := func() {
		pdf.AddPageFormat("L", fpdf.SizeType{Wd: 210, Ht: 297})
		pdf.SetFont("Arial", "B", 16)
		pdf.Cell(50, 15, "Order Totals By Room")
		pdf.Ln(15)

		header := []string{
			"Room #",
			"Drinks",
			"Cheese pizzas",
			"Cheese slices",
			"Pepperoni pizzas",
			"Pepperoni slices",
			"GF cheese pizzas",
			"GF cheese slices",
			"DF cheese pizzas",
			"DF cheese slices",
		}

		// Colors, line width and bold font
		pdf.SetFillColor(255, 0, 0)
		pdf.SetTextColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		pdf.SetLineWidth(.3)
		pdf.SetFont("", "B", 8)

		// 	Header
		w := []float64{15, 15, 25, 25, 25, 25, 25, 25, 25, 25}

		wSum := 0.0
		for _, v := range w {
			wSum += v
		}

		cellHeight := 7.0

		for j, str := range header {
			pdf.CellFormat(w[j], cellHeight, str, "1", 0, "C", true, 0, "")
		}
		pdf.Ln(-1)

		// Color and font restoration
		pdf.SetFillColor(224, 235, 255)
		pdf.SetTextColor(0, 0, 0)
		pdf.SetFont("", "", 0)

		// 	Data
		fill := false
		for _, roomNumber := range roomNumbers {

			if _, ok := orderTotalsByRoom[roomNumber]; ok {
			} else {
				continue
			}

			pdf.CellFormat(w[0], cellHeight, roomNumber, "LR", 0, "C", fill, 0, "")
			pdf.CellFormat(w[1], cellHeight, strconv.Itoa(orderTotalsByRoom[roomNumber].Drinks), "LR", 0, "C", fill, 0, "")
			pdf.CellFormat(w[2], cellHeight, strconv.Itoa(orderTotalsByRoom[roomNumber].CheesePizzas), "LR", 0, "C", fill, 0, "")
			pdf.CellFormat(w[3], cellHeight, strconv.Itoa(orderTotalsByRoom[roomNumber].CheeseSlices), "LR", 0, "C", fill, 0, "")
			pdf.CellFormat(w[4], cellHeight, strconv.Itoa(orderTotalsByRoom[roomNumber].PepperoniPizzas), "LR", 0, "C", fill, 0, "")
			pdf.CellFormat(w[5], cellHeight, strconv.Itoa(orderTotalsByRoom[roomNumber].PepperoniSlices), "LR", 0, "C", fill, 0, "")
			pdf.CellFormat(w[6], cellHeight, strconv.Itoa(orderTotalsByRoom[roomNumber].GlutenFreeCheesePizzas), "LR", 0, "C", fill, 0, "")
			pdf.CellFormat(w[7], cellHeight, strconv.Itoa(orderTotalsByRoom[roomNumber].GlutenFreeCheeseSlices), "LR", 0, "C", fill, 0, "")
			pdf.CellFormat(w[8], cellHeight, strconv.Itoa(orderTotalsByRoom[roomNumber].DairyFreeCheesePizzas), "LR", 0, "C", fill, 0, "")
			pdf.CellFormat(w[9], cellHeight, strconv.Itoa(orderTotalsByRoom[roomNumber].DairyFreeCheeseSlices), "LR", 0, "C", fill, 0, "")
			pdf.Ln(-1)
			fill = !fill
		}
		pdf.CellFormat(wSum, 0, "", "T", 0, "", false, 0, "")
	}

	orderTotalsTable()
	orderTotalsByRoomTable()

	fileName := "pizza_day_" + time.Now().Format("20060102") + ".pdf"
	err := pdf.OutputFileAndClose(fileName)
	if err != nil {
		panic(err)
	}
}
