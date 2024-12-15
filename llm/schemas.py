from typing import List

from pydantic import BaseModel


class FormulasRequest(BaseModel):
    input_formula: str
    array_formulas: List[str]
