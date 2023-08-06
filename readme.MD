cara untuk running project

pertama jalankan perintah go mod download

setelah itu ubah config.go dan ubah username password dan nama dbnya sesuai dengan environtment yang dipakai

dbURI := "<user>:<password>@tcp(localhost:3306)/<dbname>?charset=utf8&parseTime=True&loc=Local"

setelah itu jalankan go run main.go

setelah itu import postman collection untuk menjalankan api yang sudah terdokumentasi terima kasih :)
