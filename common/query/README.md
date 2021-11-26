## Send a request

```
GET /users

?sorts=created_at, username:desc // sort by created_at, then descendingly by username
&filters=username|full_name in system,systemistrator | full_name in system, systemistrator
&page=1
&page=2
```

More formally:
- sorts is a comma-delimited ordered list of {Field}:{Order} wher
  - {Field} is the name of a propery for TEntity
  - {Order} (optional) is sorted ASC | DESC. If {Order} is empty, it is implicitly ASC
- selects
- filters
- page is the number of page to return
- page_size is the number of items returned per page
