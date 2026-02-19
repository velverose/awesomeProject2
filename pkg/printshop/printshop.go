// Package printshop предоставляет функции для расчёта стоимости
// заказов в типографии, применения скидок и формирования отчётов.
package printshop

import (
	"fmt"
)

// PageCost вычисляет стоимость печати одного набора страниц.
// Возвращает общую стоимость или ошибку при некорректных данных.
func PageCost(pages int, pricePerPage float64) (float64, error) {
	if pages <= 0 {
		return 0, fmt.Errorf("количество страниц должно быть больше 0")
	}
	if pricePerPage <= 0 {
		return 0, fmt.Errorf("цена за страницу должна быть больше 0")
	}

	return float64(pages) * pricePerPage, nil
}

// OrderCost вычисляет стоимость всего заказа
// (страницы × количество копий × цена за страницу).
func OrderCost(pages, copies int, pricePerPage float64) (float64, error) {
	if pages <= 0 {
		return 0, fmt.Errorf("количество страниц должно быть больше 0")
	}
	if copies <= 0 {
		return 0, fmt.Errorf("количество копий должно быть больше 0")
	}
	if pricePerPage <= 0 {
		return 0, fmt.Errorf("цена за страницу должна быть больше 0")
	}

	total := float64(pages*copies) * pricePerPage
	return total, nil
}

// ApplyBulkDiscount применяет скидку к стоимости заказа.
// Изменяет значение по указателю.
// percent должен быть в диапазоне (0-100).
func ApplyBulkDiscount(cost *float64, percent float64) error {
	if cost == nil {
		return fmt.Errorf("указатель на стоимость не может быть nil")
	}
	if percent <= 0 || percent >= 100 {
		return fmt.Errorf("процент скидки должен быть в диапазоне 0-100")
	}

	*cost = *cost * (1 - percent/100)
	return nil
}

// FormatPrintReport формирует строку отчёта по заказу.
func FormatPrintReport(orderID string, pages, copies int, cost float64) (string, error) {
	if orderID == "" {
		return "", fmt.Errorf("ID заказа не может быть пустым")
	}
	if pages <= 0 || copies <= 0 {
		return "", fmt.Errorf("некорректные данные заказа")
	}
	if cost < 0 {
		return "", fmt.Errorf("стоимость не может быть отрицательной")
	}

	report := fmt.Sprintf(
		"Заказ №%s\nСтраниц: %d\nКопий: %d\nИтоговая стоимость: %.2f руб.\n",
		orderID,
		pages,
		copies,
		cost,
	)

	return report, nil
}
