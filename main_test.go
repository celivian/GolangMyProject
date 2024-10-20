package main

import (
	"testing"
)

func TestGetUTFLength(t *testing.T) {
	// набор тестов
	cases := []struct {
		// имя теста
		name string
		// значения на вход тестируемой функции
		value []byte
		// желаемый результат
		want int
		err  error
	}{
		// тестовые данные № 1
		{
			name:  "test1",
			value: []byte{1, 2, 3},
			want:  3,
			err:   nil,
		},
		// тестовые данные № 2
		{
			name:  "test2",
			value: []byte{0xff},
			want:  0,
			err:   ErrInvalidUTF8,
		},
		{
			name:  "test3",
			value: []byte{},
			want:  0,
			err:   nil,
		},
	}
	// перебор всех тестов
	for _, tc := range cases {
		tc := tc
		// запуск отдельного теста
		t.Run(tc.name, func(t *testing.T) {
			// тестируем функцию Sum
			got, err := GetUTFLength(tc.value)
			// проверим полученное значение
			if got != tc.want {
				t.Fatalf("GetUTFLength() got = %v, expected = %v", got, tc.want)
			}
			if err != tc.err {
				t.Errorf("GetUTFLength() error = %v, expected = %v", err, tc.err)
			}
		})
	}
}
