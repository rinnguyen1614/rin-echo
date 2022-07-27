# Rin-echo

### Directories

- [deploy](deploy/) deployment
- [internal](internal/) application code
- [scripts](scripts/) development scripts
- <del>[terraform](terraform/) - infrastructure definition</del>
- [web](web/) - frontend reatjs code

### Features

- Dynamic paramaters including sorting, filtering, selecting and pagination in RESTful API

- Authority management: Authority management based on jwt and casbin.

- User management: The system administrator assigns user roles and role permissions.

- Role management: Create the main object of permission control, and then assign different resource permissions and menu permissions to the role.

- Menu management: User dynamic menu configuration implementation, assigning different menus to different roles.

- Resource management: Different users can call different API permissions.

- Setting management: dynamically configure settings (global and current user) for the system.

- Languages management: dynamically configure languages for the system.

- Audit log: system normal operation log record and query; system abnormal information log record and query.

- Login log: The system login log record query contains login exceptions.

- File: Upload & download files from the system.

### Dependent Library

- Backend: using [echo](https://echo.labstack.com/) to quickly build basic RESTful API. Echo is high performance, extensible, minimalist Go web framework.

- DB: [PostgreSQL](https://www.postgresql.org/)，using [gorm](https://gorm.io/) to implement data manipulation, added support for MSSQL, MySQL databases.

- Cache: using [Redis](https://redis.io/) to implement the recording of the JWT token of the currently active user and implement the multi-login restriction.

- API: using Swagger to auto generate APIs docs (based on [swaggo](github.com/swaggo/swag/)

- Job Scheduling: [gocron](https://github.com/jasonlvhit/gocron)

- Log: using [zap](https://github.com/uber-go/zap) record logs

- WebSocket: [websocket](https://github.com/gorilla/websocket)

### Running locally

```go
> docker-compose up

```
