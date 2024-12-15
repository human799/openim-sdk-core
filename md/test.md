## 注册

integration_test$ go run main.go -reg -u 100000

![1734262156609](image/test/1734262156609.png)


integration_test$ go run main.go -reg -u 1000000

![1734262319772](image/test/1734262319772.png)

数据库用户表

![1734262530967](image/test/1734262530967.png)



启动5万在线用户

msgtest/main$ go run main.go -o 50000 -u true

![1734262841296](image/test/1734262841296.png)

时间坚持2分钟以上,CPU会有突高

![1734263519371](image/test/1734263519371.png)


五千个用户基本没有问题

msgtest/main$ go run main.go -o 5000 -u true

![1734263843101](image/test/1734263843101.png)
