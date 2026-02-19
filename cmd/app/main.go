package main

import (
	"fmt"

	"github.com/velverose/awesomeProject2/pkg/printshop"

	"github.com/google/uuid"
)

func main() {
	pages := 120
	copies := 10
	pricePerPage := 3.75

	// Генерация ID заказа (внешний пакет GitHub)
	orderID := uuid.New().String()

	// 1. Расчёт стоимости заказа
	totalCost, err := printshop.OrderCost(pages, copies, pricePerPage)
	if err != nil {
		fmt.Printf("Ошибка расчёта: %v\n", err)
		return
	}

	fmt.Printf("Стоимость без скидки: %10.2f руб.\n", totalCost)

	// 2. Применение скидки 15%
	err = printshop.ApplyBulkDiscount(&totalCost, 15)
	if err != nil {
		fmt.Printf("Ошибка применения скидки: %v\n", err)
		return
	}

	fmt.Printf("Стоимость со скидкой:  %10.2f руб.\n", totalCost)

	// 3. Формирование отчёта
	report, err := printshop.FormatPrintReport(orderID, pages, copies, totalCost)
	if err != nil {
		fmt.Printf("Ошибка формирования отчёта: %v\n", err)
		return
	}

	fmt.Println("\n=== ОТЧЁТ ===")
	fmt.Println(report)
}
