#!/bin/sh

curl --location 'localhost:8080/fis/stress' \
--header 'Content-Type: application/json' \
--data '{
    
    "fis":{
        "sec":60,
        "experimentName":"stress-chaos-testing"
    },
    "chaosSpec":{
        "duration": "30s",
        "containerNames":["demo"],
        "selector":{
            "namespaces": ["default"],
            "labelSelectors":{
                "app":"demo"
            }
        },
        "stressors":{
        "cpu":{
            "workers": 1,
            "load": 100
            }
        },
        "mode":"all"
    }
}'