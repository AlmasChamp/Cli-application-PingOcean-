FROM golang:1.17.6
RUN mkdir /app
WORKDIR /app
ADD . /app
RUN go build -o main .
CMD ["/app/main", "-urls", "https://habr.com,https://leetcode.com,https://kolesa.kz,https://medium.com,https://github.com,https://www.youtube.com,https://www.baidu.com,https://www.vk.com,https://www.google.de,https://www.yandex.ru", "-search", "script"]