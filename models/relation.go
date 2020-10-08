package models

/* Relation es el modelo para insertar una relacion entre un usuario y otro. */
type Relation struct {
	UsuarioID string `bson:"userid" json:"userId"`
	RelationID string `bson:"relationid" json:"relationId"`
}
