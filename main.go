package main

import (
	"aws-golang/src/infraestructure"
)


func main() {
	infraestructure.S3.Region="us-east-1"
	infraestructure.S3.NewSession(infraestructure.S3.Region)
	infraestructure.S3.Ls()

	// infraestructure.S3.Upload("subeme.txt","golandia-julian","subido.txt")
	// infraestructure.S3.GenerateUrl("golandia","subido.txt")
}