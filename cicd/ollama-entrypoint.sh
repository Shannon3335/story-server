#! bin/bash

# Start a process
/bin/ollama serve &

# Get the Process Id of the most recently executed process
pid=$!

sleep 5

echo "Retrieving model..."
ollama pull llama3.3

wait $pid
