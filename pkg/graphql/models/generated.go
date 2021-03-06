// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
)

type Category struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Dishes      []*Dish       `json:"dishes"`
	Restaurants []*Restaurant `json:"restaurants"`
}

type City struct {
	ID   int    `json:"id"`
	Code int    `json:"code"`
	Name string `json:"name"`
	Uf   string `json:"uf"`
}

type Dish struct {
	ID         int         `json:"id"`
	Category   *Category   `json:"category"`
	Restaurant *Restaurant `json:"restaurant"`
	Name       string      `json:"name"`
	Price      float64     `json:"price"`
	CookTime   int         `json:"cookTime"`
}

type Restaurant struct {
	ID          int         `json:"id"`
	Category    []*Category `json:"category"`
	Dishes      []*Dish     `json:"dishes"`
	OpenHour    string      `json:"openHour"`
	CloseHour   string      `json:"closeHour"`
	OpenDays    []Weekdays  `json:"openDays"`
	City        *City       `json:"city"`
	Name        string      `json:"name"`
	Description *string     `json:"description"`
	PhoneNumber string      `json:"phoneNumber"`
	Address     string      `json:"address"`
}

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
}

type CreateDishInput struct {
	Name         string  `json:"name"`
	RestaurantID int     `json:"restaurantId"`
	CategoryID   *int    `json:"categoryId"`
	Price        float64 `json:"price"`
	CookTime     int     `json:"cookTime"`
}

type CreateRestaurantInput struct {
	CategoryID  int        `json:"categoryId"`
	OpenHour    string     `json:"openHour"`
	CloseHour   string     `json:"closeHour"`
	OpenDays    []Weekdays `json:"openDays"`
	CityID      int        `json:"cityId"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	PhoneNumber string     `json:"phoneNumber"`
	Address     string     `json:"address"`
}

type CreateUserInput struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Senha       string `json:"senha"`
}

type UpdateCategoryInput struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UpdateDishInput struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	CategoryID *int    `json:"categoryId"`
	Price      float64 `json:"price"`
	CookTime   int     `json:"cookTime"`
}

type UpdateRestaurantInput struct {
	ID          int        `json:"id"`
	OpenHour    *string    `json:"openHour"`
	CloseHour   *string    `json:"closeHour"`
	OpenDays    []Weekdays `json:"openDays"`
	Name        *string    `json:"name"`
	Description *string    `json:"description"`
	PhoneNumber *string    `json:"phoneNumber"`
	Address     *string    `json:"address"`
}

type Weekdays string

const (
	WeekdaysMonday    Weekdays = "MONDAY"
	WeekdaysTuesday   Weekdays = "TUESDAY"
	WeekdaysWednesday Weekdays = "WEDNESDAY"
	WeekdaysThursday  Weekdays = "THURSDAY"
	WeekdaysFriday    Weekdays = "FRIDAY"
	WeekdaysSaturday  Weekdays = "SATURDAY"
	WeekdaysSunday    Weekdays = "SUNDAY"
)

var AllWeekdays = []Weekdays{
	WeekdaysMonday,
	WeekdaysTuesday,
	WeekdaysWednesday,
	WeekdaysThursday,
	WeekdaysFriday,
	WeekdaysSaturday,
	WeekdaysSunday,
}

func (e Weekdays) IsValid() bool {
	switch e {
	case WeekdaysMonday, WeekdaysTuesday, WeekdaysWednesday, WeekdaysThursday, WeekdaysFriday, WeekdaysSaturday, WeekdaysSunday:
		return true
	}
	return false
}

func (e Weekdays) String() string {
	return string(e)
}

func (e *Weekdays) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Weekdays(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Weekdays", str)
	}
	return nil
}

func (e Weekdays) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
