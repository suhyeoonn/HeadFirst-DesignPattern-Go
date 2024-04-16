package main

type UsersUpdateBuilder struct {
	updateMap map[string]interface{}
}

func NewUsersUpdateBuilder() *UsersUpdateBuilder {
	return &UsersUpdateBuilder{
		updateMap: map[string]interface{}{},
	}		
}

func (ub *UsersUpdateBuilder) Name(name string) *UsersUpdateBuilder {
	ub.updateMap["name"] = name
	return ub
}

func (ub *UsersUpdateBuilder) Tel(tel string) *UsersUpdateBuilder {
	ub.updateMap["tel"] = tel
	return ub
}

func (ub *UsersUpdateBuilder) Password(pw string) *UsersUpdateBuilder {
	ub.updateMap["password"] = pw
	return ub
}

func (ub *UsersUpdateBuilder) Build() map[string]interface{} {
	return ub.updateMap
}

func main() {
	type Users struct {
		Id int
		User_id string
		Name string
		Tel string
	}

	var user Users
	updateBuilder := NewUsersUpdateBuilder().
		Name("elly_builder").
		Tel("010-1111-2222").
		Password("****").
		Build()
	GetORM().Model(&user).Table("user").Where("id = ?", 21).Updates(updateBuilder)

	// updateMap := map[string]interface{}{"name": "elly_22", "tel": "2222", "password": "abcs"}
	// GetORM().Model(&user).Table("user").Where("id = ?", 21).Updates(updateMap)
}
