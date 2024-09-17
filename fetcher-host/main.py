from fastapi import FastAPI, Request
import json

app = FastAPI()

@app.post("/")
async def index(request: Request):

    data_in = await request.json()

    # Считать данные из файла JSON
    with open("data.json", "r") as f:
        data = json.load(f)
    
    res = []
    for pair in data_in:
        search = pair['artist'] + '-' + pair['track']
        if stat := data.get(search, None):
            res.append(stat)
            print(res)

    # Отправить данные в теле ответа
    return res


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="localhost", port=1337)