module DatabaseGolangLearningProject

go 1.16

replace (
	github.com/go-sql-driver/mysql v1.6.0 => ./localthirdparty/mysql@v1.6.0
	github.com/mattn/go-sqlite3 v1.14.8 => ./localthirdparty/go-sqlite3@v1.14.8
)

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/mattn/go-sqlite3 v1.14.8
)
