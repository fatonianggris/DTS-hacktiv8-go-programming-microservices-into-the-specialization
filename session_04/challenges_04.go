package main

import (
	"fmt"
	"strings"
)

type Person interface {
	PrintPerson() string
}

type ComponentPerson struct {
	nama         string
	alamat       string
	jeniskelamin string
	pekerjaan    string
	warganegara  string
	alasan       string
}

func (comp ComponentPerson) PrintPerson() string {

	var print string

	print = "Nama: " + strings.Title(strings.ToLower(comp.nama)) + "\n" +
		"Alamat: " + strings.Title(strings.ToLower(comp.alamat)) + "\n" +
		"Jenis Kelamin: " + comp.jeniskelamin + "\n" +
		"Pekerjaan: " + strings.ToUpper(comp.pekerjaan) + "\n" +
		"Warga Negara: " + comp.warganegara + "\n" +
		"Alasan Memilih Kelas Golang: " + strings.ToLower(comp.alamat)

	return print
}

func main() {
	var nm Person = ComponentPerson{nama: "M FATONI ANGGRIS", alamat: "jln. Sumbersari no 4", jeniskelamin: "Laki-laki", pekerjaan: "Pelajar", warganegara: "WNI", alasan: "mencari skill dan pengalaman"}
	fmt.Print(nm)

}
