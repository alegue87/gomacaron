# Grafana Data Source Plugin Template

[![CircleCI](https://circleci.com/gh/grafana/simple-datasource/tree/master.svg?style=svg)](https://circleci.com/gh/grafana/simple-datasource/tree/master)

This template is a starting point for building Grafana Data Source Plugins

## What is Grafana Data Source Plugin?
Grafana supports a wide range of data sources, including Prometheus, MySQL, and even Datadog. There’s a good chance you can already visualize metrics from the systems you have set up. In some cases, though, you already have an in-house metrics solution that you’d like to add to your Grafana dashboards. Grafana Data Source Plugins enables integrating such solutions with Grafana.

## Getting started
1. Install dependencies
```BASH
npm install
```
2. Build plugin in development mode or run in watch mode
```BASH
npm run dev
```
or
```BASH
npm run watch
```
3. Build plugin in production mode
```BASH
npm run build
```

4. Install 

Posizionare il plugin in: moqui-framework/docker/grafana/data/plugins/moqui-websocket-datasource/
Quindi aggiungere il plugin e configurare l'endpoint del websocket in grafana ( es:  ws://localhost:8080/notws )
Indicare il topic sul quale registrarsi ( creato in precedenza in moqui )

Il datasource può essere utilizzato come gauge ( per visualizzare un valore ) o su un grafico di tipo timeseries;
per utilizzarlo come timeseries è necessaria aggiungere una trasformation di tipo 1 "Prepare time series" ed indicando
come formato "Multi-frame time series".

I dati inviati da moqui devono essere [time: ..., value: ... ] es in un servizio con:

     ec.makeNotificationMessage().topic('testws')
        .title('ws')
        .message([value:r.nextInt(10), time: new Date().format("yyyy-MM-dd HH:mm:ss.SSSSSS")])
        .send()
        
Per evitare il refresh dei pannelli con il live stream (timeseries), disabilitare il refresh della dashboard indicando
nella configurazione generale della dashboard:

    Refresh live dashbord -> disabled
    (Continuously re-draw panels where the time range references 'now')
    

## Learn more
- [Build a data source plugin tutorial](https://grafana.com/tutorials/build-a-data-source-plugin)
- [Grafana documentation](https://grafana.com/docs/)
- [Grafana Tutorials](https://grafana.com/tutorials/) - Grafana Tutorials are step-by-step guides that help you make the most of Grafana
- [Grafana UI Library](https://developers.grafana.com/ui) - UI components to help you build interfaces using Grafana Design System
