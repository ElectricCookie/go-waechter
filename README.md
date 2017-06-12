<p align="center"><img src="http://imgur.com/WWeVTGh.png" height="30"/></p>


---

Go Wächter [german = guard] is a library to implement JWT auth in go apps. It is supposed to bootstrap a web application using all necessary features.

---

## Features

* Auth using JWT, Access/Refresh Token architecture
* Registration
* Password Recovery
* Administration Notifications
* Brute-Force Protection
* Invite System
* Notification Settings
* Email Verification 
* Adapter based allows to add custom email and database connections
* Fully internationalized. Allows to translate all messages and emails
* Comes with beautiful email templates, but you can add your own
* Default adapters included for mongoDB (using mgo) and SMTP
* Configurable access token generation builds the foundation for complex auth needs.
* Planned: OAuth, Check IP location

## Current status
**Currently working on the basic functionality. This package is not ready for consumption yet. There are not tests either yet.**

## Security

Wäechter uses scrypt to hash passwords using an application secret and a salt unique to every user. It also hashes email activation tokens as well. All source code is open source.




