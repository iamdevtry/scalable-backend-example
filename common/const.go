package common

type DbType int

const (
	DbTypeRestaurant DbType = 1
	DbTypeUser       DbType = 2
)

const (
	CurrentUser = "user"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
