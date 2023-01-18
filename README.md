# go-api
Api For Peactice

# Config
File: ./configs/apiserver.toml

# DataBase ( migration )
Directory: ./migrations

# Main Page
http://127.0.0.1:8080

# Register company 
127.0.0.1:8080/company/register

POST <br>
{<br>
    "Name":      "Test",<br>
	"Email":       "test@mail.ru",<br>
	"Password":    "123456",<br>
	"Description": "test",<br>
	"Country":     "test",<br>
	"City":        "test"<br>
}

* The password is encrypted and is not sent in the response

Response
{
    "id":18,
    "name":"Test",
    "email":"test@mail.ru",
    "description":"test",
    "country":"test",
    "city":"test"
}
