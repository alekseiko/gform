package models

import (
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type Field struct {
    FieldId string `json:"field_id" bson:"field_id"`
    Name string `json:"name" bson:"name"`
    MappedName string `json:"mapped_name" bson:"mapped_name"`// this can be equal name
}

type Form struct {
    Id string `json:"id" bson:"_id,"`
    PageUrl string `json:"page_url" bson:"page_url"`
    Url string `json:"url" bson:"url"`
    Fields []* Field `json:"fields" bson:"fields"`
    Method string `json:"method" bson:"method"`
}

type Storage interface {
    GetForm(id string) (form *Form, e error)
    GetForms() (forms []*Form, e error)
    SaveForm(form *Form) (e error)
    DeleteForm(form *Form) (e error)
}

type MongoStorage struct {
   Session *mgo.Session
}

const(
    DB_NAME = "gform"
    FORM_COLLECTION = "forms"
)


func (s MongoStorage) GetForm(id string) (form *Form, e error) {
    return nil, nil
}

func (s MongoStorage) GetForms() (form []*Form, e error) {
    return nil, nil
}

func (s MongoStorage) SaveForm(form *Form) (e error) {
    var c = s.Session.DB(DB_NAME).C(FORM_COLLECTION)
    if (form.Id == "") {
        form.Id = bson.NewObjectId().Hex()
        return c.Insert(form)
    }
    return c.UpdateId(form.Id, bson.M{"$set": form})
}

func (s MongoStorage) DeleteForm(form *Form) (e error) {
    return nil
}
