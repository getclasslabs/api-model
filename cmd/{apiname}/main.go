package main

import (
	"github.com/getclasslabs/{apiname}/internal"
	"github.com/getclasslabs/{apiname}/internal/config"
	"github.com/getclasslabs/{apiname}/internal/repository"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerConf "github.com/uber/jaeger-client-go/config"
	jaegerLog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"os"
)

const name = "{apiname}"

func main() {

	f, err := os.Open("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config.Config)
	if err != nil {
		panic(err)
	}

	repository.Start()

	cfg := jaegerConf.Configuration{
		ServiceName: name,
		Sampler: &jaegerConf.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegerConf.ReporterConfig{
			LogSpans:           false,
			LocalAgentHostPort: config.Config.Jaeger.Host + ":" + config.Config.Jaeger.Port,
		},
	}

	jLogger := jaegerLog.StdLogger
	jMetricsFactory := metrics.NullFactory

	tracer, closer, err := cfg.NewTracer(
		jaegerConf.Logger(jLogger),
		jaegerConf.Metrics(jMetricsFactory),
	)

	if err != nil {
		log.Fatal("failed to initialize tracer")
	}
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	s := internal.NewServer()
	log.Println("waiting routes...")
	log.Fatal(http.ListenAndServe(config.Config.Server.Port, s.Router))
}
