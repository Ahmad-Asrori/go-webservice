### Golang net/http Flexible Architecture

jika kita mendevelop aplikasi berbasis web di golang dengan tidak menggunakan framework, 
maka akan ada banyak fleksibilitas yang bisa kita peroleh. untuk mengingatkan kembali bahwa net/http mempunyai arsitektur 
seperti yang terlihat di gambar berikut : 

![](Screenshot_1.png)

kode standar dalam membuat aplikasi web di golang
```text
    //definisikan servernya mau dijalankan diport berapa
    server := http.Server{
        Addr: ":3000",
    }
    
    //multiplexer yang digunakan (serve mux adalah multiplexer bawaan package net/http)
    router := http.NewServeMux()

    //handler
    router.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
        writer.Write([]byte("hello world"))
    })
    
    //pasang multiplexer ke server
    server.Handler = router
    
    //jalankan server
    server.ListenAndServe()
```

__CASE 1__ : multiplexer bawaan golang tidak mempunyai fitur bawaan untuk menghandle URL dengan path variable
atau dengan url pattern disertai dengan regex sehingga kita harus melakukan string operation yang cukup 
rumit kepada URL.

contoh URL

```text
/products/{key}
/articles/{category}/{id:[0-9]}
``` 

__Solusi__ : ganti multiplexer bawaan golang ke third party library lain.
seperti [mux](https://github.com/gorilla/mux#examples), [httprouter](https://github.com/julienschmidt/httprouter)

contoh kode jika kita menggunakan mux sebagai multiplexer. multiplexer mux membawa fitur yang lebih banyak
daripada multiplexer bawaaan package net/http oleh karena itu library ini menjadi library multiplexer yang paling
banyak digunakan. jika yang dicari adalah fitur, pilih mux. jika yang dicari adalah performa, pilih httprouter. 

sedikit contoh kode program 
```text
    //definisikan servernya mau dijalankan diport berapa
    server := http.Server{
        Addr: ":3000",
    }

    //multiplexer yang digunakan (dalam hal ini kita menggunakan mux)
    router := mux.NewRouter()
    
    //handler
    router.HandleFunc("/article/{month}", func(writer http.ResponseWriter, request *http.Request) {
        key := mux.Vars(request)
        writer.Write([]byte("hello " + key["month"]))
    })

    //pasang multiplexer ke server
    server.Handler = router

    //jalankan server
    server.ListenAndServe()
``` 

__CASE 2__ : Jika performa dari server bawaan package net/http dirasa kurang memuaskan.
 
__Solusi__ : Kita bisa menggantinya dengan [fasthttp](https://github.com/valyala/fasthttp) server

sedikit contoh kode program 

```text
    router := fasthttprouter.New()

    router.GET("/hello", func(ctx *fasthttp.RequestCtx) {
	    ctx.Write([]byte("hello world"))
    })

    fasthttp.ListenAndServe(":3000", router.Handler)
```

__CASE 3__ : Bagaimana caranya agar dapat server dapat mensuppport penggunaan protokol HTTP2 ?

__Solusi__ : Protokol HTTP2 bisa digunakan jika TLS/SSL dipasang.

sedikit contoh program
```text
    //definisikan servernya mau dijalankan di port berapa
    server := http.Server{Addr: ":3000"}
    serverHttp2 := http2.Server{}

    // multiplexer yang digunakan (dalam hal ini kita menggunakan httprouter)
    router := httprouter.New()

    //handler
    router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
        writer.Write([]byte("hello world"))
    })

    //tambahan support http2
    http2.ConfigureServer(&server, &serverHttp2)
    
    //pasang multiplexer ke server
    server.Handler = router

    //jalankan server
    server.ListenAndServeTLS("cert.pem", "key.pem")
```

__Q&A__

Q: Apakah kita bisa memakai multiplexer mux di fasthttp server?

A: Tidak, karena arsitekturnya sudah beda. 

Q: Mengapa framework seperti gin dan echo hanya mempunyai satu parameter
masukan di handler

```text
    router.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })

    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello, World!",
        })
    })
```

A : Back to basic, parameter tersebut adalah gabungan dari 2 parameter
dasar (ResponseWriter & Request) yang ada di package net/http. kemudian mereka
tambahkan fungsi fungsi untuk memudahkan pembuatan middleware, 
data binding, data rendering dll.

Q : fitur apakah yang banyak membuat orang tertarik untuk menggunakan framework

A : middleware. middleware terletak diantara multiplexer dan handler, biasanya middleware
digunakan untuk authentifikasi informasi (token, session. dll), mencatat request yang masuk (logging) dan lain sebagainya.
didalam framework terdapat banyak middleware yang siap digunakan.
namun jika ada kebutuhkan yang berbeda dan framework belum menyediakan middlewarenya, maka kita harus 
bisa untuk mengimplementasikan sendiri. Jika dilihat dari sudut pandang yang lain kita akan menemukan
bahwa middleware sebenarnya adalah sebuah library (bisa dari third party library diluar framework) yang dijalankan sesuai fungsinya 
dan ditempatkan di didalam sebuah fungsi, yang dimana fungsi tersebut di jalankan sebelum handler.

contoh : middleware jwt di framework echo, middleware tersebut menggunakan library [jwt-go](https://github.com/dgrijalva/jwt-go), middleware casbin juga demikian.
di project native, saya memakai library jwt-go dan casbin, workflownya setiap request yang masuk akan diperiksa JWT tokennya kemudian diperiksa otorisasinya apakah request yang masuk 
boleh mengakses resource (URL Endpoint) menggunakan HTTP Method yang diizinkan.

NB : istilah multiplexer hampir sama dengan istilah dispatcher servlet di spring. saya lebih suka menyebutnya sebagai router