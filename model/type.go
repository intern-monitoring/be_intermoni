package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           	primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	FirstName  		string             `bson:"firstname,omitempty" json:"firstname,omitempty"`
	LastName  		string             `bson:"lastname,omitempty" json:"lastname,omitempty"`
	Email  			string             `bson:"email,omitempty" json:"email,omitempty"`
	Password        string         	   `bson:"password,omitempty" json:"password,omitempty"`
	Confirmpassword string         	   `bson:"confirmpass,omitempty" json:"confirmpass,omitempty"`
	Salt 			string			   `bson:"salt,omitempty" json:"salt,omitempty"`
	Role     		string			   `bson:"role,omitempty" json:"role,omitempty" `

}

type Lowongan struct {
	ID              		primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Logo            		string			   `bson:"logo,omitempty" json:"logo,omitempty"`
	Jabatan         		string			   `bson:"jabatan,omitempty" json:"jabatan,omitempty"`
	Perusahaan      		string			   `bson:"perusahaan,omitempty" json:"perusahaan,omitempty"`
	Lokasi          		string			   `bson:"lokasi,omitempty" json:"lokasi,omitempty"`
	CreatedAt       		string			   `bson:"createdat,omitempty" json:"createdat,omitempty"`
	DeskripsiPekerjaan  	string			   `bson:"deskripsipekerjaan,omitempty" json:"deskripsipekerjaan,omitempty"`
	InfoTambahanPekerjaan   string			   `bson:"infotambahanpekerjaan,omitempty" json:"infotambahanpekerjaan,omitempty"`
	TentangPerusahaan   	string			   `bson:"tentangperusahaan,omitempty" json:"tentangperusahaan,omitempty"`
	InfoTambahanPerusahaan  string			   `bson:"infotambahanperusahaan,omitempty" json:"infotambahanperusahaan,omitempty"`
}

type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type Payload struct {
	Id           	primitive.ObjectID `json:"id"`
	Exp 			time.Time 	 	   `json:"exp"`
	Iat 			time.Time 		   `json:"iat"`
	Nbf 			time.Time 		   `json:"nbf"`
	FirstName  		string             `json:"firstname"`
	LastName  		string             `json:"lastname"`
	Email  			string             `json:"email"`
	Role     		string			   `json:"role"`
}