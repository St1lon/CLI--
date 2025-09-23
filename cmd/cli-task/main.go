package main
import(
	//"cli-track/internal/domain"
	"cli-track/internal/application/services"
)

func main(){
	taskManager := services.NewTaskManager()
	taskManager.AddTask("kdmkdakda", "to-do")
	taskManager.PrintTasks()


}