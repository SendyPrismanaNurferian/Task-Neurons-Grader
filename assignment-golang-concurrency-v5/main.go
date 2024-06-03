package main

import "fmt"

type RowData struct {
	RankWebsite int
	Domain      string
	TLD         string
	IDN_TLD     string
	Valid       bool
	RefIPs      int
}

func GetTLD(domain string) (TLD string, IDN_TLD string) {
	var ListIDN_TLD = map[string]string{
		".com": ".co.id",
		".org": ".org.id",
		".gov": ".go.id",
	}

	for i := len(domain) - 1; i >= 0; i-- {
		if domain[i] == '.' {
			TLD = domain[i:]
			break
		}
	}

	if _, ok := ListIDN_TLD[TLD]; ok {
		return TLD, ListIDN_TLD[TLD]
	} else {
		return TLD, TLD
	}
}

func ProcessGetTLD(website RowData, ch chan RowData, chErr chan error) {
	if website.Domain == "" {
		chErr <- fmt.Errorf("domain name is empty")
		return
	}

	if !website.Valid {
		chErr <- fmt.Errorf("domain not valid")
		return
	}

	if website.RefIPs == -1 {
		chErr <- fmt.Errorf("domain RefIPs not valid")
		return
	}

	TLD, IDN_TLD := GetTLD(website.Domain)
	website.TLD = TLD
	website.IDN_TLD = IDN_TLD

	ch <- website
}

/*
Fungsi ini menerima parameter website bertipe RowData, yang berisi informasi tentang sebuah website.
Pertama-tama, fungsi ini melakukan pengecekan terhadap data website:
Jika domain kosong, maka kirimkan error melalui channel chErr dengan pesan "domain name is empty".
Jika Valid bernilai false, kirimkan error dengan pesan "domain not valid".
Jika RefIPs bernilai -1, kirimkan error dengan pesan "domain RefIPs not valid".
Jika semua pengecekan berhasil, fungsi akan menggunakan fungsi GetTLD untuk mendapatkan TLD dan IDN_TLD dari Domain.
Setelah itu, TLD dan IDN_TLD disimpan di dalam website, dan website dikirimkan melalui channel ch.
*/

// Gunakan variable ini sebagai goroutine di fungsi FilterAndGetDomain
var FuncProcessGetTLD = ProcessGetTLD

func FilterAndFillData(TLD string, data []RowData) ([]RowData, error) {
	ch := make(chan RowData, len(data))
	errCh := make(chan error)

	for _, website := range data {
		go FuncProcessGetTLD(website, ch, errCh)
	}

	var filteredData []RowData

	for i := 0; i < len(data); i++ {
		select {
		case website := <-ch:
			if website.TLD == TLD {
				filteredData = append(filteredData, website)
			}
		case err := <-errCh:
			return nil, err
		}
	}

	return filteredData, nil
}

/*
Fungsi ini menerima parameter TLD yang menentukan TLD yang akan difilter, dan data yang merupakan slice dari RowData.
Pertama-tama, kita membuat dua channel, ch untuk menerima data dari goroutine ProcessGetTLD dan errCh untuk menerima error.
Selanjutnya, kita menjalankan goroutine FuncProcessGetTLD untuk setiap elemen data. Setiap goroutine akan memproses sebuah website.
Setelah itu, kita melakukan iterasi sebanyak len(data), dan dengan menggunakan select, kita menunggu untuk menerima data atau error dari channel.
Jika kita menerima data website dari channel ch, kita cek apakah TLD-nya sama dengan TLD yang ditentukan. Jika ya, kita masukkan ke dalam filteredData.
Jika kita menerima error dari channel errCh, kita langsung kembalikan error tersebut.
Setelah selesai, kita kembalikan filteredData dan nil karena tidak ada error.
Dengan demikian, kita telah menyelesaikan implementasi untuk fungsi ProcessGetTLD dan FilterAndFillData. Kedua fungsi ini akan bekerja secara konkuren untuk mengisi dan memproses data sesuai spesifikasi yang diberikan.
*/

// gunakan untuk melakukan debugging
func main() {
	rows, err := FilterAndFillData(".com", []RowData{
		{1, "google.com", "", "", true, 100},
		{2, "facebook.com", "", "", true, 100},
		{3, "golang.org", "", "", true, 100},
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rows)
}
