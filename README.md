# reunion
a chat app for small groups, private, secure, something, something

## Sample execution stack
Here is an example of the execution stack for a sample task. Register a user:

```bash
main.go
   server.go
      router.go
         handler.go
            user.go
               datastore.go
                  leveldb.go
```

## Adapter pattern
The application makes use of the adapter pattern for external dependencies. This allows extensibility and abstracts away vendor details from the core application.

For example, the `datastore` interface can be used by any persistence store.