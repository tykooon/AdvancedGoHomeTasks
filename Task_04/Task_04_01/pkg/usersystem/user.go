package usersystem

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type User struct {
	Id       int64
	Email    string
	passHash string
	FullName string
	IsActive bool
}

func NewUser(firstName, lastName string) (res *User, ok bool) {
	firstName, lastName, ok = NameCheck(firstName, lastName)
	if !ok {
		return nil, false
	}
	var email = fmt.Sprintf("%s%s@coolcompany.com", firstName, lastName)
	var hash, _ = bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)

	res = &User{
		FullName: fmt.Sprintf("%s %s", firstName, lastName),
		Email:    email,
		passHash: string(hash),
		IsActive: true,
	}
	return res, true
}

func NewUserFromString(str string) (user *User, ok bool) {
	nameSlice := strings.Split(str, " ")
	if len(nameSlice) != 2 {
		return nil, false
	}
	user, ok = NewUser(nameSlice[0], nameSlice[1])
	return
}

func (user *User) ToString() string {
	status := "Active"
	if !user.IsActive {
		status = "Inactive"
	}
	return fmt.Sprintf("Id: %2d, %20s, Email: %35s, %s", user.Id, user.FullName, user.Email, status)
}

func NameCheck(first, last string) (f, l string, ok bool) {
	var nameReg regexp.Regexp = *regexp.MustCompile(`^[a-zA-Z]+$`)
	if !nameReg.MatchString(first) || !nameReg.MatchString(last) {
		return "", "", false
	}
	nameCaser := cases.Title(language.AmericanEnglish)
	f = nameCaser.String(first)
	l = nameCaser.String(last)
	return f, l, true
}

func (user *User) PassCheck(str string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.passHash), []byte(str))
	return err == nil
}
