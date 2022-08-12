package model

type RoleListInput struct {
	Page     int
	PageSize int
}

type RoleDetailInput struct {
	Id int // role id
}

type RoleUpdateInput struct {
	Id int
}

type RoleDeleteInput struct {
	Id int
}
