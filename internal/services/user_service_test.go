package services

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTopicNameByID(t *testing.T) {

	testCases := []struct {
		a, b, expected int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 1, 0},
		{-1, -1, -2},
	}
	for _, tc := range testCases {
		t.Run(
			// Создаем имя для теста на основе входных значений
			// для удобства чтения отчетов при запуске тестов
			fmt.Sprintf("Add(%d, %d)", tc.a, tc.b),
			func(t *testing.T) {
				result := Add(tc.a, tc.b)
				assert.Equal(t, tc.expected, result)
			},
		)
	}

}
