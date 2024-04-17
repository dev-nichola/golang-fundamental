package main

type Luas interface {
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

func (pp PersegiPanjang) HitungLuas() int {
	return pp.Panjang * pp.Lebar
}

func main() {

}
