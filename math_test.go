package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Add(2, 3) = %d, want 5", result)
	}
}

// Намеренная ошибка 5: тест который падает
func TestDivide(t *testing.T) {
	result := Divide(10, 2)
	if result != 6 { // неправильное ожидаемое значение — должно быть 5
		t.Errorf("Divide(10, 2) = %d, want 6", result)
	}
}

// Намеренная ошибка 6: паника при делении на ноль
func TestDivideByZero(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic on divide by zero")
		}
	}()
	_ = Divide(1, 0)
}
