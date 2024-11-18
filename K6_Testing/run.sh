#!/bin/bash

run_join() {
  echo "Running join_Test.js..."
  k6 run join_Test.js
}

run_leave() {
  echo "Running leave_Test.js..."
  k6 run leave_Test.js
}

run_messages() {
  echo "Running messages_Test.js..."
  k6 run messages_Test.js
}

run_send() {
  echo "Running send_Test.js..."
  k6 run send_Test.js
}

run_all() {
  echo "Running all tests..."
  k6 run join_Test.js &
  k6 run leave_Test.js &
  k6 run messages_Test.js &
  k6 run send_Test.js &
  wait
}

case "$1" in
  "join")
    run_join
    ;;
  "leave")
    run_leave
    ;;
  "messages")
    run_messages
    ;;
  "send")
    run_send
    ;;
  "all")
    run_all
    ;;
  *)
    echo "Usage: $0 {join|leave|messages|send|all}"
    exit 1
    ;;
esac
