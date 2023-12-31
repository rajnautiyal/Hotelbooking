#!/bin/bash

# Assuming 'go' is in your PATH. If not, provide the full path to the 'go' executable.
go build -o bookings cmd/web/*.go && ./bookings