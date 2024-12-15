from fastapi import FastAPI
from routers import router as main_router

app = FastAPI(
    title="formula-processor",
    description="Сервис обрабатывает изображения и формулы",
    version="0.0.9"
)

app.include_router(main_router)

# if __name__ == "__main__":
#     import uvicorn
#     uvicorn.run(app, host="0.0.0.0", port=8000, log_level="info")
