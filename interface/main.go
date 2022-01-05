package main

import "fmt"

type BangunDatar interface {
	HitungLuas() int
}

type Persegi struct {
	Sisi int
}

func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}
func main() {
	persegiPanjang := PersegiPanjang{Panjang: 6, Lebar: 5}
	luas := SeberapaLuas(persegiPanjang)
	fmt.Println("Luas: ", luas)

	persegi := Persegi{Sisi: 4}
	luas2 := SeberapaLuas(persegi)
	fmt.Println("Luas persegi: ", luas2)
}

func SeberapaLuas(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
}
