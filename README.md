# citra-digital

### 0.
Buat fungsi sederhana untuk menghitung jumlah huruf hidup dan huruf mati dari suatu kalimat. Contoh:
<br>&nbsp;&nbsp;&nbsp; input: "omama"   
<br>&nbsp;&nbsp;&nbsp; hasil: "huruf mati: 2, huruf hidup: 2"  (huruf hidup yaitu o dan a. a tidak perlu muncul 2x)

### 1. 
Buat fungsi sederhana untuk mengurutkan abjad dari suatu kalimat, dengan memisahkan huruf hidup dan huruf mati. Contoh:
   <br>&nbsp;&nbsp;&nbsp; input: "omama"   hasil: "aaomm" ('a', dan 'o' sebagai huruf hidup dan 'm' sebagai huruf mati)
   <br>&nbsp;&nbsp;&nbsp; input: "osama"   hasil: "aaoms"

### 2. Webservice

untuk web service bisa menggunakan endpoint

* localhost:8080/sortVowels

dengan contoh body 

{
	"Text": "ueemamipop"
}
    
    
### 3. Docker for no 2

* `docker build -t cdigital .`
* `docker run -p 8080:8080 -t cdigital`

gunakan localhost atau ip 192.168.99.100 jika menggunakan docker toolbox <br>
jika terdapat masalah di dockerfile, bisa menggunakan dockerfile2 dengan cara mengganti nama menjadi dockerfile
