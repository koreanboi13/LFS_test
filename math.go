package main

// Add returns sum of two integers
func Add(a, b int) int {
	return a + b
}

// Divide returns a divided by b
// Намеренная ошибка 4: нет проверки деления на ноль
func Divide(a, b int) int {
	return a / b
}
