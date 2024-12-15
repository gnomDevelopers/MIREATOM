from fastapi import APIRouter, Body, Query

from schemas import FormulasRequest
from services import process_formulas, process_image

router = APIRouter()

@router.get("/image/process", response_model=str, tags=["image"])
def handle_image(image_path: str = Query(..., description="Путь к изображению")):
    result = process_image(image_path)
    return result

@router.post("/formulas/process", tags=["formulas"])
def handle_formulas(data: FormulasRequest = Body(...)):
    result = process_formulas(data.input_formula, data.array_formulas)
    return result
