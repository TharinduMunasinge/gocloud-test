package main

import (
	"fmt"
	"crypto/rand"
	"github.com/scorelab/gocloud/lib/commons/aws"
	"github.com/scorelab/gocloud/lib/commons/aws/ec2"
	"encoding/hex"
	"time"
)

func main(){

	token,_:=clientToken()
	fmt.Println(token)
	auth,_ := aws.GetAuth("","",token,time.Time{});
	ec2I := ec2.New(
		auth,
		aws.Regions["ap-southeast-1"],
	)
	options := ec2.RunInstancesOptions{
		ImageId:      "ami-1ddc0b7e",
		InstanceType: "t2.micro",
		MaxCount: 1,
		MinCount:1,
		AvailabilityZone:"ap-southeast-1",
		KeyName: "goclouds-test",
	}
	res,er:=ec2I.RunInstances(&options)
	if(er != nil)	{
		fmt.Println("Error Occured");
	}else{
		fmt.Println(res)
	}

}
func clientToken() (string, error) {
	// Maximum EC2 client token size is 64 bytes.
	buf := make([]byte, 32)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(buf), nil
}
