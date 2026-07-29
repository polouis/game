# game

## Engine development
If you want to develop the engine and test it with your application :
- go to the application directory
- `go work init .`
- `go work use /path/to/local/engine`
- Ignore workspace file : 
  - `echo "go.work" >> .gitignore`
  - `echo "go.work.sum" >> .gitignore`
