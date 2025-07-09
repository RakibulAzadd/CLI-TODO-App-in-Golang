package main

func main(){
	 todos := Todos{}
     storage := NewStorage[Todos]("todos.json")
	 storage.Load(&todos)
     cmdFlag := NewCmdFlags()
	 cmdFlag.Execute(&todos)
	 storage.Save(todos)
}