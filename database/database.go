package database

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

func GetTasks(backet string) ([]string, error) {

	var err error
	db, err = bolt.Open("tasks.db", 0600, &bolt.Options{Timeout: 3 * time.Second})
	if err != nil {
		return nil, err
	}
	defer db.Close()

	tasks := []string{}
	db.View(func(tx *bolt.Tx) error {

		b, err := tx.CreateBucketIfNotExists([]byte(backet))
		if err != nil {
			fmt.Println("Could not create bucket", backet)
			return nil
		}

		c := b.Cursor()

		index := 1

		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, string(v))
			index++
		}

		if len(tasks) <= 0 {
			return nil
		}

		if backet == Uncompleted {
			fmt.Println("Uncompleted tasks:")
		} else if backet == Completed {
			fmt.Println("Completed tasks:")
		}
		return nil
	})

	return tasks, nil
}

func AddTask(task string) error {
	var err error
	db, err = bolt.Open("tasks.db", 0600, &bolt.Options{Timeout: 3 * time.Second})
	if err != nil {
		return err
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte(Uncompleted))
		id, _ := b.NextSequence()
		return b.Put(itob(id), []byte(task))
	})
	return nil
}

func DoTasks(indexes []int) ([]string, error) {
	var err error
	db, err = bolt.Open("tasks.db", 0600, &bolt.Options{Timeout: 3 * time.Second})
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var tasksToMark []Task
	var marksSuccessful []string

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(Uncompleted))
		c := b.Cursor()
		var counter = 0
		for k, v := c.First(); k != nil; k, v = c.Next() {
			counter++
			for _, inx := range indexes {
				if counter == inx {
					taskToMark := Task{
						id:          k,
						description: v,
					}
					tasksToMark = append(tasksToMark, taskToMark)
				}
			}
		}
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte(Completed))
		b1, _ := tx.CreateBucketIfNotExists([]byte(Uncompleted))
		for _, val := range tasksToMark {
			err := b.Put(val.id, val.description)
			if err != nil {
				fmt.Println("Unable to save", string(val.description))
			}
			err1 := b1.Delete(val.id)
			if err1 != nil {
				fmt.Println("Unable to delete", string(val.description))
			}
			if err == nil && err1 == nil {
				marksSuccessful = append(marksSuccessful, string(val.description))
			}

		}

		return err
	})
	return marksSuccessful, nil
}
