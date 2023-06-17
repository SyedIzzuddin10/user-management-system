package service

import (
	"context"
	"fmt"
	"userManagementSystem/ent"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

// type UserDatabase struct {
// 	Users map[uuid.UUID]User
// }

// func NewUserDatabase() *UserDatabase {
// 	return &UserDatabase{Users: make(map[uuid.UUID]User)}
// }

type UserService struct {
	c      echo.Context
	ctx    context.Context
	client *ent.Client
}

func NewUserService(c echo.Context, ctx context.Context) *UserService {
	return &UserService{
		c:      c,
		ctx:    ctx,
		client: OpenDB(),
	}
}

func (db *UserService) GetUserList() ([]*ent.User, error) {
	// c := db.c
	ctx := db.ctx
	client := db.client

	// var users []User
	// for _, item := range db.Users {
	// 	users = append(users, item)
	// }

	// fmt.Println("Get user list: ", db.Users)

	userList, err := client.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return userList, nil
}

func (db *UserService) GetUserById() (*ent.User, error) {
	c := db.c
	ctx := db.ctx
	client := db.client

	userRequest := new(User)
	err := c.Bind(userRequest)
	if err != nil {
		return nil, err
	}

	// idUUID, err := uuid.Parse(string(userRequest.ID))
	// if err != nil {
	// 	return nil, err
	// }

	user, err := client.User.Get(ctx, userRequest.ID)
	if err != nil {
		return nil, err
	}

	// user, ok := db.Users[idUUID]
	// if !ok {
	// 	return nil, fmt.Errorf("user ID: %v does not exist", id)
	// }

	// fmt.Println("Get user by ID: ", user)

	return user, nil
}

func (db *UserService) CreateUser() (*ent.User, error) {
	c := db.c
	ctx := db.ctx
	client := db.client

	userRequest := new(User)
	err := c.Bind(userRequest)
	if err != nil {
		return nil, err
	}

	fmt.Printf("userRequest: %+v\n", userRequest)

	id := uuid.New()

	user, err := client.User.Create().SetID(id).SetName(userRequest.Name).SetEmail(userRequest.Email).SetPassword(userRequest.Password).Save(ctx)
	if err != nil {
		return nil, err
	}

	// db.Users[id] = User{ID: id, Name: name, Email: email, Password: password}
	// fmt.Println("New user: ", db.Users[id])

	return user, nil
}

func (db *UserService) UpdateUser() (*ent.User, error) {
	c := db.c
	ctx := db.ctx
	client := db.client

	userRequest := new(User)
	err := c.Bind(userRequest)
	if err != nil {
		return nil, err
	}

	fmt.Printf("userRequest: %+v\n", userRequest)

	// idUUID, err := uuid.Parse(id)
	// if err != nil {
	// 	return nil, err
	// }

	user, err := client.User.Get(ctx, userRequest.ID)
	if err != nil {
		return nil, err
	}

	fmt.Printf("user: %+v\n", user)

	update := client.User.UpdateOne(user)

	if userRequest.Name != "" {
		update = update.SetName(userRequest.Name)
	}
	if userRequest.Email != "" {
		update = update.SetEmail(userRequest.Email)
	}
	if userRequest.Password != "" {
		update = update.SetPassword(userRequest.Password)
	}

	// updatedUser, err := user.Update().SetName(userRequest.Name).SetEmail(userRequest.Email).SetPassword(userRequest.Password).Save(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	updatedUser, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Printf("updatedUser: %+v\n", updatedUser)

	// _, ok := db.Users[idUUID]
	// if !ok {
	// 	return nil, fmt.Errorf("userID: %v does not exist", id)
	// }

	// db.Users[idUUID] = User{ID: idUUID, Name: name, Email: email, Password: password}
	// updatedUser := db.Users[idUUID]
	// fmt.Println("Updated user: ", updatedUser)

	return updatedUser, nil
}

func (db *UserService) DeleteUser() (*ent.User, error) {
	c := db.c
	ctx := db.ctx
	client := db.client

	userRequest := new(User)
	err := c.Bind(userRequest)
	if err != nil {
		return nil, err
	}

	fmt.Printf("userRequest: %+v\n", userRequest)

	// idUUID, err := uuid.Parse(id)
	// if err != nil {
	// 	return nil, err
	// }
	// deletedUser, ok := db.Users[idUUID]
	// if !ok {
	// 	return fmt.Errorf("userID: %v does not exist", id)
	// }

	// delete(db.Users, idUUID)

	user, err := client.User.Get(ctx, userRequest.ID)
	if err != nil {
		return nil, err
	}

	err = client.User.DeleteOneID(userRequest.ID).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}
