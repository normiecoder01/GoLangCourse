package main

import(
	"encoding/json"
	"fmt"
)

func main(){
	type User struct {
		Name string `json:"name"`
		Age int `json:"age"`
		Phone string `json:"phone"`
		Username string `json:"username"`
		Password string `json:"-"`
	}

	userJson:=`{
	"name":"Prathamesh",
	"age":39,
	"phone":"7028843375",
	"username":"excellentcoder",
	"password": "Pass@123"
	}
	`

	u:= User{
		Name:"Anupal",
		Age:30,
		Phone:"7028843375",
		Username:"normiecoder",
		Password:"Pass@123",
	}
	byteJson,err := json.Marshal(u)
	if err!= nil{
		 fmt.Println(err)
	}
		fmt.Println(string(byteJson))

		var user User
	err = json.Unmarshal([]byte(userJson), &user)
		if err!= nil{
		 fmt.Println(err)
	}

	fmt.Printf("%+v",user)

}