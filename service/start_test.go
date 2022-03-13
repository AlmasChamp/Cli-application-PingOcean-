package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {

	testTable := []struct {
		name     string
		args     []string
		expected error
	}{
		{
			name:     "invalidFlagUrls",
			args:     []string{"-u", "https:habr.com,https://leetcode.com,https://kolesa.kz,https://medium.com,https://github.com,https://www.youtube.com,https://www.baidu.com,https://www.vk.com,https://www.google.de,https://www.yandex.ru", "-search", "script"},
			expected: errors.New("ERROR: invalid flag, should be -urls"),
		},

		{
			name:     "invalidFlagSearch",
			args:     []string{"-urls", "https:habr.com,https://leetcode.com,https://kolesa.kz,https://medium.com,https://github.com,https://www.youtube.com,https://www.baidu.com,https://www.vk.com,https://www.google.de,https://www.yandex.ru", "-sarc", "script"},
			expected: errors.New("ERROR: invalid flag, should be -search"),
		},
		{
			name:     "not enough arguments",
			args:     []string{"-urls", "https:habr.com,https://leetcode.com,https://kolesa.kz,https://medium.com,https://github.com,https://www.youtube.com,https://www.baidu.com,https://www.vk.com,https://www.google.de,https://www.yandex.ru", "-search"},
			expected: errors.New("ERROR: not enough arguments. Example: -urls https://habr.com -search script"),
		},
		{
			name:     "many arguments",
			args:     []string{"-urls", "https:habr.com,https://leetcode.com,https://kolesa.kz,https://medium.com,https://github.com,https://www.youtube.com,https://www.baidu.com,https://www.vk.com,https://www.google.de,https://www.yandex.ru", "-search", "script", "too"},
			expected: errors.New("ERROR: more than 4 arguments. Example: -urls https://habr.com -search script"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			//Act
			err := Start(testCase.args)

			//Assert
			assert.Equal(t, testCase.expected, err)

		})
	}

}
