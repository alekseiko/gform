package models

import (
    "gopkg.in/mgo.v2"
)

type Field struct {
    Id string
    FieldId string
    Name string
    MappedName string // this can be equal name
}

type Form struct {
    Id string
    Url string
    Fields []* Field
}

type Storage interface {
    GetForm(id string) (form *Form, e error)
    GetForms() (forms []*Form, e error)
    SaveForm(form *Form) (r bool, e error)
    DeleteForm(form *Form) (r bool, e error)
}

type MongoStorage struct {
   Session *mgo.Session
}

func (s MongoStorage) GetForm(id string) (form *Form, e error) {
    return nil, nil
}

func (s MongoStorage) GetForms() (form []*Form, e error) {
    return nil, nil
}

func (s MongoStorage) SaveForm(form *Form) (r bool, e error) {
    return false, nil
}

func (s MongoStorage) DeleteForm(form *Form) (r bool, e error) {
    return false, nil
}
