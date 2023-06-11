package api

import (
	"fmt"
	_ "fmt"
	"sharath/database"

	"github.com/gofiber/fiber"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct{
	gorm.Model
	Name   string   `gorm: not null `
	Email  string   `gorm: not null `
	Password string `gorm: not null `
	Sid     string  `gorm: not null `
	Notes  []Notes
}

type Notes struct{
	gorm.Model
	Note    string `gorm: not null`
	Sid     string `gorm: not null `
}

type NoteRequest struct{
	Note   string   `gorm: not null`
}

type Login struct{
	Email  string   `json:email`
	Password string `json:password`
}



func RegisterNewUser(c *fiber.Ctx){
	db:=database.DBConn
	user:=new(User)
	if err:=c.BodyParser(&user); err!=nil{
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unable to parse JSON"})
		return
	}
	var existingUser User
	db.Where("Email=?",user.Email).First(&existingUser)

	if existingUser.Email!=""{
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User With This Email Already Exists"})
		return
	}
	db.Create(user)
	c.Status(fiber.StatusOK)
}



func UserLogin(c *fiber.Ctx){
	db:=database.DBConn
	userLogin:=new(Login)

	if err := c.BodyParser(&userLogin); err!=nil{
		c.Status(fiber.StatusBadRequest)
		return
	}

	//Check whether user with the email exists
	var checkUser User
	if err:=db.Where("Email= ?",userLogin.Email).First(&checkUser).Error; err!=nil{
		if err == gorm.ErrRecordNotFound{
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"unauthorized": "Email not found"})
			return
		}
		c.Status(fiber.StatusInternalServerError)
        return
	}

	//Validation of user credentials
	if checkUser.Password == userLogin.Password{
		uuid:=uuid.New()
		checkUser.Sid=uuid.String()
	}else{
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"unauthorised":"Email and password combination does not match"})
		return
	}
	db.Save(&checkUser)
	c.Status(fiber.StatusOK).JSON(fiber.Map{"sid":checkUser.Sid})
}

func GetAllNotes(c *fiber.Ctx){
	db:=database.DBConn
	sid:=c.Params("sid")

	var existingUser User
	if err:=db.Where("sid= ?",sid).First(&existingUser).Error; err !=nil{
		if err == gorm.ErrRecordNotFound{
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"Sid is invalid"})
		    return
		}
		c.Status(fiber.StatusInternalServerError)
		return
	}
	fmt.Println(existingUser.Notes)
	var notes[] Notes
	if err:=db.Where("sid= ?",sid).Find(&notes).Error; err!=nil{
		if err == gorm.ErrRecordNotFound{
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
			return
		}
		c.Status(fiber.StatusBadRequest)
		return
	}
	c.Status(fiber.StatusOK).JSON(fiber.Map{"notes":notes})
}

func CreateNote(c *fiber.Ctx){
	db:=database.DBConn
	sid:=c.Params("sid")

	//Find the user with SID
	var existingUser User
	if err:= db.Where("sid= ?",sid).First(&existingUser).Error; err!=nil{
		if err == gorm.ErrRecordNotFound{
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"Invalid Sid"})
			return
		}
		c.Status(fiber.StatusInternalServerError)
		return
	}

	var noteRequest NoteRequest
	if err:= c.BodyParser(&noteRequest); err!=nil{
		c.Status(fiber.StatusBadRequest)
		return
	}

	note:= Notes{
		Sid: sid,
		Note: noteRequest.Note,
	}

	result := db.Create(&note)
	if result.Error!=nil{
		c.Status(fiber.StatusInternalServerError)
		return
	}
	c.Status(fiber.StatusOK).JSON(fiber.Map{"Id":note.ID})
}