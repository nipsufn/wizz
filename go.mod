module github.com/nipsufn/wizz

require (
	github.com/FerdinaKusumah/wizz v0.0.0-20210831035532-213ef7abcd4e
	github.com/sirupsen/logrus v1.9.0
)

require golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect

replace github.com/FerdinaKusumah/wizz => ./

go 1.20
