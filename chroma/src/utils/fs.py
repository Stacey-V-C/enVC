import os
import json

def get_config():
  directory = os.getenv("NVC_CONFIG_DIR")

  if directory is None:
      directory = "../.."

  config = open(directory + "/nvc.config.json", "r")

  config = json.load(config)

  return config