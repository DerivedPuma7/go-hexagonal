### Explanation
This adapter allows to isolate the main part, the application core (entities, rules, etc), from external vendors. It makes the core protected from external world.
From this point foward, if necessary, is easy to create other database implementations, like mysql, dynnamoDb, etc. Just follow the contract (interface) and all should work fine.

Instead of product.go, maybe we could call it with a more specific name, like: sqlite.go, product.sqlite.go. Thats because this implementation is unic from sqlite. But lets keep it that way, since we only have one database implementation.

Of course, its not very common to change an application database, but it could happen at some point. The main ideia is to understand the concept behind the decoupling. Once this idea is clear, you can see the bigger picture, witch means to isolate not only the database, but every external vendor from your application.