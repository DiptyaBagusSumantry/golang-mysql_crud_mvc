package pasiencontroller

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/jeypc/go-crud/entities"
	"github.com/jeypc/go-crud/models"
)

var pasienModel = models.NewPasienModel()


func Index(response http.ResponseWriter, request *http.Request){
	fmt.Println("Its workss")
	temp, err := template.ParseFiles("views/pasien/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, nil)
}
func Add(response http.ResponseWriter, request *http.Request){
	if request.Method == http.MethodGet{
		temp, err := template.ParseFiles("views/pasien/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost{
		request.ParseForm()

		var pasien entities.Pasien 
		pasien.NamaLengkap = request.Form.Get("nama_lengkap")
		pasien.NIK = request.Form.Get("nik")
		pasien.JenisKelamin = request.Form.Get("jenis_kelamin")
		pasien.TempatLahir = request.Form.Get("tempat_lahir")
		pasien.TanggalLahir = request.Form.Get("tanggal_lahir")
		pasien.Alamat = request.Form.Get("alamat")
		pasien.NoHp = request.Form.Get("nomer_hp")

		// fmt.Println(pasien)
		pasienModel.Create(pasien)

		data := map[string]interface{}{
			"pesan": "Data Berhhasil di Simpan",
		}
		temp, _ := template.ParseFiles("views/pasien/add.html")
		temp.Execute(response, data)
	}
	
}
func Edit(response http.ResponseWriter, request *http.Request){

}
func Delete(response http.ResponseWriter, request *http.Request){

}