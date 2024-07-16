from aiogram import types, F
from misc import dp, bot
from aiogram.types import Message, CallbackQuery, FSInputFile
from aiogram.filters import Command, StateFilter

from aiogram.fsm.context import FSMContext
from aiogram.fsm.state import StatesGroup, State

from models.states import GivingDataForMakeRequest

@dp.message(F.text, Command("start"))
async def StartCommandHandler(msg: Message, state: FSMContext):
    if msg.from_user.username is None:
        await bot.send_message(chat_id=msg.from_user.id, text="❌ Вашему аккаунту не присвоен username. Установите его и используйте кнопку /start после этого")
        return
    
    await bot.send_message(chat_id=msg.from_user.id, text="👋 Здравствуйте, приступим к созданию задачи для просмотра историй.\n\nСначала передайте <b>API_ID</b>, инструкция по получению API_ID и API_HASH: https://clck.ru/3BwV4B\n\n❗️ <b>Обратите внимание: в силу технической структуры приложения данные для авторизации не проверяются (проверяйте данные сами и вводите их внимательно)</b>")
    await state.set_state(GivingDataForMakeRequest.GiveApiId)