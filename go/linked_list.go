package main

// import (
// 	"errors"
// 	"fmt"
// )

// // func main() {
// // 	db := infrastructure.NewGormConnect()

// // 	userRepository := repository.NewUserRepository(db)
// // 	userUsecase := usecase.NewUserUsecase(userRepository)
// // 	userController := controller.NewUserController(userUsecase)

// // 	grpcServer := infrastructure.NewGrpcServer()
// // 	user.RegisterUserServiceServer(grpcServer, userController)

// // 	listener, err := net.Listen("tcp", ":"+os.Getenv("GRPC_SERVICE_PORT")) // [::]:50051
// // 	if err != nil {
// // 		log.Fatalf("failed to listen: %v\n", err)
// // 		return
// // 	}
// // 	log.Print("grpcServer serve")
// // 	if err := grpcServer.Serve(listener); err != nil {
// // 		log.Fatal("serve err", err)
// // 	}
// // }

// type Address struct {
// 	name          string
// 	phone_number  string
// 	email_address string
// }

// type Node struct {
// 	addr *Address
// 	next *Node
// }

// type AddressList struct {
// 	root *Node
// 	len  int
// }

// type AddressListInterface interface {
// 	Insert(currentNode *Node, name, phone_number, email_address string) *Node
// 	GetAddress(n int) *Node
// 	Remove(n int, currentNode *Node) (*Node, error)
// 	Index()
// }

// func NewAddressList() AddressListInterface {
// 	return &AddressList{}
// }

// func (a *AddressList) Insert(currentNode *Node, name, phone_number, email_address string) *Node {
// 	address := &Address{
// 		name:          name,
// 		phone_number:  phone_number,
// 		email_address: email_address,
// 	}
// 	node := &Node{addr: address}
// 	if a.len > 0 {
// 		// 1番目以降のNodeを作る
// 		// 生成したnodeのnextに入れることで2個目を表現
// 		node.next = currentNode.next
// 		// rootが渡された時（？）はnextが取れないので、明示的に代入
// 		currentNode.next = node
// 		// 現在のnodeを更新する
// 		currentNode = currentNode.next
// 	} else {
// 		// 0番目のNodeを作る
// 		// rootに突っ込むけどnextがいない。nextはどうやったら分かるんだ？
// 		a.root = node
// 		currentNode = a.root
// 	}

// 	a.len++
// 	fmt.Printf("現在のlen: %v\n", a.len)
// 	return currentNode
// }

// func (a *AddressList) Index() {
// 	fmt.Println("------------------index-------------------")
// 	for i := 1; i <= a.len; i++ {
// 		node := a.GetAddress(i)
// 		fmt.Println(node.addr)
// 	}
// 	fmt.Println("------------------index-------------------")
// }

// func main() {
// 	// Interfaceを持ったstructを生成
// 	addressList := NewAddressList()
// 	// rootNodeはnextが空のstructを持つ。{name: "kanagawa", next: nil}
// 	currentNode := &Node{}
// 	currentNode = addressList.Insert(currentNode, "Kanagawa", "1111", "ohishikaito@gmail.com")
// 	// secondは{name: Tokyo, next: }
// 	currentNode = addressList.Insert(currentNode, "Tokyo", "2222", "kaito990626@yahoo.co.jp")
// 	currentNode = addressList.Insert(currentNode, "World", "3333", "kaito990626@yahoo.co.jp")

// 	// FIXME: 最後のnodeを消すとバグる。理由はnextが参照できないからだと思う！
// 	currentNode, err := addressList.Remove(3, currentNode)
// 	if err != nil {
// 		panic(err)
// 	}
// 	currentNode = addressList.Insert(currentNode, "4D", "4444", "kaito990626@yahoo.co.jp")

// 	// 一覧を表示
// 	addressList.Index()

// 	// 単体を取得
// 	// node := addressList.GetAddress(2)
// 	// fmt.Println(node.next.addr)
// 	// fmt.Println(node.addr)
// }

// func (a *AddressList) GetAddress(n int) *Node {
// 	if n > a.len {
// 		return nil
// 	}

// 	targetNode := a.root
// 	// 条件分からん➡️i++する時の条件を書けばokだった！
// 	for i := 1; i < n; i++ {
// 		targetNode = targetNode.next
// 	}
// 	return targetNode
// }

// func (a *AddressList) Remove(n int, currentNode *Node) (*Node, error) {
// 	// nilエラーをハンドリングしよう。returnするNodeは削除したやつかな？
// 	// なんでcurrentNodeを渡す必要あるんだろーw

// 	removeNode := a.GetAddress(n)
// 	if removeNode == nil {
// 		return nil, errors.New("not found!!!!!!!")
// 	}
// 	if n == 1 {
// 		// 1 == rootなので特別な処理をする
// 		a.root = removeNode.next
// 	} else {
// 		// root以外は入れ替える処理
// 		beforeNode := a.GetAddress(n - 1)
// 		// 最後尾だとnextがないので参照できない。
// 		if n == a.len {
// 			beforeNode.next = nil
// 			currentNode = beforeNode
// 		} else {
// 			// 最後尾じゃないので参照できる
// 			beforeNode.next = removeNode.next // nextがnilの可能性あり
// 		}
// 	}

// 	a.len--
// 	return currentNode, nil
// }

// var AddressList []*Address

// func main() {
// 	var AddressList []Address
// 	AddressList = append(AddressList, Address{
// 		name:          "Kanagawa",
// 		phone_number:  "1111",
// 		email_address: "ohishikaito@gmail.com",
// 	})
// 	AddressList = append(AddressList, Address{
// 		name:          "Tokyo",
// 		phone_number:  "2222",
// 		email_address: "ohishikaito@gmail.com",
// 	})
// 	AddressList = append(AddressList, Address{
// 		name:          "World",
// 		phone_number:  "3333",
// 		email_address: "ohishikaito@gmail.com",
// 	})

// 	remove(AddressList, 2)
// }

// func remove(array []Address, num int) {
// 	fmt.Println("array[:num]", array[:num])
// 	fmt.Println("array[num+1:]...", array[num+1:])
// 	hoge := append(array[:num], array[num+1:]...)
// 	fmt.Println(hoge)
// }
