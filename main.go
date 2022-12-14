package main

import (
	"fmt"
	pb "go_proto/pb"
	"google.golang.org/protobuf/proto"
	"log"
)

func main() {
	products := &pb.Products{
		Pagination: &pb.Pagination{
			Total:       100,
			PerPage:     10,
			CurrentPage: 1,
			LastPage:    10,
		},
		Data: []*pb.Product{
			{
				Id:    1,
				Name:  "Mike Black T-Shirt",
				Price: 10000,
				Stock: 10,
				ProductCategory: &pb.ProductCategory{
					Id:   1,
					Name: "Shirt",
				},
			}, {
				Id:    2,
				Name:  "Mike Air Jordan",
				Price: 10000,
				Stock: 10,
				ProductCategory: &pb.ProductCategory{
					Id:   2,
					Name: "Shoe",
				},
			},
		},
	}

	// encoding
	// compact binary wire format
	data, err := proto.Marshal(products)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Data hasil encode : %v", data)

	//	decoding
	testProduct := &pb.Products{}

	if err = proto.Unmarshal(data, testProduct); err != nil {
		log.Fatal(err)
	}
	log.Printf("Data hasil dencode : %v", testProduct)
	fmt.Println("")

	for _, product := range testProduct.Data {
		log.Println(product.GetId())
		log.Println(product.GetName())
		log.Println(product.GetProductCategory().GetName())
		fmt.Println("")
	}
}
