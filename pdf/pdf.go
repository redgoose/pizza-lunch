package pdf

import (
	"strconv"
	"strings"

	"github.com/go-pdf/fpdf"
	"github.com/redgoose/pizza-lunch/config"
	"github.com/redgoose/pizza-lunch/order"
)

func GeneratePDF(roomNumbers []string,
	roomInfo map[string]config.Room,
	ordersByRoom map[string][]order.Order,
	orderTotalsByRoom map[string]*order.OrderTotal,
	orderTotals order.OrderTotal,
	fileName string) {
	pdf := fpdf.New("P", "mm", "A4", "")

	orderTotalsPage := func() {
		pdf.AddPage()

		pdf.SetFont("Arial", "B", 24)
		pdf.Cell(50, 15, "Pizza lunch")
		pdf.Ln(15)

		pdf.SetFont("Arial", "B", 16)
		pdf.Cell(50, 15, "Order totals")
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

		// Header
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

		pdf.CellFormat(w[0], cellHeight, "GF Cheese", "LR", 0, "C", fill, 0, "")
		pdf.CellFormat(w[1], cellHeight, strconv.Itoa(orderTotals.GlutenFreeCheeseSlices), "LR", 0, "C", fill, 0, "")
		pdf.CellFormat(w[2], cellHeight, strconv.Itoa(orderTotals.GlutenFreeCheesePizzas), "LR", 0, "C", fill, 0, "")
		pdf.Ln(-1)
		fill = !fill

		pdf.CellFormat(w[0], cellHeight, "GF Pepperoni", "LR", 0, "C", fill, 0, "")
		pdf.CellFormat(w[1], cellHeight, strconv.Itoa(orderTotals.GlutenFreePepperoniSlices), "LR", 0, "C", fill, 0, "")
		pdf.CellFormat(w[2], cellHeight, strconv.Itoa(orderTotals.GlutenFreePepperoniPizzas), "LR", 0, "C", fill, 0, "")
		pdf.Ln(-1)
		fill = !fill

		pdf.CellFormat(w[0], cellHeight, "DF Cheese", "LR", 0, "C", fill, 0, "")
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

	orderTotalsByRoomPage := func() {
		pdf.AddPageFormat("L", fpdf.SizeType{Wd: 210, Ht: 297})
		pdf.SetFont("Arial", "B", 16)
		pdf.Cell(50, 15, "Order totals by room")
		pdf.Ln(15)

		header := []string{
			"Room #\n ",
			"Drinks\n ",
			"Cheese pizzas",
			"Cheese slices",
			"Pepperoni pizzas",
			"Pepperoni slices",
			"GF Cheese pizzas",
			"GF Cheese slices",
			"GF Pepperoni pizzas",
			"GF Pepperoni slices",
			"DF Cheese pizzas",
			"DF Cheese slices",
		}

		// Colors, line width and bold font
		pdf.SetFillColor(255, 0, 0)
		pdf.SetTextColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		pdf.SetLineWidth(.3)
		pdf.SetFont("", "B", 10)

		// Header
		w := []float64{18, 18, 18, 18, 23, 23, 23, 23, 27, 27, 23, 23}

		wSum := 0.0
		for _, v := range w {
			wSum += v
		}

		cellHeight := 7.0

		for j, str := range header {
			x, y := pdf.GetXY()
			pdf.MultiCell(w[j], cellHeight, str, "1", "C", true)
			pdf.SetXY(x+w[j], y)
		}
		pdf.Ln(-1)
		x, y := pdf.GetXY()
		pdf.SetXY(x, y+cellHeight)

		// Color and font restoration
		pdf.SetFillColor(224, 235, 255)
		pdf.SetTextColor(0, 0, 0)
		pdf.SetFont("", "", 12)

		// Data
		fill := false
		for _, roomNumber := range roomNumbers {
			pdf.CellFormat(w[0], cellHeight, roomNumber, "LR", 0, "C", fill, 0, "")
			pdf.CellFormat(w[1], cellHeight, removeZero(orderTotalsByRoom[roomNumber].Drinks), "LR", 0, "C", fill, 0, "")
			pdf.CellFormat(w[2], cellHeight, removeZero(orderTotalsByRoom[roomNumber].CheesePizzas), "LR", 0, "C", fill, 0, "")
			pdf.CellFormat(w[3], cellHeight, removeZero(orderTotalsByRoom[roomNumber].CheeseSlices), "LR", 0, "C", fill, 0, "")
			pdf.CellFormat(w[4], cellHeight, removeZero(orderTotalsByRoom[roomNumber].PepperoniPizzas), "LR", 0, "C", fill, 0, "")
			pdf.CellFormat(w[5], cellHeight, removeZero(orderTotalsByRoom[roomNumber].PepperoniSlices), "LR", 0, "C", fill, 0, "")
			pdf.CellFormat(w[6], cellHeight, removeZero(orderTotalsByRoom[roomNumber].GlutenFreeCheesePizzas), "LR", 0, "C", fill, 0, "")
			pdf.CellFormat(w[7], cellHeight, removeZero(orderTotalsByRoom[roomNumber].GlutenFreeCheeseSlices), "LR", 0, "C", fill, 0, "")
			pdf.CellFormat(w[8], cellHeight, removeZero(orderTotalsByRoom[roomNumber].GlutenFreePepperoniPizzas), "LR", 0, "C", fill, 0, "")
			pdf.CellFormat(w[9], cellHeight, removeZero(orderTotalsByRoom[roomNumber].GlutenFreePepperoniSlices), "LR", 0, "C", fill, 0, "")
			pdf.CellFormat(w[10], cellHeight, removeZero(orderTotalsByRoom[roomNumber].DairyFreeCheesePizzas), "LR", 0, "C", fill, 0, "")
			pdf.CellFormat(w[11], cellHeight, removeZero(orderTotalsByRoom[roomNumber].DairyFreeCheeseSlices), "LR", 0, "C", fill, 0, "")
			pdf.Ln(-1)
			fill = !fill
		}
		pdf.CellFormat(wSum, 0, "", "T", 0, "", false, 0, "")
	}

	ordersForRoomPage := func(roomInfoByNumber map[string]config.Room, roomNumber string) {
		pdf.AddPage()
		pdf.SetFont("Arial", "B", 16)
		pdf.Cell(50, 15, "Orders for Room "+strings.TrimLeft(roomNumber, "0")+" - "+roomInfoByNumber[roomNumber].Teacher)
		pdf.Ln(10)

		pdf.SetFont("Arial", "B", 14)
		pdf.Cell(50, 15, "Student orders")
		pdf.Ln(15)

		header := []string{
			"Student name",
			"Order",
		}

		// Colors, line width and bold font
		pdf.SetFillColor(255, 0, 0)
		pdf.SetTextColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		pdf.SetLineWidth(.3)
		pdf.SetFont("", "B", 10)

		// Header
		w := []float64{50, 130}

		wSum := 0.0
		for _, v := range w {
			wSum += v
		}

		cellHeight := 7.0

		for j, str := range header {
			pdf.CellFormat(w[j], cellHeight, str, "1", 0, "", true, 0, "")
		}
		pdf.Ln(-1)

		// Color and font restoration
		pdf.SetFillColor(224, 235, 255)
		pdf.SetTextColor(0, 0, 0)
		pdf.SetFont("", "", 0)
		cellHeight = 6.0

		// Data
		fill := false
		for _, order := range ordersByRoom[roomNumber] {
			pdf.CellFormat(w[0], cellHeight, order.Name, "LR", 0, "", fill, 0, "")
			pdf.CellFormat(w[1], cellHeight, order.Raw, "LR", 0, "", fill, 0, "")
			pdf.Ln(-1)
			fill = !fill
		}
		pdf.CellFormat(wSum, 0, "", "T", 0, "", false, 0, "")
		pdf.Ln(-1)

		pdf.SetFont("Arial", "B", 14)
		pdf.Cell(50, 15, "Order totals")
		pdf.Ln(15)

		// Colors, line width and bold font
		pdf.SetFillColor(255, 0, 0)
		pdf.SetTextColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		pdf.SetLineWidth(.3)
		pdf.SetFont("", "B", 12)
		cellHeight = 7.0

		pdf.CellFormat(35, cellHeight, "Pizza type", "1", 0, "C", true, 0, "")
		pdf.CellFormat(25, cellHeight, "Pizzas", "1", 0, "C", true, 0, "")
		pdf.CellFormat(25, cellHeight, "Slices", "1", 0, "C", true, 0, "")
		x, y := pdf.GetXY() // save position for drinks table
		pdf.Ln(-1)

		// Color and font restoration
		pdf.SetFillColor(224, 235, 255)
		pdf.SetTextColor(0, 0, 0)
		pdf.SetFont("", "", 0)

		fill = false
		pdf.CellFormat(35, cellHeight, "Cheese", "LR", 0, "Cß", fill, 0, "")
		pdf.CellFormat(25, cellHeight, strconv.Itoa(orderTotalsByRoom[roomNumber].CheesePizzas), "LR", 0, "C", fill, 0, "")
		pdf.CellFormat(25, cellHeight, strconv.Itoa(orderTotalsByRoom[roomNumber].CheeseSlices), "LR", 0, "C", fill, 0, "")
		pdf.Ln(-1)

		fill = !fill
		pdf.CellFormat(35, cellHeight, "Pepperoni", "LR", 0, "C", fill, 0, "")
		pdf.CellFormat(25, cellHeight, strconv.Itoa(orderTotalsByRoom[roomNumber].PepperoniPizzas), "LR", 0, "C", fill, 0, "")
		pdf.CellFormat(25, cellHeight, strconv.Itoa(orderTotalsByRoom[roomNumber].PepperoniSlices), "LR", 0, "C", fill, 0, "")
		pdf.Ln(-1)

		fill = !fill
		pdf.CellFormat(35, cellHeight, "GF Cheese", "LR", 0, "C", fill, 0, "")
		pdf.CellFormat(25, cellHeight, strconv.Itoa(orderTotalsByRoom[roomNumber].GlutenFreeCheesePizzas), "LR", 0, "C", fill, 0, "")
		pdf.CellFormat(25, cellHeight, strconv.Itoa(orderTotalsByRoom[roomNumber].GlutenFreeCheeseSlices), "LR", 0, "C", fill, 0, "")
		pdf.Ln(-1)

		fill = !fill
		pdf.CellFormat(35, cellHeight, "GF Pepperoni", "LR", 0, "C", fill, 0, "")
		pdf.CellFormat(25, cellHeight, strconv.Itoa(orderTotalsByRoom[roomNumber].GlutenFreePepperoniPizzas), "LR", 0, "C", fill, 0, "")
		pdf.CellFormat(25, cellHeight, strconv.Itoa(orderTotalsByRoom[roomNumber].GlutenFreePepperoniSlices), "LR", 0, "C", fill, 0, "")
		pdf.Ln(-1)

		fill = !fill
		pdf.CellFormat(35, cellHeight, "DF Cheese", "LR", 0, "C", fill, 0, "")
		pdf.CellFormat(25, cellHeight, strconv.Itoa(orderTotalsByRoom[roomNumber].DairyFreeCheesePizzas), "LR", 0, "C", fill, 0, "")
		pdf.CellFormat(25, cellHeight, strconv.Itoa(orderTotalsByRoom[roomNumber].DairyFreeCheeseSlices), "LR", 0, "C", fill, 0, "")
		pdf.Ln(-1)

		pdf.CellFormat(85, 0, "", "T", 0, "", false, 0, "")
		pdf.Ln(5)

		// position drinks table beside previous table rather than below
		x += 10
		pdf.SetXY(x, y)

		pdf.SetFillColor(255, 0, 0)
		pdf.SetTextColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		pdf.SetLineWidth(.3)
		pdf.SetFont("", "B", 12)
		pdf.CellFormat(25, cellHeight, "Drinks", "1", 0, "C", true, 0, "")

		y += cellHeight
		pdf.SetXY(x, y)
		pdf.SetTextColor(0, 0, 0)
		pdf.SetFont("", "", 0)
		pdf.CellFormat(25, cellHeight, strconv.Itoa(orderTotalsByRoom[roomNumber].Drinks), "LR", 0, "C", false, 0, "")

		y += cellHeight
		pdf.SetXY(x, y)
		pdf.CellFormat(25, 0, "", "T", 0, "", false, 0, "")

	}

	var roomInfoByNumber = make(map[string]config.Room)
	for _, room := range roomInfo {
		roomInfoByNumber[room.Room] = room
	}

	orderTotalsPage()
	orderTotalsByRoomPage()

	for _, roomNumber := range roomNumbers {
		ordersForRoomPage(roomInfoByNumber, roomNumber)
	}

	fileName = strings.Replace(fileName, ".xlsx", ".pdf", 1)
	err := pdf.OutputFileAndClose(fileName)
	if err != nil {
		panic(err)
	}
}

func removeZero(n int) string {
	s := strconv.Itoa(n)
	if s == "0" {
		s = ""
	}
	return s
}
