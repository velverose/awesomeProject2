package main

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
	"github.com/google/uuid"

	"github.com/AlihanSDev/lab4-variant11/pkg/printshop"
)

func main() {

	pages := 120
	copies := 10
	pricePerPage := 2.75
	discount := 15.0

	orderID := uuid.New().String()

	cost, err := printshop.OrderCost(pages, copies, pricePerPage)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	err = printshop.ApplyBulkDiscount(&cost, discount)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	report, err := printshop.FormatPrintReport(orderID, pages, copies, cost)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Цветной вывод
	color.Green("=== ОТЧЁТ О ЗАКАЗЕ ===")
	fmt.Println(report)

	// Форматированный вывод
	fmt.Printf("\nСтоимость (%.2f руб.)\n", cost)

	// humanize
	fmt.Println("Стоимость прописью:", humanize.Commaf(cost))
}
