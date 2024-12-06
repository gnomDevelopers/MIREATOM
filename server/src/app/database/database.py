import asyncpg
import os

pool = None

async def init_database():
    global pool
    pool = await asyncpg.create_pool(
        host="localhost",
        port="5432",
        user="postgres",
        password="12345678",
        database="template"
    )
    async with pool.acquire() as conn:
        await conn.execute("""
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            username VARCHAR,
            first_name VARCHAR,
            second_name VARCHAR,
            photo_url TEXT,
            authorized_at TIMESTAMP
        );
        """)
        print("Таблица 'users' создана успешно")

async def get_connection() -> asyncpg.Connection:
    return pool.acquire()

async def close() -> None:
    if pool:
        await pool.close()
