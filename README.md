# go-api
Api For Practice
<b> SITE WORK ON PHP </b>

# Config
File: ./configs/apiserver.toml

# DataBase ( migration )
Directory: ./migrations

# Main Page
http://127.0.0.1:8080

# Register company 
127.0.0.1:8080/company/register

<b>POST</b><br>
{<br>
    "Name":      "Test",<br>
	"Email":       "test@mail.ru",<br>
	"Password":    "123456",<br>
	"Description": "test",<br>
	"Country":     "test",<br>
	"City":        "test"<br>
}

* The password is encrypted and is not sent in the response

<b>Response</b><br>
{<br>
    "id":18,<br>
    "name":"Test",<br>
    "email":"test@mail.ru",<br>
    "description":"test",<br>
    "country":"test",<br>
    "city":"test"<br>
}

# Login company
127.0.0.1:8080/company/login
<b>POST</b><br>
{<br>
	"Email":       "test@mail.ru",<br>
	"Password":    "123456"<br>
}

* 

<b>Response</b><br>
{<br>
"id":18,<br>
"name":"Test",<br>
"email":"test@mail.ru",<br>
"password":"$2a$04$nJ.releYHd0rFdPXIwegk.KSLWXJ.6LETXn51UDUcdeI0H.1TB0Z2",<br>
"description":"test",<br>
"country":"test",<br>
"city":"test"<br>
}

# Vacancy Create
127.0.0.1:8080/vacancy/create

<b>POST</b><br>
{<br>
    "Position":     "Manager",<br>
	"Experience":   "1",<br>
	"Description":  "Description vacancy",<br>
    "Company_id":   "18"<br>
}

<b>Response</b><br>
{<br>
    "id":11,<br>
    "position":"Manager",<br>
    "experience":1,<br>
    "description":"Description vacancy",<br>
    "company_id":18<br>
}

# Vacancy Find

127.0.0.1:8080/vacancy/find

<b>POST</b><br>
{<br>
    "Company_id":   "18"<br>
}


<b>Response</b><br>
[<br>
{<br>
    "id":11,<br>
    "position":"Manager",<br>
    "experience":1,<br>
    "description":"Description vacancy",<br>
    "company_id":18<br>
},<br>
{<br>
    "id":12,<br>
    "position":"Manager",<br>
    "experience":1,<br>
    "description":"Description vacancy",<br>
    "company_id":18<br>
},<br>
{<br>
    "id":13,<br>
    "position":"Manager",<br>
    "experience":1,<br>
    "description":"Description vacancy",<br>
    "company_id":18<br>
    }<br>
]

# Validation 
 * All required
 * Password - Length(6, 40)
 * Email - isEmail?
