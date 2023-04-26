package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	companyProto = ProtoConfig{
		Location:    "./pkg/gen/company_api/proto",
		Output:      "./pkg/gen/company_api",
		File:        "xm.proto",
		BuildServer: true,
	}
)

var (
	gen = flag.Bool("gen", false, "regenerate files")
)

type ProtoConfig struct {
	Location    string
	Output      string
	File        string
	BuildServer bool
}

func main() {
	flag.Parse()
	if *gen {
		GenerateProtobuf(companyProto)
	} else {
		flag.PrintDefaults()
	}
}

func GenerateProtobuf(conf ProtoConfig) {
	out, err := generateProtobuf(conf)
	if err != nil {
		log.Printf("Error generating protobuf: %s - %s", err, out)
		return
	}

	log.Printf("Resources proto files generated successfully at %s using proto at %s", conf.Output, conf.Location)
}

func generateProtobuf(conf ProtoConfig) ([]byte, error) {
	if _, err := os.Stat(conf.Location); err != nil {
		return nil, err
	}

	p := strings.Split(conf.Output, "/")
	pluginPath := strings.TrimSuffix(conf.Output, p[len(p)-1])

	cmd := exec.Command("protoc",
		fmt.Sprintf("--proto_path=%s", conf.Location),
		fmt.Sprintf("--go_out=%s", conf.Output),
		fmt.Sprintf("--go-grpc_out=%s", conf.Output), "--go-grpc_opt=paths=source_relative",
		"--go_opt=paths=source_relative")

	if conf.BuildServer {
		cmd.Args = append(cmd.Args,
			fmt.Sprintf("--validate_out=lang=go:%s", pluginPath),
			fmt.Sprintf("--gofullmethods_out=%s", pluginPath))
	}
	cmd.Args = append(cmd.Args, conf.File)

	return cmd.CombinedOutput()
}
