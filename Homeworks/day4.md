**Задания**

1. Реализовать handler: `GetAccount`
```
URL: /account/1
use: ctx.ShouldBindUri
```
2. Реализовать handler: `ListAccounts`
```
URL: /accounts/?page_id=10&page_size=5
use: ctx.ShouldBindQuery
```
Body -> Query parameters  
\`json:\` -> \`form:\`