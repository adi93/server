# server
Golang server for serving vimwiki files and acting as a task manager

The files for serving static files is pretty simple, and can be gleaned by main.go, middleware and config folders.

For the task server, I am using sqlite3 as backend, and a layered architecture with mvc pattern.
