# Design Pattern Go

## Prerequisites
- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
## Usage
- 23 Classic Design Patterns demo source code in Go (Golang)
- Each pattern has a unit test file
- Run all unit tests
## Patterns
### abstract-factory
- [ ] [Abstract Factory]()
- [ ] [Builder]()
```
- Độc lập với việc tạo ra các object phức tạp
- Tức là A được build từ B,C thì B,C không cần biết A, và B và C không cần biết nhau
- Chạy song song
```
- [ ] [Factory Method]()
- [ ] [Prototype]()
```
1 Class chỉ tạo ra 1 object
```
- [ ] [Singleton]()
### structural
- [ ] [Adapter]()
- [ ] [Bridge]()
- [ ] [Composite]()
- [ ] [Decorator]()
- [ ] [Facade]()
- [ ] [Flyweight]()
- [ ] [Proxy]()
- [x] [Façade]()
```
Gom module về thành 1 service
```
### behavioral
- [ ] [Chain of Responsibility]()

```
- Đưa request vào 1 chuối handler. Mỗi handler sẽ xử lý request hoặc chuyển request sang handler tiếp theo
- Giống middware django
- Cho request là chủ yếu
- Input của 1 handler là output của handler trước
- Chạy tuần tự
```
- [ ] [Command]()
- [ ] [Interpreter]()
- [ ] [Iterator]()
- [ ] [Mediator]()
- [ ] [Memento]()
- [ ] [Observer]()
- [ ] [State]()
- [ ] [Strategy]()
- [ ] [Template Method]()
- [ ] [Visitor]()
## References
- [Design Patterns](https://refactoring.guru/design-patterns)
- [go-patterns](https://github.com/viettranx/go-design-pattern/tree/main)
