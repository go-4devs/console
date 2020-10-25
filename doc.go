// Package console eases the creation of beautiful and testable command line interfaces.
// The Console package allows you to create command-line commands.
// Your console commands can be used for any recurring task, such as cronjobs, imports, or other batch jobs.
// console application can be written as follows:
//   //cmd/console/main.go
//   func main() {
//     console.New().Execute(context.Background())
//   }
// Then, you can register the commands using Add():
//   package main
//
//   import (
//     "context"
//
//     "gitoa.ru/go-4devs/console"
//     "gitoa.ru/go-4devs/console/example/pkg/command"
//   )
//
//   func main() {
//     console.
//       New().
//         Add(
//           command.Hello(),
//           command.Args(),
//           command.Hidden(),
//           command.Namespace(),
//        ).
//        Execute(context.Background())
//   }
package console
