package main

// import (
// 	"fmt"
// )

// func main() {
// 	db := infrastructure.NewGormConnect()

// 	userRepository := repository.NewUserRepository(db)
// 	userUsecase := usecase.NewUserUsecase(userRepository)
// 	userController := controller.NewUserController(userUsecase)

// 	grpcServer := infrastructure.NewGrpcServer()
// 	user.RegisterUserServiceServer(grpcServer, userController)

// 	listener, err := net.Listen("tcp", ":"+os.Getenv("GRPC_SERVICE_PORT")) // [::]:50051
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v\n", err)
// 		return
// 	}
// 	log.Print("grpcServer serve")
// 	if err := grpcServer.Serve(listener); err != nil {
// 		log.Fatal("serve err", err)
// 	}
// }

// var maxInt int = 2

// type Stack struct {
// 	data []int
// }

// func (s *Stack) isEmpty() bool {
// 	return len(s.data) == 0
// }

// func (s *Stack) isFull() bool {
// 	return len(s.data) == maxInt
// }

// func (s *Stack) push(n int) {
// 	if s.isFull() {
// 		fmt.Println("is full")
// 		return
// 	}
// 	s.data = append(s.data, n)
// }

// func (s *Stack) pop() int {
// 	if s.isEmpty() {
// 		fmt.Println("is empty")
// 		return 0
// 	}
// 	last := len(s.data) - 1
// 	popData := s.data[last]
// 	s.data = s.data[:last]
// 	return popData
// }

// type StackInterface interface {
// 	isEmpty() bool
// 	isFull() bool
// 	push(n int)
// 	pop() int
// }

// func main() {
// 	stack := &Stack{}
// 	stack.push(3)
// 	stack.push(5)
// 	stack.push(7)
// 	stack.push(7)
// 	stack.push(7)
// }

// type Pair struct {
// 	begin int
// 	end   int
// }

// type Stack struct {
// 	data []int
// }

// func (s *Stack) isEmpty() bool {
// 	return len(s.data) == 0
// }

// func (s *Stack) push(n int) {
// 	s.data = append(s.data, n)
// }

// func (s *Stack) pop() int {
// 	if s.isEmpty() {
// 		fmt.Println("is empty")
// 		return 0
// 	}
// 	last := len(s.data) - 1
// 	popData := s.data[last]
// 	s.data = s.data[:last]
// 	return popData
// }

// type StackInterface interface {
// 	isEmpty() bool
// 	push(n int)
// 	pop() int
// }

// func check(brackets string) error {
// 	stack := &Stack{}
// 	var pairs []Pair
// 	fmt.Println("---------begin")

// 	for index, str := range brackets {
// 		currentStr := fmt.Sprintf("%c", str)
// 		fmt.Println("currentStr", currentStr)
// 		fmt.Println("index", index)

// 		if fmt.Sprintf("%c", str) == "(" {
// 			stack.push(index) // 1:[0], 2:[0,1]
// 			fmt.Println("stack.data", stack.data)
// 		} else if fmt.Sprintf("%c", str) == ")" {
// 			// } else {
// 			if stack.isEmpty() {
// 				return errors.New("しね")
// 			}
// 			begindex := stack.pop() //ペア成立 1 => 1, 2 => 0
// 			pair := Pair{
// 				begin: begindex,
// 				end:   index,
// 			}
// 			pairs = append(pairs, pair)
// 			// なんでこのコードでペアになるんだろうか？
// 		}
// 	}
// 	fmt.Println("pairs", pairs)
// 	return nil
// }

// func main() {
// 	// check("())()")
// 	check("(1")
// 	// check("()")
// }
