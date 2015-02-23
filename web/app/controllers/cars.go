package controllers

import (
    "tvners.com/web/app/models"
    "github.com/revel/revel"
    "encoding/json"
)

type CarsController struct {
    GorpController
}

func (c CarsController) parseCar() (models.Car, error) {
    car := models.Car{}
    err := json.NewDecoder(c.Request.Body).Decode(&car)

    return car, err
}

func (c CarsController) Add() revel.Result {
    if car, err := c.parseCar(); err != nil {
        return c.RenderText("Unable to parse the Car from JSON.")
    } else {
        // Validate the model
        car.Validate(c.Validation)
        if c.Validation.HasErrors() {
            // Do something better here!
            return c.RenderText("You have error in your Car.")
        } else {
            if err := c.Txn.Insert(&car); err != nil {
                return c.RenderText("Error inserting record into database!")
            } else {
                return c.RenderJson(car)
            }
        }
    }
}

func (c CarsController) Get(id int64) revel.Result {
    car := new(models.Car)
    err := c.Txn.SelectOne(car, `SELECT * FROM Car WHERE id = ?`, id)

    if err != nil {
      revel.ERROR.Printf("%s", err)
        return c.RenderText("Error.  Item probably doesn't exist.")
    }
    return c.RenderJson(car)
}

func (c CarsController) List() revel.Result {
    lastId := parseIntOrDefault(c.Params.Get("lid"), -1)
    limit := parseUintOrDefault(c.Params.Get("limit"), uint64(25))
    cars, err := c.Txn.Select(models.Car{}, 
        `SELECT * FROM Car WHERE Id > ? LIMIT ?`, lastId, limit)
    if err != nil {
        return c.RenderText(
            "Error trying to get records from DB.")
    }
    return c.RenderJson(cars)
}

func (c CarsController) Update(id int64) revel.Result {
    car, err := c.parseCar()
    if err != nil {
        return c.RenderText("Unable to parse the Car from JSON.")
    }
    // Ensure the Id is set.
    car.Id = id
    success, err := c.Txn.Update(&car)
    if err != nil || success == 0 {
        return c.RenderText("Unable to update car.")
    }
    return c.RenderText("Updated %v", id)
}

func (c CarsController) Delete(id int64) revel.Result {
    success, err := c.Txn.Delete(&models.Car{Id: id})
    if err != nil || success == 0 {
        return c.RenderText("Failed to remove car")
    }
    return c.RenderText("Deleted %v", id)
}