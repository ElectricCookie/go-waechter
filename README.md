<p align="center">
    <img alt="Waechter - Logo" src="http://imgur.com/WWeVTGh.png" height="30"/>
    <br/>
    <br/>
    <a href="https://travis-ci.org/ElectricCookie/go-waechter"><img alt="Travis CI build status" src="https://travis-ci.org/ElectricCookie/go-waechter.svg?branch=master" /></a>
    <a href="https://goreportcard.com/report/github.com/ElectricCookie/go-waechter"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/ElectricCookie/go-waechter" /></a>
    <a href="https://codecov.io/gh/ElectricCookie/go-waechter"><img alt="Codecov Card" src="https://codecov.io/gh/ElectricCookie/go-waechter/branch/master/graph/badge.svg" /></a>
 
</p>

---

Go Wächter [German = guard] is a library to implement JWT auth in go apps. It is supposed to bootstrap a web application bringing all necessary features.

---

## Current status
**Currently working on the basic functionality. This package is not ready for consumption yet.**

## Dependencies and Adapters
Wächter does not rely on external services, but it comes with adapters for SMTP and MongoDB. You can also implement your own connections by following the interfaces defined in **db-adapter.go** and **email-adapter.go**

## Features

- [ ] Administration Notifications
- [ ] Brute-Force Protection
- [ ] Invite System 
- [ ] Notification Settings
- [ ] OAuth
- [ ] IP Location check
- [x] Auth using JWT, Access/Refresh Token architecture
- [x] Registration
- [x] Password Recovery
- [x] Email Verification 
- [x] Adapter based allows to add custom email and database connections 
- [x] Fully internationalized. Allows to translate all messages and emails
- [x] Comes with beautiful email templates, but you can add your own
- [x] Configurable access token generation builds the foundation for complex auth needs.



## Security

Wäechter uses scrypt to hash passwords using an application secret and a salt unique to every user. It also hashes email activation tokens.

The refresh tokens are saved as JWT to httpOnly cookies (if you're using the gin adapter). This is recommended to prevent XSS!

## Getting started

```bash
    go get github.com/ElectricCookie/go-waechter
```

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

