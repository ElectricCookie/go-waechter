<p align="center">
    <img alt="Waechter - Logo" src="http://imgur.com/WWeVTGh.png" height="30"/>
    <br/>
    <br/>
    <a href="https://goreportcard.com/report/github.com/ElectricCookie/go-waechter"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/ElectricCookie/go-waechter" /></a>
</p>


---

Go Wächter [german = guard] is a library to implement JWT auth in go apps. It is supposed to bootstrap a web application bringing all necessary features.

---

## Current status
**Currently working on the basic functionality. This package is not ready for consumption yet.**

## Features

- [ ] Administration Notifications
- [ ] Brute-Force Protection
- [ ] Invite System 
- [ ] Notification Settings
- [ ] Planned: OAuth, Check IP location
- [x] Auth using JWT, Access/Refresh Token architecture
- [x] Registration
- [x] Password Recovery
- [x] Email Verification 
- [x] Adapter based allows to add custom email and database connections 
- [x] Fully internationalized. Allows to translate all messages and emails
- [x] Comes with beautiful email templates, but you can add your own
- [x] Default adapters included for mongoDB (using mgo) and SMTP
- [x] Configurable access token generation builds the foundation for complex auth needs.



## Security

Wäechter uses scrypt to hash passwords using an application secret and a salt unique to every user. It also hashes email activation tokens.

## Getting started

```golang
    import(
        "github.com/ElectricCookie/go-waechter"
    )

    func main(){

        // Setup a db adapter

        dbAdapter := NewMongoAdapter("localhost:27017", "waechter-test-db")

        // Setup an SMTP email adapter
        emailAdapter := NewSMTPAdapter("myemailhost",1337,)


        // TODO i18n adapter

        // Application secret is used to hash jwts!
        w := New("some-application-secret", "jwt-issuedby-claim", dbAdapter, emailAdapter)

        // You can now use waechter. To get started quickly you can also use the gin adapter

    }
```


