package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
	"github.com/gophercloud/gophercloud/pagination"
)

func main() {

	// Option 1: Pass in the values yourself
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: "http://192.168.1.13",
		Username:         "admin",
		Password:         "whoelse105105",
	}
	provider, err := openstack.AuthenticatedClient(opts)
	fmt.Println(provider)
	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	//client.HTTPClient.Get()
	fmt.Println(client)
	imageListOpts := images.ListOpts{}

	/*// 创建计算服务客户端
	computeClient, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("创建计算服务客户端失败:", err)
		return
	}
	*/

	imageListPager := images.List(client, imageListOpts)
	err = imageListPager.EachPage(func(page pagination.Page) (bool, error) {
		imageList, err := images.ExtractImages(page)
		if err != nil {
			fmt.Println("提取镜像列表失败:", err)
			return false, err
		}

		for _, image := range imageList {
			fmt.Printf("镜像名称: %s, ID: %s\n", image.Name, image.ID)
		}

		return true, nil
	})
	if err != nil {
		fmt.Println("查询镜像列表失败:", err)
		return
	}
}
