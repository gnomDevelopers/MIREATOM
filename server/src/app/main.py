from fastapi import FastAPI
from contextlib import asynccontextmanager
from server.src.app.database.database import init_database, close, get_connection

@asynccontextmanager
async def lifespan(app: FastAPI):
    # Инициализация базы данных
    await init_database()
    yield
    # Закрытие соединения с базой данных
    await close()

app = FastAPI(lifespan=lifespan)

@app.get("/")
async def root():
    return {"message": "Hello World"}