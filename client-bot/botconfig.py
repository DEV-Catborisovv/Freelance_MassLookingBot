import tomllib

def parseTomlLib():
    with open("./configs/config.toml", "rb") as f:
        data = tomllib.load(f)
    return data