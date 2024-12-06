from fastapi import FastAPI
from contextlib import asynccontextmanager
from database.database import init_database, close

app = FastAPI()

@asynccontextmanager
async def lifespan(app: FastAPI):
    # Инициализация базы данных
    await init_database()
    yield
    # Закрытие соединения с базой данных
    await close()

app = FastAPI(lifespan=lifespan)