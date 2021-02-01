package main

import (
	"context"
	"fmt"

	dbinternal "github.com/OchiengEd/afya/cmd/initdb/internal"
	"github.com/OchiengEd/afya/pkg/platform"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	fmt.Println("Hello world!")

	client, err := platform.Client()
	if err != nil {
		fmt.Printf("Error getting k8s client: %v\n", err)
	}

	pod, err := client.CoreV1().Pods("default").Get(context.Background(), "postgresql-59f98db469-ptqjj", metav1.GetOptions{})
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Pod: %+v\n", pod)

	dbinternal.InitDB()
}
