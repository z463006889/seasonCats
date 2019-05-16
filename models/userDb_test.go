package models

import "testing"

func init()  {
	
}

func TestCreateUser(t *testing.T) {
	_,err:=CreateUser("zhaohu","123456","橙子")
	t.Errorf("%s",err)
}
