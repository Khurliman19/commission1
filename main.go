package main

import (
	"commission1/internal/commission"
	"commission1/internal/utils"
	"fmt"
)

func main() {
	var sName, sSurname string
	var rName, rSurname string
	var card string
	var amount int64
	var isAlifInput int

	fmt.Print("Имя отправителя: ")
	fmt.Scan(&sName)
	fmt.Print("Фамилия отправителя: ")
	fmt.Scan(&sSurname)
	fmt.Print("Имя получателя: ")
	fmt.Scan(&rName)
	fmt.Print("Фамилия получателя: ")
	fmt.Scan(&rSurname)
	fmt.Print("Номер карты (16 цифр): ")
	fmt.Scan(&card)
	fmt.Print("Введите сумму: ")
	fmt.Scan(&amount)
	fmt.Print("Alif карта? (1-да/0-нет): ")
	fmt.Scan(&isAlifInput)

	isAlif := isAlifInput == 1

	if !commission.Validate(amount) {
		fmt.Println("Ошибка: сумма вне диапазона")
		return
	}

	comm := commission.Calculate(amount, isAlif)
	total := amount + comm

	txID := utils.GenerateTransactionID()
	date := utils.CurrentDateTime()
	maskedCard := utils.MaskCard(card)

	sender := utils.ToUpperFullName(sName, sSurname)
	receiver := utils.ToUpperFullName(rName, rSurname)

	fmt.Println("========== Электронный чек ==========")
	fmt.Println("Отправитель:", sender)
	fmt.Println("Получатель:", receiver)
	fmt.Println("Номер транзакции:", txID)
	fmt.Println("Счёт зачисления:", maskedCard)
	fmt.Println("Дата и время:", date)
	fmt.Println("Сумма:", "%d сум", amount)
	fmt.Println("Комиссия:", "%d сум", comm)
	fmt.Println("Итого:", "%d сум", total)
	fmt.Println()
	fmt.Println("AO ALIF TECH • Лицензия ЦБ РУз Nº 000010")
	fmt.Println("Статус:", "Исполнено")
	fmt.Println("Спасибо за использование Alif mobi")
	fmt.Println("=====================================")
}
