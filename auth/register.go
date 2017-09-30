package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kaaryasthan/kaaryasthan/db"
	"github.com/kaaryasthan/kaaryasthan/jsonapi"
	"golang.org/x/crypto/scrypt"
)

func (obj *Schema) register() error {
	salt := randomSalt()
	password, err := scrypt.Key([]byte(obj.Password), salt, 16384, 8, 1, 32)
	if err != nil {
		return err
	}
	_, err = db.DB.Exec(`INSERT INTO "members"
		(username, name, email, password, salt)
		VALUES ($1, $2, $3, $4, $5)`,
		obj.Username,
		obj.Name,
		obj.Email,
		password,
		salt)
	if err != nil {
		return err
	}
	return nil
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	payload := make(map[string]jsonapi.Data)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		log.Println("Unable to decode body: ", err)
		return
	}
	s := New(payload["data"])
	err = s.register()
	if err != nil {
		log.Println("Unable save data: ", err)
		return
	}
	tmpData := payload["data"]
	tmpData.ID = s.Username
	delete(tmpData.Attributes, "password")
	payload["data"] = tmpData
	b, err := json.Marshal(payload)
	if err != nil {
		log.Println("Unable marshal data: ", err)
		return
	}
	w.Write(b)
}
