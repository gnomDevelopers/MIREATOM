import hashlib
from collections import defaultdict
from difflib import SequenceMatcher

from ollama import ChatResponse, chat
from sympy import simplify, sympify


def process_image(image_path: str) -> str:
    response: ChatResponse = chat(
        model='llama3.2-vision',
        messages=[
            {
                'role': 'user',
                'content': 'Give me ONLY LaTeX formula from photo. Be very strict and careful!!!',
                'images': [image_path]
            },
        ],
    )
    if response.message.content is not None:
        return response.message.content.strip()
    else:
        return ""

def process_formulas(input_formula: str, formulas: list[str]) -> dict:
    # Проводим текстовый и семантический анализ
    all_formulas = [input_formula] + formulas
    semantic_groups = group_semantically_equivalent_formulas(all_formulas)

    threshold = 4
    matches = []
    total_matched_length = 0

    for f in formulas:
        common_ranges = find_common_substrings(input_formula, f, threshold)
        # common_ranges - список кортежей (start, end, length)
        for (start, end, length) in common_ranges:
            matches.append(f"{start}-{end}")
            total_matched_length += length

    percent = 0.0
    if len(input_formula) > 0:
        percent = (total_matched_length / len(input_formula)) * 100.0

    match_str = ";".join(matches)

    return {
        "origin": input_formula,
        "percent": percent,
        "match": match_str
    }

def hash_formula(formula):
    return hashlib.md5(formula.encode("utf-8")).hexdigest()

def are_semantically_equivalent(latex1, latex2):
    try:
        expr1 = sympify(latex1)
        expr2 = sympify(latex2)
        return simplify(expr1 - expr2) == 0
    except Exception:
        return False

def find_common_substrings(s1, s2, threshold=4):
    matcher = SequenceMatcher(None, s1, s2)
    blocks = matcher.get_matching_blocks()
    results = []
    for block in blocks:
        i, j, n = block
        if n > threshold:
            results.append((i, i + n - 1, n))
    return results

def group_semantically_equivalent_formulas(formulas):
    semantic_groups = defaultdict(list)
    for formula in formulas:
        is_unique = True
        for key_formula, group in semantic_groups.items():
            if are_semantically_equivalent(formula, key_formula):
                group.append(formula)
                is_unique = False
                break
        if is_unique:
            semantic_groups[formula].append(formula)
    return semantic_groups
