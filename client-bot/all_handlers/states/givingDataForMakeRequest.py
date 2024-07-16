from aiogram import types, F
from misc import dp, bot
from aiogram.types import Message, CallbackQuery, FSInputFile
from aiogram.filters import Command, StateFilter

from aiogram.fsm.context import FSMContext
from aiogram.fsm.state import StatesGroup, State

from models.states import GivingDataForMakeRequest

from middlewares.phonechecker.phonechecker import is_valid_phone_number
from middlewares.stringformater.stringformeter import format_string

import requests

@dp.message(GivingDataForMakeRequest.GiveApiId)
async def handleGivingApiId(msg: Message, state: FSMContext):
    await state.update_data(uAPI_ID=msg.text)

    await bot.send_message(chat_id=msg.from_user.id, text="✏️ Отлично! Пожалуйста, теперь отправьте <b>API_HASH</b>")
    await state.set_state(state=GivingDataForMakeRequest.GiveApiHash)

@dp.message(GivingDataForMakeRequest.GiveApiHash)
async def handleGivingApiHash(msg: Message, state: FSMContext):
    await state.update_data(uAPI_HASH=msg.text)

    await bot.send_message(chat_id=msg.from_user.id, text="✏️ Отлично! Пожалуйста, отправьте список чатов через запятую.\n\nПример: Чат города екатеринбург, SlapChat, AllsDams")
    await state.set_state(state=GivingDataForMakeRequest.GiveChats)

@dp.message(GivingDataForMakeRequest.GiveChats)
async def handleGivingChats(msg: Message, state: FSMContext):
    list, ok = format_string(msg.text)
    if ok == False:
        await bot.send_message(chat_id=msg.from_user.id, text="Пожалуйста, отправьте корректный формат список")
        return 

    await bot.send_message(chat_id=msg.from_user.id, text="✏️ Отлично! Пожалуйста, теперь отправьте номер телефона.\n\n<b>Обратите внимание, что требуется соблюсти формат: +(код региона страны)(номер)</b>")
    await state.update_data(chats=list)
    await state.set_state(state=GivingDataForMakeRequest.GivePhoneNumber)
    pass

@dp.message(GivingDataForMakeRequest.GivePhoneNumber)
async def handleGivingPhonenumber(msg: Message, state: FSMContext):
    if not is_valid_phone_number(msg.text):
        await bot.send_message(chat_id=msg.from_user.id, text="Пожалуйста, отправьте корректный формат номера телефона")
        return

    data = await state.get_data()

    dataOfRequest = {
        "api_id": data['uAPI_ID'],
        "api_hash": data["uAPI_HASH"],
        "phone_number": msg.text,
        "chats": data["chats"]
    }

    print(dataOfRequest)

    headers = {
        'Content-Type': 'application/json',
    }
    
    try:
        response = requests.put("http://localhost:8081/api/add_task", json=dataOfRequest, headers=headers)
        response.raise_for_status()
    except requests.exceptions.RequestException as e:
        await bot.send_message(chat_id=msg.from_user.id, text="❌ Не удалось выполнить запрос, попробуйте использовать другие данные или повторите попытку позже (начать сначала - /start)")
        return
    
    if response.status_code == 102:
        await state.update_data(phone_number=msg.text)
        await bot.send_message(chat_id=msg.from_user.id, text="❗️ Данные были отправлены на сервер.\n\n✏️ <b>Сейчас Вам должен прийти код аутентификации, пожалуйста, отправьте его</b>")
        await state.set_state(GivingDataForMakeRequest.GiveCode)
    else:
        await bot.send_message(chat_id=msg.from_user.id, text="❌ Не удалось выполнить запрос, попробуйте использовать другие данные или повторите попытку позже (начать сначала - /start)")
        await state.set_state(state=None)
        return

@dp.message(GivingDataForMakeRequest.GiveCode)
async def handleGivingCode(msg: Message, state: FSMContext):
    try:
        code = int(msg.text)
    except:
        await bot.send_message(chat_id=msg.from_user.id, text="❌ Пожалуйста, отправьте корректный код")
        return

    data = await state.get_data()

    dataOfRequest = {
        "phone_number": data["phone_number"],
        "code": str(code)
    }

    headers = {
        'Content-Type': 'application/json',
    }
    
    try:
        response = requests.put("http://localhost:8081/api/verify", json=dataOfRequest, headers=headers)
        response.raise_for_status()

    except requests.exceptions.RequestException as e:
        await bot.send_message(chat_id=msg.from_user.id, text="❌ Не удалось выполнить запрос, попробуйте использовать другие данные или повторите попытку позже (начать сначала - /start)")
        return

    if response.status_code == 200:
        await bot.send_message(chat_id=msg.from_user.id, text="❇️ Задача была создана, ожидайте окончания её выполнения (после завершения задачи на заданный аккаунт придет сообщение о завершении).\n\nЧтобы создать новую задачу используйте команду /start")
        await state.set_state(state=None)
    else:
        await bot.send_message(chat_id=msg.from_user.id, text="❌ Не удалось выполнить запрос, попробуйте использовать другие данные или повторите попытку позже (начать сначала - /start)")
        await state.set_state(state=None)
        return