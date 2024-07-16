import botconfig as botConfig

from aiogram import Bot, Dispatcher
from aiogram.enums.parse_mode import ParseMode
from aiogram.fsm.storage.memory import MemoryStorage

# Parse config.toml
config = botConfig.parseTomlLib()

# Init objects to bot
bot = Bot(token=config['telegram']['token'], parse_mode=ParseMode.HTML)
dp = Dispatcher(storage=MemoryStorage())