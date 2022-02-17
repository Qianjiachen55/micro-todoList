module task

go 1.16

require (
	github.com/golang/protobuf v1.5.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/rabbitmq/amqp091-go v1.3.0
	github.com/spf13/viper v1.10.1
	google.golang.org/protobuf v1.27.1
	gorm.io/driver/mysql v1.2.3
	gorm.io/gorm v1.22.5
)

replace google.golang.org/grpc => google.golang.org/grpc v1.19.0
