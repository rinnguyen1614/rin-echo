# Rin-echo

### Directories

- [deploy](deploy/) deployment
- [internal](internal/) application code
- [scripts](scripts/) development scripts
- <del>[terraform](terraform/) - infrastructure definition</del>
- [web](web/) - frontend reatjs code

### Features

<ul>
   
<li> Dynamic paramaters including sorting, filtering, selecting and pagination in RESTful API</li>

<li>Authority management: Authority management based on jwt and casbin.</li>

<li>User management: The system administrator assigns user roles and role permissions.</li>

<li>Role management: Create the main object of permission control, and then assign different resource permissions and menu permissions to the role.</li>

<li>Menu management: User dynamic menu configuration implementation, assigning different menus to different roles.</li>

<li>Resource management: Different users can call different API permissions.</li>

<li>Setting management: dynamically configure settings (global and current user) for the system.</li>

<li>Languages management: dynamically configure languages for the system.</li>

<li>Audit log: system normal operation log record and query; system abnormal information log record and query.</li>

<li>Login log: The system login log record query contains login exceptions.</li>

<li>File: Upload & download files from the system</li>

</ul>

### Technical selection

<ul>

<li>Frontend: using [react-admin](https://marmelab.com/react-admin//) based on [React](https://reactjs.org/)，to code the page.</li>

<li>Backend: using [echo](https://echo.labstack.com/) to quickly build basic RESTful API. Echo is high performance, extensible, minimalist Go web framework.</li>

<li>DB: [PostgreSQL](https://www.postgresql.org/)，using [gorm](https://gorm.io/) to implement data manipulation, added support for MSSQL, MySQL databases.</li>

<li>Cache: using [Redis](https://redis.io/) to implement the recording of the JWT token of the currently active user and implement the multi-login restriction.</li>

<li>API: using Swagger to auto generate APIs docs  (based on [swaggo](github.com/swaggo/swag/)</li>

<li>Config: using fsnotify and viper to implement yaml config file。</li>

<li>Job Scheduling: [gocron](https://github.com/jasonlvhit/gocron)</li>

<li>Log: using [zap](https://github.com/uber-go/zap) record logs</li>

<li>WebSocket: [websocket](https://github.com/gorilla/websocket)</li>

</ul>

### Running locally

```go
> docker-compose up

```
