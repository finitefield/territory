package housemanager

import (
    "fmt"
    "net/http"
    "html/template"
    "time"
    "encoding/json"
    "log"
)

func init() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/api/getHouses", getHousesHandler)
    http.HandleFunc("/api/getHouse", getHouseHandler)
    http.HandleFunc("/api/updateHouse", updateHouseHandler)
    http.HandleFunc("/api/removeHouse", removeHouseHandler)
    http.HandleFunc("/api/getFolders", getFoldersHandler)
    http.HandleFunc("/api/getFolder", getFolderHandler)
    http.HandleFunc("/api/addTag", addTagHandler)
    http.HandleFunc("/api/removeTag", removeTagHandler)
    http.HandleFunc("/api/getTowns", getTownsHandler)
    http.HandleFunc("/api/getTown", getTownHandler)
    http.HandleFunc("/api/removeTown", removeTownHandler)
    http.HandleFunc("/api/addTown", addTownHandler)
    http.HandleFunc("/api/updateTown", updateTownHandler)
    http.HandleFunc("/api/getUsers", getUsersHandler)
    http.HandleFunc("/api/getUser", getUserHandler)
    http.HandleFunc("/api/addUser", addUserHandler)
    http.HandleFunc("/api/updateUser", updateUserHandler)
    http.HandleFunc("/api/removeUser", removeUserHandler)
    http.HandleFunc("/api/getMeta", getMetaHandler)

        /*
    http.HandleFunc("/api/getHouse", getHouseHandler)
    http.HandleFunc("/api/addHouse", addHouseHandler)
    http.HandleFunc("/api/editHouse", editHouseHandler)
    http.HandleFunc("/api/deleteHouse", deleteHouseHandler)
    http.HandleFunc("/api/getTags", getTagsHandler)
    http.HandleFunc("/api/addTag", addTagHandler)
    http.HandleFunc("/api/editTag", editTagHandler)
    http.HandleFunc("/api/deleteTag", deleteTagHandler)
    http.HandleFunc("/api/getUsers", getUsersHandler)
    http.HandleFunc("/api/addUser", addUserHandler)
    http.HandleFunc("/api/deleteUser", deleteUserHandler)
    http.HandleFunc("/api/editUser", editUserHandler)*/
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world!")
}

func viewHouseHandler(w http.ResponseWriter, r *http.Request) {
	if err := viewHouseTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func housesHandler(w http.ResponseWriter, r *http.Request) {
	if err := housesTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type Meta struct {
	Towns []Town
	Folders []Folder
}

func getMetaHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("getMeta")
	var resp Meta = Meta {
		Towns: []Town {
			{E: "Adachi, Mutsugi", J:"足立区六ツ木"},
			{E: "Edogawa, Kasai", J:"江戸川区葛西"},
		},
		Folders: []Folder {
			{FolderCode:"TerritoryNo", FolderType:FolderTypeSelect, Tags: []Tag {{TagCode:"1000"},{TagCode:"1001"},{TagCode:"1002"}}},
			{FolderCode:"Sex", FolderType:FolderTypeSelect, Tags: []Tag {{TagCode:"M"},{TagCode:"F"},{TagCode:"Family"}}},
		},
	}
	json.NewEncoder(w).Encode(resp)
}

type PutHouseRequest struct {
	tenantId string
	boxId string
	h House
}

func putHouseHandler(w http.ResponseWriter, r *http.Request) {
	var req PutHouseRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

type GetHousesRequest struct {
	BoxId string
}

func getHousesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("getHouses")
	defer r.Body.Close()
	var req GetHousesRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("%v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("%v", req)
	var resp []HouseD
	var h HouseD
	h.HouseId = 1000
	h.T = "Adachi, Mutsugi"
	h.Tags = make(map[string]string)
	h.Tags["TerritoryNo"] = "123"
	resp = append(resp, h)
	json.NewEncoder(w).Encode(resp)
}

type GetHouseRequest struct {
	HouseId string
}

func getHouseHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("getHouse")
	defer r.Body.Close()
	var req GetHouseRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("%v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("%v", req)
	var h HouseD
	h.HouseId = 1000
	h.T = "Adachi, Mutsugi"
	h.Tags = make(map[string]string)
	h.Tags["TerritoryNo"] = "123"
	json.NewEncoder(w).Encode(h)
}

func updateHouseHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("updateHouse")
}

func removeHouseHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("removeHouse")
}


type GetFoldersRequest struct {
	BoxId string
}

func getFoldersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("getFolders")
	defer r.Body.Close()
	var req GetFoldersRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("%v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("%v", req)
	var resp []Folder
	resp = make([]Folder, 0)
	var f Folder
	f.FolderCode = "TerritoryNo"
	f.Tags = []Tag{Tag{TagCode:"1001"}, Tag{TagCode:"1002"}}
	resp = append(resp, f)
	json.NewEncoder(w).Encode(resp)
}

type GetFolderRequest struct {
	BoxId string
	FolderCode string
}

func getFolderHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("getFolder")
	defer r.Body.Close()
	var req GetFolderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("%v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("%v", req)
	var f Folder
	f.FolderCode = "TerritoryNo"
	f.Tags = []Tag{{TagCode:"1001"}, {TagCode:"1002"}}
	json.NewEncoder(w).Encode(f)
}

func removeFolderHandler(w http.ResponseWriter, r *http.Request) {
}

type AddTagRequest struct {
	BoxId string
	FolderCode string
	TagCode string
}

func addTagHandler(w http.ResponseWriter, r *http.Request) {

}

type RemoveTagRequest struct {
	BoxId string
	FolderCode string
	TagCode string
}

func removeTagHandler(w http.ResponseWriter, r *http.Request) {

}

func getTownsHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var resp []Town = []Town{
		{E:"Adachi, Kameari", J:"足立区亀有"},
		{E:"Adachi, Kanamachi", J:"足立区金町"},
	}
	json.NewEncoder(w).Encode(resp)
}

func getTownHandler(w http.ResponseWriter, r *http.Request) {
}

func removeTownHandler(w http.ResponseWriter, r *http.Request) {
}

func updateTownHandler(w http.ResponseWriter, r *http.Request) {
}

func addTownHandler(w http.ResponseWriter, r *http.Request) {
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	var resp []User = []User {
		{UserId: "a@gmail.com", UserName: "A", Authority: 0},
		{UserId: "b@gmail.com", UserName: "B", Authority: 1},
	}
	json.NewEncoder(w).Encode(resp)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	var resp User = User{UserId:"a@gmail.com", UserName: "A", Authority: 0}
	json.NewEncoder(w).Encode(resp)
}

func removeUserHandler(w http.ResponseWriter, r *http.Request) {
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
}

var housesTemplate = template.Must(template.ParseFiles("templates/houses.html"))
var viewHouseTemplate = template.Must(template.ParseFiles("templates/viewHouse.html"))
var editHouseTemplate = template.Must(template.ParseFiles("templates/editHouse.html"))

// App Engine Datastore
type Tenant struct {
	TenantId string
	TownData []byte
	FolderData []byte
	UserIds []string
	UserNames []string
	Authorities []byte
	NextHouseId int
}

type HouseList struct {
	HouseListId string // TenantId-DataNum。Parentは、Tenant
	TenantId string
	ListId int
	HouseData []byte
}

// END App Engine Datastore

type Town struct {
	E, J string // 足立区六木, Adachi, Mutsugi
}

type HouseD struct {
	House
	TownJ, TownE string
}

type House struct {
	HouseId int
	T string // Adachi, Mutsugi
	C1,C2,C3 string // 1,2,3
	AJ string // 東京マンション
	AE string // Tokyo Mansion
	R string // 203
	NJ string // マリア
	NE string // Maria
	Memo string
	Tags map[string]string // {"TerritoryNo":"1001"}
	CreatedBy string
	CreatedDate time.Time
	EditedBy string
	EditedDate time.Time
}

const (
	Viewer = 0
	Editor = 1
	Owner = 2
)

type Folder struct {
	FolderCode string
	FolderType string
	Tags []Tag
}

const (
	FolderTypeSelect = "Select"
	FolderTypeInput = "Input"
	FolderTypeTextarea = "Textarea"
	FolderTypeDate = "Date"
)

type Tag struct {
	TagCode string
	HouseCondition string
}

type UserAuthority struct {
	UserId string
	FolderCode string
	Authority int
}

type User struct {
	UserId string
	UserName string
	Authority int
}
