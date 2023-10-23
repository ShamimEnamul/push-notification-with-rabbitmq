Not finished yet......

### RabbitMQ
it accepts, stores, and forwards binary blobs of data ‒ messages.A message buffer that RabbitMQ keeps on behalf of the consumer.
If you want to check on the queue, try using rabbitmqctl list_queues.






First, install amqp using go get:
```shell
go get github.com/rabbitmq/amqp091-go
```

### docker
```shell
docker run --name rabbitmq -p 5672:5672 rabbitmq // add -d to run in background
```
here, 5672 is default port, locally we also define the same port. you can replace with yours



### Reference

- https://rohinivsenthil.medium.com/rabbitmq-in-golang-getting-started-34c65e6c7f92
- https://www.rabbitmq.com/tutorials/tutorial-one-go.html
- https://pkg.go.dev/github.com/rabbitmq/amqp091-go 
- https://www.rabbitmq.com/configure.html#config-items (Core Server Variables Configurable in rabbitmq.conf)
