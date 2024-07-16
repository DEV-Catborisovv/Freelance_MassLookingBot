import logging
import json
import requests
from aiogram import Bot, Dispatcher, types
from aiogram.client.default import DefaultBotProperties
from aiogram.enums import ParseMode
from aiogram.filters import Command
from aiogram.types import Message

import asyncio
import logging
import sys
from os import getenv

API_TOKEN = '7063971479:AAHdLf9JUYnNTtUA8mcesRSKPrx6lSw1TRQ'
BACKEND_URL = 'http://your_backend_url/api/add_task'

logging.basicConfig(level=logging.INFO)

dp = Dispatcher()

# Хранение данных пользователей
user_data = {}

@dp.message(Command("start"))
async def send_welcome(message: Message):
    await message.reply("Привет! Отправь мне API_ID.")

@dp.message(lambda message: message.text.isdigit())
async def process_api_id(message: Message):
    user_id = message.from_user.id
    user_data[user_id] = {'api_id': message.text}
    await message.reply("Теперь отправь мне API_HASH.")

@dp.message(lambda message: not message.text.isdigit())
async def process_api_hash(message: Message):
    user_id = message.from_user.id
    if 'api_id' in user_data.get(user_id, {}):
        user_data[user_id]['api_hash'] = message.text
        await message.reply("Теперь отправь мне список чатов через запятую.")
    else:
        await message.reply("Пожалуйста, сначала отправьте API_ID.")

@dp.message(lambda message: ',' in message.text)
async def process_chats(message: Message):
    user_id = message.from_user.id
    if 'api_id' in user_data.get(user_id, {}) and 'api_hash' in user_data[user_id]:
        chats = [chat.strip() for chat in message.text.split(',')]
        user_data[user_id]['chats'] = chats

        api_id = user_data[user_id]['api_id']
        api_hash = user_data[user_id]['api_hash']

        payload = {
            "api_id": api_id,
            "api_hash": api_hash,
            "chats": chats
        }

        headers = {
            'Content-Type': 'application/json'
        }

        response = requests.post(BACKEND_URL, json=payload, headers=headers)

        if response.status_code == 200:
            await message.reply("Задача создана и будет выполнена в ближайшее время.")
        else:
            await message.reply(f"Ошибка при создании задачи: {response.text}")
    else:
        await message.reply("Пожалуйста, сначала отправьте API_ID и API_HASH.")

async def main() -> None:
    # Initialize Bot instance with default bot properties which will be passed to all API calls
    bot = Bot(token=API_TOKEN, default=DefaultBotProperties(parse_mode=ParseMode.HTML))

    # And the run events dispatching
    await dp.start_polling(bot)

if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO, stream=sys.stdout)
    asyncio.run(main())