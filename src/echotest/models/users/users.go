package users

import "fmt"
import "context"
// import "errors"

import "github.com/Iljaaa/echotest/src/common/db"; 


type User struct {
	Id        int    
	Name      string 
	// Phone string 
  Login	    string	
  Password	string	
  CreatedAt string	// string ?
  UpdatedAt	string  // string ?
}


func Insert (user User) User {

	/*row := conn.QueryRow(context.Background(),
    "INSERT INTO phonebook (name, phone) VALUES ($1, $2) RETURNING id",
    rec.Name, rec.Phone)
  var id uint64
  err = row.Scan(&id)
  if err != nil {
    log.Errorf("Unable to INSERT: %v\n", err)
    w.WriteHeader(500)
    return
  }*/

	// update id fild

	return User{}
}

//
// find user by id
//
func FindById (id int) (*User, error) {

  query := fmt.Sprintf("select id, name FROM users where id = %d limit 1", id)  
  // fmt.Printf("query %s\n", query)

  u := User{}

	err := db.GetPool().QueryRow(context.Background(), query).Scan(&u.Id, &u.Name)
  if err != nil {
    return nil, err
  }

  return &u, nil;
}

//
// find user by login
//
func FindByLogin (login string) (*User, error) {

  query := fmt.Sprintf("select id, name, login, password FROM users where login = '%s' limit 1", login)  
  // fmt.Printf("query %s\n", query)

  u := User{}

	err := db.GetPool().QueryRow(context.Background(), query).Scan(&u.Id, &u.Name, &u.Login, &u.Password)
  if err != nil {
    return nil, err
  }

  return &u, nil;
}
