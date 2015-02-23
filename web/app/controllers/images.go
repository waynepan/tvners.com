package controllers

import (
    "tvners.com/web/app/models"
    "github.com/revel/revel"
    "encoding/json"
)

type ImagesController struct {
    GorpController
}

func (i ImagesController) parseImage() (models.Image, error) {
    image := models.Image{}
    err := json.NewDecoder(i.Request.Body).Decode(&image)

    return image, err
}

func (i ImagesController) Add(car_id int64) revel.Result {
    if image, err := i.parseImage(); err != nil {
        return i.RenderText("Unable to parse the Image from JSON.")
    } else {
        // Validate the model
        image.Validate(i.Validation)
        if i.Validation.HasErrors() {
            // Do something better here!
            return i.RenderText("You have error in your Image.")
        } else {
            if err := i.Txn.Insert(&image); err != nil {
                return i.RenderText("Error inserting record into database!")
            } else {
                return i.RenderJson(image)
            }
        }
    }
}

func (i ImagesController) Get(car_id int64, id int64) revel.Result {
    image := new(models.Image)
    err := i.Txn.SelectOne(image, `SELECT * FROM Image WHERE id = ? and car_id = ?` , id, car_id)

    if err != nil {
        return i.RenderText("Error.  Item probably doesn't exist.")
    }
    return i.RenderJson(image)
}

func (i ImagesController) List(car_id int64) revel.Result {
    limit := parseUintOrDefault(i.Params.Get("limit"), uint64(25))
    images, err := i.Txn.Select(models.Image{}, `SELECT * FROM Image WHERE car_id = ? LIMIT ?`, car_id, limit)
    if err != nil {
        return i.RenderText("Error trying to get records from DB.")
    }
    return i.RenderJson(images)
}
