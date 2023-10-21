package Fatwa_1214038

import (
	"fmt"
	"testing"

	"github.com/intern-monitoring/be_intermoni/model"
	"github.com/intern-monitoring/be_intermoni/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var db = module.MongoConnect()

func TestGetUserFromEmail(t *testing.T) {
	email := "admin@gmail.com"
	hasil, err := module.GetUserFromEmail(email, db, "user")
	if err != nil {
		t.Errorf("Error TestGetUserFromEmail: %v", err)
	} else {
		fmt.Println(hasil)
	}
}

func TestInsertOneLowongan(t *testing.T) {
	var doc model.Lowongan
   doc.Logo = "https://fatwaff.github.io/bk-image/user/ford.jpg"
   doc.Jabatan = "Network Engineer"
   doc.Perusahaan = "Ford Company Etc"
   doc.Lokasi = "Bandung"
   doc.CreatedAt = "07-08-2023"
   doc.DeskripsiPekerjaan = "<div><ul><li>Mengurus administrasi bagian marketing</li><li>Membuat Sales Order,membuat Penawaran Harga</li><li>Menerima Purchase Order (PO) Customer</li><li>Membina hubungan baik antara Perusahaan dan Customer</li><li>Bisa bekerja secara akurat dan memperhatikan detail sehingga bisa memproses pesanan dengan cepat dan efisien</li><li>Jujur, pekerja keras,ulet,tekun,bertanggung jawab,punya komitmen yang tinggi, percaya diri, memiliki kemampuan komunikasi yang baik</li></ul></div>"
   doc.InfoTambahanPekerjaan = "<div><ul><li>Pengalaman 3 tahun kerja</li><li>Pegawai tetap</li></ul></div>"
   doc.TentangPerusahaan = "<div>Ford Motor Company (commonly known as Ford) is an American multinational automobile manufacturer headquartered in Dearborn, Michigan, United States. It was founded by Henry Ford and incorporated on June 16, 1903. The company sells automobiles and commercial vehicles under the Ford brand, and luxury cars under its Lincoln brand.</div>"
   doc.InfoTambahanPerusahaan = "<div><ul><li>1000-2000 Pekerja</li><li>Industri Manufaktur/Produksi</li></ul></div>"
   if doc.Logo == "" || doc.Jabatan == "" || doc.Perusahaan == "" || doc.Lokasi == "" || doc.CreatedAt == "" || doc.DeskripsiPekerjaan == "" || doc.InfoTambahanPekerjaan == "" || doc.TentangPerusahaan == "" || doc.InfoTambahanPerusahaan == "" {
	   t.Errorf("mohon untuk melengkapi data")
   } else {
	   insertedID, err := module.InsertOneDoc(db, "lowongan", doc)
	   if err != nil {
		   t.Errorf("Error inserting document: %v", err)
		   fmt.Println("Data tidak berhasil disimpan")
	   } else {
	   fmt.Println("Data berhasil disimpan dengan id :", insertedID.Hex())
	   }
   }
}

func TestGetAllDoc(t *testing.T) {
	var docs []model.Lowongan
	hasil := module.GetAllDocs(db, "lowongan", docs)
	fmt.Println(hasil)
}

func TestUpdateOneDoc(t *testing.T) {
 	var docs model.User
	id := "649063d3ad72e074286c61e8"
	objectId, _ := primitive.ObjectIDFromHex(id)
	docs.FirstName = "Aufah"
	docs.LastName = "Auliana"
	docs.Email = "aufa@gmail.com"
	docs.Password = "123456"
	if docs.FirstName == "" || docs.LastName == "" || docs.Email == "" || docs.Password == "" {
		t.Errorf("mohon untuk melengkapi data")
	} else {
		err := module.UpdateOneDoc(db, "user", objectId, docs)
		if err != nil {
			t.Errorf("Error inserting document: %v", err)
			fmt.Println("Data tidak berhasil diupdate")
		} else {
			fmt.Println("Data berhasil diupdate")
		}
	}
}

func TestGetLowonganFromID(t *testing.T){
	id := "64d0b1104255ba95ba588512"
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		t.Fatalf("error converting id to objectID: %v", err)
	}
	doc, err := module.GetLowonganFromID(objectId)
	if err != nil {
		t.Fatalf("error calling GetDocFromID: %v", err)
	}
	fmt.Println(doc)
}

func TestSignUpMahasiswa(t *testing.T) {
	var doc model.User
	doc.FirstName = "Dimas"
	doc.LastName = "Ardianto"
	doc.Email = "dimas@gmail.com"
	doc.Password = "fghjkliow"
	doc.Confirmpassword = "fghjkliow"
	insertedID, err := module.SignUpMahasiswa(db, "user", doc)
	if err != nil {
		t.Errorf("Error inserting document: %v", err)
	} else {
	fmt.Println("Data berhasil disimpan dengan id :", insertedID.Hex())
	}
}

func TestSignUpIndustri(t *testing.T) {
	var doc model.User
	doc.FirstName = "Dimas"
	doc.LastName = "Ardianto"
	doc.Email = "dimas@gmail.com"
	doc.Password = "fghjkliow"
	doc.Confirmpassword = "fghjkliow"
	insertedID, err := module.SignUpIndustri(db, "user", doc)
	if err != nil {
		t.Errorf("Error inserting document: %v", err)
	} else {
	fmt.Println("Data berhasil disimpan dengan id :", insertedID.Hex())
	}
}

func TestLogIn(t *testing.T) {
	var doc model.User
	doc.Email = "dimas@gmail.com"
	doc.Password = "fghjkliow"
	user, err := module.LogIn(db, "user", doc)
	if err != nil {
		t.Errorf("Error getting document: %v", err)
	} else {
		fmt.Println("Welcome bang:", user)
	}
}

// func TestGeneratePrivateKeyPaseto(t *testing.T) {
// 	privateKey, publicKey := module.GenerateKey()
// 	fmt.Println("ini private key :", privateKey)
// 	fmt.Println("ini public key :", publicKey)
// 	hasil, err := module.Encode("fatwaff@gmail.com", privateKey)
// 	fmt.Println("ini hasil :", hasil, err)
// }

func TestWatoken2(t *testing.T) {
	var user model.User
	privateKey, publicKey := module.GenerateKey()
	fmt.Println("privateKey : ", privateKey)
	fmt.Println("publicKey : ", publicKey)
	id := "649063d3ad72e074286c61e8"
	objectId, _ := primitive.ObjectIDFromHex(id)
	user.FirstName = "Fatwa"
	user.LastName = "Fatahillah"
	user.Email = "fax@gmail.com"
	user.Role = "pelamar"
	tokenstring, err := module.Encode(objectId, user.FirstName, user.LastName, user.Email, user.Role,  privateKey)
	if err != nil {
		t.Errorf("Error getting document: %v", err)
	} else {
		body, err := module.Decode(publicKey, tokenstring)
		fmt.Println("signed : ", tokenstring)
		fmt.Println("isi : ", body)
		if err != nil {
			t.Errorf("Error getting document: %v", err)
		} else {
			fmt.Println("Berhasil yey!")
		}
	}
}

func TestWatoken(t *testing.T) {
	body, err := module.Decode("3108501ddf2f9aa33f3a7c0e387339b131d8a4d22818feb511b0fe1bc3a16b36", "v4.public.eyJleHAiOiIyMDIzLTEwLTE5VDA4OjUyOjQzWiIsImlhdCI6IjIwMjMtMTAtMTlUMDY6NTI6NDNaIiwiaWQiOiJkaW1hc0BnbWFpbC5jb20iLCJuYmYiOiIyMDIzLTEwLTE5VDA2OjUyOjQzWiJ99kyEH6cp4FMg5wM4XdZYdCZFJCb2-OhP6JLaQ7imddJKTbLVt4F5QtlRpNXxCYYs_tQ-UhEpjhySz-i7W8mmBQ")
	fmt.Println("isi : ", body, err)
}