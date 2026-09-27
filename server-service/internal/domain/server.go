package domain



type UserServer struct {
	ID int  // id в табле
	Addr string // должно быть уникально
	Name string 
	UserName string
	Password string
	UserId int
}