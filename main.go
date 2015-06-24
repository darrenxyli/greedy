package main

import
// "fmt"
(
	"runtime"

	"github.com/darrenxyli/greedy/processor"
)

// "sync"

// "github.com/darrenxyli/greedy/database/postgre"
// dedis "github.com/darrenxyli/greedy/database/redis"

// _ "github.com/lib/pq"

// 生成连接池

// var redisClient = dedis.NewClient("192.80.146.5", "6379")

// func put(wg *sync.WaitGroup, value string) {
// 	for i := 1; i < 10; i++ {
// 		redisClient.Put("test", value)
// 	}
// 	wg.Done()
// }

// func get(wg *sync.WaitGroup) {

// 	for i := 1; i < 10; i++ {
// 		value, error := redisClient.Get("test")

// 		if error != nil {
// 			i--
// 		} else {
// 			fmt.Println(value)
// 		}
// 	}
// 	wg.Done()
// }

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	processor.Clear()
	// web.Run()
	// var wg sync.WaitGroup
	//
	// wg.Add(1)
	// go put(&wg, "run")
	//
	// time.Sleep(time.Second)
	//
	// len, _ := redisClient.QueueSize("test")
	//
	// fmt.Println(len)
	//
	// wg.Add(1)
	// go get(&wg)
	//
	// wg.Wait()

	//= This initiates a form session to the database. see Open

	// taskDB := postgre.NewTaskDB(
	// 	"amazon.cbtwp3cmfmsx.us-west-2.rds.amazonaws.com",
	// 	5432,
	// 	"taskdb",
	// 	"darrenxyli",
	// 	"2jaqx97j",
	// 	[]string{"test"})

	// test connection with
	// MyGorm.SingularTable(true)

	// MyGorm.CreateTable(&task.Task{})

	// taskItem := task.NewTask()
	//
	// taskDB.Insert(taskItem)
	// //
	// // MyGorm.NewRecord(taskItem)
	// //
	// // MyGorm.Create(&taskItem)
	//
	// var user = new(task.Task)
	//
	// taskDB.Get("test", user)
	// fmt.Println(user.Project)

	// taskDB.LoadTasks("ACTIVE", "test", 10)

}
