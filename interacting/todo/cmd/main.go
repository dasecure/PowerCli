package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"ooi.cc/interacting/todo"
)

var todoFileName = ".todo.json"

func main() {
	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed for DaSecure\n", os.Args[0])
		fmt.Fprintln(flag.CommandLine.Output(), "Copyright 2022")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()

	}
	add := flag.Bool("add", false, "Task to be added in the TODO list")
	list := flag.Bool("list", false, "List all task")
	del := flag.Int("del", 0, "Delete task")
	complete := flag.Int("complete", 0, "Item to be completed")
	verbose := flag.Bool("verbose", false, "Display verbose output")
	noshow := flag.Bool("noshow", false, "Do not show completed")

	flag.Parse()

	l := &todo.List{}
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *noshow:
		for i, item := range *l {
			if !item.Done {
				fmt.Printf("%d: %s \n", i+1, item.Task)
			}
		}
	case *verbose:
		for i, item := range *l {
			if !item.Done {
				fmt.Printf("%d: %s Created: %s\n", i+1, item.Task, item.CreatedAt)
			} else {
				fmt.Printf("%d: %s Completed: %s Created: %s\n", i+1, item.Task, item.CompletedAt, item.CreatedAt)
			}

		}
	case *del > 0:
		if err := l.Delete(*del); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *list:
		fmt.Print(l)
	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *add:
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		l.Add(t)
		// l.Add(*task)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}
	if len(s.Text()) == 0 {
		return "", fmt.Errorf("Task cannot be blank")
	}
	return s.Text(), nil
}
