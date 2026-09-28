package main

// Пишите тесты в этом файле
import "testing"

func TestGenerateRandomElements(t *testing.T) {
	data := generateRandomElements(10)

	if len(data) != 10 {
		t.Errorf("ожидалась длина 10, получена %d", len(data))
	}

	emptyData := generateRandomElements(0)

	if len(emptyData) != 0 {
		t.Errorf("ожидался пустой слайс")
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{
			name:     "несколько элементов",
			data:     []int{1, 5, 3, 10, 2},
			expected: 10,
		},
		{
			name:     "один элемент",
			data:     []int{5},
			expected: 5,
		},
		{
			name:     "пустой слайс",
			data:     []int{},
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := maximum(test.data)

			if result != test.expected {
				t.Errorf("ожидалось %d, получено %d", test.expected, result)
			}
		})
	}
}