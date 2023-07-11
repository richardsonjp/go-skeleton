#!/bin/sh
migrate create -ext sql -dir cmd/apiserver/app/migrations/ -seq schema
