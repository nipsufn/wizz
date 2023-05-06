module github.com/nipsufn/wizz

replace (
    log => github.com/sirupsen/logrus
    github.com/FerdinaKusumah/wizz => github.com/nipsufn/wizz
    github.com/FerdinaKusumah/wizz/connection => github.com/nipsufn/wizz/connection
    github.com/FerdinaKusumah/wizz/models => github.com/nipsufn/wizz/models
    github.com/FerdinaKusumah/wizz/utils => github.com/nipsufn/wizz/utils
)

go 1.20
