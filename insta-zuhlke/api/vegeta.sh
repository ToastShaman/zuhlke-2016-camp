#!/bin/zsh
echo "GET http://localhost:4140/images\nHost: web" | vegeta attack -duration=30s -rate=120 | tee results.bin | vegeta report
