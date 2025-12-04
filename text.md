go mod init (nama folder) merupakan perintah untuk membuat file go.mod yang menandai folder sebagai module go. 

module ini pneting, yang nantikan go akan mencatat semua depedencies atau packages atau library yang dipakai dengan go.mod

go mod init dijalankan pertama kali karena tanpa go mod, go tidak tahu folder ini adalah sebuah module jika melakukan import library pihak ke 3 atau trith party (jwt), golang akan menolak menambahkan depedencies sebelum ada nya go mod.


