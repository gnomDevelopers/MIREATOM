import asyncpg
import os
import logging

pool = None

async def init_database():
    global pool
    try:
        pool = await asyncpg.create_pool(
            host="localhost",
            port="5432",
            user="postgres",
            password="12345678",
            database="template"
        )
        async with pool.acquire() as conn:
            await create_tables(conn)

    except asyncpg.exceptions.PostgresError as e:
        print(f"Ошибка работы с PostgreSQL: {e}")
        if pool:
            await pool.close()
        raise
    except Exception as e:
        print(f"Непредвиденная ошибка: {e}")
        if pool:
            await pool.close()
        raise

async def get_connection() -> asyncpg.Connection:
    return pool.acquire()

async def close() -> None:
    if pool:
        await pool.close()


async def create_tables(conn):
    """
    Создание таблиц в базе данных.
    """
    try:
        # Таблица пользователей
        await conn.execute("""
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            name VARCHAR NOT NULL,
            surname VARCHAR NOT NULL,
            third_name VARCHAR,
            role VARCHAR NOT NULL,
            email VARCHAR NOT NULL UNIQUE,
            password VARCHAR NOT NULL
        );
        """)
        print("Таблица 'users' создана успешно.")

        # Таблица статей
        await conn.execute("""
        CREATE TABLE IF NOT EXISTS articles (
            id SERIAL PRIMARY KEY,
            user_id INT NOT NULL,
            path VARCHAR NOT NULL,
            originality VARCHAR[] NOT NULL,
            FOREIGN KEY (user_id) REFERENCES users (id)
        );
        """)
        print("Таблица 'articles' создана успешно.")

        # Таблица управления версиями статей
        await conn.execute("""
        CREATE TABLE IF NOT EXISTS article_vcs (
            id SERIAL PRIMARY KEY,
            article_id INT NOT NULL,
            difference VARCHAR NOT NULL,
            hash VARCHAR NOT NULL,
            FOREIGN KEY (article_id) REFERENCES articles (id)
        );
        """)
        print("Таблица 'article_vcs' создана успешно.")

        # Таблица истории
        await conn.execute("""
        CREATE TABLE IF NOT EXISTS history (
            id SERIAL PRIMARY KEY,
            formula VARCHAR NOT NULL,
            user_id INT NOT NULL,
            FOREIGN KEY (user_id) REFERENCES users (id)
        );
        """)
        print("Таблица 'history' создана успешно.")

        # Таблица формул
        await conn.execute("""
        CREATE TABLE IF NOT EXISTS formula (
            id SERIAL PRIMARY KEY,
            value VARCHAR NOT NULL,
            user_id INT NOT NULL,
            FOREIGN KEY (user_id) REFERENCES users (id)
        );
        """)
        print("Таблица 'formula' создана успешно.")

    except asyncpg.exceptions.PostgresError as e:
        print(f"Ошибка при создании таблицы: {e}")
        raise
    except Exception as e:
        print(f"Непредвиденная ошибка при создании таблиц: {e}")
        raise