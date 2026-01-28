package main

import "fmt"

func main() {
	var name1 string = "Muhammad"
	var name2 string = "Muhammad"

	var result bool = name1 == name2
	fmt.Println(result) // true

	var value1 int = 100
	var value2 int = 200
	var result2 bool = value1 < value2
	fmt.Println(result2) // true

	var value3 int = 100
	var value4 int = 100
	var result3 bool = value3 >= value4
	fmt.Println(result3) // true

	// logic AND (&&)
	var nilaiAkhir = 90
	var absensi = 80
	var lulusNilaiAkhir = nilaiAkhir >= 80
	var lulusAbsensi = absensi >= 80
	var lulus = lulusNilaiAkhir && lulusAbsensi
	fmt.Println(lulus) // true

	// logic OR (||)
	var nilaiUjianAkhir = 70
	var absensiKedua = 90
	var lulusNilaiUjianAkhir = nilaiUjianAkhir >= 80
	var lulusAbsensiKedua = absensiKedua >= 80
	var lulusKedua = lulusNilaiUjianAkhir || lulusAbsensiKedua
	fmt.Println(lulusKedua) // true

	// logic NOT (!)
	var nilaiUjianAkhirKetiga = 70
	var lulusNilaiUjianAkhirKetiga = nilaiUjianAkhirKetiga >= 80
	var tidakLulus = !lulusNilaiUjianAkhirKetiga
	fmt.Println(tidakLulus) // true
}