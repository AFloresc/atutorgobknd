package rbac

import jwt "github.com/dgrijalva/jwt-go"

func NewJWTTokenForTest() (token *jwt.Token, err error) {
	return jwt.Parse(NewStringTokenForTest(), func(token *jwt.Token) (interface{}, error) {
		return []byte(NewAuthenticationKeyForTest()), nil
	})
}

func NewStringTokenForTest() string {
	return "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ0b2tlbiI6ImYyMWYwMTY1MWExMDczNDU3MDAwZjNiOTVjNjlmMDY3OGE5Yzk3MjUiLCJ1c2VyX2lkIjoiMTczMSIsInJvbGVzIjpbImxvZ2luIiwibW9kaWZ5Iiwic2hhcmUiLCJhcGkiLCJwdWJsaXNoZXIiLCJtb2RpZnktY2xpcHBpbmdzLWFuZC1wcm9maWxlIiwibW9kaWZ5LXRvcGljcyIsInRhZ3MiXSwiaWF0IjoiMTUyNTI0OTU3OCIsImV4cCI6IjE1Mjc4NDE1NzgifQ.vV9HpexQ3aJz7hSHo-2XSUV7U-GepAyaPxDH909lmHw"
}

func NewAuthenticationKeyForTest() string {
	return "12d190c6e28fc4f855e5a87570a456810d590fa6e676k8ec1f"
}

func NewAdminStringTokenForTest() string {
	return "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ0b2tlbiI6ImVlOGIxZTJkNmEyM2I3OWQ4MzhkOTFmMmY2YTU5Mjk5ZGQ0YzBlZmQiLCJ1c2VyX2lkIjoiMTg0NyIsInJvbGVzIjpbImxvZ2luIiwiYWRtaW4iLCJkZXZlbG9wZXIiLCJtb2RpZnkiLCJzaGFyZSIsInB1Ymxpc2hlciIsIm1vZGlmeS1jbGlwcGluZ3MtYW5kLXByb2ZpbGUiLCJtb2RpZnktdG9waWNzIl0sImlhdCI6IjE1NTE1NDk1MDciLCJleHAiOiIxNTU0MTQxNTA3In0.Nj0heSZa1C4hjD8t0cLrJAjGVgQahMXOveXg-in46t8"
}

func NewAnalystStringTokenForTest() string {
	return "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ0b2tlbiI6IjcxNTFhOWZiZGRmOWVmMzJmMDI5ODdkN2FmMTA0NDBkMWVjYWUzYTciLCJ1c2VyX2lkIjoiMTg0NyIsInJvbGVzIjpbImxvZ2luIiwiYWRtaW4iLCJkZXZlbG9wZXIiLCJtb2RpZnkiLCJzaGFyZSIsInB1Ymxpc2hlciIsIm1vZGlmeS1jbGlwcGluZ3MtYW5kLXByb2ZpbGUiLCJtb2RpZnktdG9waWNzIiwiYW5hbHlzdCJdLCJpYXQiOiIxNTUzMTg2OTUzIiwiZXhwIjoiMTU1NTc3ODk1MyJ9.Icg2PTNYbjfZdmwuq6_oJLLD2bC9J1O3xvd-GWfA_xg"
}