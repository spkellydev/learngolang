# Learn GoLang.io

## Dependencies

- Mux
- GoDotEnv
- PostGres driver

## Roadmap

- ~~Console logger for HTTP requests~~
- ~~File logger~~
- ~~.env~~
- ~~Connect to database~~
- ~Grouped routing
- Validate response before hitting database
- JWT
- Authentication system
- User roles
- ~Standardized CRUD API
- ~Documentation MVC
- etc

#### Helpful functions

- `iptables -t nat -A PREROUTING -p tcp --dport 80 -j REDIRECT --to-port 8080`
- redirect connections on port 80 to some other port you can open as normal user
- run as root
- delete `iptables -t nat --line-numbers -n -L` && `iptables -t nat -D PREROUTING 2`
