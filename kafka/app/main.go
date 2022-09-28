package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"strings"
	"net/http"

	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
//	brokers             = []string{"kafka-1:9092"}
	brokers	 	     []string
	topic   goka.Stream = "monitoring"
	group   goka.Group  = "monitoring-group"
)

var readCounter = prometheus.NewCounter(
   prometheus.CounterOpts{
       Name: "kafka_monitoring_read_request_count",
       Help: "No of request handled by Read handler",
   },
)

var readProcessorError = prometheus.NewCounter(
   prometheus.CounterOpts{
       Name: "kafka_monitoring_read_processor_error",
       Help: "No of request handled by Read handler",
   },
)

var writeCounter = prometheus.NewCounter(
   prometheus.CounterOpts{
       Name: "kafka_monitoring_write_request_count",
       Help: "No of request handled by Write handler",
   },
)

var writeErrorCounter = prometheus.NewCounter(
   prometheus.CounterOpts{
       Name: "kafka_monitoring_write_error_count",
       Help: "No of request handled by Write handler",
   },
)

// Emit messages forever every second
func runEmitter() {
	emitter, err := goka.NewEmitter(brokers, topic, new(codec.String))
	if err != nil {
		log.Fatalf("error creating emitter: %v", err)
	}
	defer emitter.Finish()
	for {
		time.Sleep(1 * time.Second)
		err = emitter.EmitSync("some-key", "some-value")
		writeCounter.Inc()
		if err != nil {
			writeErrorCounter.Inc()
			log.Printf("error emitting message: %v", err)
		}
	}
}

// process messages until ctrl-c is pressed
func runProcessor() {
	// process callback is invoked for each message delivered from
	// "example-stream" topic.
	cb := func(ctx goka.Context, msg interface{}) {
		var counter int64
		// ctx.Value() gets from the group table the value that is stored for
		// the message's key.
		if val := ctx.Value(); val != nil {
			counter = val.(int64)
		}
		//counter++
		// SetValue stores the incremented counter in the group table for in
		// the message's key.
		ctx.SetValue(counter)
		log.Printf("key = %s, counter = %v, msg = %v", ctx.Key(), counter, msg)
		readCounter.Inc()
	}

	// Define a new processor group. The group defines all inputs, outputs, and
	// serialization formats. The group-table topic is "example-group-table".
	g := goka.DefineGroup(group,
		goka.Input(topic, new(codec.String), cb),
		goka.Persist(new(codec.Int64)),
	)

	p, err := goka.NewProcessor(brokers, g)
	if err != nil {
		log.Fatalf("error creating processor: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan bool)
	go func() {
		defer close(done)
		if err = p.Run(ctx); err != nil {
			readProcessorError.Inc()
			log.Printf("error running processor: %v", err)
		} else {
			log.Printf("Processor shutdown cleanly")
		}
	}()

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)
	<-wait   // wait for SIGINT/SIGTERM
	cancel() // gracefully stop processor
	<-done
}

func main() {
	brokers = strings.Split(os.Getenv("BROKERS"), ",")
	topic = goka.Stream(os.Getenv("TOPIC"))
	group = goka.Group(os.Getenv("GROUP"))

	prometheus.MustRegister(writeCounter)
	prometheus.MustRegister(writeErrorCounter)
	prometheus.MustRegister(readCounter)
	prometheus.MustRegister(readProcessorError)

	go runEmitter() // emits one message and stops
	go runProcessor()  // press ctrl-c to stop
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8090", nil)
}
