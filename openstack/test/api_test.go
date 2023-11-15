package main

import (
	"fmt"
	"log"

	"github.com/gophercloud/gophercloud"
	//"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

func main() {
	// 认证信息
	//auth := aksk.NewCredentials("YOUR_ACCESS_KEY", "YOUR_SECRET_KEY", "YOUR_PROJECT_ID", "YOUR_DOMAIN_ID")
	auth := gophercloud.AuthOptions{
		IdentityEndpoint: "https://openstack.example.com:5000/v2.0",
		Username:         "{username}",
		Password:         "{password}",
		TenantID:         "{tenant_id}",
	}
	// 创建 OpenStack 客户端
	provider, err := openstack.AuthenticatedClient(auth)
	if err != nil {
		log.Fatalf("Failed to create OpenStack client: %v", err)
	}
	fmt.Println(provider)
	// 创建 Compute 客户端
	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		log.Fatalf("Failed to create Compute client: %v", err)
	}

	// 创建虚拟机
	createOpts := servers.CreateOpts{
		Name:       "my-instance",
		ImageName:  "Ubuntu 20.04",
		FlavorName: "m1.small",
		Networks: []servers.Network{
			{
				UUID: "YOUR_NETWORK_UUID",
			},
		},
	}
	server, err := servers.Create(client, createOpts).Extract()
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	fmt.Printf("Server created: ID=%s, Name=%s\n", server.ID, server.Name)

	// 停止虚拟机
	//err = servers.WaitForStatus(client, server.ID).ExtractErr()
	if err != nil {
		log.Fatalf("Failed to stop server: %v", err)
	}

	fmt.Println("Server stopped")

	// 调整虚拟机规格
	resizeOpts := servers.ResizeOpts{
		FlavorRef: "NEW_FLAVOR_ID",
	}
	err = servers.Resize(client, server.ID, resizeOpts).ExtractErr()
	if err != nil {
		log.Fatalf("Failed to resize server: %v", err)
	}

	fmt.Println("Server resized")

	// 删除虚拟机
	err = servers.Delete(client, server.ID).ExtractErr()
	if err != nil {
		log.Fatalf("Failed to delete server: %v", err)
	}

	fmt.Println("Server deleted")
}
