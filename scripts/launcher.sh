#! /bin/bash

OUTPUT_DIR="output"
INPUT_FILE=""
COMMAND=""

ui-graphics-tools "$@"

POSITIONAL_ARGS=()

while [[ $# -gt 0 ]]; do
  case $1 in
    --output-dir)
      OUTPUT_DIR="$2"
      shift # past argument
      shift # past value
      ;;
    --config-file)
      INPUT_FILE="$2"
      shift # past argument
      shift # past value
      ;;
    *)
      POSITIONAL_ARGS+=("$1") # save positional arg
      shift # past argument
      ;;
  esac
done

if [ ${POSITIONAL_ARGS[0]} == "icons" ];then
    INPUT_FILE=${POSITIONAL_ARGS[1]}
fi

OWNER_ID=$(stat -c %u $INPUT_FILE)
GROUP_ID=$(stat -c %g $INPUT_FILE)
chown -R $OWNER_ID:$GROUP_ID $OUTPUT_DIR
