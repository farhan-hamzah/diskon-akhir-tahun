package main
import "fmt"
func akhirTahun(totalBelanja float64, member bool){
	var belanja float64
	var diskon, cashback bool
	belanja = totalBelanja
	diskon = false
	cashback = false
	if totalBelanja >= 200000 && member == true{
		belanja -= 75000
		cashback = true
	}else if totalBelanja >= 100000{
		belanja = totalBelanja*(0.10)
		belanja = belanja-totalBelanja
		diskon = true
	}
	fmt.Println("kartu?", member)
	fmt.Println("Diskon?", diskon)
	fmt.Println("Cashback?", cashback)
	fmt.Println("Rp. ", belanja)
}
func main(){
	var totalBelanja float64
	var mem bool
	fmt.Scan(&totalBelanja, &mem)
	akhirTahun(totalBelanja, mem)
}