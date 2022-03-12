
run:
	go run -race . -urls https://habr.com,https://leetcode.com,https://kolesa.kz,https://medium.com,https://github.com,https://www.youtube.com,https://www.baidu.com,https://www.vk.com,https://www.google.de,https://www.yandex.ru -search script

build :
	go build -o bin/main main.go