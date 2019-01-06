# go-cleanarch-boilerplate
## Inspiration
- A small but handy tool in order to create a quick scaffolding web application in Golang with supported popular web frameworks (Martini, Gin, ...) and databases (MongoDB, PostgreSQL, MySQL, ...)
- Especially in Microservices Architecture as we usually have to create many projects (file by file, package by package) for microservices. The tool will help us not have to repeat that boring procedure anymore so we will be able to get a new microservice instance up and run very quickly.
- The generated code is organized by following Clean Architecture by Robert C. Martin (Uncle Bob)
- One of the most beautiful ideas of Clean Architecture is Dependency Rule that must be applied when building software systems whose parts are to be testable and independent of frameworks, UIs, or databases. When following this rule, one ends up with a loosely coupled system with clear separation of concerns.
## Clean Architecture
### From architectural idea
![alt text](https://github.com/minhduccm/go-cleanarch-boilerplate/blob/master/images/clean-arch-idea.png "Clean Architecture")

For more detail about Clean Architecture, please look at the post here https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
### To implementation in Golang
// Add soon
## TODO
* [x] Scaffolding with Clean Architecture
* [x] Unit testing samples
* [ ] Generating code by commands
* [ ] Dockerization
* [ ] PostgreSQL
* [ ] MySQL
* [ ] Gin (web framework)
* [ ] Logging
* [ ] Grateful shutting down
