source code: https://gitlab.com/codingthailand/gobasic
		   : https://gitlab.com/codingthailand/gin-backend-api

slide :
	https://drive.google.com/drive/folders/1-uY9tTeEwCBOUGnI1J3PzmHLUzpiGGu4

***************** DAY 1*****************
***************** set up
gin web framework 
gorm ต่อ db postgresSQL

คำสั่ง command
	go env ดูว่า go อยู่ path ไหนกันแน่ในเครื่อง
windows 
	GOTOOLDIR=C:\Program Files\Go\pkg\tool\windows_amd64

	go help
		-install lib ใช้ go get
		-go clean เพื่อ clear cach
		
vscode extension
	-install Go - by To team at google
ไปที่ menu View -> command Palettle
		พิมพ์ Go:Install/Update Tools
	
	แล้วก็ install all เพื่อ install เครื่องมือสำหรับ go จนถึง
	All tools successfully installed. You are ready to Go :).
	สามารถลงได้เรื่อยๆ
	
Link สำหรับ อธิบาย extension go
https://github.com/golang/vscode-go/blob/HEAD/docs/tools.md

*******************
Start
	https://golang.org/ref/spec
	https://golang.org/doc/effective_go 
	
1.create project

	Module คือ collection ของ go package คล้ายๆ workspace เป็น collection ของ source file
		ใน module คือหลายๆ package -> source code คล้ายๆ java
	
		1 module คือ 1 project

	1.1 สร้าง folder : D:\jworkspace\git\gobasic
		-cmd : go mod init <ชื่อ repository ห้ามซ้ำกันทั่วโลก , domain name ขององค์กรณ์ก็ได้>
		go mod init github.com/frankenstyll/gobasic
		
		*github.com/frankenstyll/gobasic สามารถเปลี่ยนได้ตลอด
		
		
		จะได้ go.mod เทียบเท่า package.json ถ้าใช้ nodejs
	
	1.2 create new file main.go
		package main
		func main() {
			fmt.Println("Hello World!")
		}
		ภาษา go ค่อนข้าง specification เพราะ compiler จะปิดด้วย semi-colon ให้

		link standard package เพื่อเรียกใช้คล้ายๆ lib อยากใช้อะไรก็ import
		https://pkg.go.dev/std
		
		พิมพ์ fmt.Println() คือคำสั่ง print ตัว go จะ auto import -> import "fmt" มาให้
		
	1.3 compile 
		-cmd go run main.go หรือ อาจจะ build ก่อน 
			หรือจะรันด้วย go run . แล้วมันจะเข้ามารันผ่าน main.go เป็น funcion แรก
		-cmd go build . ถ้าเป็น windows จะได้ file .exe มา ที่นี้จะได้ gobasic.exe
		
	การ comment ก็คล้ายๆทั่วๆไป แล้วกด save ตัว tools จะ remove import ให้
	
****
	เราได้ package gobasic มาเป็นของเราแล้ว
	
	https://pkg.go.dev/std ศูนย์รวม package 
2. เรียกใช้ lib ภายนอก
	2.1 https://pkg.go.dev/rsc.io/quote/v3 ใช้ quote 
		ด้วยคำสั่ง go get rsc.io/quote/v3 คือการติดตั้ง lib จาก internet 
		lib ตัวนี้จะไปโผล่ที่ go.mod ส่วน go.sum ไม่ต้องไปเตะก็ได้
	2.2 เรียกใช้ quote/v3 ที่ main
		fmt.Println(quote.HelloV3())
		ระบบจะ auto import lib มาให้
		
	2.3 ถ้ามี problem cmd: go mod tidy เป็นคำสั่ง clear warning และลบ unuse module
		ก่อน build prod ก็ใช้คำสั่ง go mod tidy
	
3. variables
	เป็นภาษาที่มี type ทั้ง int, float64, 32, bool, string
	3.1 สร้าง package ใหม่ เหมือนการส้าง folder ใหม่
		ให้เป็นตัวเล็กหมด ข้างใน package ก็จะมี source code  variables.go
	3.2
	
4 function 
	ตัวแปรถ้าเป็นตัวพิมพ์ใหญ่ จะสามารถ เรียกใช้จาก package อื่นได้ คล้ายๆ static
	4.1 // func add(x, y int) ถ้า x, y มีค่าเท่ากันก็ประกาศแบบนี้ได้
		func add(x int, y int) int {
	4.2 go สามารถ retunr 2 ค่าออกไปได้	
		func converter(title, email string) (string, string) {
			
5. ***************************คำสั่ง go mod vendor เพื่อ copy lib ใน paht ที่ติดตั้งมาไว้ใน project
 go mod download เอาไว้ download lib กรณีเอา project เพื่อนมา คล้ายๆ maven update หรือ npm install

6.struct type  คล้ายๆกับ class model bean
	ตัวแรกพิมพ์เล็กคือ private ตัวแรกพิมพ์ใหญ่คือ public
7.pointer สำคัญ******************************************
	
	ตัวแปร ที่เก็บ memory address ของตัวแปรอื่น


***************** DAY 2*****************
1. connect db by pgadmin https://www.pgadmin.org/download/

	username = doadmin
	password = ug7WgnEElsM1mxFz
	host = db-postgres-do-user-998710-0.b.db.ondigitalocean.com
	port = 25060
	database = defaultdb
	sslmode = require
	
2. gin web framework - RESTFull api
	2.1 https://gin-gonic.com/
	2.2 https://gin-gonic.com/docs/quickstart/
	2.3 install cmd: go get -u github.com/gin-gonic/gin  เพื่อ install gin webframework

	2.4 copy : เพื่อลองรัน api localhost 
		func main() {
			r := gin.Default()
			r.GET("/ping", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "pong",
				})
			})
			r.Run() // listen and serve on 0.0.0.0:8080
		}
	2.5 go run . ด้วย http://localhost:8080/
		หรือเปลี่ยน port ด้วยคำสั่ง r.Run(":3000")
	2.6 ลง live-reloading เพื่อให้ gin webframework รันไปต่อโดยไม่ต้อง stop
		cmd: go get github.com/codegangsta/gin
		รันด้วย cmd: gin โปแกรมจะรันด้วย port 3000 (Listening on port 3000) เหมือนถูก ontop ไปอีกที
	เปลี่ยน port โดยใช้คำสั่ง
		gin -a 5000 -p 3000 , คือ port ที่เราต้องการคือ -a 5000 และ port ของ gin คือ -p 3000
		โดยสามารถเข้าได้ทั้ง 2 port
		
3. ลง go environment เพื่อให้มันจัดการพวก config ต่างๆเช่น database, username 
	https://github.com/joho/godotenv
	cmd: go get github.com/joho/godotenv
	3.1 สร้างไฟล์ .env ไว้นอกสุด



