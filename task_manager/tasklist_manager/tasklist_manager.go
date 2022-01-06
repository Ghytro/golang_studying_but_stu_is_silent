package tasklistmanager

import (
	"errors"
	"gotutorial/task_manager/pb"
	"io/fs"
	"os"

	"google.golang.org/protobuf/proto"
)

var taskListFilename = "tasklist.pb.bin"

type Task struct {
	Name     string
	Deadline string
}

func getTaskList() (*pb.TaskList, error) {
	taskListBytes, err := os.ReadFile(taskListFilename)
	if err != nil {
		return nil, err
	}
	taskList := new(pb.TaskList)
	if err := proto.Unmarshal(taskListBytes, taskList); err != nil {
		return nil, err
	}
	return taskList, nil
}

func writeTaskList(taskList *pb.TaskList) error {
	if marshaledProto, err := proto.Marshal(taskList); err != nil {
		return err
	} else {
		os.WriteFile(taskListFilename, marshaledProto, 0644)
		return nil
	}
}

func AddTask(taskName string, deadline string) error {
	taskList, err := getTaskList()
	if err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			return err
		}
		taskList = &pb.TaskList{}
	}
	taskList.Tasks = append(taskList.Tasks, &pb.Task{
		Name:     taskName,
		Deadline: deadline,
	})
	return writeTaskList(taskList)
}

func getTaskIndexByName(tasks []*pb.Task, name string) int {
	for i, t := range tasks {
		if t.Name == name {
			return i
		}
	}
	return -1
}

func DoTaskByName(taskName string) error {
	taskList, err := getTaskList()
	if err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			return err
		}
		taskList = &pb.TaskList{}
	}
	taskIndex := getTaskIndexByName(taskList.Tasks, taskName)
	if taskIndex == -1 {
		return errors.New("Cannot find task " + taskName + " in task list")
	}
	taskList.Tasks = append(taskList.Tasks[:taskIndex], taskList.Tasks[taskIndex+1:]...)
	return writeTaskList(taskList)
}

func DoTaskByIndex(taskIndex int) error {
	taskList, err := getTaskList()
	if err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			return err
		}
		taskList = &pb.TaskList{}
	}
	taskList.Tasks = append(taskList.Tasks[:taskIndex], taskList.Tasks[taskIndex+1:]...)
	return writeTaskList(taskList)
}

func GetTaskList() ([]Task, error) {
	taskList, err := getTaskList()
	if err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			return nil, err
		}
		taskList = &pb.TaskList{}
	}
	result := make([]Task, len(taskList.Tasks))
	for i, t := range taskList.Tasks {
		result[i] = Task{t.Name, t.Deadline}
	}
	return result, nil
}
