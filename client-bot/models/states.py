from aiogram.fsm.context import FSMContext
from aiogram.fsm.state import StatesGroup, State

class GivingDataForMakeRequest(StatesGroup):
    GiveApiId = State()
    GiveApiHash = State()
    GiveChats = State()
    GivePhoneNumber = State()
    GiveCode = State()